# -*- coding: utf-8 -*-
"""
OCR 队列工作进程：从 RabbitMQ 队列中获取任务并处理 OCR 识别
"""

import os
import json
import pika
import requests
import traceback
from datetime import datetime
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


def log(message):
    print(f"[{datetime.now().strftime('%H:%M:%S')}] {message}", flush=True)


class OCRWorker:
    def __init__(self):
        log(f"连接 RabbitMQ {RABBITMQ_HOST}:{RABBITMQ_PORT}, 队列={OCR_QUEUE}")
        credentials = pika.PlainCredentials(RABBITMQ_USER, RABBITMQ_PASSWORD)
        self.connection = pika.BlockingConnection(
            pika.ConnectionParameters(host=RABBITMQ_HOST, port=RABBITMQ_PORT, credentials=credentials)
        )
        self.channel = self.connection.channel()
        self.channel.queue_declare(queue=OCR_QUEUE, durable=True)  # OCR任务队列
        self.channel.basic_qos(prefetch_count=1)  # 设置每次只处理一个消息
        log("RabbitMQ 已连接，等待任务")

    def process_task(self, ch, method, properties, body):
        callback_url = "/api/v1/callback/ocr-result"
        tmp_file_name = None
        temp_dir = None
        image_path = None
        task_data = json.loads(body)
        task_id = task_data.get('task_id')
        file_path = task_data.get('file_path')
        file_id = task_data.get('file_id', 0)  # 从任务数据中获取 file_id
        invoice_id = task_data.get('invoice_id', 0) # 获取 invoice_id
        log(f"收到任务 task_id={task_id}, file_id={file_id}, invoice_id={invoice_id}")

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
            log(f"下载文件成功: {file_url}")

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
            log(f"OCR完成，文本长度={len(ocr_text)}")

            # 如果 OCR 没有识别出文字，返回错误
            if not ocr_text.strip():
                raise ValueError("未能识别出图片中的文字")

            # 使用AI模块进行发票信息提取，传递OCR结果以避免重复识别
            log("开始模型处理")
            invoice_info = extract_invoice_info(image_path, use_ai=True, ocr_text=ocr_text)
            extracted_fields = invoice_info["extracted_fields"]

            raw_amount = extracted_fields.get("amount", 0)
            try:
                amount = float(raw_amount) if raw_amount not in (None, "") else 0.0
            except (TypeError, ValueError):
                amount = 0.0

            vendor = extracted_fields.get("vendor") or extracted_fields.get("merchant_name") or ""
            invoice_date = extracted_fields.get("invoice_date") or extracted_fields.get("date") or ""
            category = extracted_fields.get("category") or extracted_fields.get("invoice_type") or ""

            # 准备回调数据
            callback_data = {
                "task_id": task_id,
                "status": "success",
                "data": {
                    "file_id": file_id,
                    "invoice_id": invoice_id,
                    "invoice_number": extracted_fields.get("invoice_number", ""),
                    "vendor": vendor,
                    "description": extracted_fields.get("description", ""),
                    "amount": amount,
                    "category": category,
                    "payment_source": extracted_fields.get("payment_source", ""),
                    "tax_id": extracted_fields.get("tax_id", ""),
                    "invoice_date": invoice_date,
                    "ocr_text": ocr_text
                }
            }

            log("OCR处理结果:")
            log(json.dumps(callback_data, ensure_ascii=False))

            log("调用回调函数")
            # 发送回调请求到Go后端
            if callback_url:
                try:
                    callback_response = requests.post(
                        "http://localhost:8080" + callback_url,
                        json=callback_data,
                        headers={"Content-Type": "application/json"},
                        timeout=30
                    )
                    log(f"回调请求状态码: {callback_response.status_code}")
                    log(f"回调响应: {callback_response.text}")
                except Exception as callback_error:
                    log(f"发送回调请求失败: {str(callback_error)}")
                    # 即使回调失败，也继续处理

            # 清理临时文件
            log("清理临时文件")
            if image_path and os.path.exists(image_path):
                os.unlink(image_path)
            if temp_dir:
                shutil.rmtree(temp_dir)

        except Exception as e:
            log(f"处理任务 {task_id} 时发生错误: {str(e)}")
            log(traceback.format_exc())
            # 准备错误回调数据
            error_callback_data = {
                "task_id": task_id,
                "status": "failure",
                "data": {
                    "file_id": file_id,
                    "invoice_id": invoice_id
                },
                "error_message": str(e)
            }

            log("OCR处理错误:")
            log(json.dumps(error_callback_data, ensure_ascii=False))

            # 发送错误回调请求到Go后端
            if callback_url:
                try:
                    callback_response = requests.post(
                        "http://localhost:8080" + callback_url,
                        json=error_callback_data,
                        headers={"Content-Type": "application/json"},
                        timeout=30
                    )
                    log(f"错误回调请求状态码: {callback_response.status_code}")
                    log(f"错误回调响应: {callback_response.text}")
                except Exception as callback_error:
                    log(f"发送错误回调请求失败: {str(callback_error)}")

            # 清理临时文件
            try:
                if tmp_file_name and os.path.exists(tmp_file_name):
                    os.unlink(tmp_file_name)
                if temp_dir:
                    shutil.rmtree(temp_dir)
            except:
                pass

        finally:
            # 确认消息已处理
            ch.basic_ack(delivery_tag=method.delivery_tag)
    
    def run(self):
        log("开始消费队列")
        self.channel.basic_consume(queue=OCR_QUEUE, on_message_callback=self.process_task)
        self.channel.start_consuming()


def main():
    log("OCR Python Worker 启动")
    worker = OCRWorker()
    worker.run()


if __name__ == "__main__":
    main()
