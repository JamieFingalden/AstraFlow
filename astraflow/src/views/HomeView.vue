<template>
  <MainLayout>
    <!-- 英雄区 -->
    <div class="min-h-screen flex flex-col items-center justify-center px-6 relative">
      <!-- 主标题 -->
      <div class="text-center mb-12">
        <h1 class="text-6xl md:text-8xl lg:text-9xl font-black mb-6 animate-float">
          <span class="bg-clip-text text-transparent bg-gradient-to-r from-blue-400 via-purple-500 to-pink-500 drop-shadow-2xl">
            星流
          </span>
        </h1>
        <div class="text-2xl md:text-4xl lg:text-5xl font-bold mb-4 text-gray-300">
          AI智能求职平台
        </div>
        <p class="text-lg md:text-xl text-gray-400 max-w-2xl mx-auto leading-relaxed">
          让AI为你点亮职业星空，智能匹配理想职位，定制专属职业发展路径
        </p>
      </div>

      <!-- 核心功能 -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 max-w-6xl w-full mb-12">
        <div class="group relative p-8 rounded-2xl bg-gradient-to-br from-blue-900/30 to-purple-900/30 backdrop-blur-md border border-blue-500/20 hover:border-blue-400/50 transition-all duration-500 hover:scale-105">
          <div class="absolute -inset-1 bg-gradient-to-r from-blue-500 to-purple-500 rounded-2xl opacity-0 group-hover:opacity-30 blur transition-opacity duration-500"></div>
          <div class="relative">
            <div class="text-4xl mb-4">✨</div>
            <h3 class="text-xl font-bold mb-2 text-blue-300">智能简历</h3>
            <p class="text-gray-400">AI驱动的简历优化，突出你的核心竞争力</p>
          </div>
        </div>

        <div class="group relative p-8 rounded-2xl bg-gradient-to-br from-purple-900/30 to-pink-900/30 backdrop-blur-md border border-purple-500/20 hover:border-purple-400/50 transition-all duration-500 hover:scale-105">
          <div class="absolute -inset-1 bg-gradient-to-r from-purple-500 to-pink-500 rounded-2xl opacity-0 group-hover:opacity-30 blur transition-opacity duration-500"></div>
          <div class="relative">
            <div class="text-4xl mb-4">🎯</div>
            <h3 class="text-xl font-bold mb-2 text-purple-300">精准匹配</h3>
            <p class="text-gray-400">深度分析岗位需求，智能推荐最适合机会</p>
          </div>
        </div>

        <div class="group relative p-8 rounded-2xl bg-gradient-to-br from-pink-900/30 to-orange-900/30 backdrop-blur-md border border-pink-500/20 hover:border-pink-400/50 transition-all duration-500 hover:scale-105">
          <div class="absolute -inset-1 bg-gradient-to-r from-pink-500 to-orange-500 rounded-2xl opacity-0 group-hover:opacity-30 blur transition-opacity duration-500"></div>
          <div class="relative">
            <div class="text-4xl mb-4">🚀</div>
            <h3 class="text-xl font-bold mb-2 text-pink-300">职业规划</h3>
            <p class="text-gray-400">可视化职业路径，助力长远发展</p>
          </div>
        </div>
      </div>

      <!-- CTA按钮 -->
      <div class="flex flex-col sm:flex-row gap-4">
        <router-link to="/resume" class="px-8 py-4 bg-gradient-to-r from-blue-500 to-purple-500 rounded-full font-bold text-lg hover:from-blue-600 hover:to-purple-600 transition-all duration-300 hover:scale-105 hover:shadow-2xl hover:shadow-blue-500/50">
          开启星流之旅
        </router-link>
        <router-link to="/job-analysis" class="px-8 py-4 border-2 border-purple-500/50 rounded-full font-bold text-lg hover:bg-purple-500/10 transition-all duration-300">
          探索更多
        </router-link>
      </div>
    </div>

    <!-- 流动效果 -->
    <div class="absolute bottom-0 left-0 right-0 h-32 bg-gradient-to-t from-black to-transparent"></div>
  </MainLayout>
</template>

<script setup>
import MainLayout from '../layouts/MainLayout.vue'
</script>

<style scoped>
/* 浮动动画 */
@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-20px); }
}

.animate-float {
  animation: float 6s ease-in-out infinite;
}

/* 渐变文字效果 */
@keyframes gradient {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.bg-gradient-to-r {
  background-size: 200% 200%;
  animation: gradient 3s ease infinite;
}

/* 主题支持 */
/* 暗色主题（默认） */
.hero-gradient {
  background: linear-gradient(135deg, #1e293b, #0f172a);
}

/* 亮色主题 */
[data-theme="light"] .hero-gradient {
  background: linear-gradient(135deg, #e0f2fe, #f0f9ff);
}

[data-theme="light"] .text-gray-300 {
  color: #64748b;
}

[data-theme="light"] .text-gray-400 {
  color: #475569;
}

/* 功能卡片主题适配 */
[data-theme="light"] .glass {
  background: rgba(255, 255, 255, 0.7);
  border-color: rgba(59, 130, 246, 0.2);
  box-shadow: 0 8px 32px rgba(59, 130, 246, 0.1);
}

/* 按钮主题适配 */
[data-theme="light"] .hover\:from-blue-600:hover {
  background: linear-gradient(to bottom right, #2563eb, #7c3aed);
}

[data-theme="light"] .hover\:bg-purple-500\/10:hover {
  background-color: rgba(168, 85, 247, 0.1);
}

/* 标题颜色适配 */
[data-theme="light"] .text-blue-400 {
  color: #3b82f6;
}

[data-theme="light"] .text-purple-300 {
  color: #a855f7;
}

[data-theme="light"] .text-pink-300 {
  color: #ec4899;
}
</style>