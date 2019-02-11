/*
An ECB/CBC detection oracle

Now that you have ECB and CBC working:

Write a function to generate a random AES key; that's just 16 random bytes.

Write a function that encrypts data under an unknown key --- that is,
a function that generates a random key and encrypts under it.

The function should look like:

encryption_oracle(your-input)
=> [MEANINGLESS JIBBER JABBER]

Under the hood, have the function append 5-10 bytes (count chosen randomly)
before the plaintext and 5-10 bytes after the plaintext.

Now, have the function choose to encrypt under ECB 1/2 the time,
and under CBC the other half (just use random IVs each time for CBC).
Use rand(2) to decide which to use.

Detect the block cipher mode the function is using each time.
You should end up with a piece of code that, pointed at a block box that might
be encrypting ECB or CBC, tells you which one is happening.
*/

package set2

import (
	crypto_rand "crypto/rand"
	math_rand "math/rand"

	"github.com/acifani/cryptopals-challenges/set1"
)

func generateRandomKey(n int) []byte {
	output := make([]byte, n)
	_, err := crypto_rand.Read(output)

	if err != nil {
		panic(err)
	}

	return output
}

// EncryptionOracle randomly encrypts the input with either CBC or ECB
// with random key and IV
func EncryptionOracle(input []byte) []byte {
	prefix := generateRandomKey(5 + math_rand.Intn(6))
	suffix := generateRandomKey(5 + math_rand.Intn(6))
	randInput := append(append(prefix, input...), suffix...)
	paddedInput := PadRight(randInput, 16)

	key := generateRandomKey(16)
	if math_rand.Intn(2) == 0 {
		iv := generateRandomKey(16)
		return EncryptAESinCBC(paddedInput, key, iv)
	}

	return EncryptAESinECB(paddedInput, key)
}

// DetectAESMode will return "ECB" is the given string is encrypted in
// ECB mode, CBC otherwise. Input must be at least 32 byte long
func DetectAESMode(input []byte) string {
	for i := 0; i < len(input); i++ {
		paddedInput := PadRight(input[i:], 16)
		if set1.DetectAESinECB(paddedInput, 16) {
			return "ECB"
		}
	}
	return "CBC"
}
