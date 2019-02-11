package set2

import (
	"bytes"
	"testing"

	"github.com/acifani/cryptopals-challenges/set1"
	"github.com/acifani/cryptopals-challenges/utils"
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

func TestEncryptAESinCBC(t *testing.T) {
	input := []byte("YELLOW SUBMARINE")
	key := []byte("YELLOW SUBMARINE")
	iv := bytes.Repeat([]byte("\x00"), 16)

	output := DecryptAESinCBC(EncryptAESinCBC(input, key, iv), key, iv)

	if !bytes.Equal(output, input) {
		t.Fatalf("Expected %s, but got %s", input, output)
	}
}

func TestEncryptAESinEBC(t *testing.T) {
	input := []byte("YELLOW SUBMARINE")
	key := []byte("YELLOW SUBMARINE")

	output := set1.DecryptAESinECB(EncryptAESinECB(input, key), key)

	if !bytes.Equal(output, input) {
		t.Fatalf("Expected %s, but got %s", input, output)
	}
}
