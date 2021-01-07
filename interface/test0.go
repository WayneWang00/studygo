package main

import "fmt"

type any interface{}

func main() {
	var sli = make([]any, 3)
	sli[0] = 1
	sli[1] = "hello world"
	sli[2] = true
	fmt.Println(sli)

	newIo := new(MyIO)
	s := "hello world"
	ioPrint(newIo, []byte(s))
}

type Write interface {
	Write(p []byte) (n int, err error)
}

type Read interface {
	Read(p []byte) (n int, err error)
}

type ReadWrite interface {
	Read
	Write
}

type MyIO struct{}

func (io *MyIO) Read(p []byte) (n int, err error) {
	s := string(p)
	return len(s), nil
}

func (io *MyIO) Write(p []byte) (n int, err error) {
	s := string(p)
	return len(s), nil
}

func ioPrint(rw ReadWrite, p []byte) {
	fmt.Println(rw.Read(p))
	fmt.Println(rw.Write(p))
}
