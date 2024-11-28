package services

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/pkg/db"
)

type AdminPage struct {
	baseService

	Model models.AdminPage
}

// GetSchemaBySign 根据页面标识获取页面结构
func (p *AdminPage) GetSchemaBySign(sign string) string {
	var page models.AdminPage

	result := db.GetDB().Where("sign = ?", sign).First(&page)

	if result.RowsAffected == 0 {
		return ""
	}

	return page.Schema
}
