package crawldata

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
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
	MD5        string
}

// 生成唯一key
func (news *News) getUniqueKey() string {
	return news.Url
}

var user = "root"
var pwd = "123456"
var dbName = "Test"

//TODO:use a pool
func GetDBConnect() *sql.DB {
	db, err := sql.Open("mysql", user+":"+pwd+"@/"+dbName+"?charset=utf8")
	checkErr(err)
	return db
}

func CreateNewsTable(user, password, dbNmae string) {
	str := `CREATE TABLE news (
		id  int NOT NULL AUTO_INCREMENT , 
		title  text NOT NULL , content  text NULL , 
		url  varchar(255) NULL , 
		image  varchar(255) NULL , 
		time  date NULL , 
		modify_time  date NULL , PRIMARY KEY (id));`
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
	str := "INSERT news set title = ?, content = ?, url = ?, time = ?, modify_time = ?"
	db := GetDBConnect()
	defer db.Close()
	stmt, err := db.Prepare(str)
	checkErr(err)
	res, err := stmt.Exec(news.Title, news.Content, news.Url, news.Time, time.Now().Format("2006-01-02"))
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("current id :" + string(id))

	InsertToRedis(news)
}

func InsertToRedis(news News) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	key := news.getUniqueKey()
	str, _ := json.Marshal(news)

	_, err = c.Do("SET", key, string(str))
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}

func GetNewsByUrl(url string) (news News) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	str, err := redis.Bytes(c.Do("get", url))
	if err != nil {
		fmt.Println("redis get failed:", err)
		return getNewsInMysql(url)
	}
	news = News{}
	json.Unmarshal(str, &news)
	return
}

func getNewsInMysql(url string) News {
	str := "select * from news where url=" + url
	db := GetDBConnect()
	defer db.Close()
	rows, err := db.Query(str)
	checkErr(err)
	for rows.Next() {
		var id int
		var title string
		var content string
		var url string
		var image string
		var time string
		var modifyTime string
		var md5 string
		err = rows.Scan(&id, &title, &content, &url, &image, &time, &modifyTime, &md5)
		checkErr(err)
		return News{id, title, content, url, image, time, modifyTime, md5}
	}
	return nil
}

func UpdateNews(news News) {

}

//更新Redis
func UpdateToRedis(news News, i int64) {

}

func main() {
	CreateNewsTable("root", "root123", "test1")
	//InsertNews(News{})
	//InsertToRedis(News{})

	v := News{1, "a", "b", "sfsdf", "imge", "2018-9-10", "2018-9-11", "sdfd"}
	//InsertToRedis(v)
	//fmt.Println(v)
	fmt.Println(GetNewsByUrl(v.Url))

}
