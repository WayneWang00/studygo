package main

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	RtxMsg    int = 100 + iota // 消息
	RtxNotice                  // 通知
)

var Req = struct {
	Url       string `json:"url"`
	SecretKey string `json:"secret_key"`
	AppId     string `json:"app_id"`
	Method    string `json:"method"`
}{
	//Url:       "http://rtx.boyaa.com:8012/rtx.php",
	//SecretKey: "&(9",
	Url:       "http://list.oa.com/api/rest.php",
	SecretKey: "6fe7f10342bfae481c09665dbb86f5e7",
	AppId:     "1",
	Method:    "Message.Send",
}

type Rtx struct {
	typ      string
	receiver []string
	title    string
	content  string
	clientip string
}

func (r *Rtx) Send() {
	client := http.Client{}

	postData := r.MkBody()
	fmt.Println("data:", postData)
	req, err := http.NewRequest(http.MethodPost, Req.Url, strings.NewReader(postData))
	if err != nil {
		fmt.Println("new request failed:", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 4e9)
	req = req.WithContext(ctx)
	req.Header.Set("USER_AGENT", "BOYAA.COM API GoLang Client")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("post failed:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var b = make(map[string]interface{})
	err = json.Unmarshal(body, &b)
	if err != nil {
		fmt.Println("unmarshal failed:", err)
	}
	fmt.Printf("b:%+v\n", b)
	fmt.Printf("code:%d body:%s err:%v\n", resp.StatusCode, string(body), err)
}

func (r *Rtx) MkBody() string {
	receivers := strings.Join(r.receiver, ",")
	params := []string{
		fmt.Sprintf("time=%d", time.Now().Unix()),
		fmt.Sprintf("param[type]=%s", url.QueryEscape(r.typ)),
		fmt.Sprintf("param[title]=%s", url.QueryEscape(r.title)),
		fmt.Sprintf("param[receiver]=%s", url.QueryEscape(receivers)),
		fmt.Sprintf("param[content]=%s", url.QueryEscape(r.content)),
		fmt.Sprintf("param[clientip]=%s", url.QueryEscape(r.clientip)),
		fmt.Sprintf("method=%s", Req.Method),
		fmt.Sprintf("appid=%s", Req.AppId),
	}

	sort.Sort(sort.Reverse(sort.StringSlice(params))) // 逆序
	paramStr := strings.Join(params, "&")
	paramMd5 := fmt.Sprintf("%s%s%s", Req.SecretKey, paramStr, Req.SecretKey)
	sum := md5.Sum([]byte(paramMd5))

	return fmt.Sprintf("%s&sig=%x", paramStr, sum)
}

func SendRtx(typ int, receiver []string, title, content string) {
	if len(receiver) <= 0 || content == "" {
		fmt.Println("参数错误")
		return
	}

	var msg = &Rtx{
		receiver: receiver,
		title:    title,
		content:  content,
		clientip: "0.0.0.0",
	}
	switch typ {
	case RtxMsg:
		msg.typ = "rtx"
	case RtxNotice:
		msg.typ = "notify"
	default:
		msg.typ = "rtx;notify"
	}

	msg.Send()
}

func main() {
	SendRtx(0, []string{"WayneWang"}, "rtx测试", "test")
}
