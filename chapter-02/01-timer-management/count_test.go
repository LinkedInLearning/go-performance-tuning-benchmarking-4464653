package count

import (
	"bufio"
	"os"
	"testing"
)

// Count unique entries in a slice of strings
func Count(s []string) int {
	m := map[string]int{}
	for _, v := range s {
		m[v]++
	}
	return len(m)
}

func LoadFile(p string) ([]string, error) {
	in, err := os.Open(p)
	if err != nil {
		return nil, err
	}

	scn := bufio.NewScanner(in)

	scn.Split(bufio.ScanWords)

	var s []string
	for scn.Scan() {
		s = append(s, scn.Text())
	}

	return s, nil
}

func BenchmarkCount(b *testing.B) {
	words, err := LoadFile("testdata/alice-in-wonderland.txt")
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		Count(allWords)
	}
}
