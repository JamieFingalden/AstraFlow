import { request } from '../request'
import { API_URLS } from '../apiUrls'

/**
 * 获取仪表板数据
 * @param {Object} params - 查询参数
 * @param {string} params.start_date - 开始日期 (格式: YYYY-MM-DD)
 * @param {string} params.end_date - 结束日期 (格式: YYYY-MM-DD)
 * @returns {Promise<Object>} 仪表板数据响应
 */
export const getDashboardData = async (params = {}) => {
    return request({
        url: API_URLS.ANALYTICS.DASHBOARD,
        method: 'GET',
        params
    })
}

/**
 * 获取关键指标数据
 * @param {Object} params - 查询参数
 * @param {string} params.start_date - 开始日期 (格式: YYYY-MM-DD)
 * @param {string} params.end_date - 结束日期 (格式: YYYY-MM-DD)
 * @returns {Promise<Object>} 指标数据响应
 */
export const getMetrics = async (params = {}) => {
    return request({
        url: API_URLS.ANALYTICS.METRICS,
        method: 'GET',
        params
    })
}

/**
 * 获取支出类别数据
 * @param {Object} params - 查询参数
 * @param {string} params.start_date - 开始日期 (格式: YYYY-MM-DD)
 * @param {string} params.end_date - 结束日期 (格式: YYYY-MM-DD)
 * @returns {Promise<Object>} 支出类别数据响应
 */
export const getExpenseCategories = async (params = {}) => {
    return request({
        url: API_URLS.ANALYTICS.CATEGORIES,
        method: 'GET',
        params
    })
}

/**
 * 获取每周支出趋势数据
 * @param {Object} params - 查询参数
 * @param {string} params.start_date - 开始日期 (格式: YYYY-MM-DD)
 * @param {string} params.end_date - 结束日期 (格式: YYYY-MM-DD)
 * @returns {Promise<Object>} 每周支出趋势响应
 */
export const getWeeklyExpenses = async (params = {}) => {
    return request({
        url: API_URLS.ANALYTICS.TRENDS,
        method: 'GET',
        params
    })
}

/**
 * 获取近期账单数据
 * @param {Object} params - 查询参数
 * @param {string} params.start_date - 开始日期 (格式: YYYY-MM-DD)
 * @param {string} params.end_date - 结束日期 (格式: YYYY-MM-DD)
 * @param {number} params.limit - 限制返回数量，默认10
 * @returns {Promise<Object>} 近期账单数据响应
 */
export const getRecentBills = async (params = {}) => {
    return request({
        url: API_URLS.ANALYTICS.RECENT_BILLS,
        method: 'GET',
        params
    })
}

/**
 * 获取报销统计数据
 * @param {Object} params - 查询参数
 * @param {string} params.start_date - 开始日期 (格式: YYYY-MM-DD)
 * @param {string} params.end_date - 结束日期 (格式: YYYY-MM-DD)
 * @returns {Promise<Object>} 报销统计数据响应
 */
export const getReimbursementStatisticsData = async (params = {}) => {
    return request({
        url: API_URLS.ANALYTICS.REIMBURSEMENT.LIST,
        method: 'GET',
        params
    })
}

/**
 * 获取报销统计详情
 * @param {Object} params - 查询参数
 * @param {string} params.start_date - 开始日期 (格式: YYYY-MM-DD)
 * @param {string} params.end_date - 结束日期 (格式: YYYY-MM-DD)
 * @returns {Promise<Object>} 报销统计详情响应
 */
export const getReimbursementStatistics = async (params = {}) => {
    return request({
        url: API_URLS.ANALYTICS.REIMBURSEMENT.STATISTICS,
        method: 'GET',
        params
    })
}

/**
 * 获取报销月度趋势数据
 * @param {Object} params - 查询参数
 * @param {string} params.start_date - 开始日期 (格式: YYYY-MM-DD)
 * @param {string} params.end_date - 结束日期 (格式: YYYY-MM-DD)
 * @returns {Promise<Object>} 报销月度趋势响应
 */
export const getReimbursementMonthlyTrends = async (params = {}) => {
    return request({
        url: API_URLS.ANALYTICS.REIMBURSEMENT.TRENDS,
        method: 'GET',
        params
    })
}