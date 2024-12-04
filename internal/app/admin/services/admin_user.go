package services

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/app/admin/types"
	"amis-base/internal/pkg/db"
	"errors"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strings"
)

type AdminUser struct {
	baseService
}

// List 获取列表
func (r *AdminUser) List(filters fiber.Map) ([]fiber.Map, int64) {
	var count int64
	var list []models.AdminUser

	query := db.Query().Model(models.AdminUser{})

	if filters["name"].(string) != "" {
		query.Where("name like ?", "%"+filters["name"].(string)+"%")
	}
	if filters["username"].(string) != "" {
		query.Where("username like ?", "%"+filters["username"].(string)+"%")
	}

	query.Count(&count)
	r.ListQuery(query, filters).
		Preload("Roles").
		Omit("password").
		Order("updated_at desc").
		Find(&list)

	items := make([]fiber.Map, len(list))
	for i, item := range list {
		items[i] = fiber.Map{
			"id":         item.ID,
			"avatar":     item.Avatar,
			"name":       item.Name,
			"username":   item.Username,
			"enabled":    item.Enabled,
			"created_at": item.CreatedAt,
			"updated_at": item.UpdatedAt,

			"roles": slice.Map(item.Roles, func(_ int, role models.AdminRole) fiber.Map {
				return fiber.Map{"id": role.ID, "name": role.Name}
			}),
		}
	}

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
		// 清除原有角色关联信息
		del := db.Query().Table("admin_user_role").Where("admin_user_id = ?", data.ID).Delete(nil)
		if del.Error != nil {
			return del.Error
		}

		// 保存角色关联信息
		if err := insertRoles(data.ID); err != nil {
			return err
		}

		saveQuery := db.Query()

		// 如果密码为空, 则不保存
		if data.Password == "" {
			saveQuery = saveQuery.Omit("password")
		}

		return saveQuery.Save(&data).Error
	})
}

// GetDetailById 获取详情
func (r *AdminUser) GetDetailById(id int) fiber.Map {
	var user models.AdminUser

	db.Query().Preload("Roles").Omit("password").First(&user, id)

	result := fiber.Map{
		"id":       user.ID,
		"avatar":   user.Avatar,
		"name":     user.Name,
		"username": user.Username,
		"enabled":  user.Enabled,
		"roleIds": strings.Join(slice.Map(user.Roles, func(_ int, role models.AdminRole) string {
			return convertor.ToString(role.ID)
		}), ","),
	}

	return result
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

// GetRoleOptions 获取用户角色选项
func (r *AdminUser) GetRoleOptions(isAdministrator bool) []types.Options {
	query := db.Query()

	// 非超管, 不可设置超管
	if !isAdministrator {
		query.Where("sign <> ?", "administrator")
	}

	var list []models.AdminRole
	query.Find(&list)

	var options []types.Options
	for _, item := range list {
		options = append(options, types.Options{
			Label: item.Name,
			Value: convertor.ToString(item.ID),
		})
	}

	return options
}

// QuickSave 快速保存
func (r *AdminUser) QuickSave(user models.AdminUser) error {
	return db.Query().Select("enabled").Save(user).Error
}
