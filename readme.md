# 毕设

# **AstraFlow 智能报销与账单管理平台 — 项目文档（V1）**

## **📌1. 项目简介（Project Overview）**

**AstraFlow** 是一个面向中小微企业的 **AI 智能报销 + 账单管理平台**。

核心目标：减少财务人员的重复劳动，让企业的报销流程更自动化、更透明。

平台采用**异步微服务架构**，由以下核心组件组成：

- **前端可视化层 (`astraflow`)**：基于 Vue 3 + Vite 的现代化 SPA。
- **业务服务层 (`AstraFlow-go`)**：基于 Go (Gin) 的高性能 RESTful API，负责业务逻辑、数据持久化及任务分发。
- **AI 任务工作流 (`ocr-ai-mac`)**：基于 Python + RabbitMQ 的异步消费者，集成 PaddleOCR 与 LLM，负责繁重的图像识别与数据提取任务。

主要能力包括：

- 自动识别发票/支付记录（微信/支付宝/PDF/图片）
- 异步处理 OCR 任务，通过 WebSocket 或轮询反馈结果
- 自动解析金额、日期、品类、商户等关键信息
- 多租户企业管理（支持多个公司账户隔离）
- 可视化账单分析与趋势图表 (ECharts)
- 审批流（规划中）

---

# **2. 系统架构（System Architecture）**

系统采用异步消息队列解耦业务逻辑与 AI 计算密集型任务。

```mermaid
graph TD
    User[用户 (Web/Mobile)] -->|上传发票| API[Go 业务后端 (Gin)]
    API -->|1. 存储文件| LocalStorage/S3[文件存储]
    API -->|2. 发布识别任务| MQ[RabbitMQ (ocr_tasks)]
    
    subgraph AI Worker [Python AI Worker]
        Worker[ocr-ai-mac] -->|3. 消费任务| MQ
        Worker -->|4. 下载文件| API
        Worker -->|5. OCR + LLM 推理| Model[PaddleOCR + LLM]
    end
    
    Worker -->|6. 回调结果 (Webhook)| API
    API -->|7. 存储结构化数据| DB[(MySQL)]
    API -->|8. 推送/展示结果| User
```

---

# **3. 技术栈（Tech Stack）**

### **前端（Frontend - `astraflow`）**

- **核心框架**: Vue 3 (Composition API) + Vite
- **状态管理**: Pinia
- **UI 组件库**: Element Plus, Naive UI
- **可视化**: ECharts, Chart.js
- **网络请求**: Axios
- **图标**: Lucide Vue Next

### **后端（Backend - `AstraFlow-go`）**

- **语言**: Go 1.22+
- **Web 框架**: Gin
- **数据库 ORM**: GORM (MySQL 8.0)
- **消息队列**: RabbitMQ (amqp)
- **配置管理**: Viper
- **日志**: Zap
- **认证**: JWT (JSON Web Tokens)
- **架构**: 清洁架构 (Router -> Handler -> Service -> Repository -> Model)

### **AI 服务（AI Worker - `ocr-ai-mac`）**

- **语言**: Python 3.10+
- **核心模式**: RabbitMQ Consumer (Pika)
- **OCR 引擎**: PaddleOCR
- **PDF 处理**: PyMuPDF
- **大模型集成**: OpenAI SDK (适配各类 LLM API)
- **任务流**: 异步消费 -> 图像预处理 -> OCR 文本提取 -> LLM 实体抽取 -> 结构化回调

---

# **4. 目录结构说明 (Directory Structure)**

| 目录 | 说明 | 备注 |
| :--- | :--- | :--- |
| `astraflow/` | 前端源代码 | Vue 3 + Vite 项目 |
| `AstraFlow-go/` | 后端源代码 | Go 模块，核心业务 API |
| `ocr-ai-mac/` | **核心 AI Worker** | 生产环境使用的 Python RabbitMQ 消费者 |
| `ocr-ai/` | 旧版 AI 服务 | (已废弃) 基于 Flask 的同步 HTTP 服务 |
| `document/` | 项目文档 | 毕设相关文档资料 |

---

# **5. 快速开始（Quick Start）**

## **5.1 环境准备**
- **基础设施**: MySQL 8.0+, RabbitMQ 3.x+
- **开发环境**: Go 1.22+, Node.js 18+, Python 3.10+

## **5.2 启动步骤**

### **1. 启动后端 (AstraFlow-go)**
确保 `pkg/config/config.yml` 中数据库和 RabbitMQ 配置正确。

```bash
cd AstraFlow-go
# 运行数据库迁移
go run cmd/migrate/main.go
# 启动 API 服务
go run cmd/server/main.go
```

### **2. 启动 AI Worker (ocr-ai-mac)**
确保 `.env` 中 RabbitMQ 配置正确。

```bash
cd ocr-ai-mac
# 安装依赖
pip install -r requirements.txt
# 启动消费者
python worker.py
```

### **3. 启动前端 (astraflow)**

```bash
cd astraflow
npm install
npm run dev
```

---

# **6. 项目功能模块（Features）**

## **6.1 基础架构**

| 模块 | 描述 |
| --- | --- |
| **多租户体系** | 支持个人用户与企业租户并存；数据逻辑隔离 |
| **安全认证** | 基于 JWT 的双 Token 机制 (Access/Refresh Token) |
| **异步处理** | 大文件上传与 AI 识别全异步化，提升系统吞吐量 |

## **6.2 核心业务**

- **智能票据采集**:
    - 支持 PDF 发票转图片处理
    - 支持微信/支付宝账单截图识别
- **AI 结构化引擎**:
    - 结合 OCR 文字坐标与语义理解
    - 自动提取：`金额`、`日期`、`发票号`、`销售方`、`税号`、`消费类型`
- **财务数据中心**:
    - 个人/企业报销单管理
    - 支出可视化分析 (按月/按类目)
    - 原始凭证云存储

---

# **7. 数据库设计概要**

主要包含以下实体表：
- `tenant`: 租户/公司信息
- `user`: 用户信息（关联 tenant）
- `invoice`: 发票主表（存储基础元数据）
- `ocr_result`: AI 识别的详细结构化数据（JSON 格式存储复杂结果）
- `attachment`: 原始文件记录
- `reimbursement`: 报销单据
- `reimbursement_item`: 报销单明细

---

# **8. 项目愿景（Vision）**

AstraFlow 致力于成为 **“AI 驱动的中小企业财务自动化基础设施”**，通过技术手段极大地降低企业财务合规成本与人力成本。
