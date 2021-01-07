package main

import "runtime"
import "log"

func main() {
	test()
}

func test() {
	test2()
}

func test2() {
	pc, file, line, ok := runtime.Caller(2)
	log.Println(pc)   //栈标识符
	log.Println(file) //文件名
	log.Println(line) //该调用在文件中的行数
	log.Println(ok)   //是否获取信息成功
	f := runtime.FuncForPC(pc)
	log.Println(f.Name())

	pc, file, line, ok = runtime.Caller(0)
	log.Println(pc)
	log.Println(file)
	log.Println(line)
	log.Println(ok)
	f = runtime.FuncForPC(pc)
	log.Println(f.Name())

	pc, file, line, ok = runtime.Caller(1)
	log.Println(pc)
	log.Println(file)
	log.Println(line)
	log.Println(ok)
	f = runtime.FuncForPC(pc)
	log.Println(f.Name())
}
