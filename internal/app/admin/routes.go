package admin

import (
	"amis-base/internal/app/admin/handlers"
	"amis-base/internal/app/admin/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func BootRoutes(app *fiber.App) {
	adminRoute := app.Group(viper.GetString("admin.api_prefix"))

	// 认证中间件
	adminRoute.Use(middleware.Auth)

	// 系统相关
	systemHandler := handlers.AdminSystem{}
	adminRoute.Get("/user", systemHandler.User)              // 获取用户信息
	adminRoute.Get("/menus", systemHandler.Menus)            // 获取菜单
	adminRoute.Post("/login", systemHandler.Login)           // 登录
	adminRoute.Get("/logout", systemHandler.Logout)          // 退出登录
	adminRoute.Post("/upload", systemHandler.Upload)         // 文件上传
	adminRoute.Get("/settings", systemHandler.Settings)      // 获取系统配置
	adminRoute.Get("/page_schema", systemHandler.PageSchema) // 获取页面结构
	adminRoute.Post("/settings", systemHandler.SaveSettings) // 保存配置

	// 系统内置功能
	bootSystemRoutes(adminRoute)
}

func bootSystemRoutes(adminRoute fiber.Router) {
	// 系统内置功能
	systemRoute := adminRoute.Group("/system")

	// 页面管理
	systemRoute.Route("/pages", func(router fiber.Router) {
		handler := handlers.AdminPage{}

		router.Get("", handler.Index)                 // 列表
		router.Post("", handler.Save)                 // 保存
		router.Get("/copy", handler.Copy)             // 复制
		router.Get("/detail", handler.Detail)         // 详情
		router.Post("/delete", handler.Destroy)       // 删除
		router.Post("/quick_save", handler.QuickSave) // 快速编辑
	})

	// 角色管理
	systemRoute.Route("/roles", func(router fiber.Router) {
		handler := handlers.AdminRole{}

		router.Get("", handler.Index)           // 列表
		router.Post("", handler.Save)           // 新增/修改
		router.Get("/detail", handler.Detail)   // 详情
		router.Post("/delete", handler.Destroy) // 删除
	})

	// 用户管理
	systemRoute.Route("/users", func(router fiber.Router) {
		handler := handlers.AdminUser{}

		router.Get("", handler.Index)                    // 列表
		router.Post("", handler.Save)                    // 新增/修改
		router.Get("/detail", handler.Detail)            // 详情
		router.Post("/delete", handler.Destroy)          // 删除
		router.Get("/role_options", handler.RoleOptions) // 角色选项列表
		router.Post("/quick_save", handler.QuickSave)    // 快速编辑
	})

	// 菜单管理
	systemRoute.Route("/menus", func(router fiber.Router) {
		handler := handlers.AdminMenu{}

		router.Get("", handler.Index)                        // 列表
		router.Post("", handler.Save)                        // 新增/修改
		router.Get("/detail", handler.Detail)                // 详情
		router.Post("/delete", handler.Destroy)              // 删除
		router.Post("/quick_save", handler.QuickSave)        // 快速编辑
		router.Post("/sort", handler.SaveSort)               // 排序
		router.Get("/parent_options", handler.ParentOptions) // 父级选项
		router.Get("/page_options", handler.PageOptions)     // 页面选项
	})
}
