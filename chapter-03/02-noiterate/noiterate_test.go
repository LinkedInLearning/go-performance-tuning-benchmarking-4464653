package noiterate

import (
	"math"
	"math/rand"
	"testing"
)

//go:noinline
func Hypot(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}

func BenchmarkWithoutIterating(b *testing.B) {
	x, y := rand.Float64(), rand.Float64()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Hypot(x,y)
	}
}
