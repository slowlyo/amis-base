package db

import (
	"amis-base/internal/app/admin/models"
	baseModel "amis-base/internal/models"
	"github.com/gofiber/fiber/v2/log"
)

func Migration() {
	err := db.AutoMigrate(
		models.AdminRole{},
		models.AdminUser{},
		models.AdminPermission{},
		models.AdminMenu{},
		models.AdminPage{},
		models.AdminSetting{},
		baseModel.Token{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database: " + err.Error())
	}

	// 填充数据
	Seed()
}
