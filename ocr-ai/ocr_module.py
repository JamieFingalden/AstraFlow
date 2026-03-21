# -*- coding: utf-8 -*-
"""
OCR模块：使用PaddleOCR进行中文文本识别
同时集成AI模块，支持从图片中提取结构化发票信息
"""

from paddleocr import PaddleOCR
import os
import ai_module

# 通过环境变量控制是否启用 GPU（Windows + NVIDIA 可设置为 1）
OCR_USE_GPU = os.getenv("OCR_USE_GPU", "0").lower() in ("1", "true", "yes", "on")

# CPU 模式下关闭 OneDNN/MKLDNN，避免部分环境报错
if not OCR_USE_GPU:
    os.environ["CUDA_VISIBLE_DEVICES"] = "-1"
    os.environ["FLAGS_allocator_strategy"] = "naive_best_fit"
    os.environ["FLAGS_use_mkldnn"] = "0"
    os.environ["PADDLE_DISABLE_ONEDNN"] = "1"

# Global variable to hold the OCR engine instance
_ocr_engine = None


def get_ocr_engine():
    """
    Lazy initialization of the OCR engine to avoid OneDNN issues during startup.
    Creates the OCR engine instance only when first needed.
    """
    global _ocr_engine
    if _ocr_engine is None:
        _ocr_engine = PaddleOCR(
            lang='ch',
            ocr_version='PP-OCRv3',
            use_textline_orientation=True,
            use_gpu=OCR_USE_GPU,
        )
    return _ocr_engine


def ocr_image(image_path):
    """
    对图片进行OCR识别，返回拼接后的文本

    Args:
        image_path (str): 图片文件路径

    Returns:
        str: 识别出的文本内容
    """
    if not os.path.exists(image_path):
        raise FileNotFoundError(f"图片文件不存在: {image_path}")

    ocr_engine = get_ocr_engine()
    result = ocr_engine.predict(image_path)

    texts = []
    if result and len(result) > 0:
        first_result = result[0]
        if isinstance(first_result, dict) and 'rec_texts' in first_result:
            recognized_texts = first_result['rec_texts']
            if isinstance(recognized_texts, list):
                for text in recognized_texts:
                    if text:
                        texts.append(str(text))

    return '\n'.join(texts)


def extract_invoice_info(image_path, use_ai=True, ocr_text=None):
    """
    完整的发票信息提取流程：先进行OCR识别，然后使用AI提取结构化字段

    Args:
        image_path (str): 图片文件路径
        use_ai (bool): 是否使用AI模型进行字段提取，如果为False则只返回OCR文本
        ocr_text (str, optional): 已经进行OCR识别的结果，如果未提供则自行识别

    Returns:
        dict: 包含OCR文本和提取的结构化字段
    """
    if ocr_text is None:
        ocr_text = ocr_image(image_path)

    result = {
        "ocr_text": ocr_text,
        "extracted_fields": {}
    }

    if use_ai:
        try:
            ai_result = ai_module.classify_expense(image_path=image_path)
            result["extracted_fields"] = ai_result
        except Exception as e:
            print(f"AI字段提取失败: {str(e)}")
            result["extracted_fields"] = {
                "tax_id": "",
                "invoice_number": "",
                "amount": 0.0,
                "merchant_name": "",
                "date": "",
                "payment_method": "",
                "invoice_type": ""
            }
    else:
        result["extracted_fields"] = {
            "tax_id": "",
            "invoice_number": "",
            "amount": 0.0,
            "merchant_name": "",
            "date": "",
            "payment_method": "",
            "invoice_type": ""
        }

    return result


if __name__ == "__main__":
    test_image = "test_bill.jpg"
    if os.path.exists(test_image):
        ocr_text = ocr_image(test_image)
        print("OCR识别结果:")
        print(ocr_text)

        print("\n" + "=" * 50)

        full_result = extract_invoice_info(test_image, use_ai=True)
        print("完整发票信息提取结果:")
        print(f"OCR文本:\n{full_result['ocr_text']}")
        print(f"\n提取的字段:\n{full_result['extracted_fields']}")
    else:
        print(f"测试图片 {test_image} 不存在")
