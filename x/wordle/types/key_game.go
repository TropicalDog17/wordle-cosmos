package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// GameKeyPrefix is the prefix to retrieve all Game
	GameKeyPrefix = "Game/value/"
)

// GameKey returns the store key to retrieve a Game from the index fields
func GameKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
