package hash

import (
	"crypto/sha256"
	"fmt"
)

func GetHash(key string) ([]byte, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(key))
	if err != nil {
		return []byte{}, err
	}
	return hash.Sum(nil), nil
}

func GetMapIndexFromHash(hash []byte, cap int) (int, error) {
	if hash == nil {
		return -1, fmt.Errorf("hash is nil")
	}
	checksum := sha256.Sum256(hash)
	return int(checksum[0]) % cap, nil
}
