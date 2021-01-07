package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

var (
	TimeFormat = "20060102"
	BasePath   = "./file"
)

func main() {
	_, err := os.Stat(BasePath)
	if os.IsNotExist(err) {
		fmt.Println("not exist")
		return
	}
	srcName := srcInit()
	fmt.Printf("src name:%+v\n", srcName)
	fmt.Println("name:", srcName.Name())
	copyFile(srcName)
}

func srcInit() *os.File {
	now := time.Now()
	fileName := fmt.Sprintf("src.%s", now.Format(TimeFormat))
	filePath := filepath.Join(BasePath, fileName)
	var lastFile *os.File
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("src not exist")
		lastFile, err = os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("openFile failed:", err)
			return nil
		}
	} else {
		lastFile, err = os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("src open failed:", err)
			return nil
		}
		return lastFile
	}
	defer lastFile.Close()

	_, err = lastFile.WriteString(now.Format(TimeFormat) + "input\n")
	if err != nil {
		fmt.Println("file write failed:", err)
		return nil
	}
	return lastFile
}

func copyFile(src *os.File) {
	now := time.Now()
	fileName := fmt.Sprintf("dst.%s", now.Format(TimeFormat))
	sourceFileStat, err := os.Stat(src.Name())
	if err != nil {
		fmt.Println("src stat failed:", err)
		return
	}
	if !sourceFileStat.Mode().IsRegular() {
		fmt.Printf("%s not found", src.Name())
		return
	}
	sourceFile, err := os.Open(src.Name())
	if err != nil {
		fmt.Println("src open failed:", err)
		return
	}
	defer sourceFile.Close()

	var dstFile *os.File
	filePath := filepath.Join(BasePath, fileName)
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("dst stat failed:", err)
		dstFile, err = os.Create(filePath)
		if err != nil {
			fmt.Println("dst create failed:", err)
			return
		}
	} else {
		dstFile, err = os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("dst open failed:", err)
			return
		}
	}
	defer dstFile.Close()

	_, err = dstFile.WriteString(src.Name() + "\n")
	if err != nil {
		fmt.Println("dst write failed:", err)
		return
	}
	n, err := io.Copy(dstFile, sourceFile)
	if err != nil {
		fmt.Println("copy failed:", err)
		return
	}
	fmt.Println("copy bytes:", n)
}
