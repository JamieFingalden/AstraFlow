<template>
  <div :class="'page-header ' + (isDark ? 'dark-theme' : 'light-theme')">
    <div class="container">
      <div class="header-content">
        <div class="header-left">
          <!-- 返回按钮 -->
          <router-link
            v-if="showBackButton"
            :to="backTo || '/'"
            :class="'back-button ' + (isDark ? 'dark-theme' : 'light-theme')"
            :title="backTitle || '返回上一页'"
          >
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="m15 18-6-6 6-6"/>
            </svg>
            <span class="back-button-text">{{ backText || '返回' }}</span>
          </router-link>

          <div v-if="showBrand" class="brand-text">
            AstraFlow
          </div>
          <h1 :class="'page-title ' + (isDark ? 'dark-theme' : 'light-theme')">
            {{ title }}
          </h1>
        </div>

        <div class="header-right">
          <!-- 主题切换按钮 -->
          <button
            v-if="showThemeToggle"
            @click="toggleTheme"
            :class="'theme-toggle-btn ' + (isDark ? 'dark-theme' : 'light-theme')"
          >
            <SunIcon v-if="isDark" :size="20" />
            <MoonIcon v-else :size="20" />
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useTheme } from '../../composables/useTheme'
import { SunIcon, MoonIcon } from 'lucide-vue-next'

const props = defineProps({
  title: {
    type: String,
    required: true
  },
  showBackButton: {
    type: Boolean,
    default: true
  },
  backTo: {
    type: String,
    default: null
  },
  backTitle: {
    type: String,
    default: '返回上一页'
  },
  backText: {
    type: String,
    default: '返回'
  },
  showBrand: {
    type: Boolean,
    default: true
  },
  showThemeToggle: {
    type: Boolean,
    default: true
  }
})

const router = useRouter()
const { isDark, toggleTheme } = useTheme()
</script>

<style scoped>
/* 头部 */
.page-header {
  position: relative;
  z-index: 10;
  backdrop-filter: blur(12px);
  border-bottom: 1px solid;
  transition: all 0.3s ease;
}

.page-header.dark-theme {
  background-color: rgba(31, 41, 55, 0.7);
  border-color: rgba(55, 65, 81, 0.5);
}

.page-header.light-theme {
  background-color: rgba(255, 255, 255, 0.7);
  border-color: rgba(229, 231, 235, 0.5);
}

.container {
  max-width: 80rem;
  margin: 0 auto;
  padding: 0 1rem;
}

@media (min-width: 640px) {
  .container {
    padding: 0 1.5rem;
  }
}

@media (min-width: 1024px) {
  .container {
    padding: 0 2rem;
  }
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 4rem;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 1rem;
}

/* 返回按钮 */
.back-button {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 0.75rem;
  border-radius: 0.5rem;
  transition: all 0.2s ease;
  text-decoration: none;
  border: none;
  background: transparent;
  cursor: pointer;
}

.back-button.dark-theme {
  color: #d1d5db;
}

.back-button.light-theme {
  color: #4b5563;
}

.back-button:hover.dark-theme {
  background-color: rgba(55, 65, 81, 0.5);
  color: #ffffff;
}

.back-button:hover.light-theme {
  background-color: #f3f4f6;
  color: #1f2937;
}

.back-button-text {
  display: none;
}

@media (min-width: 640px) {
  .back-button-text {
    display: inline;
  }
}

/* 品牌文字 */
.brand-text {
  font-size: 1.5rem;
  font-weight: 700;
  background: linear-gradient(135deg, #06b6d4, #3b82f6);
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
}

/* 页面标题 */
.page-title {
  font-size: 1.25rem;
  font-weight: 600;
  transition: color 0.3s ease;
}

.page-title.dark-theme {
  color: #ffffff;
}

.page-title.light-theme {
  color: #1f2937;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 1rem;
}

/* 主题切换按钮 */
.theme-toggle-btn {
  padding: 0.5rem;
  border-radius: 0.5rem;
  transition: all 0.2s ease;
  border: none;
  cursor: pointer;
}

.theme-toggle-btn.dark-theme {
  background-color: rgba(55, 65, 81, 0.5);
  color: #d1d5db;
}

.theme-toggle-btn.light-theme {
  background-color: #f3f4f6;
  color: #374151;
}

.theme-toggle-btn:hover.dark-theme {
  background-color: rgba(75, 85, 99, 0.5);
}

.theme-toggle-btn:hover.light-theme {
  background-color: #e5e7eb;
}
</style>