-- Demo data for AstraFlow database
-- 3 months of comprehensive demo data (September - November 2025)
-- Supports both personal and tenant users with proper multi-tenant architecture

-- Clean existing data (delete in proper order due to foreign key relationships)
DELETE FROM reimbursement_items;
DELETE FROM reimbursements;
DELETE FROM ocr_results;
DELETE FROM invoices;
DELETE FROM users;
DELETE FROM tenants;

-- Reset auto-increment sequences
ALTER TABLE tenants AUTO_INCREMENT = 1;
ALTER TABLE users AUTO_INCREMENT = 1;
ALTER TABLE invoices AUTO_INCREMENT = 1;
ALTER TABLE ocr_results AUTO_INCREMENT = 1;
ALTER TABLE reimbursements AUTO_INCREMENT = 1;
ALTER TABLE reimbursement_items AUTO_INCREMENT = 1;

-- Insert demo tenant
INSERT INTO tenants (id, name, industry, contact_name, contact_phone, contact_email, created_at, updated_at) VALUES
(1, '科技有限公司', '软件开发', '张经理', '13900139000', 'contact@demo-tenant.com', NOW(), NOW());

-- Insert users (preserve original demo_user credentials)
INSERT INTO users (id, tenant_id, username, password, email, phone, role, created_at, updated_at) VALUES
(1, NULL, 'demo_user', '$2a$10$L4748PJjRPOVGaQGXon/c.091eSsA66ND8Hd9IZ/hSyJLJni87QGq', 'demo@astraflow.com', '13800138000', 'normal', NOW(), NOW()),
(2, 1, 'tenant_user', '$2a$10$L4748PJjRPOVGaQGXon/c.091eSsA66ND8Hd9IZ/hSyJLJni87QGq', 'tenant@demo-tenant.com', '13900139001', 'admin', NOW(), NOW());

-- Insert September 2025 invoices (10 total: 5 personal, 5 tenant)
INSERT INTO invoices (id, tenant_id, user_id, invoice_number, invoice_date, amount, vendor, tax_id, category, payment_source, status, created_at, updated_at) VALUES
-- Personal user invoices (September)
(1, NULL, 1, 'BILL-20250901', '2025-09-03', 35.50, '星巴克', NULL, '餐饮', '微信支付', 'confirmed', NOW(), NOW()),
(2, NULL, 1, 'BILL-20250902', '2025-09-05', 45.00, '滴滴出行', NULL, '交通', '支付宝', 'confirmed', NOW(), NOW()),
(3, NULL, 1, 'BILL-20250903', '2025-09-08', 128.90, '京东', NULL, '办公', '银行卡', 'pending', NOW(), NOW()),
(4, NULL, 1, 'BILL-20250904', '2025-09-12', 89.60, '美团外卖', NULL, '餐饮', '支付宝', 'confirmed', NOW(), NOW()),
(5, NULL, 1, 'BILL-20250905', '2025-09-15', 200.00, '地铁月卡', NULL, '交通', '手动添加', 'confirmed', NOW(), NOW()),
-- Tenant user invoices (September)
(6, 1, 2, 'TENANT-20250901', '2025-09-02', 52.00, '肯德基', NULL, '餐饮', '微信支付', 'confirmed', NOW(), NOW()),
(7, 1, 2, 'TENANT-20250902', '2025-09-07', 280.00, '如家酒店', NULL, '住宿', '银行卡', 'confirmed', NOW(), NOW()),
(8, 1, 2, 'TENANT-20250903', '2025-09-10', 67.30, '办公用品店', NULL, '办公', '现金', 'confirmed', NOW(), NOW()),
(9, 1, 2, 'TENANT-20250904', '2025-09-18', 78.50, '出租车', NULL, '交通', '现金', 'pending', NOW(), NOW()),
(10, 1, 2, 'TENANT-20250905', '2025-09-22', 168.00, '海底捞', NULL, '餐饮', '支付宝', 'confirmed', NOW(), NOW());

