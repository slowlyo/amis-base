package models

import (
	base "amis-base/internal/models"
)

type AdminMenu struct {
	base.BaseModel

	ParentId  uint   `gorm:"type:int(11);not null;default:0"`
	Name      string `gorm:"type:varchar(255);not null"`
	Icon      string `gorm:"type:varchar(255);not null"`
	Path      string `gorm:"type:varchar(255);not null"`
	Visible   int    `gorm:"type:tinyint(1);default:1;not null"`
	Sort      int    `gorm:"type:int(11);default:0;not null"`
	IsHome    int    `gorm:"type:tinyint(1);default:0;not null"`
	IsFull    int    `gorm:"type:tinyint(1);default:0;not null"`
	PageSign  string `gorm:"type:varchar(255);index"`
	KeepAlive int    `gorm:"type:tinyint(1);default:0;not null"`

	Permissions []AdminPermission `gorm:"many2many:admin_menu_permission;"`
}

// DevMenus 开发菜单
func (AdminMenu) DevMenus() []AdminMenu {
	return []AdminMenu{
		{
			BaseModel: base.BaseModel{ID: 100000},
			Name:      "开发",
			Icon:      "fluent:window-dev-tools-20-regular",
			Path:      "/dev",
			Visible:   1,
		},
		{
			BaseModel: base.BaseModel{ID: 3},
			ParentId:  100000,
			Name:      "页面管理",
			Icon:      "iconoir:multiple-pages",
			Path:      "/dev/admin_page",
			Visible:   1,
			PageSign:  "admin_page",
		},
	}
}
