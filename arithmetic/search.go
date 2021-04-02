package main

import "fmt"

func main() {
	//a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//fmt.Println(binarySearchRecursive(a, 2))
	//fmt.Println(binarySearch(a, 2))

	//b := []int{1, 1, 1, 2, 2, 3, 3, 4, 5, 6, 6}
	//fmt.Println(binarySearchFirst(b, 3))
	//fmt.Println(binarySearchLast(b, 3))

	c := []int{1, 1, 2, 2, 3, 3, 5, 5, 6, 7, 8}
	//fmt.Println(binarySearchFirstGT(c, 3))
	fmt.Println(binarySearchLastLT(c, 4))

	a := 3
	b := 9
	fmt.Println((a+b)/2, a+(b-a)>>1)
}

/*
	二分查找
	取中间元素跟要查找的元素比较。若大于，则取前半部分再进行之前的操作；若小于，则取后半部分再进行之前的操作；若等于，则返回对应的下标。
*/
func binarySearchRecursive(a []int, v int) int {
	l := len(a)
	if l < 1 {
		return -1
	}

	return bs(a, v, 0, l-1)
}

func bs(a []int, v int, low, high int) int {
	if low > high {
		return -1
	}
	var ret int
	mid := (low + high) / 2

	if a[mid] == v {
		return mid
	} else if a[mid] > v {
		ret = bs(a, v, low, mid-1)
	} else {
		ret = bs(a, v, mid+1, high)
	}

	return ret
}

// 非递归二分查找
func binarySearch(a []int, v int) int {
	l := len(a)
	if l < 1 {
		return -1
	}

	low := 0
	high := l - 1
	for low <= high {
		mid := (low + high) / 2
		if a[mid] == v {
			return mid
		} else if a[mid] > v {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

// 查询第一个该元素出现的位置
func binarySearchFirst(a []int, v int) int {
	l := len(a)
	if l < l {
		return -1
	}

	low := 0
	high := l - 1
	for low <= high {
		mid := low + (high-low)>>1
		if a[mid] > v {
			high = mid - 1
		} else if a[mid] < v {
			low = mid + 1
		} else {
			if mid == 0 || a[mid-1] != v {
				return mid
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

// 查询最后一个该元素出现的位置
func binarySearchLast(a []int, v int) int {
	l := len(a)
	if l < 1 {
		return -1
	}

	low := 0
	high := l - 1
	for low <= high {
		mid := low + (high-low)>>1
		if a[mid] > v {
			high = mid - 1
		} else if a[mid] < v {
			low = mid + 1
		} else {
			if mid == l-1 || a[mid+1] != v {
				return mid
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}

// 查询第一个大于该元素的位置
func binarySearchFirstGT(a []int, v int) int {
	l := len(a)
	if l < 1 {
		return -1
	}

	low := 0
	high := l - 1
	for low <= high {
		mid := low + (high-low)>>1
		if a[mid] > v {
			high = mid - 1
		} else if a[mid] < v {
			low = mid + 1
		} else {
			if mid != l-1 && a[mid+1] > v {
				return mid + 1
			} else {
				low = mid + 1
			}
		}
	}

	return -1
}

// 查询最后一个小于该元素的位置
func binarySearchLastLT(a []int, v int) int {
	l := len(a)
	if l < 1 {
		return -1
	}

	low := 0
	high := l - 1
	for low <= high {
		mid := low + (high-low)>>1
		if a[mid] > v {
			high = mid - 1
		} else if a[mid] < v {
			low = mid + 1
		} else {
			if mid != 0 && a[mid-1] < v {
				return mid - 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}
