package main

import "fmt"

type test struct {
	key1 int
	key2 string
	key3 []test1
	key4 int
}

type test1 struct {
	k1 int
	k2 string
	k3 string
}

func main() {
	t := test{key1: 1, key2: "key", key3: []test1{{k1: 1, k2: "111", k3: "122"}, {k1: 2, k2: "211", k3: "222"}}, key4: 0}
	t3 := t.key3
	for k, _ := range t3 {
		//v.k2 = v.k3
		t3[k].k2 = t3[k].k3
		//t.key3[k].k2 = t.key3[k].k3
	}
	fmt.Println(t)
}
