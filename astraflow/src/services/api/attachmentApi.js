import request from '../request'

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
      url: '/attachments',
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
      url: `/attachments/${id}`,
      method: 'get'
    })
  },

  /**
   * 获取用户附件列表
   * @returns {Promise}
   */
  getAttachmentsByUser() {
    return request({
      url: '/attachments',
      method: 'get'
    })
  },

  /**
   * 获取租户附件列表
   * @returns {Promise}
   */
  getAttachmentsByTenant() {
    return request({
      url: '/attachments/tenant',
      method: 'get'
    })
  },

  /**
   * 根据发票ID获取附件列表
   * @param {number} invoiceId - 发票ID
   * @returns {Promise}
   */
  getAttachmentsByInvoice(invoiceId) {
    return request({
      url: `/attachments/invoice/${invoiceId}`,
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
      url: `/attachments/${id}`,
      method: 'delete'
    })
  }
}