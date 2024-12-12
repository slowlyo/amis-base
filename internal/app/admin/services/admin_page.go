package services

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/pkg/db"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type AdminPage struct {
	baseService
}

// List 获取列表
func (p *AdminPage) List(filters fiber.Map) ([]models.AdminPage, int64) {
	var count int64
	var items []models.AdminPage

	query := db.Query().Model(models.AdminPage{})

	if filters["name"].(string) != "" {
		query.Where("name like ?", "%"+filters["name"].(string)+"%")
	}
	if filters["sign"].(string) != "" {
		query.Where("sign like ?", "%"+filters["sign"].(string)+"%")
	}

	query.Count(&count)
	p.ListQuery(query, filters).
		Omit("schema").
		Order("updated_at desc").
		Find(&items)

	return items, count
}

// Save 保存
func (p *AdminPage) Save(data models.AdminPage) error {
	query := db.Query().Where("sign = ?", data.Sign)

	if data.ID == 0 {
		if query.First(&models.AdminPage{}).RowsAffected > 0 {
			return errors.New("页面标识已存在")
		}

		return db.Query().Create(&data).Error
	}

	if query.Where("id != ?", data.ID).First(&models.AdminPage{}).RowsAffected > 0 {
		return errors.New("页面标识已存在")
	}

	return db.Query().Save(&data).Error
}

// Copy 复制
func (p *AdminPage) Copy(id int) error {
	var original models.AdminPage
	db.Query().First(&original, id)
	if original.ID == 0 {
		return errors.New("页面不存在")
	}

	newPage := models.AdminPage{
		Name:   original.Name + " (副本)",
		Sign:   original.Sign + "_copy",
		Schema: original.Schema,
	}

	return p.Save(newPage)
}

// GetDetailById 获取详情
func (p *AdminPage) GetDetailById(id int) models.AdminPage {
	var page models.AdminPage

	db.Query().First(&page, id)

	return page
}

// Delete 删除
func (p *AdminPage) Delete(ids []string) error {
	return db.Query().Where("id in ?", ids).Delete(&models.AdminPage{}).Error
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

// QuickSave 快速保存
func (p *AdminPage) QuickSave(page models.AdminPage) error {
	//return db.Query().Model(&models.AdminPage{}).Where("id = ?", page.ID).Update("schema", string(page.Schema)).Error
	return db.Query().Select("Schema").Save(page).Error
}
