import { request } from '../request'
import { API_URLS } from '../apiUrls'

/**
 * 附件相关API
 */
export const attachmentApi = {
  /**
   * 上传文件
   * @param {File} file - 要上传的文件
   * @param {number} [invoiceId] - 关联的发票ID
   * @returns {Promise}
   */
  uploadFile(data) {
    const formData = new FormData()
    formData.append('file', data.file)
    if (data.invoiceId) {
      formData.append('invoice_id', data.invoiceId)
    }

    return request({
      url: API_URLS.ATTACHMENT.UPLOAD,
      method: 'post',
      data: formData,
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
  },

  /**
   * 根据ID获取附件信息
   * @param {number} id - 附件ID
   * @returns {Promise}
   */
  getAttachmentById(id) {
    return request({
      url: API_URLS.ATTACHMENT.DETAIL(id),
      method: 'get'
    })
  },

  /**
   * 获取用户附件列表
   * @returns {Promise}
   */
  getAttachmentsByUser() {
    return request({
      url: API_URLS.ATTACHMENT.LIST,
      method: 'get'
    })
  },

  /**
   * 获取租户附件列表
   * @param {Object} params - 查询参数 {page, page_size}
   * @returns {Promise}
   */
  getAttachmentsByTenant(params = {}) {
    return request({
      url: API_URLS.ATTACHMENT.LIST_BY_TENANT,
      method: 'get',
      params
    })
  },

  /**
   * 根据发票ID获取附件列表
   * @param {number} invoiceId - 发票ID
   * @returns {Promise}
   */
  getAttachmentsByInvoice(invoiceId) {
    return request({
      url: API_URLS.ATTACHMENT.LIST_BY_INVOICE(invoiceId),
      method: 'get'
    })
  },

  /**
   * 删除附件
   * @param {number} id - 附件ID
   * @returns {Promise}
   */
  deleteAttachment(id) {
    return request({
      url: API_URLS.ATTACHMENT.DELETE(id),
      method: 'delete'
    })
  }
}