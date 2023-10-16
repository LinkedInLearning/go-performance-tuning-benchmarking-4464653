package fibonacci

type Sequence struct {
	memo map[int]int
	n    int
}

func (s *Sequence) fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}
	return s.fib(n-1) + s.fib(n-2)
}

func (s *Sequence) Next() int {
	s.n++
	return s.fib(s.n)
}

func (s *Sequence) N(n int) int {
	return s.fib(n)
}
