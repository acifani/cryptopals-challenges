package set2

import (
	"math"
	"testing"
)

func TestDetectAESMode(t *testing.T) {
	input := []byte("YELLOW SUBMARINEYELLOW SUBMARINEYELLOW SUBMARINE")
	ecb := 0
	cbc := 0
	for i := 0; i < 1000; i++ {
		encrypted := EncryptionOracle(input)
		if DetectAESMode(encrypted) == "ECB" {
			ecb++
		} else {
			cbc++
		}
	}

	if math.Abs(float64(ecb-cbc)) > 50 {
		t.Fatalf("ECB: %v, CBC: %v", ecb, cbc)
	}
}
