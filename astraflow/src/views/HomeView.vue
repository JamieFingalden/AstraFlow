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
        <router-link to="/resume" class="cta-primary px-8 py-4 bg-gradient-to-r from-blue-500 to-purple-500 rounded-full font-bold text-lg hover:from-blue-600 hover:to-purple-600 transition-all duration-300 hover:scale-105 hover:shadow-2xl hover:shadow-blue-500/50">
          开启星流之旅
        </router-link>
        <router-link to="/job-analysis" class="cta-secondary px-8 py-4 border-2 border-purple-500/50 rounded-full font-bold text-lg hover:bg-purple-500/10 transition-all duration-300">
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

/* 亮色主题 - 全面优化 */
[data-theme="light"] .hero-gradient {
  background: linear-gradient(135deg, #f8fafc, #e2e8f0);
}

/* 主要文字 - 强制覆盖 */
[data-theme="light"] .text-gray-300 {
  color: var(--color-text-primary) !important;
  font-weight: 500;
}

[data-theme="light"] .text-gray-400 {
  color: var(--color-text-secondary) !important;
}

/* 主标题渐变文字 - 亮色主题专用 */
[data-theme="light"] .bg-gradient-to-r {
  background: linear-gradient(135deg, var(--color-stellar-blue), var(--color-nebula-purple), var(--color-aurora-pink)) !important;
  background-clip: text !important;
  -webkit-background-clip: text !important;
  color: transparent !important;
}

/* 功能卡片 - 完全重新设计 */
[data-theme="light"] .group {
  /* 覆盖Tailwind的背景渐变 */
  background: var(--color-bg-card) !important;
  border: 1px solid var(--color-border-light) !important;
  box-shadow: 0 8px 32px rgba(37, 99, 235, 0.08) !important;
}

[data-theme="light"] .group:hover {
  background: rgba(255, 255, 255, 0.98) !important;
  border-color: var(--color-stellar-blue) !important;
  box-shadow: 0 12px 40px rgba(37, 99, 235, 0.15) !important;
  transform: scale(1.02) !important;
}

/* 覆盖功能卡片的悬停光晕效果 */
[data-theme="light"] .group .absolute {
  background: linear-gradient(135deg, var(--color-stellar-blue), var(--color-nebula-purple)) !important;
  opacity: 0.1 !important;
}

[data-theme="light"] .group:hover .absolute {
  opacity: 0.2 !important;
}

/* 功能卡片标题 - 更高对比度 */
[data-theme="light"] .text-blue-300 {
  color: var(--color-stellar-blue) !important;
  font-weight: 700 !important;
}

[data-theme="light"] .text-purple-300 {
  color: var(--color-nebula-purple) !important;
  font-weight: 700 !important;
}

[data-theme="light"] .text-pink-300 {
  color: var(--color-aurora-pink) !important;
  font-weight: 700 !important;
}

/* 功能卡片描述文字 */
[data-theme="light"] .group .text-gray-400 {
  color: var(--color-text-secondary) !important;
  line-height: 1.6 !important;
}

/* CTA按钮区域优化 - 统一白色文字 */

/* "开始星流之旅" 主按钮 - 亮色主题 */
[data-theme="light"] .cta-primary {
  background: linear-gradient(135deg, #1e40af, #7c3aed) !important;
  color: white !important;
  font-weight: 700 !important;
  box-shadow: 0 8px 25px rgba(30, 64, 175, 0.3) !important;
  border: none !important;
  position: relative;
  overflow: hidden;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1) !important;
}

[data-theme="light"] .cta-primary:hover {
  background: linear-gradient(135deg, #1e3a8a, #6d28d9) !important;
  box-shadow: 0 12px 35px rgba(30, 64, 175, 0.4) !important;
  transform: translateY(-3px) scale(1.02) !important;
  color: white !important;
}

[data-theme="light"] .cta-primary:active {
  transform: translateY(-1px) scale(1.01) !important;
  color: white !important;
}

/* "开始星流之旅" 按钮内部光晕效果 */
[data-theme="light"] .cta-primary::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.8s;
}

[data-theme="light"] .cta-primary:hover::before {
  left: 100%;
}

/* "探索更多" 次按钮 - 亮色主题 */
[data-theme="light"] .cta-secondary {
  background: rgba(124, 58, 237, 0.9) !important;
  border: 2px solid #7c3aed !important;
  color: white !important;
  font-weight: 600 !important;
  backdrop-filter: blur(10px);
  position: relative;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2) !important;
}

[data-theme="light"] .cta-secondary:hover {
  background: rgba(37, 99, 235, 0.9) !important;
  border-color: #2563eb !important;
  color: white !important;
  box-shadow: 0 8px 25px rgba(37, 99, 235, 0.3) !important;
  transform: translateY(-2px) !important;
}

/* 确保常态下就是白色文字 - 超强优先级 */
.cta-primary:not(:hover):not(:active):not(:focus),
.cta-secondary:not(:hover):not(:active):not(:focus),
a.cta-primary:not(:hover):not(:active):not(:focus),
a.cta-secondary:not(:hover):not(:active):not(:focus) {
  color: white !important;
}

/* 确保所有状态下都是白色文字 */
.cta-primary,
.cta-secondary,
a.cta-primary,
a.cta-secondary,
.cta-primary *,
.cta-secondary * {
  color: white !important;
}

.cta-primary:hover,
.cta-secondary:hover,
a.cta-primary:hover,
a.cta-secondary:hover,
.cta-primary:hover *,
.cta-secondary:hover * {
  color: white !important;
}

/* 针对所有可能的状态 */
.cta-primary:active,
.cta-secondary:active,
.cta-primary:focus,
.cta-secondary:focus {
  color: white !important;
}

/* 特别针对router-link的内部文字 */
.cta-primary span,
.cta-secondary span,
.cta-primary .text-lg,
.cta-secondary .text-lg {
  color: white !important;
}

/* 按钮通用优化 */
[data-theme="light"] .cta-primary,
[data-theme="light"] .cta-secondary {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
  border-radius: 9999px !important;
}

/* 确保按钮容器在亮色主题下的间距 */
[data-theme="light"] .flex-col.sm\:flex-row {
  gap: 1.25rem !important;
}

/* 整体背景和布局优化 */
[data-theme="light"] .min-h-screen {
  background: linear-gradient(135deg, #ffffff, #f8fafc, #ffffff);
}

/* 确保所有文字在亮色主题下都有足够对比度 */
[data-theme="light"] .text-center {
  color: var(--color-text-primary);
}

[data-theme="light"] .text-2xl,
[data-theme="light"] .text-4xl,
[data-theme="light"] .text-5xl {
  color: var(--color-text-primary) !important;
}

[data-theme="light"] .text-lg,
[data-theme="light"] .text-xl {
  color: var(--color-text-secondary) !important;
}
</style>