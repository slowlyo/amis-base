package handlers

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/app/admin/services"
	base "amis-base/internal/models"
	"amis-base/internal/pkg/helper"
	"amis-base/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type AdminUser struct {
	baseHandler

	Service services.AdminUser
}

// Index 获取列表
func (u *AdminUser) Index(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(models.AdminUser)
	filters := u.ParseParams(ctx)

	filters["name"] = ctx.Query("name")
	filters["username"] = ctx.Query("username")

	items, total := u.Service.List(filters, user.IsSuperAdmin())

	return response.Success(ctx, fiber.Map{"items": items, "total": total})
}

// Save 保存
func (u *AdminUser) Save(ctx *fiber.Ctx) error {
	var params struct {
		ID       int    `json:"id"`
		Avatar   string `json:"avatar"`
		Username string `json:"username"`
		Name     string `json:"name"`
		Password string `json:"password"`
		RoleIds  string `json:"roleIds"`
		Enabled  int    `json:"enabled"`
	}
	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	user := models.AdminUser{
		BaseModel: base.BaseModel{ID: uint(params.ID)},
		Avatar:    params.Avatar,
		Username:  params.Username,
		Name:      params.Name,
		Enabled:   params.Enabled,
	}

	// 传了密码才修改
	if params.Password != "" {
		user.Password = helper.BcryptString(params.Password)
	}

	if err := u.Service.Save(user, params.RoleIds); err != nil {
		return response.Error(ctx, err.Error())
	}

	return response.Ok(ctx, "保存成功")
}

// Detail 详情
func (u *AdminUser) Detail(ctx *fiber.Ctx) error {
	return response.Success(ctx, u.Service.GetDetailById(ctx.QueryInt("id")))
}

// Destroy 删除
func (u *AdminUser) Destroy(ctx *fiber.Ctx) error {
	var params idsReq
	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	if len(params.Ids) == 0 {
		return response.Error(ctx, "请选择要删除的数据")
	}

	if err := u.Service.Delete(strings.Split(params.Ids, ",")); err != nil {
		return response.Error(ctx, err.Error())
	}

	return response.Ok(ctx, "删除成功")
}

// RoleOptions 角色下拉框
func (u *AdminUser) RoleOptions(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(models.AdminUser)

	return response.Success(ctx, u.Service.GetRoleOptions(user.IsSuperAdmin()))
}

// QuickSave 快速保存
func (u *AdminUser) QuickSave(ctx *fiber.Ctx) error {
	var params struct {
		ID      int `json:"id"`
		Enabled int `json:"enabled"`
	}
	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	user := models.AdminUser{
		BaseModel: base.BaseModel{ID: uint(params.ID)},
		Enabled:   params.Enabled,
	}

	if err := u.Service.QuickSave(user); err != nil {
		return response.Error(ctx, err.Error())
	}

	return response.Ok(ctx, "保存成功")
}
