package middleware

import (
	"amis-base/internal/app/admin/services"
	"amis-base/internal/pkg/auth"
	"amis-base/internal/pkg/helper"
	"amis-base/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"strings"
)

var authService services.Auth

// Auth 认证中间件
func Auth(ctx *fiber.Ctx) error {
	if !viper.GetBool("admin.auth.enabled") {
		return ctx.Next()
	}

	// 白名单
	exclude := viper.GetStringSlice("admin.auth.exclude")
	if len(exclude) > 0 {
		for _, v := range exclude {
			if helper.IsAllowRequest(v, ctx.Method(), ctx.OriginalURL()) {
				return ctx.Next()
			}
		}
	}

	// 获取 token
	token := strings.Replace(ctx.Get("authorization"), "Bearer ", "", -1)
	if token == "" {
		return response.UnAuthorized(ctx)
	}

	// 查询 token 信息
	tokenModel := auth.QueryToken("admin_users", token)
	if tokenModel == nil {
		return response.UnAuthorized(ctx)
	}

	// 查询用户信息
	user, err := authService.GetUserById(tokenModel.UserId)
	if err != nil {
		return response.UnAuthorized(ctx)
	}

	// 存储用户信息
	ctx.Locals("user", user)
	ctx.Locals("token", token)

	// 清除过期 token
	go auth.CleanExpiredToken()

	return ctx.Next()
}
