package main

import (
	"bytes"
	"fmt"
)

func main() {
	strOneElem()
	byteByStr()
	printStr()
	changeStr()
}

func strOneElem() {
	//serverName := "String"
	//svr := []byte(serverName)
	//fmt.Println(svr)
	//if svr[0] < 'a' || svr[0] > 'z' {
	//	fmt.Println("srvname的首字母要小写")
	//	os.Exit(1)
	//}
	//fmt.Println(serverName)
	test := "string"
	fmt.Printf("value %T\n", test[0])
	fmt.Println(test[0])
	fmt.Println(test)
}

func byteByStr() {
	str := `{MatchId:*11775,Status:*2,LastPlayerNum:*0,SignPlayerNum:*0,PrizePool:*0,RewardPlayerNum:nil,ChipAvg:nil,ChipMax:nil,ChipMin:nil,BlindLevel:nil}`
	str1 := `RuntimeUserMatchReq{Mid:*29005,MatchId:*1166,Status:*1}`
	//str2 := `{\"Nick\":\"\xe5\x86\x8d\xe7\xbe\x8e\xe7\x9a\x84\xe8\x9d\xb4\xe8\x9d\xb6\xe6\xb2\x92\xe4\xba\x86\xe7\xbf\x85\xe8\x86\x80\xe4\xb9\x9f\xe6\x98\xaf\xe6\xaf\x9b\xe6\xaf\x9b\xe8\x9f\xb2\",\"Avatar\":\"https://pclpthpk04-static.boyaagame.com/robot/demo/33679.jpeg\"}`
	str2 := `{\"Nick\":\"\xe7\x95\xb6\xe6\x88\x91\xe4\xb8\x8d\xe5\xad\x98\xe5\x9c\xa8\xe5\x97\x8e\",\"Avatar\":\"https://pclpthpk04-static.boyaagame.com/robot/demo/33649.jpeg\"}`

	fmt.Println(len([]rune(str)))
	fmt.Println("str:", bytes.Count([]byte(str), nil)-1)
	fmt.Println("str1:", bytes.Count([]byte(str1), nil)-1)
	fmt.Println("str2:", bytes.Count([]byte(str2), nil)-1)
}

func printStr() {
	a := `"string", "字符串文字" \n`
	b := "string\n"
	c := "str"

	fmt.Println(a)
	fmt.Printf("%s", b)
	fmt.Println(c)
}

func changeStr() {
	a := "string"
	fmt.Println("before:", a, &a)
	a = "String"
	fmt.Println("after:", a, &a)
}
