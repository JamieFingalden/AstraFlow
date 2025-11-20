import axios from 'axios'
import { useUserStore } from '@/stores/userStore'

// 创建axios实例
const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    // 添加JWT token到请求头
    const userStore = useUserStore()
    const token = userStore.accessToken

    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }

    // 添加请求ID用于追踪
    config.headers['X-Request-ID'] = generateRequestId()

    return config
  },
  (error) => {
    console.error('Request interceptor error:', error)
    return Promise.reject(error)
  }
)

// 响应拦截器
api.interceptors.response.use(
  (response) => {
    // AstraFlow API统一响应格式处理
    const { data, code, message } = response.data

    // 如果后端返回的标准格式包含code和message
    if (code !== undefined && message !== undefined) {
      if (code === 200) {
        return { data, message }
      } else {
        // 后端返回错误状态
        const error = new Error(message || '请求失败')
        error.response = {
          ...response,
          data: { code, message }
        }
        return Promise.reject(error)
      }
    }

    // 如果返回的是直接数据，直接返回
    return response.data
  },
  async (error) => {
    const userStore = useUserStore()

    console.error('API请求错误:', error)

    // 处理401未授权错误
    if (error.response?.status === 401) {
      // 尝试刷新token
      if (userStore.refreshToken) {
        try {
          await userStore.refreshTokens()
          // 重试原请求
          const originalRequest = error.config
          originalRequest.headers.Authorization = `Bearer ${userStore.accessToken}`
          return api(originalRequest)
        } catch (refreshError) {
          // 刷新失败，清除认证状态并跳转到登录页
          userStore.clearAuthState()
          window.location.href = '/login'
          return Promise.reject(refreshError)
        }
      } else {
        // 没有refresh token，直接清除状态
        userStore.clearAuthState()
        window.location.href = '/login'
      }
    }

    // 处理403禁止访问错误
    if (error.response?.status === 403) {
      const forbiddenError = new Error('权限不足，无法访问该资源')
      forbiddenError.response = error.response
      return Promise.reject(forbiddenError)
    }

    // 处理404资源不存在错误
    if (error.response?.status === 404) {
      const notFoundError = new Error('请求的资源不存在')
      notFoundError.response = error.response
      return Promise.reject(notFoundError)
    }

    // 处理429请求频率限制错误
    if (error.response?.status === 429) {
      const rateLimitError = new Error('请求过于频繁，请稍后重试')
      rateLimitError.response = error.response
      return Promise.reject(rateLimitError)
    }

    // 处理500服务器错误
    if (error.response?.status >= 500) {
      const serverError = new Error('服务器内部错误，请稍后重试')
      serverError.response = error.response
      return Promise.reject(serverError)
    }

    // 网络错误处理
    if (error.code === 'ECONNREFUSED' || error.code === 'NETWORK_ERROR') {
      // 如果是开发环境且后端未启动，允许使用mock数据
      if (import.meta.env.DEV) {
        console.log('后端服务未启动，某些功能将使用mock数据')
        error.useMock = true
      } else {
        const networkError = new Error('网络连接失败，请检查您的网络设置')
        return Promise.reject(networkError)
      }
    }

    // 超时错误处理
    if (error.code === 'ECONNABORTED') {
      const timeoutError = new Error('请求超时，请稍后重试')
      return Promise.reject(timeoutError)
    }

    return Promise.reject(error)
  }
)

// 生成请求ID的工具函数
function generateRequestId() {
  return `req_${Date.now()}_${Math.random().toString(36).substr(2, 9)}`
}

// Authentication related APIs
export const authAPI = {
  // Login
  login: (credentials) => {
    // 转换字段名以匹配后端API
    const loginData = {
      username: credentials.username || credentials.email,
      password: credentials.password
    }
    return api.post('/auth/login', loginData)
  },

  // Register
  register: (userData) => {
    // 转换字段名以匹配后端API
    const registerData = {
      username: userData.username || userData.email,
      password: userData.password,
      email: userData.email,
      phone: userData.phone || '',
      tenant_id: userData.tenant_id || null
    }
    return api.post('/auth/register', registerData)
  },

  // Logout
  logout: () => api.post('/auth/logout'),

  // Refresh token - 使用Authorization header发送refresh token
  refresh: () => {
    const userStore = useUserStore()
    const refresh_token = userStore.refreshToken

    if (!refresh_token) {
      return Promise.reject(new Error('No refresh token available'))
    }

    return api.post('/auth/refresh', {}, {
      headers: {
        'Authorization': `Bearer ${refresh_token}`
      }
    })
  },

  // Get current user
  getCurrentUser: () => api.get('/auth/me'),

  // Update user profile
  updateProfile: (userData) => api.put('/auth/profile', userData),

  // Change password
  changePassword: (passwordData) => api.put('/auth/password', passwordData),

  // Forgot password
  forgotPassword: (email) => api.post('/auth/forgot-password', { email }),

  // Reset password
  resetPassword: (resetData) => api.post('/auth/reset-password', resetData)
}

