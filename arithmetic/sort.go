package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := randSlice(30)
	//BubbleSort(a)
	//SelectSort(a)
	//InsertSort(a)
	//ShellSort(a)
	//MergeSort(a)
	//QuickSort(a)
	//HeapSort(a)
	CountSort(a)
}

// 随机slice
func randSlice(n int) []int {
	var a = make([]int, n)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100)
	}

	fmt.Println("a:\n", a)
	return a
}

// 冒泡排序 O(n2) 稳定
/*
	第1次：取第一个元素跟之后的所有元素依次比较，大于（小于）就互换位置
	第2次：取第二个元素跟之后的所有元素依次比较，大于（小于）就互换位置
	依次类推，直到倒数第二个结束。
*/
func BubbleSort(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	fmt.Println("bubble:\n", a)
	return a
}

// 选择排序 O(n2) 不稳定
/*
	第1次：在未排序的序列中选出最小（大）的元素，放到序列起始位置
	第2次：在剩余未排序的序列中同样选出最小（大）的元素，放到已排序的序列末尾
	依此类推，直到未排序序列只剩最后一个结束。
*/

func SelectSort(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}

	for i := 0; i < l-1; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}

	fmt.Println("selection:\n", a)
	return a
}

// 插入排序 O(n2) 稳定
/*
	第1次：首先在未排序的序列中选择第一个元素作为已排序的元素
	第2次：在未排序的序列中选择第一个元素，跟已排序的序列从后向前（从前向后）比较。当小于已排序的元素时，将已排序的元素向后移一位；反之，将该元素插入到已排序的元素后
	重复第2次的步骤，直到所有未排序的元素取完
*/
func InsertSort(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}

	for i := 1; i < l; i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}

	fmt.Println("insert:\n", a)
	return a
}

// 希尔排序（缩小增量排序/优化的插入排序） O(n*log n) 不稳定
/*
	第1次：确定一个增量将未排序的序列分成几组，然后分别进行插入排序
	第2次：缩小增量将未排序的序列分成几组，然后分别进行插入排序
	以此类推，直到增量缩小为1，进行插入排序
*/
func ShellSort(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}

	gap := 2
	step := l / gap
	for step >= 1 {
		for i := step; i < l; i++ {
			for j := i; j >= step && a[j] < a[j-step]; j -= step {
				a[j], a[j-step] = a[j-step], a[j]
			}
		}
		fmt.Println("step:", step, "\n", a)
		step /= gap
	}

	fmt.Println("shell:\n", a)
	return a
}

// 归并排序 O(n*log n) 稳定
/*
	第1次：将未排序的序列分成两部分，然后分别进行归并排序，然后将排好序的序列合并成一条序列
	第2次：将未排序的序列再分成两部分，然后分别进行归并排序，然后将排好序的序列合并成一条序列
	以此类推，直到所有部分排好序，合并成一条序列
*/
func MergeSort(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}

	mergeSort(a, 0, l-1)
	fmt.Println("merge:\n", a)
	return a
}

func mergeSort(a []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSort(a, start, mid)
	mergeSort(a, mid+1, end)
	merge(a, start, mid, end)
}

