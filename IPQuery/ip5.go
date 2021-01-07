package main

import (
	"Wayne/IPQuery/ipquery"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ip := "14.116.139.99"
	item := strings.SplitN(ip, ".", 4)
	itemint := [4]int{}
	for i := 0; i < len(item); i++ {
		itemint[i], _ = strconv.Atoi(item[i])
		fmt.Println(itemint[i])
	}
	begin := itemint[0]<<24 | itemint[1]<<16 | itemint[2]<<8 | itemint[3]
	fmt.Println("begin: ", begin)
	begin32 := uint32(begin)
	fmt.Println("begin32: ", begin32)
	for k, v := range item {
		fmt.Println("k: ", k, " v: ", v)
	}
	city, err := ipquery.IpData.City(ip)
	fmt.Println("city: ", city, " err: ", err)
}
