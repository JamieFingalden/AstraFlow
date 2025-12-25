package main

import (
	"AstraFlow-go/internal/database"
	"AstraFlow-go/internal/worker"
	"AstraFlow-go/pkg/config"
	"log"
)

func main() {
	// 1. 初始化配置
	config.InitConfig()

	// 2. 初始化数据库连接
	database.InitDB()

	log.Println("Worker started (Python consumer handles OCR tasks)...")

	// 3. 启动Go端OCR工作进程作为备用方案
	go func() {
		log.Println("Starting Go OCR worker as fallback option...")
		if err := worker.StartOCRWorker(); err != nil {
			log.Printf("Go OCR worker stopped: %v", err)
		}
	}()

	// 4. 启动Python工作服务
	// 注意：实际的OCR处理主要由Python消费者完成，Go后端只负责发送任务
	if err := worker.StartWorker(); err != nil {
		log.Fatal(err)
	}
}
