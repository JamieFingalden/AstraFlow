<template>
  <div class="auth-layout">
    <!-- Background with stellar effects -->
    <div class="auth-background">
      <div class="stars"></div>
      <div class="nebula"></div>
      <div class="aurora"></div>
    </div>

    <!-- Main content -->
    <div class="auth-container">
      <!-- Logo/Brand section -->
      <header class="auth-header">
        <div class="brand-logo">
          <h1 class="brand-name">AstraFlow</h1>
        </div>
        <p class="brand-tagline">智慧财务管理，开启星际之旅</p>
      </header>

      <!-- Auth forms wrapper -->
      <main class="auth-content">
        <div class="auth-form-container glass glass-dark">
          <slot />
        </div>
      </main>

      <!-- Footer -->
      <footer class="auth-footer">
        <p class="copyright">© 2024 AstraFlow. All rights reserved.</p>
        <div class="footer-links">
          <a href="#" class="footer-link">隐私政策</a>
          <a href="#" class="footer-link">服务条款</a>
          <a href="#" class="footer-link">帮助中心</a>
        </div>
      </footer>
    </div>

    <!-- Back to Home button -->
    <div class="back-home-wrapper">
      <button
        @click="handleBackToHome"
        class="back-home-btn"
        title="返回首页"
      >
        <svg class="back-home-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m15 18-6-6 6-6"/>
        </svg>
        <span class="back-home-text">首页</span>
      </button>
    </div>

    <!-- Theme toggle -->
    <div class="theme-toggle-wrapper">
      <button
        @click="toggleTheme"
        class="theme-toggle-btn glass"
        :title="isDark ? '切换到浅色主题' : '切换到深色主题'"
      >
        <div class="theme-icon">
          <div v-if="isDark" class="sun-icon">☀️</div>
          <div v-else class="moon-icon">🌙</div>
        </div>
      </button>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useTheme } from '@/composables/useTheme'

const router = useRouter()
const { theme, toggleTheme } = useTheme()

const isDark = computed(() => theme.value === 'dark')

const handleBackToHome = () => {
  router.push('/')
}

// Initialize stellar background animations
onMounted(() => {
  createStellarBackground()
})

const createStellarBackground = () => {
  // Create dynamic starfield
  const starsContainer = document.querySelector('.stars')
  if (starsContainer) {
    for (let i = 0; i < 100; i++) {
      const star = document.createElement('div')
      star.className = 'star'
      star.style.cssText = `
        position: absolute;
        width: ${Math.random() * 3}px;
        height: ${Math.random() * 3}px;
        background: white;
        border-radius: 50%;
        top: ${Math.random() * 100}%;
        left: ${Math.random() * 100}%;
        animation: twinkle ${3 + Math.random() * 4}s ease-in-out infinite;
        animation-delay: ${Math.random() * 3}s;
        opacity: ${Math.random() * 0.8 + 0.2};
      `
      starsContainer.appendChild(star)
    }
  }
}
</script>

<style scoped>
.auth-layout {
  min-height: 100vh;
  position: relative;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  font-family: var(--font-primary);
}

/* Background Effects */
.auth-background {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1;
  background: var(--color-deep-space);
}

