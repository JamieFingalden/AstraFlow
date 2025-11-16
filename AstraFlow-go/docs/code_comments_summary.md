# 代码注释总结

本文档总结了AstraFlow认证模块中各个文件的函数级注释。

## 目录结构

```
AstraFlow-go/
├── cmd/
│   ├── server/main.go          # 服务启动入口
│   └── migrate/main.go         # 数据库迁移脚本
├── internal/
│   ├── api/router.go           # API路由配置
│   ├── database/database.go    # 数据库连接
│   ├── handler/auth_handler.go # 认证处理器
│   ├── model/user.go           # 用户数据模型
│   ├── repository/user_repository.go # 用户数据访问层
│   └── service/auth_service.go # 认证业务逻辑层
├── pkg/
│   ├── config/config.go        # 配置管理
│   ├── jwt/jwt.go             # JWT工具
│   ├── logger/logger.go       # 日志模块
│   └── middleware/auth.go     # 认证中间件
```

## 主要模块注释

### 1. 数据模型层 (`internal/model/user.go`)

- **User结构体**: 定义了用户的基本信息和关联关系
  - 包含用户ID、租户ID、用户名、密码、邮箱、手机号、角色等字段
  - 使用GORM标签定义数据库约束
  - 密码字段在JSON序列化时忽略

### 2. 数据库连接层 (`internal/database/database.go`)

- **InitDB()函数**: 初始化数据库连接
  - 从配置文件读取数据库连接参数
  - 构建DSN并建立MySQL连接
  - 提供全局数据库连接实例

### 3. 数据访问层 (`internal/repository/user_repository.go`)

- **UserRepository结构体**: 封装对用户表的CRUD操作
- **NewUserRepository()**: 创建用户仓库实例
- **Create()**: 创建新用户
- **FindByUsername()**: 根据用户名查找用户
- **FindByID()**: 根据ID查找用户
- **Update()**: 更新用户信息
- **Delete()**: 删除用户

### 4. 业务逻辑层 (`internal/service/auth_service.go`)

- **AuthService结构体**: 处理用户注册、登录等认证相关业务逻辑
- **NewAuthService()**: 创建认证服务实例
- **Register()**: 用户注册，包含用户名唯一性检查和密码加密
- **Login()**: 用户登录，验证用户名和密码
- **GetUserByID()**: 根据ID获取用户信息

### 5. JWT工具层 (`pkg/jwt/jwt.go`)

- **Claims结构体**: 定义JWT Token中包含的用户信息和标准声明
- **GenerateToken()**: 生成访问令牌（24小时过期）
- **ParseToken()**: 验证并解析JWT Token
- **GenerateRefreshToken()**: 生成刷新令牌（7天过期）

### 6. 中间件层 (`pkg/middleware/auth.go`)

- **AuthMiddleware()**: JWT认证中间件
  - 验证请求中的JWT Token
  - 提取用户信息并存储到上下文中
  - 支持Bearer前缀的Token格式

### 7. 控制器层 (`internal/handler/auth_handler.go`)

- **AuthHandler结构体**: 处理认证相关的HTTP请求
- **NewAuthHandler()**: 创建认证处理器实例
- **Register()**: 处理用户注册请求
- **Login()**: 处理用户登录请求
- **RefreshToken()**: 处理Token刷新请求
- **GetCurrentUser()**: 处理获取当前用户信息请求

### 8. 路由层 (`internal/api/router.go`)

- **InitRouter()函数**: 初始化API路由
  - 初始化数据库连接
  - 配置认证相关路由
  - 应用认证中间件到需要保护的接口

## 注释规范

所有函数都遵循以下注释规范：

1. **函数名注释**: 简要说明函数功能
2. **详细说明**: 解释函数的具体实现逻辑
3. **参数说明**: 描述函数参数的含义和用途
4. **返回值说明**: 描述函数返回值的含义
5. **业务逻辑说明**: 解释重要的业务处理逻辑

## 安全特性

1. **密码加密**: 使用bcrypt进行密码哈希
2. **JWT认证**: 使用JWT进行用户身份验证
3. **Token过期**: 设置合理的Token过期时间
4. **参数验证**: 对输入参数进行严格验证
5. **错误处理**: 完善的错误处理和用户友好的错误信息

## 数据安全

1. **密码不返回**: 在响应中不包含密码字段
2. **Token安全**: JWT Token使用配置的密钥签名
3. **中间件保护**: 敏感接口通过认证中间件保护
4. **软删除**: 支持用户软删除功能