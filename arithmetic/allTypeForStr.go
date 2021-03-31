package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	n := 3
	s := "aBcD123中文"

	//res := caseString(s)
	start := time.Now()
	res := letterCasePermutation(s)
	res666 := case666(n)
	fmt.Println(time.Now().Sub(start))

	fmt.Println(len(res), res)
	fmt.Println(len(res666), res666)
}

func caseString(s string) []string {
	s = strings.ToLower(s)
	res := []string{s}

	for i, v := range s {
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			l := len(res)
			for j := 0; j < l; j++ {
				byt := []byte(res[j])
				byt[i] = byt[i] - 'a' + 'A'
				res = append(res, string(byt))
			}
		}
	}

	return res
}

func letterCasePermutation(s string) []string {
	byt := []byte(s)
	end := len(s)

	return dfs(0, end, byt)
}

func dfs(start, end int, sub []byte) []string {
	var res []string
	if start == end {
		res = append(res, string(sub))
		return res
	}

	res = append(res, dfs(start+1, end, sub)...)

	if sub[start] >= 'a' && sub[start] <= 'z' || sub[start] >= 'A' && sub[start] <= 'Z' {
		sub[start] ^= 32
		res = append(res, dfs(start+1, end, sub)...)
	}

	return res
}

func case666(n int) []string {
	byt := make([]byte, n)
	for i := range byt {
		byt[i] = '1'
	}
	s := string(byt)
	res := []string{s}

	for i := range s {
		l := len(res)
		for j := 0; j < l; j++ {
			bt := []byte(res[j])
			for k := 0; k < 5; k++ {
				bt[i] = bt[i] + 1
				res = append(res, string(bt))
			}
		}
	}

	return res
}
