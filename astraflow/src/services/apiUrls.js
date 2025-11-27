// API URL 列表
export const API_URLS = {
  // 认证相关
  AUTH: {
    LOGIN: '/v1/auth/login',
    REGISTER: '/v1/auth/register',
    LOGOUT: '/v1/auth/logout',
    REFRESH: '/v1/auth/refresh',
    ME: '/v1/auth/me',
  },

  // 租户相关
  TENANT: {
    LIST: '/v1/tenants',
    DETAIL: (id) => `/v1/tenants/${id}`,
    CREATE: '/v1/tenants',
    UPDATE: (id) => `/v1/tenants/${id}`,
  },

  // 发票相关
  INVOICE: {
    LIST: '/v1/invoices',
    DETAIL: (id) => `/v1/invoices/${id}`,
    CREATE: '/v1/invoices',
    UPDATE: (id) => `/v1/invoices/${id}`,
    DELETE: (id) => `/v1/invoices/${id}`,
    BY_USER: '/v1/invoicesByUser',
    BY_TENANT: '/v1/invoicesByTenant',
    UPLOAD: '/v1/invoices/upload',
  },

  // OCR相关
  OCR: {
    PROCESS: '/v1/ocr/process',
    RESULT: (id) => `/v1/ocr/${id}`,
    HISTORY: '/v1/ocr/history',
  },

  // 报销相关
  REIMBURSEMENT: {
    LIST: '/v1/reimbursements',
    DETAIL: (id) => `/v1/reimbursements/${id}`,
    CREATE: '/v1/reimbursements',
    UPDATE: (id) => `/v1/reimbursements/${id}`,
    SUBMIT: (id) => `/v1/reimbursements/${id}/submit`,
    APPROVE: (id) => `/v1/reimbursements/${id}/approve`,
    REJECT: (id) => `/v1/reimbursements/${id}/reject`,
  },

  // 分析和报告
  ANALYTICS: {
    DASHBOARD: '/v1/analytics/dashboard',
    METRICS: '/v1/analytics/metrics',
    CATEGORIES: '/v1/analytics/categories',
    TRENDS: '/v1/analytics/weekly',
    MONTHLY: '/v1/analytics/monthly',
    RECENT_BILLS: '/v1/analytics/recent-bills',
    REIMBURSEMENT: {
      LIST: '/v1/analytics/reimbursement',
      STATISTICS: '/v1/analytics/reimbursement/statistics',
      TRENDS: '/v1/analytics/reimbursement/trends',
    },
  },

  // 设置相关
  SETTINGS: {
    USER: '/v1/settings',
    CONFIG: '/v1/config',
    NOTIFICATIONS: '/v1/notifications/preferences',
  }
}

export default API_URLS