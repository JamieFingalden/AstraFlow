import request from '../utils/request'

export interface DashboardStats {
  pendingCount: number
  todayCount: number
  totalAmount: number
  aiAccuracy: number
  aiStatus: 'online' | 'offline'
  rabbitMqStatus: {
    connected: boolean
    queueDepth: number
  }
}

export const getStats = () => {
  return request({
    url: '/dashboard/stats',
    method: 'get'
  })
}
