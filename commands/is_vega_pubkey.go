package commands

import "encoding/hex"

const vegaPubkeyLen = 64

// IsVegaPubkey check if a string is a valid vega public vega public key.
// A vega public key is a string of 64 characters containing only hexadecimal characters.
func IsVegaPubkey(pk string) bool {
	var (
		len    = len(pk)
		_, err = hex.DecodeString(pk)
	)
	return len == 64 && err == nil
}
