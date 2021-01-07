package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var st Stack

type Stack struct {
	i    int
	data [10]int
}

func (s *Stack) Push(k int) {
	if s.i+1 > 9 {
		return
	}
	s.data[s.i] = k
	s.i++
}

func (s *Stack) Pop() (res int) {
	s.i--
	if s.i < 0 {
		s.i = 0
		return
	}
	res = s.data[s.i]
	return res
}
func main() {
	for {
		s, err := reader.ReadString('\n')
		var token string
		if err != nil {
			return
		}
		var sum int
		for _, c := range s {
			switch {
			case c >= '0' && c <= '9':
				token = token + string(c)
			case c == ' ':
				r, _ := strconv.Atoi(token)
				if r >= 1 && r <= 9 {
					st.Push(r)
				}
				token = ""
			case c == '+':
				sum = st.Pop() + st.Pop()
				fmt.Printf("sum= %d, i=%d\n", sum, st.i)
				st.Push(sum)
				fmt.Printf("%d\n", st.data)
			case c == '*':
				fmt.Printf("%d\n", st.Pop()*st.Pop())
				fmt.Printf("%d\n", st.data)
			case c == '-':
				//st.Push(sum)
				fmt.Printf("sum= %d\n", sum)
				p := st.Pop()
				q := st.Pop()
				diff := q - p
				fmt.Printf("%d\n", diff)
				st.Push(diff)
				fmt.Printf("%d\n", st.data)
			case c == 'q':
				return
			default:
			}
		}
		//fmt.Printf("%d\n", st.data)
	}
}
