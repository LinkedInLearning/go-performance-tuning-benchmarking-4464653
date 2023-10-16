package fib

type Sequence struct {
	n int
}

func New() *Sequence {
	return &Sequence{}
}

func (s *Sequence) Next() int {
	v := s.fib(s.n)
	s.n++
	return v
}

func (s *Sequence) fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1, 2:
		return 1
	case 3:
		return 2
	default:
		return s.fib(n-1) + s.fib(n-2)
	}
}
