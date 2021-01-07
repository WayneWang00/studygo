package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func main() {
	f, e := os.OpenFile("src/pprof/testAdd/test.pprof", os.O_RDWR|os.O_CREATE, 644)
	if e != nil {
		fmt.Print("创建文件失败")
	}
	defer f.Close()

	p := pprof.NewProfile("test")
	add(p)
	p.Add("cpu", 0)
	err := pprof.Lookup("test").WriteTo(f, 2)
	if err != nil {
		fmt.Println("mutex 文件写入错误")
	}

}
func add(p *pprof.Profile) {
	p.Add("heap", 0)
}
