package set1

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestFixedXOR(t *testing.T) {
	inputA := HexDecode([]byte("1c0111001f010100061a024b53535009181c"), t)
	inputB := HexDecode([]byte("686974207468652062756c6c277320657965"), t)
	expected := HexDecode([]byte("746865206b696420646f6e277420706c6179"), t)

	output, _ := FixedXOR(inputA, inputB)

	if !bytes.Equal(expected, output) {
		t.Fatalf("Expected %v, but got %v", expected, output)
	}
}

func HexDecode(input []byte, t *testing.T) []byte {
	output := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(output, input)
	if err != nil {
		t.Fatal(err)
	}

	return output
}
