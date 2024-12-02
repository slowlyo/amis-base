package models

import base "amis-base/internal/models"

type AdminUser struct {
	base.BaseModel

	Username string      `gorm:"unique;type:varchar(255);not null" json:"username"`
	Password string      `gorm:"type:varchar(255);not null;" json:"-"`
	Name     string      `gorm:"type:varchar(255);not null" json:"name"`
	Avatar   string      `gorm:"type:varchar(255);not null" json:"avatar"`
	Enabled  int         `gorm:"type:tinyint(1);default:1;not null" json:"enabled"`
	Roles    []AdminRole `gorm:"many2many:admin_user_role;" json:"roles"`
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
