package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
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
func (news *News) getUniqueKey() interface{} {
	return news.Id
}

var user = "root"
var pwd = "123456"
var dbName = "Test"
var tabName = "news"
var exTime = 60 // 过期时间 s

//TODO:use a pool

func GetDBConnect() *sql.DB {
	db, err := sql.Open("mysql", user+":"+pwd+"@/"+dbName+"?charset=utf8")
	checkErr(err)
	return db
}

func CreateNewsTable(tabName string) {
	str := "CREATE TABLE " + tabName + " (id  int NOT NULL AUTO_INCREMENT , title  text NOT NULL , content  text NULL , url  varchar(255) NULL , image  varchar(512) NULL , time  date NULL , modify_time  date NULL , PRIMARY KEY (id), FULLTEXT (title))ENGINE = MyISAM DEFAULT CHARSET=utf8"
	db := GetDBConnect()
	defer db.Close()
	stmt, err := db.Prepare(str)
	checkErr(err)
	stmt.Exec()
}

func DeleteAllNews(tabName string) {
	str := "truncate  table " + tabName
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
	t := strings.Replace(news.Time, " / ", "-", -1)
	args := []interface{}{news.Title, news.Content, news.Url, news.Image, t, time.Now().Format("2006-01-02")}
	if t == "" {
		str = "INSERT " + tabName + " set title = ?, content = ?, url = ?, image = ?, modify_time = ?"
		args = []interface{}{news.Title, news.Content, news.Url, news.Image, time.Now().Format("2006-01-02")}
	}

	db := GetDBConnect()
	defer db.Close()
	stmt, err := db.Prepare(str)
	checkErr(err)
	res, err := stmt.Exec(args...)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("current id :", id)

	InsertToRedis(news)
}

func InsertAllNews(info []News) {
	str := "INSERT " + tabName + " set title = ?, content = ?, url = ?, image = ?, time = ?, modify_time = ?"
	db := GetDBConnect()
	defer db.Close()
	stmt, err := db.Prepare(str)
	checkErr(err)
	for _, news := range info {
		t := strings.Replace(news.Time, " / ", "-", -1)
		args := []interface{}{news.Title, news.Content, news.Url, news.Image, t, time.Now().Format("2006-01-02")}
		if t == "" {
			str = "INSERT " + tabName + " set title = ?, content = ?, url = ?, image = ?, modify_time = ?"
			args = []interface{}{news.Title, news.Content, news.Url, news.Image, time.Now().Format("2006-01-02")}
		}
		res, err := stmt.Exec(args...)
		checkErr(err)
		id, err := res.LastInsertId()
		checkErr(err)
		fmt.Println("current id :", id)
		InsertToRedis(news)
	}
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

func InsertAllToRedis(t []News) {
	c, err := redis.Dial("tcp", "192.168.56.102:6379")
	checkErr(err)
	defer c.Close()

	for _, news := range t {
		key := news.getUniqueKey()
		str, _ := json.Marshal(news)

		_, err = c.Do("SET", key, string(str), "EX", exTime)
		if err != nil {
			fmt.Println("redis set failed:", err)
		}
	}
}

func GetNewsByKey(key interface{}) (news News) {
	c, err := redis.Dial("tcp", "192.168.56.102:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	str, err := redis.Bytes(c.Do("get", key))
	if err != nil {
		fmt.Println("redis get failed:", err)
		news = getNewsFromMysql(key)
		if news.Url != "" {
			InsertToRedis(news) // add to redis
		}
		return
	}
	news = News{}
	json.Unmarshal(str, &news)
	fmt.Println("get from redis")
	return news
}

func getNewsFromMysql(key interface{}) News {
	str := "select * from " + tabName + " where id=?"
	db := GetDBConnect()
	defer db.Close()
	rows, err := db.Query(str, key)
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

func SearchFromMysql(keyword string) []News {
	str := "select * from " + tabName + " where title like '%" + keyword + "%'"
	db := GetDBConnect()
	defer db.Close()
	rows, err := db.Query(str)
	checkErr(err)

	infos := make([]News, 0, 10)
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
		infos = append(infos, News{id, title, content, url, image, time, modifyTime})
	}

	return infos
}

//todo:关键词的搜索缓存
func SearchNewsByKeyword(keyword string) []News {
	return SearchFromMysql(keyword)
	//c, err := redis.Dial("tcp", "127.0.0.1:6379")
	//checkErr(err)
	//defer c.Close()
	//keys := make([]interface{}, 0)
	//str, err := redis.ByteSlices(c.Do("keys", keyword))
	//checkErr(err)
	//for _, v := range str {
	//	keys = append(keys, v)
	//}
	//return make([]News, 0)
}

// 显示第page页, 每页显示pagecount条数据
func GetNews(page, pageCount int) []News {
	limit1 := pageCount * (page - 1)
	limit2 := pageCount

	str := "select * from " + tabName + " limit ?,?"
	db := GetDBConnect()
	defer db.Close()
	rows, err := db.Query(str, limit1, limit2)
	checkErr(err)
	infos := make([]News, 0, 10)
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
		infos = append(infos, News{id, title, content, url, image, time, modifyTime})
	}
	return infos
}

func GetNewsCount() int {
	str := "select id from " + tabName
	db := GetDBConnect()
	defer db.Close()
	stmt, err := db.Query(str)
	checkErr(err)
	count := 0
	for stmt.Next() {
		count++
	}
	return count
}

func UpdateNews(news News) {

}

//更新Redis
func UpdateToRedis(news News, i int64) {

}

//func main() {
//	CreateNewsTable(tabName) //
//	//StartRefreshNews()
//	//select {}
//}
