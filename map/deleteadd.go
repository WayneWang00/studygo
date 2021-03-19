package main

import (
	"fmt"
	"sync"
)

func main() {
	//deleteKey()
	//OnceDelete()
	addKey()
}

func deleteKey() {
	var pc map[string]string
	pc = make(map[string]string)
	pc["qingdao"] = "青岛"
	pc["jinan"] = "济南"
	pc["yantai"] = "烟台"
	delete(pc, "qingdao")
	qingdao, ok := pc["qingdao"]
	if ok {
		fmt.Println(qingdao)
	} else {
		fmt.Println("元素不存在")
	}
	var map1 = map[string]int{"key1": 100, "key2": 200, "key3": 300}
	for k, v := range map1 {
		fmt.Println(k, v)
		if k == "key2" {
			delete(map1, "key2")
		}
		if k == "key3" {
			map1["key4"] = 400
		}
	}
	fmt.Println(map1)
}

func OnceDelete() {
	var m = map[int]int{1: 1, 2: 2, 3: 3}
	var once sync.Once

	for k := range m {
		once.Do(func() {
			for _, v := range []int{1, 2, 3} {
				if k != v {
					fmt.Println("k:", k, " v:", v)
					delete(m, v)
					break
				}
			}
		})

		fmt.Println(k, m[k])
	}

	fmt.Println("m:", m)
}

func addKey() {
	for i := 0; i < 10; i++ {
		create()
	}
}

func create() {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	for i := range m {
		m[4] = 4
		fmt.Println(i, m[i]) // 有时候会遍历到4 4，有时候遍历不到。跟map内部实现有关
	}

	fmt.Println(m)
}
