package set2

import (
	"bytes"
	"testing"

	"../utils"
)

func TestDecryptAESinCBC(t *testing.T) {
	input := utils.ReadFromBase64File("../data/10.txt", t)
	key := []byte("YELLOW SUBMARINE")
	iv := []byte("\x00\x00\x00")

	output := DecryptAESinCBC(input, key, iv)

	expected := []byte("I'm back and I'm ringin' the bell ")
	if !bytes.Equal(expected, output[:len(expected)]) {
		t.Fatalf("Expected %s, but got %s", expected, output[:len(expected)])
	}
}
