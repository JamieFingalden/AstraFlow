📘 AstraFlow API 文档（v1）

AstraFlow 是一个面向中小微企业与个人用户的财务管理系统，支持多租户模式与个人用户模式。

本接口文档覆盖以下模块：
	•	🧩 认证（Auth）
	•	🏢 租户（Tenant）
	•	👤 用户（User）
	•	🧾 发票（Invoice）
	•	🪪 OCR 识别结果（OCR Result）
	•	📎 附件（Attachment）
	•	💵 报销单（Reimbursement）

⸻

🌐 API 总览

1. Auth 认证接口
功能
方法
URL
用户注册
POST
/api/v1/auth/register
用户登录
POST
/api/v1/auth/login
刷新 Token
POST
/api/v1/auth/refresh
获取当前登录用户信息
GET
/api/v1/auth/me
