package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println(i, " = FizzBuzz")
			}
			fmt.Println(i, " = Fizz")
			continue
		}
		if i%5 == 0 {
			if i%3 == 0 {
				fmt.Println(i, " = FizzBuzz")
			}
			fmt.Println(i, " = Buzz")
			continue
		}
		fmt.Println(i)
	}
}
