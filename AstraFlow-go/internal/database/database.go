package database

import (
	"AstraFlow-go/internal/model"
	"AstraFlow-go/pkg/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 全局数据库连接实例
var DB *gorm.DB

// InitDB 初始化数据库连接
// 从配置文件读取数据库连接参数，建立MySQL连接
func InitDB() {
	cfg := config.Cfg.MySQL

	// 构建DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.Charset, cfg.ParseTime, cfg.Loc)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // 使用默认配置（复数表名）
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 自动迁移数据库表 - 确保所有模型表结构同步
	err = DB.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Tenant{},
		&model.Invoice{},
		&model.OCRResult{},
		&model.Attachment{},
		&model.Reimbursement{},
		&model.ReimbursementItem{},
	)
	if err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	log.Println("数据库连接成功")
}
