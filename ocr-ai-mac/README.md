# AstraFlow OCR AI MAC

这是一个独立的 OCR 处理项目，专门作为消费者，使用 RabbitMQ 从队列中获取并处理来自 AstraFlow 系统的 OCR 任务。

## 项目结构

```
ocr-ai-mac/
├── worker.py            # 任务消费者 - 从队列获取任务并处理
├── online_ocr_module.py # OCR 处理模块
├── ai_module.py         # AI 分类模块
└── requirements.txt     # 依赖包列表
```

## 安装依赖

```bash
pip install pika
```

## 环境变量配置

```bash
# RabbitMQ 配置
RABBITMQ_HOST=localhost
RABBITMQ_PORT=5672
RABBITMQ_USER=guest
RABBITMQ_PASSWORD=guest

# OpenAI API 配置（用于 ai_module.py）
OPENAI_API_KEY=your_openai_api_key
OPENAI_BASE_URL=https://api.openai.com/v1
```

## 使用方法

### 1. 启动消费者（worker）

```bash
python worker.py
```

注意：任务由 AstraFlow 的 Go 后端通过 RabbitMQ 队列发送，Python 项目只负责消费处理。

## 队列说明

- `ocr_tasks`: OCR 任务队列，存储待处理的任务
- `ocr_results`: OCR 结果队列，存储处理完成的结果

## 消息格式

### 任务消息格式
```json
{
  "task_id": "unique_task_id",
  "file_path": "/path/to/image.jpg",
  "callback_url": "optional_callback_url",
  "created_at": 1234567890.123
}
```

### 结果消息格式
```json
{
  "task_id": "unique_task_id",
  "status": "success|error",
  "data": {
    // AI 提取的发票字段
  },
  "ocr_text": "识别出的文本",
  "error_message": "错误信息（仅在失败时）",
  "processed_at": 1234567890.123
}
```