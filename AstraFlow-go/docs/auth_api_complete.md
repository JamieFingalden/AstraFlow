# AstraFlow API 文档（Auth认证模块）

## **📘 AstraFlow API 文档（v1）**

AstraFlow 是一个面向中小微企业与个人用户的财务管理系统，支持多租户、个人用户、OCR 发票识别、报销流转等核心功能。

## **🌐 Auth 认证接口**

| **功能** | **方法** | **URL** |
| --- | --- | --- |
| 用户注册 | POST | /api/v1/auth/register |
| 用户登录 | POST | /api/v1/auth/login |
| 刷新 Token | POST | /api/v1/auth/refresh |
| 获取当前用户信息 | GET | /api/v1/auth/me |

### **1. 用户注册 - POST /api/v1/auth/register**

#### 请求参数
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| username | string | 是 | 登录用户名，长度3-50 |
| password | string | 是 | 密码，长度6-100 |
| email | string | 否 | 邮箱 |
| phone | string | 否 | 手机号 |
| tenant_id | int64 | 否 | 租户ID，NULL表示个人用户 |

#### 响应参数
| 参数名 | 类型 | 说明 |
|--------|------|------|
| code | int | 状态码，200表示成功 |
| message | string | 响应消息 |
| data | object | 响应数据 |

#### data字段说明
| 参数名 | 类型 | 说明 |
|--------|------|------|
| user | object | 用户信息 |
| token | string | JWT Token |

##### user字段说明
| 参数名 | 类型 | 说明 |
|--------|------|------|
| id | int64 | 用户ID |
| username | string | 用户名 |
| email | string | 邮箱 |
| phone | string | 手机号 |
| role | string | 用户角色 |
| created_at | string | 创建时间 |
| updated_at | string | 更新时间 |

#### 示例请求
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123",
    "email": "test@example.com"
  }'
```

#### 示例响应
```json
{
  "code": 200,
  "message": "注册成功",
  "data": {
    "user": {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com",
      "phone": "",
      "role": "normal",
      "created_at": "2025-11-16T11:10:04.472+08:00",
      "updated_at": "2025-11-16T11:10:04.472+08:00"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### **2. 用户登录 - POST /api/v1/auth/login**

#### 请求参数
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| username | string | 是 | 登录用户名 |
| password | string | 是 | 密码 |

#### 响应参数
| 参数名 | 类型 | 说明 |
|--------|------|------|
| code | int | 状态码，200表示成功 |
| message | string | 响应消息 |
| data | object | 响应数据 |

#### data字段说明
| 参数名 | 类型 | 说明 |
|--------|------|------|
| user | object | 用户信息 |
| token | string | JWT Token |
| refresh_token | string | 刷新Token |

#### 示例请求
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }'
```

#### 示例响应
```json
{
  "code": 200,
  "message": "登录成功",
  "data": {
    "user": {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com",
      "phone": "",
      "role": "normal",
      "created_at": "2025-11-16T11:10:04.472+08:00",
      "updated_at": "2025-11-16T11:10:04.472+08:00"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### **3. 刷新Token - POST /api/v1/auth/refresh**

#### 请求头
| 头部字段 | 值 | 说明 |
|----------|----|------|
| Authorization | Bearer {refresh_token} | 刷新Token |

#### 响应参数
| 参数名 | 类型 | 说明 |
|--------|------|------|
| code | int | 状态码，200表示成功 |
| message | string | 响应消息 |
| data | object | 响应数据 |

#### data字段说明
| 参数名 | 类型 | 说明 |
|--------|------|------|
| token | string | 新的JWT Token |

#### 示例请求
```bash
curl -X POST http://localhost:8080/api/v1/auth/refresh \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

#### 示例响应
```json
{
  "code": 200,
  "message": "Token刷新成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### **4. 获取当前用户信息 - GET /api/v1/auth/me**

#### 请求头
| 头部字段 | 值 | 说明 |
|----------|----|------|
| Authorization | Bearer {token} | 访问Token |

#### 响应参数
| 参数名 | 类型 | 说明 |
|--------|------|------|
| code | int | 状态码，200表示成功 |
| message | string | 响应消息 |
| data | object | 响应数据 |

#### data字段说明
| 参数名 | 类型 | 说明 |
|--------|------|------|
| user | object | 用户信息 |

#### 示例请求
```bash
curl -X GET http://localhost:8080/api/v1/auth/me \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

#### 示例响应
```json
{
  "code": 200,
  "message": "获取用户信息成功",
  "data": {
    "user": {
      "id": 1,
      "username": "testuser",
      "email": "test@example.com",
      "phone": "",
      "role": "normal",
      "created_at": "2025-11-16T11:10:04.472+08:00",
      "updated_at": "2025-11-16T11:10:04.472+08:00"
    }
  }
}
```

## 错误码说明

| 错误码 | 说明 |
|--------|------|
| 200 | 成功 |
| 400 | 请求参数错误 |
| 401 | 未授权或Token无效 |
| 404 | 用户不存在 |
| 409 | 用户名已存在 |
| 500 | 服务器内部错误 |

## **技术实现**

- **框架**: Gin Web框架
- **数据库**: MySQL + GORM ORM
- **认证**: JWT Token
- **密码加密**: bcrypt
- **配置管理**: viper
- **日志**: zap logger

后端项目结构:
```
AstraFlow-go/
├── cmd/
│   ├── server/main.go          # 服务启动入口
│   └── migrate/main.go         # 数据库迁移脚本
├── internal/
│   ├── api/                    # API路由
│   ├── database/               # 数据库连接
│   ├── handler/                # 控制器层
│   ├── model/                  # 数据模型
│   ├── repository/             # 数据访问层
│   └── service/                # 业务逻辑层
├── pkg/
│   ├── config/                 # 配置管理
│   ├── jwt/                    # JWT工具
│   ├── logger/                 # 日志模块
│   └── middleware/             # 中间件
├── docs/                       # 文档
└── go.mod/go.sum              # Go依赖
```