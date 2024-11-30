package handlers

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/app/admin/services"
	base "amis-base/internal/models"
	"amis-base/internal/pkg/response"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type AdminPage struct {
	baseHandler

	Service services.AdminPage
}

// Index 获取列表
func (h *AdminPage) Index(ctx *fiber.Ctx) error {
	items, total := h.Service.List(ctx.QueryInt("page", 1), ctx.QueryInt("perPage", 10))

	return response.Success(ctx, fiber.Map{
		"items": items,
		"total": total,
	})
}

type savePageReq struct {
	ID   int                        `json:"id"`
	Name string                     `json:"name"`
	Sign string                     `json:"sign"`
	Page map[string]json.RawMessage `json:"page"`
}

func (h *AdminPage) Update(ctx *fiber.Ctx) error {
	var params savePageReq
	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	page := models.AdminPage{
		BaseModel: base.BaseModel{ID: uint(params.ID)},
		Name:      params.Name,
		Sign:      params.Sign,
		Schema:    params.Page["schema"],
	}

	if err := h.Service.Save(page); err != nil {
		return response.Error(ctx, "保存失败")
	}

	return response.Ok(ctx, "保存成功")
}

func (h *AdminPage) Edit(ctx *fiber.Ctx) error {
	page := h.Service.GetDetailById(ctx.QueryInt("id"))

	return response.Success(ctx, fiber.Map{
		"id":   page.ID,
		"name": page.Name,
		"sign": page.Sign,
		"page": fiber.Map{
			"schema": page.Schema,
		},
	})
}

func (h *AdminPage) Store(ctx *fiber.Ctx) error {
	var params savePageReq
	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	page := models.AdminPage{
		Name:   params.Name,
		Sign:   params.Sign,
		Schema: params.Page["schema"],
	}

	if err := h.Service.Save(page); err != nil {
		return response.Error(ctx, "保存失败")
	}

	return response.Ok(ctx, "保存成功")
}

func (h *AdminPage) Destroy(ctx *fiber.Ctx) error {
	return nil
}
