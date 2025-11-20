-- ==========================================================
-- AstraFlow 多租户 + 个人用户 模式数据库（软连接版）
-- MySQL 8.0
-- ==========================================================

-- -----------------------------
-- 租户表（企业模式才会用）
-- -----------------------------
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


-- -----------------------------
-- 用户表（租户用户 或 个人用户）
-- -----------------------------
CREATE TABLE user (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    tenant_id BIGINT DEFAULT NULL COMMENT '租户ID，NULL 表示个人用户',
    role_id BIGINT DEFAULT NULL COMMENT '角色ID，引用role表，NULL表示默认普通用户',
    username VARCHAR(50) NOT NULL UNIQUE COMMENT '登录用户名',
    password VARCHAR(255) NOT NULL COMMENT '加密后的密码',
    email VARCHAR(100) DEFAULT NULL COMMENT '邮箱',
    phone VARCHAR(50) DEFAULT NULL COMMENT '手机号',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_user_tenant (tenant_id),
    INDEX idx_user_role (role_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表，支持租户用户与个人用户';


-- -----------------------------
-- 发票信息表（invoice）
-- -----------------------------
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


-- -----------------------------
-- OCR 识别结果表（结构化 + 非结构化）
-- -----------------------------
CREATE TABLE ocr_result (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    tenant_id BIGINT DEFAULT NULL COMMENT '租户ID / 个人用户 NULL',
    user_id BIGINT NOT NULL COMMENT '关联用户ID',
    invoice_id BIGINT NOT NULL COMMENT '关联发票ID（软连接）',
    -- 结构化字段（便于查询）
    amount DECIMAL(12,2) DEFAULT NULL COMMENT '金额',
    invoice_date DATE DEFAULT NULL COMMENT '开票日期',
    vendor VARCHAR(255) DEFAULT NULL COMMENT '商户',
    category VARCHAR(100) DEFAULT NULL COMMENT '类别',
    invoice_number VARCHAR(100) DEFAULT NULL COMMENT '发票号',
    confidence DECIMAL(5,2) DEFAULT NULL COMMENT '识别置信度，例如 0.97',
    -- 全量 OCR 内容
    raw_text LONGTEXT COMMENT 'OCR 原始文本',
    parsed_json JSON COMMENT 'OCR 解析后的 JSON',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_ocr_tenant (tenant_id),
    INDEX idx_ocr_user (user_id),
    INDEX idx_ocr_invoice (invoice_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='OCR 识别结果表';


-- -----------------------------
-- 附件表（发票照片/支付截图/报销凭证）
-- -----------------------------
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


-- -----------------------------
-- 报销单主表
-- -----------------------------
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


-- -----------------------------
-- 报销单子项（关联发票）
-- -----------------------------
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


-- -----------------------------
-- 角色表
-- -----------------------------
CREATE TABLE role (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL UNIQUE COMMENT '角色名称',
    display_name VARCHAR(100) NOT NULL COMMENT '角色显示名称',
    description VARCHAR(255) DEFAULT NULL COMMENT '角色描述',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_role_name (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表，定义用户角色';


-- -----------------------------
-- 初始化角色数据
-- -----------------------------
INSERT INTO role (name, display_name, description) VALUES
('admin', '管理员', '拥有系统所有权限的管理员'),
('normal', '普通用户', '普通租户用户，具有基本操作权限'),
('personal', '个人用户', '个人用户，仅管理自己的数据');
