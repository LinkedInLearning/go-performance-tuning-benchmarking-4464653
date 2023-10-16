package ratio

import (
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:noinline
func Add(x, y int) int {
	return x + y
}

func BenchmarkBadRatio(b *testing.B) {
	// walk a file system, open a text file, read the contents, then benchmark adding them.

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		root := os.DirFS("testdata/numbers")
		fs.WalkDir(root, ".", func(p string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				return nil
			}

			in, err := root.Open(p)
			if err != nil {
				b.Fatal(err)
			}
			defer in.Close()

			var x, y int
			if n, err := fmt.Fscanf(in, "%d, %d\n", &x, &y); n != 2 || err != nil {
				b.Fatalf("could not read integer pair from %q", p)
			}

		b.StartTimer()
			Add(x, y)
			b.StopTimer()
			return nil
		})
	}
}
