package main

import (
	"fmt"
	"time"
)

func main() {
	//UnBuffChannel()
	//BuffChannel()
	//NewUnBuffChannel()
	//NewBuffChannel()
	//NilChannel()
	stopGoroutine()
	fmt.Println("end")
}

func loop(ch chan int) {
	for {
		select {
		case v := <-ch:
			fmt.Println("this value of unbuff channel", v)
		}
	}
}

func UnBuffChannel() {
	ch := make(chan int)
	go loop(ch)
	ch <- 1
	//go loop(ch)	// fatal error:all goroutine are asleep - deadlock! 因为ch<-1 发送了，但是同时没有接受者，所以发生阻塞
	time.Sleep(1 * time.Second)
}

var c = make(chan int)

func f() {
	c <- 'C'
	fmt.Println("在goroutine内")
}

func NewUnBuffChannel() {
	go f()

	c <- 'c'
	<-c
	<-c
	fmt.Println("外部调用")
}

func BuffChannel() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	go loop(ch)
	ch <- 4
	//go loop(ch)	// fatal error: all goroutine are asleep - deadlock! 因为ch的大小为3，而这里要传入4个数据，所以就阻塞
	time.Sleep(1 * time.Second)
}

func writeRoutine(test_chan chan int, value int) {
	test_chan <- value
}

func readRoutine(test_chan chan int) {
	<-test_chan
	return
}

func NewBuffChannel() {
	ch := make(chan int)
	v := 100

	go writeRoutine(ch, v)
	readRoutine(ch)
	fmt.Println(v)
}

func NilChannel() {
	var inCh = make(chan int)
	var outCh = make(chan int)

	go func() {
		var in <-chan int = inCh
		var out chan<- int
		var val int

		for {
			select {
			case out <- val:
				fmt.Println("----------")
				fmt.Println("out val:", val)
				out = nil
				in = inCh
			case val = <-in:
				fmt.Println("++++++++++")
				fmt.Println("in val:", val)
				in = nil
				out = outCh
			}
		}
	}()

	go func() {
		for r := range outCh {
			fmt.Println("Result:", r)
		}
	}()

	time.Sleep(0)
	inCh <- 1
	inCh <- 2
	time.Sleep(3 * time.Second)
}

func stopGoroutine() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
