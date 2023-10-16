package fibonacci

import (
	"fmt"
	"testing"
	"sync"
)

type Sequence struct {
	memo map[int]int
	mu *sync.RWMutex
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

// Our benchmark shows that for fib(n), as n increases, so does the run time!
//
// This progression doesn't seem to be linear -- it seems exponential in fact.
//
// Can we modify the Sequence type to improve performance?  Would it help if
// the function memoized previous answers?

func BenchmarkFibonacci(b *testing.B) {
  seq := Sequence{
		mu: &sync.RWMutex{},
		memo: map[int]int{
			0: 0,
			1: 1,
			2: 1,
			3: 2,
		},
	}
	for n := 0; n < 5; n++ {
		b.Run(fmt.Sprintf("fib %d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				seq.N(n)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	fibonacci := []int{0, 1, 1, 2, 3, 5, 8, 13, 21}

	seq := Sequence{
		mu: &sync.RWMutex{},
		memo: map[int]int{
			0: 0,
			1: 1,
			2: 1,
			3: 2,
		},
	}
	
	for i := 0; i < len(fibonacci); i++ {
		t.Run(fmt.Sprintf("fib %d", i), func(t *testing.T) {
			if got, expected := seq.N(i), fibonacci[i]; got != expected {
				t.Fatalf("expected fib(%d) to be %d; got %d", i, expected, got)
			}
		})
	}
}
