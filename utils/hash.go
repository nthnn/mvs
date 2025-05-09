package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"os"
)

func Hash(data []byte) string {
	hash := sha512.New()
	hash.Write(data)

	return hex.EncodeToString(hash.Sum(nil))
}

func HashFile(path string) string {
	data, _ := os.ReadFile(path)
	return Hash(data)
}
