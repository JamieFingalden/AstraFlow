package mq

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

/*
Consumer：
RabbitMQ 消费者封装
职责：
- 从指定队列中消费消息
- 将消息反序列化为 OCRTask
- 交给业务 handler 处理
- 根据处理结果决定 Ack / Nack
*/
type Consumer struct {
	// RabbitMQ 的 channel
	// 注意：channel 不是并发安全的，一般一个 consumer 用一个 channel
	ch *amqp.Channel
}

/*
NewConsumer：
创建一个消费者实例

做的事情：
1. 从 connection 中创建一个 channel
2. 设置 QoS（限流）
  - prefetchCount = 1
  - 表示：RabbitMQ 在当前消息未 ack 前，不会再推送下一条消息
  - 适合 OCR / AI 这种“慢任务 + 重任务”的场景
*/
func NewConsumer(conn *amqp.Connection) (*Consumer, error) {
	// 从连接中创建一个 channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	// QoS 设置：一次只处理一条消息
	// 参数说明：
	// 1: prefetchCount，当前 consumer 同时最多处理 1 条未确认消息
	// 0: prefetchSize，按消息大小限流（一般不用）
	// false: 是否作用于整个 channel（false 表示只作用于当前 consumer）
	if err := ch.Qos(1, 0, false); err != nil {
		return nil, err
	}

	return &Consumer{ch: ch}, nil
}

/*
Consume：
开始消费消息

参数：
  - handler：业务处理函数
    接收一个 OCRTask，返回 error
  - error == nil   → 处理成功，Ack
  - error != nil   → 处理失败，Nack（重新入队）

流程：
1. 订阅队列
2. 循环读取消息
3. 反序列化消息体
4. 调用业务 handler
5. 根据结果 Ack / Nack
*/
func (c *Consumer) Consume(handler func(task OCRTask) error) error {
	// 开始消费队列
	msgs, err := c.ch.Consume(
		OCRQueueName, // 队列名
		"",           // consumerTag，空字符串表示由 RabbitMQ 自动生成
		false,        // autoAck=false，关闭自动确认，改为手动 Ack
		false,        // exclusive=false，非独占消费者
		false,        // noLocal=false（RabbitMQ 已废弃该参数，基本无影响）
		false,        // noWait=false，等待服务器响应
		nil,          // 额外参数
	)
	if err != nil {
		return err
	}

	// range 会一直阻塞，直到 channel 关闭
	for msg := range msgs {
		var task OCRTask

		// 1. 反序列化消息体
		// 如果消息格式不合法，说明是“脏数据”
		// → Nack + 不重新入队（否则会死循环）
		if err := json.Unmarshal(msg.Body, &task); err != nil {
			log.Println("解析消息失败:", err)
			msg.Nack(false, false) // requeue=false
			continue
		}

		// 2. 执行业务处理逻辑（OCR / AI 推理等）
		// 如果失败：
		// → 说明是临时错误（模型服务不可用等）
		// → Nack + 重新入队，稍后重试
		if err := handler(task); err != nil {
			log.Println("任务处理失败:", err)
			msg.Nack(false, true) // requeue=true
			continue
		}

		// 3. 处理成功
		// 手动 Ack，通知 RabbitMQ 可以删除该消息
		msg.Ack(false)
	}

	return nil
}
