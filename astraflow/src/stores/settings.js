import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useSettingsStore = defineStore('settings', () => {
  // 当前选中的菜单项
  const activeTab = ref('account')

  // 账户信息
  const accountInfo = ref({
    avatar: '/default-avatar.png',
    username: '张三',
    email: 'zhangsan@example.com',
    company: '科技有限公司'
  })

  // 报销偏好
  const reimbursementPreferences = ref({
    autoCategory: true,
    spendingLimit: 5000,
    currency: 'CNY'
  })

  // AI接口配置
  const aiConfig = ref({
    flaskUrl: 'http://localhost:5000/api/ocr',
    apiKey: ''
  })

  // 系统外观
  const appearance = ref({
    theme: 'light', // light, dark, system
    primaryColor: 'blue' // blue, green, purple, gray
  })

  // 计算属性
  const isAutoCategoryEnabled = computed(() => reimbursementPreferences.value.autoCategory)
  const getSpendingLimit = computed(() => reimbursementPreferences.value.spendingLimit)
  const getCurrentCurrency = computed(() => reimbursementPreferences.value.currency)
  const getFlaskUrl = computed(() => aiConfig.value.flaskUrl)
  const getApiKey = computed(() => aiConfig.value.apiKey)
  const getCurrentTheme = computed(() => appearance.value.theme)
  const getCurrentPrimaryColor = computed(() => appearance.value.primaryColor)

  // 方法
  const setActiveTab = (tab) => {
    activeTab.value = tab
  }

  const updateAccountInfo = (info) => {
    accountInfo.value = { ...accountInfo.value, ...info }
    saveToLocalStorage()
  }

  const updateReimbursementPreferences = (prefs) => {
    reimbursementPreferences.value = { ...reimbursementPreferences.value, ...prefs }
    saveToLocalStorage()
  }

  const updateAiConfig = (config) => {
    aiConfig.value = { ...aiConfig.value, ...config }
    saveToLocalStorage()
  }

  const updateAppearance = (app) => {
    appearance.value = { ...appearance.value, ...app }
    saveToLocalStorage()
  }

  // 保存到localStorage
  const saveToLocalStorage = () => {
    localStorage.setItem('astraflow-settings', JSON.stringify({
      accountInfo: accountInfo.value,
      reimbursementPreferences: reimbursementPreferences.value,
      aiConfig: aiConfig.value,
      appearance: appearance.value
    }))
  }

  // 从localStorage加载
  const loadFromLocalStorage = () => {
    const saved = localStorage.getItem('astraflow-settings')
    if (saved) {
      try {
        const data = JSON.parse(saved)
        accountInfo.value = { ...accountInfo.value, ...data.accountInfo }
        reimbursementPreferences.value = { ...reimbursementPreferences.value, ...data.reimbursementPreferences }
        aiConfig.value = { ...aiConfig.value, ...data.aiConfig }
        appearance.value = { ...appearance.value, ...data.appearance }
      } catch (error) {
        console.error('Failed to load settings from localStorage:', error)
      }
    }
  }

  // 测试AI连接
  const testAiConnection = async () => {
    try {
      const response = await fetch(aiConfig.value.flaskUrl, {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${aiConfig.value.apiKey}`
        }
      })
      return response.ok
    } catch (error) {
      console.error('AI connection test failed:', error)
      return false
    }
  }

  // 初始化时加载设置
  loadFromLocalStorage()

  return {
    // 状态
    activeTab,
    accountInfo,
    reimbursementPreferences,
    aiConfig,
    appearance,

    // 计算属性
    isAutoCategoryEnabled,
    getSpendingLimit,
    getCurrentCurrency,
    getFlaskUrl,
    getApiKey,
    getCurrentTheme,
    getCurrentPrimaryColor,

    // 方法
    setActiveTab,
    updateAccountInfo,
    updateReimbursementPreferences,
    updateAiConfig,
    updateAppearance,
    testAiConnection,
    saveToLocalStorage,
    loadFromLocalStorage
  }
})