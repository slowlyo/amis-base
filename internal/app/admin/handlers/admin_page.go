package handlers

import (
	"amis-base/internal/app/admin/services"
	"amis-base/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type AdminPage struct {
	BaseHandler
}

var (
	service services.AdminPage
)

// Index 获取列表
func (h *AdminPage) Index(ctx *fiber.Ctx) error {
	items, total := service.List(ctx.QueryInt("page", 1), ctx.QueryInt("perPage", 10))

	return response.Success(ctx, fiber.Map{
		"items": items,
		"total": total,
	})
}

func (h *AdminPage) Show(ctx *fiber.Ctx) error {
	return nil
}

func (h *AdminPage) Update(ctx *fiber.Ctx) error {
	return nil
}

func (h *AdminPage) Edit(ctx *fiber.Ctx) error {
	return nil
}

func (h *AdminPage) Store(ctx *fiber.Ctx) error {
	return nil
}

func (h *AdminPage) Destroy(ctx *fiber.Ctx) error {
	return nil
}

func (h *AdminPage) DestroyBatch(ctx *fiber.Ctx) error {
	return nil
}
