package db

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"
)

var (
	db   *gorm.DB
	once sync.Once
)

func Bootstrap() {
	Connect()

	if viper.GetBool("database.migrate") {
		Migration()
	}
}

func Connect() {
	once.Do(func() {
		var err error

		driver := viper.GetString("database.driver")

		switch driver {
		case "mysql":
			db, err = gorm.Open(mysql.Open(viper.GetString("database.dsn")), &gorm.Config{
				DisableForeignKeyConstraintWhenMigrating: true,
				Logger:                                   getLogger(),
			})
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

func getLogger() logger.Interface {
	logLevel := logger.Silent

	// 开发时打印所有 SQL
	if viper.GetBool("app.dev") {
		logLevel = logger.Info
	}

	return logger.New(
		log.New(os.Stdout, "\r\n", log.Flags()),
		logger.Config{
			SlowThreshold:             5 * time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			LogLevel:                  logLevel,
		},
	)
}
