package main

import "fmt"

func main() {
	var i, j, row, col, max int
	var f = [3][4]int{{1, 3, 4}, {2, 3, 5}, {1, 3, 4, 6}}
	max = f[0][0]
	for i = 0; i < 3; i++ {
		for j = 0; j < 4; j++ {
			if f[i][j] > max {
				max = f[i][j]
				row = i
				col = j
			}
		}
	}
	fmt.Printf("max = %d, row = %d, col = %d", max, row, col)
}
