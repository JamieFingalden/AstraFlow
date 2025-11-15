# 毕设

# **AstraFlow 智能报销与账单管理平台 — 项目文档（V1）**

## **📌1. 项目简介（Project Overview）**

**AstraFlow** 是一个面向中小微企业的 **AI 智能报销 + 账单管理平台**。

核心目标：减少财务人员的重复劳动，让企业的报销流程更自动化、更透明。

平台由两条主线组成：

- **AI 识别层（Flask / OCR / LLM）**
- **业务服务层（Go / Gin / MySQL）**
- **前端可视化层（Vue3 + Element Plus）**

主要能力包括：

- 自动识别发票/支付记录（微信/支付宝）
- 自动解析金额、日期、品类
- 自动生成账单并推送到财务
- 多租户企业管理（支持多个公司账户）
- 可视化账单分析与趋势图表
- 审批流（后续版本）

---

# **2. 系统架构（System Architecture）**

```
     +-----------------------------+
     |        Vue3 前端层          |
     |  发票上传 / 可视化 / 审批页  |
     +---------------+-------------+
                     |
                     v
     +------------------------------+
     |       Go 业务后端 (REST)     |
     | Gin + GORM + MySQL + JWT     |
     +-------+-----------+----------+
             |           |
 +-----------+           +------------------+
 |                                          |
 v                                          v
+-----------+                            +-------------------+
| MySQL DB  |                            |  Flask AI 服务    |
| 多租户表   |                            | OCR + LLM 推理    |
+-----------+                            +-------------------+
```

# **3. 技术栈（Tech Stack）**

### **前端（Frontend）**

- Vue 3（Composition API）
- Vite
- Pinia
- Element Plus
- Axios
- ECharts

### **后端（Backend - Go）**

- Go 1.22+
- Gin
- GORM
- MySQL 8.0
- Viper（配置）
- Zap（日志）
- JWT（认证）
- 分层架构：router → handler → service → repository → model

### **AI 服务**

- Python 3.10+
- Flask
- PaddleOCR / RapidOCR / 你选择的 OCR 套件
- Llama / Qwen（本地大模型）
- 文本清洗与结构化处理

---

# **4. 项目功能模块（Features）**

## **4.1 基础功能**

| **模块** | **描述** |
| --- | --- |
| 多租户管理 | 一个用户可属于多个公司；每个公司有独立账单数据 |
| 用户登录注册 | JWT 鉴权、密码加盐、Session 验证 |
| 公司管理 | 创建公司、加入公司、切换公司 |
| 用户角色 | 管理员 / 财务 / 普通员工 |

## **4.2 发票上传模块（Invoice Upload）**

- 支持上传多种文件格式：PNG/JPG/PDF
- 支持上传支付记录截图（微信/支付宝）
- 上传后自动调用 Flask AI 服务识别
- 返回识别的金额 / 日期 / 类目 / 商户 / 备注

---

## **4.3 AI 识别结果页**

- 展示 AI 解析结果
- 允许用户手动修改字段
- 可以保存为账单
- 提供“重新识别”功能

---

## **4.4 账单可视化（Analytics）**

- 支出趋势折线图
- 分类统计饼图
- 月度对比柱状图
- 最近 30 天列表
- 支持分类筛选、时间筛选、公司筛选

---

## **4.5 后续扩展模块（Future Work）**

- 自动审批流
- 企业预算管理
- 对接银行流水
- 自动生成报销凭证 PDF
- API-Gateway（后端多版本管理）

## **5. 数据库设计（ER）**

## **5.1 租户表（tenant）**

企业模式下用于区分不同公司；个人用户不涉及租户。

```sql
CREATE TABLE tenant (
id BIGINT PRIMARY KEY AUTO_INCREMENT,
name VARCHAR(100) NOT NULL COMMENT '租户名称（公司名称）',
industry VARCHAR(100) DEFAULT NULL COMMENT '行业',
contact_name VARCHAR(100) DEFAULT NULL COMMENT '联系人姓名',
contact_phone VARCHAR(50) DEFAULT NULL COMMENT '联系电话',
contact_email VARCHAR(100) DEFAULT NULL COMMENT '联系邮箱',
created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='租户表，多租户模式使用';
```

## **5.2 用户表（user）**

支持租户用户 / 个人用户。

```sql
CREATE TABLE user (
id BIGINT PRIMARY KEY AUTO_INCREMENT,
tenant_id BIGINT DEFAULT NULL COMMENT '租户ID，NULL 表示个人用户',
username VARCHAR(50) NOT NULL UNIQUE COMMENT '登录用户名',
password VARCHAR(255) NOT NULL COMMENT '加密后的密码',
email VARCHAR(100) DEFAULT NULL COMMENT '邮箱',
phone VARCHAR(50) DEFAULT NULL COMMENT '手机号',
role VARCHAR(30) DEFAULT 'normal' COMMENT '用户角色：admin/normal/personal',
created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
INDEX idx_user_tenant (tenant_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表，支持租户用户与个人用户';
```

## **5.3 发票表（invoice）**

发票 / 支付记录基础信息。

