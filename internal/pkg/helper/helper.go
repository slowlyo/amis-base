package helper

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

// IsAllowRequest 判断是否允许请求
func IsAllowRequest(rule, method, originalURL string) bool {
	matchedMethod := true
	cleanRule := strings.TrimLeft(rule, "^")
	methodList := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS", "CONNECT", "TRACE"}

	// 移除请求路径中的参数以及前缀
	url := strings.TrimLeft(strings.Split(originalURL, "?")[0], "/")
	url = strings.TrimLeft(url, strings.Split(url, "/")[0])
	url = strings.TrimLeft(url, "/")

	// 校验请求方式
	for _, v := range methodList {
		prefixUpper := v + ":"
		prefixLower := strings.ToLower(prefixUpper)
		// 如果匹配上了, 则说明规则中限定了请求方式
		if strings.HasPrefix(cleanRule, prefixUpper) || strings.HasPrefix(cleanRule, prefixLower) {
			// 请求方式是否匹配
			matchedMethod = v == method
			// 去除规则中的请求方式前缀
			cleanRule = strings.TrimLeft(strings.TrimLeft(cleanRule, prefixLower), prefixUpper)
			break
		}
	}

	// 非正则匹配
	if !strings.HasPrefix(rule, "^") {
		return matchedMethod && (strings.TrimLeft(cleanRule, "/") == url)
	}

	// 正则匹配
	re, err := regexp.Compile(cleanRule)
	if err == nil {
		return matchedMethod && re.MatchString(url)
	}

	return false
}

// JsonEncode 将任意数据编码为 JSON 字符串
func JsonEncode(data any) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Errorf("Error encoding data to JSON: %v\n", err)
		return ""
	}
	return string(jsonData)
}

// JsonDecode 将 JSON 字符串解码为指定类型。如果解析失败，返回零值。
func JsonDecode[T any](data string) T {
	var result T
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		log.Errorf("Error decoding JSON data: %v\n", err)
	}
	return result
}
