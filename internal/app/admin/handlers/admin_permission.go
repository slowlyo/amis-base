package handlers

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/app/admin/services"
	base "amis-base/internal/models"
	"amis-base/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
	"strings"
)

// AdminPermission 权限
type AdminPermission struct {
	baseHandler

	Service services.AdminPermission
}

// Index 获取列表
func (p *AdminPermission) Index(ctx *fiber.Ctx) error {
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
func (p *AdminPermission) Save(ctx *fiber.Ctx) error {
	var params struct {
		ID       int    `json:"id"`
		ParentId int    `json:"parent_id"`
		Name     string `json:"name"`
		Sign     string `json:"sign"`
		Api      string `json:"api"`
		Sort     uint   `json:"sort"`
		MenuIds  string `json:"menu_ids"`
	}

	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	permission := models.AdminPermission{
		BaseModel: base.BaseModel{ID: uint(params.ID)},
		ParentId:  uint(params.ParentId),
		Name:      params.Name,
		Sign:      params.Sign,
		Api:       params.Api,
		Sort:      params.Sort,
	}

	if err := p.Service.Save(permission, params.MenuIds); err != nil {
		return response.Error(ctx, err.Error())
	}

	return response.Ok(ctx, "保存成功")
}

// Detail 获取详情
func (p *AdminPermission) Detail(ctx *fiber.Ctx) error {
	return response.Success(ctx, p.Service.GetDetailById(ctx.QueryInt("id")))
}

// Destroy 删除
func (p *AdminPermission) Destroy(ctx *fiber.Ctx) error {
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

// SaveSort 保存排序
func (p *AdminPermission) SaveSort(ctx *fiber.Ctx) error {
	var params struct {
		Menus []models.AdminPermission `json:"rows"`
	}
	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	if err := p.Service.Reorder(params.Menus); err != nil {
		return response.Error(ctx, err.Error())
	}

	return response.Ok(ctx, "保存成功")
}

// ParentOptions 获取父级权限选项
func (p *AdminPermission) ParentOptions(ctx *fiber.Ctx) error {
	return response.Success(ctx, p.Service.GetParentOptions())
}