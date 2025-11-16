# API

# **📘 AstraFlow API 文档（v1）**

AstraFlow 是一个面向中小微企业与个人用户的财务管理系统，支持多租户、个人用户、OCR 发票识别、报销流转等核心功能。

本接口文档基于以下数据库结构设计：

- 租户（tenant）
- 用户（user）
- 发票（invoice）
- OCR 结果（ocr_result）
- 附件（attachment）
- 报销单（reimbursement）
- 报销单子项（reimbursement_item）

---

# **🌐 API 总览**

---

## **1. Auth 认证接口**

| **功能** | **方法** | **URL** |
| --- | --- | --- |
| 用户注册 | POST | /api/v1/auth/register |
| 用户登录 | POST | /api/v1/auth/login |
| 刷新 Token | POST | /api/v1/auth/refresh |
| 获取当前用户信息 | GET | /api/v1/auth/me |

## **2. Tenant 租户接口**

| **功能** | **方法** | **URL** |
| --- | --- | --- |
| 创建租户 | POST | /api/v1/tenants |
| 获取租户列表 | GET | /api/v1/tenants |
| 获取租户详情 | GET | /api/v1/tenants/{id} |

## **3. Invoice 发票接口**

| **功能** | **方法** | **URL** |
| --- | --- | --- |
| 创建发票 | POST | /api/v1/invoices |
| 获取发票列表 | GET | /api/v1/invoices |
| 获取发票详情 | GET | /api/v1/invoices/{id} |
| 更新发票 | PUT | /api/v1/invoices/{id} |
| 删除发票 | DELETE | /api/v1/invoices/{id} |

## **4. OCR Result 识别结果接口**

| **功能** | **方法** | **URL** |
| --- | --- | --- |
| 创建 OCR 记录 | POST | /api/v1/ocr |
| 根据发票 ID 查询 | GET | /api/v1/ocr/invoice/{invoice_id} |
| OCR 详情 | GET | /api/v1/ocr/{id} |

## **5. Attachment 附件接口**

| **功能** | **方法** | **URL** |
| --- | --- | --- |
| 上传附件 | POST | /api/v1/attachments |
| 附件列表 | GET | /api/v1/attachments |
| 附件详情 | GET | /api/v1/attachments/{id} |
| 删除附件 | DELETE | /api/v1/attachments/{id} |

## **6. Reimbursement 报销单接口**

| **功能** | **方法** | **URL** |
| --- | --- | --- |
| 创建报销单 | POST | /api/v1/reimbursements |
| 报销单列表 | GET | /api/v1/reimbursements |
| 报销单详情 | GET | /api/v1/reimbursements/{id} |
| 更新报销单 | PUT | /api/v1/reimbursements/{id} |
| 删除报销单 | DELETE | /api/v1/reimbursements/{id} |

## **7. Reimbursement Item 报销子项**

| **功能** | **方法** | **URL** |
| --- | --- | --- |
| 添加子项 | POST | /api/v1/reimbursements/{id}/items |
| 获取子项 | GET | /api/v1/reimbursements/{id}/items |
| 删除子项 | DELETE | /api/v1/reimbursements/items/{item_id} |