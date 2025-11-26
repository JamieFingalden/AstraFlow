<template>
  <n-config-provider :locale="zhCN" :date-locale="dateZhCN">
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
import { onMounted } from 'vue'
import { useUserStore } from './stores/userStore'
import { NConfigProvider, NMessageProvider, NDialogProvider, NNotificationProvider, zhCN, dateZhCN } from 'naive-ui'

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