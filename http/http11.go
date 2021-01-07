package main

import (
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	VMaxIdle             = 200  //最大空闲连接数
	VMaxActive           = 200  //最大连接数(包含了空闲的)
	VIdleTimeout         = 180  //池中的连接空闲多久之后超时（单位：秒）(即超过180秒的)
	VRedisConnectTimeout = 5000 //连接超时时间（单位:毫秒）
	VRedisReadTimeout    = 5000 //读超时时间（单位:毫秒）
	VRedisWriteTimeout   = 5000 //写超时时间（单位:毫秒）
	VMAXCOUNT            = 3
)

//type Pool struct {
//	Dail         func() (*http.Client, error)
//	TestOnBorrow func(c *http.Client, t time.Time) error
//	MaxIdle      int           //最大空闲
//	MaxActive    int           //最大连接数
//	IdleTimeout  time.Duration //空闲超时
//	Wait         bool
//	mu           sync.Mutex
//	cond         *sync.Cond
//	closed       bool
//	active       int
//	idle         list.List
//}

//创建一个新的连接池
func NewPool() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:    VMaxIdle,
			MaxConnsPerHost: VMaxActive,
			IdleConnTimeout: VIdleTimeout * time.Second,
		},
	}
}

var (
	httpClient *http.Client
)

func init() {
	httpClient = NewPool()
}

func Get(params string) []byte {
	resp, err := httpClient.Get("http://texas-demo-13.boyaa.com/texas/api/api.php?api=" + params)
	CheckErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)
	return body
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
