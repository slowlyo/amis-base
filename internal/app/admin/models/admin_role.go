package models

import base "amis-base/internal/models"

type AdminRole struct {
	base.BaseModel

	Name string `gorm:"type:varchar(255);not null;unique" json:"name"`
	Sign string `gorm:"type:varchar(255);not null;unique" json:"sign"`

	Users       []AdminUser       `gorm:"many2many:admin_user_role;" json:"-"`
	Permissions []AdminPermission `gorm:"many2many:admin_role_permission;" json:"-"`
}
