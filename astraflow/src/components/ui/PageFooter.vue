<template>
  <footer :class="'page-footer ' + (isDark ? 'dark-theme' : 'light-theme')">
    <div class="footer-container">
      <div class="footer-content">
        <div class="footer-brand">
          <div class="footer-indicator"></div>
          <span class="footer-brand-text">
            ASTRAFLOW
          </span>
        </div>
        <p class="footer-tagline">
          Smart Expense Made Simple ✨
        </p>
        <p :class="'footer-text ' + (isDark ? 'dark-theme' : 'light-theme')">
          © 2025 AstraFlow. All rights reserved.
        </p>
      </div>
    </div>
  </footer>
</template>

<script setup>
import { useTheme } from '../../composables/useTheme'

const { isDark } = useTheme()
</script>

<style scoped>
/* 页脚 */
.page-footer {
  flex-shrink: 0;
  position: relative;
  z-index: 10;
}

.footer-container {
  border-top: 1px solid rgba(59, 130, 246, 0.2);
  background-color: rgba(0, 0, 0, 0.3);
  backdrop-filter: blur(4px);
  margin: 0 auto;
  padding:  1rem;
  text-align: center;
}

.footer-content {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.page-footer.dark-theme {
  background-color: rgba(31, 41, 55, 0.7);
  border-color: rgba(55, 65, 81, 0.5);
}

.page-footer.light-theme {
  background-color: rgba(255, 255, 255, 0.7);
  border-color: rgba(229, 231, 235, 0.5);
}

.footer-brand {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
}

.footer-indicator {
  width: 0.5rem;
  height: 0.5rem;
  background-color: #3b82f6;
  border-radius: 50%;
  animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

.footer-brand-text {
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
  background-image: linear-gradient(to right, #60a5fa, #06b6d4);
  font-size: 1.125rem;
  font-weight: 700;
}

.footer-tagline {
  color: #9ca3af;
  font-size: 0.875rem;
  margin: 0;
}

.footer-text {
  text-align: center;
  font-size: 0.75rem;
  transition: color 0.3s ease;
}

.footer-text.dark-theme {
  color: #6b7280;
}

.footer-text.light-theme {
  color: #6b7280;
  margin: 0;
}

/* 主题支持 */
[data-theme="light"] .footer-container {
  background: rgba(255, 255, 255, 0.95);
  border-top: 1px solid rgba(229, 231, 235, 0.5);
}

[data-theme="light"] .footer-tagline {
  color: #6b7280;
}

[data-theme="light"] .footer-brand-text {
  background-image: linear-gradient(135deg, #3b82f6, #8b5cf6);
}

/* 脉冲动画 */
@keyframes pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.5;
  }
}
</style>