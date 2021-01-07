package main

import "fmt"

func main() {
	a := map[string]interface{}{
		"A": "",
		"B": 0,
	}
	b := "A"
	c, err := getMapString(a, b)
	fmt.Println("c: ", c)
	fmt.Println("err: ", err)
}

func getMapString(m map[string]interface{}, key string) (string, error) {
	i, ok := m[key]
	if !ok {
		return "", fmt.Errorf("map %+v key %s not found", m, key)
	}
	s, y := i.(string)
	if !y {
		return "", fmt.Errorf("map %+v key %s is not string", m, key)
	}
	return s, nil
}
