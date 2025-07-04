package utils

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var newLogger = logger.New(
	// os.Stdout: 将日志输出到控制台
	// \r\n: 每条日志前缀，这里换行
	// log.LstdFlags: 日志中自动包含日期和时间戳
	log.New(os.Stdout, "\r\n", log.LstdFlags),
	logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		// 忽略 gorm.ErrRecordNotFound 错误的日志输出。避免因为正常业务逻辑中的“未找到记录”错误而产生过多警告日志。
		IgnoreRecordNotFoundError: true,
		// 输出参数化后的 SQL 查询（而不是实际传入的参数值），提高安全性，同时减少日志重复。例如输出 WHERE id = ? 而不是 WHERE id = 123。
		ParameterizedQueries: true,
		// 在终端中启用彩色日志输出，便于阅读。比如 Info 是绿色，Warn 是黄色，Error 是红色等。
		Colorful: true,
	},
)

func DBConnect(dbName string) (*gorm.DB, error) {
	dsn := "root:xL1210...@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "svc_", // user => svc_user
			SingularTable: true,   // 表明单数形式
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

func Paginate(page, pageSize int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		// Offset: 跳过前 n 条记录
		// Limit: 取 n 条记录
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