func merge(a []int, start, mid, end int) {
	temp := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0

	for ; i <= mid && j <= end; k++ {
		if a[i] < a[j] {
			temp[k] = a[i]
			i++
		} else {
			temp[k] = a[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		temp[k] = a[i]
		k++
	}

	for ; j <= end; j++ {
		temp[k] = a[j]
		k++
	}

	copy(a[start:end+1], temp)
}

// 快速排序 O(n*log n) 不稳定
/*
	第1次：取第一个元素作为标准，将所有小于该元素的放到该元素左边，大于该元素的放到该元素右边。形成两个新的序列
	第2次：将新的两个序列按照步骤1分别进行快速排序
	以此类推，直到不能再分新的序列
*/
func QuickSort(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}

	//quickSort(a, 0, l-1)
	//quickAscSort(a, 0, l-1)
	quickDescSort(a, 0, l-1)
	fmt.Println("quick:\n", a)
	return a
}

func quickSort(a []int, start, end int) {
	if start >= end {
		return
	}

	p := partition(a, start, end)
	quickSort(a, start, p-1)
	quickSort(a, p+1, end)
}

func partition(a []int, start, end int) int {
	temp := a[start]

	i := end
	for j := end; j > start; j-- {
		if a[j] > temp {
			a[i], a[j] = a[j], a[i]
			i--
		}
	}
	a[i], a[start] = a[start], a[i]

	return i
}

// 升序
/*
	第1次：取中间的元素，从前往后取元素跟中间的元素判断，当元素大于中间元素时停止；再从后往前取元素跟中间的元素判断，当元素小于中间元素时停止。去这两个元素的下标，当i <= j时，将这两个元素互换
	第2次：将新生成的两个序列再重复步骤一
	以此类推，直到不能再生成新的序列
*/
func quickAscSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			quickAscSort(arr, start, j)
		}
		if end > i {
			quickAscSort(arr, i, end)
		}
	}
}

//降序
func quickDescSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] > key {
				i++
			}
			for arr[j] < key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			quickDescSort(arr, start, j)
		}
		if end > i {
			quickDescSort(arr, i, end)
		}
	}
}

// 堆排序 O(n*log n) 不稳定
/*
	升序排序：需先构建一个大顶堆，然后将堆顶元素和末尾元素互换，即将最大的元素放到末尾。之后根据前n-1个元素构建大顶堆，再互换；再取前n-2个元素。。。直到只剩一个元素，即升序排序成功。
	第1次：首先将未排序的序列构成一个大顶堆，然后交换序列收尾两个元素，即将该序列中的最大值移动到末尾
	第2次：将除最后一个元素的序列再构成一个大顶堆，然后交换该序列的收尾两个元素，同样将新序列中的最大值移动到末尾
	以此类推，直到序列只剩一个元素
*/
func HeapSort(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}

	buildMaxHeap(a, l)
	for i := l - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		adjustHeap(a, 0, i)
	}

	fmt.Println("heap:\n", a)
	return a
}

func buildMaxHeap(a []int, n int) {
	for i := n/2 - 1; i >= 0; i-- { // n/2 - 1 堆中最后一个父节点
		adjustHeap(a, i, n)
	}
}

func adjustHeap(a []int, i, n int) {
	left := 2*i + 1  // 左子节点
	right := 2*i + 2 // 右子节点
	largest := i     // 父节点

	if left < n && a[left] > a[largest] {
		largest = left
	}
	if right < n && a[right] > a[largest] {
		largest = right
	}
	if largest != i {
		a[largest], a[i] = a[i], a[largest]
		adjustHeap(a, largest, n)
	}
}

// 计数排序 O(n+k) 稳定
/*
	第1次：选出未排序的序列中的最大值max，初始化一个长度为max+1的数组，每个元素对应的值为未排序的序列中值为新数组下标的个数。
	第2次：根据新数组的值将相应数量的下标从前向后赋值到未排序的序列中
*/
func CountSort(a []int) []int {
	l := len(a)
	if l <= 1 {
		return a
	}

	max := max(a)
	countSort(a, max)

	fmt.Println("count:\n", a)
	return a
}

func max(a []int) int {
	var max = 0
	for k := range a {
		if a[k] > max {
			max = a[k]
		}
	}

	return max
}

func countSort(a []int, max int) {
	arr := make([]int, max+1)
	index := 0

	for i := 0; i < len(a); i++ {
		arr[a[i]] += 1
	}

	for j := 0; j < max+1; j++ {
		for arr[j] > 0 {
			a[index] = j
			index++
			arr[j]--
		}
	}
}

// 桶排序 O(n+k)
/**/
func BucketSort(a []int) []int {
	return a
}

// 基数排序 O(n*k)
/**/
func BaseSort(a []int) []int {
	return a
}
