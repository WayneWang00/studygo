package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("say hello name")
	fmt.Println("method:", r.Method)
	r.ParseForm()
	fmt.Println("form:", r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("value: ", strings.Join(v, ""))
	}
	w.Write([]byte("say hello\n"))
	//fmt.Fprintf(w, "sayhello")
}

func getJson(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("E:/GOPATH/Global/src/cdn.json")
		log.Println(t.Execute(w, nil))
	}
}

func main() {
	fmt.Println("web server start...")
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/getjson", getJson)
	http.HandleFunc("/sayHello", sayHello)
	http.HandleFunc("/sayHi", sayHi)
	err := http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("say hello")
	fmt.Println("method: ", r.Method)
	if _, err := fmt.Fprintf(w, "hello\n"); err != nil {
		fmt.Println("sayHello Error: ", err)
	}
}

func sayHi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("say hi")
	fmt.Println("method: ", r.Method)
	fmt.Println("url:", r.URL.Path)
	r.ParseForm()
	fmt.Printf("form type:%T\n", r.Form)
	fmt.Printf("form:%+v\n", r.Form)
	fmt.Println("len form:", len(r.Form))
	fmt.Println("id:", r.Form["id"])
	if _, err := fmt.Fprintf(w, "hi\n"); err != nil {
		fmt.Println("sayHi Error: ", err)
	}
}
