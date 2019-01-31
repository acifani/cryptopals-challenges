package set1

import (
	"strings"
	"testing"
)

func TestDetectAESinECB(t *testing.T) {
	input := ReadFile("../data/8.txt", t)
	inputLines := strings.Split(string(input), "\n")

	for i, line := range inputLines {
		decodedLine := HexDecode([]byte(line), t)
		output := DetectAESinECB([]byte(decodedLine))
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
