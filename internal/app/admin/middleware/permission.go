package middleware

import (
	"amis-base/internal/app/admin/models"
	"github.com/gofiber/fiber/v2"
)

// Permission 鉴权中间件
func Permission(ctx *fiber.Ctx) error {
	// 未获取到用户信息, 无法鉴权
	if ctx.Locals("user") == nil {
		return ctx.Next()
	}

	user := ctx.Locals("user").(models.AdminUser)

	// 超级管理员
	if user.IsSuperAdmin() {
		return ctx.Next()
	}

	// todo ...

	return ctx.Next()
}
