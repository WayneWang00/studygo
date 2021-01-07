package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("start goroutines")
	go func() {
		defer wg.Done()
		for count :=0; count < 3; count++{
			for char := 'a'; char < 'a'+26; char++{
				fmt.Printf("%c ",char)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for count :=0; count < 3; count++{
			for char := 'A'; char < 'A'+26; char++{
				fmt.Printf("%c ",char)
			}
		}
	}()
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
	defer_call()
}
func defer_call() {
	defer func() {fmt.Println(111)}()
	defer func() {fmt.Println(222)}()
	defer func() {fmt.Println(333)}()
	fmt.Println(444)
	fmt.Println(555)
	fmt.Println(666)
}
