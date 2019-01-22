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
	xoredString := xorAgainstByte(decodedString, byte(candidate))
	score := 0
	for _, char := range xoredString {
		score += calcLetterRarity(char)
	}

	return score, xoredString
}

func xorAgainstByte(buff []byte, char byte) []byte {
	xoredString := make([]byte, len(buff))
	for i := range buff {
		xoredString[i] = buff[i] ^ char
	}

	return xoredString
}

func calcLetterRarity(char byte) int {
	mostCommonLetters := []byte("etaoin shrdlu")
	commonLettersLength := len(mostCommonLetters)
	rarity := -1
	index := bytes.IndexByte(mostCommonLetters, char)
	if index > -1 {
		rarity = commonLettersLength - index
	}
	return rarity
}

// BruteForceXORCypher takes an hex encoded string that has been
// XORed against a single character and tries to decypher it
func BruteForceXORCypher(hexString string) (string, int, error) {
	decodedString, err := hex.DecodeString(hexString)
	if err != nil {
		return "", 0, err
	}

	maxScore := 0
	var bestMatch []byte

	for i := 0; i < 256; i++ {
		candidate := rune(i)
		score, match := scoreString(decodedString, candidate)
		if score > maxScore {
			maxScore = score
			bestMatch = match
		}
	}

	return string(bestMatch), maxScore, nil
}
