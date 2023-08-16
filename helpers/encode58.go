package helpers

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/btcsuite/btcd/btcutil/base58"
)

func Encode58(input string) (string, error) {
	if input == "" {
		return "", nil
	}
	b, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}
	hash0 := hash(b)
	hash1 := hash(hash0)
	// Since hash (sha256) never fails, the hash should always have length >4
	inputCheck := append(b, hash1[:4]...)
	result := base58.Encode(inputCheck)

	return result, nil
}

func hash(s []byte) []byte {
	h := sha256.New()
	h.Write(s)
	return h.Sum(nil)
}
