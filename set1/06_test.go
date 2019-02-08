package set1

import (
	"bytes"
	"testing"

	"github.com/acifani/cryptopals-challenges/utils"
)

func TestEstimateKeySize(t *testing.T) {
	expected := 29
	input := utils.ReadFromBase64File("../data/6.txt", t)

	output := EstimateKeySize(input)

	if output != expected {
		t.Fatalf("Expected %v, but got %v", expected, output)
	}
}

func TestHammingDistance(t *testing.T) {
	expected := 37
	inputA := []byte("this is a test")
	inputB := []byte("wokka wokka!!!")

	output := HammingDistance(inputA, inputB)

	if output != expected {
		t.Fatalf("Expected %v, but got %v", expected, output)
	}
}

func TestBreakRepeatinKeyXOR(t *testing.T) {
	expected := []byte("I'm back and I'm ringin' the bell")
	input := utils.ReadFromBase64File("../data/6.txt", t)

	output := BreakRepeatingKeyXOR(input)

	if !bytes.Equal(expected, output[:len(expected)]) {
		t.Fatalf("Expected %s, but got %s", expected, output[:len(expected)])
	}
}
