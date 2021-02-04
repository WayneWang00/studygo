package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

func main() {
	//map2Bool()
	mapInterface()
}

func map2Bool() {

	a := map[string]interface{}{
		"acc": "111",
		"t":   make(map[string]interface{}),
	}
	if aim, ok := a["t"].(bool); ok {
		fmt.Println("bool:", aim)
	}
	if v, ok := a["t"].(map[string]interface{}); ok {
		fmt.Println("map:", v)
	}
	fmt.Println("a[t]:", reflect.TypeOf(a["t"]))
}

func mapInterface() {
	s := []byte(`{"status":200}`)
	var value map[string]interface{}

	if err := json.Unmarshal(s, &value); err != nil {
		log.Fatal("unmarshal failed:", err)
	}

	fmt.Println("status:", value["status"])
	fmt.Printf("%T\n", value["status"])
}
