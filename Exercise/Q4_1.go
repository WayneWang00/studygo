package main

import "fmt"

func main() {
	slice := []float64{1, 2, 3, 4, 5, 6}
	var sum float64 = 0
	for i, _ := range slice {
		sum = sum + slice[i]
	}
	num := float64(len(slice))
	average := sum / num
	fmt.Println("切片slice的平均数为: ", average)
}
