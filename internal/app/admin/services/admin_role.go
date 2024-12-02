package services

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/pkg/db"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type AdminRole struct {
	baseService
}

// List 获取列表
func (r *AdminRole) List(filters fiber.Map) ([]models.AdminRole, int64) {
	var count int64
	var items []models.AdminRole

	query := db.Query().Model(models.AdminRole{})

	if filters["name"].(string) != "" {
		query.Where("name like ?", "%"+filters["name"].(string)+"%")
	}
	if filters["sign"].(string) != "" {
		query.Where("sign like ?", "%"+filters["sign"].(string)+"%")
	}

	query.Count(&count)
	r.ListQuery(query, filters).Order("updated_at desc").Find(&items)

	return items, count
}

// Save 保存
func (r *AdminRole) Save(data models.AdminRole) error {
	query := db.Query().Where("name = ? or sign = ?", data.Name, data.Sign)

	if data.ID == 0 {
		if query.First(&models.AdminRole{}).RowsAffected > 0 {
			return errors.New("角色名或标识已存在")
		}

		return db.Query().Create(&data).Error
	}

	if query.Where("id != ?", data.ID).First(&models.AdminRole{}).RowsAffected > 0 {
		return errors.New("角色名或标识已存在")
	}

	return db.Query().Save(&data).Error
}

// GetDetailById 获取详情
func (r *AdminRole) GetDetailById(id int) models.AdminRole {
	var role models.AdminRole

	db.Query().First(&role, id)

	return role
}

func (r *AdminRole) Delete(ids []string) error {
	return db.Query().Where("id in ?", ids).Delete(&models.AdminRole{}).Error
}
