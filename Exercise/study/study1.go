package main

import "fmt"

func main() {
	pase_student()
}
type student struct {
	Name string
}
func pase_student() {
	m := make(map[string]*student)
	stus := []student {
		{"zhou"},{"li"},{"wang"},
	}
	for _,stu :=range stus {
		m[stu.Name] = &stu
	}
	for _, v :=range m {
		fmt.Println(v.Name)
	}
}
