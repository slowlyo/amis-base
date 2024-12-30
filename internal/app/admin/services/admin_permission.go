package services

import (
	"amis-base/internal/app/admin/models"
	base "amis-base/internal/models"
	"amis-base/internal/pkg/db"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strings"
)

type AdminPermission struct {
	baseService
}

// GetTree 获取树形权限
func (p *AdminPermission) GetTree(menus []models.AdminPermission, parentId int) []models.AdminPermission {
	result := make([]models.AdminPermission, 0)
	for _, item := range menus {
		if item.ParentId == uint(parentId) {
			children := p.GetTree(menus, int(item.ID))
			if children != nil {
				item.Children = children
			}
			result = append(result, item)
		}
	}

	return result
}

// List 获取列表
func (p *AdminPermission) List(filters fiber.Map) ([]models.AdminPermission, int64) {
	var count int64
	var items []models.AdminPermission

	query := db.Query().Model(models.AdminPermission{})

	if filters["name"].(string) != "" {
		query.Where("name like ?", "%"+filters["name"].(string)+"%")
	}
	if filters["sign"].(string) != "" {
		query.Where("sign like ?", "%"+filters["sign"].(string)+"%")
	}

	query.Count(&count)
	p.ListGet(query, filters).Order("sort asc").Find(&items)

	return p.GetTree(items, 0), count
}

// Save 保存
func (p *AdminPermission) Save(data models.AdminPermission, menuIdChar string) error {
	// 菜单信息
	insertMenus := func(permissionId uint) error {
		var err error
		// 如果是修改, 清除原有菜单
		if data.ID == 0 {
			err = db.Query().Table("admin_menu_permission").Where("admin_permission_id = ?", permissionId).Delete(nil).Error
		}

		if err != nil {
			return err
		}

		if menuIdChar == "" {
			return nil
		}

		menuIds := strings.Split(menuIdChar, ",")
		insertMenus := make([]map[string]interface{}, 0)
		for _, menuId := range menuIds {
			insertMenus = append(insertMenus, map[string]interface{}{
				"admin_permission_id": permissionId,
				"admin_menu_id":       menuId,
			})
		}
		return db.Query().Table("admin_menu_permission").Create(insertMenus).Error
	}

	query := db.Query().Where("sign = ?", data.Sign)

	if data.ID == 0 {
		if query.First(&models.AdminPermission{}).RowsAffected > 0 {
			return errors.New("权限标识已存在")
		}

		return db.Query().Transaction(func(tx *gorm.DB) error {
			if err := db.Query().Create(&data).Error; err != nil {
				return err
			}

			return insertMenus(data.ID)
		})
	}

	if query.Where("id != ?", data.ID).First(&models.AdminPermission{}).RowsAffected > 0 {
		return errors.New("权限标识已存在")
	}

	return db.Query().Transaction(func(tx *gorm.DB) error {
		// 清除原有菜单关联信息
		del := db.Query().Table("admin_menu_permission").Where("admin_permission_id = ?", data.ID).Delete(nil)
		if del.Error != nil {
			return del.Error
		}

		// 保存菜单关联信息
		if err := insertMenus(data.ID); err != nil {
			return err
		}

		return db.Query().Save(&data).Error
	})
}

// GetDetailById 获取详情
func (p *AdminPermission) GetDetailById(id int) models.AdminPermission {
	var permission models.AdminPermission

	db.Query().First(&permission, id)

	return permission
}

// Delete 删除
func (p *AdminPermission) Delete(ids []string) error {
	return db.Query().Transaction(func(tx *gorm.DB) error {
		var err error

		// 删除权限权限关联信息
		err = db.Query().Table("admin_menu_permission").Where("admin_permission_id in ?", ids).Delete(nil).Error
		if err != nil {
			return err
		}

		return db.Query().Where("id in ?", ids).Delete(&models.AdminPermission{}).Error
	})
}

// Reorder 排序
func (p *AdminPermission) Reorder(menus []models.AdminPermission) error {
	sortMap := map[uint]int{}
	p.computeSort(menus, sortMap)

	sql := "update admin_permissions set `sort` = case id "
	for id, sort := range sortMap {
		sql += fmt.Sprintf("when %d then %d ", id, sort)
	}
	sql += "end where 1 = 1"

	return db.Query().Exec(sql).Error
}

// 递归计算排序
func (p *AdminPermission) computeSort(menus []models.AdminPermission, sortMap map[uint]int) {
	if len(menus) == 0 {
		return
	}

	for index, permission := range menus {
		sortMap[permission.ID] = index * 10

		p.computeSort(permission.Children, sortMap)
	}
}

// GetParentOptions 获取父级权限选项 (树)
func (p *AdminPermission) GetParentOptions() []models.AdminPermission {
	var menus []models.AdminPermission

	db.Query().Model(models.AdminPermission{}).Order("sort asc").Find(&menus)

	return append([]models.AdminPermission{{
		BaseModel: base.BaseModel{ID: 0},
		Name:      "无",
	}}, p.GetTree(menus, 0)...)
}
