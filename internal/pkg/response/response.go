package response

import "github.com/gofiber/fiber/v2"

func Success(ctx *fiber.Ctx, data any) error {
	return ctx.JSON(fiber.Map{
		"status": 0,
		"code":   200,
		"msg":    "ok",
		"data":   data,
	})
}

func Ok(ctx *fiber.Ctx, msg string) error {
	data := fiber.Map{
		"status": 0,
		"code":   200,
		"msg":    msg,
		"data":   fiber.Map{},
	}

	return ctx.JSON(data)
}

func Fail(ctx *fiber.Ctx, data any) error {
	return ctx.JSON(fiber.Map{
		"status": 1,
		"code":   500,
		"msg":    "fail",
		"data":   data,
	})
}

func Error(ctx *fiber.Ctx, msg string) error {
	data := fiber.Map{
		"status": 1,
		"code":   500,
		"msg":    msg,
		"data":   fiber.Map{},
	}

	return ctx.JSON(data)
}

func UnAuthorized(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"status": 1,
		"code":   401,
		"msg":    "请先登录",
		"data":   fiber.Map{},
	})
}
