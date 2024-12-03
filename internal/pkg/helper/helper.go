package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"os"
	"path/filepath"
)

// Sha256Hash sha256 加密
func Sha256Hash(str string) string {
	sha := sha256.New()

	sha.Write([]byte(str))

	return hex.EncodeToString(sha.Sum(nil))
}

// BcryptString 加密
func BcryptString(str string) string {
	hashedStr, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)

	if err != nil {
		return ""
	}

	return string(hashedStr)
}

// MakeDir 创建目录
func MakeDir(path string) {
	path = filepath.Dir(path)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.MkdirAll(path, os.ModePerm)
	}
}
