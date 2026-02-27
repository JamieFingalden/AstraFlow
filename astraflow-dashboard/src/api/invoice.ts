import request from '../utils/request'

/**
 * AI 智能识别流上传
 * @param file 图片文件
 */
export const uploadInvoiceForOCR = (file: File) => {
  const formData = new FormData()
  formData.append('file', file)

  return request({
    url: '/invoices/upload-ocr',
    method: 'post',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}

/**
 * 手动提交流
 * @param data 包含文件和表单数据的对象
 */
export const uploadInvoiceManually = (data: {
  file: File
  invoice_number: string
  amount: number
  vendor: string
  invoice_date: string // YYYY-MM-DD
  payment_method: string
  category: string
  description: string
}) => {
  const formData = new FormData()
  formData.append('file', data.file)
  formData.append('invoice_number', data.invoice_number)
  formData.append('amount', data.amount.toString())
  formData.append('vendor', data.vendor)
  formData.append('invoice_date', data.invoice_date)
  formData.append('payment_method', data.payment_method)
  formData.append('category', data.category)
  formData.append('description', data.description)

  return request({
    url: '/invoices/upload-manual',
    method: 'post',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}

/**
 * 获取当前用户的发票列表
 * @param params 查询参数，例如 { status: 'recognizing', page: 1, pageSize: 10 }
 */
export const getMyInvoices = (params: any) => {
  return request({
    url: '/invoices/my-invoices',
    method: 'get',
    params,
  });
};

/**
 * 获取单据详情
 * @param id 发票ID
 */
export const getInvoiceDetail = (id: number) => {
  return request({
    url: `/invoices/${id}`,
    method: 'get',
  });
};

/**
 * 员工确认动作
 * @param id 发票ID
 * @param data 确认的数据
 */
export const confirmInvoice = (id: number, data: any) => {
  return request({
    url: `/invoices/${id}/confirm`,
    method: 'put',
    data,
  });
};

/**
 * 正式发布
 * @param ids 发票ID数组
 */
export const publishInvoices = (ids: number[]) => {
  return request({
    url: `/invoices/publish`,
    method: 'post',
    data: { ids },
  });
};

/**
 * 更新发票信息
 * @param id 发票ID
 * @param data 更新的数据
 */
export const updateInvoice = (id: number, data: any) => {
  return request({
    url: `/invoices/${id}`,
    method: 'put',
    data,
  });
};

/**
 * 删除发票（仅待确认和待发布）
 * @param id 发票ID
 */
export const deleteInvoice = (id: number) => {
  return request({
    url: `/invoices/${id}`,
    method: 'delete',
  });
};
