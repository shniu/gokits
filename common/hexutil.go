package common

import (
	"encoding/hex"
	"errors"
)

var (
	ErrEmptyString = errors.New("empty hex string")
	ErrMissingPrefix = errors.New("hex string without 0x prefix")
)

// Encodes bytes as a hex string with 0x prefix.
func Encode(b []byte) string {
	encode := make([]byte, len(b)*2+2)
	copy(encode, "0x")
	hex.Encode(encode[2:], b)
	return string(encode)
}

func Decode(hexstring string) ([]byte, error) {
	if len(hexstring) == 0 {
		return nil, ErrEmptyString
	}

	if !has0xPrefix(hexstring) {
		return nil, ErrMissingPrefix
	}

	b, err := hex.DecodeString(hexstring[2:])
	if err != nil {
		err = errors.New("decode error")
	}
	return b, err
}

func has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}
