package rand

import (
	cryptorand "crypto/rand"
	mathrand "math/rand"
	"testing"
)

func BenchmarkRandRead(b *testing.B) {
	funcs := map[string]func([]byte) (int, error){
		"mathrand":   mathrand.Read,
		"cryptorand": cryptorand.Read,
	}

	for name, f := range funcs {
		buf := make([]byte, 128)
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(buf)
			}
		})
	}
}

func BenchmarkRandReadInterleved(b *testing.B) {
	funcs := map[string]func([]byte) (int, error){
		"mathrand":   mathrand.Read,
		"cryptorand": cryptorand.Read,
	}

	for name, f := range funcs {
		buf := make([]byte, 128)
		b.Run(name, func(b *testing.B) {
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				b.StartTimer()
				f(buf)
				b.StopTimer()
			}
		})
	}
}
