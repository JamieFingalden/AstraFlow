# -*- coding: utf-8 -*-
"""
AI模型调用模块：支持多种AI服务进行发票字段提取
包括阿里云百炼API和OpenAI API，支持图像识别，可以从发票图片中提取税号、发票号、金额、商户名称、日期、支付方式等信息
"""

import json
import os
import re
import base64
from io import BytesIO
from typing import List
from dotenv import load_dotenv
from openai import OpenAI
from PIL import Image

# 加载环境变量
load_dotenv()

# 默认分类标签
DEFAULT_CATEGORIES = ["餐饮", "交通", "购物", "住宿", "办公用品", "其他"]

# Ollama服务地址
OLLAMA_URL = "http://localhost:11434/api/generate"

# 阿里云百炼API配置
DASHSCOPE_API_KEY = os.getenv('DASHSCOPE_API_KEY', 'your_api_key_here')
DASHSCOPE_URL = os.getenv('DASHSCOPE_URL', 'https://api-inference.modelscope.cn/v1/')

# OpenAI API配置
OPENAI_API_KEY = os.getenv('DASHSCOPE_API_KEY', 'your_openai_api_key_here')
OPENAI_BASE_URL = os.getenv('DASHSCOPE_URL', 'https://api.openai.com/v1')

MAX_IMAGE_SIDE = int(os.getenv("OPENAI_IMAGE_MAX_SIDE", "2048"))
LONG_IMAGE_RATIO = float(os.getenv("OPENAI_LONG_IMAGE_RATIO", "2.2"))
LONG_IMAGE_TILE_HEIGHT = int(os.getenv("OPENAI_LONG_IMAGE_TILE_HEIGHT", "1800"))
LONG_IMAGE_OVERLAP = int(os.getenv("OPENAI_LONG_IMAGE_OVERLAP", "120"))
MAX_IMAGE_SEGMENTS = int(os.getenv("OPENAI_MAX_IMAGE_SEGMENTS", "4"))

def encode_image_to_base64(image_path):
    """
    将图片文件编码为base64字符串

    Args:
        image_path (str): 图片文件路径

    Returns:
        str: base64编码的图片字符串
    """
    with open(image_path, "rb") as image_file:
        return base64.b64encode(image_file.read()).decode('utf-8')


def _resize_to_limit(image: Image.Image, limit: int) -> Image.Image:
    width, height = image.size
    max_side = max(width, height)
    if max_side <= limit:
        return image

    scale = limit / float(max_side)
    new_size = (max(1, int(width * scale)), max(1, int(height * scale)))
    return image.resize(new_size, Image.Resampling.LANCZOS)


