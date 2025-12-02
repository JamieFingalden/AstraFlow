import { request } from '../request'
import { API_URLS } from '../apiUrls'

/**
 * 获取租户下的用户列表
 * @param {string} tenantId - 租户ID
 * @param {Object} params - 查询参数
 * @returns {Promise<Object>} 用户列表响应
 */
export const getTenantUsers = async (tenantId, params = {}) => {
    return request({
        url: `/v1/tenants/${tenantId}/users`,
        method: 'GET',
        params
    })
}

/**
 * 获取用户详情
 * @param {string} userId - 用户ID
 * @returns {Promise<Object>} 用户详情响应
 */
export const getUserById = async (userId) => {
    return request({
        url: `/v1/users/${userId}`,
        method: 'GET'
    })
}

/**
 * 创建租户用户
 * @param {Object} userData - 用户数据
 * @param {string} userData.name - 用户名
 * @param {string} userData.email - 邮箱
 * @param {string} userData.role - 用户角色 (admin, normal)
 * @param {boolean} userData.is_active - 是否激活
 * @param {string} userData.tenant_id - 租户ID
 * @returns {Promise<Object>} 用户创建响应
 */
export const createTenantUser = async (userData) => {
    return request({
        url: '/v1/users',
        method: 'POST',
        data: userData
    })
}

/**
 * 更新租户用户
 * @param {string} userId - 用户ID
 * @param {Object} userData - 用户数据
 * @param {string} [userData.name] - 用户名
 * @param {string} [userData.email] - 邮箱
 * @param {string} [userData.role] - 用户角色
 * @param {boolean} [userData.is_active] - 是否激活
 * @returns {Promise<Object>} 用户更新响应
 */
export const updateTenantUser = async (userId, userData) => {
    return request({
        url: `/v1/users/${userId}`,
        method: 'PUT',
        data: userData
    })
}

/**
 * 删除租户用户
 * @param {string} userId - 用户ID
 * @returns {Promise<Object>} 用户删除响应
 */
export const deleteTenantUser = async (userId) => {
    return request({
        url: `/v1/users/${userId}`,
        method: 'DELETE'
    })
}

/**
 * 邀请用户加入租户
 * @param {Object} inviteData - 邀请数据
 * @param {string} inviteData.email - 邮箱
 * @param {string} inviteData.role - 用户角色 (admin, normal)
 * @param {string} inviteData.tenant_id - 租户ID
 * @returns {Promise<Object>} 邀请响应
 */
export const inviteTenantUser = async (inviteData) => {
    return request({
        url: '/v1/users/invite',
        method: 'POST',
        data: inviteData
    })
}

/**
 * 批量操作用户
 * @param {Object} batchData - 批量操作数据
 * @param {string[]} batchData.user_ids - 用户ID数组
 * @param {string} batchData.action - 操作类型 (activate, deactivate, delete)
 * @returns {Promise<Object>} 批量操作响应
 */
export const batchUpdateUsers = async (batchData) => {
    return request({
        url: '/v1/users/batch',
        method: 'PUT',
        data: batchData
    })
}

/**
 * 获取用户角色列表
 * @returns {Promise<Object>} 角色列表响应
 */
export const getUserRoles = async () => {
    return request({
        url: '/v1/roles',
        method: 'GET'
    })
}