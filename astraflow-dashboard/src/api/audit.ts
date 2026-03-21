import request from '../utils/request'

interface ApiResponse<T> {
  code: number
  message: string
  data: T
}

type AxiosOrApiResponse<T> =
  | ApiResponse<T>
  | {
      data: ApiResponse<T>
    }

export interface PendingInvoiceItem {
  id: number
  tenant_id?: number
  user_id: number
  user_name: string
  attachment_id: number
  invoice_number: string
  invoice_date?: string
  vendor: string
  amount: number
  category: string
  description: string
  status: 'pending' | 'approved' | 'rejected'
  pre_audit_status?: 'pre_approved' | 'need_review'
  pre_audit_score?: number
  pre_audit_reasons?: string[]
  created_at: string
  updated_at: string
}

export interface PendingInvoicesResult {
  items: PendingInvoiceItem[]
  page: number
  size: number
  total: number
  total_pages: number
}

export interface InvoiceDetail {
  id: number
  tenant_id?: number
  user_id: number
  user_name: string
  user_username: string
  attachment_id: number
  image_url: string
  invoice_number: string
  invoice_date?: string
  amount: number
  vendor: string
  category: string
  description: string
  status: 'pending' | 'approved' | 'rejected' | 'paid' | 'draft' | 'recognizing' | 'unconfirmed'
  pre_audit_status?: 'pre_approved' | 'need_review'
  pre_audit_score?: number
  pre_audit_reasons?: string[]
  reviewer_id?: number
  review_remarks: string
  paid_at?: string
  created_at: string
  updated_at: string
}

export interface ReviewInvoicePayload {
  action: 'approve' | 'reject'
  final_amount?: number
  remarks?: string
}

export interface ReviewInvoiceResult {
  invoice_id: number
  status: 'approved' | 'rejected'
  amount: number
  reviewer_id?: number
  review_remarks: string
}

/**
 * 统一解析接口返回结构。
 * @param response axios 原始返回或被拦截器处理后的返回对象。
 * @returns 业务 data 数据。
 */
const unwrapApiData = <T>(response: AxiosOrApiResponse<T>): T => {
  if (
    typeof response === 'object' &&
    response !== null &&
    'data' in response &&
    typeof response.data === 'object' &&
    response.data !== null &&
    'code' in response.data
  ) {
    return response.data.data as T
  }
  return (response as ApiResponse<T>).data
}

/**
 * 获取待审核任务池。
 * @param params 分页参数。
 * @returns 待审核分页结果。
 */
export const getPendingInvoices = async (params?: {
  page?: number
  size?: number
  pre_audit_status?: 'pre_approved' | 'need_review' | ''
}) => {
  const response = await request<ApiResponse<PendingInvoicesResult>>({
    url: '/audit/invoices/pending',
    method: 'get',
    params,
  })
  return unwrapApiData<PendingInvoicesResult>(response as AxiosOrApiResponse<PendingInvoicesResult>)
}

/**
 * 获取单据详情。
 * @param id 单据 ID。
 * @returns 单据详情数据。
 */
export const getInvoiceDetail = async (id: string | number) => {
  const response = await request<ApiResponse<InvoiceDetail>>({
    url: `/audit/invoices/${id}`,
    method: 'get'
  })
  return unwrapApiData<InvoiceDetail>(response as AxiosOrApiResponse<InvoiceDetail>)
}

/**
 * 提交审核处理结果。
 * @param id 单据 ID。
 * @param data 审核动作和参数。
 * @returns 审核结果。
 */
export const processAudit = async (id: string | number, data: ReviewInvoicePayload) => {
  const response = await request<ApiResponse<ReviewInvoiceResult>>({
    url: `/audit/invoices/${id}/review`,
    method: 'put',
    data
  })
  return unwrapApiData<ReviewInvoiceResult>(response as AxiosOrApiResponse<ReviewInvoiceResult>)
}
