<template>
  <MainLayout>
    <div class="min-h-screen flex items-center justify-center px-6 py-12 relative">
      <div class="max-w-4xl w-full">
        <!-- 标题区域 -->
        <div class="text-center mb-12">
          <h1 class="text-5xl md:text-6xl font-black mb-6">
            <span class="bg-clip-text text-transparent bg-gradient-to-r from-blue-400 via-purple-500 to-pink-500">
              岗位分析
            </span>
          </h1>
          <p class="text-xl text-gray-400 max-w-2xl mx-auto">
            深度解析职位需求，智能匹配最适合的机会
          </p>
        </div>

        <!-- 分析工具卡片 -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-8 mb-12">
          <!-- 职位匹配度分析 -->
          <div class="group relative p-8 rounded-2xl bg-gradient-to-br from-blue-900/30 to-purple-900/30 backdrop-blur-md border border-blue-500/20 hover:border-blue-400/50 transition-all duration-500 hover:scale-105">
            <div class="absolute -inset-1 bg-gradient-to-r from-blue-500 to-purple-500 rounded-2xl opacity-0 group-hover:opacity-30 blur transition-opacity duration-500"></div>
            <div class="relative">
              <div class="text-4xl mb-4">🎯</div>
              <h3 class="text-2xl font-bold mb-4 text-blue-300">职位匹配度分析</h3>
              <p class="text-gray-400 mb-6">AI智能分析你的简历与目标职位的匹配程度</p>
              <div class="space-y-4">
                <div class="flex justify-between items-center">
                  <span class="text-gray-300">技能匹配度</span>
                  <span class="text-blue-400 font-bold">85%</span>
                </div>
                <div class="w-full bg-gray-700 rounded-full h-2">
                  <div class="bg-gradient-to-r from-blue-500 to-purple-500 h-2 rounded-full" style="width: 85%"></div>
                </div>

                <div class="flex justify-between items-center">
                  <span class="text-gray-300">经验匹配度</span>
                  <span class="text-purple-400 font-bold">78%</span>
                </div>
                <div class="w-full bg-gray-700 rounded-full h-2">
                  <div class="bg-gradient-to-r from-purple-500 to-pink-500 h-2 rounded-full" style="width: 78%"></div>
                </div>
              </div>
            </div>
          </div>

          <!-- 技能差距分析 -->
          <div class="group relative p-8 rounded-2xl bg-gradient-to-br from-purple-900/30 to-pink-900/30 backdrop-blur-md border border-purple-500/20 hover:border-purple-400/50 transition-all duration-500 hover:scale-105">
            <div class="absolute -inset-1 bg-gradient-to-r from-purple-500 to-pink-500 rounded-2xl opacity-0 group-hover:opacity-30 blur transition-opacity duration-500"></div>
            <div class="relative">
              <div class="text-4xl mb-4">🔍</div>
              <h3 class="text-2xl font-bold mb-4 text-purple-300">技能差距分析</h3>
              <p class="text-gray-400 mb-6">识别与目标职位的关键技能差距</p>
              <div class="space-y-3">
                <div class="flex items-center space-x-3">
                  <div class="w-2 h-2 bg-green-500 rounded-full"></div>
                  <span class="text-gray-300">JavaScript (已具备)</span>
                </div>
                <div class="flex items-center space-x-3">
                  <div class="w-2 h-2 bg-yellow-500 rounded-full"></div>
                  <span class="text-gray-300">TypeScript (需提升)</span>
                </div>
                <div class="flex items-center space-x-3">
                  <div class="w-2 h-2 bg-red-500 rounded-full"></div>
                  <span class="text-gray-300">React Native (需学习)</span>
                </div>
                <div class="flex items-center space-x-3">
                  <div class="w-2 h-2 bg-red-500 rounded-full"></div>
                  <span class="text-gray-300">GraphQL (需学习)</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 市场分析 -->
        <div class="group relative p-8 rounded-2xl bg-gradient-to-br from-pink-900/30 to-orange-900/30 backdrop-blur-md border border-pink-500/20 hover:border-pink-400/50 transition-all duration-500 mb-12">
          <div class="absolute -inset-1 bg-gradient-to-r from-pink-500 to-orange-500 rounded-2xl opacity-0 group-hover:opacity-30 blur transition-opacity duration-500"></div>
          <div class="relative">
            <div class="text-4xl mb-4">📊</div>
            <h3 class="text-2xl font-bold mb-6 text-pink-300">市场薪资分析</h3>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div class="text-center p-4 rounded-xl bg-black/20">
                <div class="text-3xl font-bold text-blue-400">¥25K</div>
                <div class="text-gray-400">平均薪资</div>
              </div>
              <div class="text-center p-4 rounded-xl bg-black/20">
                <div class="text-3xl font-bold text-purple-400">¥35K</div>
                <div class="text-gray-400">最高薪资</div>
              </div>
              <div class="text-center p-4 rounded-xl bg-black/20">
                <div class="text-3xl font-bold text-pink-400">¥18K</div>
                <div class="text-gray-400">最低薪资</div>
              </div>
            </div>
          </div>
        </div>

        <!-- 分析输入区域 -->
        <div class="bg-gradient-to-br from-blue-900/20 to-purple-900/20 backdrop-blur-md border border-blue-500/20 rounded-2xl p-8">
          <div class="text-center mb-6">
            <h3 class="text-2xl font-bold text-gray-200 mb-2">开始岗位分析</h3>
            <p class="text-gray-400">输入职位信息，获取智能分析报告</p>
          </div>

          <div class="space-y-6">
            <div>
              <label class="block text-gray-300 mb-3">职位名称</label>
              <input
                type="text"
                placeholder="例如：高级前端开发工程师"
                class="w-full px-4 py-3 bg-black/30 border border-blue-500/30 rounded-xl text-white placeholder-gray-500 focus:outline-none focus:border-blue-400 transition-colors"
              />
            </div>

            <div>
              <label class="block text-gray-300 mb-3">公司要求</label>
              <textarea
                placeholder="粘贴职位描述或要求..."
                rows="6"
                class="w-full px-4 py-3 bg-black/30 border border-blue-500/30 rounded-xl text-white placeholder-gray-500 focus:outline-none focus:border-blue-400 transition-colors resize-none"
              ></textarea>
            </div>

            <button class="w-full py-4 bg-gradient-to-r from-blue-500 to-purple-500 rounded-xl font-bold text-lg hover:from-blue-600 hover:to-purple-600 transition-all duration-300 hover:scale-105 hover:shadow-2xl hover:shadow-blue-500/50">
              开始智能分析
            </button>
          </div>
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

[data-theme="light"] .bg-black\/30 {
  background: rgba(255, 255, 255, 0.7) !important;
}

[data-theme="light"] .border-blue-500\/30 {
  border-color: var(--color-border-light) !important;
}

[data-theme="light"] .placeholder-gray-500 {
  color: var(--color-text-muted) !important;
}

[data-theme="light"] input,
[data-theme="light"] textarea {
  background: rgba(255, 255, 255, 0.8) !important;
  border: 1px solid var(--color-border-light) !important;
  color: var(--color-text-primary) !important;
}

[data-theme="light"] input::placeholder,
[data-theme="light"] textarea::placeholder {
  color: var(--color-text-muted) !important;
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
</style>