# -*- coding: utf-8 -*-
"""
OCR模块：使用PaddleOCR进行中文文本识别
同时集成AI模块，支持从图片中提取结构化发票信息
"""

from paddleocr import PaddleOCR
import os
import ai_module

# 设置环境变量 to avoid OneDNN issues
os.environ["CUDA_VISIBLE_DEVICES"] = "-1"

# 禁用 OneDNN / MKLDNN（避免 OneDNN Filter 报错）
os.environ['FLAGS_use_mkldnn'] = '0'
os.environ['PADDLE_DISABLE_ONEDNN'] = '1'

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
            use_textline_orientation=True,  # use_angle_cls is deprecated in 3.x, use use_textline_orientation instead
            lang='ch'
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
    # 检查文件是否存在
    if not os.path.exists(image_path):
        raise FileNotFoundError(f"图片文件不存在: {image_path}")

    # Get the OCR engine (lazy initialization)
    ocr_engine = get_ocr_engine()

    # 执行OCR识别
    result = ocr_engine.predict(image_path)

    # 提取文本内容并拼接
    # PaddleOCR 3.x returns a list of dictionaries, each containing recognition results
    texts = []
    if result and len(result) > 0:
        first_result = result[0]  # Get the first result (for single image)
        if isinstance(first_result, dict) and 'rec_texts' in first_result:
            # Extract recognized texts from the rec_texts field
            recognized_texts = first_result['rec_texts']
            if isinstance(recognized_texts, list):
                for text in recognized_texts:
                    if text:  # Check if text is not empty
                        texts.append(str(text))

    # 返回拼接后的文本
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
    # 如果OCR文本未提供，则执行OCR识别
    if ocr_text is None:
        ocr_text = ocr_image(image_path)

    result = {
        "ocr_text": ocr_text,
        "extracted_fields": {}
    }

    # 如果启用AI，则使用AI模型提取结构化字段
    if use_ai:
        try:
            # 使用AI模块提取发票字段
            ai_result = ai_module.classify_expense(image_path=image_path)
            result["extracted_fields"] = ai_result
        except Exception as e:
            print(f"AI字段提取失败: {str(e)}")
            # 如果AI提取失败，返回空的字段字典
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
        # 如果不使用AI，只返回OCR文本
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


# 测试代码（可选）
if __name__ == "__main__":
    # 示例用法
    test_image = "test_bill.jpg"
    if os.path.exists(test_image):
        # 只进行OCR识别
        ocr_text = ocr_image(test_image)
        print("OCR识别结果:")
        print(ocr_text)

        print("\n" + "="*50)

        # 完整的发票信息提取（OCR + AI字段提取）
        full_result = extract_invoice_info(test_image, use_ai=True)
        print("完整发票信息提取结果:")
        print(f"OCR文本:\n{full_result['ocr_text']}")
        print(f"\n提取的字段:\n{full_result['extracted_fields']}")
    else:
        print(f"测试图片 {test_image} 不存在")