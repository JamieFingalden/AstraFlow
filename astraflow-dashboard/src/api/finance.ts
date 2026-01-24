import request from '../utils/request'

export interface Invoice {
  id: number | string
  invoice_no: string
  employee_name: string
  amount: number
  approved_date?: string
  date?: string
  status: 'pending' | 'approved' | 'paid' | 'rejected'
}

export const getApprovedInvoices = (params?: any) => {
  return request({
    url: '/invoices',
    method: 'get',
    params: { ...params, status: 'approved' }
  })
}

export const markBatchPaid = (ids: (number | string)[]) => {
  return request({
    url: '/invoices/batch-pay',
    method: 'post',
    data: { ids }
  })
}

export const getAllInvoices = (params?: any) => {
  return request({
    url: '/invoices',
    method: 'get',
    params
  })
}

export const exportInvoices = (params?: any) => {
  return request({
    url: '/invoices/export',
    method: 'get',
    params,
    responseType: 'blob'
  })
}
