<template>
  <MainLayout>
    <!-- 英雄区 -->
    <div class="hero-container">
      <!-- 主标题 -->
      <div class="title-section">
        <h1 class="main-title animate-float">
          <span class="gradient-text">
            星流
          </span>
        </h1>
        <div class="subtitle">
          AI智能报销管理平台
        </div>
        <p class="description">
          AI自动识别发票账单，一键完成智能报销
        </p>
      </div>

      <!-- 核心功能 -->
      <div class="features-grid">
        <div class="feature-card feature-card-blue">
          <div class="feature-glow feature-glow-blue"></div>
          <div class="feature-content">
            <div class="feature-icon">📄</div>
            <h3 class="feature-title">智能识别</h3>
            <p class="feature-desc">AI自动识别发票、收据和账单信息</p>
          </div>
        </div>

        <div class="feature-card feature-card-purple">
          <div class="feature-glow feature-glow-purple"></div>
          <div class="feature-content">
            <div class="feature-icon">🎯</div>
            <h3 class="feature-title">自动分类</h3>
            <p class="feature-desc">智能分类管理，账单结构化一目了然</p>
          </div>
        </div>

        <div class="feature-card feature-card-pink">
          <div class="feature-glow feature-glow-pink"></div>
          <div class="feature-content">
            <div class="feature-icon">🚀</div>
            <h3 class="feature-title">一键报销</h3>
            <p class="feature-desc">快速生成报销单，告别繁琐流程</p>
          </div>
        </div>
      </div>

      <!-- CTA按钮 -->
      <div class="cta-container">
        <router-link to="/upload" class="cta-primary">
          上传发票账单
        </router-link>
        <router-link to="/dashboard" class="cta-secondary">
          账单可视化
        </router-link>
      </div>
    </div>

    <!-- 流动效果 -->
    <div class="flow-gradient"></div>
  </MainLayout>
</template>

<script setup>
import MainLayout from '../layouts/MainLayout.vue'
</script>

<style scoped>
/* 主容器 */
.hero-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 0 1.5rem;
  position: relative;
}

/* 标题区域 */
.title-section {
  text-align: center;
  margin-bottom: 3rem;
}

.main-title {
  font-size: 3.5rem;
  font-weight: 900;
  margin-bottom: 1.5rem;
}

@media (min-width: 768px) {
  .main-title {
    font-size: 5rem;
  }
}

@media (min-width: 1024px) {
  .main-title {
    font-size: 6rem;
  }
}

.gradient-text {
  background: linear-gradient(135deg, #60a5fa, #a855f7, #ec4899);
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
  filter: drop-shadow(0 25px 25px rgba(0, 0, 0, 0.25));
}

.subtitle {
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 1rem;
  color: #d1d5db;
}

@media (min-width: 768px) {
  .subtitle {
    font-size: 2.25rem;
  }
}

@media (min-width: 1024px) {
  .subtitle {
    font-size: 3rem;
  }
}

.description {
  font-size: 1.125rem;
  color: #9ca3af;
  max-width: 42rem;
  margin: 0 auto;
  line-height: 1.75;
}

@media (min-width: 768px) {
  .description {
    font-size: 1.25rem;
  }
}

/* 功能网格 */
.features-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1.5rem;
  max-width: 72rem;
  width: 100%;
  margin-bottom: 3rem;
}

@media (min-width: 768px) {
  .features-grid {
    grid-template-columns: repeat(3, 1fr);
  }
}

/* 功能卡片 */
.feature-card {
  position: relative;
  padding: 2rem;
  border-radius: 1rem;
  backdrop-filter: blur(12px);
  border: 1px solid;
  transition: all 0.5s ease;
}

.feature-card:hover {
  transform: scale(1.05);
}

.feature-card-blue {
  background: linear-gradient(135deg, rgba(30, 58, 138, 0.3), rgba(91, 33, 182, 0.3));
  border-color: rgba(59, 130, 246, 0.2);
}

.feature-card-blue:hover {
  border-color: rgba(96, 165, 250, 0.5);
}

.feature-card-purple {
  background: linear-gradient(135deg, rgba(91, 33, 182, 0.3), rgba(159, 18, 57, 0.3));
  border-color: rgba(139, 92, 246, 0.2);
}

.feature-card-purple:hover {
  border-color: rgba(168, 85, 247, 0.5);
}

.feature-card-pink {
  background: linear-gradient(135deg, rgba(159, 18, 57, 0.3), rgba(221, 107, 32, 0.3));
  border-color: rgba(236, 72, 153, 0.2);
}

.feature-card-pink:hover {
  border-color: rgba(251, 146, 60, 0.5);
}

/* 功能卡片光晕 */
.feature-glow {
  position: absolute;
  inset: -0.25rem;
  border-radius: 1rem;
  opacity: 0;
  filter: blur(1rem);
  transition: opacity 0.5s ease;
}

.feature-card:hover .feature-glow {
  opacity: 0.3;
}

