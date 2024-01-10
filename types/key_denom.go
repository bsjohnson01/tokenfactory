package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	DenomKeyPrefix = "Denom/value/"
)

func DenomKey(denom string) []byte {
	var key []byte

	denomBytes := []byte(denom)
	key = append(key, denomBytes...)
	key = append(key, []byte("/")...)

	return key
}
