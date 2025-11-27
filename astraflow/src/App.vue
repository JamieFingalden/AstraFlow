<template>
  <n-config-provider :locale="zhCN" :date-locale="dateZhCN" :theme-overrides="themeOverrides">
    <n-message-provider>
      <n-dialog-provider>
        <n-notification-provider>
          <div id="app">
            <router-view v-slot="{ Component }">
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
  await userStore.initializeAuth()
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
</style>