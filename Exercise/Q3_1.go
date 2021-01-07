package main

import "fmt"

func main() {
	for i := 0; ; {
		for j := 0; ; j++ {
			k := 0
			for ; ; k++ {
				i++
				if i > 100 {
					break
				}
				if j > k {
					fmt.Print("A")
				}
				if j == k {
					fmt.Println("A")
					break
				}
			}
			if i > 100 {
				break
			}
		}
		break
	}
}
