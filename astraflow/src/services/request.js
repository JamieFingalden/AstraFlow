import axios from 'axios'

// 创建axios实例
const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
})

/**
 * 通用请求方法
 * @param {Object} config - 请求配置对象
 * @param {string} config.url - 请求地址
 * @param {string} config.method - 请求方法
 * @param {Object} config.data - 请求数据（用于POST、PUT等）
 * @param {Object} config.params - URL参数（用于GET、DELETE等）
 * @param {Object} config.headers - 请求头
 * @returns {Promise} - 返回请求结果的Promise
 */
const request = async (config) => {
  const response = await api(config)
  return response.data
}

// 请求拦截器 - 添加token
api.interceptors.request.use(
  (config) => {
    // 直接从localStorage获取token，而不是通过useUserStore()
    const token = localStorage.getItem('accessToken')

    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }

    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

export { request }
export default api