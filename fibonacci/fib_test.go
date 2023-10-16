package fib

import (
	"fmt"
	"testing"
)

// Can you write a benchmark that shows the performance of fib(0) - fib(10)?
//
// Once you've written a benchmark, do you have an instinct about how you could
// improve performance?
//
// If so, can you alter the fib() function and demonstrate that your new
// version is faster?
//
// Maybe this is a good place to wrie BenchmarkFib()

func BenchmarkFib(b *testing.B) {
	for n := 0; n < 20; n++ {
		b.Run(fmt.Sprintf("fib-%d", n), func(b *testing.B) {
			seq := New()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				seq.fib(n)
			}
		})
	}
}

func TestFib(t *testing.T) {
	seq := New()

	tests := []int{0, 1, 1, 2, 3, 5, 8, 13}

	for i, v := range tests {
		if got, expected := seq.fib(i), v; got != expected {
			t.Fatalf("expected seq.fib(%d) to be %d; got %d", i, expected, got)
		}
	}
}
