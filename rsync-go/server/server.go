package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/minio/rsync-go"
)

type SyncRequest struct {
	Path string `json:"path"`
}

type FileListRequest struct {
	Path string `json:"path"`
}

type FileListResponse struct {
	Files []string `json:"files"`
}

func handleSync(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var syncReq SyncRequest
		if err := json.NewDecoder(r.Body).Decode(&syncReq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		path := syncReq.Path
		if fileInfo, err := os.Stat(path); err == nil {
			if fileInfo.IsDir() {
				handleDirectorySync(w, r, path)
			} else {
				handleFileSync(w, r, path)
			}
			return
		}
		http.Error(w, "路径不存在", http.StatusNotFound)
	}
}

func handleFileSync(w http.ResponseWriter, r *http.Request, path string) {
	// 同步文件逻辑
}

func handleDirectorySync(w http.ResponseWriter, r *http.Request, dirPath string) {
	// 生成目录的签名
	sig, err := generateDirectorySignature(dirPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(sig)
}

func generateDirectorySignature(dirPath string) ([]byte, error) {
	var sig []byte
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fileSig, err := generateFileSignature(path)
		if err != nil {
			return err
		}
		sig = append(sig, fileSig...)
		return nil
	})
	return sig, err
}

func generateFileSignature(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	rs := &rsync.RSync{}
	fileSig := make([]byte, 0)

	writeSignature := func(bl rsync.BlockHash) error {
		sigData, err := json.Marshal(bl)
		if err != nil {
			return err
		}
		fileSig = append(fileSig, sigData...)
		return nil
	}

	rs.CreateSignature(file, writeSignature)
	return fileSig, nil
}

func handlePatch(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		path := r.URL.Query().Get("path")
		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 判断是文件还是目录
		if fileInfo, err := os.Stat(path); err == nil && fileInfo.IsDir() {
			handleDirectoryPatch(w, r, path)
			return
		}

		// 处理文件差异
		decoder := rsync.NewOperationDecoder(r.Body)
		if err := applyDelta(path, decoder); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func handleDirectoryPatch(w http.ResponseWriter, r *http.Request, dirPath string) {
	// 处理目录差异逻辑
}

func applyDelta(filePath string, decoder *rsync.OperationDecoder) error {
	// 获取旧文件和新文件
	oldFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer oldFile.Close()

	newFile, err := os.Create(filePath + ".new")
	if err != nil {
		return err
	}
	defer newFile.Close()

	rs := &rsync.RSync{}
	return rs.ApplyDelta(oldFile, newFile, decoder)
}

func handleFileList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var fileListReq FileListRequest
		if err := json.NewDecoder(r.Body).Decode(&fileListReq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		path := fileListReq.Path
		files, err := getFileList(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := FileListResponse{Files: files}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func getFileList(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		relPath, err := filepath.Rel(path, p)
		if err != nil {
			return err
		}
		files = append(files, relPath)
		return nil
	})
	return files, err
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var deleteReq map[string]string
		if err := json.NewDecoder(r.Body).Decode(&deleteReq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		path := deleteReq["path"]
		if err := deletePath(path); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func deletePath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("path does not exist")
	}

	if err := os.RemoveAll(path); err != nil {
		return err
	}

	return nil
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		path := r.Header.Get("X-File-Path")
		if path == "" {
			http.Error(w, "Path is required", http.StatusBadRequest)
			return
		}

		// 创建目录
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 保存文件
		file, err := os.Create(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		_, err = io.Copy(file, r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	http.HandleFunc("/sync", handleSync)
	http.HandleFunc("/patch", handlePatch)
	http.HandleFunc("/filelist", handleFileList)
	http.HandleFunc("/delete", handleDelete)
	http.HandleFunc("/upload", handleUpload)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
