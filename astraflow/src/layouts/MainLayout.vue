<template>
  <div class="min-h-screen bg-black text-white relative overflow-hidden">
    <!-- 星空背景 -->
    <div class="fixed inset-0 bg-gradient-to-b from-black via-blue-950/20 to-black">
      <div class="absolute inset-0">
        <div class="stars absolute inset-0"></div>
        <div class="stars2 absolute inset-0"></div>
      </div>
    </div>

    <!-- 导航栏 -->
    <nav class="relative z-50 flex items-center justify-between p-6 backdrop-blur-lg bg-gradient-to-r from-black/20 via-blue-900/20 to-black/20 rounded-2xl mx-6 my-4 transition-all duration-300 ease-in-out">
      <div class="text-2xl font-bold flex items-center space-x-2">
        <div class="w-3 h-3 bg-blue-500 rounded-full animate-pulse shadow-lg shadow-blue-500/50"></div>
        <span class="bg-clip-text text-transparent bg-gradient-to-r from-blue-400 via-purple-400 to-pink-400">
          ASTRAFLOW
        </span>
      </div>
      <div class="hidden md:flex items-center space-x-8">
        <router-link to="/" class="text-gray-300 hover:text-white transition-all duration-300 px-3 py-2 rounded-lg">首页</router-link>
        <router-link to="/expense" class="text-gray-300 hover:text-white transition-all duration-300 px-3 py-2 rounded-lg">费用管理</router-link>
        <router-link to="/analytics" class="text-gray-300 hover:text-white transition-all duration-300 px-3 py-2 rounded-lg">数据分析</router-link>
        <router-link to="/profile" class="text-gray-300 hover:text-white transition-all duration-300 px-3 py-2 rounded-lg">个人中心</router-link>
      </div>

      <!-- 右侧操作区 -->
      <div class="flex items-center space-x-4">
        <button class="hidden md:inline-flex px-4 py-2 text-gray-300 hover:text-white font-medium transition-colors duration-300">
          登录
        </button>
        <button class="inline-flex px-5 py-2 bg-gradient-to-r from-blue-500 to-purple-500 text-white font-medium rounded-lg hover:from-blue-600 hover:to-purple-600 transition-all duration-300 transform hover:scale-105 shadow-lg shadow-blue-500/25">
          注册
        </button>
        <!-- 主题切换按钮 -->
        <ThemeToggle />
      </div>
    </nav>

    <!-- 主容器 -->
    <div class="relative z-10 min-h-screen flex flex-col">
      <main class="flex-grow">
        <slot></slot>
      </main>

      <!-- 页脚 -->
      <footer class="flex-shrink-0">
        <div class="border-t border-blue-500/20 bg-black/30 backdrop-blur-sm">
          <div class="container mx-auto px-6 py-8">
            <div class="text-center">
              <div class="flex items-center justify-center space-x-2 mb-4">
                <div class="w-2 h-2 bg-blue-500 rounded-full animate-pulse"></div>
                <span class="bg-clip-text text-transparent bg-gradient-to-r from-blue-400 to-cyan-400 text-lg font-bold">
                  ASTRAFLOW
                </span>
              </div>
              <p class="text-gray-400 text-sm">
                Smart Expense Made Simple ✨
              </p>
              <p class="text-gray-500 text-xs mt-2">
                © 2025 AstraFlow. All rights reserved.
              </p>
            </div>
          </div>
        </div>
      </footer>
    </div>
  </div>
</template>

<script setup>
import ThemeToggle from '../components/ui/ThemeToggle.vue'
</script>

<style scoped>
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

/* 主题支持 */
/* 亮色主题下的背景调整 - 增强对比度 */
[data-theme="light"] .stars {
  background: transparent url('data:image/svg+xml,%3Csvg width="100" height="100" xmlns="http://www.w3.org/2000/svg"%3E%3Ccircle cx="50" cy="50" r="0.5" fill="%232563eb"/%3E%3C/svg%3E') repeat;
  opacity: 0.4;
}

[data-theme="light"] .stars2 {
  background: transparent url('data:image/svg+xml,%3Csvg width="100" height="100" xmlns="http://www.w3.org/2000/svg"%3E%3Ccircle cx="50" cy="50" r="0.8" fill="%237c3aed" opacity="0.4"/%3E%3C/svg%3E') repeat;
  opacity: 0.3;
}

[data-theme="light"] .bg-gradient-to-b {
  background: linear-gradient(to bottom, #ffffff, #f1f5f9, #ffffff);
}

/* 亮色主题下的导航栏调整 - 增强可读性 */
[data-theme="light"] nav {
  background: rgba(255, 255, 255, 0.95);
  border: 1px solid var(--color-border-light);
  backdrop-filter: blur(20px);
}

[data-theme="light"] .hover\:border-blue-400\/40:hover {
  border-color: var(--color-stellar-blue);
  box-shadow: 0 8px 32px rgba(37, 99, 235, 0.15);
}

/* 导航链接亮色主题 */
[data-theme="light"] .text-gray-300 {
  color: var(--color-text-secondary);
  transition: color 0.3s ease;
}

[data-theme="light"] .text-gray-300:hover {
  color: var(--color-stellar-blue);
}

[data-theme="light"] .text-blue-400 {
  color: var(--color-stellar-blue);
  font-weight: 600;
}

/* 亮色主题下的页脚调整 - 更好的对比度 */
[data-theme="light"] footer > div {
  background: rgba(255, 255, 255, 0.95);
  border-top: 1px solid var(--color-border-light);
}

[data-theme="light"] .text-gray-400 {
  color: var(--color-text-muted);
}

[data-theme="light"] .text-gray-500 {
  color: var(--color-text-muted);
  opacity: 0.7;
}

/* 品牌文字亮色主题适配 */
[data-theme="light"] .bg-clip-text {
  background: linear-gradient(135deg, var(--color-stellar-blue), var(--color-nebula-purple), var(--color-aurora-pink));
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
  filter: none;
}
</style>