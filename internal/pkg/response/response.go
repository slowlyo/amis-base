package response

import "github.com/gofiber/fiber/v2"

// Success 响应成功
func Success(ctx *fiber.Ctx, data any) error {
	return ctx.JSON(fiber.Map{
		"status":            0,
		"code":              200,
		"msg":               "ok",
		"doNotDisplayToast": 1,
		"data":              data,
	})
}

// Ok 响应成功消息
func Ok(ctx *fiber.Ctx, msg string) error {
	data := fiber.Map{
		"status":            0,
		"code":              200,
		"msg":               msg,
		"doNotDisplayToast": 0,
		"data":              fiber.Map{},
	}

	return ctx.JSON(data)
}

// Fail 响应失败
func Fail(ctx *fiber.Ctx, data any) error {
	return ctx.JSON(fiber.Map{
		"status":            1,
		"code":              500,
		"msg":               "fail",
		"doNotDisplayToast": 0,
		"data":              data,
	})
}

// Error 响应失败消息
func Error(ctx *fiber.Ctx, msg string) error {
	data := fiber.Map{
		"status":            1,
		"code":              500,
		"msg":               msg,
		"doNotDisplayToast": 0,
		"data":              fiber.Map{},
	}

	return ctx.JSON(data)
}

// UnAuthorized 未登录
func UnAuthorized(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"status":            1,
		"code":              401,
		"msg":               "请先登录",
		"doNotDisplayToast": 0,
		"data":              fiber.Map{},
	})
}

// Forbidden 无权访问
func Forbidden(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"status":            1,
		"code":              403,
		"msg":               "无权访问",
		"doNotDisplayToast": 0,
		"data":              fiber.Map{},
	})
}
