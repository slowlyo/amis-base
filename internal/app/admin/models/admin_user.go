package models

import (
	"amis-base/internal/app/admin/types"
	base "amis-base/internal/models"
	"encoding/json"
	"github.com/duke-git/lancet/v2/slice"
)

type AdminUser struct {
	base.BaseModel

	Username string      `gorm:"unique;type:varchar(255);not null" json:"username"`
	Password string      `gorm:"type:varchar(255);not null;" json:"-"`
	Name     string      `gorm:"type:varchar(255);not null" json:"name"`
	Avatar   string      `gorm:"type:varchar(255);not null" json:"avatar"`
	Enabled  int         `gorm:"type:tinyint(1);default:1;not null" json:"enabled"`
	Roles    []AdminRole `gorm:"many2many:admin_user_role;" json:"roles"`
}

// IsSuperAdmin 判断是否是超级管理员
func (u *AdminUser) IsSuperAdmin() bool {
	return slice.Some(u.Roles, func(index int, item AdminRole) bool {
		return item.Sign == types.SuperAdminSign
	})
}

// Permissions 获取用户权限
func (u *AdminUser) Permissions() []AdminPermission {
	result := make([]AdminPermission, 0)
	for _, role := range u.Roles {
		for _, permission := range role.Permissions {
			result = append(result, permission)
		}
	}

	return result
}

// PermissionApiRules 获取用户权限的 api 规则
func (u *AdminUser) PermissionApiRules() []string {
	result := make([]string, 0)
	permissions := u.Permissions()
	for _, permission := range permissions {
		var apis []string
		if err := json.Unmarshal([]byte(permission.Api), &apis); err == nil {
			result = append(result, apis...)
		}
	}

	return result
}
