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

// AdminPage 页面管理
type AdminPage struct {
	baseHandler

	Service services.AdminPage
}

// Index 获取列表
func (p *AdminPage) Index(ctx *fiber.Ctx) error {
	filters := p.ParseParams(ctx)

	filters["name"] = ctx.Query("name")
	filters["sign"] = ctx.Query("sign")

	items, total := p.Service.List(filters)

	return response.Success(ctx, fiber.Map{
		"items": items,
		"total": total,
	})
}

// Save 保存
func (p *AdminPage) Save(ctx *fiber.Ctx) error {
	var params struct {
		ID   int                        `json:"id"`
		Name string                     `json:"name"`
		Sign string                     `json:"sign"`
		Page map[string]json.RawMessage `json:"page"`
	}
	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	page := models.AdminPage{
		BaseModel: base.BaseModel{ID: uint(params.ID)},
		Name:      params.Name,
		Sign:      params.Sign,
		Schema:    params.Page["schema"],
	}

	if err := p.Service.Save(page); err != nil {
		return response.Error(ctx, err.Error())
	}

	return response.Ok(ctx, "保存成功")
}

// Copy 复制
func (p *AdminPage) Copy(ctx *fiber.Ctx) error {
	if err := p.Service.Copy(ctx.QueryInt("id")); err != nil {
		return response.Error(ctx, err.Error())
	}

	return response.Ok(ctx, "复制成功")
}

// Detail 详情
func (p *AdminPage) Detail(ctx *fiber.Ctx) error {
	page := p.Service.GetDetailById(ctx.QueryInt("id"))

	return response.Success(ctx, fiber.Map{
		"id":   page.ID,
		"name": page.Name,
		"sign": page.Sign,
		"page": fiber.Map{
			"schema": page.Schema,
		},
	})
}

// Destroy 删除
func (p *AdminPage) Destroy(ctx *fiber.Ctx) error {
	var params idsReq
	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	if len(params.Ids) == 0 {
		return response.Error(ctx, "请选择要删除的数据")
	}

	if err := p.Service.Delete(strings.Split(params.Ids, ",")); err != nil {
		return response.Error(ctx, err.Error())
	}

	return response.Ok(ctx, "删除成功")
}
