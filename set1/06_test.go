package set1

import (
	"encoding/base64"
	"io/ioutil"
	"testing"
)

func TestEstimateKeySize(t *testing.T) {
	expected := 29
	encryptedInput, _ := ioutil.ReadFile("../data/6.txt")
	input := make([]byte, base64.StdEncoding.DecodedLen(len(encryptedInput)))
	base64.StdEncoding.Decode(input, encryptedInput)

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
