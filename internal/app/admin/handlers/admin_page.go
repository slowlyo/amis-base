package handlers

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/app/admin/services"
	base "amis-base/internal/models"
	"amis-base/internal/pkg/response"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type AdminPage struct {
	baseHandler

	Service services.AdminPage
}

// Index 获取列表
func (h *AdminPage) Index(ctx *fiber.Ctx) error {
	filters := h.ParseParams(ctx)

	filters["name"] = ctx.Query("name")
	filters["sign"] = ctx.Query("sign")

	items, total := h.Service.List(filters)

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

func (h *AdminPage) Save(ctx *fiber.Ctx) error {
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
		return response.Error(ctx, err.Error())
	}

	return response.Ok(ctx, "保存成功")
}

// Copy 复制
func (h *AdminPage) Copy(ctx *fiber.Ctx) error {
	if err := h.Service.Copy(ctx.QueryInt("id")); err != nil {
		return response.Error(ctx, err.Error())
	}

	return response.Ok(ctx, "复制成功")
}

func (h *AdminPage) Detail(ctx *fiber.Ctx) error {
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

type deletePageReq struct {
	Ids string `json:"ids"`
}

func (h *AdminPage) Destroy(ctx *fiber.Ctx) error {
	var params deletePageReq
	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	if len(params.Ids) == 0 {
		return response.Error(ctx, "请选择要删除的数据")
	}

	if err := h.Service.Delete(strings.Split(params.Ids, ",")); err != nil {
		return response.Error(ctx, err.Error())
	}

	return response.Ok(ctx, "删除成功")
}
