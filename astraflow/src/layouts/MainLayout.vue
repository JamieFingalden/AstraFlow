<template>
  <div class="app-container">
    <!-- 星空背景 -->
    <div class="background-container">
      <div class="stars-container">
        <div class="stars"></div>
        <div class="stars2"></div>
      </div>
    </div>

    <!-- 导航栏 -->
    <nav class="navbar">
      <div class="brand">
        <div class="brand-indicator"></div>
        <span class="brand-text">
          AstraFlow 星流
        </span>
      </div>

      <!-- 桌面端导航 -->
      <div class="nav-links desktop-nav">
        <router-link
          to="/"
          class="nav-link"
          :class="{ 'nav-link-active': $route.path === '/' }"
        >
          首页
        </router-link>
        <router-link
          to="/upload"
          class="nav-link"
          :class="{ 'nav-link-active': $route.path === '/upload' }"
        >
          发票上传
        </router-link>
        <router-link
          to="/bills"
          class="nav-link"
          :class="{ 'nav-link-active': $route.path === '/bills' }"
        >
          账单管理
        </router-link>
        <router-link
          to="/visualization"
          class="nav-link"
          :class="{ 'nav-link-active': $route.path === '/visualization' }"
        >
          可视化
        </router-link>
        <router-link
          to="/ai-result"
          class="nav-link"
          :class="{ 'nav-link-active': $route.path === '/ai-result' }"
        >
          AI识别
        </router-link>
        <router-link
          to="/statistics"
          class="nav-link"
          :class="{ 'nav-link-active': $route.path === '/statistics' }"
        >
          报销统计
        </router-link>
        <router-link
          to="/settings"
          class="nav-link"
          :class="{ 'nav-link-active': $route.path === '/settings' }"
        >
          设置中心
        </router-link>
      </div>

      <!-- 移动端汉堡菜单按钮 -->
      <button
        class="mobile-menu-button"
        @click="toggleMobileMenu"
      >
        <svg class="hamburger-icon" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path :d="mobileMenuOpen ? 'M6 18L18 6M6 6l12 12' : 'M4 6h16M4 12h16M4 18h16'" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
        </svg>
      </button>

      <!-- 右侧操作区 -->
      <div class="nav-actions desktop-nav">
        <button class="login-button">
          登录
        </button>
        <button class="register-button">
          注册
        </button>
        <!-- 主题切换按钮 -->
        <ThemeToggle />
      </div>
    </nav>

    <!-- 移动端导航菜单 -->
    <div
      v-if="mobileMenuOpen"
      class="mobile-menu-overlay"
      @click="closeMobileMenu"
    >
      <div class="mobile-menu-content" @click.stop>
        <router-link
          to="/"
          class="mobile-nav-link"
          :class="{ 'mobile-nav-link-active': $route.path === '/' }"
          @click="closeMobileMenu"
        >
          首页
        </router-link>
        <router-link
          to="/upload"
          class="mobile-nav-link"
          :class="{ 'mobile-nav-link-active': $route.path === '/upload' }"
          @click="closeMobileMenu"
        >
          发票上传
        </router-link>
        <router-link
          to="/bills"
          class="mobile-nav-link"
          :class="{ 'mobile-nav-link-active': $route.path === '/bills' }"
          @click="closeMobileMenu"
        >
          账单管理
        </router-link>
        <router-link
          to="/visualization"
          class="mobile-nav-link"
          :class="{ 'mobile-nav-link-active': $route.path === '/visualization' }"
          @click="closeMobileMenu"
        >
          可视化
        </router-link>
        <router-link
          to="/ai-result"
          class="mobile-nav-link"
          :class="{ 'mobile-nav-link-active': $route.path === '/ai-result' }"
          @click="closeMobileMenu"
        >
          AI识别
        </router-link>
        <router-link
          to="/statistics"
          class="mobile-nav-link"
          :class="{ 'mobile-nav-link-active': $route.path === '/statistics' }"
          @click="closeMobileMenu"
        >
          报销统计
        </router-link>
        <router-link
          to="/settings"
          class="mobile-nav-link"
          :class="{ 'mobile-nav-link-active': $route.path === '/settings' }"
          @click="closeMobileMenu"
        >
          设置中心
        </router-link>
        <div class="mobile-actions">
          <button class="mobile-login-button">
            登录
          </button>
          <button class="mobile-register-button">
            注册
          </button>
          <!-- 移动端主题切换按钮 -->
          <ThemeToggle />
        </div>
      </div>
    </div>

    <!-- 主容器 -->
    <div class="main-container">
      <main class="main-content">
        <slot></slot>
      </main>

      <!-- 页脚 -->
      <footer class="footer">
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
            <p class="footer-copyright">
              © 2025 AstraFlow. All rights reserved.
            </p>
          </div>
        </div>
      </footer>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import ThemeToggle from '../components/ui/ThemeToggle.vue'

// 移动端菜单状态
const mobileMenuOpen = ref(false)

// 切换移动端菜单
const toggleMobileMenu = () => {
  mobileMenuOpen.value = !mobileMenuOpen.value
}

