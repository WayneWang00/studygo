package main

import (
	"context"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/gokrazy/rsync/rsyncclient"
	ignore "github.com/sabhiram/go-gitignore"
)

func main() {
	projectDir := "/test"
	ignore, err := ignore.CompileIgnoreFile(filepath.Join(projectDir, ".gitignore"))
	if err != nil {
		log.Fatal(err)
	}

	// 收集需要同步的文件/目录
	var paths []string
	filepath.Walk(projectDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		rel, _ := filepath.Rel(projectDir, path)
		if !ignore.MatchesPath(rel) {
			paths = append(paths, rel)
		}
		return nil
	})

	// 创建 rsync 客户端
	client, err := rsyncclient.New([]string{"-avz"}, rsyncclient.WithSender())
	if err != nil {
		log.Fatal(err)
	}

	// 连接服务端
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 执行同步
	_, err = client.Run(context.Background(), conn, paths)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("同步完成")
}
