package main

import (
	"fmt"
	"testing"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		need := target - num
		if j, ok := m[need]; ok {
			return []int{j, i}
		}
		m[num] = i
	}

	return []int{}
}

func TestTwoSum(t *testing.T) {

	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
	if len(res) != 2 {
		t.Fatal("twoSum test case1 failed")
	}
	if res[0] != 0 || res[1] != 1 {
		t.Fatal("twoSum test case1 failed")
	}
	t.Log("twoSum test case1 passed")

	nums = []int{3, 2, 4}
	target = 6
	res = twoSum(nums, target)
	fmt.Println(res)
	if len(res) != 2 {
		t.Fatal("twoSum test case2 failed")
	}
	if res[0] != 1 || res[1] != 2 {
		t.Fatal("twoSum test case2 failed")
	}
	t.Log("twoSum test case2 passed")

	nums = []int{3, 3}
	target = 6
	res = twoSum(nums, target)
	fmt.Println(res)
	if len(res) != 2 {
		t.Fatal("twoSum test case3 failed")
	}
	if res[0] != 0 || res[1] != 1 {
		t.Fatal("twoSum test case3 failed")
	}
	t.Log("twoSum test case3 passed")
}
