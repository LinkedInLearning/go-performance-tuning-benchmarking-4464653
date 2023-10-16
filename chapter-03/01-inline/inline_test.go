package inline

import (
	"math/rand"
	"testing"
	"runtime"
)

func Add(x, y int) int {
	return x + y
}

func BenchmarkAdd(b *testing.B) {
	x, y := rand.Int(), rand.Int()
	var sum int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum += Add(x, y)
	}
	runtime.KeepAlive(&sum)
}
