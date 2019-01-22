/*
Convert hex to base64

The string:
49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d

Should produce:
SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t

So go ahead and make that happen. You'll need to use this code for the rest of the exercises.
*/

package set1

import (
	"encoding/base64"
	"encoding/hex"
)

// HexToBase64 takes an hex encoded string and transforms it to a base64 encoded byte array
func HexToBase64(hexString string) ([]byte, error) {
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}

	base64Bytes := make([]byte, base64.StdEncoding.EncodedLen(len(hexBytes)))
	base64.StdEncoding.Encode(base64Bytes, hexBytes)

	return base64Bytes, nil
}
