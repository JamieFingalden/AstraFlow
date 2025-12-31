# -*- coding: utf-8 -*-
"""
OCR 队列工作进程：从 RabbitMQ 队列中获取任务并处理 OCR 识别
"""

import json
import pika
import time
import os
import ssl
import requests
from pathlib import Path
from online_ocr_module import ocr_image
from ai_module import extract_invoice_fields_with_openai
import logging
from requests.adapters import HTTPAdapter
from urllib3.util.retry import Retry

# 配置日志
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)

# RabbitMQ 配置
RABBITMQ_HOST = os.getenv('RABBITMQ_HOST', 'localhost')
RABBITMQ_PORT = int(os.getenv('RABBITMQ_PORT', 5672))
RABBITMQ_USER = os.getenv('RABBITMQ_USER', 'admin')
RABBITMQ_PASSWORD = os.getenv('RABBITMQ_PASSWORD', 'password')
OCR_QUEUE = 'ocr_tasks'
RESULT_QUEUE = 'ocr_results'

class OCRWorker:
    def __init__(self):
        """初始化 OCR 工作进程"""
        credentials = pika.PlainCredentials(RABBITMQ_USER, RABBITMQ_PASSWORD)
        self.connection = pika.BlockingConnection(
            pika.ConnectionParameters(host=RABBITMQ_HOST, port=RABBITMQ_PORT, credentials=credentials)
        )
        self.channel = self.connection.channel()
        
        # 声明队列（如果不存在则创建）
        self.channel.queue_declare(queue=OCR_QUEUE, durable=True)  # OCR任务队列
        self.channel.queue_declare(queue=RESULT_QUEUE, durable=True)  # 结果队列
        
        # 设置每次只处理一个消息
        self.channel.basic_qos(prefetch_count=1)
        
        self.running = True
        
        # 创建带有重试策略的请求会话
        self.session = self.create_session_with_retries()
    
    def create_session_with_retries(self):
        """创建具有重试策略的请求会话"""
        session = requests.Session()
        
        # 配置重试策略
        retry_strategy = Retry(
            total=3,  # 总重试次数
            backoff_factor=1,  # 重试间隔的倍数因子
            status_forcelist=[429, 500, 502, 503, 504],  # 需要重试的状态码
            allowed_methods=["HEAD", "GET", "OPTIONS", "POST"]  # 允许重试的HTTP方法
        )
        
        adapter = HTTPAdapter(max_retries=retry_strategy)
        session.mount("http://", adapter)
        session.mount("https://", adapter)
        
        # 设置SSL上下文选项以处理SSL连接问题
        session.verify = True  # 验证SSL证书
        
        return session
        
    def process_task(self, ch, method, properties, body):
        """
        处理单个 OCR 任务
        
        Args:
            ch: 通道对象
            method: 方法帧
            properties: 属性
            body: 消息体
        """
        try:
            # 解析任务数据
            task_data = json.loads(body)
            task_id = task_data.get('task_id')
            file_path = task_data.get('file_path')
            callback_url = task_data.get('callback_url')  # 可选的回调 URL
            
            logger.info(f"开始处理任务 {task_id}，文件路径: {file_path}")
            
            import tempfile
            
            # 从Go后端获取文件
            # 队列中的路径是 Go 后端的格式，如 "./uploads/filename"，我们需要提取文件名并从Go后端的HTTP服务获取
            # 提取文件名部分
            if file_path.startswith('./'):
                file_path = file_path[2:]  # 移除 "./" 前缀
            
            # 获取Go后端基础URL
            backend_base_url = os.getenv('BACKEND_BASE_URL', 'http://localhost:8080')
            
            # 确保backend_base_url格式正确
            if not backend_base_url.endswith('/'):
                backend_base_url += '/'
            
            # 确保file_path以正确的格式（不以/开头）
            if file_path.startswith('/'):
                file_path = file_path[1:]
            
            # 构造完整的文件URL
            file_url = f"{backend_base_url}{file_path}"
            
            # 从Go后端下载文件到临时位置
            # 使用会话和重试机制以处理网络问题
            try:
                response = self.session.get(file_url, timeout=30)
                response.raise_for_status()
            except requests.exceptions.SSLError as ssl_err:
                logger.error(f"SSL错误: {str(ssl_err)}")
                # 尝试不验证SSL证书（仅用于开发环境）
                response = requests.get(file_url, timeout=30, verify=False)
                response.raise_for_status()
            except requests.exceptions.RequestException as e:
                logger.error(f"下载文件失败: {str(e)}")
                raise e
            
            # 创建临时文件
            file_extension = os.path.splitext(file_path)[1]
            tmp_file_name = tempfile.NamedTemporaryFile(delete=False, suffix=file_extension).name
            
            with open(tmp_file_name, 'wb') as tmp_file:
                tmp_file.write(response.content)
            
            full_file_path = tmp_file_name
            
            # 执行 OCR 识别
            ocr_text = ocr_image(full_file_path)
            
            # 如果 OCR 没有识别出文字，返回错误
            if not ocr_text.strip():
                raise ValueError("未能识别出图片中的文字")
            
            # 使用 AI 模型进行分类
            # 使用原始文件路径（从消息中获取）而不是临时文件路径
            category = extract_invoice_fields_with_openai(task_data.get('file_path', ''), ocr_text)
            
            # 准备结果
            result = {
                'task_id': task_id,
                'status': 'success',
                'data': category,
                'ocr_text': ocr_text,
                'processed_at': time.time()
            }
            
            logger.info(f"任务 {task_id} 处理成功")
            
            # 将结果发送到结果队列
            self.channel.basic_publish(
                exchange='',
                routing_key=RESULT_QUEUE,
                body=json.dumps(result),
                properties=pika.BasicProperties(
                    delivery_mode=2,  # 消息持久化
                )
            )
            
            # 如果有回调 URL，发送结果
            if callback_url:
                self.send_callback(callback_url, result)
                
            # 确认消息已处理
            ch.basic_ack(delivery_tag=method.delivery_tag)
            
            # 清理临时文件
            try:
                os.unlink(full_file_path)
                logger.info(f"临时文件已清理: {full_file_path}")
            except Exception as e:
                logger.error(f"清理临时文件失败: {e}")
            
        except Exception as e:
            logger.error(f"任务处理失败: {str(e)}", exc_info=True)  # 添加详细的异常信息
            
            # 如果任务数据解析失败，获取任务ID
            task_id = 'unknown'
            try:
                task_data = json.loads(body)
                task_id = task_data.get('task_id', 'unknown')
            except json.JSONDecodeError as json_err:
                logger.error(f"解析任务数据失败: {str(json_err)}")
            except Exception as parse_err:
                logger.error(f"解析任务数据时发生未知错误: {str(parse_err)}")
            
            # 记录错误结果
            error_result = {
                'task_id': task_id,
                'status': 'error',
                'error_message': str(e),
                'processed_at': time.time()
            }
            
            # 尝试将错误结果发送到结果队列
            try:
                self.channel.basic_publish(
                    exchange='',
                    routing_key=RESULT_QUEUE,
                    body=json.dumps(error_result),
                    properties=pika.BasicProperties(
                        delivery_mode=2,  # 消息持久化
                    )
                )
                logger.info(f"错误结果已发送到结果队列，任务ID: {task_id}")
            except Exception as publish_err:
                logger.error(f"发送错误结果到队列失败: {str(publish_err)}")
            
            # 确认消息已处理（即使处理失败）
            try:
                ch.basic_ack(delivery_tag=method.delivery_tag)
                logger.info(f"消息确认完成，任务ID: {task_id}")
            except Exception as ack_err:
                logger.error(f"确认消息失败: {str(ack_err)}")
            
            # 清理临时文件（如果存在）
            try:
                if 'full_file_path' in locals() and full_file_path and os.path.exists(full_file_path):
                    os.unlink(full_file_path)
                    logger.info(f"临时文件已清理: {full_file_path}")
            except FileNotFoundError:
                logger.warning(f"临时文件不存在，无需清理: {full_file_path}")
            except Exception as cleanup_error:
                logger.error(f"清理临时文件失败: {cleanup_error}")
    
    def send_callback(self, callback_url, result):
        """
        向指定 URL 发送回调结果
        
        Args:
            callback_url (str): 回调 URL
            result (dict): 处理结果
        """
        try:
            # 使用会话和重试机制以处理网络问题
            try:
                response = self.session.post(
                    callback_url,
                    json=result,
                    timeout=30
                )
                response.raise_for_status()
                logger.info(f"回调成功发送到 {callback_url}")
            except requests.exceptions.SSLError as ssl_err:
                logger.error(f"回调发送时发生SSL错误: {str(ssl_err)}")
                # 尝试不验证SSL证书（仅用于开发环境）
                response = requests.post(
                    callback_url,
                    json=result,
                    timeout=30,
                    verify=False
                )
                response.raise_for_status()
                logger.info(f"回调成功发送到 {callback_url} (忽略SSL验证)")
            except requests.exceptions.RequestException as e:
                logger.error(f"回调发送失败到 {callback_url}: {str(e)}")
                raise e
        except Exception as e:
            logger.error(f"回调发送失败到 {callback_url}: {str(e)}")
    
    def run(self):
        """运行工作进程，持续监听队列"""
        logger.info("OCR 工作进程启动，开始监听任务队列...")
        
        # 设置消息处理函数
        self.channel.basic_consume(
            queue=OCR_QUEUE,
            on_message_callback=self.process_task
        )
        
        try:
            # 开始消费消息
            logger.info("等待任务消息...")
            self.channel.start_consuming()
        except KeyboardInterrupt:
            logger.info("收到中断信号，正在停止工作进程...")
            self.channel.stop_consuming()
            self.connection.close()
            logger.info("OCR 工作进程已停止")

def main():
    """主函数"""
    worker = OCRWorker()
    worker.run()

if __name__ == "__main__":
    main()