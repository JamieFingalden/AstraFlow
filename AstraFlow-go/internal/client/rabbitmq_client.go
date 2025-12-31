package client

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

// RabbitMQOCRClient RabbitMQ OCR 客户端
type RabbitMQOCRClient struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	// 用于存储等待结果的回调
	resultCallbacks map[string]chan *OCRQueueResult
	callbackMutex   sync.RWMutex
}

// OCRQueueTask OCR 队列任务结构
type OCRQueueTask struct {
	TaskID      string  `json:"task_id"`
	FileID      int64   `json:"file_id"`
	FilePath    string  `json:"file_path"`
	CallbackURL *string `json:"callback_url,omitempty"`
	CreatedAt   float64 `json:"created_at"`
}

// OCRQueueResult OCR 队列结果结构
type OCRQueueResult struct {
	TaskID      string                 `json:"task_id"`
	Status      string                 `json:"status"`
	Data        map[string]interface{} `json:"data,omitempty"`
	OCRText     string                 `json:"ocr_text,omitempty"`
	ErrorMsg    string                 `json:"error_message,omitempty"`
	ProcessedAt float64                `json:"processed_at"`
}

// NewRabbitMQOCRClient 创建新的 RabbitMQ OCR 客户端
func NewRabbitMQOCRClient() (*RabbitMQOCRClient, error) {
	rabbitmqHost := getEnvOrDefault("RABBITMQ_HOST", "localhost")
	rabbitmqPort := getEnvOrDefault("RABBITMQ_PORT", "5672")
	rabbitmqUser := getEnvOrDefault("RABBITMQ_USER", "admin")
	rabbitmqPassword := getEnvOrDefault("RABBITMQ_PASSWORD", "password")

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

	// 声明队列
	_, err = channel.QueueDeclare(
		"ocr_tasks", // 队列名
		true,        // durable - 持久化
		false,       // autoDelete
		false,       // exclusive
		false,       // noWait
		nil,         // args
	)
	if err != nil {
		channel.Close()
		conn.Close()
		return nil, fmt.Errorf("声明任务队列失败: %v", err)
	}

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

	client := &RabbitMQOCRClient{
		conn:            conn,
		channel:         channel,
		resultCallbacks: make(map[string]chan *OCRQueueResult),
	}

	// 启动结果监听器
	go client.listenForResults()

	return client, nil
}

// AddTask 添加 OCR 任务到队列
func (c *RabbitMQOCRClient) AddTask(fileID int64, filePath string, callbackURL *string) (string, error) {
	taskID := generateUUID()
	task := OCRQueueTask{
		TaskID:      taskID,
		FileID:      fileID,
		FilePath:    filePath, // 使用Go服务端的文件路径格式，如 "/uploads/filename"
		CallbackURL: callbackURL,
		CreatedAt:   float64(time.Now().Unix()),
	}

	taskJSON, err := json.Marshal(task)
	if err != nil {
		return "", fmt.Errorf("序列化任务失败: %v", err)
	}

	// 创建结果通道用于等待结果
	resultChan := make(chan *OCRQueueResult, 1)
	c.callbackMutex.Lock()
	c.resultCallbacks[taskID] = resultChan
	c.callbackMutex.Unlock()

	// 发布消息到队列
	err = c.channel.Publish(
		"",          // exchange
		"ocr_tasks", // routing key
		false,       // mandatory
		false,       // immediate
		amqp091.Publishing{
			ContentType:  "application/json",
			Body:         taskJSON,
			DeliveryMode: amqp091.Persistent, // 消息持久化
		},
	)
	if err != nil {
		// 清理结果通道
		c.callbackMutex.Lock()
		delete(c.resultCallbacks, taskID)
		c.callbackMutex.Unlock()
		close(resultChan)
		return "", fmt.Errorf("发布任务消息失败: %v", err)
	}

	log.Printf("任务 %s 已添加到队列，文件路径: %s", taskID, filePath)
	return taskID, nil
}

// GetResult 等待并获取 OCR 任务结果
func (c *RabbitMQOCRClient) GetResult(taskID string, timeout int) (*OCRQueueResult, error) {
	c.callbackMutex.RLock()
	resultChan, exists := c.resultCallbacks[taskID]
	c.callbackMutex.RUnlock()

	if !exists {
		return nil, fmt.Errorf("任务ID不存在: %s", taskID)
	}

	// 等待结果或超时
	select {
	case result := <-resultChan:
		// 清理结果通道
		c.callbackMutex.Lock()
		delete(c.resultCallbacks, taskID)
		c.callbackMutex.Unlock()
		close(resultChan)
		return result, nil
	case <-time.After(time.Duration(timeout) * time.Second):
		// 清理结果通道
		c.callbackMutex.Lock()
		delete(c.resultCallbacks, taskID)
		c.callbackMutex.Unlock()
		close(resultChan)
		return nil, fmt.Errorf("获取结果超时，任务ID: %s", taskID)
	}
}

// listenForResults 监听结果队列
func (c *RabbitMQOCRClient) listenForResults() {
	msgs, err := c.channel.Consume(
		"ocr_results", // 队列名
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		log.Printf("监听结果队列失败: %v", err)
		return
	}

	log.Println("开始监听结果队列...")
	for msg := range msgs {
		var result OCRQueueResult
		if err := json.Unmarshal(msg.Body, &result); err != nil {
			log.Printf("解析结果消息失败: %v", err)
			continue
		}

		// 查找对应的任务回调
		c.callbackMutex.RLock()
		resultChan, exists := c.resultCallbacks[result.TaskID]
		c.callbackMutex.RUnlock()

		if exists {
			// 发送结果到对应的通道
			select {
			case resultChan <- &result:
				// 成功发送结果
			default:
				// 通道已满或已关闭，忽略
				log.Printf("结果通道已满或已关闭，任务ID: %s", result.TaskID)
			}
		} else {
			log.Printf("未找到对应的任务回调，任务ID: %s", result.TaskID)
		}
	}
}

// ConsumeTasks 消费任务队列，处理传入的 OCR 任务
func (c *RabbitMQOCRClient) ConsumeTasks(processFunc func(OCRQueueTask) error) error {
	msgs, err := c.channel.Consume(
		"ocr_tasks", // 队列名
		"",          // consumer
		false,       // auto-ack - 手动确认以确保消息处理完成
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		return fmt.Errorf("监听任务队列失败: %v", err)
	}

	log.Println("开始监听 OCR 任务队列...")
	for msg := range msgs {
		var task OCRQueueTask
		if err := json.Unmarshal(msg.Body, &task); err != nil {
			log.Printf("解析任务消息失败: %v", err)
			// 拒绝消息并重新入队
			msg.Nack(false, true)
			continue
		}

		log.Printf("接收到 OCR 任务: taskID=%s fileID=%d", task.TaskID, task.FileID)

		// 处理任务
		if err := processFunc(task); err != nil {
			log.Printf("处理 OCR 任务失败: %v", err)
			// 拒绝消息并重新入队
			msg.Nack(false, true)
			continue
		}

		// 手动确认消息
		if err := msg.Ack(false); err != nil {
			log.Printf("确认消息失败: %v", err)
		}

		log.Printf("OCR 任务处理完成: taskID=%s", task.TaskID)
	}

	return nil
}

// getEnvOrDefault 获取环境变量，如果不存在则返回默认值
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// generateUUID 生成简单的 UUID（实际应用中可以使用更复杂的实现）
func generateUUID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// Close 关闭连接
func (c *RabbitMQOCRClient) Close() error {
	if c.channel != nil {
		c.channel.Close()
	}
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