def _split_long_image(image: Image.Image) -> List[Image.Image]:
    width, height = image.size
    if width <= 0 or height <= 0:
        return [image]

    ratio = height / float(width)
    if ratio < LONG_IMAGE_RATIO or height <= MAX_IMAGE_SIDE:
        return [image]

    tile_height = max(600, min(LONG_IMAGE_TILE_HEIGHT, MAX_IMAGE_SIDE))
    overlap = max(0, min(LONG_IMAGE_OVERLAP, tile_height // 3))
    step = tile_height - overlap
    if step <= 0:
        step = tile_height

    segments: List[Image.Image] = []
    top = 0
    while top < height and len(segments) < MAX_IMAGE_SEGMENTS:
        bottom = min(height, top + tile_height)
        segments.append(image.crop((0, top, width, bottom)))
        if bottom >= height:
            break
        top += step

    return segments if segments else [image]


def _encode_image_obj_to_base64(image: Image.Image) -> str:
    buffer = BytesIO()
    image.save(buffer, format="JPEG", quality=90, optimize=True)
    return base64.b64encode(buffer.getvalue()).decode("utf-8")


def prepare_image_payloads(image_path: str) -> List[str]:
    with Image.open(image_path) as image:
        normalized = image.convert("RGB")
        normalized = _resize_to_limit(normalized, MAX_IMAGE_SIDE)
        segments = _split_long_image(normalized)
        return [_encode_image_obj_to_base64(segment) for segment in segments]


def _extract_response_text(response) -> str:
    content = response.choices[0].message.content
    if isinstance(content, str):
        return content.strip()
    return ""


def extract_invoice_fields_with_openai(image_path, ocr_result):
    """
    使用OpenAI API，通过图片提取发票字段

    Args:
        image_path (str): 图片文件路径
        ocr_result (str): OCR识别结果

    Returns:
        dict: 发票字段提取结果
    """
    # 将图片预处理后编码为base64（自动缩放+长图切片）
    base64_images = prepare_image_payloads(image_path)

    # 创建OpenAI客户端
    client = OpenAI(
        api_key=OPENAI_API_KEY,
        base_url=OPENAI_BASE_URL  # 可选，用于指定API基础URL
    )

    # 构造系统提示词
    system_prompt = f"""你是一个发票信息整理分类助手。
        你将接收一张发票图片和OCR识别结果，任务是：
        1. 分辨税号（Tax ID / 统一社会信用代码）
        2. 分辨发票号码（Invoice Number）
        3. 分辨总金额（Total Amount）
        4. 分辨商户名称（Merchant Name / Seller Name）
        5. 分辨开票日期（Invoice Date）
        6. 分辨支付方式（Payment Method）
        7. 分辨发票类型（Invoice Type）
        8. 如果某个字段在图片中不存在或识别不出，请留空

        商户名称需要根据情况来分辨，如以下例子
        {{
            {{
                "input": "网易云音乐主站",
                "output": "网易云音乐"
            }},
            {{
                "input": "中国平安保险股份有限公司",
                "output": "中国平安保险"
            }}
        }}

        
        你需要根据发票信息来判断该发票数据下面几个类别中的哪一类并放入invoice_type：{DEFAULT_CATEGORIES}
        
        发票OCR识别结果：{ocr_result}

        返回格式严格按照以下JSON结构：
        {{
            "tax_id": null, // 为空时返回null
            "invoice_number": "发票号码",
            "amount": "金额数字",
            "merchant_name": "商户名称",
            "date": "YYYY-MM-DD格式日期",
            "payment_method": "支付方式",
            "invoice_type": "发票类型"
        }}

        请严格按照上述JSON格式返回，不要包含其他文字。"""

    # 构造请求数据
    user_content = [
        {"type": "image_url", "image_url": {"url": f"data:image/jpeg;base64,{image_base64}"}}
        for image_base64 in base64_images
    ]
    user_content.append({"type": "text", "text": f"OCR识别结果: {ocr_result}\n请提取这张发票的详细信息"})

    messages = [
        {"role": "system", "content": system_prompt},
        {"role": "user", "content": user_content}
    ]

    try:
        # 调用OpenAI API
        response = client.chat.completions.create(
            model="Qwen/Qwen3-VL-8B-Instruct",
            messages=messages,
            response_format={"type": "json_object"}  # 确保返回JSON格式
        )

        response_text = _extract_response_text(response)
        if not response_text:
            raise ValueError("模型返回内容为空")

        # 解析JSON响应
        json_match = re.search(r'\{.*\}', response_text, re.DOTALL)
        if json_match:
            result_json = json_match.group()
            result = json.loads(result_json)
            return result
        else:
            raise ValueError("无法从响应中提取JSON")

    except Exception as e:
        raise RuntimeError(f"OpenAI API调用失败: {str(e)}")


# 测试代码（可选）
if __name__ == "__main__":
    # 初始化模型（在实际应用中，这应该在服务启动时调用一次）
    # initialize_model()

    # 示例1：基于文本分类
    print("=== 基于文本的分类 ===")
    test_text = "星巴克咖啡 价格: 35元"
    # Note: This function is for invoice field extraction, not text classification
    print(f"文本内容: {test_text}")

    # 示例2：基于图像提取发票字段
    print("\n=== 基于图像的发票字段提取 (DashScope) ===")
    test_image = "test_bill.jpg"  # 请确保测试图片存在

    # 示例3：基于图像提取发票字段 using OpenAI
    print("\n=== 基于图像的发票字段提取 (OpenAI) ===")
    if os.path.exists(test_image):
        result = extract_invoice_fields_with_openai(image_path=test_image, ocr_result="OCR result placeholder")
        print(f"OpenAI图像提取结果: {result}")
    else:
        print(f"测试图片 {test_image} 不存在")
