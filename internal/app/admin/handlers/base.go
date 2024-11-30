package handlers

import "github.com/gofiber/fiber/v2"

// BaseHandler 控制器的默认实现
type baseHandler struct {
}

func (h baseHandler) ParseParams(ctx *fiber.Ctx) fiber.Map {
	return fiber.Map{
		"page":     ctx.QueryInt("page", 1),
		"perPage":  ctx.QueryInt("perPage", 20),
		"orderBy":  ctx.Query("orderBy"),
		"orderDir": ctx.Query("orderDir", "asc"),
	}
}
