package main

import (
	"AstraFlow-go/internal/database"
	"AstraFlow-go/internal/model"
	"AstraFlow-go/pkg/config"
	"fmt"
	"log"
)

func main() {
	// 1. 初始化配置
	config.InitConfig()

	// 2. 初始化数据库
	database.InitDB()

	// 3. 自动迁移数据库表结构
	fmt.Println("正在同步数据库表结构...")
	err := database.DB.AutoMigrate(
		&model.Role{},
		&model.Tenant{},
		&model.User{},
		&model.Invoice{},
		&model.Attachment{},
	)
	if err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	fmt.Println("数据库迁移完成")
}
