package main

import (
	"AstraFlow-go/internal/database"
	"AstraFlow-go/internal/model"
	"AstraFlow-go/pkg/config"
	"fmt"
	"log"
)

func main() {
	// 初始化配置
	config.InitConfig()

	// 初始化数据库
	database.InitDB()

	// 自动迁移数据库表
	err := database.DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	fmt.Println("数据库迁移完成")
}
