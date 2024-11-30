package handlers

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/app/admin/services"
	base "amis-base/internal/models"
	"amis-base/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type AdminRole struct {
	baseHandler

	Service services.AdminRole
}

// Index 获取列表
func (h *AdminRole) Index(ctx *fiber.Ctx) error {
	filters := h.ParseParams(ctx)

	filters["name"] = ctx.Query("name")
	filters["sign"] = ctx.Query("sign")

	items, total := h.Service.List(filters)

	return response.Success(ctx, fiber.Map{
		"items": items,
		"total": total,
	})
}

type saveRoleReq struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Sign string `json:"sign"`
}

func (h *AdminRole) Save(ctx *fiber.Ctx) error {
	var params saveRoleReq
	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	page := models.AdminRole{
		BaseModel: base.BaseModel{ID: uint(params.ID)},
		Name:      params.Name,
		Sign:      params.Sign,
	}

	if err := h.Service.Save(page); err != nil {
		return response.Error(ctx, err.Error())
	}

	return response.Ok(ctx, "保存成功")
}

func (h *AdminRole) Detail(ctx *fiber.Ctx) error {
	return response.Success(ctx, h.Service.GetDetailById(ctx.QueryInt("id")))
}

type deleteRoleReq struct {
	Ids string `json:"ids"`
}

func (h *AdminRole) Destroy(ctx *fiber.Ctx) error {
	var params deleteRoleReq
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
