package challenge

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

// Challenge 01
func Solve01(input []byte) ([]byte, error) {
	dec, err := decodeHex(input)
	if err != nil {
		fmt.Printf("failed to decode hex: %s", err)
		return nil, err
	}

	res := base64Encode(dec)

	fmt.Printf("%s\n", res)

	return res, nil
}

// Challenge 02
func Solve02(input1, input2 []byte) ([]byte, error) {
	dec1, err := decodeHex(input1)
	if err != nil {
		fmt.Printf("failed to decode hex: %s", err)
		return nil, err
	}

	dec2, err := decodeHex(input2)
	if err != nil {
		fmt.Printf("failed to decode hex: %s", err)
		return nil, err
	}

	res, err := fixedXorDecrypt(dec1, dec2)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s\n", res)

	return res, nil
}

//
// Shared functions
//

func decodeHex(input []byte) ([]byte, error) {
	dec := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(dec, input)
	if err != nil {
		return nil, err
	}
	return dec, nil
}

func encodeHex(input []byte) []byte {
	eb := make([]byte, hex.EncodedLen(len(input)))
	hex.Encode(eb, input)

	return eb
}

func base64Encode(input []byte) []byte {
	eb := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(eb, input)

	return eb
}

func fixedXorDecrypt(input1, input2 []byte) ([]byte, error) {
	if len(input1) != len(input2) {
		return nil, errors.New("inputs have mismatched lenghts")
	}
	dec := make([]byte, len(input1))
	for i := 0; i < len(input1); i++ {
		dec[i] = input1[i] ^ input2[i]
	}

	fmt.Printf("FixedXORDecrypt %s\n", dec)
	return dec, nil
}
