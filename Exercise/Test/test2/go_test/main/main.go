package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var pageCount = 8

var funcMap = map[string]func(http.ResponseWriter, *http.Request){
	"/list":   ShowList,
	"/search": Search,
	"/detail": ShowDetail,
}

func main() {
	//CreateNewsTable(tabName)
	StartRefreshNews() // 定时更新
	fmt.Println("server start ...")
	for name, f := range funcMap {
		http.HandleFunc(name, f)
	}
	http.Handle("/img/", http.FileServer(http.Dir("./html")))
	http.Handle("/css/", http.FileServer(http.Dir("./html")))
	http.Handle("/js/", http.FileServer(http.Dir("./html")))

	http.ListenAndServe(":9999", nil)
}

func Search(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	keyword := r.Form.Get("keyword")
	allNews := SearchFromMysql(keyword)

	if keyword == "" {
		t := template.Must(template.ParseFiles("./view/search.html"))
		t.Execute(w, nil)
	} else {
		t := template.Must(template.ParseFiles("./view/searchmsg.html"))
		n := make([]interface{}, 0)
		for i := 0; i < len(allNews); i++ {
			t := struct {
				News
				Title template.HTML
			}{allNews[i], template.HTML(strings.Replace(allNews[i].Title, keyword, "<span style=\"color:red\">"+keyword+"</span>", -1))}
			n = append(n, t)
		}
		t.Execute(w, n)
	}
	//test
	//todo:根据allNews去填充界面
	for _, v := range allNews {
		fmt.Println(v.Id, v.Title)
	}
}

func ShowList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	page := r.Form.Get("page")
	if page == "" {
		page = "1"
	}
	page2, _ := strconv.Atoi(page)
	allNews := GetNews(page2, pageCount)

	t := template.Must(template.ParseFiles("./view/index.html"))
	t.Execute(w, allNews)

	//test
	//todo:根据allNews去填充界面
	for _, v := range allNews {
		fmt.Println(v.Id, v.Title)
	}
}

func ShowDetail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")
	news := GetNewsByKey(id)

	t := template.Must(template.ParseFiles("./view/content.html"))

	ss := strings.Split(news.Image, ";")
	for i, v := range ss {
		if len(v) == 0 {
			ss = append(ss[:i], ss[i+1:]...)
		}
	}

	n := struct {
		News
		Images []string
	}{news, ss}

	t.Execute(w, n)

	fmt.Println(news)
}
