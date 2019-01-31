package set1

import (
	"bytes"
	"testing"
)

func TestDecryptAESinECB(t *testing.T) {
	expected := []byte(`I'm back and I'm ringin' the bell
A rockin' on the mike while the fly girls yell
In ecstasy in the back of me
Well that's my DJ Deshay cuttin' all them Z's
Hittin' hard and the girlies goin' crazy
Vanilla's on the mike, man I'm not lazy.`)
	key := []byte("YELLOW SUBMARINE")
	input := ReadFromBase64("../data/7.txt", t)

	output := DecryptAESinECB(key, input)

	t.Log(string(output))
	if !bytes.Equal(expected, output[:len(expected)]) {
		t.Fatalf("Expected %s, but got %s", expected, output[:len(expected)])
	}
}
