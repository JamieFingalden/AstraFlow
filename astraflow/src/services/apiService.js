import axios from 'axios'

// 创建axios实例
const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    // 可以在这里添加token等认证信息
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
api.interceptors.response.use(
  (response) => {
    return response.data
  },
  (error) => {
    console.error('API请求错误:', error)

    // 如果是开发环境，使用mock数据
    if (import.meta.env.DEV && error.code === 'ECONNREFUSED') {
      console.log('后端服务未启动，使用mock数据')
      return Promise.reject({ useMock: true })
    }

    return Promise.reject(error)
  }
)

// 简历相关API
export const resumeAPI = {
  // 生成简历
  generateResume: (data) => api.post('/resume/generate', data),

  // 获取简历建议
  getResumeSuggestions: (data) => api.post('/resume/suggestions', data),

  // 下载简历
  downloadResume: (format = 'pdf') => api.get(`/resume/download?format=${format}`, { responseType: 'blob' })
}

// 岗位分析相关API
export const jobAPI = {
  // 分析岗位匹配度
  analyzeJob: (data) => api.post('/job/analyze', data),

  // 获取岗位推荐
  getJobRecommendations: (data) => api.post('/job/recommendations', data),

  // 获取市场薪资分析
  getSalaryAnalysis: (jobTitle, location) => api.get(`/job/salary?title=${jobTitle}&location=${location}`)
}

// 用户画像相关API
export const profileAPI = {
  // 获取用户画像
  getProfile: (userId) => api.get(`/profile/${userId}`),

  // 更新用户画像
  updateProfile: (data) => api.put('/profile', data),

  // 获取技能评估
  getSkillAssessment: (data) => api.post('/profile/skills/assessment', data),

  // 获取职业发展建议
  getCareerAdvice: (data) => api.post('/profile/career/advice', data)
}

// AI相关API（连接心流平台）
export const aiAPI = {
  // AI文本生成
  generateText: (prompt, type) => api.post('/ai/generate', { prompt, type }),

  // AI简历优化
  optimizeResume: (resumeData, targetJob) => api.post('/ai/resume/optimize', { resume: resumeData, job: targetJob }),

  // AI面试准备
  prepareInterview: (jobData, resumeData) => api.post('/ai/interview/prepare', { job: jobData, resume: resumeData }),

  // AI职业规划
  careerPlanning: (profileData) => api.post('/ai/career/planning', profileData)
}

export default api