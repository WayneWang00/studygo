package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

type info struct {
	Name  string
	Total string
	Used  string
	Free  string
	Buff  string
	Avail string
}

type Info struct {
	Infos []info
}

var myInfo Info

func setInfo(i string) {
	cmd := exec.Command("top", "-bn 1")
	var out bytes.Buffer
	cmd.Stdout = &out //输出
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 5; i++ {
		line, _ := out.ReadString('\n')
		if i == 3 || i == 4 {
			fmt.Println(line)
		}
	}
	m := info{"Mem", "100", "50", "50", "30", "20"}
	s := info{"Swap", "101", "50", "51", "30", "20"}
	myInfo = Info{[]info{m, s}}
}

func getInfo() Info {
	setInfo("")
	return myInfo
}

func showTopInfo(w http.ResponseWriter, r *http.Request) {
	f := template.Must(template.ParseFiles("home/website/src/Wayne/Test/test1/topInfo.html"))
	f.Execute(w, getInfo())
	//fmt.Fprint(w, "hello")
}

func main() {
	http.HandleFunc("/", showTopInfo)
	http.ListenAndServe(":9999", nil)
}
