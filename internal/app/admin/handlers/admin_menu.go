package handlers

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/app/admin/services"
	base "amis-base/internal/models"
	"amis-base/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
	"strings"
)

// AdminMenu 菜单
type AdminMenu struct {
	baseHandler

	Service services.AdminMenu
}

// Index 获取列表
func (m *AdminMenu) Index(ctx *fiber.Ctx) error {
	filters := m.ParseParams(ctx)

	filters["name"] = ctx.Query("name")
	filters["path"] = ctx.Query("path")

	items, total := m.Service.List(filters)

	return response.Success(ctx, fiber.Map{
		"items": items,
		"total": total,
	})
}

// Save 保存
func (m *AdminMenu) Save(ctx *fiber.Ctx) error {
	var params struct {
		ID        int    `json:"id"`
		ParentId  int    `json:"parent_id"`
		Name      string `json:"name"`
		Icon      string `json:"icon"`
		Path      string `json:"path"`
		Visible   int    `json:"visible"`
		Sort      int    `json:"sort"`
		IsHome    int    `json:"is_home"`
		IsFull    int    `json:"is_full"`
		PageSign  string `json:"page_sign"`
		KeepAlive int    `json:"keep_alive"`
	}

	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	role := models.AdminMenu{
		BaseModel: base.BaseModel{ID: uint(params.ID)},
		ParentId:  uint(params.ParentId),
		Name:      params.Name,
		Icon:      params.Icon,
		Path:      params.Path,
		Visible:   params.Visible,
		Sort:      params.Sort,
		IsHome:    params.IsHome,
		IsFull:    params.IsFull,
		PageSign:  params.PageSign,
		KeepAlive: params.KeepAlive,
	}

	if !strings.HasPrefix(role.Path, "http") {
		role.Path = "/" + strings.TrimLeft(role.Path, "/")
	}

	if err := m.Service.Save(role); err != nil {
		return response.Error(ctx, err.Error())
	}

	return response.Ok(ctx, "保存成功")
}

// Detail 获取详情
func (m *AdminMenu) Detail(ctx *fiber.Ctx) error {
	return response.Success(ctx, m.Service.GetDetailById(ctx.QueryInt("id")))
}

// Destroy 删除
func (m *AdminMenu) Destroy(ctx *fiber.Ctx) error {
	var params idsReq
	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	if len(params.Ids) == 0 {
		return response.Error(ctx, "请选择要删除的数据")
	}

	if err := m.Service.Delete(strings.Split(params.Ids, ",")); err != nil {
		return response.Error(ctx, err.Error())
	}

	return response.Ok(ctx, "删除成功")
}

// QuickSave 快速保存
func (m *AdminMenu) QuickSave(ctx *fiber.Ctx) error {
	var params struct {
		ID      int `json:"id"`
		Visible int `json:"visible"`
		IsHome  int `json:"is_home"`
	}
	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	user := models.AdminMenu{
		BaseModel: base.BaseModel{ID: uint(params.ID)},
		Visible:   params.Visible,
		IsHome:    params.IsHome,
	}

	if err := m.Service.QuickSave(user); err != nil {
		return response.Error(ctx, err.Error())
	}

	return response.Ok(ctx, "保存成功")
}

// SaveSort 保存排序
func (m *AdminMenu) SaveSort(ctx *fiber.Ctx) error {
	var params struct {
		Menus []models.AdminMenu `json:"rows"`
	}
	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	if err := m.Service.Reorder(params.Menus); err != nil {
		return response.Error(ctx, err.Error())
	}

	return response.Ok(ctx, "保存成功")
}

// ParentOptions 获取父级菜单选项
func (m *AdminMenu) ParentOptions(ctx *fiber.Ctx) error {
	return response.Success(ctx, m.Service.GetParentOptions())
}

// PageOptions 获取页面选项
func (m *AdminMenu) PageOptions(ctx *fiber.Ctx) error {
	return response.Success(ctx, m.Service.GetPageOptions())
}
