package cache

import (
	"amis-base/internal/pkg/cache/drivers"
	"amis-base/internal/pkg/helper"
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"time"
)

// 封装自 fiber 的 storage

// storage 实例
var cacheInstance fiber.Storage

// Bootstrap 初始化缓存
func Bootstrap() {
	driver := viper.GetString("cache.driver")

	switch driver {
	case "memory":
		cacheInstance = drivers.NewMemory()
	case "redis":
		cacheInstance = drivers.NewRedis()
	// other driver
	default:
		cacheInstance = drivers.NewMemory()
	}
}

// BuildKey 构建缓存 key
func BuildKey(key string) string {
	return viper.GetString("cache.prefix") + ":" + key
}

// Get 获取缓存数据(byte)
func Get(key string) ([]byte, error) {
	return cacheInstance.Get(BuildKey(key))
}

// GetString 获取缓存数据(string)
func GetString(key string) string {
	bytesValue, err := Get(key)
	if err != nil {
		return ""
	}

	return string(bytesValue)
}

// GetObject 获取缓存数据(object)
func GetObject[T any](key string) T {
	var value T
	bytesValue, err := Get(key)
	if err != nil {
		return value
	}

	return helper.JsonDecode[T](string(bytesValue))
}

// Set 设置缓存数据
func Set(key string, val []byte, exp time.Duration) error {
	return cacheInstance.Set(BuildKey(key), val, exp)
}

// SetString 设置缓存数据(string)
func SetString(key, value string, exp time.Duration) error {
	return Set(key, []byte(value), exp)
}

// SetObject 设置缓存数据(object)
func SetObject(key string, val interface{}, exp time.Duration) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(val)
	if err != nil {
		return err
	}
	return Set(key, buf.Bytes(), exp)
}

// Delete 删除缓存数据
func Delete(key string) error {
	return cacheInstance.Delete(BuildKey(key))
}

// Reset 重置缓存
func Reset() error {
	return cacheInstance.Reset()
}

// Close 关闭缓存
func Close() error {
	return cacheInstance.Close()
}

// Remember 缓存数据
func Remember[T any](key string, exp time.Duration, callback func() (T, error)) (T, error) {
	// 先尝试从缓存中获取数据
	var value T
	bytesValue, err := Get(key)
	if err == nil && len(bytesValue) > 0 {
		// 如果缓存中有数据，反序列化并返回
		return helper.JsonDecode[T](string(bytesValue)), nil
	}

	// 如果缓存中没有数据，执行回调函数获取数据
	value, err = callback()
	if err != nil {
		return value, err
	}

	// 执行完回调后将数据缓存
	err = SetObject(key, value, exp)
	if err != nil {
		return value, err
	}

	return value, nil
}

// RememberForever 缓存数据，不过期
func RememberForever[T any](key string, callback func() (T, error)) (T, error) {
	return Remember(key, 0, callback)
}
