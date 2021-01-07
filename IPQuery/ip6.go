package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	pwd, _ := os.Getwd()
	fmt.Println("pwd: ", pwd)
	pwd1 := filepath.Dir(pwd)
	fmt.Println("pwd1: ", pwd1)
	pwd2 := filepath.Base(pwd1)
	fmt.Println("pwd1: ", pwd1)
	fmt.Println("pwd2: ", pwd2)
	fmt.Println(os.Getenv("GOPATH"))
}
