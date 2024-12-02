package drivers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/memory/v2"
	"github.com/spf13/viper"
	"time"
)

// NewMemory 内存
func NewMemory() fiber.Storage {
	interval := viper.GetInt64("cache.options.memory.gc_interval")

	return memory.New(memory.Config{
		GCInterval: time.Duration(interval) * time.Second,
	})
}
