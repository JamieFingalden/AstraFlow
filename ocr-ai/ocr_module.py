# -*- coding: utf-8 -*-
"""
OCR模块：使用PaddleOCR进行中文文本识别
"""

from paddleocr import PaddleOCR
import os

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

# 测试代码（可选）
if __name__ == "__main__":
    # 示例用法
    test_image = "test_bill.jpg"
    if os.path.exists(test_image):
        ocr_text = ocr_image(test_image)
        print("OCR识别结果:")
        print(ocr_text)
    else:
        print(f"测试图片 {test_image} 不存在")