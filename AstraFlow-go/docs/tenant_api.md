# 租户(Tenant) API 接口文档

## 接口概述

租户接口用于管理企业租户信息，支持多租户架构下的企业信息管理功能。所有租户相关接口都需要用户身份认证。

## 接口列表

### 1. 创建租户
- **接口路径**: `POST /api/v1/tenants`
- **接口描述**: 创建新的企业租户
- **认证要求**: 需要有效的JWT Token
- **请求头**:
  - `Authorization: Bearer {token}`
  - `Content-Type: application/json`

#### 请求体参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| name | string | 是 | 租户名称（公司名称），长度1-100字符 |
| industry | string | 否 | 行业 |
| contact_name | string | 否 | 联系人姓名 |
| contact_phone | string | 否 | 联系电话 |
| contact_email | string | 否 | 联系邮箱 |

#### 请求示例

```json
{
  "name": "ABC科技有限公司",
  "industry": "互联网",
  "contact_name": "张三",
  "contact_phone": "13800138000",
  "contact_email": "zhangsan@abc.com"
}
```

#### 响应参数

| 参数名 | 类型 | 说明 |
|--------|------|------|
| code | int | 响应状态码 |
| message | string | 响应消息 |
| data.tenant | object | 创建的租户信息 |

#### 响应示例

```json
{
  "code": 200,
  "message": "创建租户成功",
  "data": {
    "tenant": {
      "id": 1,
      "name": "ABC科技有限公司",
      "industry": "互联网",
      "contact_name": "张三",
      "contact_phone": "13800138000",
      "contact_email": "zhangsan@abc.com",
      "created_at": "2024-01-01T12:00:00Z",
      "updated_at": "2024-01-01T12:00:00Z"
    }
  }
}
```

#### 错误响应

- `400`: 请求参数错误
- `401`: 未授权访问
- `409`: 租户名称已存在
- `500`: 服务器内部错误

---

### 2. 获取租户列表
- **接口路径**: `GET /api/v1/tenants`
- **接口描述**: 获取租户列表，支持分页
- **认证要求**: 需要有效的JWT Token
- **请求头**:
  - `Authorization: Bearer {token}`

#### 查询参数

| 参数名 | 类型 | 必填 | 默认值 | 说明 |
|--------|------|------|--------|------|
| page | int | 否 | 1 | 页码 |
| page_size | int | 否 | 10 | 每页数量 |

#### 请求示例

```
GET /api/v1/tenants?page=1&page_size=10
```

#### 响应参数

| 参数名 | 类型 | 说明 |
|--------|------|------|
| code | int | 响应状态码 |
| message | string | 响应消息 |
| data.tenants | array | 租户列表 |
| data.page | int | 当前页码 |
| data.size | int | 当前页记录数 |

#### 响应示例

```json
{
  "code": 200,
  "message": "获取租户列表成功",
  "data": {
    "tenants": [
      {
        "id": 1,
        "name": "ABC科技有限公司",
        "industry": "互联网",
        "contact_name": "张三",
        "contact_phone": "13800138000",
        "contact_email": "zhangsan@abc.com",
        "created_at": "2024-01-01T12:00:00Z",
        "updated_at": "2024-01-01T12:00:00Z"
      }
    ],
    "page": 1,
    "size": 1
  }
}
```

#### 错误响应

- `401`: 未授权访问
- `500`: 服务器内部错误

---

### 3. 获取租户详情
- **接口路径**: `GET /api/v1/tenants/{id}`
- **接口描述**: 根据ID获取租户详细信息
- **认证要求**: 需要有效的JWT Token
- **请求头**:
  - `Authorization: Bearer {token}`

#### 路径参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int64 | 是 | 租户ID |

#### 请求示例

```
GET /api/v1/tenants/1
```

#### 响应参数

| 参数名 | 类型 | 说明 |
|--------|------|------|
| code | int | 响应状态码 |
| message | string | 响应消息 |
| data.tenant | object | 租户详细信息 |

#### 响应示例

