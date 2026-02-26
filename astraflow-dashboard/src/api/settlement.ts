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

export interface SettlementUser {
  id: number
  username: string
  name?: string
}

export interface SettlementInvoiceItem {
  id: number
  invoice_number: string
  amount: number
  category: string
  status: 'approved' | 'paid'
  updated_at: string
  user?: SettlementUser
}

export interface ApprovedInvoicesPage {
  items: SettlementInvoiceItem[]
  page: number
  size: number
  total: number
  total_pages: number
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
 * 获取已审核通过的待打款列表。
 * @param params 分页参数，page 为页码，size 为每页条数。
 * @returns 待打款分页结果。
 */
export const getApprovedInvoices = async (params: { page: number; size: number }) => {
  const response = await request<ApiResponse<ApprovedInvoicesPage>>({
    url: '/invoices/approved',
    method: 'get',
    params,
  })

  return unwrapApiData<ApprovedInvoicesPage>(response as AxiosOrApiResponse<ApprovedInvoicesPage>)
}

/**
 * 批量确认打款。
 * @param data 需要打款的单据 ID 数组。
 * @returns 无返回数据，成功即表示打款状态更新完成。
 */
export const batchPayInvoices = async (data: { invoice_ids: number[] }) => {
  await request<ApiResponse<null>>({
    url: '/invoices/batch-pay',
    method: 'post',
    data,
  })
}