// Tenant related APIs
export const tenantAPI = {
  // Get current tenant
  getCurrentTenant: () => api.get('/tenants/current'),

  // Get tenant by ID
  getTenant: (id) => api.get(`/tenants/${id}`),

  // Create tenant
  createTenant: (tenantData) => api.post('/tenants', tenantData),

  // Update tenant
  updateTenant: (id, tenantData) => api.put(`/tenants/${id}`, tenantData),

  // Get tenant users
  getTenantUsers: (id) => api.get(`/tenants/${id}/users`),

  // Add user to tenant
  addUserToTenant: (id, userData) => api.post(`/tenants/${id}/users`, userData),

  // Remove user from tenant
  removeUserFromTenant: (tenantId, userId) => api.delete(`/tenants/${tenantId}/users/${userId}`)
}

// Invoice related APIs
export const invoiceAPI = {
  // Get invoices
  getInvoices: (params = {}) => api.get('/invoices', { params }),

  // Get invoice by ID
  getInvoice: (id) => api.get(`/invoices/${id}`),

  // Create invoice
  createInvoice: (invoiceData) => api.post('/invoices', invoiceData),

  // Update invoice
  updateInvoice: (id, invoiceData) => api.put(`/invoices/${id}`, invoiceData),

  // Delete invoice
  deleteInvoice: (id) => api.delete(`/invoices/${id}`),

  // Upload invoice attachment
  uploadInvoice: (formData) => api.post('/invoices/upload', formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  }),

  // Get OCR results
  getOcrResults: (invoiceId) => api.get(`/invoices/${invoiceId}/ocr`)
}

// OCR related APIs
export const ocrAPI = {
  // Process document
  processDocument: (formData) => api.post('/ocr/process', formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  }),

  // Get OCR result
  getOcrResult: (id) => api.get(`/ocr/${id}`),

  // Get OCR results history
  getOcrHistory: (params = {}) => api.get('/ocr/history', { params })
}

// Reimbursement related APIs
export const reimbursementAPI = {
  // Get reimbursements
  getReimbursements: (params = {}) => api.get('/reimbursements', { params }),

  // Get reimbursement by ID
  getReimbursement: (id) => api.get(`/reimbursements/${id}`),

  // Create reimbursement
  createReimbursement: (reimbursementData) => api.post('/reimbursements', reimbursementData),

  // Update reimbursement
  updateReimbursement: (id, reimbursementData) => api.put(`/reimbursements/${id}`, reimbursementData),

  // Delete reimbursement
  deleteReimbursement: (id) => api.delete(`/reimbursements/${id}`),

  // Submit reimbursement
  submitReimbursement: (id) => api.post(`/reimbursements/${id}/submit`),

  // Approve reimbursement
  approveReimbursement: (id) => api.post(`/reimbursements/${id}/approve`),

  // Reject reimbursement
  rejectReimbursement: (id) => api.post(`/reimbursements/${id}/reject`)
}

// Analytics and reporting APIs
export const analyticsAPI = {
  // Get dashboard data
  getDashboardData: () => api.get('/analytics/dashboard'),

  // Get financial reports
  getFinancialReports: (params) => api.get('/analytics/reports', { params }),

  // Get billing trends
  getBillingTrends: (params) => api.get('/analytics/trends', { params }),

  // Export data
  exportData: (type, params) => api.get(`/analytics/export/${type}`, {
    params,
    responseType: 'blob'
  })
}

// Configuration and settings APIs
export const configAPI = {
  // Get user settings
  getUserSettings: () => api.get('/settings'),

  // Update user settings
  updateUserSettings: (settings) => api.put('/settings', settings),

  // Get system configuration
  getSystemConfig: () => api.get('/config'),

  // Get notification preferences
  getNotificationPreferences: () => api.get('/notifications/preferences'),

  // Update notification preferences
  updateNotificationPreferences: (preferences) => api.put('/notifications/preferences', preferences)
}

export default api