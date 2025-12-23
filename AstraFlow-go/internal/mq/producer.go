package mq

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

/*
OCRQueueName 定义 OCR 任务队列的名字
- Web 服务负责往这个队列里投递任务
- Worker 服务负责消费这个队列
- 名字是 Web 和 Worker 之间的“契约”，必须一致
*/
const OCRQueueName = "ocr_tasks"

/*
Producer 是消息生产者
本质上就是对 RabbitMQ Channel 的一层封装
职责：只负责「把任务丢进队列」
*/
type Producer struct {
	ch *amqp.Channel // RabbitMQ 的 channel，用于发布消息
}

/*
NewProducer 创建一个新的 Producer

参数：
- conn: 已经建立好的 RabbitMQ 连接（amqp.Connection）

这里做了两件事：
1. 从连接中创建一个 channel
2. 确保 OCR 任务队列存在（QueueDeclare 是幂等的）
*/
func NewProducer(conn *amqp.Connection) (*Producer, error) {
	// 从连接中创建一个 channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	/*
		声明一个队列（如果已存在则不会重复创建）

		参数说明：
		- name: 队列名
		- durable: true
			→ 队列持久化（RabbitMQ 重启后队列仍然存在）
		- autoDelete: false
			→ 不会因为没有消费者就自动删除
		- exclusive: false
			→ 不只允许一个连接使用（Web 和 Worker 都能用）
		- noWait: false
			→ 等待服务器确认
		- args: nil
			→ 不使用额外参数（如 TTL、死信队列等）
	*/
	_, err = ch.QueueDeclare(
		OCRQueueName,
		true,  // durable：队列持久化
		false, // auto-delete：不自动删除
		false, // exclusive：非独占
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}

	return &Producer{ch: ch}, nil
}

/*
PublishOCRTask 向 OCR 队列发送一条任务消息

参数：
- task: OCRTask（你定义的 OCR 任务结构体）

流程：
1. 将任务结构体序列化为 JSON
2. 通过 channel.Publish 投递到队列
*/
func (p *Producer) PublishOCRTask(task OCRTask) error {
	// 将 OCR 任务结构体序列化成 JSON
	body, err := json.Marshal(task)
	if err != nil {
		return err
	}

	/*
		发布消息到队列

		参数说明：
		- exchange: ""
			→ 使用默认交换机
		- routingKey: OCRQueueName
			→ 默认交换机下，routingKey = 队列名
		- mandatory: false
		- immediate: false
	*/
	return p.ch.Publish(
		"",
		OCRQueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json", // 标明消息类型
			Body:        body,               // 实际消息内容
		},
	)
}
