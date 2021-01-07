package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

var json = `[{"gameId":1001,"list":[{"vip":1,"bankruptType":0,"number":5,"grandMoney":3000,"bankruptMoney":1500},{"vip":0,"bankruptType":0,"number":3,"grandMoney":1500,"bankruptMoney":1000}]}]`

const json1 = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	value := gjson.Get(json, "0.list")
	fmt.Println("value: ", value)
	valuearr := value.Array()
	valuemap := valuearr[0].Map()
	for k, v := range valuemap {
		fmt.Println("k: ", k, "v: ", v)
	}
	fmt.Println("valuearr: ", valuearr)
	value1 := gjson.Get(json, "0.list.0.number")
	fmt.Println("value1: ", value1.Value())
	value2 := gjson.Get(json1, "name.last")
	fmt.Println("value2type: ", value2.Type)
	fmt.Println("value2: ", value2.String())
	value3 := gjson.Get(json1, "age")
	fmt.Println("value3.Index: ", value3.Index)
	jsonbyte := []byte(json)
	fmt.Println("jsonbyte: ", jsonbyte)
	raw := jsonbyte[value3.Index : value3.Index+len(value3.Raw)]
	fmt.Println("raw: ", raw)
	fmt.Println("rawstr: ", string(raw))
	raw1 := []byte(value3.Raw)
	fmt.Println("raw: ", raw1)
	fmt.Println("rawstr: ", string(raw1))
	fmt.Println("value3string: ", value3.Type.String())
	fmt.Println("value3: ", value3)
	values := gjson.GetMany(json1, "name.last", "age")
	fmt.Println("values: ", values)
	list := [2]int{}
	for i := 0; i < len(list); i++ {
		list[i] = int(values[i].Int())
		fmt.Println(i, ":", list[i])
	}
	fmt.Println(list)
	if !gjson.Valid(json) {
		fmt.Println("json error!")
	}
	fmt.Println("json 是否有效: ", gjson.Valid(json))
	valuebyte := gjson.GetBytes([]byte(json), "0.list.1.number")
	fmt.Println("valuebyte.Index: ", valuebyte.Index)
	fmt.Println("valuebyte.Type: ", valuebyte.Type)
	fmt.Println("valuebyte: ", valuebyte)
	valuebyte1 := valuebyte.Value()
	fmt.Println("valuebyte.Value: ", valuebyte1)
	parsejson := gjson.Parse(json)
	parsejson1 := parsejson.Value()
	fmt.Println("parsejson1:", parsejson1)
	fmt.Println("parsejsontype: ", parsejson.Type)
	parsearr := parsejson.Array()
	if parsejson.IsArray() {
		fmt.Println("json is a array")
	}
	fmt.Println("parsearr.len: ", len(parsearr))
	bytejson := gjson.ParseBytes([]byte(json))
	fmt.Println("bytejsontype: ", bytejson.Type)
	fmt.Println("bytejson: ", bytejson)
	json1byte := gjson.GetBytes([]byte(json1), "age")
	//fmt.Println("json1byte: ", json1byte.String())
	fmt.Println("json1byte: ", json1byte.Int())
	var rawbyte []byte
	if json1byte.Index > 0 {
		rawbyte = []byte(json1)[json1byte.Index : json1byte.Index+len(json1byte.Raw)]
		fmt.Println("rawbyte1: ", rawbyte)
		fmt.Println("rawbyte1str: ", string(rawbyte))
	} else {
		rawbyte = []byte(json1byte.Raw)
		fmt.Println("rawbyte2: ", rawbyte)
	}
}
