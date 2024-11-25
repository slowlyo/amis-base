package services

import (
	"amis-base/internal/app/admin/models"
	"amis-base/internal/pkg/db"
)

type Auth struct{}

func (a Auth) GetUserByUsername(username string) (models.AdminUser, error) {
	var user models.AdminUser

	return user, db.GetDB().Where("username = ?", username).First(&user).Error
}
