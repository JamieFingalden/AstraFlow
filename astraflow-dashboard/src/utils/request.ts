import axios from 'axios'
import type { InternalAxiosRequestConfig, AxiosResponse } from 'axios'
import { ElMessage } from 'element-plus'
import router from '../router'
import { useUserStore } from '../stores/userStore'

const service = axios.create({
	baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1',
	timeout: 5000,
})

let isHandlingUnauthorized = false

service.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    // 动态获取 Pinia store
    const userStore = useUserStore()
    if (userStore.token) {
      config.headers.Authorization = `Bearer ${userStore.token}`
    }
    return config
  },
  (error: any) => {
    return Promise.reject(error)
  }
)

service.interceptors.response.use(
  (response: AxiosResponse) => {
    return response.data
  },
  (error: any) => {
    const status = error?.response?.status
    const code = error?.response?.data?.code
    const isUnauthorized = status === 401 || code === 401

    if (isUnauthorized && !isHandlingUnauthorized) {
      isHandlingUnauthorized = true

      const userStore = useUserStore()
      userStore.logout()

      ElMessage.warning('登录已过期，请重新登录')

      const currentPath = router.currentRoute.value.fullPath
      if (currentPath !== '/login') {
        router.push({ path: '/login', query: { redirect: currentPath } })
      }

      setTimeout(() => {
        isHandlingUnauthorized = false
      }, 500)
    }

    console.error('Request Error:', error)
    return Promise.reject(error)
  }
)

export default service
