import request from '../utils/request'

export interface LoginParams {
  username: string
  password?: string
}

export interface LoginResponse {
  token: string
  user_info: {
    name: string
    avatar?: string
    role_key: string
    [key: string]: any
  }
}

export const login = (data: LoginParams) => {
  return request({
    url: '/auth/login',
    method: 'post',
    data
  })
}

export const logout = () => {
  return request({
    url: '/auth/logout',
    method: 'post'
  })
}
