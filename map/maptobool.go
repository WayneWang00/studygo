package main

import "fmt"

func main() {
	a := map[string]interface{}{
		"acc": "111",
		"t":   make(map[string]interface{}),
	}
	if aim, ok := a["t"].(bool); ok {
		fmt.Println("Success", aim)
	}
}