```sql
CREATE TABLE invoice (
id BIGINT PRIMARY KEY AUTO_INCREMENT,
tenant_id BIGINT DEFAULT NULL COMMENT '租户ID，NULL 表示个人用户',
user_id BIGINT NOT NULL COMMENT '上传用户ID',
invoice_number VARCHAR(100) DEFAULT NULL COMMENT '发票编号',
invoice_date DATE DEFAULT NULL COMMENT '开票日期',
amount DECIMAL(12,2) DEFAULT NULL COMMENT '金额',
vendor VARCHAR(255) DEFAULT NULL COMMENT '商户名称',
tax_id VARCHAR(100) DEFAULT NULL COMMENT '税号',
category VARCHAR(100) DEFAULT NULL COMMENT '分类：餐饮、交通、购物、住宿、办公用品等',
payment_source VARCHAR(50) DEFAULT NULL COMMENT '支付方式：微信/支付宝/现金/银行卡 等',
status VARCHAR(50) DEFAULT 'pending' COMMENT '状态：pending/recognized/confirmed/rejected',
created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
INDEX idx_invoice_tenant (tenant_id),
INDEX idx_invoice_user (user_id),
INDEX idx_invoice_vendor (vendor),
INDEX idx_invoice_date (invoice_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='发票主表';
```

## **5.4 OCR 识别结果表（ocr_result）**

结构化 + 非结构化 OCR 承载表。

```sql
CREATE TABLE ocr_result (
id BIGINT PRIMARY KEY AUTO_INCREMENT,
tenant_id BIGINT DEFAULT NULL COMMENT '租户ID / 个人用户 NULL',
user_id BIGINT NOT NULL COMMENT '关联用户ID',
invoice_id BIGINT NOT NULL COMMENT '关联发票ID（软连接）',
amount DECIMAL(12,2) DEFAULT NULL COMMENT '金额',
invoice_date DATE DEFAULT NULL COMMENT '开票日期',
vendor VARCHAR(255) DEFAULT NULL COMMENT '商户',
category VARCHAR(100) DEFAULT NULL COMMENT '类别',
invoice_number VARCHAR(100) DEFAULT NULL COMMENT '发票号',
confidence DECIMAL(5,2) DEFAULT NULL COMMENT '识别置信度，例如 0.97',
raw_text LONGTEXT COMMENT 'OCR 原始文本',
parsed_json JSON COMMENT 'OCR 解析后的 JSON',
created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
INDEX idx_ocr_tenant (tenant_id),
INDEX idx_ocr_user (user_id),
INDEX idx_ocr_invoice (invoice_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='OCR 识别结果表';
```

## **3.5 附件表（attachment）**

存储发票照片 / 支付截图。

```sql
CREATE TABLE attachment (
id BIGINT PRIMARY KEY AUTO_INCREMENT,
tenant_id BIGINT DEFAULT NULL COMMENT '租户ID（个人用户为 NULL）',
user_id BIGINT NOT NULL COMMENT '上传者 ID',
invoice_id BIGINT DEFAULT NULL COMMENT '关联发票 ID，可为空（软连接）',
file_url VARCHAR(500) NOT NULL COMMENT '文件存储路径',
file_type VARCHAR(50) DEFAULT NULL COMMENT '文件类型：jpg/png/pdf 等',
file_size BIGINT DEFAULT NULL COMMENT '大小（字节）',
created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
INDEX idx_attachment_tenant (tenant_id),
INDEX idx_attachment_user (user_id),
INDEX idx_attachment_invoice (invoice_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='附件表：发票照片/支付截图';
```

## **5.6 报销单主表（reimbursement）**

```sql
CREATE TABLE reimbursement (
id BIGINT PRIMARY KEY AUTO_INCREMENT,
tenant_id BIGINT DEFAULT NULL COMMENT '租户ID（个人用户为 NULL）',
user_id BIGINT NOT NULL COMMENT '申请人',
title VARCHAR(255) NOT NULL COMMENT '报销单标题',
total_amount DECIMAL(12,2) DEFAULT NULL COMMENT '总金额',
status VARCHAR(50) DEFAULT 'draft' COMMENT '状态：draft/submitted/approved/rejected/paid',
created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
INDEX idx_reim_tenant (tenant_id),
INDEX idx_reim_user (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='报销单主表';
```

## **5.7 报销单子项（reimbursement_item）**

关联多张发票。

```sql
CREATE TABLE reimbursement_item (
id BIGINT PRIMARY KEY AUTO_INCREMENT,
reimbursement_id BIGINT NOT NULL COMMENT '报销单 ID（软连接）',
invoice_id BIGINT NOT NULL COMMENT '发票 ID（软连接）',
amount DECIMAL(12,2) DEFAULT NULL COMMENT '发票金额快照',
note VARCHAR(255) DEFAULT NULL COMMENT '备注',
created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
INDEX idx_item_reim (reimbursement_id),
INDEX idx_item_invoice (invoice_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='报销单发票子项';
```

## **6. 项目愿景（Vision）**

AstraFlow 的定位不是“做个简单工具”，而是一个：

> “AI 驱动的中小企业财务自动化平台”
> 

未来可以拓展为商业化 SaaS。