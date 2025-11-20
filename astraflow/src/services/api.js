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

// 请求拦截器 - 添加token
api.interceptors.request.use(
  (config) => {
    const userStore = useUserStore()
    const token = userStore.accessToken

    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }

    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器 - 处理错误状态码
api.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    // 统一错误处理
    if (error.response?.status === 401) {
      // token过期，跳转到登录页
      const userStore = useUserStore()
      userStore.clearAuthState()
      window.location.href = '/login'
    } else if (error.response?.status === 403) {
      console.error('权限不足')
    } else if (error.response?.status >= 500) {
      console.error('服务器错误')
    }
    return Promise.reject(error)
  }
)

export default api