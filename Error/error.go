package main

import (
	"errors"
	"fmt"
)

func main() {
	var err = errors.New("test err")
	var errNil error
	fmt.Printf("print err:%s\n", err)
	fmt.Printf("print errNil:%s\n", errNil)
}
