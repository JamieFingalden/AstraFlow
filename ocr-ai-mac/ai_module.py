# -*- coding: utf-8 -*-
"""
AI模型调用模块：支持多种AI服务进行发票字段提取
包括阿里云百炼API和OpenAI API，支持图像识别，可以从发票图片中提取税号、发票号、金额、商户名称、日期、支付方式等信息
"""

import json
import os
import re
import base64
from datetime import date
from dotenv import load_dotenv
from openai import OpenAI

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


def classify_expense(image_path):
    """
    从图片中提取发票字段信息

    Args:
        image_path (str): 图片文件路径

    Returns:
        dict: 包含发票字段的字典
    """
    # 首先进行OCR识别获取文本内容
    from ocr_module import ocr_image
    ocr_result = ocr_image(image_path)
    
    # 调用OpenAI API提取发票字段
    return extract_invoice_fields_with_openai(image_path, ocr_result)


def extract_invoice_fields_with_openai(image_path, ocr_result):
    """
    使用OpenAI API，通过图片提取发票字段

    Args:
        image_path (str): 图片文件路径
        ocr_result (str): OCR识别结果

    Returns:
        dict: 发票字段提取结果
    """
    # 将图片编码为base64
    base64_image = encode_image_to_base64(image_path)

    # 创建OpenAI客户端
    client = OpenAI(
        api_key=OPENAI_API_KEY,
        base_url=OPENAI_BASE_URL  # 可选，用于指定API基础URL
    )

    today_str = date.today().strftime("%Y-%m-%d")

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
    messages = [
        {"role": "system", "content": system_prompt},
        {"role": "user", "content": [
            {"type": "image_url", "image_url": {"url": f"data:image/jpeg;base64,{base64_image}"}},
            {"type": "text", "text": f"OCR识别结果: {ocr_result}\n请提取这张发票的详细信息"}
        ]}
    ]

    try:
        # 调用OpenAI API
        response = client.chat.completions.create(
            model="Qwen/Qwen3-VL-8B-Instruct",
            messages=messages,
            response_format={"type": "json_object"}  # 确保返回JSON格式
        )

        response_text = response.choices[0].message.content.strip()

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