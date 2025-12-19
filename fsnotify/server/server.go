package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var projects = make(map[string]string)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10MB
		http.Error(w, "Error parsing the form: "+err.Error(), http.StatusBadRequest)
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Printf("获取表单文件错误: %v", err)
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	fmt.Println("Uploaded file:", handler.Filename)

	// 打印整个请求的Form数据
	fmt.Println("Form data:")
	for key, values := range r.Form {
		fmt.Printf("%s: %v\n", key, values)
	}

	// 打印请求头
	fmt.Println("Headers:")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}

	project := r.FormValue("project")
	fmt.Println("Project:", project)
	if project == "" {
		http.Error(w, "Project name is required", http.StatusBadRequest)
		return
	}

	fileRelPath := r.FormValue("relPath")
	fmt.Println("File relPath:", fileRelPath)
	if fileRelPath == "" {
		http.Error(w, "File relPath is required", http.StatusBadRequest)
		return
	}

	savePath := filepath.Join("D:/test_upload", project)
	fmt.Println("Save path:", savePath)
	if err := os.MkdirAll(savePath, os.ModePerm); err != nil {
		log.Printf("创建项目目录失败: %v", err)
		http.Error(w, "Unable to create directory", http.StatusInternalServerError)
		return
	}
	log.Printf("保存项目目录: %s", savePath)
	// 确保多级目录存在
	fullPath := filepath.Join(savePath, fileRelPath)
	if err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
		log.Printf("创建项目文件目录失败: %v", err)
		http.Error(w, "Unable to create directories", http.StatusInternalServerError)
		return
	}

	dst, err := os.Create(fullPath)
	if err != nil {
		http.Error(w, "Unable to save the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	projects[project+"/"+fileRelPath] = time.Now().String()
	fmt.Fprintf(w, "File %s uploaded successfully", fileRelPath)
}

func projectInfoHandler(w http.ResponseWriter, r *http.Request) {
	project := mux.Vars(r)["project"]
	files := make([]string, 0)
	for k := range projects {
		if strings.HasPrefix(k, project+"/") {
			files = append(files, k)
		}
	}
	fmt.Fprintf(w, "Files in project %s: %v", project, files)
}

type ProjectConfig struct {
	ClientID     string            `json:"clientId"`
	CodebaseName string            `json:"codebaseName"`
	CodebasePath string            `json:"codebasePath"`
	CodebaseId   string            `json:"codebaseId"`
	HashTree     map[string]string `json:"hashTree"`
	LastSync     time.Time         `json:"lastSync"`
}

type CodebaseHashResp struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    CodebaseHashRespData `json:"data"`
}

type CodebaseHashRespData struct {
	List []HashItem `json:"list"`
}

type HashItem struct {
	Path string `json:"path"`
	Hash string `json:"hash"`
}

func CodebaseHashHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 解析请求参数
	clientId := r.URL.Query().Get("clientId")
	codebasePath := r.URL.Query().Get("codebasePath")
	fmt.Println("clientId:", clientId, " codebasePath:", codebasePath)

	filePath := "D:/test_upload/sync_config.json"

	// 如果是文件，返回文件内容
	content, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read file: %v", err), http.StatusInternalServerError)
		return
	}

	var config = new(ProjectConfig)
	err = json.Unmarshal(content, config)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to unmarshal json: %v", err), http.StatusInternalServerError)
		return
	}

	data := make([]HashItem, 0)
	for path, hash := range config.HashTree {
		data = append(data, HashItem{Hash: hash, Path: path})
	}

	resp, err := json.Marshal(CodebaseHashResp{Code: 0, Data: CodebaseHashRespData{List: data}})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal json: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func uploadZipHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10MB
		http.Error(w, "Error parsing the form: "+err.Error(), http.StatusBadRequest)
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Printf("获取表单文件错误: %v", err)
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	fmt.Println("Uploaded file:", handler.Filename)

	// 打印整个请求的Form数据
	fmt.Println("Form data:")
	for key, values := range r.Form {
		fmt.Printf("%s: %v\n", key, values)
	}

	// 打印请求头
	fmt.Println("Headers:")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}

	savePath := filepath.Join("D:/test_upload", handler.Filename)
	fmt.Println("Save path:", savePath)
	dst, err := os.Create(savePath)
	if err != nil {
		http.Error(w, "Unable to save the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	projects[handler.Filename] = time.Now().String()
	fmt.Fprintf(w, "File %s uploaded successfully", handler.Filename)
}

func clientConfigHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request for client config", time.Now())
	configPath := "D:/test_upload/config.json"

	// 检查文件是否存在
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		http.Error(w, fmt.Sprintf("config file not found: %v", err), http.StatusNotFound)
		return
	}

	// 读取文件内容
	content, err := os.ReadFile(configPath)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read config file: %v", err), http.StatusInternalServerError)
		return
	}

	// 直接写入原始JSON内容，避免双重编码
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

func clientConfigByVersionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request for client config", time.Now())
	configPath := "D:/test_upload/0.1.8/config.json"

	// 检查文件是否存在
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		http.Error(w, fmt.Sprintf("config file not found: %v", err), http.StatusNotFound)
		return
	}

	// 读取文件内容
	content, err := os.ReadFile(configPath)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read config file: %v", err), http.StatusInternalServerError)
		return
	}

	// 直接写入原始JSON内容，避免双重编码
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

type TokenRequest struct {
	ClientId     string `json:"clientId"`
	CodebasePath string `json:"codebasePath"`
}

type TokenResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    TokenResponseData `json:"data"`
}

type TokenResponseData struct {
	UploadToken string `json:"uploadToken"`
	UserCount   int    `json:"userCount"`
}

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received token request ", time.Now())
	var tokenReq TokenRequest
	if err := json.NewDecoder(r.Body).Decode(&tokenReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("tokenReq:", tokenReq)

	token := fmt.Sprintf("%s_%s", tokenReq.ClientId, tokenReq.CodebasePath)
	response := TokenResponse{Code: 200, Message: "success", Data: TokenResponseData{UploadToken: token}}
	resp, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal json: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

type FileStatusReq struct {
	ClientId     string `json:"clientId"`
	CodebasePath string `json:"codebasePath"`
	CodebaseName string `json:"codebaseName"`
	SyncId       string `json:"syncId"`
}

// FileStatusResp 文件状态响应
type FileStatusResp struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    FileStatusRespData `json:"data"`
}

type FileStatusRespData struct {
	Process       string                       `json:"process"`
	TotalProcess  int                          `json:"totalProcess"`
	FileSuceesNum int                          `json:"fileSuceesNum"`
	FileList      []FileStatusRespFileListItem `json:"fileList"`
}

type FileStatusRespFileListItem struct {
	Path    string `json:"path"`
	Operate string `json:"operate"`
	Status  string `json:"status"`
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received file status request ", time.Now())
	var fileStatusReq FileStatusReq
	if err := json.NewDecoder(r.Body).Decode(&fileStatusReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("fileStatusReq:", fileStatusReq)
	response := FileStatusResp{Code: 200, Message: "success", Data: FileStatusRespData{
		Process:       "complete",
		TotalProcess:  0,
		FileSuceesNum: 0,
		FileList:      []FileStatusRespFileListItem{},
	}}
	resp, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshal json: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/upload", uploadHandler).Methods("POST")
	r.HandleFunc("/project/{project}", projectInfoHandler).Methods("GET")
	r.HandleFunc("/codebase-embedder/api/v1/codebases/hash", CodebaseHashHandler).Methods("GET")
	r.HandleFunc("/codebase-embedder/api/v1/files/upload", uploadZipHandler).Methods("POST")
	r.HandleFunc("/codebase-embedder/api/v1/files/token", tokenHandler).Methods("POST")
	r.HandleFunc("/codebase-embedder/api/v1/files/status", statusHandler).Methods("POST")
	r.HandleFunc("/shenma/api/v1/config/codebase-indexer-config.json", clientConfigHandler).Methods("GET")
	r.HandleFunc("/shenma/api/v1/config/0.1.8/codebase-indexer-config.json", clientConfigByVersionHandler).Methods("GET")

	log.Println("Starting server at :8888")
	log.Fatal(http.ListenAndServe(":8888", r))
}
