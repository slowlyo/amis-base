package handlers

import "github.com/gofiber/fiber/v2"

type Common struct {
}

// Settings 获取系统设置
func (c *Common) Settings(ctx *fiber.Ctx) error {
	return nil
}

// SaveSettings 保存系统设置
func (c *Common) SaveSettings(ctx *fiber.Ctx) error {
	return nil
}