/* Light theme background adjustments */
[data-theme="light"] .auth-background {
  background: linear-gradient(135deg, #f0f9ff 0%, #e0f2fe 50%, #dbeafe 100%);
}

.stars {
  position: absolute;
  width: 100%;
  height: 100%;
}

/* Light theme star adjustments */
[data-theme="light"] .stars {
  display: none; /* Hide stars in light mode for better readability */
}

.nebula {
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(
    circle at 20% 50%,
    var(--color-nebula-glow) 0%,
    transparent 50%
  ),
  radial-gradient(
    circle at 80% 80%,
    var(--color-aurora-glow) 0%,
    transparent 50%
  );
  animation: nebulaFloat 20s ease-in-out infinite;
  opacity: var(--nebula-opacity);
}

/* Light theme nebula adjustments */
[data-theme="light"] .nebula {
  background: radial-gradient(
    circle at 20% 50%,
    rgba(56, 189, 248, 0.3) 0%,
    transparent 50%
  ),
  radial-gradient(
    circle at 80% 80%,
    rgba(244, 114, 182, 0.3) 0%,
    transparent 50%
  );
  opacity: 0.5;
}

.aurora {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(
    45deg,
    transparent 30%,
    var(--color-nebula-glow) 50%,
    transparent 70%
  );
  animation: auroraFlow 15s ease-in-out infinite;
  opacity: 0.3;
}

/* Light theme aurora adjustments */
[data-theme="light"] .aurora {
  background: linear-gradient(
    45deg,
    transparent 30%,
    rgba(56, 189, 248, 0.2) 50%,
    transparent 70%
  );
  opacity: 0.2;
}

/* Main Container */
.auth-container {
  width: 100%;
  max-width: 1200px;
  padding: var(--space-lg);
  z-index: 1;
  position: relative;
}

/* Header */
.auth-header {
  text-align: center;
  margin-bottom: var(--space-2xl);
  animation: fadeInDown 0.8s ease-out;
}

.brand-logo {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: var(--space-sm);
}

.brand-name {
  font-family: var(--font-display);
  font-size: 2.5rem;
  font-weight: 700;
  color: var(--color-starlight);
  margin: 0;
  text-shadow: 0 0 30px var(--color-nebula-glow);
}

/* Light theme brand name adjustments */
[data-theme="light"] .brand-name {
  color: var(--color-text-primary);
  text-shadow: 0 0 20px rgba(56, 189, 248, 0.4);
}

.brand-tagline {
  font-size: 1rem;
  color: var(--color-starlight);
  opacity: 0.8;
  margin: 0;
  font-weight: 300;
}

/* Light theme brand tagline adjustments */
[data-theme="light"] .brand-tagline {
  color: var(--color-text-secondary);
  opacity: 0.7;
}

/* Content Area */
.auth-content {
  animation: fadeInUp 0.8s ease-out 0.2s both;
}

.auth-form-container {
  max-width: 500px;
  margin: 0 auto;
  padding: var(--space-2xl);
  border-radius: var(--radius-2xl);
  position: relative;
  overflow: hidden;
  background: var(--glass-bg);
  backdrop-filter: blur(16px);
  -webkit-backdrop-filter: blur(16px);
  border: 1px solid var(--glass-border);
  box-shadow: var(--glass-shadow);
}

/* Light theme form container adjustments */
[data-theme="light"] .auth-form-container {
  background: var(--glass-bg);
  border: 1px solid var(--glass-border);
  box-shadow: var(--glass-shadow);
}

/* Footer */
.auth-footer {
  text-align: center;
  margin-top: var(--space-2xl);
  animation: fadeIn 0.8s ease-out 0.4s both;
}

.copyright {
  color: var(--color-starlight);
  opacity: 0.6;
  font-size: 0.875rem;
  margin-bottom: var(--space-sm);
}

/* Light theme copyright adjustments */
[data-theme="light"] .copyright {
  color: var(--color-text-muted);
}

.footer-links {
  display: flex;
  justify-content: center;
  gap: var(--space-lg);
  flex-wrap: wrap;
}

.footer-link {
  color: var(--color-starlight);
  opacity: 0.6;
  text-decoration: none;
  font-size: 0.875rem;
  transition: all var(--duration-normal) ease;
}

.footer-link:hover {
  opacity: 1;
  color: var(--color-stellar-blue);
  text-shadow: 0 0 10px var(--color-nebula-glow);
}

/* Light theme footer link adjustments */
[data-theme="light"] .footer-link {
  color: var(--color-text-secondary);
}

[data-theme="light"] .footer-link:hover {
  color: var(--color-stellar-blue);
  text-shadow: 0 0 8px rgba(56, 189, 248, 0.4);
}

/* Back to Home Button */
.back-home-wrapper {
  position: fixed;
  top: var(--space-lg);
  left: var(--space-lg);
  z-index: 1000;
}

.back-home-btn {
  display: flex;
  align-items: center;
  gap: var(--space-xs);
  padding: var(--space-sm) var(--space-md);
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: var(--radius-lg);
  color: var(--color-starlight);
  cursor: pointer;
  transition: all var(--duration-normal) ease;
  font-size: 0.875rem;
  font-weight: 500;
  text-decoration: none;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
}

.back-home-btn:hover {
  background: rgba(255, 255, 255, 0.15);
  border-color: rgba(255, 255, 255, 0.3);
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.2);
}

