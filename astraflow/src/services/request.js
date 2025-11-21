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
  console.log('=== request.js START ===')
  console.log('Request: Config received:', config)
  console.log('Request: Base URL:', api.defaults.baseURL)
  console.log('Request: Full URL will be:', api.defaults.baseURL + config.url)

  try {
    console.log('Request: CALLING AXIOS NOW...')
    const response = await api(config)
    console.log('Request: AXIOS RESPONSE RECEIVED:', response)
    console.log('Request: Response data:', response.data)
    console.log('=== request.js END SUCCESS ===')
    return response.data
  } catch (error) {
    console.error('=== request.js ERROR ===')
    console.error('Request: Error details:', error)
    console.error('Request: Error response:', error.response)
    console.error('Request: Error config:', error.config)
    console.error('=== request.js END ERROR ===')
    // 错误已经在拦截器中处理，这里可以进行额外的错误处理
    throw error
  }
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

// 响应拦截器 - 处理错误状态码
api.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    // 统一错误处理
    if (error.response?.status === 401) {
      console.log('Request: 401 error detected, removing tokens')
      // token过期，清除localStorage中的token
      localStorage.removeItem('accessToken')
      localStorage.removeItem('refreshToken')

      // 只有在不在登录页面时才重定向
      if (!window.location.pathname.includes('/login')) {
        console.log('Request: Redirecting to login page')
        // 使用 router 导航而不是页面刷新
        import('@/router').then(router => {
          router.default.push('/login')
        }).catch(() => {
          // 如果 router 导入失败，回退到页面刷新
          window.location.href = '/login'
        })
      }
    } else if (error.response?.status === 403) {
      console.error('权限不足')
    } else if (error.response?.status >= 500) {
      console.error('服务器错误')
    }
    return Promise.reject(error)
  }
)

export { request }
export default api