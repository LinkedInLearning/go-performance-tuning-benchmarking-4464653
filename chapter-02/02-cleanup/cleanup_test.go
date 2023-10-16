package cleanup

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"testing"
)

func BenchmarkWriteBytes(b *testing.B) {
	b.StopTimer()

	dir, err := os.MkdirTemp("testdata", "write*")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		var buf [16]byte
		rand.Read(buf[:])
		fn := path.Join(dir, fmt.Sprintf("%x", buf[:]))

		out, err := os.Create(fn)
		if err != nil {
			b.Fatal(err)
		}
		defer out.Close()
		message := "benchmarking in go!"
		b.StartTimer()
		fmt.Fprintf(out, message)
		b.StopTimer()
		b.SetBytes(int64(len(message)))
	}
}
