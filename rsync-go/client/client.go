package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
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

func main() {
	client := &http.Client{}
	localPath := "example_dir" // 示例本地目录路径
	serverURL := "http://localhost:8080"

	// 获取本地文件列表
	localFiles, err := getFileList(localPath)
	if err != nil {
		fmt.Println("Error getting local file list:", err)
		return
	}

	// 获取远程文件列表
	remoteFiles, err := getRemoteFileList(client, serverURL, localPath)
	if err != nil {
		fmt.Println("Error getting remote file list:", err)
		return
	}

	// 识别新增和修改的文件
	filesToAddOrUpdate := diffFileLists(localFiles, remoteFiles)

	// 识别删除的文件
	filesToDelete := diffFileLists(remoteFiles, localFiles)

	// 同步新增和修改的文件
	for _, file := range filesToAddOrUpdate {
		fullPath := filepath.Join(localPath, file)
		if fileInfo, err := os.Stat(fullPath); err == nil {
			if fileInfo.IsDir() {
				// 同步目录
				if err := syncDirectory(client, serverURL, fullPath); err != nil {
					fmt.Println("Error syncing directory:", err)
				}
			} else {
				// 同步文件
				if err := syncFile(client, serverURL, fullPath); err != nil {
					fmt.Println("Error syncing file:", err)
				}
			}
		}
	}

	// 删除远程已删除的文件
	for _, file := range filesToDelete {
		fullPath := filepath.Join(localPath, file)
		if err := deleteRemoteFile(client, serverURL, fullPath); err != nil {
			fmt.Println("Error deleting file:", err)
		}
	}

	fmt.Println("Sync completed successfully")
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

func getRemoteFileList(client *http.Client, serverURL string, path string) ([]string, error) {
	url := serverURL + "/filelist"
	reqBody, _ := json.Marshal(FileListRequest{Path: path})
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response FileListResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response.Files, nil
}

func diffFileLists(localFiles, remoteFiles []string) []string {
	// 创建远程文件的映射
	remoteMap := make(map[string]bool)
	for _, file := range remoteFiles {
		remoteMap[file] = true
	}

	// 找出本地存在但远程不存在的文件
	var diff []string
	for _, file := range localFiles {
		if !remoteMap[file] {
			diff = append(diff, file)
		}
	}

	return diff
}

func syncDirectory(client *http.Client, serverURL string, dirPath string) error {
	url := serverURL + "/sync"
	reqBody, _ := json.Marshal(SyncRequest{Path: dirPath})
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		sig, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		// 创建差异操作
		opsOut := make(chan rsync.Operation)
		go func() {
			defer close(opsOut)
			if err := walkDirectoryAndGenerateDeltas(dirPath, sig, opsOut); err != nil {
				fmt.Println("Error generating deltas:", err)
			}
		}()

		// 上传差异操作
		patchUrl := url + "?path=" + dirPath
		patchReq, _ := http.NewRequest(http.MethodPost, patchUrl, nil)
		patchReq.Header.Set("Content-Type", "application/octet-stream")

		// 使用多部分请求发送差异操作
		writer := multipart.NewWriter(patchReq.Body)
		part, _ := writer.CreateFormFile("file", "delta.dat")
		encoder := rsync.NewOperationEncoder(part)
		for op := range opsOut {
			err := encoder.Encode(op)
			if err != nil {
				fmt.Println("Error encoding operation:", err)
				return err
			}
		}
		writer.Close()

		patchReq.Header.Set("Content-Type", writer.FormDataContentType())

		_, err = client.Do(patchReq)
		if err != nil {
			return err
		}

		fmt.Println("Directory patched successfully")
	} else {
		// 目录不存在或遇到其他问题，进行全量上传
		if err := uploadDirectory(client, dirPath, url); err != nil {
			return err
		}

		fmt.Println("Directory uploaded successfully")
	}
	return nil
}

func syncFile(client *http.Client, serverURL string, filePath string) error {
	url := serverURL + "/sync"
	reqBody, _ := json.Marshal(SyncRequest{Path: filePath})
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// 文件存在，进行增量上传
		sig, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		// 创建差异操作
		opsOut := make(chan rsync.Operation)
		go func() {
			defer close(opsOut)
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer file.Close()

			rs := &rsync.RSync{}
			rs.CreateDelta(file, opsOut, bytes.NewReader(sig))
		}()

		// 上传差异操作
		patchUrl := url + "?path=" + filePath
		patchReq, _ := http.NewRequest(http.MethodPost, patchUrl, nil)
		patchReq.Header.Set("Content-Type", "application/octet-stream")

		// 使用多部分请求发送差异操作
		writer := multipart.NewWriter(patchReq.Body)
		part, _ := writer.CreateFormFile("file", "delta.dat")
		encoder := rsync.NewOperationEncoder(part)
		for op := range opsOut {
			err := encoder.Encode(op)
			if err != nil {
				fmt.Println("Error encoding operation:", err)
				return err
			}
		}
		writer.Close()

		patchReq.Header.Set("Content-Type", writer.FormDataContentType())

		_, err = client.Do(patchReq)
		if err != nil {
			return err
		}

		fmt.Println("File patched successfully")
	} else {
		// 文件不存在，进行全量上传
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		uploadUrl := serverURL + "/upload"
		uploadReq, _ := http.NewRequest(http.MethodPost, uploadUrl, file)
		uploadReq.Header.Set("Content-Type", "application/octet-stream")
		uploadReq.Header.Set("X-File-Path", filePath)

		_, err = client.Do(uploadReq)
		if err != nil {
			return err
		}

		fmt.Println("File uploaded successfully")
	}
	return nil
}

func walkDirectoryAndGenerateDeltas(dirPath string, sig []byte, opsOut chan<- rsync.Operation) error {
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
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
		if bytes.Contains(sig, fileSig) {
			// 文件已存在，进行增量上传
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			rs := &rsync.RSync{}
			err = rs.CreateDelta(file, opsOut, bytes.NewReader(sig))
			if err != nil {
				return err
			}
		} else {
			// 文件不存在，进行全量上传
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(io.Discard, file)
			if err != nil {
				return err
			}
			// 发送全量文件逻辑
		}
		return nil
	})
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

func uploadDirectory(client *http.Client, dirPath string, url string) error {
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		relativePath, err := filepath.Rel(dirPath, path)
		if err != nil {
			return err
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		uploadUrl := url + "?path=" + filepath.Join(dirPath, relativePath)
		uploadReq, _ := http.NewRequest(http.MethodPost, uploadUrl, file)
		uploadReq.Header.Set("Content-Type", "application/octet-stream")
		uploadReq.Header.Set("X-File-Name", info.Name())

		_, err = client.Do(uploadReq)
		if err != nil {
			return err
		}

		return nil
	})
}

func deleteRemoteFile(client *http.Client, serverURL string, filePath string) error {
	deleteUrl := serverURL + "/delete"
	reqBody, _ := json.Marshal(map[string]string{"path": filePath})
	req, _ := http.NewRequest(http.MethodPost, deleteUrl, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error deleting file: %s", string(body))
	}

	return nil
}
