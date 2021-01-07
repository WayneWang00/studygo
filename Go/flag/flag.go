package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//testXxx()
	testXxxVar()
	//testNginx()
}

func testXxx() {
	var a = flag.Bool("a", false, "bool类型参数")
	var b = flag.String("b", "", "string类型参数")
	var c = flag.Int("c", 0, "int类型参数")

	flag.Parse()
	fmt.Println("a:", *a)
	fmt.Println("b:", *b)
	fmt.Println("c:", *c)
	fmt.Println("other:", flag.Args())
}

func testXxxVar() {
	var a bool
	var b string
	var c int

	flag.BoolVar(&a, "a", false, "bool类型参数")
	flag.StringVar(&b, "b", "", "string类型参数")
	flag.IntVar(&c, "c", 0, "int类型参数")

	flag.Parse()
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("len(args):", flag.NArg())
	fmt.Println("other:", flag.Args())
	for k, v := range flag.Args() {
		fmt.Printf("args %d:%s\t arg[%d]:%s\n", k, v, k, flag.Arg(k))
	}
}

var (
	h bool

	v, V bool
	t, T bool
	q    *bool

	s string
	p string
	c string
	g string
)

func initArg() {
	flag.NewFlagSet("testNginx", flag.ExitOnError)

	flag.BoolVar(&h, "h", false, "this help")

	flag.BoolVar(&v, "v", false, "show version and end")
	flag.BoolVar(&V, "V", false, "show version and configure options then exit")

	flag.BoolVar(&t, "t", false, "test configuration and exit")
	flag.BoolVar(&T, "T", false, "test configuration, dump it and exit")

	// 另一种绑定方式
	q = flag.Bool("q", false, "suppress non-error message during configuration testing")

	// 默认值是 -s string，有了`signal`之后，变成为 -s signal
	flag.StringVar(&s, "s", "", "send `signal` to master process: stop, quit, reopen, reload")
	flag.StringVar(&p, "p", "/usr/local/nginx/", "set `prefix` path")
	flag.StringVar(&c, "c", "conf/nginx.conf", "set configuration `file`")
	flag.StringVar(&g, "g", "conf/nginx.conf", "set global `directive` out of configuration file")

	// 改变默认的Usage，flag中的Usage是函数类型。
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.16.1
Usage: nginx [-?hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directive]

Options:
`)

	flag.PrintDefaults()
}

func testNginx() {
	initArg()

	flag.Parse()
	if h {
		flag.Usage()
	}
}
