package set1

import (
	"testing"
)

func TestBruteForceXORCypher(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expected := "Cooking MC's like a pound of bacon"

	output, _ := BruteForceXORCypher(input)

	if output != expected {
		t.Fatalf("Expected %v, but got %v", expected, output)
	}
}