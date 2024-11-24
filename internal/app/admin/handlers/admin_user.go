package handlers

import (
	"amis-base/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type AdminUser struct {
}

func (a *AdminUser) GetUserInfo(ctx *fiber.Ctx) error {
	return response.Success(ctx, fiber.Map{
		"user": ctx.Locals("user"),
	})
}
