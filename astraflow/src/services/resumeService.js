// This service is resume related and not part of core AstraFlow functionality
// Import resumeAPI when backend endpoints are available

// 简历服务
export const useResumeService = () => {
  // 生成简历（支持真实API和Mock数据）
  const generateResume = async (resumeData) => {
    try {
      // 暂时使用Mock数据，待后端API实现
      if (import.meta.env.VITE_MOCK_API !== 'false') {
        console.log('使用Mock数据生成简历')
        return await generateMockResume(resumeData)
      }
      throw new Error('Resume generation API not implemented yet')
    } catch (error) {
      if (error.useMock || import.meta.env.VITE_MOCK_API === 'true') {
        console.log('使用Mock数据生成简历')
        return await generateMockResume(resumeData)
      }
      throw error
    }
  }

  // 获取简历建议（支持真实API和Mock数据）
  const getResumeSuggestions = async (resumeData) => {
    try {
      // 暂时使用Mock数据，待后端API实现
      if (import.meta.env.VITE_MOCK_API !== 'false') {
        console.log('使用Mock数据获取建议')
        return await getMockSuggestions(resumeData)
      }
      throw new Error('Resume suggestions API not implemented yet')
    } catch (error) {
      if (error.useMock || import.meta.env.VITE_MOCK_API === 'true') {
        console.log('使用Mock数据获取建议')
        return await getMockSuggestions(resumeData)
      }
      throw error
    }
  }

  // 模拟生成简历
  const generateMockResume = async (resumeData) => {
    // 模拟网络延迟
    await new Promise(resolve => setTimeout(resolve, 1000))

    // 处理技能数据
    const skills = resumeData.skills.split(',').map(skill => skill.trim()).filter(skill => skill)

    return {
      name: resumeData.name || '张三',
      email: resumeData.email || 'zhangsan@example.com',
      phone: resumeData.phone || '138-0000-0000',
      targetPosition: resumeData.targetPosition || '前端开发工程师',
      experience: resumeData.experience || '3年前端开发经验，熟悉Vue.js和React，有丰富的项目经验',
      education: resumeData.education || '计算机科学与技术专业，本科',
      skills: skills.length > 0 ? skills : ['Vue.js', 'React', 'JavaScript', 'HTML/CSS']
    }
  }

  // 模拟获取AI建议
  const getMockSuggestions = async (resumeData) => {
    // 模拟网络延迟
    await new Promise(resolve => setTimeout(resolve, 800))

    return [
      {
        title: '简历优化建议',
        content: '建议在工作经验部分增加具体的数据和成果，例如"提升页面加载速度30%"等量化指标'
      },
      {
        title: '技能匹配度提升',
        content: '根据您意向的岗位，建议补充关于Node.js后端开发的经验描述'
      },
      {
        title: '格式美化建议',
        content: '简历整体结构清晰，建议使用更专业的术语描述项目经验'
      }
    ]
  }

  return {
    generateResume,
    getResumeSuggestions,
    generateMockResume,
    getMockSuggestions
  }
}