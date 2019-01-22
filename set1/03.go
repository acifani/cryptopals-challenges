/*
Single-byte XOR cipher

The hex encoded string:
1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736

... has been XOR'd against a single character. Find the key, decrypt the message.
You can do this by hand. But don't: write code to do it for you.
How? Devise some method for "scoring" a piece of English plaintext. 
Character frequency is a good metric. Evaluate each output and choose the one with the best score.
*/

package set1

import (
	"bytes"
	"encoding/hex"
)

func scoreString(decodedString []byte, candidate rune) (int, []byte) {
	xoredString := make([]byte, len(decodedString))
	candidateByte := byte(candidate)
	for i := range decodedString {
		xoredString[i] = decodedString[i] ^ candidateByte
	}

	mostCommonLetters := []byte("etaoin shrdlu")
	commonLettersLength := len(mostCommonLetters)
	score := 0
	for _, char := range xoredString {
		rarity := -1
		index := bytes.IndexByte(mostCommonLetters, char)
		if index > -1 {
			rarity = commonLettersLength - index
		}
		score += rarity
	}
	return score, xoredString
}

// BruteForceXORCypher takes an hex encoded string that has been
// XORed against a single character and tries to decypher it
func BruteForceXORCypher(hexString string) (string, error) {
	decodedString, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	} 

	maxScore := 0
	bestMatch := []byte{}

	for i := 0; i < 256; i++ {
		candidate := rune(i)
		score, match := scoreString(decodedString, candidate)
		if score > maxScore {
			maxScore = score
			bestMatch = match
		}
	}

	return string(bestMatch), nil
}