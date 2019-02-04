package set2

import (
	"bytes"
	"testing"
)

func TestPadRight(t *testing.T) {
	input := []byte("YELLOW SUBMARINE")
	expected := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")

	output := PadRight(input, 20)

	if !bytes.Equal(expected, output) {
		t.Fatalf("Expected %s, got %s", expected, output)
	}
}
