package models

import base "amis-base/internal/models"

type AdminUser struct {
	base.BaseModel

	Username string      `gorm:"unique;type:varchar(255);not null"`
	Password string      `gorm:"type:varchar(255);not null;" json:"-"`
	Name     string      `gorm:"type:varchar(255);not null"`
	Avatar   string      `gorm:"type:varchar(255);not null"`
	Enabled  int         `gorm:"type:tinyint(1);default:1;not null"`
	Roles    []AdminRole `gorm:"many2many:admin_user_role;"`
}

// IsAdministrator 判断是否是超级管理员
func (u AdminUser) IsAdministrator() bool {
	for _, roles := range u.Roles {
		if roles.Sign == "administrator" {
			return true
		}
	}

	return false
}
