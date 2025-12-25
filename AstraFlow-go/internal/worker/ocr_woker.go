package worker

import (
	"log"
)

// StartWorker 启动Python消费者工作服务
// 注意：此函数现在仅作保留，实际的OCR处理由Python消费者完成
// Go后端只负责发送任务到队列，不处理消费逻辑
func StartWorker() error {
	log.Println("Python Consumer Worker started, but actual OCR processing handled by Python consumers...")
	log.Println("Go端仅负责发送任务到队列，Python消费者负责处理OCR任务")

	// 保持此函数运行，但不执行实际的消费任务
	// 实际的OCR处理由Python消费者完成
	select {}
}
