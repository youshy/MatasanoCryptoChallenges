package challenge

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// Challenge 01
func Solve(input []byte) ([]byte, error) {
	dec, err := decodeHex(input)
	if err != nil {
		fmt.Printf("failed to decode hex: %s", err)
		return nil, err
	}

	f := base64Encode(dec)

	fmt.Printf("%s\n", f)

	return f, nil
}

func decodeHex(input []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(dst, input)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func base64Encode(input []byte) []byte {
	eb := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(eb, input)

	return eb
}
