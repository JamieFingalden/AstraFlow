// API URL 列表
export const API_URLS = {
  // 认证相关
  AUTH: {
    LOGIN: '/auth/login',
    REGISTER: '/auth/register',
    LOGOUT: '/auth/logout',
    REFRESH: '/auth/refresh',
    ME: '/auth/me',
  },

  // 租户相关
  TENANT: {
    LIST: '/tenants',
    DETAIL: (id) => `/tenants/${id}`,
    CREATE: '/tenants',
    UPDATE: (id) => `/tenants/${id}`,
  },

  // 发票相关
  INVOICE: {
    LIST: '/invoices',
    DETAIL: (id) => `/invoices/${id}`,
    CREATE: '/invoices',
    UPDATE: (id) => `/invoices/${id}`,
    UPLOAD: '/invoices/upload',
  },

  // OCR相关
  OCR: {
    PROCESS: '/ocr/process',
    RESULT: (id) => `/ocr/${id}`,
    HISTORY: '/ocr/history',
  },

  // 报销相关
  REIMBURSEMENT: {
    LIST: '/reimbursements',
    DETAIL: (id) => `/reimbursements/${id}`,
    CREATE: '/reimbursements',
    UPDATE: (id) => `/reimbursements/${id}`,
    SUBMIT: (id) => `/reimbursements/${id}/submit`,
    APPROVE: (id) => `/reimbursements/${id}/approve`,
    REJECT: (id) => `/reimbursements/${id}/reject`,
  },

  // 分析和报告
  ANALYTICS: {
    DASHBOARD: '/analytics/dashboard',
    REPORTS: '/analytics/reports',
    TRENDS: '/analytics/trends',
  },

  // 设置相关
  SETTINGS: {
    USER: '/settings',
    CONFIG: '/config',
    NOTIFICATIONS: '/notifications/preferences',
  }
}

export default API_URLS