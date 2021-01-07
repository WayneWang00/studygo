package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

type Info struct {
	Infos []map[string]string
}

var myInfo Info

func setInfo() {
	cmd := exec.Command("top", "-bn 1")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	var mm map[string]string
	var ss map[string]string

	for i := 0; i < 5; i++ {
		line, _ := out.ReadString('\n')
		memInfo := regexp.MustCompile("^Mem:[0-9a-z,\\s]*").FindString(line)
		swapInfo := regexp.MustCompile("^Swap:[0-9a-z,\\s]*").FindString(line)

		if len(memInfo) > 0 {
			memMap := make(map[string]string, 5)
			m := strings.Trim(strings.Replace(memInfo, "Mem:", " ", -1), " ")
			mm = stringToMap(m, memMap)
		}
		if len(swapInfo) > 0 {
			swapMap := make(map[string]string, 5)
			s := strings.Trim(strings.Replace(swapInfo, "Swap:", " ", -1), " ")
			ss = stringToMap(s, swapMap)
		}
	}

	myInfo = Info{make([]map[string]string, 0)}
	myInfo.Infos = append(myInfo.Infos, mm)
	myInfo.Infos = append(myInfo.Infos, ss)
}

func stringToMap(str string, m map[string]string) map[string]string {
	strSli := strings.Split(str, ",")
	for _, v := range strSli {
		vSli := strings.Split(strings.Trim(v, " "), " ")
		m[vSli[1]] = vSli[0]
	}
	return m
}

func getInfo() Info {
	setInfo()
	return myInfo
}

func showTopInfo(w http.ResponseWriter, r *http.Request) {
	f := template.Must(template.ParseFiles("topInfo.html"))
	f.Execute(w, getInfo())
}

func main() {
	http.HandleFunc("/get_top_info", showTopInfo)
	http.ListenAndServe("localhost:8888", nil)
}
