  const register = async (userData) => {
    loading.value = true
    error.value = null

    try {
      const response = await registerApi(userData)
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