package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//s := "[\"\", \"\", \"\"]"
	//s := ""
	//s := "[]"
	s := "[\"1\", \"2\", \"\"]"
	lang := []string{"a", "b", "c"}
	ret, err := sliceToMap(s, lang)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	fmt.Printf("ret Value: %+v\n", ret)
}

func sliceToMap(s string, lang []string) (map[string]string, error) {
	if s == "" || s == "[]" {
		return nil, nil
	}
	var s1 []string
	err := json.Unmarshal([]byte(s), &s1)
	fmt.Printf("s1 value: %s, len: %d\n", s1, len(s1))
	if err != nil {
		return nil, err
	}
	//if len(s1) == 0 {
	//	return nil, errors.New("len(s1) == 0")
	//}
	var m = map[string]string{}
	for k, v := range lang {
		fmt.Printf("k: %d, v: %s\n", k, v)
		if len(s1) >= k {
			m[v] = s1[k]
		}
	}
	return m, nil
}
