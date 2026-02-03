-- Demo data for AstraFlow database
-- Only contains INSERT statements

-- Disable foreign key checks
SET FOREIGN_KEY_CHECKS = 0;

-- -----------------------------
-- Insert Data
-- -----------------------------

-- Roles
INSERT INTO roles (id, `key`, name, description, created_at, updated_at) VALUES
(1, 'admin', 'Administrator', 'System Administrator', NOW(), NOW()),
(2, 'auditor', 'Finance Auditor', 'Finance Auditor', NOW(), NOW()),
(3, 'employee', 'Employee', 'Regular Employee', NOW(), NOW());

-- Tenants
INSERT INTO tenants (id, name, industry, contact_name, contact_phone, contact_email, created_at, updated_at) VALUES
(1, 'TechNova Solutions', 'Software', 'John Manager', '13900139000', 'contact@technova.com', NOW(), NOW());

-- Users (Password: 123456)
INSERT INTO users (id, tenant_id, role_id, username, password, name, email, phone, status, created_at, updated_at) VALUES
(1, 1, 1, 'admin', '$2a$10$L4748PJjRPOVGaQGXon/c.091eSsA66ND8Hd9IZ/hSyJLJni87QGq', 'System Admin', 'admin@technova.com', '13900000001', 1, NOW(), NOW()),
(2, 1, 2, 'auditor', '$2a$10$L4748PJjRPOVGaQGXon/c.091eSsA66ND8Hd9IZ/hSyJLJni87QGq', 'Jane Auditor', 'auditor@technova.com', '13900000002', 1, NOW(), NOW()),
(3, 1, 3, 'employee', '$2a$10$L4748PJjRPOVGaQGXon/c.091eSsA66ND8Hd9IZ/hSyJLJni87QGq', 'Bob Employee', 'bob@technova.com', '13900000003', 1, NOW(), NOW());

-- Invoices
INSERT INTO invoices (id, tenant_id, user_id, invoice_number, invoice_date, amount, vendor, category, status, created_at, updated_at) VALUES
(1, 1, 3, 'INV-2025-001', '2025-08-15', 50.00, 'KFC', 'Meals', 'confirmed', NOW(), NOW()),
(2, 1, 3, 'INV-2025-002', '2025-08-16', 35.50, 'Uber', 'Transport', 'confirmed', NOW(), NOW()),
(3, 1, 3, 'INV-2025-003', '2025-08-20', 120.00, 'Hilton Hotel', 'Accommodation', 'confirmed', NOW(), NOW()),
(4, 1, 3, 'INV-2025-004', '2025-09-01', 25.00, 'Starbucks', 'Meals', 'pending', NOW(), NOW()),
(5, 1, 3, 'INV-2025-005', '2025-09-02', 450.00, 'Delta Airlines', 'Travel', 'pending', NOW(), NOW()),
(6, 1, 3, 'INV-2025-006', '2025-09-10', 89.99, 'Office Depot', 'Office Supplies', 'pending', NOW(), NOW()),
(7, 1, 3, 'INV-2025-007', '2025-09-12', 15.00, 'Subway', 'Meals', 'pending', NOW(), NOW()),
(8, 1, 3, 'INV-2025-008', '2025-09-15', 200.00, 'Marriott', 'Accommodation', 'pending', NOW(), NOW()),
(9, 1, 3, 'INV-2025-009', '2025-09-18', 60.00, 'City Taxi', 'Transport', 'rejected', NOW(), NOW()),
(10, 1, 3, 'INV-2025-010', '2025-09-20', 1200.00, 'Apple Store', 'Office Supplies', 'pending', NOW(), NOW());

-- Re-enable foreign key checks
SET FOREIGN_KEY_CHECKS = 1;