-- Insert October 2025 invoices (10 total: 5 personal, 5 tenant)
INSERT INTO invoices (id, tenant_id, user_id, invoice_number, invoice_date, amount, vendor, tax_id, category, payment_source, status, created_at, updated_at) VALUES
-- Personal user invoices (October)
(11, NULL, 1, 'BILL-20251001', '2025-10-01', 95.00, '滴滴专车', NULL, '交通', '微信支付', 'confirmed', NOW(), NOW()),
(12, NULL, 1, 'BILL-20251002', '2025-10-05', 45.60, '超市', NULL, '其他', '银行卡', 'pending', NOW(), NOW()),
(13, NULL, 1, 'BILL-20251003', '2025-10-08', 15.00, '打印服务', NULL, '办公', '手动添加', 'confirmed', NOW(), NOW()),
(14, NULL, 1, 'BILL-20251004', '2025-10-12', 38.00, '麦当劳', NULL, '餐饮', '微信支付', 'confirmed', NOW(), NOW()),
(15, NULL, 1, 'BILL-20251005', '2025-10-18', 156.00, '火车票', NULL, '交通', '支付宝', 'confirmed', NOW(), NOW()),
-- Tenant user invoices (October)
(16, 1, 2, 'TENANT-20251001', '2025-10-02', 88.00, '饿了么', NULL, '餐饮', '微信支付', 'confirmed', NOW(), NOW()),
(17, 1, 2, 'TENANT-20251002', '2025-10-09', 320.00, '汉庭酒店', NULL, '住宿', '银行卡', 'confirmed', NOW(), NOW()),
(18, 1, 2, 'TENANT-20251003', '2025-10-14', 245.00, '文具店', NULL, '办公', '支付宝', 'pending', NOW(), NOW()),
(19, 1, 2, 'TENANT-20251004', '2025-10-20', 65.00, '机场快线', NULL, '交通', '现金', 'confirmed', NOW(), NOW()),
(20, 1, 2, 'TENANT-20251005', '2025-10-25', 198.00, '药店', NULL, '其他', '银行卡', 'confirmed', NOW(), NOW());

-- Insert November 2025 invoices (10 total: 5 personal, 5 tenant)
INSERT INTO invoices (id, tenant_id, user_id, invoice_number, invoice_date, amount, vendor, tax_id, category, payment_source, status, created_at, updated_at) VALUES
-- Personal user invoices (November)
(21, NULL, 1, 'BILL-20251101', '2025-11-02', 28.50, '便利店', NULL, '其他', '现金', 'pending', NOW(), NOW()),
(22, NULL, 1, 'BILL-20251102', '2025-11-06', 225.00, '如家酒店', NULL, '住宿', '银行卡', 'confirmed', NOW(), NOW()),
(23, NULL, 1, 'BILL-20251103', '2025-11-09', 67.80, '星巴克', NULL, '餐饮', '微信支付', 'confirmed', NOW(), NOW()),
(24, NULL, 1, 'BILL-20251104', '2025-11-15', 89.90, '京东', NULL, '办公', '支付宝', 'pending', NOW(), NOW()),
(25, NULL, 1, 'BILL-20251105', '2025-11-20', 145.00, '滴滴出行', NULL, '交通', '微信支付', 'confirmed', NOW(), NOW()),
-- Tenant user invoices (November)
(26, 1, 2, 'TENANT-20251101', '2025-11-01', 76.50, '肯德基', NULL, '餐饮', '支付宝', 'confirmed', NOW(), NOW()),
(27, 1, 2, 'TENANT-20251102', '2025-11-08', 380.00, 'Airbnb', NULL, '住宿', '银行卡', 'confirmed', NOW(), NOW()),
(28, 1, 2, 'TENANT-20251103', '2025-11-12', 156.70, '办公用品店', NULL, '办公', '微信支付', 'pending', NOW(), NOW()),
(29, 1, 2, 'TENANT-20251104', '2025-11-18', 92.00, '出租车', NULL, '交通', '现金', 'confirmed', NOW(), NOW()),
(30, 1, 2, 'TENANT-20251105', '2025-11-24', 234.50, '海底捞', NULL, '餐饮', '支付宝', 'confirmed', NOW(), NOW());

