# -*- coding: utf-8 -*-
"""
Flask应用入口文件：提供OCR识别和AI分类服务
"""

from flask import Flask, request, jsonify
import os
import tempfile
from ocr_module import ocr_image
from ai_module import extract_invoice_fields_with_openai

# 创建Flask应用实例
app = Flask(__name__)

ALLOWED_MIME_TYPES = {
    "application/pdf",
    "image/png",
    "image/jpeg",
    "image/jpg",
    "image/webp"
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

        # 4. 根据文件类型决定后缀
        if file.mimetype == "application/pdf" or file.filename.lower().endswith(".pdf"):
            suffix = ".pdf"
            file_type = "pdf"
        else:
            suffix = ".jpg"   # 统一图片后缀即可
            file_type = "image"

        # 5. 保存为临时文件
        with tempfile.NamedTemporaryFile(delete=False, suffix=suffix) as tmp_file:
            file.save(tmp_file.name)
            temp_file_path = tmp_file.name
        
        try:
            # 执行OCR识别
            ocr_text = ocr_image(temp_image_path)
            
            # 如果OCR没有识别出文字，返回错误
            if not ocr_text.strip():
                return jsonify({
                    "status": "error",
                    "message": "未能识别出图片中的文字"
                }), 400
            
            # 使用AI模型进行分类
            category = extract_invoice_fields_with_openai(temp_image_path, ocr_text)
            
            # 返回成功结果
            return jsonify({
                "status": "success",
                "data": category
            })
        
        finally:
            # 清理临时文件
            if os.path.exists(temp_image_path):
                os.unlink(temp_image_path)
    
    except Exception as e:
        # 捕获并返回异常信息
        return jsonify({
            "status": "error",
            "message": f"处理图片时发生错误: {str(e)}"
        }), 500

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