```json
{
  "code": 200,
  "message": "获取租户成功",
  "data": {
    "tenant": {
      "id": 1,
      "name": "ABC科技有限公司",
      "industry": "互联网",
      "contact_name": "张三",
      "contact_phone": "13800138000",
      "contact_email": "zhangsan@abc.com",
      "created_at": "2024-01-01T12:00:00Z",
      "updated_at": "2024-01-01T12:00:00Z"
    }
  }
}
```

#### 错误响应

- `400`: 租户ID格式错误
- `401`: 未授权访问
- `404`: 租户不存在
- `500`: 服务器内部错误

---

### 4. 更新租户
- **接口路径**: `PUT /api/v1/tenants/{id}`
- **接口描述**: 根据ID更新租户信息
- **认证要求**: 需要有效的JWT Token
- **请求头**:
  - `Authorization: Bearer {token}`
  - `Content-Type: application/json`

#### 路径参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int64 | 是 | 租户ID |

#### 请求体参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| name | string | 是 | 租户名称（公司名称），长度1-100字符 |
| industry | string | 否 | 行业 |
| contact_name | string | 否 | 联系人姓名 |
| contact_phone | string | 否 | 联系电话 |
| contact_email | string | 否 | 联系邮箱 |

#### 请求示例

```json
{
  "name": "ABC科技有限公司",
  "industry": "互联网科技",
  "contact_name": "李四",
  "contact_phone": "13900139000",
  "contact_email": "lisi@abc.com"
}
```

#### 响应参数

| 参数名 | 类型 | 说明 |
|--------|------|------|
| code | int | 响应状态码 |
| message | string | 响应消息 |
| data.tenant | object | 更新后的租户信息 |

#### 响应示例

```json
{
  "code": 200,
  "message": "更新租户成功",
  "data": {
    "tenant": {
      "id": 1,
      "name": "ABC科技有限公司",
      "industry": "互联网科技",
      "contact_name": "李四",
      "contact_phone": "13900139000",
      "contact_email": "lisi@abc.com",
      "created_at": "2024-01-01T12:00:00Z",
      "updated_at": "2024-01-02T10:00:00Z"
    }
  }
}
```

#### 错误响应

- `400`: 请求参数错误或租户ID格式错误
- `401`: 未授权访问
- `404`: 租户不存在
- `409`: 租户名称已存在
- `500`: 服务器内部错误

---

### 5. 删除租户
- **接口路径**: `DELETE /api/v1/tenants/{id}`
- **接口描述**: 根据ID删除租户
- **认证要求**: 需要有效的JWT Token
- **请求头**:
  - `Authorization: Bearer {token}`

#### 路径参数

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int64 | 是 | 租户ID |

#### 请求示例

```
DELETE /api/v1/tenants/1
```

#### 响应参数

| 参数名 | 类型 | 说明 |
|--------|------|------|
| code | int | 响应状态码 |
| message | string | 响应消息 |

#### 响应示例

```json
{
  "code": 200,
  "message": "删除租户成功"
}
```

#### 错误响应

- `400`: 租户ID格式错误
- `401`: 未授权访问
- `404`: 租户不存在
- `500`: 服务器内部错误

---

## 数据模型

### 租户(Tenant)模型

| 字段名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | int64 | 是 | 租户ID，主键，自动递增 |
| name | string | 是 | 租户名称（公司名称） |
| industry | string | 否 | 行业 |
| contact_name | string | 否 | 联系人姓名 |
| contact_phone | string | 否 | 联系电话 |
| contact_email | string | 否 | 联系邮箱 |
| created_at | string | 是 | 创建时间 |
| updated_at | string | 是 | 更新时间 |

## 认证说明

所有租户相关接口都需要用户身份认证，请求头中必须包含有效的JWT Token：

```
Authorization: Bearer {token}
```

## 错误码说明

| 错误码 | 说明 |
|--------|------|
| 200 | 操作成功 |
| 400 | 请求参数错误 |
| 401 | 未授权访问 |
| 404 | 资源不存在 |
| 409 | 资源冲突（如租户名称已存在） |
| 500 | 服务器内部错误 |