import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import apiService from '@/services/apiService'

export const useUserStore = defineStore('user', () => {
  // State
  const user = ref({
    id: null,
    name: '',
    email: '',
    avatar: '',
    isAuthenticated: false,
    permissions: [],
    role: '',
    tenantId: null,
    tenantName: null
  })

  const loading = ref(false)
  const error = ref(null)

  // Token management
  const accessToken = ref(localStorage.getItem('accessToken') || null)
  const refreshToken = ref(localStorage.getItem('refreshToken') || null)

  // Demo user data for fallback
  const demoUser = {
    id: 1,
    name: '演示用户',
    email: 'demo@astraflow.com',
    avatar: '/default-avatar.png',
    isAuthenticated: true,
    permissions: ['view_bills', 'upload_invoices', 'view_reports', 'manage_users'],
    role: 'admin',
    tenantId: null,
    tenantName: null
  }

  // Getters
  const isAuthenticated = computed(() => user.value.isAuthenticated && !!accessToken.value)
  const isLoggedIn = computed(() => isAuthenticated.value)
  const getUser = computed(() => user.value)
  const getUserRole = computed(() => user.value.role)
  const getUserPermissions = computed(() => user.value.permissions)
  const hasPermission = computed(() => (permission) => {
    return user.value.permissions.includes(permission)
  })
  const isTenantUser = computed(() => !!user.value.tenantId)
  const isPersonalUser = computed(() => user.value.role === 'personal')

  // Actions
  const login = async (email, password, rememberMe = false) => {
    loading.value = true
    error.value = null

    try {
      // Try to login with real API first
      if (email === 'demo@astraflow.com' && password === 'Demo@123') {
        // Demo login - use mock data
        handleSuccessfulLogin(demoUser, null, null, rememberMe)
        return { success: true, user: demoUser }
      }

      const response = await apiService.post('/auth/login', {
        email,
        password,
        remember_me: rememberMe
      })

      const { user: userData, access_token, refresh_token } = response.data

      handleSuccessfulLogin(userData, access_token, refresh_token, rememberMe)

      return { success: true, user: userData }

    } catch (err) {
      console.error('Login error:', err)

      // Fallback to demo account if API fails and it's demo credentials
      if (email === 'demo@astraflow.com' && password === 'Demo@123') {
        handleSuccessfulLogin(demoUser, null, null, rememberMe)
        return { success: true, user: demoUser }
      }

      const errorMessage = err.response?.data?.message || '登录失败，请稍后重试'
      error.value = errorMessage
      throw new Error(errorMessage)

    } finally {
      loading.value = false
    }
  }

  const register = async (userData) => {
    loading.value = true
    error.value = null

    try {
      const response = await apiService.post('/auth/register', userData)
      const { user: newUser, access_token, refresh_token } = response.data

      handleSuccessfulLogin(newUser, access_token, refresh_token, false)

      return { success: true, user: newUser }

    } catch (err) {
      console.error('Registration error:', err)

      const errorMessage = err.response?.data?.message || '注册失败，请稍后重试'
      error.value = errorMessage
      throw new Error(errorMessage)

    } finally {
      loading.value = false
    }
  }

  const logout = async () => {
    try {
      // Call logout API to invalidate tokens
      if (accessToken.value) {
        await apiService.post('/auth/logout')
      }
    } catch (err) {
      console.error('Logout error:', err)
      // Continue with local logout even if API call fails
    } finally {
      // Clear local state
      clearAuthState()
    }
  }

  const refreshTokens = async () => {
    if (!refreshToken.value) {
      throw new Error('No refresh token available')
    }

    try {
      const response = await apiService.post('/auth/refresh', {
        refresh_token: refreshToken.value
      })

      const { access_token, refresh_token: new_refresh_token } = response.data

      accessToken.value = access_token
      localStorage.setItem('accessToken', access_token)

      if (new_refresh_token) {
        refreshToken.value = new_refresh_token
        localStorage.setItem('refreshToken', new_refresh_token)
      }

      return access_token

    } catch (err) {
      console.error('Token refresh error:', err)
      clearAuthState()
      throw err
    }
  }

  const fetchCurrentUser = async () => {
    if (!accessToken.value) {
      throw new Error('No access token available')
    }

    loading.value = true
    error.value = null

    try {
      const response = await apiService.get('/auth/me')
      const userData = response.data

      user.value = {
        ...userData,
        isAuthenticated: true
      }

      return userData

    } catch (err) {
      console.error('Fetch current user error:', err)
      error.value = '获取用户信息失败'
      throw err

    } finally {
      loading.value = false
    }
  }

  const updateUser = (userData) => {
    user.value = { ...user.value, ...userData }
  }

  const hasRole = (role) => {
    return user.value.role === role
  }

  const canAccess = (permission) => {
    if (!user.value.isAuthenticated) return false
    if (user.value.role === 'admin') return true
    return user.value.permissions.includes(permission)
  }

  // Helper functions
  const handleSuccessfulLogin = (userData, accessTokenValue, refreshTokenValue, rememberMe) => {
    user.value = {
      ...userData,
      isAuthenticated: true
    }

    // Store tokens
    if (accessTokenValue) {
      accessToken.value = accessTokenValue
      localStorage.setItem('accessToken', accessTokenValue)

      if (rememberMe && refreshTokenValue) {
        refreshToken.value = refreshTokenValue
        localStorage.setItem('refreshToken', refreshTokenValue)
      } else if (!rememberMe) {
        // Remove refresh token if not remembering
        localStorage.removeItem('refreshToken')
        refreshToken.value = null
      }
    }
  }

  const clearAuthState = () => {
    user.value = {
      id: null,
      name: '',
      email: '',
      avatar: '',
      isAuthenticated: false,
      permissions: [],
      role: '',
      tenantId: null,
      tenantName: null
    }

    accessToken.value = null
    refreshToken.value = null

    // Clear localStorage
    localStorage.removeItem('accessToken')
    localStorage.removeItem('refreshToken')
  }

  // Initialize auth state on store creation
  const initializeAuth = async () => {
    if (accessToken.value) {
      try {
        await fetchCurrentUser()
      } catch (err) {
        // Token might be expired, try to refresh
        if (refreshToken.value) {
          try {
            await refreshTokens()
            await fetchCurrentUser()
          } catch (refreshErr) {
            clearAuthState()
          }
        } else {
          clearAuthState()
        }
      }
    }
  }

  return {
    // State
    user,
    loading,
    error,
    accessToken,
    refreshToken,

    // Getters
    isAuthenticated,
    isLoggedIn,
    getUser,
    getUserRole,
    getUserPermissions,
    hasPermission,
    isTenantUser,
    isPersonalUser,

    // Actions
    login,
    register,
    logout,
    refreshTokens,
    fetchCurrentUser,
    updateUser,
    hasRole,
    canAccess,
    initializeAuth,
    clearAuthState
  }
})