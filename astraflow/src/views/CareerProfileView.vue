<template>
  <MainLayout>
    <div class="min-h-screen flex items-center justify-center px-6 py-12 relative">
      <div class="max-w-6xl w-full">
        <!-- 标题区域 -->
        <div class="text-center mb-12">
          <h1 class="text-5xl md:text-6xl font-black mb-6">
            <span class="bg-clip-text text-transparent bg-gradient-to-r from-blue-400 via-purple-500 to-pink-500">
              职业画像
            </span>
          </h1>
          <p class="text-xl text-gray-400 max-w-2xl mx-auto">
            可视化职业路径，助力长远发展
          </p>
        </div>

        <!-- 个人资料卡片 -->
        <div class="group relative p-8 rounded-2xl bg-gradient-to-br from-blue-900/30 to-purple-900/30 backdrop-blur-md border border-blue-500/20 hover:border-blue-400/50 transition-all duration-500 mb-8">
          <div class="absolute -inset-1 bg-gradient-to-r from-blue-500 to-purple-500 rounded-2xl opacity-0 group-hover:opacity-30 blur transition-opacity duration-500"></div>
          <div class="relative">
            <div class="flex flex-col md:flex-row items-center md:items-start space-y-6 md:space-y-0 md:space-x-8">
              <div class="w-32 h-32 rounded-full bg-gradient-to-br from-blue-500 to-purple-500 p-1">
                <div class="w-full h-full rounded-full bg-black flex items-center justify-center text-4xl font-bold text-white">
                  A
                </div>
              </div>
              <div class="flex-1 text-center md:text-left">
                <h2 class="text-3xl font-bold text-white mb-2">张三</h2>
                <p class="text-blue-300 mb-4">高级前端开发工程师 | 5年经验</p>
                <div class="flex flex-wrap justify-center md:justify-start gap-4 text-sm">
                  <span class="px-3 py-1 bg-blue-500/20 text-blue-300 rounded-full">JavaScript</span>
                  <span class="px-3 py-1 bg-purple-500/20 text-purple-300 rounded-full">React</span>
                  <span class="px-3 py-1 bg-pink-500/20 text-pink-300 rounded-full">Vue.js</span>
                  <span class="px-3 py-1 bg-cyan-500/20 text-cyan-300 rounded-full">TypeScript</span>
                </div>
              </div>
              <div class="text-center">
                <div class="text-3xl font-bold text-yellow-400">A+</div>
                <div class="text-gray-400">职业评级</div>
              </div>
            </div>
          </div>
        </div>

        <!-- 技能雷达图 -->
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-8">
          <div class="group relative p-8 rounded-2xl bg-gradient-to-br from-purple-900/30 to-pink-900/30 backdrop-blur-md border border-purple-500/20 hover:border-purple-400/50 transition-all duration-500">
            <div class="absolute -inset-1 bg-gradient-to-r from-purple-500 to-pink-500 rounded-2xl opacity-0 group-hover:opacity-30 blur transition-opacity duration-500"></div>
            <div class="relative">
              <h3 class="text-2xl font-bold mb-6 text-purple-300">技能雷达图</h3>
              <div class="h-64 flex items-center justify-center">
                <div class="relative w-48 h-48">
                  <!-- 雷达图背景 -->
                  <div class="absolute inset-0 border border-purple-500/30 rounded-full flex items-center justify-center">
                    <div class="w-3/4 h-3/4 border border-purple-500/30 rounded-full flex items-center justify-center">
                      <div class="w-1/2 h-1/2 border border-purple-500/30 rounded-full"></div>
                    </div>
                  </div>
                  <!-- 雷达图数据 -->
                  <div class="absolute inset-0">
                    <div class="absolute top-0 left-1/2 transform -translate-x-1/2 w-1 h-full bg-gradient-to-t from-purple-500 to-transparent opacity-50"></div>
                    <div class="absolute right-0 top-1/2 transform -translate-y-1/2 h-1 w-full bg-gradient-to-l from-purple-500 to-transparent opacity-50"></div>
                    <div class="absolute bottom-0 left-1/2 transform -translate-x-1/2 w-1 h-full bg-gradient-to-b from-purple-500 to-transparent opacity-50"></div>
                    <div class="absolute left-0 top-1/2 transform -translate-y-1/2 h-1 w-full bg-gradient-to-r from-purple-500 to-transparent opacity-50"></div>
                    <!-- 数据覆盖层 -->
                    <div class="absolute inset-0">
                      <polygon
                        points="50%,20% 80%,35% 85%,70% 50%,85% 15%,70% 20%,35%"
                        fill="url(#skillGradient)"
                        opacity="0.6"
                      />
                      <polygon
                        points="50%,20% 80%,35% 85%,70% 50%,85% 15%,70% 20%,35%"
                        fill="transparent"
                        stroke="url(#skillGradient)"
                        stroke-width="2"
                      />
                    </div>
                  </div>
                  <svg class="absolute inset-0 w-full h-full">
                    <defs>
                      <linearGradient id="skillGradient" x1="0%" y1="0%" x2="100%" y2="100%">
                        <stop offset="0%" stop-color="#8b5cf6" />
                        <stop offset="100%" stop-color="#ec4899" />
                      </linearGradient>
                    </defs>
                  </svg>
                </div>
              </div>
              <div class="grid grid-cols-2 gap-4 mt-6">
                <div class="flex justify-between">
                  <span class="text-gray-300">前端开发</span>
                  <span class="text-blue-400">95%</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-gray-300">后端能力</span>
                  <span class="text-purple-400">75%</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-gray-300">设计思维</span>
                  <span class="text-pink-400">70%</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-gray-300">项目管理</span>
                  <span class="text-cyan-400">65%</span>
                </div>
              </div>
            </div>
          </div>

          <!-- 职业发展路径 -->
          <div class="group relative p-8 rounded-2xl bg-gradient-to-br from-pink-900/30 to-orange-900/30 backdrop-blur-md border border-pink-500/20 hover:border-pink-400/50 transition-all duration-500">
            <div class="absolute -inset-1 bg-gradient-to-r from-pink-500 to-orange-500 rounded-2xl opacity-0 group-hover:opacity-30 blur transition-opacity duration-500"></div>
            <div class="relative">
              <h3 class="text-2xl font-bold mb-6 text-pink-300">职业发展路径</h3>
              <div class="space-y-6">
                <div class="flex items-center space-x-4">
                  <div class="w-12 h-12 rounded-full bg-gradient-to-r from-blue-500 to-purple-500 flex items-center justify-center text-white font-bold">
                    1
                  </div>
                  <div class="flex-1">
                    <h4 class="text-lg font-bold text-white">前端开发工程师</h4>
                    <p class="text-gray-400">1-3年经验 | 初级</p>
                  </div>
                </div>
                <div class="flex items-center space-x-4">
                  <div class="w-12 h-12 rounded-full bg-gradient-to-r from-purple-500 to-pink-500 flex items-center justify-center text-white font-bold">
                    2
                  </div>
                  <div class="flex-1">
                    <h4 class="text-lg font-bold text-white">高级前端开发工程师</h4>
                    <p class="text-gray-400">3-5年经验 | 中级</p>
                  </div>
                </div>
                <div class="flex items-center space-x-4">
                  <div class="w-12 h-12 rounded-full bg-gradient-to-r from-pink-500 to-orange-500 flex items-center justify-center text-white font-bold">
                    3
                  </div>
                  <div class="flex-1">
                    <h4 class="text-lg font-bold text-white">技术专家/架构师</h4>
                    <p class="text-gray-400">5-8年经验 | 高级</p>
                  </div>
                </div>
                <div class="flex items-center space-x-4">
                  <div class="w-12 h-12 rounded-full bg-gradient-to-r from-orange-500 to-red-500 flex items-center justify-center text-white font-bold">
                    4
                  </div>
                  <div class="flex-1">
                    <h4 class="text-lg font-bold text-white">技术总监/CTO</h4>
                    <p class="text-gray-400">8+年经验 | 管理层</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 推荐学习路径 -->
        <div class="group relative p-8 rounded-2xl bg-gradient-to-br from-cyan-900/30 to-blue-900/30 backdrop-blur-md border border-cyan-500/20 hover:border-cyan-400/50 transition-all duration-500 mb-8">
          <div class="absolute -inset-1 bg-gradient-to-r from-cyan-500 to-blue-500 rounded-2xl opacity-0 group-hover:opacity-30 blur transition-opacity duration-500"></div>
          <div class="relative">
            <h3 class="text-2xl font-bold mb-6 text-cyan-300">推荐学习路径</h3>
            <div class="space-y-4">
              <div class="p-4 bg-black/20 rounded-xl">
                <div class="flex items-center justify-between mb-2">
                  <h4 class="text-lg font-bold text-white">微前端架构实践</h4>
                  <span class="text-cyan-400 text-sm">进阶课程</span>
                </div>
                <p class="text-gray-400 mb-3">学习现代微前端架构，提升大型项目管理能力</p>
                <div class="flex items-center space-x-4 text-sm">
                  <span class="px-2 py-1 bg-blue-500/20 text-blue-300 rounded">React</span>
                  <span class="px-2 py-1 bg-purple-500/20 text-purple-300 rounded">Vue</span>
                  <span class="text-gray-500">预计 4周</span>
                </div>
              </div>
              <div class="p-4 bg-black/20 rounded-xl">
                <div class="flex items-center justify-between mb-2">
                  <h4 class="text-lg font-bold text-white">TypeScript高级特性</h4>
                  <span class="text-cyan-400 text-sm">进阶课程</span>
                </div>
                <p class="text-gray-400 mb-3">深入掌握TypeScript高级类型系统和泛型编程</p>
                <div class="flex items-center space-x-4 text-sm">
                  <span class="px-2 py-1 bg-pink-500/20 text-pink-300 rounded">TypeScript</span>
                  <span class="text-gray-500">预计 3周</span>
                </div>
              </div>
              <div class="p-4 bg-black/20 rounded-xl">
                <div class="flex items-center justify-between mb-2">
                  <h4 class="text-lg font-bold text-white">性能优化实战</h4>
                  <span class="text-cyan-400 text-sm">进阶课程</span>
                </div>
                <p class="text-gray-400 mb-3">前端性能优化策略与实践，提升用户体验</p>
                <div class="flex items-center space-x-4 text-sm">
                  <span class="px-2 py-1 bg-blue-500/20 text-blue-300 rounded">Performance</span>
                  <span class="text-gray-500">预计 3周</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 行动按钮 -->
        <div class="text-center">
          <button class="py-4 px-8 bg-gradient-to-r from-blue-500 to-purple-500 rounded-xl font-bold text-lg hover:from-blue-600 hover:to-purple-600 transition-all duration-300 hover:scale-105 hover:shadow-2xl hover:shadow-blue-500/50">
            生成我的职业报告
          </button>
        </div>
      </div>
    </div>
  </MainLayout>
