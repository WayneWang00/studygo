package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// 确保test.txt文件存在
	testPath := filepath.Join("..", "..", "test.txt")
	if _, err := os.Stat(testPath); os.IsNotExist(err) {
		fmt.Println("test.txt文件不存在，请在project根目录下创建")
		return
	}

	// 创建一个buffer存储表单数据
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 创建文件字段
	file, err := os.Open(testPath)
	if err != nil {
		fmt.Println("打开文件错误:", err)
		return
	}
	defer file.Close()

	part, err := writer.CreateFormFile("file", "test.txt")
	if err != nil {
		fmt.Println("创建表单文件错误:", err)
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println("复制文件内容错误:", err)
		return
	}

	// 添加project字段
	err = writer.WriteField("project", "test")
	if err != nil {
		fmt.Println("添加project字段错误:", err)
		return
	}

	// 关闭writer
	err = writer.Close()
	if err != nil {
		fmt.Println("关闭writer错误:", err)
		return
	}

	// 创建HTTP请求
	req, err := http.NewRequest("POST", "http://localhost:8888/upload", body)
	if err != nil {
		fmt.Println("创建请求错误:", err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求错误:", err)
		return
	}
	defer resp.Body.Close()

	// 打印响应
	fmt.Println("响应状态码:", resp.Status)
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Println("错误响应内容:", string(bodyBytes))
	}
}
