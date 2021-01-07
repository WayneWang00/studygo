package main

import (
	"fmt"
	"github.com/tabalt/ipquery"
)

func main() {
	//df := "/home/website/src/github.com/tabalt/ipquery/testdata/test_10000.txt"
	//df := "/home/website/src/github.com/tabalt/ipquery/testdata/ip_chunzhen.txt"
	//df := "/home/website/src/qqwry.dat"
	df := "/home/website/src/111.txt"
	err := ipquery.Load(df)
	if err != nil {
		fmt.Println(err)
	}
	//ip := "61.149.208.1"
	ip := "14.116.139.99"
	dt, err := ipquery.Find(ip)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ip: ", string(dt))
	}
}
