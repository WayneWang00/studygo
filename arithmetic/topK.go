package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Top K 问题
// 从长度为大数N的无序数组中，选出最大的k个数
// 方法：从数组中取出k个数，构建成小顶堆。然后依次从剩余的N-k个元素中取元素和堆顶的元素比较。如果小于堆顶的元素则舍弃；如果大于堆顶的元素则互换，然后重新构建新的小顶堆。
func main() {
	n := 100
	k := 10
	a := buildSlice(n)
	fmt.Println("a:", a)
	b := kForN(a, k)
	fmt.Println("b:", b)

	/*
		有两个长度为N的升序数组，取出合并后的第N个元素
	*/
	a1 := buildSlice(n)
	heapSort(a1, n)
	fmt.Println("a1:", a1)
	a2 := buildSlice(n)
	heapSort(a2, n)
	fmt.Println("a2:", a2)
	mid := midFor2N(a1, a2, n)
	fmt.Println("mid:", mid)
}

func buildSlice(n int) []int {
	a := make([]int, n)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100)
	}

	return a
}

func midFor2N(a, b []int, n int) int {
	for i, j := 0, 0; i < len(a) || j < len(b); {
		if i+j+2 == n {
			if a[i] > a[j] {
				return a[i]
			} else {
				return b[j]
			}
		}

		if a[i] > b[j] {
			j++
		} else {
			i++
		}
	}

	return 0
}

func kForN(a []int, n int) []int {
	l := len(a)
	heapSort(a, n)
	//fmt.Println("heapSort:", a)

	for i := n; i < l; i++ {
		if a[0] < a[i] {
			a[0], a[i] = a[i], a[0]
			heapSort(a, n)
			//fmt.Println("heapSort:\n", i, a)
		}
	}

	return a[0:n]
}

func heapSort(a []int, n int) {
	buildSmallHeap(a, n)
	//fmt.Println("buildSmallHeap:", a)

	for i := n - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		building(a, 0, i)
	}
}

func buildSmallHeap(a []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		building(a, i, n)
	}
}

func building(a []int, i, n int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i

	if left < n && a[left] > a[largest] {
		largest = left
	}
	if right < n && a[right] > a[largest] {
		largest = right
	}

	if largest != i {
		a[largest], a[i] = a[i], a[largest]
		building(a, largest, n)
	}
}
