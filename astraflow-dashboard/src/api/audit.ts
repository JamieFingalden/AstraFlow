import request from '../utils/request'

export interface InvoiceItem {
  name: string
  count: string | number
  amount: string | number
  unitPrice?: string | number
  tax?: string | number
  unit?: string
}

export interface Invoice {
  id: number | string
  trace_id: string
  user: {
    name: string
    avatar?: string
  }
  vendor: string
  amount: number
  date: string
  confidence: number
  status: 'pending' | 'approved' | 'rejected'
  file_url: string
  category?: string
  description?: string
  items?: InvoiceItem[]
}

export interface AuditResult {
  status: 'approved' | 'rejected'
  remarks?: string
  items?: InvoiceItem[]
  amount?: number
  vendor?: string
  date?: string
}

export const getPendingInvoices = (params?: any) => {
  return request({
    url: '/invoices',
    method: 'get',
    params: { ...params, status: 'pending' }
  })
}

export const getInvoiceDetail = (id: string | number) => {
  return request({
    url: `/invoices/${id}`,
    method: 'get'
  })
}

export const processAudit = (id: string | number, data: AuditResult) => {
  return request({
    url: `/invoices/${id}/audit`,
    method: 'put',
    data
  })
}
