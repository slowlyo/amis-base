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

	adminApi.Post("/settings", system.SaveSettings)
	adminApi.Get("/settings", system.Settings)
	adminApi.Get("/menus", system.Menus)
	adminApi.Get("/user", system.User)
	adminApi.Post("/login", system.Login)
}
