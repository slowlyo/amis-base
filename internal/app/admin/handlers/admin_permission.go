package handlers

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/app/admin/services"
	base "amis-base/internal/models"
	"amis-base/internal/pkg/response"
	"encoding/json"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/slice"
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
		ID       int         `json:"id"`
		ParentId int         `json:"parent_id"`
		Name     string      `json:"name"`
		Sign     string      `json:"sign"`
		Api      []fiber.Map `json:"api"`
		Sort     uint        `json:"sort"`
		MenuIds  string      `json:"menu_ids"`
	}

	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	apis := slice.Map[fiber.Map](params.Api, func(index int, item fiber.Map) string {
		return item["value"].(string)
	})

	apiJson, _ := json.Marshal(apis)

	permission := models.AdminPermission{
		BaseModel: base.BaseModel{ID: uint(params.ID)},
		ParentId:  uint(params.ParentId),
		Name:      params.Name,
		Sign:      params.Sign,
		Api:       string(apiJson),
		Sort:      params.Sort,
	}

	if err := p.Service.Save(permission, params.MenuIds); err != nil {
		return response.Error(ctx, err.Error())
	}

	return response.Ok(ctx, "保存成功")
}

// Detail 获取详情
func (p *AdminPermission) Detail(ctx *fiber.Ctx) error {
	permission := p.Service.GetDetailById(ctx.QueryInt("id"))

	menuIds := slice.Map(permission.Menus, func(_ int, item models.AdminMenu) string {
		return convertor.ToString(item.ID)
	})

	data := fiber.Map{
		"id":        permission.ID,
		"name":      permission.Name,
		"sign":      permission.Sign,
		"parent_id": permission.ParentId,
		"sort":      permission.Sort,
		"menu_ids":  strings.Join(menuIds, ","),
		"api": func() []fiber.Map {
			if permission.Api == "" {
				return []fiber.Map{}
			}

			apis := make([]string, 0)
			err := json.Unmarshal([]byte(permission.Api), &apis)
			if err != nil {
				return []fiber.Map{}
			}

			return slice.Map(apis, func(_ int, item string) fiber.Map {
				return fiber.Map{"value": item}
			})
		}(),
	}

	return response.Success(ctx, data)
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
