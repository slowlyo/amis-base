package admin

import (
	"amis-base/internal/app/admin/handlers"
	"github.com/gofiber/fiber/v2"
)

func registerRoutes(app *fiber.App) {
	adminApi := app.Group("/admin-api")

	system := handlers.System{}

	adminApi.Post("/settings", system.SaveSettings)
	adminApi.Get("/settings", system.Settings)
	adminApi.Get("/menus", system.Menus)
	adminApi.Get("/user", system.User)
	adminApi.Post("/login", system.Login)
}
