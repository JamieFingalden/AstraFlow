# -*- coding: utf-8 -*-
"""
OCR 队列工作进程：从 RabbitMQ 队列中获取任务并处理 OCR 识别
"""

import os
import json
import pika
import requests
from ocr_module import ocr_image, extract_invoice_info
from pdf_to_png import pdf_to_png
import tempfile
import shutil

# RabbitMQ 配置
RABBITMQ_HOST = 'localhost'
RABBITMQ_PORT = 5672
RABBITMQ_USER = 'admin'
RABBITMQ_PASSWORD = 'password'
OCR_QUEUE = 'ocr_tasks'


class OCRWorker:
    def __init__(self):
        credentials = pika.PlainCredentials(RABBITMQ_USER, RABBITMQ_PASSWORD)
        self.connection = pika.BlockingConnection(
            pika.ConnectionParameters(host=RABBITMQ_HOST, port=RABBITMQ_PORT, credentials=credentials)
        )
        self.channel = self.connection.channel()
        self.channel.queue_declare(queue=OCR_QUEUE, durable=True)  # OCR任务队列
        self.channel.basic_qos(prefetch_count=1)  # 设置每次只处理一个消息

    def process_task(self, ch, method, properties, body):
        task_data = json.loads(body)
        task_id = task_data.get('task_id')
        file_path = task_data.get('file_path')
        callback_url = task_data.get('/api/v1/callback/ocr-result')  # 可选的回调 URL
        file_id = task_data.get('file_id', 0)  # 从任务数据中获取 file_id

        try:
            # 从Go后端获取文件
            backend_base_url = 'http://localhost:8080'
            if not backend_base_url.endswith('/'):
                backend_base_url += '/'
            if file_path.startswith('/'):
                file_path = file_path[1:]
            file_url = f"{backend_base_url}{file_path}"

            response = requests.get(file_url, timeout=30)
            response.raise_for_status()

            file_extension = os.path.splitext(file_path)[1]
            tmp_file_name = tempfile.NamedTemporaryFile(delete=False, suffix=file_extension).name

            with open(tmp_file_name, 'wb') as tmp_file:
                tmp_file.write(response.content)

            # 如果是PDF文件，转换为PNG格式
            file_extension = os.path.splitext(tmp_file_name)[1].lower()
            if file_extension == '.pdf':
                temp_dir = tempfile.mkdtemp()
                png_files = pdf_to_png(tmp_file_name, temp_dir)
                if png_files:
                    image_path = png_files[0]  # 取第一个页面进行OCR
                else:
                    raise ValueError("PDF文件转换为PNG失败，没有生成任何图片文件")
                # 清理临时PDF文件
                os.unlink(tmp_file_name)
            else:
                image_path = tmp_file_name

            # 执行 OCR 识别
            ocr_text = ocr_image(image_path)

            # 如果 OCR 没有识别出文字，返回错误
            if not ocr_text.strip():
                raise ValueError("未能识别出图片中的文字")

            # 使用AI模块进行发票信息提取，传递OCR结果以避免重复识别
            print("开始模型处理")
            invoice_info = extract_invoice_info(image_path, use_ai=True, ocr_text=ocr_text)
            extracted_fields = invoice_info["extracted_fields"]

            # 准备回调数据
            callback_data = {
                "task_id": task_id,
                "status": "success",
                "data": {
                    "file_id": file_id,
                    "invoice_number": extracted_fields.get("invoice_number", ""),
                    "vendor": extracted_fields.get("vendor", ""),
                    "description": extracted_fields.get("description", ""),
                    "amount": extracted_fields.get("amount", 0),
                    "category": extracted_fields.get("category", ""),
                    "payment_source": extracted_fields.get("payment_source", ""),
                    "tax_id": extracted_fields.get("tax_id", ""),
                    "invoice_date": extracted_fields.get("invoice_date", ""),
                    "ocr_text": ocr_text
                }
            }

            print("OCR处理结果:")
            print(json.dumps(callback_data, indent=2, ensure_ascii=False))

            print("调用回调函数")
            callback_url = "/api/v1/callback/ocr-result"
            # 发送回调请求到Go后端
            if callback_url:
                try:
                    print("发送请求")
                    callback_response = requests.post(
                        "http://localhost:8080" + callback_url,
                        json=callback_data,
                        headers={"Content-Type": "application/json"},
                        timeout=30
                    )
                    print(f"回调请求状态码: {callback_response.status_code}")
                    print(f"回调响应: {callback_response.text}")
                except Exception as callback_error:
                    print(f"发送回调请求失败: {str(callback_error)}")
                    # 即使回调失败，也继续处理

            # 清理临时文件
            print("清理临时文件")
            os.unlink(image_path)
            if 'temp_dir' in locals():
                shutil.rmtree(temp_dir)

        except Exception as e:
            print(f"处理任务 {task_id} 时发生错误: {str(e)}")
            # 准备错误回调数据
            error_callback_data = {
                "task_id": task_id,
                "status": "failure",
                "data": {
                    "file_id": file_id
                },
                "error_message": str(e)
            }

            print("OCR处理错误:")
            print(json.dumps(error_callback_data, indent=2, ensure_ascii=False))

            # 发送错误回调请求到Go后端
            if callback_url:
                try:
                    callback_response = requests.post(
                        "http://localhost:8081" + callback_url,
                        json=error_callback_data,
                        headers={"Content-Type": "application/json"},
                        timeout=30
                    )
                    print(f"错误回调请求状态码: {callback_response.status_code}")
                    print(f"错误回调响应: {callback_response.text}")
                except Exception as callback_error:
                    print(f"发送错误回调请求失败: {str(callback_error)}")

            # 清理临时文件
            try:
                if 'tmp_file_name' in locals() and os.path.exists(tmp_file_name):
                    os.unlink(tmp_file_name)
                if 'temp_dir' in locals():
                    shutil.rmtree(temp_dir)
            except:
                pass

        finally:
            # 确认消息已处理
            ch.basic_ack(delivery_tag=method.delivery_tag)
    
    def run(self):
        self.channel.basic_consume(queue=OCR_QUEUE, on_message_callback=self.process_task)
        self.channel.start_consuming()


def main():
    worker = OCRWorker()
    worker.run()


if __name__ == "__main__":
    main()
