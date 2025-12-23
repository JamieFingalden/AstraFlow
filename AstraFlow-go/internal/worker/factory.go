package worker

import (
	"log"

	"github.com/streadway/amqp"
	"AstraFlow-go/internal/mq"
)

// NewConsumer 创建MQ消费者
// 简单的辅助函数，用于创建消费者实例
func NewConsumer(conn *amqp.Connection) (*mq.Consumer, error) {
	consumer, err := mq.NewConsumer(conn)
	if err != nil {
		log.Printf("MQ Consumer 初始化失败: %v", err)
		return nil, err
	}
	return consumer, nil
}