package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime/pprof"
)

func main() {
	f, e := os.OpenFile("src/pprof/cpu/cpu.pprof", os.O_RDWR|os.O_CREATE, 644)

	if e != nil {
		fmt.Print("创建文件失败")
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	for i := 0; i < 100; i++ {
		ff, ee := os.Open("src/pprof/webpprof.go")
		if ee != nil {
			fmt.Print("读取webpprof失败", e)
		}
		b, _ := ioutil.ReadAll(ff)
		fmt.Println(string(b))
		ff.Close()

	}
}
