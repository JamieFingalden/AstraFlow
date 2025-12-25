package service

import (
	"AstraFlow-go/internal/client"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

// OCRResultListener OCR 结果监听器
type OCRResultListener struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
}

// NewOCRResultListener 创建新的 OCR 结果监听器
func NewOCRResultListener() (*OCRResultListener, error) {
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	if rabbitmqHost == "" {
		rabbitmqHost = "localhost"
	}
	rabbitmqPort := os.Getenv("RABBITMQ_PORT")
	if rabbitmqPort == "" {
		rabbitmqPort = "5672"
	}
	rabbitmqUser := os.Getenv("RABBITMQ_USER")
	if rabbitmqUser == "" {
		rabbitmqUser = "guest"
	}
	rabbitmqPassword := os.Getenv("RABBITMQ_PASSWORD")
	if rabbitmqPassword == "" {
		rabbitmqPassword = "guest"
	}

	uri := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitmqUser, rabbitmqPassword, rabbitmqHost, rabbitmqPort)

	conn, err := amqp091.Dial(uri)
	if err != nil {
		return nil, fmt.Errorf("连接 RabbitMQ 失败: %v", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("创建通道失败: %v", err)
	}

	// 声明结果队列
	_, err = channel.QueueDeclare(
		"ocr_results", // 队列名
		true,          // durable - 持久化
		false,         // autoDelete
		false,         // exclusive
		false,         // noWait
		nil,           // args
	)
	if err != nil {
		channel.Close()
		conn.Close()
		return nil, fmt.Errorf("声明结果队列失败: %v", err)
	}

	listener := &OCRResultListener{
		conn:    conn,
		channel: channel,
	}

	return listener, nil
}

// ListenForResults 监听 OCR 结果队列
func (l *OCRResultListener) ListenForResults() error {
	msgs, err := l.channel.Consume(
		"ocr_results", // 队列名
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		return fmt.Errorf("监听结果队列失败: %v", err)
	}

	log.Println("开始监听 OCR 结果队列...")
	for msg := range msgs {
		go l.processResultMessage(msg.Body)
	}

	return nil
}

// processResultMessage 处理结果消息
func (l *OCRResultListener) processResultMessage(body []byte) {
	var result client.OCRQueueResult
	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("解析结果消息失败: %v", err)
		return
	}

	log.Printf("接收到 OCR 结果，任务ID: %s，状态: %s", result.TaskID, result.Status)

	if result.Status == "success" {
		// 从任务ID中提取文件ID（这需要在任务创建时建立映射）
		// 这里需要根据实际的实现来处理
		log.Printf("OCR 任务成功完成: %s", result.TaskID)

		// 在实际实现中，这里应该更新数据库记录
		// 但现在我们先打印日志
	} else {
		log.Printf("OCR 任务失败: %s, 错误: %s", result.TaskID, result.ErrorMsg)
	}
}

// Close 关闭连接
func (l *OCRResultListener) Close() error {
	if l.channel != nil {
		l.channel.Close()
	}
	if l.conn != nil {
		return l.conn.Close()
	}
	return nil
}

// StartOCRResultListener 启动 OCR 结果监听器
func StartOCRResultListener() {
	listener, err := NewOCRResultListener()
	if err != nil {
		log.Printf("创建 OCR 结果监听器失败: %v", err)
		return
	}
	defer listener.Close()

	// 启动监听
	if err := listener.ListenForResults(); err != nil {
		log.Printf("监听 OCR 结果时发生错误: %v", err)
	}
}
