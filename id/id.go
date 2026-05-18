package id

import (
	"crypto/rand"
	"fmt"
	"time"
)

// GenerateULID generates a ULID-style ID (timestamp + random).
func GenerateULID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(fmt.Sprintf("phi-utils/id: failed to generate random bytes: %v", err))
	}

	ts := time.Now().UnixMilli()
	b[0] = byte(ts >> 40)
	b[1] = byte(ts >> 32)
	b[2] = byte(ts >> 24)
	b[3] = byte(ts >> 16)
	b[4] = byte(ts >> 8)
	b[5] = byte(ts)

	return encodeHex(b)
}

func encodeHex(b []byte) string {
	const hexChars = "0123456789abcdef"
	out := make([]byte, len(b)*2)
	for i, v := range b {
		out[i*2] = hexChars[v>>4]
		out[i*2+1] = hexChars[v&0x0f]
	}
	return string(out)
}

// GenerateID generates a short unique ID with optional prefix.
func GenerateID(prefix string) string {
	id := GenerateULID()
	if prefix == "" {
		return id
	}
	return prefix + "_" + id
}
