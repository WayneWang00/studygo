package main

import "fmt"

const (
	First = 1<<iota
	Second
	Third
	Four = 1<<iota
	Five
)
func main() {
	fmt.Println(Third, Five)
}
