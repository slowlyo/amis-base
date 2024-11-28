package handlers

import (
	"amis-base/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type AdminPage struct {
}

func (a *AdminPage) GetUserInfo(ctx *fiber.Ctx) error {
	return response.Success(ctx, fiber.Map{
		"user": ctx.Locals("user"),
	})
}
