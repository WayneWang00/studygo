package main

import (
	"fmt"
	"sync"
)

var (
	ch1 chan bool
	ch2 chan bool
	ex  chan struct{}
)

func main() {
	//print10()
	print100()
}

func print10() {
	//runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)

	go goA(&wg)
	go goB(&wg)

	wg.Wait()
	//time.Sleep(100 * time.Millisecond)
}

func goA(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func goB(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func print100() {
	ch1 = make(chan bool, 1)
	ch2 = make(chan bool)
	ex = make(chan struct{})

	go go1()
	go go2()

	ch1 <- true
	<-ex
}

func go1() {
	for i := 1; i <= 50; i++ {
		if ok := <-ch1; ok {
			fmt.Println("goroutine1:", 2*i-1)
			ch2 <- true
		}
	}
}

func go2() {
	defer func() {
		close(ex)
	}()
	for i := 1; i <= 50; i++ {
		if ok := <-ch2; ok {
			fmt.Println("goroutine2:", 2*i)
			ch1 <- true
		}
	}
}
