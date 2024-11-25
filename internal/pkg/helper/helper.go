package helper

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256Hash(str string) string {
	sha := sha256.New()

	sha.Write([]byte(str))

	return hex.EncodeToString(sha.Sum(nil))
}
