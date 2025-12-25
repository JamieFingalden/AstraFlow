package worker

import (
	"log"
)

// NewConsumer 创建MQ消费者
// 注意：此函数现在仅作保留，实际的OCR处理由Python消费者完成
// Go后端只负责发送任务到队列，不处理消费逻辑
func NewConsumer() error {
	log.Println("实际的OCR处理由Python消费者完成，Go后端只负责发送任务")
	return nil
}
