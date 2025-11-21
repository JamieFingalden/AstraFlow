// src/utils/userUtils.js

// 辅助函数：获取角色显示名称
export const getRoleDisplayName = (role) => {
  const roleMap = {
    'admin': '系统管理员',
    'normal': '普通用户',
    'personal': '个人用户',
    'tenant_admin': '企业管理员',
    'tenant_user': '企业用户'
  }
  return roleMap[role] || role
}

// 根据角色获取租户角色
export const getTenantRoleFromRole = (role) => {
  if (role === 'admin' || role === 'tenant_admin') return 'admin'
  if (role === 'normal' || role === 'tenant_user') return 'normal'
  if (role === 'personal') return 'personal'
  return 'normal'
}