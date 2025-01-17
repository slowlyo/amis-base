package middleware

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/pkg/helper"
	"amis-base/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// Permission 鉴权中间件
func Permission(ctx *fiber.Ctx) error {
	if !viper.GetBool("admin.permission.enabled") {
		return ctx.Next()
	}

	// 白名单
	exclude := viper.GetStringSlice("admin.permission.exclude")
	if len(exclude) > 0 {
		for _, v := range exclude {
			if helper.IsAllowRequest(v, ctx.Method(), ctx.OriginalURL()) {
				return ctx.Next()
			}
		}
	}

	// 未获取到用户信息, 无法鉴权
	if ctx.Locals("user") == nil {
		return response.Forbidden(ctx)
	}

	user := ctx.Locals("user").(models.AdminUser)

	// 超级管理员
	if user.IsSuperAdmin() {
		return ctx.Next()
	}

	// 检查权限
	rules := user.PermissionApiRules()
	for _, rule := range rules {
		if helper.IsAllowRequest(rule, ctx.Method(), ctx.Path()) {
			return ctx.Next()
		}
	}

	// 无权限
	return response.Forbidden(ctx)
}
