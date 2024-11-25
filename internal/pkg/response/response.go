package response

import "github.com/gofiber/fiber/v2"

func Success(ctx *fiber.Ctx, data fiber.Map) error {
	return ctx.JSON(fiber.Map{
		"status": 0,
		"msg":    "",
		"data":   data,
	})
}

func Ok(ctx *fiber.Ctx, msg string) error {
	data := fiber.Map{
		"status": 0,
		"msg":    msg,
		"data":   fiber.Map{},
	}

	return ctx.JSON(data)
}

func Fail(ctx *fiber.Ctx, data fiber.Map) error {
	data["status"] = 1

	return ctx.JSON(data)
}

func Error(ctx *fiber.Ctx, msg string) error {
	data := fiber.Map{
		"status": 1,
		"msg":    msg,
		"data":   fiber.Map{},
	}

	return ctx.JSON(data)
}
