package main

import (
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

//type News struct {
//	Id         int
//	Title      string
//	Content    string
//	Url        string
//	Image      string
//	Time       string
//	ModifyTime string
//	MD5        string
//}

var Db *sql.DB
var Pool redis.Pool

func init() { //init 用於初始化一些參數，先於main執行
	Pool = redis.Pool{
		MaxIdle:     16,
		MaxActive:   32,
		IdleTimeout: 120,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "192.168.56.102:6379")
		},
	}
	Db, _ = sql.Open("mysql", "root:123456@/Test?charset=utf8")
}

func main() {
	conn := Pool.Get()
	r1, err := conn.Do("SET", "name", "new")
	fmt.Println(r1, err)
	r2, err := redis.String(conn.Do("GET", "name"))
	fmt.Println("res1: ", r2, "err: ", err)
	row, err := Db.Query("SELECT * FROM my_int WHERE int_5=5")
	if err != nil {
		fmt.Println("query error: ", err)
		return
	}
	for row.Next() {
		var int_1, int_2, int_3, int_4, int_5, int_6 int
		err = row.Scan(&int_1, &int_2, &int_3, &int_4, &int_5, &int_6)
		fmt.Println(int_1, int_2, int_3, int_4, int_5, int_6)
	}
	//v := News{1, "a", "b", "sfsdf", "imge", "2018-9-10", "2018-9-11", "sdfd"}
	//InsertNews(v)
	//var req string
	//res, err := GetNews(req)
	//if res == "" {
	//	if err != nil {
	//		res1, err := GetNewsSQL(req)
	//		if err != nil {
	//			fmt.Println("select err: ", err)
	//		}
	//	}
	//}
	//resp, err := http.Get("http://www.boyaa.com/news.html")
	//defer resp.Body.Close()
	//if err != nil {
	//	fmt.Println("error: ", err)
	//} else {
	//	b, _ := ioutil.ReadAll(resp.Body)
	//
	//	fmt.Println(string(b))
	//}
	doc, err := goquery.NewDocument("http://www.boyaa.com/news.html")
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		doc.Find(".section.section.section.div.ul").Each(func(i int, s *goquery.Selection) { //获取节点集合并遍历
			text := s.Find("a").Text() //获取匹配节点的文本值
			fmt.Println("text: ", text)
		})
	}
	//ExampleScrape()
}

// 插入新记录
//func InsertNews(news News) {
//	str := "INSERT news set title = ?, content = ?, url = ?, time = ?, modify_time = ?"
//	//defer Db.Close()
//	stmt, err := Db.Prepare(str)
//	fmt.Println(err)
//	res, err := stmt.Exec(news.Title, news.Content, news.Url, news.Time, time.Now().Format("2006-01-02"))
//	fmt.Println(err)
//	id, err := res.LastInsertId()
//	fmt.Println(err)
//	fmt.Println("current id :" + string(id))
//
//	InsertToRedis(news)
//}

// 生成唯一key
//func (news *News) getUniqueKey() string {
//	return news.Url
//}

//func InsertToRedis(news News) {
//	key := news.getUniqueKey()
//	str, _ := json.Marshal(news)
//	conn := Pool.Get()
//	_, err := conn.Do("SET", key, string(str))
//	if err != nil {
//		fmt.Println("redis set failed:", err)
//	}
//}

//获取信息
func GetNews(url string) (interface{}, error) {
	conn := Pool.Get()
	res, err := conn.Do("GET", url)
	if err != nil {
		return res, err
	}
	return res, nil
}

func GetNewsSQL(url string) (*sql.Rows, error) {
	format := "SELECT `title`,`content`,`url`,`content`,`time`,`modify_time` FROM %s WHERE `url` = %s "
	sql := fmt.Sprintf(format, "news", url)
	res, err := Db.Query(sql)
	if err != nil {
		fmt.Println("select error: ", err)
		return nil, err
	}
	return res, nil
}

//func ExampleScrape() {
//	// Request the HTML page.
//	res, err := http.Get("http://metalsucks.net")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer res.Body.Close()
//	if res.StatusCode != 200 {
//		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
//	}
//
//	// Load the HTML document
//	doc, err := goquery.NewDocumentFromReader(res.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Find the review items
//	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
//		// For each item found, get the band and title
//		band := s.Find("a").Text()
//		title := s.Find("i").Text()
//		fmt.Printf("Review %d: %s - %s\n", i, band, title)
//	})
//}
