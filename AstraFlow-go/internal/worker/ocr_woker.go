package worker

import (
	"AstraFlow-go/internal/mq"
	"log"
)

// StartWorker 启动OCR工作服务
// 这是一个简单的函数，直接处理从队列收到的任务
func StartWorker(consumer *mq.Consumer) error {
	log.Println("OCR Worker starting, waiting for OCR tasks...")

	// 消费队列中的消息，每个消息都是一个OCR任务
	return consumer.Consume(func(task mq.OCRTask) error {
		// 直接调用处理函数处理OCR任务
		return ProcessOCRTask(task.FileID, task.FilePath)
	})
}
