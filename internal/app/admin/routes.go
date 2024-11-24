package admin

import "github.com/gofiber/fiber/v2"

func registerRoutes(app *fiber.App) {
	adminApi := app.Group("/admin")

	adminApi.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})
}
