package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func Auth(ctx *fiber.Ctx) error {
	fmt.Println(ctx.OriginalURL(), viper.GetStringSlice("admin.auth.exclude"))
	return ctx.Next()
}