</template>

<script setup>
import MainLayout from '../layouts/MainLayout.vue'
</script>

<style scoped>
/* 亮色主题样式 */
[data-theme="light"] .bg-gradient-to-br {
  background: linear-gradient(135deg, var(--color-bg-card), var(--color-bg-card)) !important;
  border: 1px solid var(--color-border-light) !important;
}

[data-theme="light"] .backdrop-blur-md {
  backdrop-filter: blur(20px) !important;
  -webkit-backdrop-filter: blur(20px) !important;
}

[data-theme="light"] .text-gray-300 {
  color: var(--color-text-secondary) !important;
}

[data-theme="light"] .text-gray-400 {
  color: var(--color-text-muted) !important;
}

[data-theme="light"] .text-blue-300 {
  color: var(--color-stellar-blue) !important;
  font-weight: 700;
}

[data-theme="light"] .text-purple-300 {
  color: var(--color-nebula-purple) !important;
  font-weight: 700;
}

[data-theme="light"] .text-pink-300 {
  color: var(--color-aurora-pink) !important;
  font-weight: 700;
}

[data-theme="light"] .text-cyan-300 {
  color: var(--color-cosmic-cyan) !important;
  font-weight: 700;
}

[data-theme="light"] .bg-black\/20 {
  background: rgba(255, 255, 255, 0.5) !important;
}

