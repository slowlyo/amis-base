package models

import base "amis-base/internal/models"

type AdminMenu struct {
	base.BaseModel

	ParentId uint   `gorm:"type:int(11);not null;default:0"`
	Name     string `gorm:"type:varchar(255);not null"`
	Icon     string `gorm:"type:varchar(255);not null"`
	Path     string `gorm:"type:varchar(255);not null"`
	Visible  int    `gorm:"type:tinyint(1);default:1;not null"`
	Sort     int    `gorm:"type:int(11);default:0;not null"`
	IsHome   int    `gorm:"type:tinyint(1);default:0;not null"`
	IsFull   int    `gorm:"type:tinyint(1);default:0;not null"`
	PageSign string `gorm:"type:varchar(255)"`

	Permissions []AdminPermission `gorm:"many2many:admin_menu_permission;"`
}
