package main

import (
	"encoding/json"
	"fmt"
)

type List_Conf struct {
	Vip            int32 `json:"vip"`
	Bankrupt_type  int32 `json:"bankrupt_type"`
	Number         int32 `json:"number"`
	Grand_money    int64 `json:"grand_money"`
	Bankrupt_money int64 `json:"bankrupt_money"`
}

type Game_Conf struct {
	Game_id int32       `json:"game_id"`
	List    []List_Conf `json:"list"`
}

type AllGame_Conf []Game_Conf

func main() {
	//game_conf := `[{"game_id": 1,"list": [{"vip": 1,"type": 1,"number": 5,"grand_money": 3000,"bankrupt_money": 1500},{"vip": 0,"type": 1,"number": 3,"grand_money": 1500,"bankrupt_money": 1000}] }]`
	//var gamecof map[string]interface{}
	//json.Unmarshal([]byte(game_conf), &gamecof)
	//fmt.Println("game_id: ", gamecof["game_id"])
	//if gamecof["game_id"] == 1 {
	//	fmt.Println("list: ", gamecof["list"])
	//}
	//fmt.Println("gamecof: ", gamecof)
	//var newgamecof AllGame_Conf
	//json.Unmarshal([]byte(game_conf), &newgamecof)
	//if newgamecof.Game_id == 1 {
	//	fmt.Println("game_id: ", newgamecof.Game_id)
	//	if newgamecof.List[0].Vip == 1 {
	//		fmt.Println("vip: ", newgamecof.List[0].Vip)
	//	}
	//}
	//fmt.Println("newgamecof: ", newgamecof)
	var game_id int32 = 1
	var vip int32 = 1
	fmt.Println("list: ", F(game_id, vip))
}

func F(game_id int32, vip int32) List_Conf {
	game_conf := `[{"game_id": 1,"list": [{"vip": 1,"type": 1,"number": 5,"grand_money": 3000,"bankrupt_money": 1500},{"vip": 0,"type": 1,"number": 3,"grand_money": 1500,"bankrupt_money": 1000}] }]`
	var newgamecof AllGame_Conf
	json.Unmarshal([]byte(game_conf), &newgamecof)
	fmt.Println(len(newgamecof))
	fmt.Println("newgamecof: ", newgamecof[0])
	for i := 0; i < len(newgamecof); i++ {
		if newgamecof[i].Game_id == game_id {
			fmt.Println("game_id: ", newgamecof[i].Game_id)
			for j := 0; j < len(newgamecof[i].List); j++ {
				if newgamecof[i].List[j].Vip == vip {
					fmt.Println("vip: ", newgamecof[i].List[j])
					return newgamecof[i].List[j]
				}
			}
		}
	}
	return List_Conf{}
}
