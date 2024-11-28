package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type BaseHandlerInterface interface {
	Index(ctx *fiber.Ctx) error
	Show(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Edit(ctx *fiber.Ctx) error
	Store(ctx *fiber.Ctx) error
	Destroy(ctx *fiber.Ctx) error
	DestroyBatch(ctx *fiber.Ctx) error
}

// BaseHandler 控制器的默认实现
type BaseHandler struct {
}
