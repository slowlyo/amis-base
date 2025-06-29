package services

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/app/admin/types"
	base "amis-base/internal/models"
	"amis-base/internal/pkg/db"
	"amis-base/internal/schema"
	"errors"
	"fmt"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AdminMenu struct {
	baseService
}

func (m *AdminMenu) GetUserMenus(user *models.AdminUser) *[]models.AdminMenu {
	var menus []models.AdminMenu

	query := db.Query().Model(models.AdminMenu{})

	if !user.IsSuperAdmin() {
		subQuery := db.Query().
			Table("admin_menu_permission").
			Where("admin_permission_id in (?)", slice.Map(user.Permissions(), func(_ int, item models.AdminPermission) uint {
				return item.ID
			})).
			Select("admin_menu_id")

		query.Where("id in (?)", subQuery)
	}

	query.Order("sort asc").Find(&menus)

	// 追加开发者菜单
	if viper.GetBool("app.dev") && user.IsSuperAdmin() {
		menus = append(menus, models.AdminMenu{}.DevMenus()...)
	}

	// 追加系统菜单
	menus = append(menus, models.AdminMenu{}.SystemMenus()...)

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
	result := make([]models.AdminMenu, 0)
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

	query := db.Query().Model(models.AdminMenu{}).Preload("Page", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "sign")
	})

	if filters["name"].(string) != "" {
		query.Where("name like ?", "%"+filters["name"].(string)+"%")
	}
	if filters["path"].(string) != "" {
		query.Where("path like ?", "%"+filters["path"].(string)+"%")
	}

	query.Count(&count)
	m.ListGet(query, filters).Order("sort asc").Find(&items)

	return m.GetTree(items, 0), count
}

// Save 保存
func (m *AdminMenu) Save(data models.AdminMenu) error {
	if data.ParentId != 0 {
		if data.ParentId == data.ID {
			return errors.New("父级菜单不能选择自己")
		}

		parentId := data.ParentId
		for {
			parentMenu := models.AdminPermission{}
			db.Query().Where("id = ?", parentId).First(&parentMenu)

			if parentMenu.ID == data.ID {
				return errors.New("不可选择子菜单限作为父级")
			}

			if parentMenu.ParentId == 0 {
				break
			}

			parentId = parentMenu.ParentId
		}
	}

	query := db.Query().Where("path = ?", data.Path)

	if data.ID == 0 {
		if query.First(&models.AdminMenu{}).RowsAffected > 0 {
			return errors.New("菜单路径已存在")
		}

		m.updateIsHome(data)

		return db.Query().Create(&data).Error
	}

	if query.Where("id != ?", data.ID).First(&models.AdminMenu{}).RowsAffected > 0 {
		return errors.New("菜单路径已存在")
	}

	m.updateIsHome(data)

	return db.Query().Save(&data).Error
}

func (m *AdminMenu) updateIsHome(data models.AdminMenu) {
	if data.IsHome == 1 {
		db.Query().Model(&models.AdminMenu{}).Where("is_home = ?", 1).Where("id != ?", data.ID).Update("is_home", 0)
	}
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
		err = tx.Table("admin_menu_permission").Where("admin_menu_id in ?", ids).Delete(nil).Error
		if err != nil {
			return err
		}

		return tx.Where("id in ?", ids).Delete(&models.AdminMenu{}).Error
	})
}

// QuickSave 快速保存
func (m *AdminMenu) QuickSave(menu models.AdminMenu) error {
	m.updateIsHome(menu)

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

// GetParentOptions 获取父级菜单选项 (树)
func (m *AdminMenu) GetParentOptions() []models.AdminMenu {
	var menus []models.AdminMenu

	db.Query().Model(models.AdminMenu{}).Order("sort asc").Find(&menus)

	return append([]models.AdminMenu{{
		BaseModel: base.BaseModel{ID: 0},
		Name:      "无",
	}}, m.GetTree(menus, 0)...)
}

// GetPageOptions 获取页面选项
func (m *AdminMenu) GetPageOptions() []fiber.Map {
	var result []fiber.Map

	// 首先添加程序内定义的schema选项
	embeddedSchemas := schema.GetAllSchemas()
	for _, schemaInfo := range embeddedSchemas {
		result = append(result, fiber.Map{
			"label": schemaInfo.Name,
			"value": schemaInfo.Sign,
		})
	}

	// 然后添加数据库中的页面选项
	var pages []models.AdminPage
	db.Query().Model(models.AdminPage{}).Find(&pages)

	for _, page := range pages {
		// 检查是否与程序内定义的schema重复
		isDuplicate := false
		for _, schemaInfo := range embeddedSchemas {
			if schemaInfo.Sign == page.Sign {
				isDuplicate = true
				break
			}
		}

		// 如果不重复，则添加到结果中
		if !isDuplicate {
			result = append(result, fiber.Map{
				"label": page.Name,
				"value": page.Sign,
			})
		}
	}

	return result
}
