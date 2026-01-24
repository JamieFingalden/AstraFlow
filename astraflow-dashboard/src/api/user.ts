import request from '../utils/request'

export interface User {
  id: number
  username: string
  name: string
  avatar?: string
  role_key: 'admin' | 'auditor' | 'employee'
  status: 'active' | 'disabled'
  created_at?: string
  last_login?: string
}

export interface CreateUserParams {
  username: string
  name: string
  password?: string
  role_key: string
}

export const getEmployeeList = (params?: any) => {
  return request({
    url: '/users',
    method: 'get',
    params
  })
}

export const inviteEmployee = (data: CreateUserParams) => {
  return request({
    url: '/users',
    method: 'post',
    data
  })
}

export const updateUserStatus = (id: number, status: 'active' | 'disabled') => {
  return request({
    url: `/users/${id}/status`,
    method: 'patch',
    params: { status }
  })
}