[data-theme="light"] .bg-black\/30 {
  background: rgba(255, 255, 255, 0.7) !important;
}

[data-theme="light"] .text-white {
  color: var(--color-text-primary) !important;
}

[data-theme="light"] .text-yellow-400 {
  color: #d97706 !important;
}

[data-theme="light"] .border-cyan-500\/20 {
  border-color: var(--color-border-light) !important;
}

[data-theme="light"] button {
  background: linear-gradient(135deg, var(--color-stellar-blue), var(--color-nebula-purple)) !important;
  color: white !important;
  font-weight: 700 !important;
  box-shadow: 0 8px 25px rgba(30, 64, 175, 0.3) !important;
}

[data-theme="light"] button:hover {
  background: linear-gradient(135deg, var(--color-stellar-blue), var(--color-nebula-purple)) !important;
  transform: translateY(-3px) scale(1.02) !important;
  box-shadow: 0 12px 35px rgba(30, 64, 175, 0.4) !important;
}

/* 背景渐变文字 */
.bg-gradient-to-r {
  background-size: 200% 200%;
  animation: gradient 3s ease infinite;
}

@keyframes gradient {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

/* 悬停动效 */
.group {
  position: relative;
  transition: all 0.3s ease;
}

.group:hover {
  transform: translateY(-5px);
}

/* 个人头像样式 */
.w-32.h-32.rounded-full {
  background: linear-gradient(135deg, var(--color-stellar-blue), var(--color-nebula-purple));
}

/* 亮色主题头像 */
[data-theme="light"] .w-32.h-32.rounded-full {
  background: linear-gradient(135deg, var(--color-stellar-blue), var(--color-nebula-purple));
}
</style>