// 关闭移动端菜单
const closeMobileMenu = () => {
  mobileMenuOpen.value = false
}
</script>

<style scoped>
/* 主容器 */
.app-container {
  min-height: 100vh;
  background-color: #000000;
  color: #ffffff;
  position: relative;
  overflow: hidden;
}

/* 背景容器 */
.background-container {
  position: fixed;
  inset: 0;
  background: linear-gradient(to bottom, #000000, rgba(30, 58, 138, 0.2), #000000);
}

.stars-container {
  position: absolute;
  inset: 0;
}

/* 星空动画 */
.stars, .stars2 {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}

.stars {
  background: transparent url('data:image/svg+xml,%3Csvg width="100" height="100" xmlns="http://www.w3.org/2000/svg"%3E%3Ccircle cx="50" cy="50" r="0.5" fill="white"/%3E%3C/svg%3E') repeat;
  animation: animStar 100s linear infinite;
}

.stars2 {
  background: transparent url('data:image/svg+xml,%3Csvg width="100" height="100" xmlns="http://www.w3.org/2000/svg"%3E%3Ccircle cx="50" cy="50" r="0.8" fill="white" opacity="0.5"/%3E%3C/svg%3E') repeat;
  animation: animStar 150s linear infinite reverse;
}

@keyframes animStar {
  from { transform: translateY(0px); }
  to { transform: translateY(-2000px); }
}

/* 导航栏 */
.navbar {
  position: relative;
  z-index: 50;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1.5rem;
  backdrop-filter: blur(12px);
  background: linear-gradient(to right, rgba(0, 0, 0, 0.2), rgba(30, 58, 138, 0.2), rgba(0, 0, 0, 0.2));
  border-radius: 1rem;
  margin: 1rem 1.5rem;
  transition: all 0.3s ease-in-out;
  min-height: 3rem;
}

/* On mobile, ensure proper flex sizing to prevent wrapping */
@media (max-width: 767px) {
  .navbar {
    padding-left: 1rem;
    padding-right: 1rem;
  }

  .nav-links {
    display: none;
  }

  .brand {
    flex: 1;
  }

  .mobile-menu-button {
    flex: 0;
  }

  .nav-actions {
    display: none;
  }
}

/* 品牌 */
.brand {
  font-size: 1.5rem;
  font-weight: 700;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.brand-indicator {
  width: 0.75rem;
  height: 0.75rem;
  background-color: #3b82f6;
  border-radius: 50%;
  animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
  box-shadow: 0 0 25px rgba(59, 130, 246, 0.5);
}

.brand-text {
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
  background-image: linear-gradient(to right, #60a5fa, #a855f7, #ec4899);
}

/* 桌面端导航 */
.desktop-nav {
  display: none;
}

@media (min-width: 768px) {
  .desktop-nav {
    display: flex;
  }
}

.nav-links {
  display: flex;
  align-items: center;
  gap: 1.5rem;
  flex-wrap: nowrap;
  overflow-x: auto;
  scrollbar-width: none; /* Firefox */
  -ms-overflow-style: none;  /* IE 10+ */
}

.nav-links::-webkit-scrollbar {
  display: none; /* Chrome, Safari, Opera */
}

@media (min-width: 1024px) {
  .nav-links {
    gap: 2rem;
  }
}

.nav-link {
  color: #d1d5db;
  transition: all 0.3s ease;
  padding: 0.5rem 0.75rem;
  border-radius: 0.5rem;
  text-decoration: none;
}

.nav-link:hover {
  color: #ffffff;
}

.nav-link-active {
  color: #60a5fa;
  font-weight: 500;
}

/* 移动端菜单按钮 */
.mobile-menu-button {
  display: block;
  color: #d1d5db;
  background: none;
  border: none;
  cursor: pointer;
}

@media (min-width: 768px) {
  .mobile-menu-button {
    display: none;
  }
}

.hamburger-icon {
  width: 1.5rem;
  height: 1.5rem;
}

/* 右侧操作区 */
.nav-actions {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.login-button {
  display: none;
  padding: 0.5rem 1rem;
  color: #d1d5db;
  font-weight: 500;
  background: none;
  border: none;
  cursor: pointer;
  transition: color 0.3s ease;
}

@media (min-width: 768px) {
  .login-button {
    display: inline-flex;
  }
}

.login-button:hover {
  color: #ffffff;
}

.register-button {
  display: inline-flex;
  padding: 0.625rem 1.25rem;
  background: linear-gradient(to right, #3b82f6, #8b5cf6);
  color: #ffffff;
  font-weight: 500;
  border: none;
  border-radius: 0.5rem;
  cursor: pointer;
  transition: all 0.3s ease;
  transform: scale(1);
}

.register-button:hover {
  background: linear-gradient(to right, #2563eb, #7c3aed);
  transform: scale(1.05);
  box-shadow: 0 0 50px -12px rgba(59, 130, 246, 0.25);
}

/* 移动端导航菜单 */
.mobile-menu-overlay {
  display: block;
  position: fixed;
  inset: 0;
  z-index: 40;
  background-color: rgba(0, 0, 0, 0.9);
  backdrop-filter: blur(12px);
}

@media (min-width: 768px) {
  .mobile-menu-overlay {
    display: none;
  }
}

.mobile-menu-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  gap: 2rem;
}

.mobile-nav-link {
  font-size: 1.5rem;
  color: #d1d5db;
  transition: all 0.3s ease;
  padding: 0.5rem 1rem;
  border-radius: 0.5rem;
  text-decoration: none;
}

.mobile-nav-link:hover {
  color: #ffffff;
}

.mobile-nav-link-active {
  color: #60a5fa;
  font-weight: 500;
}

.mobile-actions {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-top: 2rem;
}

.mobile-login-button {
  padding: 0.5rem 1rem;
  color: #d1d5db;
  font-weight: 500;
  background: none;
  border: none;
  cursor: pointer;
  transition: color 0.3s ease;
}

.mobile-login-button:hover {
  color: #ffffff;
}

.mobile-register-button {
  padding: 0.625rem 1.25rem;
  background: linear-gradient(to right, #3b82f6, #8b5cf6);
  color: #ffffff;
  font-weight: 500;
  border: none;
  border-radius: 0.5rem;
  cursor: pointer;
  transition: all 0.3s ease;
  transform: scale(1);
}

.mobile-register-button:hover {
  background: linear-gradient(to right, #2563eb, #7c3aed);
  transform: scale(1.05);
  box-shadow: 0 0 50px -12px rgba(59, 130, 246, 0.25);
}

/* 主容器 */
.main-container {
  position: relative;
  z-index: 10;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.main-content {
  flex: 1;
}

/* 页脚 */
.footer {
  flex-shrink: 0;
}

.footer-container {
  border-top: 1px solid rgba(59, 130, 246, 0.2);
  background-color: rgba(0, 0, 0, 0.3);
  backdrop-filter: blur(4px);
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem 1.5rem;
  text-align: center;
}

.footer-brand {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  margin-bottom: 1rem;
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
}

.footer-copyright {
  color: #6b7280;
  font-size: 0.75rem;
  margin-top: 0.5rem;
}

/* 主题支持 */
[data-theme="light"] .app-container {
  background-color: #ffffff;
  color: #000000;
}

[data-theme="light"] .background-container {
  background: linear-gradient(to bottom, #ffffff, #f1f5f9, #ffffff);
}

/* 亮色主题下的星空动画 - 调整颜色和透明度 */
[data-theme="light"] .stars {
  background: transparent url('data:image/svg+xml,%3Csvg width="100" height="100" xmlns="http://www.w3.org/2000/svg"%3E%3Ccircle cx="50" cy="50" r="0.5" fill="%232563eb"/%3E%3C/svg%3E') repeat;
  opacity: 0.4;
}

[data-theme="light"] .stars2 {
  background: transparent url('data:image/svg+xml,%3Csvg width="100" height="100" xmlns="http://www.w3.org/2000/svg"%3E%3Ccircle cx="50" cy="50" r="0.8" fill="%237c3aed" opacity="0.4"/%3E%3C/svg%3E') repeat;
  opacity: 0.3;
}

/* 亮色主题下的导航栏 */
[data-theme="light"] .navbar {
  background: rgba(255, 255, 255, 0.95);
  border: 1px solid var(--color-border-light);
  backdrop-filter: blur(20px);
}

[data-theme="light"] .navbar:hover {
  border-color: var(--color-stellar-blue);
  box-shadow: 0 8px 32px rgba(37, 99, 235, 0.15);
}

/* 导航链接亮色主题 */
[data-theme="light"] .nav-link {
  color: var(--color-text-secondary);
  transition: color 0.3s ease;
}

[data-theme="light"] .nav-link:hover {
  color: var(--color-stellar-blue);
}

[data-theme="light"] .nav-link-active {
  color: var(--color-stellar-blue);
  font-weight: 600;
}

[data-theme="light"] .mobile-nav-link {
  color: var(--color-text-secondary);
}

[data-theme="light"] .mobile-nav-link:hover {
  color: var(--color-stellar-blue);
}

[data-theme="light"] .mobile-nav-link-active {
  color: var(--color-stellar-blue);
  font-weight: 600;
}

/* 亮色主题下的页脚 */
[data-theme="light"] .footer-container {
  background: rgba(255, 255, 255, 0.95);
  border-top: 1px solid var(--color-border-light);
}

[data-theme="light"] .footer-tagline {
  color: var(--color-text-muted);
}

[data-theme="light"] .footer-copyright {
  color: var(--color-text-muted);
  opacity: 0.7;
}

/* 品牌文字亮色主题适配 */
[data-theme="light"] .brand-text {
  background-image: linear-gradient(135deg, var(--color-stellar-blue), var(--color-nebula-purple), var(--color-aurora-pink));
  filter: none;
}

[data-theme="light"] .footer-brand-text {
  background-image: linear-gradient(135deg, var(--color-stellar-blue), var(--color-nebula-purple));
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