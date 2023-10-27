package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoredPollKeyPrefix is the prefix to retrieve all StoredPoll
	StoredPollKeyPrefix = "StoredPoll/value/"
)

// StoredPollKey returns the store key to retrieve a StoredPoll from the index fields
func StoredPollKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
