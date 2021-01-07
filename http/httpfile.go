package main

import (
	"net/http"
	"log"
)

func main() {
	h := http.FileServer(http.Dir("E:/GOPATH/Global/src"))//Dir是一个用sting定义的新数据类型
	//http.Handle("/static/", http.StripPrefix("/static/", h))
	http.Handle("/", http.StripPrefix("/", h))
	//http.Handle("/Wayne/", http.StripPrefix("/Wayne/", h))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
