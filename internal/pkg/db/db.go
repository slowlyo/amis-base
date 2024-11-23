package db

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func Bootstrap() {
	connect()
}

func connect() {
	once.Do(func() {
		var err error

		driver := viper.GetString("database.driver")

		switch driver {
		case "mysql":
			db, err = gorm.Open(mysql.Open(viper.GetString("database.dsn")))
		default:
			log.Fatal("Unsupported database driver: " + driver)
		}

		if err != nil {
			log.Fatal("Failed to connect to database: " + err.Error())
		}
	})
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database connection not initialized")
	}

	return db
}
