package main

import (
	"fmt"
	"time"
)

func main() {
	//参考：https://blog.csdn.net/feiwutudou/article/details/81001453
	//st := "2019-02-28 23:59:59"
	//loc, _ := time.LoadLocation("Asia/Shanghai")
	//sTime, err := time.ParseInLocation("2006-01-02 15:04:05", st, loc)
	//if err != nil {
	//	fmt.Println("time.Parse Error: ", err)
	//}
	//startTime := sTime.Unix()
	//fmt.Println("sTime: ", startTime)
	//tm := time.Unix(startTime, 0)
	//fmt.Println("st: ", tm.Format("2006-01-02 15:04:05"))
	//参考：https://blog.csdn.net/maqianQAQ/article/details/78705001

	//st := "2019-02-24 23:59:59"
	//sTime, err := time.ParseInLocation("2006-01-02 15:04:05", st, time.Local)
	//if err != nil {
	//	fmt.Println("time.Parse Error: ", err)
	//}
	//startTime := sTime.Unix()
	//fmt.Println("startTime: ", startTime)
	//tm := time.Unix(startTime, 0)
	//fmt.Println("st: ", tm.Format("2006-01-02 15:04:05"))

	testUnixNano()
}

func testUnixNano() {
	t := time.Now().UnixNano()
	fmt.Println("t:", t)

	var n31 int64 = 1 << 31
	fmt.Println("n31:", n31)

	if n31 > t {
		fmt.Println("n31>t")
	}

	fmt.Println("x16:", 0xffffff)
	fmt.Printf("%06x\n", 0xffffff)
}
