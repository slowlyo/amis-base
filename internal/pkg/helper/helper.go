package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

func Sha256Hash(str string) string {
	sha := sha256.New()

	sha.Write([]byte(str))

	return hex.EncodeToString(sha.Sum(nil))
}

func BcryptString(str string) string {
	hashedStr, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)

	if err != nil {
		return ""
	}

	return string(hashedStr)
}
