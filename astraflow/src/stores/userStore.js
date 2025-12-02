import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { me } from '../services/api/authApi'

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

  const initializeAuth = async () => {
    const token = localStorage.getItem('accessToken')
    const refreshTokenValue = localStorage.getItem('refreshToken')

    if (token) {
      // 如果有token，设置为已认证状态
      accessToken.value = token
      refreshToken.value = refreshTokenValue

      // 检查用户信息是否为空，如果为空则调用API获取用户信息
      if (!user.value.id || !user.value.name) {
        loading.value = true
        try {
          const response = await me()

          if (response?.data) {
            // 检查数据结构，可能是 response.data.user 或直接 response.data
            const backendUser = response.data.user || response.data

            // 转换后端用户数据格式到前端格式
            const frontendUserData = {
              id: backendUser.id,
              name: backendUser.username || backendUser.name,
              email: backendUser.email,
              avatar: backendUser.avatar || '',
              permissions: backendUser.permissions || [],
              role: backendUser.role,
              roleId: backendUser.role_id || null,
              roleName: getRoleDisplayName(backendUser.role),
              tenantId: backendUser.tenant_id,
              tenantName: backendUser.tenant_name || null,
              tenantRole: backendUser.tenant_id ? getTenantRoleFromRole(backendUser.role) : 'personal',
              isAuthenticated: true
            }

            setUser(frontendUserData)
          } else {
            // 如果获取用户信息失败，至少设置基本的认证状态
            user.value.isAuthenticated = true
          }
        } catch (error) {
          console.error('Failed to fetch user info during initialization:', error)
          // 即使获取用户信息失败，也要设置基本的认证状态，因为token存在
          user.value.isAuthenticated = true
          error.value = error.message || 'Failed to fetch user info'
        } finally {
          loading.value = false
        }
      } else {
        // 用户信息已经存在，不需要重新获取
        user.value.isAuthenticated = true
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