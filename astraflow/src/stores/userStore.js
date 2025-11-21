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

  
  // 纯状态管理方法
  const setUser = (userData) => {
    user.value = {
      id: userData.id || null,
      name: userData.name || '',
      email: userData.email || '',
      avatar: userData.avatar || '',
      permissions: userData.permissions || [],
      role: userData.role || '',
      roleId: userData.roleId || null,
      roleName: userData.roleName || '',
      tenantId: userData.tenantId || null,
      tenantName: userData.tenantName || null,
      tenantRole: userData.tenantRole || null,
      isAuthenticated: userData.isAuthenticated || false
    }
  }

  const setTokens = (accessTokenValue, refreshTokenValue, rememberMe = false) => {
    if (accessTokenValue) {
      accessToken.value = accessTokenValue
      localStorage.setItem('accessToken', accessTokenValue)

      if (rememberMe && refreshTokenValue) {
        refreshToken.value = refreshTokenValue
        localStorage.setItem('refreshToken', refreshTokenValue)
      } else {
        localStorage.removeItem('refreshToken')
        refreshToken.value = null
      }
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

  
  // 初始化认证状态（从localStorage恢复）
  const initializeAuth = async () => {
    const token = localStorage.getItem('accessToken')
    if (token) {
      // 如果有token，设置为已认证状态
      accessToken.value = token
      // 注意：我们不设置isAuthenticated为true，因为没有验证token是否有效
      // 这个状态应该在获取用户信息后设置
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
    clearAuthState,
    setUser,
    setTokens,
    initializeAuth
  }
})