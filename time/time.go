package main

import (
	"fmt"
	"time"
)

const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp       = "Jan _2 15:04:05"
	StampMilli  = "Jan _2 15:04:05.000"
	StampMicro  = "Jan _2 15:04:05.000000"
	StampNano   = "Jan _2 15:04:05.000000000"
	WeekdayTime = "Mon 15:04:05"
)

func main() {
	//nowtime := time.Now().Format("2006-01-02 15:04:05")
	//fmt.Println(nowtime)
	//nowtime1 := time.Now().Format("2006-01-02 15:04")
	//fmt.Println(nowtime1)
	//rfc1123 := time.Now().Format(RFC1123)
	//fmt.Println("rfc822: ", rfc1123)
	//nowTime := time.Now().Format(WeekdayTime)
	//fmt.Println("weekdayTime: ", nowTime)
	////nowTime1, err := time.Parse(WeekdayTime, time.Now().Format(WeekdayTime))
	////if err != nil {
	////	fmt.Println("nowTime1 error: ", err)
	////}
	//nowTime1, err := time.Parse(WeekdayTime, "Sat 22:00:00")
	//if err != nil {
	//	fmt.Println("nowTime1 error: ", err)
	//}
	//nowTime2, err := time.Parse(WeekdayTime, "Sat 22:00:00")
	//if err != nil {
	//	fmt.Println("nowTime2 error: ", err)
	//}
	//if nowTime1.Before(nowTime2) {
	//	fmt.Println("nowTime1 before nowTime2")
	//} else {
	//	fmt.Println("nowTime1 after nowTime2")
	//}

	//testWeekTime()
	timeWeekHour()
}

func testWeekTime() {
	nowTime, err := time.ParseInLocation(WeekdayTime, time.Now().Format(WeekdayTime), time.Local)
	if err != nil {
		fmt.Println("now time parse err:", err)
		return
	} else {
		fmt.Println("not time:", nowTime.Format(RFC1123))
	}
	endTime, err := time.ParseInLocation(WeekdayTime, "Sat 22:00:00", time.Local)
	if err != nil {
		fmt.Println("end time parse err:", err)
		return
	} else {
		fmt.Println("end time:", endTime.Format(RFC1123))
	}
	if nowTime.After(endTime) {
		fmt.Println("now time after end time")
	} else {
		fmt.Println("not time before end time")
	}
}

func timeWeekHour() {
	nowTime := time.Now()
	fmt.Println("nowTime:", nowTime)
	endTime := nowTime.Add((24 + 2) * time.Hour)
	fmt.Println("endTime:", endTime)

	var start int64 = 1596895200
	startUnix := time.Unix(start, 0)
	fmt.Println("start unix:", startUnix)
	endUnix := startUnix.Add((24*1 + 2) * time.Hour)
	fmt.Println("end unix:", endUnix)

	if endTime.After(endUnix) {
		fmt.Println("endTim after endUnix")
	} else {
		fmt.Println("endTime before endUnix")
	}
}
