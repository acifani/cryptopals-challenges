package set1

import "testing"

func Test(t *testing.T) {
	expected := "Now that the party is jumping"

	output, _ := DetectRuneXOR("../data/4.txt")

	if output != expected {
		t.Fatalf("Expected %v, but got %v", expected, output)
	}
}
