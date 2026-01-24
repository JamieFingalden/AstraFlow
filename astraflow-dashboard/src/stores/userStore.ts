import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useUserStore = defineStore('user', () => {
  const token = ref<string>(localStorage.getItem('token') || '')
  const userInfo = ref<any>(JSON.parse(localStorage.getItem('userInfo') || 'null'))

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

  // Mock Login Action for UI Development
  async function login(form: any) {
    // Simulate API delay
    await new Promise(resolve => setTimeout(resolve, 500))
    
    // Mock logic for demonstration
    const mockRole = form.username === 'admin' ? 'admin' : 'auditor'
    
    setToken('mock-jwt-token-' + Date.now())
    setUserInfo({
      name: form.username === 'admin' ? 'Administrator' : 'Jane Auditor',
      avatar: 'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png',
      role_key: mockRole,
      role: { key: mockRole, name: mockRole.toUpperCase() }
    })
    
    return true
  }

  return { token, userInfo, role_key, setToken, setUserInfo, logout, login }
})