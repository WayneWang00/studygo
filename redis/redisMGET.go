package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"strings"
)

type info struct {
	Uid    int64
	Nick   string
	Avatar string
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}

	infoList := make(map[string]interface{})
	for i := 1; i < 4; i++ {
		info := new(info)
		info.Nick = "name" + strconv.Itoa(i)
		info.Avatar = "avatar" + strconv.Itoa(i)
		infoList[strconv.Itoa(i)] = strings.Join([]string{info.Nick, info.Avatar}, " ")
	}
	mSet, err := client.HMSet("ranklist", infoList).Result()
	fmt.Println(mSet, err)

	mGet, err := client.MGet("1", "2", "3").Result()
	fmt.Println(mGet, err)
}
