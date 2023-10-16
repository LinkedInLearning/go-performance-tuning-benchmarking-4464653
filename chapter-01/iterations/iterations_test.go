package iterations

import (
	"testing"
	"time"
)

// Observe how long Go says it takes to execute the innner loop.
//
// Is it close to 100 milliseconds?
//
// How does b.N change?
func BenchmarkN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(100 * time.Millisecond)
	}
}
