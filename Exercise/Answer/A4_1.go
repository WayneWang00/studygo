package main

import "fmt"

func main() {
	var ave float64
	sum := 0.0
	xs := []float64{1, 2, 3, 4, 5, 6}
	switch len(xs) {
	case 0:
		ave = 0
	default:
		for _, v := range xs {
			sum = sum + v
		}
		ave = sum / float64(len(xs))
	}
	fmt.Printf("average: %f", ave)
}
