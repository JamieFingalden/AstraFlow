import { jobAPI } from './apiService'

// 岗位分析服务
export const useJobAnalysisService = () => {
  // 分析岗位（支持真实API和Mock数据）
  const analyzeJob = async (jobData) => {
    try {
      // 尝试调用真实API
      const response = await jobAPI.analyzeJob(jobData)
      return response.data
    } catch (error) {
      if (error.useMock || import.meta.env.VITE_MOCK_API === 'true') {
        console.log('使用Mock数据分析岗位')
        return await analyzeMockJob(jobData)
      }
      throw error
    }
  }

  // 获取岗位推荐
  const getJobRecommendations = async (jobData) => {
    try {
      const response = await jobAPI.getJobRecommendations(jobData)
      return response.data
    } catch (error) {
      if (error.useMock || import.meta.env.VITE_MOCK_API === 'true') {
        console.log('使用Mock数据获取推荐')
        return await getMockRecommendations(jobData)
      }
      throw error
    }
  }

  // 获取薪资分析
  const getSalaryAnalysis = async (jobTitle, location = '北京') => {
    try {
      const response = await jobAPI.getSalaryAnalysis(jobTitle, location)
      return response.data
    } catch (error) {
      if (error.useMock || import.meta.env.VITE_MOCK_API === 'true') {
        console.log('使用Mock数据获取薪资分析')
        return await getMockSalaryAnalysis(jobTitle, location)
      }
      throw error
    }
  }

  // 模拟岗位分析
  const analyzeMockJob = async (jobData) => {
    // 模拟网络延迟
    await new Promise(resolve => setTimeout(resolve, 1500))

    // 模拟匹配度计算
    const skillMatch = Math.floor(Math.random() * 40) + 60 // 60-99之间的随机数
    const experienceMatch = Math.floor(Math.random() * 30) + 50 // 50-79之间的随机数
    const matchScore = Math.floor((skillMatch + experienceMatch) / 2)

    return {
      matchScore,
      skillMatch,
      experienceMatch,
      suggestions: [
        {
          title: '技能匹配优化',
          content: '岗位要求具备React经验，建议在简历中突出相关项目经历'
        },
        {
          title: '经验要求分析',
          content: '该岗位偏好3-5年经验，您的经验符合要求，建议详细描述项目成果'
        },
        {
          title: '薪资谈判建议',
          content: '根据市场行情，该岗位薪资范围在15K-25K，您可以根据经验适当争取'
        }
      ]
    }
  }

  // 模拟岗位推荐
  const getMockRecommendations = async (jobData) => {
    await new Promise(resolve => setTimeout(resolve, 800))

    return [
      {
        title: '前端工程师 - 阿里巴巴',
        company: '阿里巴巴',
        location: '杭州',
        salary: '20-35K',
        matchRate: 85,
        requirements: ['Vue.js', 'React', 'Node.js']
      },
      {
        title: '高级前端开发 - 字节跳动',
        company: '字节跳动',
        location: '北京',
        salary: '25-40K',
        matchRate: 78,
        requirements: ['React', 'TypeScript', 'Webpack']
      },
      {
        title: '全栈工程师 - 腾讯',
        company: '腾讯',
        location: '深圳',
        salary: '18-30K',
        matchRate: 72,
        requirements: ['JavaScript', 'Python', 'Docker']
      }
    ]
  }

  // 模拟薪资分析
  const getMockSalaryAnalysis = async (jobTitle, location) => {
    await new Promise(resolve => setTimeout(resolve, 500))

    return {
      jobTitle,
      location,
      marketData: {
        average: '22K',
        range: '15K-35K',
        median: '20K',
        sampleSize: 1250
      },
      experienceRanges: [
        { level: '1-3年', range: '12K-18K' },
        { level: '3-5年', range: '18K-25K' },
        { level: '5-8年', range: '25K-35K' },
        { level: '8年以上', range: '35K-50K' }
      ]
    }
  }

  return {
    analyzeJob,
    getJobRecommendations,
    getSalaryAnalysis,
    analyzeMockJob,
    getMockRecommendations,
    getMockSalaryAnalysis
  }
}