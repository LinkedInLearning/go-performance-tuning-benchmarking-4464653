package fibonacci

import (
	"fmt"
	"testing"
)

// Our benchmark shows that for fib(n), as n increases, so does the run time!
//
// This progression doesn't seem to be linear -- it seems exponential in fact.
//
// Can we modify the Sequence type to improve performance?  Would it help if
// the function memoized previous ansers?
func BenchmarkFibonacci(b *testing.B) {
	var seq Sequence
	for n := 0; n < 50; n++ {
		b.Run(fmt.Sprintf("fib %d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				seq.N(n)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	fibonacci := []int{0, 1, 1, 2, 3, 5, 8, 13, 21}

	var seq Sequence

	for i := 0; i < len(fibonacci); i++ {
		t.Run(fmt.Sprintf("fib %d", i), func(t *testing.T) {
			if got, expected := seq.N(i), fibonacci[i]; got != expected {
				t.Fatalf("expected fib(%d) to be %d; got %d", i, expected, got)
			}
		})
	}
}
