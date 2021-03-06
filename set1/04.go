/*
Detect single-character XOR

One of the 60-character strings in this file has been encrypted by single-character XOR.
Find it.
(Your code from #3 should help.)
*/

package set1

import (
	"bufio"
	"encoding/hex"
	"os"
	"strings"
)

// DetectRuneXOR reads every line from the given file path
// and XORs it against every possible character,
// returning the string with the highest score
func DetectRuneXOR(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	bestMatch := []byte{}
	maxScore := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		decodedLine, err := hex.DecodeString(scanner.Text())
		if err != nil {
			return "", err
		}

		decodedString, score := BruteForceXORCypher([]byte(decodedLine))

		if score > maxScore {
			maxScore = score
			bestMatch = decodedString
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strings.TrimRight(string(bestMatch), "\r\n"), nil
}