-- Insert OCR results for all invoices
INSERT INTO ocr_results (id, tenant_id, user_id, invoice_id, amount, invoice_date, vendor, category, invoice_number, confidence, raw_text, parsed_json, created_at, updated_at) VALUES
-- September OCR results
(1, NULL, 1, 1, 35.50, '2025-09-03', '星巴克', '餐饮', 'BILL-20250901', 0.98, '星巴克咖啡 金额: 35.50', JSON_OBJECT(), NOW(), NOW()),
(2, NULL, 1, 2, 45.00, '2025-09-05', '滴滴出行', '交通', 'BILL-20250902', 0.97, '滴滴出行 金额: 45.00', JSON_OBJECT(), NOW(), NOW()),
(3, NULL, 1, 3, 128.90, '2025-09-08', '京东', '办公', 'BILL-20250903', 0.96, '京东文具 金额: 128.90', JSON_OBJECT(), NOW(), NOW()),
(4, NULL, 1, 4, 89.60, '2025-09-12', '美团外卖', '餐饮', 'BILL-20250904', 0.95, '美团外卖 金额: 89.60', JSON_OBJECT(), NOW(), NOW()),
(5, NULL, 1, 5, 200.00, '2025-09-15', '地铁', '交通', 'BILL-20250905', 0.94, '地铁月卡 金额: 200.00', JSON_OBJECT(), NOW(), NOW()),
(6, 1, 2, 6, 52.00, '2025-09-02', '肯德基', '餐饮', 'TENANT-20250901', 0.98, '肯德基 金额: 52.00', JSON_OBJECT(), NOW(), NOW()),
(7, 1, 2, 7, 280.00, '2025-09-07', '如家酒店', '住宿', 'TENANT-20250902', 0.99, '如家酒店 金额: 280.00', JSON_OBJECT(), NOW(), NOW()),
(8, 1, 2, 8, 67.30, '2025-09-10', '办公用品店', '办公', 'TENANT-20250903', 0.97, '办公用品 金额: 67.30', JSON_OBJECT(), NOW(), NOW()),
(9, 1, 2, 9, 78.50, '2025-09-18', '出租车', '交通', 'TENANT-20250904', 0.96, '出租车 金额: 78.50', JSON_OBJECT(), NOW(), NOW()),
(10, 1, 2, 10, 168.00, '2025-09-22', '海底捞', '餐饮', 'TENANT-20250905', 0.95, '海底捞 金额: 168.00', JSON_OBJECT(), NOW(), NOW()),
-- October OCR results
(11, NULL, 1, 11, 95.00, '2025-10-01', '滴滴专车', '交通', 'BILL-20251001', 0.94, '滴滴专车 金额: 95.00', JSON_OBJECT(), NOW(), NOW()),
(12, NULL, 1, 12, 45.60, '2025-10-05', '超市', '其他', 'BILL-20251002', 0.98, '超市购物 金额: 45.60', JSON_OBJECT(), NOW(), NOW()),
(13, NULL, 1, 13, 15.00, '2025-10-08', '打印服务', '办公', 'BILL-20251003', 0.97, '打印服务 金额: 15.00', JSON_OBJECT(), NOW(), NOW()),
(14, NULL, 1, 14, 38.00, '2025-10-12', '麦当劳', '餐饮', 'BILL-20251004', 0.96, '麦当劳 金额: 38.00', JSON_OBJECT(), NOW(), NOW()),
(15, NULL, 1, 15, 156.00, '2025-10-18', '火车票', '交通', 'BILL-20251005', 0.99, '火车票 金额: 156.00', JSON_OBJECT(), NOW(), NOW()),
(16, 1, 2, 16, 88.00, '2025-10-02', '饿了么', '餐饮', 'TENANT-20251001', 0.95, '饿了么外卖 金额: 88.00', JSON_OBJECT(), NOW(), NOW()),
(17, 1, 2, 17, 320.00, '2025-10-09', '汉庭酒店', '住宿', 'TENANT-20251002', 0.98, '汉庭酒店 金额: 320.00', JSON_OBJECT(), NOW(), NOW()),
(18, 1, 2, 18, 245.00, '2025-10-14', '文具店', '办公', 'TENANT-20251003', 0.96, '文具店 金额: 245.00', JSON_OBJECT(), NOW(), NOW()),
(19, 1, 2, 19, 65.00, '2025-10-20', '机场快线', '交通', 'TENANT-20251004', 0.97, '机场快线 金额: 65.00', JSON_OBJECT(), NOW(), NOW()),
(20, 1, 2, 20, 198.00, '2025-10-25', '药店', '其他', 'TENANT-20251005', 0.94, '药店购物 金额: 198.00', JSON_OBJECT(), NOW(), NOW()),
-- November OCR results
(21, NULL, 1, 21, 28.50, '2025-11-02', '便利店', '其他', 'BILL-20251101', 0.99, '便利店 金额: 28.50', JSON_OBJECT(), NOW(), NOW()),
(22, NULL, 1, 22, 225.00, '2025-11-06', '如家酒店', '住宿', 'BILL-20251102', 0.95, '如家酒店 金额: 225.00', JSON_OBJECT(), NOW(), NOW()),
(23, NULL, 1, 23, 67.80, '2025-11-09', '星巴克', '餐饮', 'BILL-20251103', 0.98, '星巴克咖啡 金额: 67.80', JSON_OBJECT(), NOW(), NOW()),
(24, NULL, 1, 24, 89.90, '2025-11-15', '京东', '办公', 'BILL-20251104', 0.96, '京东办公用品 金额: 89.90', JSON_OBJECT(), NOW(), NOW()),
(25, NULL, 1, 25, 145.00, '2025-11-20', '滴滴出行', '交通', 'BILL-20251105', 0.97, '滴滴出行 金额: 145.00', JSON_OBJECT(), NOW(), NOW()),
(26, 1, 2, 26, 76.50, '2025-11-01', '肯德基', '餐饮', 'TENANT-20251101', 0.95, '肯德基 金额: 76.50', JSON_OBJECT(), NOW(), NOW()),
(27, 1, 2, 27, 380.00, '2025-11-08', 'Airbnb', '住宿', 'TENANT-20251102', 0.98, 'Airbnb住宿 金额: 380.00', JSON_OBJECT(), NOW(), NOW()),
(28, 1, 2, 28, 156.70, '2025-11-12', '办公用品店', '办公', 'TENANT-20251103', 0.96, '办公用品 金额: 156.70', JSON_OBJECT(), NOW(), NOW()),
(29, 1, 2, 29, 92.00, '2025-11-18', '出租车', '交通', 'TENANT-20251104', 0.97, '出租车 金额: 92.00', JSON_OBJECT(), NOW(), NOW()),
(30, 1, 2, 30, 234.50, '2025-11-24', '海底捞', '餐饮', 'TENANT-20251105', 0.95, '海底捞 金额: 234.50', JSON_OBJECT(), NOW(), NOW());

