package handlers

import (
	"amis-base/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type Index struct {
}

func (i *Index) Hello(ctx *fiber.Ctx) error {
	return response.Ok(ctx, "hello AmisBase~")
}
