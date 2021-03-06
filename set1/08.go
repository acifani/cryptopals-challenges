/*
Detect AES in ECB mode

In this file are a bunch of hex-encoded ciphertexts.
One of them has been encrypted with ECB.
Detect it.

Remember that the problem with ECB is that it is stateless and deterministic; the same 16 byte plaintext block will always produce the same 16 byte ciphertext.
*/

package set1

// DetectAESinECB will return true if it encounters a block of the given size more than once
func DetectAESinECB(input []byte, blockSize int) bool {
	if len(input)%blockSize != 0 {
		panic("Input length not a multiple of the block size")
	}

	seenBlocks := make(map[string]int)
	for start := 0; start < len(input); start += blockSize {
		end := start + blockSize
		candidate := string(input[start:end])
		_, seen := seenBlocks[candidate]
		if seen {
			return true
		}
		seenBlocks[candidate] = 1
	}
	return false
}