-- Insert simple reimbursements
INSERT INTO reimbursements (id, tenant_id, user_id, title, total_amount, status, created_at, updated_at) VALUES
-- Personal user reimbursements
(1, NULL, 1, '9月餐饮交通费用报销', 280.50, 'approved', NOW(), NOW()),
(2, NULL, 1, '10-11月办公费用报销', 233.80, 'pending', NOW(), NOW()),
-- Tenant user reimbursements
(3, 1, 2, '9月差旅费用报销', 425.50, 'approved', NOW(), NOW()),
(4, 1, 2, '10-11月综合费用报销', 1202.70, 'pending', NOW(), NOW());

-- Insert reimbursement items
INSERT INTO reimbursement_items (id, reimbursement_id, invoice_id, amount, note, created_at, updated_at) VALUES
-- Personal user reimbursement items
(1, 1, 1, 35.50, '星巴克咖啡', NOW(), NOW()),
(2, 1, 2, 45.00, '滴滴出行', NOW(), NOW()),
(3, 1, 4, 89.60, '美团外卖', NOW(), NOW()),
(4, 1, 5, 200.00, '地铁月卡', NOW(), NOW()),
(5, 2, 3, 128.90, '京东办公用品', NOW(), NOW()),
(6, 2, 13, 15.00, '打印服务', NOW(), NOW()),
(7, 2, 24, 89.90, '京东办公用品', NOW(), NOW()),
-- Tenant user reimbursement items
(8, 3, 6, 52.00, '肯德基', NOW(), NOW()),
(9, 3, 7, 280.00, '如家酒店住宿', NOW(), NOW()),
(10, 3, 9, 78.50, '出租车费用', NOW(), NOW()),
(11, 4, 8, 67.30, '办公用品', NOW(), NOW()),
(12, 4, 16, 88.00, '饿了么工作餐', NOW(), NOW()),
(13, 4, 17, 320.00, '汉庭酒店住宿', NOW(), NOW()),
(14, 4, 18, 245.00, '文具采购', NOW(), NOW()),
(15, 4, 27, 380.00, 'Airbnb住宿', NOW(), NOW()),
(16, 4, 28, 156.70, '办公用品采购', NOW(), NOW());