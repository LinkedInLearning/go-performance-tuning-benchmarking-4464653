package benchmark

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

/*
* execute barebones benchmark with a count of 1.

* observe how many test files are left.

* try using t.Cleanup() to os.RemoveAll() entire path
 b.Cleanup(func() {
	 if err := os.RemoveAll(dir); err != nil {
		 b.Fatalf("failed to remove temp directory %q: %v", dir, err)
	 }
 })

* execute barebones benchmark with a count of 20.  observe if time it takes increases

* try closing the file

* remember that every iteration is re-using the same directory.  try creating a new directory every iteration with:

	dir, err := os.MkdirTemp("./testdata", "create*")
	if err != nil {
		b.Fatal(err)
	}

* demonstrate b.ResetTimer()

* demonstrate b.StopTimer() and b.StartTimer()

* demonstrate interleved Stop/StartTimer operation

*/

func BenchmarkCreate(b *testing.B) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	dir := "./testdata/create"

	if err := os.MkdirAll(dir, 0o755); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var bytes [12]byte
		rng.Read(bytes[:])
		fn := fmt.Sprintf("%x", bytes[:])
		_, err := os.Create("./testdata/create/" + fn)
		if err != nil {
			b.Fatal(err)
		}
	}
}
