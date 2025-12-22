# -*- coding: utf-8 -*-
"""
Flask应用入口文件：提供OCR识别和AI分类服务
"""

from flask import Flask, request, jsonify
import os
import tempfile
# from ocr_module import ocr_image
from pathlib import Path
from pdf_to_png import pdf_to_png, convert_image_to_jpg, png_to_jpg
from online_ocr_module import ocr_image
from ai_module import extract_invoice_fields_with_openai

# 创建Flask应用实例
app = Flask(__name__)

ALLOWED_MIME_TYPES = {
    "application/pdf",        # PDF文档
    "image/png",              # PNG图片
    "image/jpeg",             # JPEG图片
    "image/jpg",              # JPG图片
    "image/webp",             # WebP图片
    "image/bmp",              # BMP图片
    "image/tiff",             # TIFF图片
    "image/tif"               # TIF图片
}

@app.route('/process_image', methods=['POST'])
def process_image():
    """
    处理上传的账单图片，进行OCR识别和AI分类

    请求格式：
        POST /process_image
        Content-Type: multipart/form-data
        Form-Data: image (图片文件)

    返回格式：
        JSON: {
            "status": "success",
            "ocr_text": "识别出的文本",
            "category": "分类结果"
        }
    """
    temp_file_path = None  # 初始化临时文件路径变量
    temp_dir = None        # 初始化临时目录路径变量

    try:
        # 1. 是否上传文件
        if 'image' not in request.files:
            return jsonify({
                "status": "error",
                "message": "没有上传文件"
            }), 400

        file = request.files['image']

        # 2. 文件名校验
        if file.filename == '':
            return jsonify({
                "status": "error",
                "message": "文件名为空"
            }), 400

        # 3. 文件类型校验
        if file.mimetype not in ALLOWED_MIME_TYPES:
            return jsonify({
                "status": "error",
                "message": "仅支持上传 PDF 或图片文件"
            }), 400

        # 4. 判断是否 PDF
        is_pdf = (
                file.mimetype == "application/pdf"
                or file.filename.lower().endswith(".pdf")
        )

        # 5. 保存 & 转换
        if is_pdf:
            # 保存 PDF
            with tempfile.NamedTemporaryFile(delete=False, suffix=".pdf") as tmp_pdf:
                file.save(tmp_pdf.name)
                pdf_path = tmp_pdf.name

            try:
                # PDF -> PNG
                png_dir = Path(tempfile.mkdtemp())
                png_paths = pdf_to_png(
                    pdf_path=pdf_path,
                    output_dir=png_dir,
                    zoom=3.0
                )

                # PNG -> JPG
                jpg_dir = Path(tempfile.mkdtemp())
                jpg_paths = png_to_jpg(
                    png_paths=png_paths,
                    output_dir=jpg_dir
                )

                # 使用第一张转换后的图片进行OCR（如果有多张PDF页面）
                temp_file_path = jpg_paths[0]
                temp_dir = Path(tempfile.mkdtemp())
            finally:
                # 清理原始PDF文件
                if os.path.exists(pdf_path):
                    os.unlink(pdf_path)
        else:
            # 处理图片文件 - 统一转换为JPG格式
            # 首先保存上传的文件到临时位置
            with tempfile.NamedTemporaryFile(delete=False, suffix=Path(file.filename).suffix) as tmp_img:
                file.save(tmp_img.name)
                original_file_path = tmp_img.name

            try:
                # 将各种图片格式转换为JPG
                temp_dir = Path(tempfile.mkdtemp())
                temp_file_path = convert_image_to_jpg(
                    image_path=original_file_path,
                    output_dir=temp_dir
                )
            finally:
                # 清理原始上传的文件
                if os.path.exists(original_file_path):
                    os.unlink(original_file_path)

        # 执行OCR识别
        ocr_text = ocr_image(temp_file_path)

        # 如果OCR没有识别出文字，返回错误
        if not ocr_text.strip():
            return jsonify({
                "status": "error",
                "message": "未能识别出图片中的文字"
            }), 400

        # 使用AI模型进行分类
        category = extract_invoice_fields_with_openai(temp_file_path, ocr_text)

        # 返回成功结果
        return jsonify({
            "status": "success",
            "data": category
        })

    except Exception as e:
        # 捕获并返回异常信息
        return jsonify({
            "status": "error",
            "message": f"处理图片时发生错误: {str(e)}"
        }), 500
    finally:
        # 清理临时文件
        if temp_file_path and os.path.exists(temp_file_path):
            os.unlink(temp_file_path)
        if temp_dir and temp_dir.exists():
            import shutil
            shutil.rmtree(temp_dir, ignore_errors=True)

@app.route('/health', methods=['GET'])
def health_check():
    """
    健康检查接口
    """
    return jsonify({
        "status": "healthy",
        "message": "服务运行正常"
    })

if __name__ == '__main__':
    # 启动Flask应用
    # 在生产环境中，应该使用WSGI服务器如Gunicorn
    app.run(
        host='0.0.0.0',  # 监听所有接口
        port=5000,       # 端口号
        debug=True      # 关闭调试模式
    )