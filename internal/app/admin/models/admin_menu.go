package models

import (
	base "amis-base/internal/models"
)

type AdminMenu struct {
	base.BaseModel

	ParentId  uint   `gorm:"type:int(11);not null;default:0" json:"parent_id"`
	Name      string `gorm:"type:varchar(255);not null" json:"name"`
	Icon      string `gorm:"type:varchar(255);not null" json:"icon"`
	Path      string `gorm:"type:varchar(255);not null" json:"path"`
	Visible   int    `gorm:"type:tinyint(1);default:0;not null" json:"visible"`
	Sort      int    `gorm:"type:int(11);default:0;not null" json:"sort"`
	IsHome    int    `gorm:"type:tinyint(1);default:0;not null" json:"is_home"`
	IsFull    int    `gorm:"type:tinyint(1);default:0;not null" json:"is_full"`
	PageSign  string `gorm:"type:varchar(255);index" json:"page_sign"`
	KeepAlive int    `gorm:"type:tinyint(1);default:0;not null" json:"keep_alive"`

	Permissions []AdminPermission `gorm:"many2many:admin_menu_permission;" json:"permissions"`
	Children    []AdminMenu       `gorm:"-" json:"children"`
	Page        AdminPage         `gorm:"foreignKey:PageSign;references:Sign" json:"page"`
}

// DevMenus 开发菜单
func (AdminMenu) DevMenus() []AdminMenu {
	return []AdminMenu{
		{
			BaseModel: base.BaseModel{ID: 100001},
			Name:      "开发",
			Icon:      "fluent:window-dev-tools-20-regular",
			Path:      "/dev",
			Visible:   1,
		},
		{
			BaseModel: base.BaseModel{ID: 100002},
			ParentId:  100001,
			Name:      "页面管理",
			Icon:      "iconoir:multiple-pages",
			Path:      "/dev/admin_page",
			Visible:   1,
			PageSign:  "admin_page",
		},
	}
}

// SystemMenus 系统菜单
func (AdminMenu) SystemMenus() []AdminMenu {
	return []AdminMenu{
		{
			BaseModel: base.BaseModel{ID: 100000},
			Name:      "个人中心",
			Path:      "/user_center",
			Visible:   0,
			PageSign:  "user_center",
		},
	}
}
