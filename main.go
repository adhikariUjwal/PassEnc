package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func generateKey(password string) []byte {
	hash := sha256.Sum256([]byte(password))
	return hash[:]
}

func main() {
	password := "monev@001"
	key := generateKey((password))

	fmt.Println("Generated Key", hex.EncodeToString(key))
}
