import axios from 'axios'
import type { InternalAxiosRequestConfig, AxiosResponse } from 'axios'

const service = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8081/api/v1',
  timeout: 5000,
})

service.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    // Add token here when available
    // const token = useUserStore().token
    // if (token) {
    //   config.headers.Authorization = `Bearer ${token}`
    // }
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
