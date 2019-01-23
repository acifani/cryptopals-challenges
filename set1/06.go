/*
Break repeating-key XOR

There's a file here. It's been base64'd after being encrypted with repeating-key XOR.
Decrypt it.

Here's how:

1. Let KEYSIZE be the guessed length of the key; try values from 2 to (say) 40.

2. Write a function to compute the edit distance/Hamming distance between two strings.
   The Hamming distance is just the number of differing bits. The distance between:
   "this is a test" and "wokka wokka!!!" is 37. Make sure your code agrees before you proceed.

3. For each KEYSIZE, take the first KEYSIZE worth of bytes, and the second KEYSIZE worth of bytes,
   and find the edit distance between them. Normalize this result by dividing by KEYSIZE.

4. The KEYSIZE with the smallest normalized edit distance is probably the key.
   You could proceed perhaps with the smallest 2-3 KEYSIZE values.
   Or take 4 KEYSIZE blocks instead of 2 and average the distances.

5. Now that you probably know the KEYSIZE: break the ciphertext into blocks of KEYSIZE length.

6. Now transpose the blocks: make a block that is the first byte of every block,
   and a block that is the second byte of every block, and so on.

7. Solve each block as if it was single-character XOR. You already have code to do this.

8. For each block, the single-byte XOR key that produces the best looking histogram
   is the repeating-key XOR key byte for that block. Put them together and you have the key.

This code is going to turn out to be surprisingly useful later on.
Breaking repeating-key XOR ("Vigenere") statistically is obviously an academic exercise,
a "Crypto 101" thing. But more people "know how" to break it than can actually break it,
and a similar technique breaks something much more important.
*/

package set1

import (
	"encoding/base64"
	"math"
	"math/bits"
)

func BreakRepeatinKeyXOR(input []byte) []byte {
	//keySize := EstimateKeySize(input)
	var result []byte
	base64.StdEncoding.Decode(result, input)
	return result
}

func EstimateKeySize(input []byte) int {
	numOfChunks := 4
	bestKeySize := 2
	bestScore := math.Inf(1)
	for keySize := 2; keySize <= 40; keySize++ {
		chunks := make([][]byte, numOfChunks)
		for i := range chunks {
			start := keySize * i
			chunks[i] = input[start : start+keySize]
		}
		sumOfDistances := 0
		for i, a := range chunks {
			for j, b := range chunks {
				if i != j {
					sumOfDistances += HammingDistance(a, b) / keySize
				}
			}
		}
		score := float64(sumOfDistances) / float64(numOfChunks)
		if score < bestScore {
			bestKeySize = keySize
			bestScore = score
		}
	}

	return bestKeySize
}

// HammingDistance calculates the edit distance given two byte arrays of the same length
func HammingDistance(a, b []byte) int {
	distance := 0
	for i := range a {
		// Identify the different bits with the XOR and then count them
		xor := a[i] ^ b[i]
		distance += bits.OnesCount(uint(xor))
	}

	return distance
}
