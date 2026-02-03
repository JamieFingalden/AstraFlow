import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { login as loginApi } from '../api/auth'
import type { LoginParams } from '../api/auth'

export const useUserStore = defineStore('user', () => {
  const token = ref<string>(localStorage.getItem('token') || '')
  // Try to parse stored user info safely
  const storedUserInfo = localStorage.getItem('userInfo')
  const userInfo = ref<any>(storedUserInfo ? JSON.parse(storedUserInfo) : null)

  const role_key = computed(() => userInfo.value?.role?.key || userInfo.value?.role_key || 'guest')

  function setToken(newToken: string) {
    token.value = newToken
    localStorage.setItem('token', newToken)
  }

  function setUserInfo(info: any) {
    userInfo.value = info
    localStorage.setItem('userInfo', JSON.stringify(info))
  }

  function logout() {
    token.value = ''
    userInfo.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('userInfo')
  }

  async function login(form: LoginParams) {
    try {
      const res: any = await loginApi(form)
      
      // Handle response structure variations (direct object vs wrapped in data)
      const data = res.data || res
      
      if (data.token) {
        setToken(data.token)
        
        // Handle user_info (snake_case from API vs potential camelCase in frontend)
        const user = data.user_info || data.userInfo || {}
        setUserInfo(user)
        
        return data
      } else {
        throw new Error('Token not found in response')
      }
    } catch (error) {
      throw error
    }
  }

  return { token, userInfo, role_key, setToken, setUserInfo, logout, login }
})
