import { request } from '../request'
import { API_URLS } from '../apiUrls'

/**
 * 获取发票列表（支持分页）
 * @param {Object} params - 查询参数
 * @param {number} params.page - 页码，默认1
 * @param {number} params.page_size - 每页数量，默认10
 * @returns {Promise<Object>} 发票列表响应
 */
export const getInvoices = async (params = {}) => {
    return request({
        url: API_URLS.INVOICE.LIST,
        method: 'GET',
        params
    })
}

/**
 * 根据用户ID获取发票列表
 * @param {Object} params - 查询参数
 * @param {number} params.page - 页码，默认1
 * @param {number} params.page_size - 每页数量，默认10
 * @returns {Promise<Object>} 用户发票列表响应
 */
export const getInvoicesByUser = async (params = {}) => {
    return request({
        url: API_URLS.INVOICE.BY_USER,
        method: 'GET',
        params
    })
}

/**
 * 根据租户ID获取发票列表
 * @param {Object} params - 查询参数
 * @param {number} params.page - 页码，默认1
 * @param {number} params.page_size - 每页数量，默认10
 * @returns {Promise<Object>} 租户发票列表响应
 */
export const getInvoicesByTenant = async (params = {}) => {
    return request({
        url: API_URLS.INVOICE.BY_TENANT,
        method: 'GET',
        params
    })
}

/**
 * 根据ID获取单个发票详情
 * @param {number} id - 发票ID
 * @returns {Promise<Object>} 发票详情响应
 */
export const getInvoiceById = async (id) => {
    return request({
        url: API_URLS.INVOICE.DETAIL(id),
        method: 'GET'
    })
}

/**
 * 创建新发票
 * @param {Object} invoiceData - 发票数据
 * @param {string} invoiceData.invoice_number - 发票号码（必填）
 * @param {string} invoiceData.invoice_date - 发票日期（必填）
 * @param {number} invoiceData.amount - 金额（必填）
 * @param {string} invoiceData.vendor - 供应商/商户名称（必填）
 * @param {string} [invoiceData.image_url] - 发票图片URL（选填）
 * @param {string} [invoiceData.description] - 发票描述/备注（选填）
 * @param {string} [invoiceData.taxId] - 税号（选填）
 * @param {string} invoiceData.category - 分类（必填）
 * @param {string} [invoiceData.payment_source] - 支付来源（选填）
 * @param {string} [invoiceData.status] - 状态（选填，默认pending）
 * @returns {Promise<Object>} 创建发票响应
 */
export const createInvoice = async (invoiceData) => {
    return request({
        url: API_URLS.INVOICE.CREATE,
        method: 'POST',
        data: invoiceData
    })
}

/**
 * 更新发票信息
 * @param {number} id - 发票ID
 * @param {Object} invoiceData - 更新的发票数据
 * @param {string} [invoiceData.invoice_number] - 发票号码
 * @param {string} [invoiceData.invoice_date] - 发票日期
 * @param {number} [invoiceData.amount] - 金额
 * @param {string} [invoiceData.vendor] - 供应商/商户名称
 * @param {string} [invoiceData.image_url] - 发票图片URL（选填）
 * @param {string} [invoiceData.description] - 发票描述/备注（选填）
 * @param {string} [invoiceData.taxId] - 税号（选填）
 * @param {string} [invoiceData.category] - 分类
 * @param {string} [invoiceData.payment_source] - 支付来源（选填）
 * @param {string} [invoiceData.status] - 状态（选填）
 * @returns {Promise<Object>} 更新发票响应
 */
export const updateInvoice = async (id, invoiceData) => {
    return request({
        url: API_URLS.INVOICE.UPDATE(id),
        method: 'PUT',
        data: invoiceData
    })
}

/**
 * 删除发票
 * @param {number} id - 发票ID
 * @returns {Promise<Object>} 删除发票响应
 */
export const deleteInvoice = async (id) => {
    return request({
        url: API_URLS.INVOICE.DELETE(id),
        method: 'DELETE'
    })
}

/**
 * 上传发票文件（OCR预留）
 * @param {FormData} formData - 包含文件和表单数据的FormData
 * @returns {Promise<Object>} 上传响应
 */
export const uploadInvoice = async (formData) => {
    return request({
        url: API_URLS.INVOICE.UPLOAD,
        method: 'POST',
        data: formData,
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    })
}