package set1

import (
	"strings"
	"testing"

	"github.com/acifani/cryptopals-challenges/utils"
)

func TestDetectAESinECB(t *testing.T) {
	input := utils.ReadFile("../data/8.txt", t)
	inputLines := strings.Split(string(input), "\n")

	for i, line := range inputLines {
		decodedLine := HexDecode([]byte(line), t)
		output := DetectAESinECB([]byte(decodedLine), 16)
		if output {
			if i != 132 {
				t.Fail()
			}
		} else {
			if i == 132 {
				t.Fail()
			}
		}
	}
}
