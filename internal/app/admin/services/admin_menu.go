package services

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/app/admin/types"
	"amis-base/internal/pkg/db"
	"errors"
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AdminMenu struct {
	baseService
}

func (m *AdminMenu) GetUserMenus(user models.AdminUser) *[]models.AdminMenu {
	var menus []models.AdminMenu

	query := db.Query().Model(models.AdminMenu{})

	if !user.IsSuperAdmin() {
		// todo 权限
	}

	query.Order("sort asc").Find(&menus)

	// 追加开发者菜单
	if viper.GetBool("app.dev") {
		menus = append(menus, models.AdminMenu{}.DevMenus()...)
	}

	return &menus
}

// BuildRoutes 构建路由
func (m *AdminMenu) BuildRoutes(menus *[]models.AdminMenu, parentId uint) *[]types.AdminRoutes {
	var routes []types.AdminRoutes
	for _, menu := range *menus {
		if menu.ParentId == parentId {
			route := types.AdminRoutes{
				Name:      cryptor.Sha256(convertor.ToString(menu.ID)),
				Path:      menu.Path,
				PageSign:  menu.PageSign,
				Component: "amis",
				IsFull:    menu.IsFull,
				IsHome:    menu.IsHome,
				Meta: types.AdminRouteMeta{
					Title: menu.Name,
					Icon:  menu.Icon,
					Sort:  menu.Sort,
					Hide:  menu.Visible == 0,
				},
			}

			children := m.BuildRoutes(menus, menu.ID)

			if children != nil {
				route.Children = children
			}

			routes = append(routes, route)
		}
	}

	return &routes
}

// GetTree 获取树形菜单
func (m *AdminMenu) GetTree(menus []models.AdminMenu, parentId int) []models.AdminMenu {
	var result []models.AdminMenu
	for _, item := range menus {
		if item.ParentId == uint(parentId) {
			children := m.GetTree(menus, int(item.ID))
			if children != nil {
				item.Children = children
			}
			result = append(result, item)
		}
	}

	return result
}

// List 获取列表
func (m *AdminMenu) List(filters fiber.Map) ([]models.AdminMenu, int64) {
	var count int64
	var items []models.AdminMenu

	query := db.Query().Model(models.AdminMenu{})

	if filters["name"].(string) != "" {
		query.Where("name like ?", "%"+filters["name"].(string)+"%")
	}
	if filters["path"].(string) != "" {
		query.Where("path like ?", "%"+filters["path"].(string)+"%")
	}

	query.Count(&count)
	m.ListQuery(query, filters).Order("sort asc").Find(&items)

	return m.GetTree(items, 0), count
}

// Save 保存
func (m *AdminMenu) Save(data models.AdminMenu) error {
	query := db.Query().Where("path = ?", data.Path)

	if data.ID == 0 {
		if query.First(&models.AdminMenu{}).RowsAffected > 0 {
			return errors.New("菜单路径已存在")
		}

		return db.Query().Create(&data).Error
	}

	if query.Where("id != ?", data.ID).First(&models.AdminMenu{}).RowsAffected > 0 {
		return errors.New("菜单路径已存在")
	}

	return db.Query().Save(&data).Error
}

// GetDetailById 获取详情
func (m *AdminMenu) GetDetailById(id int) models.AdminMenu {
	var menu models.AdminMenu

	db.Query().First(&menu, id)

	return menu
}

// Delete 删除
func (m *AdminMenu) Delete(ids []string) error {
	return db.Query().Transaction(func(tx *gorm.DB) error {
		var err error

		// 删除菜单权限关联信息
		err = db.Query().Table("admin_menu_permission").Where("admin_menu_id in ?", ids).Delete(nil).Error
		if err != nil {
			return err
		}

		return db.Query().Where("id in ?", ids).Delete(&models.AdminMenu{}).Error
	})
}

// QuickSave 快速保存
func (m *AdminMenu) QuickSave(menu models.AdminMenu) error {
	// 首页只能有一个
	if menu.IsHome == 1 {
		db.Query().Model(&models.AdminMenu{}).Where("is_home = ?", 1).Update("is_home", 0)
	}

	return db.Query().Select("visible", "is_home").Save(menu).Error
}

// Reorder 排序
func (m *AdminMenu) Reorder(menus []models.AdminMenu) error {
	sortMap := map[uint]int{}
	m.computeSort(menus, sortMap)

	sql := "update admin_menus set `sort` = case id "
	for id, sort := range sortMap {
		sql += fmt.Sprintf("when %d then %d ", id, sort)
	}
	sql += "end where 1 = 1"

	return db.Query().Exec(sql).Error
}

// 递归计算排序
func (m *AdminMenu) computeSort(menus []models.AdminMenu, sortMap map[uint]int) {
	if len(menus) == 0 {
		return
	}

	for index, menu := range menus {
		sortMap[menu.ID] = index * 10

		m.computeSort(menu.Children, sortMap)
	}
}
