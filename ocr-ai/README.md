# OCR + AI发票信息提取系统

## 项目概述

该系统结合了PaddleOCR的文本识别能力和AI模型的结构化信息提取能力，能够从发票图片中提取以下信息：

- 税号 (Tax ID)
- 发票号码 (Invoice Number)
- 金额 (Amount)
- 商户名称 (Merchant Name)
- 开票日期 (Date)
- 支付方式 (Payment Method)

## 功能特点

### 1. OCR文本识别
- 使用PaddleOCR 3.3.2进行中文文本识别
- 支持多种格式的图片输入
- 高精度的文本检测和识别

### 2. AI驱动的结构化提取
- 支持图像直接输入到AI模型进行端到端提取
- 提供两种AI服务支持：
  - 本地Ollama服务（推荐）
  - 阿里云百炼API（备用方案）

### 3. 完整的发票处理流程
- OCR识别 → AI字段提取 → 结构化输出
- 自动错误处理和容错机制
- 支持仅OCR模式和完整模式

## 模块说明

### ocr_module.py
- `ocr_image(image_path)`: 执行OCR文本识别
- `extract_invoice_info(image_path, use_ai=True)`: 完整的发票信息提取流程
- 返回包含OCR文本和提取字段的字典

### ai_module.py
- `extract_invoice_fields(image_path)`: 从图像直接提取发票字段
- `classify_expense(ocr_text=None, image_path=None, user_id=None)`: 支持文本或图像输入的分类函数

## 使用方法

### 基本使用
```python
from ocr_module import extract_invoice_info

# 完整提取（OCR + AI）
result = extract_invoice_info("invoice.jpg", use_ai=True)

# 仅OCR识别
result = extract_invoice_info("invoice.jpg", use_ai=False)

# 输出格式
{
    "ocr_text": "识别的文本内容...",
    "extracted_fields": {
        "tax_id": "税号",
        "invoice_number": "发票号码",
        "amount": 金额数字,
        "merchant_name": "商户名称",
        "date": "YYYY-MM-DD",
        "payment_method": "支付方式"
    }
}
```

### 直接使用AI模块
```python
from ai_module import classify_expense

# 从图像直接提取
fields = classify_expense(image_path="invoice.jpg")

# 从OCR文本分类
classification = classify_expense(ocr_text="OCR识别的文本", user_id=123)
```

## 配置要求

### 本地Ollama服务（推荐）
1. 安装Ollama
2. 拉取支持视觉的模型：
   ```bash
   ollama pull llava
   ```
3. 启动Ollama服务：
   ```bash
   ollama serve
   ```

### 阿里云百炼API（备用）
在`.env`文件中配置：
```
DASHSCOPE_API_KEY=your_api_key_here
DASHSCOPE_URL=https://dashscope.aliyuncs.com/api/v1
```

## 依赖项

- paddleocr==3.3.2
- paddlepaddle
- requests
- python-dotenv
- base64（内置）

## 错误处理

- 如果AI服务不可用，系统会自动返回空字段值
- OCR识别独立于AI服务，即使AI服务失败也能获取文本内容
- 提供详细的错误日志用于调试

## 输出示例

```json
{
  "ocr_text": "WESVIP\n网易云音乐主站\nRenewal Voucher\nRenewal Amount\n￥5.00\n...",
  "extracted_fields": {
    "tax_id": "123456789012345678",
    "invoice_number": "NO12345678",
    "amount": 5.00,
    "merchant_name": "网易云音乐",
    "date": "2024-01-15",
    "payment_method": "支付宝"
  }
}
```

## 集成示例

```bash
python integration_example.py
```

这将处理`test_bill.jpg`并输出完整的发票信息提取结果到`invoice_extraction_result.json`。