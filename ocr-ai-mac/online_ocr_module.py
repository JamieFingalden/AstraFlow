# -*- coding: utf-8 -*-
"""
OCR模块（线上版）：使用免费 OCR API 进行中文文本识别
适合作为 PaddleOCR 的轻量替代方案
"""

import os
import requests
import time
import ai_module
from requests.adapters import HTTPAdapter
from urllib3.util.retry import Retry

# OCR.space 免费 API（无需注册）
OCR_API_URL = "https://api.ocr.space/parse/image"
OCR_API_KEY = "helloworld"  # 官方提供的免费 key


def ocr_image(image_path):
    """
    使用线上 OCR API 对图片进行识别，返回拼接后的文本

    Args:
        image_path (str): 图片文件路径

    Returns:
        str: OCR 识别后的文本
    """
    if not os.path.exists(image_path):
        raise FileNotFoundError(f"图片文件不存在: {image_path}")

    # Create a session with retry strategy
    session = requests.Session()
    
    # Define retry strategy
    retry_strategy = Retry(
        total=3,  # Total number of retries
        backoff_factor=2,  # Exponential backoff factor (1s, 2s, 4s)
        status_forcelist=[429, 500, 502, 503, 504],  # HTTP status codes to retry
        allowed_methods=["POST"]  # HTTP methods to retry
    )
    
    # Mount adapter with retry strategy
    adapter = HTTPAdapter(max_retries=retry_strategy)
    session.mount("http://", adapter)
    session.mount("https://", adapter)

    # Attempt the request with retries
    last_exception = None
    for attempt in range(3):
        try:
            with open(image_path, "rb") as f:
                response = session.post(
                    OCR_API_URL,
                    files={"file": f},
                    data={
                        "apikey": OCR_API_KEY,
                        "language": "chs",           # 中文简体
                        "isOverlayRequired": False,
                        "OCREngine": 2               # OCR 引擎 2，效果更好
                    },
                    timeout=120,
                    verify=True  # Enable SSL verification
                )

            response.raise_for_status()
            result = response.json()

            if result.get("IsErroredOnProcessing"):
                error_msg = result.get("ErrorMessage", "未知 OCR 错误")
                raise RuntimeError(f"OCR 识别失败: {error_msg}")

            texts = []
            for item in result.get("ParsedResults", []):
                text = item.get("ParsedText", "")
                if text.strip():
                    texts.append(text.strip())

            return "\n".join(texts)
            
        except requests.exceptions.SSLError as e:
            last_exception = e
            if attempt < 2:  # If not the last attempt
                print(f"OCR API SSL error (attempt {attempt + 1}/3): {str(e)}")
                time.sleep(2 ** attempt)  # Exponential backoff
            else:
                raise RuntimeError(f"OCR API SSL error after 3 attempts: {str(e)}")
                
        except requests.exceptions.ConnectionError as e:
            last_exception = e
            if attempt < 2:  # If not the last attempt
                print(f"OCR API connection error (attempt {attempt + 1}/3): {str(e)}")
                time.sleep(2 ** attempt)  # Exponential backoff
            else:
                raise RuntimeError(f"OCR API connection error after 3 attempts: {str(e)}")
                
        except requests.exceptions.RequestException as e:
            last_exception = e
            if attempt < 2:  # If not the last attempt
                print(f"OCR API request error (attempt {attempt + 1}/3): {str(e)}")
                time.sleep(2 ** attempt)  # Exponential backoff
            else:
                raise RuntimeError(f"OCR API request error after 3 attempts: {str(e)}")
    
    # If we get here, all attempts failed
    raise RuntimeError(f"OCR API failed after 3 attempts: {str(last_exception)}")


def extract_invoice_info(image_path, use_ai=True):
    """
    发票信息提取流程：
    1. 调用线上 OCR API
    2. 可选：调用 AI 模块提取结构化字段
    """
    ocr_text = ocr_image(image_path)

    result = {
        "ocr_text": ocr_text,
        "extracted_fields": {}
    }

    print(result)

    if use_ai:
        try:
            from ai_module import extract_invoice_fields_with_openai
            ai_result = extract_invoice_fields_with_openai(
                image_path=image_path,
                ocr_text=ocr_text  # ✅ 把 OCR 文本也传给 AI
            )
            result["extracted_fields"] = ai_result
        except Exception as e:
            print(f"AI字段提取失败: {str(e)}")
            result["extracted_fields"] = _empty_invoice_fields()
    else:
        result["extracted_fields"] = _empty_invoice_fields()

    return result


def _empty_invoice_fields():
    """统一的空字段结构"""
    return {
        "tax_id": "",
        "invoice_number": "",
        "amount": 0.0,
        "merchant_name": "",
        "date": "",
        "payment_method": "",
        "invoice_type": ""
    }


# 测试代码
if __name__ == "__main__":
    test_image = "test_bill.jpg"
    if os.path.exists(test_image):
        print("OCR 识别结果：")
        print(ocr_image(test_image))

        print("\n" + "=" * 50)

        result = extract_invoice_info(test_image, use_ai=True)
        print("完整结果：")
        print(result)
    else:
        print("测试图片不存在")
