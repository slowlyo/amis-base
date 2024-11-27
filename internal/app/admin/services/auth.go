package services

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/pkg/db"
)

type Auth struct{}

// GetUserByUsername 根据用户名获取用户
func (a *Auth) GetUserByUsername(username string) (models.AdminUser, error) {
	var user models.AdminUser

	return user, db.GetDB().Where("username = ?", username).First(&user).Error
}

// GetUserById 根据用户id获取用户
func (a *Auth) GetUserById(id uint) (models.AdminUser, error) {
	var user models.AdminUser

	return user, db.GetDB().Preload("Roles").Where("enabled = ?", 1).First(&user, id).Error
}
