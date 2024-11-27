package handlers

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/app/admin/services"
	"amis-base/internal/pkg/auth"
	"amis-base/internal/pkg/response"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type System struct {
}

var (
	authService    services.Auth
	menuService    services.AdminMenu
	settingService services.AdminSetting
)

// Settings 获取系统设置
func (s *System) Settings(ctx *fiber.Ctx) error {
	return response.Success(ctx, fiber.Map{
		"dev":     viper.GetBool("app.dev"),
		"appName": viper.GetString("app.name"),
		"logo":    viper.GetString("app.logo"),
		"theme":   settingService.Get("system.theme"),
	})
}

type saveSettingReq struct {
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
}

// SaveSettings 保存系统设置
func (s *System) SaveSettings(ctx *fiber.Ctx) error {
	var params saveSettingReq

	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	if err := settingService.Set(params.Key, params.Value); err != nil {
		return response.Error(ctx, "保存失败")
	}

	return response.Ok(ctx, "保存成功")
}

// Menus 菜单
func (s *System) Menus(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(models.AdminUser)
	menus := menuService.GetUserMenus(user)

	return response.Success(ctx, *menuService.BuildRoutes(menus, 0))
}

// User 获取用户信息
func (s *System) User(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(models.AdminUser)

	return response.Success(ctx, fiber.Map{
		"name":   user.Name,
		"avatar": user.Avatar,
	})
}

type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login 登录
func (s *System) Login(ctx *fiber.Ctx) error {
	var params loginReq

	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	user, err := authService.GetUserByUsername(params.Username)

	if err != nil || !auth.CheckHash(params.Password, user.Password) {
		return response.Error(ctx, "用户名或密码错误")
	}

	if user.Enabled == 0 {
		return response.Error(ctx, "用户已被禁用")
	}

	return response.Success(ctx, fiber.Map{
		"token": auth.GenerateToken("admin_users", user.ID),
	})
}

// Logout 退出登录
func (s *System) Logout(ctx *fiber.Ctx) error {
	auth.RemoveToken(ctx.Locals("token").(string))

	return response.Success(ctx, nil)
}

// PageSchema 获取页面结构
func (s *System) PageSchema(ctx *fiber.Ctx) error {
	return response.Success(ctx, fiber.Map{
		"type": "tpl",
		"tpl":  "hello " + ctx.Query("sign"),
	})
}
