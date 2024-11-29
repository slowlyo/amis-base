package services

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/pkg/db"
	"encoding/json"
)

type AdminPage struct {
	baseService
}

// List 获取列表
func (s baseService) List(page, perPage int) ([]models.AdminPage, int64) {
	var count int64
	var items []models.AdminPage

	query := db.Query().Model(models.AdminPage{}).Omit("Schema")

	query.Count(&count)
	query.Offset((page - 1) * perPage).Limit(perPage).Find(&items)

	return items, count
}

// Save 保存
func (p *AdminPage) Save(data models.AdminPage) error {
	if data.ID == 0 {
		return db.Query().Create(&data).Error
	}

	return db.Query().Save(&data).Error
}

// GetDetailById 获取详情
func (p *AdminPage) GetDetailById(id int) models.AdminPage {
	var page models.AdminPage

	db.Query().First(&page, id)

	return page
}

// GetSchemaBySign 根据页面标识获取页面结构
func (p *AdminPage) GetSchemaBySign(sign string) json.RawMessage {
	var page models.AdminPage

	result := db.Query().Where("sign = ?", sign).First(&page)

	if result.RowsAffected == 0 {
		return nil
	}

	return page.Schema
}
