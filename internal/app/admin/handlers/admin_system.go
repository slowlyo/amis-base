package handlers

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/app/admin/services"
	"amis-base/internal/pkg/auth"
	"amis-base/internal/pkg/helper"
	"amis-base/internal/pkg/response"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"strings"
)

// AdminSystem 系统基本内容
type AdminSystem struct {
	AuthService    services.Auth
	MenuService    services.AdminMenu
	PageService    services.AdminPage
	SettingService services.AdminSetting
}

// Settings 获取系统设置
func (s *AdminSystem) Settings(ctx *fiber.Ctx) error {
	return response.Success(ctx, fiber.Map{
		"dev":     viper.GetBool("app.dev"),
		"appName": viper.GetString("app.name"),
		"logo":    viper.GetString("app.logo"),
		"theme":   s.SettingService.Get("system.theme"),
	})
}

// SaveSettings 保存系统设置
func (s *AdminSystem) SaveSettings(ctx *fiber.Ctx) error {
	var params struct {
		Key   string          `json:"key"`
		Value json.RawMessage `json:"value"`
	}

	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	if err := s.SettingService.Set(params.Key, params.Value); err != nil {
		return response.Error(ctx, "保存失败")
	}

	return response.Ok(ctx, "保存成功")
}

// Menus 菜单
func (s *AdminSystem) Menus(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(models.AdminUser)
	menus := s.MenuService.GetUserMenus(user)

	return response.Success(ctx, *s.MenuService.BuildRoutes(menus, 0))
}

// User 获取用户信息
func (s *AdminSystem) User(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(models.AdminUser)

	return response.Success(ctx, fiber.Map{
		"name":   user.Name,
		"avatar": user.Avatar,
	})
}

// Login 登录
func (s *AdminSystem) Login(ctx *fiber.Ctx) error {
	var params struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	user, err := s.AuthService.GetUserByUsername(params.Username)

	if err != nil || user.ID == 0 || !auth.CheckHash(params.Password, user.Password) {
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
func (s *AdminSystem) Logout(ctx *fiber.Ctx) error {
	auth.RemoveToken(ctx.Locals("token").(string))

	return response.Success(ctx, nil)
}

// Upload 文件上传
func (s *AdminSystem) Upload(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return response.Error(ctx, "参数错误")
	}

	// 资源文件夹
	resourcePath := "./assets"
	// 文件上传目录
	uploadDir := "uploads"
	pathSlice := []string{resourcePath, uploadDir}

	// 如果传递了自定义的目录, 则添加到路径中
	dir := ctx.FormValue("dir")
	if dir != "" {
		pathSlice = append(pathSlice, dir)
	}

	// 重命名文件
	nameSplit := strings.Split(file.Filename, ".")
	fileName := uuid.NewString() + "." + nameSplit[len(nameSplit)-1]
	pathSlice = append(pathSlice, fileName)

	filePath := strings.Join(pathSlice, "/")

	helper.MakeDir(filePath)

	err = ctx.SaveFile(file, filePath)
	if err != nil {
		return response.Error(ctx, "文件保存失败: "+err.Error())
	}

	url := fmt.Sprintf("%s://%s%s", ctx.Protocol(), ctx.Hostname(), strings.ReplaceAll(filePath, resourcePath, ""))

	return response.Success(ctx, fiber.Map{"url": url})
}

// PageSchema 获取页面结构
func (s *AdminSystem) PageSchema(ctx *fiber.Ctx) error {
	pageSign := ctx.Query("sign")
	schemaStr := s.PageService.GetSchemaBySign(pageSign)

	if schemaStr == nil {
		return response.Success(ctx, fiber.Map{
			"type": "page",
			"body": fiber.Map{
				"type":     "alert",
				"showIcon": true,
				"level":    "danger",
				"body":     fmt.Sprintf("页面 %s 不存在", pageSign),
			},
		})
	}

	var schema any

	err := json.Unmarshal([]byte(schemaStr), &schema)
	if err != nil {
		return response.Error(ctx, "页面结构解析失败")
	}

	return response.Success(ctx, schema)
}
