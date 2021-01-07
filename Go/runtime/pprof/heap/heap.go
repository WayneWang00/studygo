package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime/pprof"
)

func main() {
	f, e := os.OpenFile("src/pprof/heap/heap.pprof", os.O_RDWR|os.O_CREATE, 644)
	if e != nil {
		fmt.Print("创建文件失败")
	}
	defer f.Close()
	read()
	pprof.WriteHeapProfile(f)

}

func read() {
	buff := make([]byte, 0)
	for i := 0; i < 10000; i++ {
		ff, ee := os.Open("src/pprof/webpprof.go")
		if ee != nil {
			fmt.Print("读取webpprof失败", ee)
		}
		b, _ := ioutil.ReadAll(ff)
		buff = append(buff, b...)
		fmt.Println(string(b))

		ff.Close()

	}
}
