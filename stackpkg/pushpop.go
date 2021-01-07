package stackpkg

type Stack struct {
	i    int
	data [10]int
}

func (s *Stack) Push(k int) {
	s.data[s.i] = k
	s.i++
}

func (s *Stack) Pop() int {
	s.i--
	return s.data[s.i]
}
