package main

import "fmt"

func main() {
	sliceForStruct()
	structForSlice()
}

type ListCof struct {
	Vip           int32
	BankruptType  int32
	Number        int32
	BankruptMoney int64
	GrandMoney    int64
}

type GameCof struct {
	GameId int32
	List   []ListCof
}

type AllGameCof []GameCof

func sliceForStruct() {
	list := new(ListCof)
	lists := make([]ListCof, 1)
	lists[0] = ListCof{1, 0, 5, 3000, 1500}
	*list = ListCof{1, 0, 5, 3000, 1500}
	//l1 := List_cof{1, 0, 5, 3000, 1500}
	//g1 := AllGameCof{{1, []ListCof{{1, 0, 5, 3000, 1500}, {0, 0, 3, 1500, 1000}}}}
	g2 := AllGameCof{{1, lists}}
	fmt.Println("g2: ", g2)
}

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

func structForSlice() {
	t := test{key1: 1, key2: "key", key3: []test1{{k1: 1, k2: "111", k3: "122"}, {k1: 2, k2: "211", k3: "222"}}, key4: 0}
	fmt.Println("before:", t)
	t3 := t.key3
	for k := range t3 {
		//v.k2 = v.k3
		t3[k].k2 = t3[k].k3
		//t.key3[k].k2 = t.key3[k].k3
	}
	fmt.Println("after:", t)
}
