package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/**
  *应用场景：1. 验证码
			2. 随机密码
			3. 抽奖
			4. 随机算法
**/

func main() {
	//testPerm()
	testRand()
	//mathRand()
}

func testPerm() {
	// 需要设置随机种子，如果不设置，每次的运行结果都一样
	rand.Seed(time.Now().UnixNano())
	s := rand.Perm(5)
	fmt.Printf("s:%+v\n", s)
}

func testRand() {
	// 没设置随机种子，默认为seed(1)
	fmt.Println("no seed:", rand.Intn(10))
	fmt.Println("no seed:", rand.Float64())
	len := 4
	rand.Seed(time.Now().Unix())
	a := rand.Intn(10 ^ len)
	fmt.Printf("a value is %d\n", a)
	b := float32(rand.Intn(10^len)) / float32(10^len)
	fmt.Printf("b type %T\n", b)
	fmt.Printf("test type %T\n", 1.0000)
	fmt.Println("b: ", b)
	c := rand.Float64()
	fmt.Println(c)
	d := round(c, len) + float64(rand.Intn(99))
	fmt.Println(d)
}

func round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc(f*pow10_n+0.5) / pow10_n
}

// 伪随机数生成器
func mathRand() {
	rand.Seed(time.Now().UTC().UnixNano())

	c := fanIn(genrt(), genrt())
	for i := 0; i < 10000; i++ {
		fmt.Println(<-c)
	}
}

func fanIn(a <-chan int, b <-chan int) <-chan string {
	c := make(chan string)

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
			c <- rand.Intn(6) + 1
			time.Sleep(time.Duration(500 * time.Millisecond))
		}
	}()

	return c
}
