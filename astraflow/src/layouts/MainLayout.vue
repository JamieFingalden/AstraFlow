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
    <nav class="relative z-50 flex items-center justify-between p-6 backdrop-blur-lg bg-gradient-to-r from-black/20 via-blue-900/20 to-black/20 border border-blue-500/20 rounded-2xl mx-6 my-4 hover:border-blue-400/40 transition-all duration-500">
      <div class="text-2xl font-bold flex items-center space-x-2">
        <div class="w-3 h-3 bg-blue-500 rounded-full animate-pulse shadow-lg shadow-blue-500/50"></div>
        <span class="bg-clip-text text-transparent bg-gradient-to-r from-blue-400 via-purple-400 to-pink-400">
          ASTRAFLOW
        </span>
      </div>
      <div class="hidden md:flex items-center space-x-8">
        <router-link to="/" class="text-gray-300 hover:text-white transition-all duration-300 px-3 py-2 rounded-lg hover:bg-white/10 drop-shadow-glow">首页</router-link>
        <router-link to="/resume" class="text-blue-400 hover:text-blue-300 transition-all duration-300 px-3 py-2 rounded-lg hover:bg-white/10 drop-shadow-glow">简历生成</router-link>
        <router-link to="/job-analysis" class="text-gray-300 hover:text-white transition-all duration-300 px-3 py-2 rounded-lg hover:bg-white/10 drop-shadow-glow">岗位分析</router-link>
        <router-link to="/profile" class="text-gray-300 hover:text-white transition-all duration-300 px-3 py-2 rounded-lg hover:bg-white/10 drop-shadow-glow">职业画像</router-link>
      </div>

      <!-- 主题切换按钮 -->
      <ThemeToggle />
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
                <span class="bg-clip-text text-transparent bg-gradient-to-r from-blue-400 via-purple-400 to-pink-400 text-lg font-bold">
                  ASTRAFLOW
                </span>
              </div>
              <p class="text-gray-400 text-sm">
                让求职更智能，让未来更清晰 ✨
              </p>
              <p class="text-gray-500 text-xs mt-2">
                © 2024 AstraFlow. All rights reserved.
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
/* 亮色主题下的背景调整 */
[data-theme="light"] .stars {
  background: transparent url('data:image/svg+xml,%3Csvg width="100" height="100" xmlns="http://www.w3.org/2000/svg"%3E%3Ccircle cx="50" cy="50" r="0.5" fill="%233b82f6"/%3E%3C/svg%3E') repeat;
  opacity: 0.3;
}

[data-theme="light"] .stars2 {
  background: transparent url('data:image/svg+xml,%3Csvg width="100" height="100" xmlns="http://www.w3.org/2000/svg"%3E%3Ccircle cx="50" cy="50" r="0.8" fill="%238b5cf6" opacity="0.3"/%3E%3C/svg%3E') repeat;
  opacity: 0.2;
}

[data-theme="light"] .bg-gradient-to-b {
  background: linear-gradient(to bottom, #ffffff, #e0f2fe, #ffffff);
}

/* 亮色主题下的导航栏调整 */
[data-theme="light"] .bg-black\/30 {
  background: rgba(255, 255, 255, 0.8);
}

[data-theme="light"] .border-blue-500\/20 {
  border-color: rgba(59, 130, 246, 0.3);
}

[data-theme="light"] .hover\:border-blue-400\/40:hover {
  border-color: rgba(59, 130, 246, 0.5);
}

/* 亮色主题下的页脚调整 */
[data-theme="light"] .bg-black\/30 {
  background: rgba(255, 255, 255, 0.9);
}

[data-theme="light"] .border-blue-500\/20 {
  border-color: rgba(59, 130, 246, 0.2);
}

[data-theme="light"] .text-gray-400 {
  color: #64748b;
}

[data-theme="light"] .text-gray-500 {
  color: #94a3b8;
}
</style>