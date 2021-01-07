package main

import "fmt"

type List_cof struct {
	Vip            int32
	BanKrupt_type  int32
	Number         int32
	Bankrupt_money int64
	Grand_money    int64
}

type Game_cof struct {
	Game_id int32
	List    []List_cof
}

type AllGame_cof []Game_cof

func main() {
	list := new(List_cof)
	list1 := make([]List_cof, 1)
	list1[0] = List_cof{1, 0, 5, 3000, 1500}
	*list = List_cof{1, 0, 5, 3000, 1500}
	//l1 := List_cof{1, 0, 5, 3000, 1500}
	//g1 := AllGame_cof{{1, {{1, 0, 5, 3000, 1500}, {0, 0, 3, 1500, 1000}}}}
	g2 := AllGame_cof{{1, list1}}
	fmt.Println("g2: ", g2)
}
