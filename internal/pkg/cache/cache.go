package cache

import (
	"amis-base/internal/pkg/cache/drivers"
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
	default:
		cacheInstance = drivers.NewMemory()
	}
}

func Get(key string) ([]byte, error) {
	return cacheInstance.Get(key)
}

func Set(key string, val []byte, exp time.Duration) error {
	return cacheInstance.Set(key, val, exp)
}

func Delete(key string) error {
	return cacheInstance.Delete(key)
}

func Reset() error {
	return cacheInstance.Reset()
}

func Close() error {
	return cacheInstance.Close()
}
