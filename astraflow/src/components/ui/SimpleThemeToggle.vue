<template>
  <button
    @click="toggleTheme"
    class="theme-toggle-btn"
    :title="`切换到${isDark ? '浅色' : '深色'}模式`"
  >
    <!-- Sun icon for dark mode (shows when in dark) -->
    <svg
      v-if="isDark"
      class="sun-icon"
      fill="currentColor"
      viewBox="0 0 20 20"
    >
      <path
        fill-rule="evenodd"
        d="M10 2a1 1 0 011 1v1a1 1 0 11-2 0V3a1 1 0 011-1zm4 8a4 4 0 11-8 0 4 4 0 018 0zm-.464 4.95l.707.707a1 1 0 001.414-1.414l-.707-.707a1 1 0 00-1.414 1.414zm2.12-10.607a1 1 0 010 1.414l-.706.707a1 1 0 11-1.414-1.414l.707-.707a1 1 0 011.414 0zM17 11a1 1 0 100-2h-1a1 1 0 100 2h1zm-7 4a1 1 0 011 1v1a1 1 0 11-2 0v-1a1 1 0 011-1zM5.05 6.464A1 1 0 106.465 5.05l-.708-.707a1 1 0 00-1.414 1.414l.707.707zm1.414 8.486l-.707.707a1 1 0 01-1.414-1.414l.707-.707a1 1 0 011.414 1.414zM4 11a1 1 0 100-2H3a1 1 0 000 2h1z"
        clip-rule="evenodd"
      />
    </svg>

    <!-- Moon icon for light mode (shows when in light) -->
    <svg
      v-else
      class="moon-icon"
      fill="currentColor"
      viewBox="0 0 20 20"
    >
      <path d="M17.293 13.293A8 8 0 016.707 2.707a8.001 8.001 0 1010.586 10.586z" />
    </svg>
  </button>
</template>

<script setup>
import { useTheme } from '../../composables/useTheme'

const { toggleTheme, isDark } = useTheme()
</script>

<style scoped>
.theme-toggle-btn {
  padding: 0.5rem;
  border-radius: 0.5rem;
  transition: all 0.3s ease;
  background-color: #e5e7eb;
  border: none;
  cursor: pointer;
  transform: scale(1);
}

.theme-toggle-btn:hover {
  transform: scale(1.1);
}

.theme-toggle-btn:active {
  transform: scale(0.95);
}

.theme-toggle-btn:focus {
  outline: none;
  box-shadow: 0 0 0 2px #3b82f6, 0 0 0 4px rgba(59, 130, 246, 0.1);
}

/* 深色主题 */
[data-theme="dark"] .theme-toggle-btn {
  background-color: #374151;
}

[data-theme="dark"] .theme-toggle-btn:hover {
  background-color: #4b5563;
}

/* 图标样式 */
.sun-icon {
  width: 1.25rem;
  height: 1.25rem;
  color: #facc15;
}

.moon-icon {
  width: 1.25rem;
  height: 1.25rem;
  color: #374151;
}

/* 亮色主题下的月亮图标 */
[data-theme="light"] .moon-icon {
  color: #374151;
}
</style>