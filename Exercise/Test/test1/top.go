package main

import (
	"fmt"
	"net/http"
)

func TopInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method)
	if r.Method == "Get" {
		fmt.Println("top")
	}
}

func main() {
	http.HandleFunc("/get_top_info", TopInfo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServer: ", err)
	}
}
