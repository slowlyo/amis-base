package app

import (
	"amis-base/internal/app/admin"
	"amis-base/internal/app/api"
	"amis-base/internal/pkg/cache"
	"amis-base/internal/pkg/config"
	"amis-base/internal/pkg/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

func Bootstrap() {
	// 加载配置
	config.Bootstrap()

	// 数据库
	db.Bootstrap()

	// 缓存
	cache.Bootstrap()

	// 初始化 fiber
	app := fiber.New(fiber.Config{
		AppName:           viper.GetString("app.name"),
		CaseSensitive:     true,
		EnablePrintRoutes: viper.GetBool("app.dev"),
	})

	// 加载 admin 模块
	admin.Bootstrap(app)
	// 加载 api 模块
	api.Bootstrap(app)

	// 启动服务
	log.Fatal(app.Listen(":" + viper.GetString("app.port")))
}
