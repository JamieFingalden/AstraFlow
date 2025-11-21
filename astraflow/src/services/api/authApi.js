import { request } from '../request'
import { API_URLS } from '../apiUrls'

export const login = async (identifier, password) => {
    try {
        // 检测输入是邮箱还是用户名
        const isEmail = /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(identifier)

        const requestData = {
            url: API_URLS.AUTH.LOGIN,
            method: 'POST',
            data: isEmail ? { email: identifier, password, username: '' } : { username: identifier, password, email: '' }
        }

        const response = await request(requestData)
        return response
    } catch (error) {
        throw error
    }
}

export const register = async (userData) => {
    return request({
        url: API_URLS.AUTH.REGISTER,
        method: 'POST',
        data: userData
    })
}

export const logout = async () => {
    return request({
        url: API_URLS.AUTH.LOGOUT,
        method: 'POST'
    })
}

export const refreshToken = async () => {
    return request({
        url: API_URLS.AUTH.REFRESH,
        method: 'POST'
    })
}

export const me = async () => {
    return request({
        url: API_URLS.AUTH.ME,
        method: 'GET',
    })
}