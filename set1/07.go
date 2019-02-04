/*
AES in ECB mode

The Base64-encoded content in this file has been encrypted via AES-128 in ECB mode under the key
"YELLOW SUBMARINE".
(case-sensitive, without the quotes; exactly 16 characters; I like "YELLOW SUBMARINE" because it's exactly 16 bytes long, and now you do too).

Decrypt it. You know the key, after all.

Easiest way: use OpenSSL::Cipher and give it AES-128-ECB as the cipher.
*/

package set1

import "crypto/aes"

// DecryptAESinECB decyphers an input encrypted via AES-128 in ECB,
// according to the given key
func DecryptAESinECB(input, key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(input)%cipher.BlockSize() != 0 {
		panic("Input length not a multiple of the block size")
	}

	result := make([]byte, len(input))
	for start := 0; start < len(input); start += cipher.BlockSize() {
		cipher.Decrypt(result[start:], input[start:])
	}

	return result
}
