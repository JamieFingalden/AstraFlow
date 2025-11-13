import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useThemeStore = defineStore('theme', () => {
  // State
  const currentTheme = ref(localStorage.getItem('theme') || 'light')

  // 初始化主题
  const initializeTheme = () => {
    const savedTheme = localStorage.getItem('theme')
    if (savedTheme) {
      currentTheme.value = savedTheme
    } else {
      // 根据系统偏好设置默认主题
      const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches
      currentTheme.value = prefersDark ? 'dark' : 'light'
    }
    applyTheme(currentTheme.value)
  }

  // 应用主题到DOM
  const applyTheme = (theme) => {
    document.documentElement.classList.remove('dark', 'light')
    document.documentElement.classList.add(theme)
    document.documentElement.setAttribute('data-theme', theme)
  }

  // Getters
  const isDark = computed(() => currentTheme.value === 'dark')
  const isLight = computed(() => currentTheme.value === 'light')

  // Actions
  const setTheme = (theme) => {
    if (theme === 'dark' || theme === 'light') {
      currentTheme.value = theme
      localStorage.setItem('theme', theme)
      applyTheme(theme)
    }
  }

  const toggleTheme = () => {
    const newTheme = currentTheme.value === 'dark' ? 'light' : 'dark'
    setTheme(newTheme)
  }

  return {
    // State
    currentTheme,

    // Getters
    isDark,
    isLight,

    // Actions
    initializeTheme,
    setTheme,
    toggleTheme
  }
})