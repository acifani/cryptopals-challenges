package set1

import (
	"bytes"
	"testing"

	"../utils"
)

func TestDecryptAESinECB(t *testing.T) {
	expected := []byte("I'm back and I'm ringin' the bell")
	key := []byte("YELLOW SUBMARINE")
	input := utils.ReadFromBase64File("../data/7.txt", t)

	output := DecryptAESinECB(key, input)

	if !bytes.Equal(expected, output[:len(expected)]) {
		t.Fatalf("Expected %s, but got %s", expected, output[:len(expected)])
	}
}
