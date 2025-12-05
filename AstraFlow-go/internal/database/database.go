package database

import (
	"AstraFlow-go/internal/model"
	"AstraFlow-go/pkg/config"
	"fmt"
	"log"
	"net/url"

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
	// 对密码和其他可能包含特殊字符的参数进行转义
	username := url.QueryEscape(cfg.User)
	password := url.QueryEscape(cfg.Password)
	host := cfg.Host // host一般不需要转义
	database := url.QueryEscape(cfg.Database)
	charset := url.QueryEscape(cfg.Charset)
	loc := url.QueryEscape(cfg.Loc)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, cfg.Port, database, charset, cfg.ParseTime, loc)

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
