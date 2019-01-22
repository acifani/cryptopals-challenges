package set1

import (
	"bytes"
	"testing"
)

func TextFixedXOR(t *testing.T) {
	inputA := []byte("1c0111001f010100061a024b53535009181c")
	inputB := []byte("686974207468652062756c6c277320657965")
	expected := []byte("746865206b696420646f6e277420706c6179")

	output, _ := FixedXOR(inputA, inputB)

	if !bytes.Equal(expected, output) {
		t.Fatalf("Expected %v, but got %v", expected, output)
	}
}