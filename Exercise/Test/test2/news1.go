package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
	"strings"
	"time"
)

type News struct {
	Id         int
	Title      string
	Content    string
	Url        string
	Image      string
	Time       string
	ModifyTime string
	//MD5        string
}

// 生成唯一key
func (news *News) getUniqueKey() string {
	return news.Url
}

var user = "root"
var pwd = "123456"
var dbName = "Test"
var tabName = "news"
var exTime = 7200 // 过期时间 s

//TODO:use a pool
func GetDBConnect() *sql.DB {
	db, err := sql.Open("mysql", user+":"+pwd+"@/"+dbName+"?charset=utf8")
	checkErr(err)
	return db
}

func CreateNewsTable(tabName string) {
	str := "CREATE TABLE " + tabName + " (id  int NOT NULL AUTO_INCREMENT , title  text NOT NULL , content  text NULL , url  varchar(255) NULL , image  varchar(255) NULL , time  date NULL , modify_time  date NULL , PRIMARY KEY (id));"
	db := GetDBConnect()
	defer db.Close()
	stmt, err := db.Prepare(str)
	checkErr(err)
	stmt.Exec()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// 插入新记录
func InsertNews(news News) {
	str := "INSERT " + tabName + " set title = ?, content = ?, url = ?, image = ?, time = ?, modify_time = ?"
	db := GetDBConnect()
	defer db.Close()
	stmt, err := db.Prepare(str)
	checkErr(err)
	res, err := stmt.Exec(news.Title, news.Content, news.Url, news.Image, news.Time, time.Now().Format("2006-01-02"))
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("current id :" + string(id))

	InsertToRedis(news)
}

func InsertToRedis(news News) {
	c, err := redis.Dial("tcp", "192.168.56.102:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	key := news.getUniqueKey()
	str, _ := json.Marshal(news)

	_, err = c.Do("SET", key, string(str), "EX", exTime)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}

func GetNewsByUrl(url string) (news News) {
	c, err := redis.Dial("tcp", "192.168.56.102:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	str, err := redis.Bytes(c.Do("get", url))
	fmt.Println(str)
	if err != nil {
		fmt.Println("redis get failed:", err)
		news = getNewsFromMysql(url)
		if news.Url != "" {
			InsertToRedis(news) // add to redis
		}
		return
	}
	news = News{}
	json.Unmarshal(str, &news)
	return news
}

func getNewsFromMysql(url string) News {
	str := "select * from " + tabName + " where url=?"
	db := GetDBConnect()
	defer db.Close()
	rows, err := db.Query(str, url)
	checkErr(err)
	for rows.Next() {
		var id int
		var title string
		var content string
		var url string
		var image string
		var time string
		var modifyTime string
		//var md5 string
		err = rows.Scan(&id, &title, &content, &url, &image, &time, &modifyTime)
		checkErr(err)
		return News{id, title, content, url, image, time, modifyTime}
	}
	return News{}
}

func UpdateNews(news News) {

}

//更新Redis
func UpdateToRedis(news News, i int64) {

}

func main() {
	//CreateNewsTable()

	v := News{Title: "a", Content: "b", Url: "bcd", Image: "imge", Time: "2018-9-10"}
	//InsertNews(v) // 插入新闻
	//fmt.Println(GetNewsByUrl(v.Url)) //查询
	n := GetNewsByUrl(v.Url)
	n1 := GetNewsByUrl("abc")
	n2 := GetNewsByUrl("cde")
	news := []News{n, n1, n2}
	NewsShow(news)
	//tmpl := template.Must(template.ParseFiles("/home/website/src/Wayne/Test/test2/news/index.html"))
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	tmpl.Execute(w, struct{ News []News }{news})
	//})
	//tmp2 := template.Must(template.ParseFiles("/home/website/src/Wayne/Test/test2/news/content.html"))
	//http.HandleFunc("/content", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println(r.URL)
	//	url := strings.Split(r.URL.String(), "?")
	//	for _, k := range news {
	//		fmt.Println("k: ", k, "url: ", url[1])
	//		if k.Url == url[1] {
	//			tmp2.Execute(w, k)
	//		}
	//	}
	//	//tmp2.Execute(w, n)
	//})
	//http.ListenAndServe(":8080", nil)
}

func NewsShow(news []News) {
	tmpl := template.Must(template.ParseFiles("/home/website/src/Wayne/Test/test2/news/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, struct{ News []News }{news})
	})
	tmp2 := template.Must(template.ParseFiles("/home/website/src/Wayne/Test/test2/news/content.html"))
	http.HandleFunc("/content", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		url := strings.Split(r.URL.String(), "?")
		for _, k := range news {
			fmt.Println("k: ", k, "url: ", url[1])
			if k.Url == url[1] {
				tmp2.Execute(w, k)
			}
		}
	})
	http.ListenAndServe(":8080", nil)
}
