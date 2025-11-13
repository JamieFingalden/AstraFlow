<template>
  <button
    class="theme-toggle"
    @click="toggleTheme"
    :title="`切换到${isDark ? '亮色' : '暗色'}主题`"
  >
    <div class="toggle-container">
      <!-- Sun Icon (Light Theme) -->
      <svg
        v-if="isDark"
        class="icon sun-icon"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
      >
        <circle cx="12" cy="12" r="5"></circle>
        <line x1="12" y1="1" x2="12" y2="3"></line>
        <line x1="12" y1="21" x2="12" y2="23"></line>
        <line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line>
        <line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line>
        <line x1="1" y1="12" x2="3" y2="12"></line>
        <line x1="21" y1="12" x2="23" y2="12"></line>
        <line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line>
        <line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line>
      </svg>

      <!-- Moon Icon (Dark Theme) -->
      <svg
        v-else
        class="icon moon-icon"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
      >
        <path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path>
      </svg>

      <!-- Animated Ring -->
      <div class="toggle-ring"></div>
    </div>
  </button>
</template>

<script setup>
import { computed } from 'vue'
import { useThemeStore } from '../../stores/themeStore'

const themeStore = useThemeStore()

const isDark = computed(() => themeStore.isDark)

const toggleTheme = () => {
  themeStore.toggleTheme()
}
</script>

<style scoped>
.theme-toggle {
  position: relative;
  width: 60px;
  height: 60px;
  border: none;
  border-radius: 50%;
  background: var(--glass-bg);
  backdrop-filter: blur(16px);
  -webkit-backdrop-filter: blur(16px);
  border: 1px solid var(--glass-border);
  cursor: pointer;
  transition: all var(--duration-normal) ease;
  overflow: hidden;
}

.theme-toggle:hover {
  transform: scale(1.05);
  box-shadow: 0 0 30px var(--color-nebula-glow);
}

.theme-toggle:active {
  transform: scale(0.95);
}

.toggle-container {
  position: relative;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.icon {
  width: 24px;
  height: 24px;
  stroke-width: 2;
  transition: all var(--duration-normal) ease;
  z-index: 2;
  position: relative;
}

.sun-icon {
  color: var(--color-stellar-blue);
  animation: rotateSun 20s linear infinite;
}

.moon-icon {
  color: var(--color-nebula-purple);
  animation: floatMoon 6s ease-in-out infinite;
}

.toggle-ring {
  position: absolute;
  width: 100%;
  height: 100%;
  border-radius: 50%;
  border: 2px solid transparent;
  background: linear-gradient(var(--glass-bg), var(--glass-bg)) padding-box,
              var(--gradient-stellar) border-box;
  opacity: 0.3;
  animation: pulseRing 3s ease-in-out infinite;
}

/* Animations */
@keyframes rotateSun {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

@keyframes floatMoon {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-3px); }
}

@keyframes pulseRing {
  0%, 100% {
    transform: scale(1);
    opacity: 0.3;
  }
  50% {
    transform: scale(1.1);
    opacity: 0.6;
  }
}

/* Light theme specific styles */
[data-theme="light"] .theme-toggle {
  background: rgba(255, 255, 255, 0.8);
  border-color: rgba(59, 130, 246, 0.2);
}

[data-theme="light"] .sun-icon {
  color: var(--color-stellar-blue);
  filter: drop-shadow(0 0 8px rgba(59, 130, 246, 0.5));
}

[data-theme="light"] .toggle-ring {
  opacity: 0.5;
}

/* Dark theme specific styles */
[data-theme="dark"] .theme-toggle {
  background: rgba(15, 23, 42, 0.6);
  border-color: rgba(139, 92, 246, 0.3);
}

[data-theme="dark"] .moon-icon {
  color: var(--color-nebula-purple);
  filter: drop-shadow(0 0 8px rgba(139, 92, 246, 0.6));
}
</style>