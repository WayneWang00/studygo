package main

import "fmt"

func main() {
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
