package admin

import (
	"amis-base/internal/app/admin/handlers"
	"amis-base/internal/app/admin/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func registerRoutes(app *fiber.App) {
	adminApi := app.Group(viper.GetString("admin.api_prefix"))

	adminApi.Use(middleware.Auth)

	system := handlers.System{}

	adminApi.Get("/settings", system.Settings)      // 获取系统配置
	adminApi.Post("/settings", system.SaveSettings) // 保存配置
	adminApi.Get("/menus", system.Menus)            // 获取菜单
	adminApi.Get("/user", system.User)              // 获取用户信息
	adminApi.Post("/login", system.Login)           // 登录
	adminApi.Get("/logout", system.Logout)          // 退出登录
	adminApi.Get("/pageSchema", system.PageSchema)  // 获取页面结构
}
