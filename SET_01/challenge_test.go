package challenge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChallenge01(t *testing.T) {
	input := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	expected := []byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t")

	res, err := Solve(input)

	assert.Equal(t, res, expected, "they should be equal")
	assert.NoError(t, err, "there should be no error")
}
