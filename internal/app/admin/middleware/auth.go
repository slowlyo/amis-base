package middleware

import (
	"amis-base/internal/app/admin/services"
	"amis-base/internal/pkg/auth"
	"amis-base/internal/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"regexp"
	"strings"
)

var authService services.Auth

// Auth 认证中间件
func Auth(ctx *fiber.Ctx) error {
	// 白名单
	exclude := viper.GetStringSlice("admin.auth.exclude")
	if len(exclude) > 0 {
		for _, v := range exclude {
			splitStr := strings.Split(v, ":")
			if ctx.Route().Method == splitStr[0] && matchingString(ctx.OriginalURL(), splitStr[1]) {
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

	return ctx.Next()
}

// 匹配字符串
func matchingString(str, pattern string) bool {
	if str == pattern {
		return true
	}

	if strings.TrimLeft(str, "/") == strings.TrimLeft(pattern, "/") {
		return true
	}

	re, err := regexp.Compile(pattern)
	if err == nil {
		return re.MatchString(str)
	}

	return false
}
