package fib

import (
	"strconv"
	"testing"
)

func BenchmarkFib(b *testing.B) {
	seq := New()
	for n := 0; n < 17; n++ {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				seq.fib(n)
			}
		})
	}
}

func TestFib(t *testing.T) {
	tests := []int{0, 1, 1, 2, 3, 5, 8, 13, 21}
	seq := New()
	for i, expected := range tests {
		i, value := i, expected
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got, expected := seq.fib(i), value; got != expected {
				t.Fatalf("fib(%v) returned %v; expected %v", i, got, expected)
			}
		})
	}
}
