-- Demo data for AstraFlow database
-- Generated from mock data in BillManagement.vue

-- Insert demo user (personal user with user_id = 1)
INSERT INTO users (id, tenant_id, role_id, username, password, email, phone, created_at, updated_at) VALUES
(1, NULL, 3, 'demo_user', '$2a$10$L4748PJjRPOVGaQGXon/c.091eSsA66ND8Hd9IZ/hSyJLJni87QGq', 'demo@astraflow.com', '13800138000', NOW(), NOW());

-- Insert demo invoices
INSERT INTO invoices (id, tenant_id, user_id, invoice_number, invoice_date, amount, vendor, tax_id, category, payment_source, status, created_at, updated_at) VALUES
(1, NULL, 1, 'BILL-001', '2025-11-08', 35.50, '星巴克', NULL, '餐饮', '微信支付', 'confirmed', NOW(), NOW()),
(2, NULL, 1, 'BILL-002', '2025-11-08', 45.00, '滴滴出行', NULL, '交通', '支付宝', 'confirmed', NOW(), NOW()),
(3, NULL, 1, 'BILL-003', '2025-11-07', 128.90, '京东', NULL, '办公', '银行卡', 'pending', NOW(), NOW()),
(4, NULL, 1, 'BILL-004', '2025-11-06', 280.00, '如家酒店', NULL, '住宿', '手动添加', 'confirmed', NOW(), NOW()),
(5, NULL, 1, 'BILL-005', '2025-11-05', 89.60, '美团', NULL, '餐饮', '支付宝', 'pending', NOW(), NOW()),
(6, NULL, 1, 'BILL-006', '2025-11-01', 200.00, '地铁', NULL, '交通', '手动添加', 'confirmed', NOW(), NOW()),
(7, NULL, 1, 'BILL-007', '2025-10-30', 52.00, '肯德基', NULL, '餐饮', '微信支付', 'pending', NOW(), NOW()),
(8, NULL, 1, 'BILL-008', '2025-10-28', 67.30, '办公用品店', NULL, '办公', '银行卡', 'confirmed', NOW(), NOW()),
(9, NULL, 1, 'BILL-009', '2025-10-25', 78.50, '出租车', NULL, '交通', '现金', 'confirmed', NOW(), NOW()),
(10, NULL, 1, 'BILL-010', '2025-10-24', 25.80, '便利店', NULL, '其他', '手动添加', 'pending', NOW(), NOW()),
(11, NULL, 1, 'BILL-011', '2025-10-22', 168.00, '海底捞', NULL, '餐饮', '支付宝', 'confirmed', NOW(), NOW()),
(12, NULL, 1, 'BILL-012', '2025-10-20', 95.00, '滴滴专车', NULL, '交通', '微信支付', 'confirmed', NOW(), NOW()),
(13, NULL, 1, 'BILL-013', '2025-10-18', 45.60, '超市', NULL, '其他', '银行卡', 'pending', NOW(), NOW()),
(14, NULL, 1, 'BILL-014', '2025-10-15', 15.00, '打印服务', NULL, '办公', '手动添加', 'confirmed', NOW(), NOW()),
(15, NULL, 1, 'BILL-015', '2025-10-12', 38.00, '麦当劳', NULL, '餐饮', '微信支付', 'pending', NOW(), NOW());

