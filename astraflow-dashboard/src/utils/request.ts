import axios from 'axios'
import type { InternalAxiosRequestConfig, AxiosResponse } from 'axios'
import { useUserStore } from '../stores/userStore'

const service = axios.create({
	baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1',
	timeout: 5000,
})

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
    console.error('Request Error:', error)
    return Promise.reject(error)
  }
)

export default service
