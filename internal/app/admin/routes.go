package admin

import (
	"amis-base/internal/app/admin/handlers"
	"amis-base/internal/app/admin/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func registerRoutes(app *fiber.App) {
	adminApi := app.Group(viper.GetString("admin.api_prefix"))

	// 认证中间件
	adminApi.Use(middleware.Auth)

	// 系统相关
	systemHandler := handlers.AdminSystem{}
	adminApi.Get("/settings", systemHandler.Settings)      // 获取系统配置
	adminApi.Post("/settings", systemHandler.SaveSettings) // 保存配置
	adminApi.Get("/menus", systemHandler.Menus)            // 获取菜单
	adminApi.Get("/user", systemHandler.User)              // 获取用户信息
	adminApi.Post("/login", systemHandler.Login)           // 登录
	adminApi.Get("/logout", systemHandler.Logout)          // 退出登录
	adminApi.Get("/pageSchema", systemHandler.PageSchema)  // 获取页面结构

	// 系统内置功能
	systemApi := adminApi.Group("/system")

	// 页面管理
	systemApi.Route("/pages", func(router fiber.Router) {
		handler := handlers.AdminPage{}

		router.Get("", handler.Index)                     // 列表
		router.Put("", handler.Update)                    // 修改
		router.Post("", handler.Store)                    // 新增
		router.Delete("", handler.Destroy)                // 删除
		router.Get("/edit", handler.Edit)                 // 编辑
		router.Get("/detail", handler.Show)               // 详情
		router.Post("/bulk-delete", handler.DestroyBatch) // 批量删除
	})
}
