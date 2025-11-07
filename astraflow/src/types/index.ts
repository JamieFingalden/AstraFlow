// 用户相关类型
export interface User {
  id: string
  email: string
  name: string
  avatar?: string
  createdAt: string
  updatedAt: string
}

// 简历相关类型
export interface Resume {
  id: string
  userId: string
  name: string
  email: string
  phone: string
  targetPosition: string
  experience: string
  education: string
  skills: string[]
  createdAt: string
  updatedAt: string
}

// 岗位相关类型
export interface Job {
  id: string
  title: string
  company: string
  location: string
  description: string
  requirements: string[]
  salary?: string
  postedAt: string
}

// 分析结果类型
export interface AnalysisResult {
  matchScore: number
  skillMatch: number
  experienceMatch: number
  suggestions: Suggestion[]
}

export interface Suggestion {
  title: string
  content: string
  type: 'optimization' | 'suggestion' | 'warning'
}

// 职业画像类型
export interface CareerProfile {
  userId: string
  skills: SkillData[]
  experience: ExperienceData[]
  education: EducationData[]
  overallScore: number
  marketCompetitiveness: number
  growthPotential: number
  recommendations: JobRecommendation[]
}

export interface SkillData {
  name: string
  score: number
  category: string
}

export interface ExperienceData {
  company: string
  position: string
  duration: string
  description: string
}

export interface EducationData {
  school: string
  degree: string
  major: string
  duration: string
}

export interface JobRecommendation {
  title: string
  company: string
  location: string
  salary: string
  matchRate: number
  requirements: string[]
}

// API响应类型
export interface ApiResponse<T> {
  success: boolean
  data: T
  message?: string
  error?: string
}

// 表单类型
export interface ResumeForm {
  name: string
  email: string
  phone: string
  targetPosition: string
  experience: string
  education: string
  skills: string
}

export interface JobAnalysisForm {
  title: string
  description: string
  company: string
}

// 路由类型
export interface RouteMeta {
  title: string
  requiresAuth?: boolean
  icon?: string
}