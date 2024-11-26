package db

import (
	"amis-base/internal/app/admin/models"
	baseModel "amis-base/internal/models"
	"github.com/gofiber/fiber/v2/log"
)

func Migration() {
	var modelsToMigrate = []interface{}{
		models.AdminRole{},
		models.AdminUser{},
		models.AdminPermission{},
		models.AdminMenu{},
		models.AdminPage{},
		models.AdminSetting{},
		baseModel.Token{},
	}

	for _, model := range modelsToMigrate {
		err := db.AutoMigrate(model)
		if err != nil {
			log.Fatal("Failed to migrate database: " + err.Error())
		}
	}

	go seed()
}

// 填充初始数据
func seed() {
}
