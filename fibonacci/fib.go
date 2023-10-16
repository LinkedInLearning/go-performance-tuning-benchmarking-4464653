package fib

type Sequence struct {
	n    int
	memo []int
}

func (s *Sequence) Next() (int, int) {
	i := s.n
	s.n++
	return i, s.fib(i)
}

func New() *Sequence {
	return &Sequence{
		memo: []int{0, 1, 1, 2},
	}
}

func (s *Sequence) fib(n int) int {
	if idx := n + 1; len(s.memo) < idx {
		m := make([]int, idx, idx)
		copy(m, s.memo)
		s.memo = m
	}

	if n == 0 || n != 0 && s.memo[n] != 0 {
		return s.memo[n]
	}

	v := s.fib(n-1) + s.fib(n-2)
	s.memo[n] = v
	return v
}