-- Insert corresponding OCR results
INSERT INTO ocr_results (id, tenant_id, user_id, invoice_id, amount, invoice_date, vendor, category, invoice_number, confidence, raw_text, parsed_json, created_at, updated_at) VALUES
(1, NULL, 1, 1, 35.50, '2025-11-08', '星巴克', '餐饮', 'BILL-001', 0.98, '星巴克咖啡 金额: 35.50', JSON_OBJECT(), NOW(), NOW()),
(2, NULL, 1, 2, 45.00, '2025-11-08', '滴滴出行', '交通', 'BILL-002', 0.97, '滴滴出行 金额: 45.00', JSON_OBJECT(), NOW(), NOW()),
(3, NULL, 1, 3, 128.90, '2025-11-07', '京东', '办公', 'BILL-003', 0.96, '京东文具 金额: 128.90', JSON_OBJECT(), NOW(), NOW()),
(4, NULL, 1, 4, 280.00, '2025-11-06', '如家酒店', '住宿', 'BILL-004', 0.99, '如家酒店 金额: 280.00', JSON_OBJECT(), NOW(), NOW()),
(5, NULL, 1, 5, 89.60, '2025-11-05', '美团', '餐饮', 'BILL-005', 0.95, '美团外卖 金额: 89.60', JSON_OBJECT(), NOW(), NOW()),
(6, NULL, 1, 6, 200.00, '2025-11-01', '地铁', '交通', 'BILL-006', 0.94, '地铁月卡 金额: 200.00', JSON_OBJECT(), NOW(), NOW()),
(7, NULL, 1, 7, 52.00, '2025-10-30', '肯德基', '餐饮', 'BILL-007', 0.98, '肯德基 金额: 52.00', JSON_OBJECT(), NOW(), NOW()),
(8, NULL, 1, 8, 67.30, '2025-10-28', '办公用品店', '办公', 'BILL-008', 0.97, '办公用品 金额: 67.30', JSON_OBJECT(), NOW(), NOW()),
(9, NULL, 1, 9, 78.50, '2025-10-25', '出租车', '交通', 'BILL-009', 0.96, '出租车 金额: 78.50', JSON_OBJECT(), NOW(), NOW()),
(10, NULL, 1, 10, 25.80, '2025-10-24', '便利店', '其他', 'BILL-010', 0.99, '便利店 金额: 25.80', JSON_OBJECT(), NOW(), NOW()),
(11, NULL, 1, 11, 168.00, '2025-10-22', '海底捞', '餐饮', 'BILL-011', 0.95, '海底捞 金额: 168.00', JSON_OBJECT(), NOW(), NOW()),
(12, NULL, 1, 12, 95.00, '2025-10-20', '滴滴专车', '交通', 'BILL-012', 0.94, '滴滴专车 金额: 95.00', JSON_OBJECT(), NOW(), NOW()),
(13, NULL, 1, 13, 45.60, '2025-10-18', '超市', '其他', 'BILL-013', 0.98, '超市购物 金额: 45.60', JSON_OBJECT(), NOW(), NOW()),
(14, NULL, 1, 14, 15.00, '2025-10-15', '打印服务', '办公', 'BILL-014', 0.97, '打印服务 金额: 15.00', JSON_OBJECT(), NOW(), NOW()),
(15, NULL, 1, 15, 38.00, '2025-10-12', '麦当劳', '餐饮', 'BILL-015', 0.96, '麦当劳 金额: 38.00', JSON_OBJECT(), NOW(), NOW());

-- Insert demo reimbursements if needed
INSERT INTO reimbursements (id, tenant_id, user_id, title, total_amount, status, created_at, updated_at) VALUES
(1, NULL, 1, '11月餐饮费用报销', 175.10, 'approved', NOW(), NOW()),
(2, NULL, 1, '11月交通费用报销', 323.50, 'approved', NOW(), NOW()),
(3, NULL, 1, '10-11月办公用品报销', 211.20, 'pending', NOW(), NOW());

-- Insert demo reimbursement items
INSERT INTO reimbursement_items (id, reimbursement_id, invoice_id, amount, note, created_at, updated_at) VALUES
(1, 1, 1, 35.50, '星巴克咖啡', NOW(), NOW()),
(2, 1, 5, 89.60, '美团外卖', NOW(), NOW()),
(3, 1, 7, 50.00, '肯德基', NOW(), NOW()),
(4, 2, 2, 45.00, '滴滴出行', NOW(), NOW()),
(5, 2, 6, 200.00, '地铁月卡', NOW(), NOW()),
(6, 2, 9, 78.50, '出租车', NOW(), NOW()),
(7, 3, 3, 128.90, '办公用品', NOW(), NOW()),
(8, 3, 8, 67.30, '打印纸笔', NOW(), NOW()),
(9, 3, 14, 15.00, '打印服务', NOW(), NOW());