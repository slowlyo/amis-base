package helper

import (
	"golang.org/x/crypto/bcrypt"
	"os"
	"path/filepath"
)

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
