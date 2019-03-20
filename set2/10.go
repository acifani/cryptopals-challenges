/*
Implement CBC mode

CBC mode is a block cipher mode that allows us to encrypt irregularly-sized messages,
despite the fact that a block cipher natively only transforms individual blocks.

In CBC mode, each ciphertext block is added to the next plaintext block before the next call to the cipher core.

The first plaintext block, which has no associated previous ciphertext block,
is added to a "fake 0th ciphertext block" called the initialization vector, or IV.

Implement CBC mode by hand by taking the ECB function you wrote earlier,
making it encrypt instead of decrypt (verify this by decrypting whatever you encrypt to test),
and using your XOR function from the previous exercise to combine them.

The file here is intelligible (somewhat) when CBC decrypted against "YELLOW SUBMARINE"
with an IV of all ASCII 0 (\x00\x00\x00 &c)
*/

package set2

import (
	"crypto/aes"

	"github.com/acifani/cryptopals-challenges/set1"
)

// DecryptAESinCBC decyphers an input encrypted via AES-128 in CBC,
// according to the given key and IV
func DecryptAESinCBC(input, key, iv []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(input)%cipher.BlockSize() != 0 {
		panic("Input length not a multiple of the block size")
	}

	if len(iv) != cipher.BlockSize() {
		panic("Wrong IV size")
	}

	blockSize := cipher.BlockSize()
	var result []byte
	prevBlock := iv
	for i := 0; i < len(input)/blockSize; i++ {
		decryptedText := make([]byte, blockSize)
		block := input[i*blockSize : (i+1)*blockSize]
		cipher.Decrypt(decryptedText, block)
		plaintext := set1.RepeatingKeyXOR(decryptedText, prevBlock)
		result = append(result, plaintext...)
		prevBlock = block
	}

	return result
}

// EncryptAESinCBC encrypts a plaintext input via AES-128 in CBC,
// according to the given key and IV
func EncryptAESinCBC(input, key, iv []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(input)%cipher.BlockSize() != 0 {
		panic("Input length not a multiple of the block size")
	}

	if len(iv) != cipher.BlockSize() {
		panic("Wrong IV size")
	}

	blockSize := cipher.BlockSize()
	result := make([]byte, len(input))
	prevBlock := iv
	for i := 0; i < len(input)/blockSize; i++ {
		block := input[i*blockSize : (i+1)*blockSize]
		xoredBlock := set1.RepeatingKeyXOR(block, prevBlock)
		cipher.Encrypt(result[i*blockSize:(i+1)*blockSize], xoredBlock)
		prevBlock = result[i*blockSize : (i+1)*blockSize]
	}

	return result
}

// EncryptAESinECB encrypts a plaintext input via AES-128 in ECB,
// according to the given key
func EncryptAESinECB(input, key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(input)%cipher.BlockSize() != 0 {
		panic("Input length not a multiple of the block size")
	}

	result := make([]byte, len(input))
	for start := 0; start < len(input); start += cipher.BlockSize() {
		cipher.Encrypt(result[start:], input[start:])
	}

	return result
}
