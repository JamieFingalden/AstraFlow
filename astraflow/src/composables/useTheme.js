import { ref } from 'vue'

const THEME_KEY = 'astraflow-theme'
const DEFAULT_THEME = 'dark'

// Theme state
const theme = ref(localStorage.getItem(THEME_KEY) || DEFAULT_THEME)

// Initialize theme on load
if (!localStorage.getItem(THEME_KEY)) {
  const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches
  theme.value = prefersDark ? 'dark' : 'light'
}

// Apply theme immediately
document.documentElement.classList.remove('dark', 'light')
document.documentElement.classList.add(theme.value)
document.documentElement.setAttribute('data-theme', theme.value)

export function useTheme() {
  const toggleTheme = () => {
    const newTheme = theme.value === 'dark' ? 'light' : 'dark'
    theme.value = newTheme

    // Apply to document
    document.documentElement.classList.remove('dark', 'light')
    document.documentElement.classList.add(newTheme)
    document.documentElement.setAttribute('data-theme', newTheme)

    // Save to localStorage
    localStorage.setItem(THEME_KEY, newTheme)
  }

  const setTheme = (newTheme) => {
    if (newTheme === 'dark' || newTheme === 'light') {
      theme.value = newTheme
      document.documentElement.classList.remove('dark', 'light')
      document.documentElement.classList.add(newTheme)
      document.documentElement.setAttribute('data-theme', newTheme)
      localStorage.setItem(THEME_KEY, newTheme)
    }
  }

  const isDark = () => theme.value === 'dark'
  const isLight = () => theme.value === 'light'

  return {
    theme,
    toggleTheme,
    setTheme,
    isDark,
    isLight
  }
}