package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Nlist struct {
	Time  string
	Url   string
	Title string
}

func main() {
	fmt.Println("开始：")
	tmpl := template.Must(template.ParseFiles("news.html"))
	nlist := []Nlist{
		{"2018-08-23", "http://www.boyaa.com/page/news/newsinfo69.html", "博雅互动公布2018年中期业绩，巩固现有市场地位，加大力度拓展海外市场及其他棋牌游戏业务"},
		{"2018-05-23", "http://www.boyaa.com/page/news/newsinfo68.html", "博雅互动公布2018年第一季度业绩，纯利增17.7% 持续加强产品精细化运营 专注巩固提升用户满意度"},
		{"2018-03-26", "http://www.boyaa.com/page/news/newsinfo67.html", "博雅互动公布2017全年业绩，经调整纯利增长约8.7% 致力打造百年棋牌品牌"},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct{ Nlist []Nlist }{nlist})
	})

	http.ListenAndServe(":8080", nil)
}
