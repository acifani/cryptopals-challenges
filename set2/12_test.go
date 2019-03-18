package set2

import (
	"bytes"
	"testing"
)

func TestECBSuffixDecryption(t *testing.T) {
	output := ECBSuffixDecryption()

	expected := []byte("Rollin' in my 5.0")
	if !bytes.Equal(expected, output[:len(expected)]) {
		t.Fatalf("Expected %s, but got %s", expected, output[:len(expected)])
	}
}
