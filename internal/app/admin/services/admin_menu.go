package services

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/app/admin/types"
	"amis-base/internal/pkg/db"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/spf13/viper"
)

type AdminMenu struct {
}

func (m *AdminMenu) GetUserMenus(user models.AdminUser) *[]models.AdminMenu {
	var menus []models.AdminMenu

	query := db.Query().Model(models.AdminMenu{})

	if !user.IsSuperAdmin() {
		// todo 权限
	}

	query.Find(&menus)

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
