package admin

import (
	"amis-base/internal/app/admin/handlers"
	"github.com/gofiber/fiber/v2"
)

func registerRoutes(app *fiber.App) {
	adminApi := app.Group("/admin-api")

	common := handlers.Common{}

	adminApi.Post("/settings", common.SaveSettings)
	adminApi.Get("/settings", common.Settings)
}
