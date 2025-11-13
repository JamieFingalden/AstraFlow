import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useUserStore = defineStore('user', () => {
  // State
  const user = ref({
    id: null,
    name: '',
    email: '',
    avatar: '',
    isAuthenticated: false,
    permissions: []
  })

  // Mock user data
  const mockUsers = [
    {
      id: 1,
      name: '张三',
      email: 'zhangsan@example.com',
      avatar: '/default-avatar.png',
      permissions: ['view_bills', 'upload_invoices', 'view_reports']
    },
    {
      id: 2,
      name: '李四',
      email: 'lisi@example.com',
      avatar: '/default-avatar.png',
      permissions: ['view_bills', 'view_reports']
    }
  ]

  // Getters
  const isLoggedIn = computed(() => user.value.isAuthenticated)
  const getUser = computed(() => user.value)

  // Actions
  const login = async (email, password) => {
    // Mock login - in real app this would call an API
    const foundUser = mockUsers.find(u => u.email === email)
    if (foundUser) {
      user.value = { ...foundUser, isAuthenticated: true }
      return { success: true, user: foundUser }
    }
    return { success: false, error: '用户不存在或密码错误' }
  }

  const logout = () => {
    user.value = {
      id: null,
      name: '',
      email: '',
      avatar: '',
      isAuthenticated: false,
      permissions: []
    }
  }

  const register = async (userData) => {
    // Mock registration
    const newUser = {
      id: mockUsers.length + 1,
      name: userData.name,
      email: userData.email,
      avatar: '/default-avatar.png',
      permissions: ['view_bills', 'upload_invoices']
    }
    mockUsers.push(newUser)
    user.value = { ...newUser, isAuthenticated: true }
    return { success: true, user: newUser }
  }

  const updateUser = (userData) => {
    user.value = { ...user.value, ...userData }
  }

  return {
    // State
    user,

    // Getters
    isLoggedIn,
    getUser,

    // Actions
    login,
    logout,
    register,
    updateUser
  }
})