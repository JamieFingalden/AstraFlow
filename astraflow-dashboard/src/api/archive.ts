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

export interface ArchiveUser {
  id: number
  username: string
  name?: string
}

export interface ArchiveInvoiceItem {
  id: number
  invoice_number: string
  invoice_date?: string
  amount: number
  category: string
  description: string
  status: 'pending' | 'approved' | 'rejected' | 'paid' | 'draft' | 'recognizing' | 'unconfirmed'
  created_at: string
  updated_at: string
  user?: ArchiveUser
}

export interface ArchiveInvoicesPage {
  items: ArchiveInvoiceItem[]
  page: number
  size: number
  total: number
  total_pages: number
}

export interface ArchiveSearchParams {
  page: number
  size: number
  status?: string
  start_date?: string
  end_date?: string
  keyword?: string
}

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
 * 历史归档高级搜索。
 * @param params 分页与筛选条件。
 * @returns 归档分页列表。
 */
export const getArchiveInvoices = async (params: ArchiveSearchParams) => {
  const response = await request<ApiResponse<ArchiveInvoicesPage>>({
    url: '/archive/invoices',
    method: 'get',
    params,
  })

  return unwrapApiData<ArchiveInvoicesPage>(response as AxiosOrApiResponse<ArchiveInvoicesPage>)
}
