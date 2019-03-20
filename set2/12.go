/*
Byte-at-a-time ECB decryption (Simple)

Copy your oracle function to a new function that encrypts buffers under ECB
mode using a consistent but unknown key
(for instance, assign a single random key, once, to a global variable).

Now take that same function and have it append to the plaintext,
BEFORE ENCRYPTING, the following string:

Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK

Base64 decode the string before appending it.
Do not base64 decode the string by hand; make your code do it.
The point is that you don't know its contents.

What you have now is a function that produces:
AES-128-ECB(your-string || unknown-string, random-key)

It turns out: you can decrypt "unknown-string" with repeated calls to the oracle function!

Here's roughly how:
1. Feed identical bytes of your-string to the function 1 at a time --- start with 1 byte ("A"),
   then "AA", then "AAA" and so on. Discover the block size of the cipher.
   You know it, but do this step anyway.
2. Detect that the function is using ECB. You already know, but do this step anyways.
3. Knowing the block size, craft an input block that is exactly 1 byte short
   (for instance, if the block size is 8 bytes, make "AAAAAAA").
   Think about what the oracle function is going to put in that last byte position.
4. Make a dictionary of every possible last byte by feeding different strings to the oracle;
   for instance, "AAAAAAAA", "AAAAAAAB", "AAAAAAAC", remembering the first block of each invocation.
5. Match the output of the one-byte-short input to one of the entries in your dictionary.
   You've now discovered the first byte of unknown-string.
6. Repeat for the next byte.
*/

package set2

import (
	"bytes"
	"encoding/base64"

	"github.com/acifani/cryptopals-challenges/set1"
)

func encryptECBWithUnknownKey(input []byte) []byte {
	key := []byte("YELLOW SUBMARINE")
	suffix := []byte(`Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK`)
	decodedSuffix := make([]byte, base64.StdEncoding.DecodedLen(len(suffix)))
	base64.StdEncoding.Decode(decodedSuffix, suffix)
	input = PadRight(append(input, decodedSuffix...), len(key))
	return EncryptAESinECB(input, key)
}

// ECBSuffixDecryption decrypts the secret appendend
// to a ECB encryption oracle
func ECBSuffixDecryption(oracle func([]byte) []byte) []byte {
	blockSize := guessBlockSize(oracle)
	if blockSize == 0 {
		panic("Couldn't guess block size")
	}

	suffixLen := len(oracle([]byte{}))
	var result []byte

	for len(result) < suffixLen {
		start := len(result)
		end := len(result) + blockSize

		for i := blockSize - 1; i >= 0; i-- {
			input := bytes.Repeat([]byte{1}, i)
			target := oracle(input)[start:end]
			encryptionMap := buildMap(oracle, input, result, start, end)
			secret := encryptionMap[string(target)]
			result = append(result, secret)
		}
	}

	return result
}

func guessBlockSize(oracle func([]byte) []byte) int {
	for i := 1; i <= 128; i++ {
		input := bytes.Repeat([]byte{1}, i*2)
		encrypted := oracle(input)
		if set1.DetectAESinECB(encrypted[:i*2], i) {
			return i
		}
	}

	return 0
}

func buildMap(oracle func([]byte) []byte, input, result []byte, start, end int) map[string]byte {
	base := append(input, result...)
	encryptionMap := make(map[string]byte)
	for i := 0; i < 256; i++ {
		b := byte(i)
		toBeEncrypted := append(base, b)
		candidate := oracle(toBeEncrypted)[start:end]
		encryptionMap[string(candidate)] = b
	}

	return encryptionMap
}
