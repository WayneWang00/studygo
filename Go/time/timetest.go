package main

import (
	"fmt"
	"time"
)

const format = "15:04"

func main() {
	t1 := "00:00"
	t2 := "00:00"
	time1, err := time.ParseInLocation(format, t1, time.Local)
	if err != nil {
		fmt.Println("t1:", err)
	}
	time2, err := time.ParseInLocation(format, t2, time.Local)
	if err != nil {
		fmt.Println("t2:", err)
	}
	fmt.Println("t1 before t2:", time1.Before(time2))
	fmt.Println("t1 after t2:", time1.After(time2))
}
