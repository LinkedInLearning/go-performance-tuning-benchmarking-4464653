package crockford

import (
	"testing"
)

func TestNewID(t *testing.T) {
	t.Logf("NewID() = %s", NewID())
}

func TestEncode(t *testing.T) {
	tests := map[string]struct {
		input    []byte
		expected string
	}{
		"empty string": {},
		"simple": {
			input:    []byte("The quick brown fox jumps over the lazy dog."),
			expected: "AHM6A83HENMP6TS0C9S6YXVE41K6YY10D9TPTW3K41QQCSBJ41T6GS90DHGQMY90CHQPEBG=",
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got, expected := Encode(test.input), test.expected; got != expected {
				t.Fatalf("Expected Encode(%q) to return %q; got %q", test.input, expected, got)
			}
		})
	}
}
