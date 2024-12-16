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

// Bootstrap 初始化
func Bootstrap() {
	// 加载配置
	config.Bootstrap()

	// 数据库
	db.Bootstrap()

	// 缓存
	cache.Bootstrap()
}

// Start 启动服务
func Start() {
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

	// 前端
	app.Static("/", "./web/dist")

	// 文件上传目录
	app.Static("/uploads", "./assets/uploads")

	// 启动服务
	log.Fatal(app.Listen(":" + viper.GetString("app.port")))
}
