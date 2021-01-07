package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

/**
  *应用场景：1. 生成随机加密串
**/

func main() {
	testPrime()
	//cryptoRand()
}

func testPrime() {
	// 返回一个具有指定位数（二进制位数）的数字，数字为质数的可能性很高
	p, err := rand.Prime(rand.Reader, 5)
	if err != nil {
		fmt.Println("prime failed:", err)
		return
	}
	fmt.Println("prime:", p)
}

// 加密安全随机数生成器
func cryptoRand() {
	c := fanIn(genrt(), genrt())
	for i := 0; i < 10000; i++ {
		fmt.Println(<-c)
	}
}

func fanIn(a <-chan int, b <-chan int) <-chan string {
	var c = make(chan string)

	go func() {
		var count int
		for {
			count += <-a
			c <- fmt.Sprintf("Tally of A is:%d", count)
		}
	}()

	go func() {
		var count int
		for {
			count += <-b
			c <- fmt.Sprintf("Tally of B is:%d", count)
		}
	}()

	return c
}

func genrt() <-chan int {
	var c = make(chan int)

	go func() {
		for i := 0; ; i++ {
			dice, err := rand.Int(rand.Reader, big.NewInt(6))
			if err != nil {
				fmt.Println("rand.int failed:", err)
				return
			}
			c <- int(dice.Int64()) + 1
			time.Sleep(time.Duration(500 * time.Millisecond))
		}
	}()

	return c
}
