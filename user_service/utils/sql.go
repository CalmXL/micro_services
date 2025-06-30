package utils

import (
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags),
	logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: true,
		ParameterizedQueries:      true,
		Colorful:                  true,
	},
)

func DBConnect(dbName string) (*gorm.DB, error) {
	dsn := "root:xL1210...@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "svc_", // user => svc_user
			SingularTable: true,
		},
	})

	return db, err
}

func DBClose(db *gorm.DB) {
	sqlDB, errDB := db.DB()
	if errDB != nil {
		log.Fatalf("Failed to start server: %v", errDB)
		return
	}

	defer sqlDB.Close()
}

func Paginate() {

}
