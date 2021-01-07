package main

import (
	"fmt"
	"github.com/yinheli/qqwry"
)

func main() {
	q := qqwry.NewQQwry("qqwry.dat")
	q.Find("180.89.94.90")
	fmt.Printf("ip:%v, country:%v, city:%v", q.Ip, q.Country, q.City)
}