.feature-glow-blue {
  background: linear-gradient(135deg, #3b82f6, #8b5cf6);
}

.feature-glow-purple {
  background: linear-gradient(135deg, #8b5cf6, #ec4899);
}

.feature-glow-pink {
  background: linear-gradient(135deg, #ec4899, #fb923c);
}

/* 功能内容 */
.feature-content {
  position: relative;
}

.feature-icon {
  font-size: 2.25rem;
  margin-bottom: 1rem;
}

.feature-title {
  font-size: 1.25rem;
  font-weight: 700;
  margin-bottom: 0.5rem;
}

.feature-card-blue .feature-title {
  color: #93c5fd;
}

.feature-card-purple .feature-title {
  color: #d8b4fe;
}

.feature-card-pink .feature-title {
  color: #fbcfe8;
}

.feature-desc {
  color: #9ca3af;
  line-height: 1.6;
}

/* CTA区域 */
.cta-container {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  margin-bottom: 15rem;
}

@media (min-width: 640px) {
  .cta-container {
    flex-direction: row;
  }
}

.cta-primary,
.cta-secondary {
  display: inline-block;
  padding: 1rem 2rem;
  font-weight: 700;
  font-size: 1.125rem;
  text-decoration: none;
  border-radius: 9999px;
  transition: all 0.3s ease;
}

.cta-primary {
  background: linear-gradient(135deg, #3b82f6, #8b5cf6);
  color: white;
}

.cta-primary:hover {
  background: linear-gradient(135deg, #2563eb, #7c3aed);
  transform: scale(1.05);
  box-shadow: 0 25px 50px -12px rgba(59, 130, 246, 0.5);
}

.cta-secondary {
  border: 2px solid rgba(139, 92, 246, 0.5);
  color: white;
}

.cta-secondary:hover {
  background: rgba(139, 92, 246, 0.1);
}

/* 流动渐变 */
.flow-gradient {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 8rem;
  background: linear-gradient(to top, #000000, transparent);
}

/* 浮动动画 */
@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-20px); }
}

.animate-float {
  animation: float 6s ease-in-out infinite;
}

/* 渐变动画 */
@keyframes gradient {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.gradient-text {
  background-size: 200% 200%;
  animation: gradient 3s ease infinite;
}

/* 亮色主题适配 */
[data-theme="light"] .hero-container {
  background: linear-gradient(135deg, #ffffff, #f8fafc, #ffffff);
}

[data-theme="light"] .subtitle {
  color: var(--color-text-primary);
  font-weight: 500;
}

[data-theme="light"] .description {
  color: var(--color-text-secondary);
}

[data-theme="light"] .gradient-text {
  background: linear-gradient(135deg, var(--color-stellar-blue), var(--color-nebula-purple), var(--color-aurora-pink));
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.1));
}

[data-theme="light"] .feature-card {
  background: var(--color-bg-card);
  border: 1px solid var(--color-border-light);
  box-shadow: 0 8px 32px rgba(37, 99, 235, 0.08);
}

[data-theme="light"] .feature-card:hover {
  background: rgba(255, 255, 255, 0.98);
  border-color: var(--color-stellar-blue);
  box-shadow: 0 12px 40px rgba(37, 99, 235, 0.15);
  transform: scale(1.02);
}

[data-theme="light"] .feature-glow {
  background: linear-gradient(135deg, var(--color-stellar-blue), var(--color-nebula-purple));
  opacity: 0.1;
}

[data-theme="light"] .feature-card:hover .feature-glow {
  opacity: 0.2;
}

[data-theme="light"] .feature-card-blue .feature-title {
  color: var(--color-stellar-blue);
  font-weight: 700;
}

[data-theme="light"] .feature-card-purple .feature-title {
  color: var(--color-nebula-purple);
  font-weight: 700;
}

[data-theme="light"] .feature-card-pink .feature-title {
  color: var(--color-aurora-pink);
  font-weight: 700;
}

[data-theme="light"] .feature-desc {
  color: var(--color-text-secondary);
}

[data-theme="light"] .cta-primary {
  background: linear-gradient(135deg, #1e40af, #7c3aed);
  color: white;
  font-weight: 700;
  box-shadow: 0 8px 25px rgba(30, 64, 175, 0.3);
  border: none;
  position: relative;
  overflow: hidden;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

[data-theme="light"] .cta-primary:hover {
  background: linear-gradient(135deg, #1e3a8a, #6d28d9);
  box-shadow: 0 12px 35px rgba(30, 64, 175, 0.4);
  transform: translateY(-3px) scale(1.02);
  color: white;
}

[data-theme="light"] .cta-primary:active {
  transform: translateY(-1px) scale(1.01);
  color: white;
}

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

[data-theme="light"] .cta-secondary {
  background: rgba(124, 58, 237, 0.9);
  border: 2px solid #7c3aed;
  color: white;
  font-weight: 600;
  backdrop-filter: blur(10px);
  position: relative;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
}

[data-theme="light"] .cta-secondary:hover {
  background: rgba(37, 99, 235, 0.9);
  border-color: #2563eb;
  color: white;
  box-shadow: 0 8px 25px rgba(37, 99, 235, 0.3);
  transform: translateY(-2px);
}

[data-theme="light"] .cta-primary,
[data-theme="light"] .cta-secondary {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  border-radius: 9999px;
}
</style>