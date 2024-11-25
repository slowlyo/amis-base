package handlers

import (
	"amis-base/internal/app/admin/services"
	"amis-base/internal/pkg/response"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type Common struct {
}

var (
	settingService services.AdminSetting
)

// Settings 获取系统设置
func (c *Common) Settings(ctx *fiber.Ctx) error {
	return response.Success(ctx, fiber.Map{
		"dev":     viper.GetBool("app.dev"),
		"appName": viper.GetString("app.name"),
		"theme":   settingService.Get("system.theme"),
	})
}

type saveSettingReq struct {
	Key   string          `json:"key"`
	Value json.RawMessage `json:"value"`
}

// SaveSettings 保存系统设置
func (c *Common) SaveSettings(ctx *fiber.Ctx) error {
	var params saveSettingReq

	if err := ctx.BodyParser(&params); err != nil {
		return response.Error(ctx, "参数错误")
	}

	if err := settingService.Set(params.Key, params.Value); err != nil {
		return response.Error(ctx, "保存失败")
	}

	return response.Ok(ctx, "保存成功")
}
