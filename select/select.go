package main

import (
	"fmt"
	"time"
)

var chanClose = make(chan struct{})
var timeout = make(chan bool)

func main() {
	go func() {
		go func() {
			fmt.Println("start sleep")
			time.Sleep(5*time.Second)
			timeout <- true
		}()
		select {
		case <-chanClose:
			fmt.Println("close goroutine")
			return
		case <-timeout:
			fmt.Println("end goroutine")
		}
	}()

	time.Sleep(2*time.Second)
	chanClose <- struct{}{}

	time.Sleep(5*time.Second)
}
