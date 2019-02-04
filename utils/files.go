package utils

import (
	"encoding/base64"
	"io/ioutil"
	"testing"
)

func ReadFile(path string, t *testing.T) []byte {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return file
}

func ReadFromBase64File(path string, t *testing.T) []byte {
	encryptedInput := ReadFile(path, t)
	input := make([]byte, base64.StdEncoding.DecodedLen(len(encryptedInput)))
	n, _ := base64.RawStdEncoding.Decode(input, encryptedInput)
	return input[:n]
}
