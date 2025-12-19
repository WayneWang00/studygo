package main

import (
	"bufio"
	"bytes"

	// "context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	// "time"

	"github.com/fsnotify/fsnotify"
	// "golang.org/x/time/rate"
)

// var (
// 	limit = rate.NewLimiter(rate.Every(100*time.Millisecond), 1<<20) // 1MB/s
// )

func loadGitIgnoreRules(root string) ([]string, error) {
	gitignore := filepath.Join(root, ".gitignore")
	if _, err := os.Stat(gitignore); os.IsNotExist(err) {
		return nil, nil
	}
	file, err := os.Open(gitignore)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rules []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		rules = append(rules, line)
	}
	return rules, scanner.Err()
}

func shouldIgnore(path string, rules []string) bool {
	for _, rule := range rules {
		matched, _ := filepath.Match(rule, path)
		if matched {
			return true
		}
	}
	return false
}

func uploadFile(path string, rel string, project string) error {
	fmt.Println("Uploading:", path, project)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(path))
	io.Copy(part, file)
	writer.WriteField("project", project)
	writer.WriteField("relPath", rel)
	writer.Close()

	req, err := http.NewRequest("POST", "http://localhost:8888/upload", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// limit.Wait(context.Background()) // 控制上传速率
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("server Response:", resp.StatusCode, resp.Status)
	return nil
}

func syncDir(root string, project string) error {
	rules, _ := loadGitIgnoreRules(root)
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		rel, _ := filepath.Rel(root, path)
		fmt.Println("rel path:", rel)
		if shouldIgnore(rel, rules) {
			return nil
		}
		return uploadFile(path, rel, project)
	})
}

func watchDir(root string, project string) {
	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			watcher.Add(path)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
				rel, _ := filepath.Rel(root, event.Name)
				fileInfo, err := os.Stat(event.Name)
				if err != nil {
					log.Println("stat error:", err)
					continue
				}
				if fileInfo.IsDir() {
					watcher.Add(event.Name)
					log.Println("Added directory to watch:", event.Name)
				} else {
					log.Println("File changed:", event.Name)
					uploadFile(event.Name, rel, project)
				}
			}
			if event.Op&fsnotify.Remove == fsnotify.Remove {
				fileInfo, err := os.Stat(event.Name)
				if err != nil {
					log.Println("stat error (file may be already removed):", err)
					fileInfo = nil
				}

				if fileInfo == nil || fileInfo.IsDir() {
					// 文件已删除或目录
					watcher.Remove(event.Name)
					log.Println("Remove watch:", event.Name)

					// 记录目录下所有文件
					if err := filepath.Walk(event.Name, func(path string, info os.FileInfo, err error) error {
						if err != nil {
							return err
						}
						if !info.IsDir() {
							log.Println("File removed/deleted:", path)
						}
						return nil
					}); err != nil {
						log.Println("walk error:", err)
					}
				} else {
					// 是文件
					log.Println("File removed:", event.Name)
				}
			}
			if event.Op&fsnotify.Rename == fsnotify.Rename {
				str := event.String()
				log.Println("Rename event:", str)
				parts := strings.Split(strings.TrimSpace(str), " ← ")
				if len(parts) != 2 {
					log.Println("invalid rename event format:", str)
					continue
				}
				oldPath := strings.TrimSpace(parts[1])
				newPath := strings.TrimSpace(event.Name)

				// 确保旧路径存在
				if _, err := os.Stat(oldPath); os.IsNotExist(err) {
					log.Println("old path does not exist:", oldPath)
					watcher.Remove(oldPath)
				}

				fmt.Printf("Rename event: old path: %s, new path: %s\n", oldPath, newPath)

				// 确保新路径存在后再添加监听
				if info, err := os.Stat(newPath); err == nil {
					// 如果是目录，
					if info.IsDir() {
						if err := watcher.Add(newPath); err != nil {
							log.Println("error adding new path:", newPath, "error:", err)
						}
					}
				}
			}
		case err := <-watcher.Errors:
			log.Println("error:", err)
		}
	}
}

func main() {
	project := "personal"
	root := "D:\\Code\\Go\\personal"
	fmt.Println("Starting full sync...")
	err := syncDir(root, project)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Starting real-time sync...")
	watchDir(root, project)
}
