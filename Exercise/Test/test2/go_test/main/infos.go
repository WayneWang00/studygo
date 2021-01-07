//抓取http://www.boyaa.com/news.html的信息，并将数据存储在mysql中，查询时，需要生成redis缓存
//一共三个页面：
//1.新闻列表页
//2.新闻详情页
//3.搜索页
//细节要求：
//1.数据需要定时抓取，并有相应的更新策略。
//2.抓取的数据要存储在mysql。
//3.做页面展示的数据，需要有cache层，统一使用redis。
//4.页面渲染，统一使用golang的html/tempalte包，页面头部和底部需要分离。
//5.页面排版，需要分离css和js等文件，展示效果可以自由发挥。
//6.列表页需要有翻页效果。
//7.能对新闻标题和内容进行搜索，并对关键词高亮显示。

package main

import (
	"crypto/md5"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron"
	"log"
	"strings"
)

var infoSli []News
var count = 0
var rootSrc = "http://www.boyaa.com/"

//func main() {
//	getNewsInfo()
//}

func getContent(str string) (content string, imgsrc string) {
	doc, err := goquery.NewDocument(rootSrc + str)
	checkError(err)

	imgAll := ""
	contentAll := ""
	doc.Find(".middle p").Each(func(i int, selection *goquery.Selection) {
		content = selection.Text()
		contentAll = contentAll + "<br><br>" + content
		img := selection.Find("img")
		for i := 0; i < len(img.Nodes); i++ {
			imgsrc, _ := img.Attr("src")
			imgsrc = strings.TrimLeft(imgsrc, "../../")
			imgsrc = rootSrc + imgsrc
			imgAll = imgCombine(imgAll, imgsrc)
		}
	})
	return contentAll, imgAll
}

func getDomInfoBySelect(doc *goquery.Document, s string) {
	imgAll := ""
	doc.Find(s).Each(func(i int, selection *goquery.Selection) {
		title := selection.Text()
		c := []byte(title)
		title = string(c[14:])

		img := selection.Find("img")
		for i := 0; i < len(img.Nodes); i++ {
			imgsrc, _ := img.Attr("src")
			imgsrc = rootSrc + imgsrc
			imgAll = imgCombine(imgAll, imgsrc)
		}

		link := selection.Find("a")
		lin, _ := link.Attr("href")

		t := selection.Find("em")
		time := t.Text()
		content, src2 := getContent(lin)
		imgAll = imgCombine(imgAll, src2)

		infoSli = append(infoSli, News{Id: count, Title: title, Url: lin, Image: imgAll, Time: time, Content: content})
		count++
		imgAll = ""
	})
}

func imgCombine(str1 string, str2 string) string {
	if str1 == str2 {
		return str1
	} else if strings.Contains(str1, str2) {
		return str1
	} else if strings.Contains(str2, str1) {
		return str2
	} else if len(str1) > 0 {
		return str1 + ";" + str2
	} else {
		return str2
	}
}

var myMD5 [16]byte

func isUpdate(s string) bool {
	t := md5.Sum([]byte(s))
	fmt.Println("mdt")
	if myMD5 == t {
		fmt.Println("not message update")
		return false
	} else {
		myMD5 = t
		fmt.Println("message update")
		return true
	}
}

func getNewsInfo() {
	//抓取数据
	doc, err := goquery.NewDocument("http://www.boyaa.com/news.html")
	checkError(err)

	s, err := doc.Html()
	if !isUpdate(s) {
		return
	}

	//logo
	//getDomInfoBySelect(doc, ".logo")

	//顶部信息
	imgAll := ""
	doc.Find(".topnews").Each(func(i int, selection *goquery.Selection) {
		title := selection.Find("h2 a").Text()
		url, _ := selection.Find("h2 a").Attr("href")

		img := selection.Find("img")
		for i := 0; i < len(img.Nodes); i++ {
			imgsrc, _ := img.Attr("src")
			imgsrc = rootSrc + imgsrc
			imgAll = imgCombine(imgAll, imgsrc)
		}
		content, src2 := getContent(url)
		imgAll = imgCombine(imgAll, src2)
		time := selection.Find(".time").Text()

		infoSli = append(infoSli, News{Id: count, Title: title, Url: url, Image: imgAll, Time: time, Content: content})
		count++
	})

	//页面主体
	getDomInfoBySelect(doc, ".newsbox ul li")

	//fmt.Print(infoSli)

	DeleteAllNews(tabName)
	InsertAllNews(infoSli)
	//return infoSli
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get:%v", err)
	}
}

//定时刷新任务
func StartRefreshNews() {
	c := cron.New()
	spec := "1/5 * * * * ?" // 23:00:00 update
	c.AddFunc(spec, func() {
		getNewsInfo()
	})
	c.Start()
}
