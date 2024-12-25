package models

import base "amis-base/internal/models"

type AdminPermission struct {
	base.BaseModel

	ParentId uint   `gorm:"type:int(11);not null;default:0" json:"parent_id"`
	Name     string `gorm:"type:varchar(255);not null;unique" json:"name"`
	Sign     string `gorm:"type:varchar(255);not null;unique" json:"sign"`
	Api      string `gorm:"type:text" json:"api"`
	Sort     uint   `gorm:"type:int(11);not null;default:0" json:"sort"`

	Roles    []AdminRole       `gorm:"many2many:admin_role_permission;" json:"roles"`
	Menus    []AdminMenu       `gorm:"many2many:admin_menu_permission;" json:"menus"`
	Children []AdminPermission `gorm:"-" json:"children"`
}
