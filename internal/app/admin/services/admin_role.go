package services

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/app/admin/types"
	"amis-base/internal/pkg/db"
	"errors"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AdminRole struct {
	baseService
}

// List 获取列表
func (r *AdminRole) List(filters fiber.Map) ([]models.AdminRole, int64) {
	var count int64
	var items []models.AdminRole

	query := db.Query().
		Model(models.AdminRole{}).
		Where("sign <> ?", types.SuperAdminSign) // 超管角色不显示

	if filters["name"].(string) != "" {
		query.Where("name like ?", "%"+filters["name"].(string)+"%")
	}
	if filters["sign"].(string) != "" {
		query.Where("sign like ?", "%"+filters["sign"].(string)+"%")
	}

	query.Count(&count)
	r.ListPaginate(query, filters).Order("updated_at desc").Find(&items)

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

// Delete 删除
func (r *AdminRole) Delete(ids []string) error {
	return db.Query().Transaction(func(tx *gorm.DB) error {
		var err error

		// 删除用户角色关联信息
		err = tx.Table("admin_user_role").Where("admin_role_id in ?", ids).Delete(nil).Error
		if err != nil {
			return err
		}

		// 删除角色权限关联信息
		err = tx.Table("admin_role_permission").Where("admin_role_id in ?", ids).Delete(nil).Error
		if err != nil {
			return err
		}

		return tx.Where("id in ?", ids).Delete(&models.AdminRole{}).Error
	})
}

// GetPermissionsById 根据角色id获取权限
func (r *AdminRole) GetPermissionsById(id int) []int {
	var list []map[string]interface{}

	db.Query().
		Table("admin_role_permission").
		Where("admin_role_id = ?", id).
		Select("admin_permission_id as id").
		Find(&list)

	return slice.Map(list, func(_ int, item map[string]interface{}) int {
		return int(item["id"].(uint64))
	})
}

// SavePermissions 保存权限
func (r *AdminRole) SavePermissions(id int, permissions []int) error {
	return db.Query().Transaction(func(tx *gorm.DB) error {
		err := tx.Table("admin_role_permission").Where("admin_role_id = ?", id).Delete(nil).Error
		if err != nil {
			return err
		}

		insert := make([]map[string]interface{}, 0)
		for _, v := range permissions {
			insert = append(insert, map[string]interface{}{
				"admin_role_id":       id,
				"admin_permission_id": v,
			})
		}

		if len(insert) == 0 {
			return nil
		}

		return tx.Table("admin_role_permission").Create(insert).Error
	})
}
