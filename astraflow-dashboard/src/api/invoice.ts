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
  amount: number
  invoice_date: string // YYYY-MM-DD
  category: string
  description?: string
}) => {
  const formData = new FormData()
  formData.append('file', data.file)
  formData.append('amount', data.amount.toString())
  formData.append('invoice_date', data.invoice_date)
  formData.append('category', data.category)
  if (data.description) {
    formData.append('description', data.description)
  }

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
