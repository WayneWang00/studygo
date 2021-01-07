package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"regexp"
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
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	//for i := 0; i < 5; i++ {
	//	line, _ := out.ReadString('\n')
	//	if i == 3 || i == 4 {
	//		fmt.Println(line)
	//	}
	//}
	//m := info{"Mem", "100", "50", "50", "30", "20"}
	//s := info{"Swap", "101", "50", "51", "30", "20"}
	//myInfo = Info{[]info{m, s}}
	for i := 0; i < 5; i++ {
		//if i == 3 || i == 4 {
		line, _ := out.ReadString('\n')
		memInfo := regexp.MustCompile("^Mem:[0-9a-z,\\s]*").FindString(line)
		fmt.Println("memInfo: ", memInfo)

		mg := regexp.MustCompile("([0-9]*k)").FindAllString(memInfo, -1)
		//refnm := reflect.ValueOf(m).Elem()
		//var m Info
		//var j = 0
		//for k := range m.Info1s[0] {
		//	m.Info1s[0][k] = mg[j]
		//	j++
		//}
		//for i := 0; i < len(mg); i++ {
		//
		//}
		fmt.Println("mg: ", mg)
		//}
	}
}

//r := regexp.MustCompile("([0-9]*k)\\s*total")
//t:= r.FindAllStringSubmatch("Mem:800 805034k total 0999k total", -1)
//fmt.Println(t[0][1])

//Mem:   1019924k total,   827520k used,   192404k free,    15216k buffers
//
//Swap:  4194300k total,        8k used,  4194292k free,   291920k cached

func getInfo() Info {
	setInfo("")
	return myInfo
}

func showTopInfo(w http.ResponseWriter, r *http.Request) {
	f := template.Must(template.ParseFiles("./topInfo.html"))
	f.Execute(w, getInfo())
}

func main() {
	getInfo()
	//http.HandleFunc("/get_top_info", showTopInfo)
	//http.ListenAndServe("localhost:8888", nil)
}
