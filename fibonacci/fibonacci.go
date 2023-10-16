package fibonacci

import "sync"

type Sequence struct {
	mu   *sync.RWMutex
	memo map[int]int
	n    int
}

func (s *Sequence) fib(n int) int {
	if v, exists := s.memo[n]; exists {
		return v
	}
	v := s.fib(n-1) + s.fib(n-2)
	s.memo[n] = v
	return v
}

func (s *Sequence) Next() int {
	s.n++
	return s.fib(s.n)
}

func (s *Sequence) N(n int) int {
	return s.fib(n)
}
