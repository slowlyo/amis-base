package api

import "github.com/gofiber/fiber/v2"

func Bootstrap(app *fiber.App) {
	// 注册路由
	registerRoutes(app)
}
