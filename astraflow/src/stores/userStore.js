import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import api from '@/services/api'
import { API_URLS } from '@/services/apiUrls'

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
    tenantName: null,
    roleId: null,  // Add roleId to track the specific role
    roleName: null,  // Add roleName for better identification
    tenantRole: null // Add tenant role (admin, normal, personal)
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
    roleId: 1,
    roleName: '系统管理员',
    tenantId: null,
    tenantName: null,
    tenantRole: 'admin'
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
  const getIsTenantUserSimple = computed(() => !!user.value.tenantId)
  const getIsPersonalUserSimple = computed(() => user.value.role === 'personal')

  // Actions
  const login = async (username, password, rememberMe = false) => {
    loading.value = true
    error.value = null

    try {
      // Try to login with real API first
      if (username === 'demo@astraflow.com' && password === 'Demo@123') {
        // Demo login - use mock data
        handleSuccessfulLogin(demoUser, null, null, rememberMe)
        return { success: true, user: demoUser }
      }

      const response = await api.post(API_URLS.AUTH.LOGIN, {
        username,
        password
      })

      const { user: userData, token: access_token, refresh_token } = response.data

      // 转换后端用户数据格式到前端格式
      const frontendUserData = {
        id: userData.id,
        name: userData.username, // 后端使用username作为显示名
        email: userData.email,
        avatar: '', // 后端暂无avatar字段
        permissions: [], // 后端暂无permissions字段，根据role生成
        role: userData.role,
        roleId: userData.role_id || null, // 后端现在提供roleId
        roleName: getRoleDisplayName(userData.role),
        tenantId: userData.tenant_id,
        tenantName: null, // 需要额外获取
        tenantRole: userData.tenant_id ? getTenantRoleFromRole(userData.role) : 'personal',
        isAuthenticated: true
      }

      handleSuccessfulLogin(frontendUserData, access_token, refresh_token, rememberMe)

      return { success: true, user: frontendUserData }

    } catch (err) {
      console.error('Login error:', err)

      // Fallback to demo account if API fails and it's demo credentials
      if (username === 'demo@astraflow.com' && password === 'Demo@123') {
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
      const response = await api.post(API_URLS.AUTH.REGISTER, userData)
      const { user: backendUser, token: access_token, refresh_token } = response.data

      // 转换后端用户数据格式到前端格式
      const frontendUserData = {
        id: backendUser.id,
        name: backendUser.username, // 后端使用username作为显示名
        email: backendUser.email,
        avatar: '', // 后端暂无avatar字段
        permissions: [], // 后端暂无permissions字段，根据role生成
        role: backendUser.role,
        roleId: backendUser.role_id || null, // 后端现在提供roleId
        roleName: getRoleDisplayName(backendUser.role),
        tenantId: backendUser.tenant_id,
        tenantName: null, // 需要额外获取
        tenantRole: backendUser.tenant_id ? getTenantRoleFromRole(backendUser.role) : 'personal',
        isAuthenticated: true
      }

      handleSuccessfulLogin(frontendUserData, access_token, refresh_token, false)

      return { success: true, user: frontendUserData }

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
        await api.post(API_URLS.AUTH.LOGOUT)
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
    try {
      const response = await api.post(API_URLS.AUTH.REFRESH, {}, {
        headers: {
          'Authorization': `Bearer ${refreshToken.value}`
        }
      })

      const { token: access_token } = response.data

      accessToken.value = access_token
      localStorage.setItem('accessToken', access_token)

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
      const response = await api.get(API_URLS.AUTH.ME)
      const backendUser = response.data

      user.value = {
        id: backendUser.id || null,
        name: backendUser.username || '',
        email: backendUser.email || '',
        avatar: '', // 后端暂无avatar字段
        permissions: [], // 后端暂无permissions字段，根据role生成
        role: backendUser.role || '',
        roleId: null, // 后端暂无roleId
        roleName: getRoleDisplayName(backendUser.role),
        tenantId: backendUser.tenant_id || null,
        tenantName: null, // 需要额外获取
        tenantRole: backendUser.tenant_id ? getTenantRoleFromRole(backendUser.role) : 'personal',
        isAuthenticated: true
      }

      return user.value

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

  const hasRoleId = (roleId) => {
    return user.value.roleId === roleId
  }

  const hasTenantRole = (tenantRole) => {
    return user.value.tenantRole === tenantRole
  }

  const isSystemAdmin = () => {
    return user.value.role === 'admin' || user.value.roleId === 1
  }

  const isTenantAdmin = () => {
    return user.value.tenantRole === 'admin' || user.value.role === 'tenant_admin'
  }

  const isTenantUser = () => {
    return user.value.tenantRole === 'normal' || user.value.role === 'tenant_user'
  }

  const isPersonalUser = () => {
    return user.value.role === 'personal' || user.value.tenantRole === 'personal'
  }

  const canAccess = (permission) => {
    if (!user.value.isAuthenticated) return false
    if (user.value.role === 'admin') return true
    if (user.value.tenantRole === 'admin') return true
    return user.value.permissions.includes(permission)
  }

  const canAccessTenantFeature = (permission) => {
    if (!user.value.isAuthenticated) return false
    if (!user.value.tenantId) return false  // Not a tenant user
    if (user.value.tenantRole === 'admin') return true
    return user.value.permissions.includes(permission)
  }

  // Helper functions
  const handleSuccessfulLogin = (userData, accessTokenValue, refreshTokenValue, rememberMe) => {
    user.value = {
      id: userData.id || null,
      name: userData.name || '',
      email: userData.email || '',
      avatar: userData.avatar || '',
      permissions: userData.permissions || [],
      role: userData.role || '',
      roleId: userData.roleId || null,
      roleName: userData.roleName || getRoleDisplayName(userData.role),
      tenantId: userData.tenantId || null,
      tenantName: userData.tenantName || null,
      tenantRole: userData.tenantRole || getTenantRoleFromRole(userData.role),
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

  // 辅助函数：获取角色显示名称
  const getRoleDisplayName = (role) => {
    const roleMap = {
      'admin': '系统管理员',
      'normal': '普通用户',
      'personal': '个人用户',
      'tenant_admin': '企业管理员',
      'tenant_user': '企业用户'
    }
    return roleMap[role] || role
  }

  // 根据角色获取租户角色
  const getTenantRoleFromRole = (role) => {
    if (role === 'admin' || role === 'tenant_admin') return 'admin'
    if (role === 'normal' || role === 'tenant_user') return 'normal'
    if (role === 'personal') return 'personal'
    return 'normal'
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
      roleId: null,
      roleName: null,
      tenantId: null,
      tenantName: null,
      tenantRole: null
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

    // Actions
    login,
    register,
    logout,
    refreshTokens,
    fetchCurrentUser,
    updateUser,
    hasRole,
    hasRoleId,
    hasTenantRole,
    isSystemAdmin,
    isTenantAdmin,
    isTenantUser,
    isPersonalUser,
    canAccess,
    canAccessTenantFeature,
    initializeAuth,
    clearAuthState
  }
})