.back-home-btn:active {
  transform: translateY(0);
}

.back-home-icon {
  width: 18px;
  height: 18px;
  flex-shrink: 0;
}

.back-home-text {
  white-space: nowrap;
}

/* Theme Toggle */
.theme-toggle-wrapper {
  position: fixed;
  top: var(--space-lg);
  right: var(--space-lg);
  z-index: 1000;
}

.theme-toggle-btn {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all var(--duration-normal) ease;
  border: 1px solid var(--glass-border);
}

.theme-toggle-btn:hover {
  transform: scale(1.1);
  box-shadow: 0 0 20px var(--color-nebula-glow);
}

.theme-icon {
  font-size: 1.2rem;
}

/* Animations */
@keyframes fadeInDown {
  from {
    opacity: 0;
    transform: translateY(-30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

@keyframes nebulaFloat {
  0%, 100% { transform: translate(0, 0) rotate(0deg); }
  33% { transform: translate(30px, -30px) rotate(120deg); }
  66% { transform: translate(-20px, 20px) rotate(240deg); }
}

@keyframes auroraFlow {
  0%, 100% { transform: translateX(-100%) translateY(0); }
  50% { transform: translateX(100%) translateY(-20px); }
}

@keyframes coreGlow {
  0%, 100% {
    transform: scale(1);
    opacity: 1;
  }
  50% {
    transform: scale(1.1);
    opacity: 0.8;
  }
}

@keyframes ringRotate {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

@keyframes ringPulse {
  0%, 100% {
    transform: scale(1);
    opacity: 0.6;
  }
  50% {
    transform: scale(1.05);
    opacity: 0.3;
  }
}

@keyframes twinkle {
  0%, 100% { opacity: 0.2; }
  50% { opacity: 1; }
}

/* Text selection */
::selection {
  background: var(--color-stellar-blue);
  color: white;
}

/* Light theme text selection */
[data-theme="light"] ::selection {
  background: var(--color-stellar-blue);
  color: white;
}

/* Light theme back home button adjustments */
[data-theme="light"] .back-home-btn {
  background: rgba(0, 0, 0, 0.05);
  border-color: rgba(0, 0, 0, 0.1);
  color: var(--color-text-primary);
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
}

[data-theme="light"] .back-home-btn:hover {
  background: rgba(0, 0, 0, 0.1);
  border-color: rgba(0, 0, 0, 0.15);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

/* Responsive Design */
@media (max-width: 768px) {
  .auth-container {
    padding: var(--space-md);
  }

  .brand-name {
    font-size: 2rem;
  }

  .auth-form-container {
    padding: var(--space-xl);
  }

  .footer-links {
    gap: var(--space-md);
  }

  .back-home-wrapper {
    top: var(--space-md);
    left: var(--space-md);
  }

  .back-home-text {
    display: none; /* Hide text on smaller screens */
  }

  .back-home-btn {
    padding: var(--space-sm);
    border-radius: 50%; /* Make it circular on mobile */
  }

  .theme-toggle-wrapper {
    top: var(--space-md);
    right: var(--space-md);
  }
}

@media (max-width: 480px) {
  .brand-logo {
    flex-direction: column;
    gap: var(--space-sm);
  }

  .brand-name {
    font-size: 1.75rem;
  }

  .brand-tagline {
    font-size: 0.875rem;
  }

  .auth-form-container {
    padding: var(--space-lg);
  }
}
</style>