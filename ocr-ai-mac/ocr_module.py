# -*- coding: utf-8 -*-
"""
OCR 模块：使用 PaddleOCR 识别中文票据文本，并提取结构化发票字段。
"""

import os
import re

import paddle
from paddleocr import PaddleOCR

import ai_module

# 当前 mac 目录沿用 CPU 模式，避免 OneDNN/MKLDNN 在本地环境报错。
os.environ["CUDA_VISIBLE_DEVICES"] = "-1"
paddle.set_device("cpu")
os.environ["FLAGS_allocator_strategy"] = "naive_best_fit"
os.environ["FLAGS_use_mkldnn"] = "0"
os.environ["PADDLE_DISABLE_ONEDNN"] = "1"

_ocr_engine = None


def get_ocr_engine():
    """懒加载 OCR 引擎，避免服务启动时就占用模型资源。"""
    global _ocr_engine
    if _ocr_engine is None:
        _ocr_engine = PaddleOCR(
            lang="ch",
            ocr_version="PP-OCRv3",
            use_textline_orientation=True,
        )
    return _ocr_engine


def ocr_image(image_path):
    """识别图片中的文本并按行拼接返回。"""
    if not os.path.exists(image_path):
        raise FileNotFoundError(f"图片文件不存在: {image_path}")

    result = get_ocr_engine().predict(image_path)

    texts = []
    if result and len(result) > 0:
        first_result = result[0]
        if isinstance(first_result, dict) and "rec_texts" in first_result:
            recognized_texts = first_result["rec_texts"]
            if isinstance(recognized_texts, list):
                texts.extend(str(text) for text in recognized_texts if text)

    return "\n".join(texts)


def extract_invoice_info(image_path, use_ai=True, ocr_text=None):
    """
    返回 OCR 原文和结构化字段。
    AI 提取失败时使用 OCR 文本保底解析，避免把已识别出的关键信息丢掉。
    """
    if ocr_text is None:
        ocr_text = ocr_image(image_path)

    result = {
        "ocr_text": ocr_text,
        "extracted_fields": {},
    }

    if use_ai:
        try:
            result["extracted_fields"] = ai_module.classify_expense(
                image_path=image_path,
                ocr_result=ocr_text,
            )
        except Exception as e:
            print(f"AI字段提取失败: {str(e)}")
            result["extracted_fields"] = parse_invoice_fields_from_ocr_text(ocr_text)
    else:
        result["extracted_fields"] = parse_invoice_fields_from_ocr_text(ocr_text)

    return result


def parse_invoice_fields_from_ocr_text(ocr_text):
    """
    从 OCR 文本中保底解析结构化字段。
    仅提取明确标签附近的信息，避免把不确定内容写入后端。
    """
    text = ocr_text or ""
    lines = [line.strip() for line in text.splitlines() if line.strip()]

    def first_match(pattern, source=text):
        match = re.search(pattern, source, re.S)
        return match.group(1).strip() if match else ""

    invoice_number = first_match(r"发票号码[:：]?\s*([0-9A-Za-z]+)")

    invoice_date = ""
    date_match = re.search(r"开票日期[:：]?\s*(\d{4})年\s*(\d{1,2})月\s*(\d{1,2})日", text)
    if date_match:
        invoice_date = f"{date_match.group(1)}-{int(date_match.group(2)):02d}-{int(date_match.group(3)):02d}"

    amount_text = first_match(r"[（(]小写[）)]\s*[￥¥]?\s*([0-9]+(?:\.[0-9]+)?)")
    if not amount_text:
        amount_text = first_match(r"价税合计.*?[￥¥]\s*([0-9]+(?:\.[0-9]+)?)")
    try:
        amount = float(amount_text) if amount_text else 0.0
    except ValueError:
        amount = 0.0

    vendor = ""
    vendor_match = re.search(r"销售方信息.*?名称[:：]\s*([^\n]+)", text, re.S)
    if vendor_match:
        vendor = vendor_match.group(1).strip()

    tax_id = ""
    seller_block = text.split("销售方信息", 1)[1] if "销售方信息" in text else text
    tax_matches = re.findall(r"统一社会信用代码/纳税人识别号[:：]\s*([0-9A-Z]+)", seller_block)
    if tax_matches:
        tax_id = tax_matches[-1].strip()

    payment_method = first_match(r"支付渠道[:：]\s*([^；;\n]+)")

    description = ""
    for line in lines:
        if line.startswith("*"):
            description = line.strip("*")
            break

    category = ""
    category_source = f"{vendor} {description} {text}"
    if re.search(r"餐|饮|奶茶|咖啡|食品|饭|茶", category_source):
        category = "餐饮"
    elif re.search(r"车|票|高铁|铁路|航班|打车|交通", category_source):
        category = "交通"
    elif re.search(r"办公|文具|耗材", category_source):
        category = "办公用品"

    return {
        "tax_id": tax_id,
        "invoice_number": invoice_number,
        "amount": amount,
        "merchant_name": vendor,
        "vendor": vendor,
        "date": invoice_date,
        "invoice_date": invoice_date,
        "payment_method": payment_method,
        "payment_source": payment_method,
        "category": category,
        "invoice_type": category,
        "description": description,
    }


if __name__ == "__main__":
    test_image = "test_bill.jpg"
    if os.path.exists(test_image):
        recognized_text = ocr_image(test_image)
        print("OCR识别结果:")
        print(recognized_text)

        print("\n" + "=" * 50)

        full_result = extract_invoice_info(test_image, use_ai=True)
        print("完整发票信息提取结果:")
        print(f"OCR文本:\n{full_result['ocr_text']}")
        print(f"\n提取的字段:\n{full_result['extracted_fields']}")
    else:
        print(f"测试图片 {test_image} 不存在")
