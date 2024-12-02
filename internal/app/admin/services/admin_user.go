package services

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/pkg/db"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strings"
)

type AdminUser struct {
	baseService
}

// List 获取列表
func (r *AdminUser) List(filters fiber.Map) ([]models.AdminUser, int64) {
	var count int64
	var items []models.AdminUser

	query := db.Query().Model(models.AdminUser{})

	if filters["name"].(string) != "" {
		query.Where("name like ?", "%"+filters["name"].(string)+"%")
	}
	if filters["username"].(string) != "" {
		query.Where("username like ?", "%"+filters["username"].(string)+"%")
	}

	query.Count(&count)
	r.ListQuery(query, filters).
		Omit("password").
		Order("updated_at desc").
		Find(&items)

	return items, count
}

// Save 保存
func (r *AdminUser) Save(data models.AdminUser, roleIdChar string) error {
	query := db.Query().Where("username = ?", data.Username)

	// 角色信息
	insertRoles := func(userId uint) error {
		var err error
		// 如果是修改, 清除原有角色
		if data.ID == 0 {
			err = db.Query().Table("admin_user_role").Where("admin_user_id = ?", userId).Delete(nil).Error
		}

		if err != nil {
			return err
		}

		if roleIdChar == "" {
			return nil
		}

		roleIds := strings.Split(roleIdChar, ",")
		insertRoles := make([]map[string]interface{}, 0)
		for _, roleId := range roleIds {
			insertRoles = append(insertRoles, map[string]interface{}{
				"admin_user_id": userId,
				"admin_role_id": roleId,
			})
		}
		return db.Query().Table("admin_user_role").Create(insertRoles).Error
	}

	if data.ID == 0 {
		if query.First(&models.AdminUser{}).RowsAffected > 0 {
			return errors.New("用户名已存在")
		}

		return db.Query().Transaction(func(tx *gorm.DB) error {
			if err := db.Query().Create(&data).Error; err != nil {
				return err
			}

			return insertRoles(data.ID)
		})
	}

	if query.Where("id != ?", data.ID).First(&models.AdminUser{}).RowsAffected > 0 {
		return errors.New("用户名已存在")
	}

	return db.Query().Transaction(func(tx *gorm.DB) error {
		if err := insertRoles(data.ID); err != nil {
			return err
		}

		return db.Query().Save(&data).Error
	})
}

// GetDetailById 获取详情
func (r *AdminUser) GetDetailById(id int) models.AdminUser {
	var user models.AdminUser

	db.Query().Omit("password").First(&user, id)

	return user
}

func (r *AdminUser) Delete(ids []string) error {
	return db.Query().Transaction(func(tx *gorm.DB) error {
		// 删除用户角色关联信息
		result := db.Query().Table("admin_user_role").Where("admin_user_id in ?", ids).Delete(nil)

		if result.Error != nil {
			return result.Error
		}

		return db.Query().Where("id in ?", ids).Delete(&models.AdminUser{}).Error
	})
}
