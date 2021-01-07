package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var numberflag = flag.Bool("n", false, "number each line")

func cat(r *bufio.Reader) {
	i := 1
	for {
		buffer, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if *numberflag {
			fmt.Fprintf(os.Stdout, "%5d %s", i, buffer)
			i++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buffer)
		}
	}
	return
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, e := os.Open(flag.Arg(i))
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading form %s: %s\n", os.Args[0], flag.Arg(i), e.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}
