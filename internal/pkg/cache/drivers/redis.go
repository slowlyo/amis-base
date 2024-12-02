package drivers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/redis/v3"
	"github.com/spf13/viper"
)

// NewRedis redis
func NewRedis() fiber.Storage {
	return redis.New(redis.Config{
		Host:     viper.GetString("cache.options.redis.host"),
		Port:     viper.GetInt("cache.options.redis.port"),
		Username: viper.GetString("cache.options.redis.username"),
		Password: viper.GetString("cache.options.redis.password"),
		Database: viper.GetInt("cache.options.redis.database"),
		Reset:    viper.GetBool("cache.options.redis.reset"),
		PoolSize: viper.GetInt("cache.options.redis.pool_size"),
	})
}
