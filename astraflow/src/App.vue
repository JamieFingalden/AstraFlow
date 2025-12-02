<template>
  <n-config-provider :locale="zhCN" :date-locale="dateZhCN" :theme-overrides="themeOverrides">
    <n-message-provider>
      <n-dialog-provider>
        <n-notification-provider>
          <div id="app">
            <!-- 全局加载状态 -->
            <div v-if="isInitializing" class="app-loading">
              <div class="loading-spinner"></div>
              <p>正在初始化应用...</p>
            </div>

            <!-- 正常内容 -->
            <router-view v-else v-slot="{ Component }">
              <transition name="fade" mode="out-in">
                <component :is="Component" />
              </transition>
            </router-view>
          </div>
        </n-notification-provider>
      </n-dialog-provider>
    </n-message-provider>
  </n-config-provider>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { useUserStore } from './stores/userStore'
import { NConfigProvider, NMessageProvider, NDialogProvider, NNotificationProvider, zhCN, dateZhCN } from 'naive-ui'

// 应用初始化状态
const isInitializing = ref(true)

// Theme overrides for stellar colors
const themeOverrides = ref({
  common: {
    // Primary colors for buttons and interactive elements
    primaryColor: '#3b82f6', // --color-stellar-blue
    primaryColorHover: '#60a5fa', // Lighter blue for hover
    primaryColorPressed: '#2563eb', // Darker blue for pressed
    primaryColorSuppl: '#3b82f6',

    // Info colors
    infoColor: '#3b82f6', // --color-stellar-blue
    infoColorHover: '#60a5fa',
    infoColorPressed: '#2563eb',
    infoColorSuppl: '#3b82f6',

    // Success colors
    successColor: '#8b5cf6', // --color-nebula-purple
    successColorHover: '#a78bfa',
    successColorPressed: '#7c3aed',
    successColorSuppl: '#8b5cf6',

    // Warning colors
    warningColor: '#06b6d4', // --color-cosmic-cyan
    warningColorHover: '#22d3ee',
    warningColorPressed: '#0891b2',
    warningColorSuppl: '#06b6d4',

    // Error colors
    errorColor: '#ec4899', // --color-aurora-pink
    errorColorHover: '#f472b6',
    errorColorPressed: '#db2777',
    errorColorSuppl: '#ec4899',

    // Button colors
    buttonColor: '#3b82f6',
    buttonColorHover: '#60a5fa',
    buttonColorPressed: '#2563eb',
    buttonColorFocus: '#3b82f6',

    // Input field colors
    inputColor: '#f9fafb',
    inputColorDisabled: '#f3f4f6',
    inputBorderColor: '#d1d5db',
    inputBorderHoverColor: '#9ca3af',
    inputBorderFocusColor: '#3b82f6',

    // Card and modal background colors
    cardColor: '#ffffff',
    popoverColor: '#ffffff',
    modalColor: '#ffffff',

    // Dark theme colors
    bodyColor: '#f9fafb',
    invertedColor: '#1f2937',

    // Text colors
    textColorBase: '#1f2937',
    textColor1: '#1f2937',
    textColor2: '#374151',
    textColor3: '#6b7280',
    textColorDisabled: '#9ca3af',

    // Border colors
    borderColor: '#e5e7eb',
    borderRadius: '0.375rem'
  }
})

// Initialize authentication on app startup
onMounted(async () => {
  const userStore = useUserStore()

  try {
    // 如果有 token，则初始化用户信息
    if (userStore.accessToken) {
      await userStore.initializeAuth()
    } else {
      // 没有 token，直接结束初始化
      isInitializing.value = false
      return
    }
  } catch (error) {
    console.error('Failed to initialize auth:', error)
  } finally {
    // 无论成功与否，都要结束初始化状态
    isInitializing.value = false
  }
})
</script>

<style>
/* Global styles are in style.css */

/* 页面切换动画 */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* 全局加载状态 */
.app-loading {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(to bottom, #000000, rgba(30, 58, 138, 0.2), #000000);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  z-index: 9999;
}

.app-loading .loading-spinner {
  width: 50px;
  height: 50px;
  border: 4px solid rgba(255, 255, 255, 0.1);
  border-top: 4px solid #3b82f6;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

.app-loading p {
  margin-top: 20px;
  color: #ffffff;
  font-size: 16px;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

/* 亮色主题适配 */
[data-theme="light"] .app-loading {
  background: linear-gradient(to bottom, #ffffff, #f1f5f9, #ffffff);
}

[data-theme="light"] .app-loading .loading-spinner {
  border-color: rgba(0, 0, 0, 0.1);
  border-top-color: #3b82f6;
}

[data-theme="light"] .app-loading p {
  color: #1f2937;
}
</style>