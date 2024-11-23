package app

import (
	"amis-base/internal/pkg/config"
	"amis-base/internal/pkg/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

func Bootstrap() {
	// configs
	config.Bootstrap()

	// database
	db.Bootstrap()

	// fiber
	app := fiber.New(fiber.Config{
		AppName: viper.GetString("app.name"),
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// start
	log.Fatal(app.Listen(":" + viper.GetString("app.port")))
}
