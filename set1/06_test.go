package set1

import (
	"bytes"
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

func TestBreakRepeatinKeyXOR(t *testing.T) {
	expected := []byte(`I'm back and I'm ringin' the bell 
A rockin' on the mike while the fly girls yell 
In ecstasy in the back of me 
Well that's my DJ Deshay cuttin' all them Z's 
Hittin' hard and the girlies goin' crazy 
Vanilla's on the mike, man I'm not lazy.`)
	encryptedInput, _ := ioutil.ReadFile("../data/6.txt")
	input := make([]byte, base64.StdEncoding.DecodedLen(len(encryptedInput)))
	base64.StdEncoding.Decode(input, encryptedInput)

	output := BreakRepeatingKeyXOR(input)

	if !bytes.Equal(expected, output[:len(expected)]) {
		t.Fatalf("Expected %s, but got %s", expected, output[:len(expected)])
	}
}
