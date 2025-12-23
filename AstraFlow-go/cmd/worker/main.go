package main

import (
	"AstraFlow-go/internal/database"
	"AstraFlow-go/internal/mq"
	"AstraFlow-go/internal/worker"
	"log"
)

func main() {
	// 1. 建立与RabbitMQ的连接
	conn, err := mq.NewConnection()
	if err != nil {
		log.Fatal("MQ 连接失败:", err)
	}
	defer conn.Close()

	// 2. 初始化数据库连接
	database.InitDB()

	// 3. 创建消费者实例
	consumer, err := mq.NewConsumer(conn)
	if err != nil {
		log.Fatal("MQ Consumer 初始化失败:", err)
	}

	log.Println("Worker started, waiting for OCR tasks...")

	// 4. 启动工作服务，开始消费OCR任务
	// ProcessOCRTask函数直接处理每个收到的任务
	if err := worker.StartWorker(consumer); err != nil {
		log.Fatal(err)
	}
}
