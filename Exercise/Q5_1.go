package main

import "fmt"

func Average(xs []float64) (ave float64) {
	sum := 0.0
	for _, v := range xs {
		sum = sum + v
	}
	ave = sum / float64(len(xs))
	return
}

func main() {
	xs := []float64{1, 2, 3, 4, 5, 6}
	fmt.Printf("average: %f", Average(xs))
}
