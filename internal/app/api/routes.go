package api

import (
	"amis-base/internal/app/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func BootRoutes(app *fiber.App) {
	api := app.Group("/api")

	index := handlers.Index{}
	api.Get("/hello", index.Hello)
}
