<template>
  <MainLayout>
    <div class="career-profile-container">
      <div class="career-profile-content">
        <!-- 标题区域 -->
        <div class="title-section">
          <h1 class="title-text">
            <span class="title-gradient">
              职业画像
            </span>
          </h1>
          <p class="title-subtext">
            可视化职业路径，助力长远发展
          </p>
        </div>

        <!-- 个人资料卡片 -->
        <div class="profile-card">
          <div class="profile-card-bg"></div>
          <div class="profile-card-content">
            <div class="profile-layout">
              <div class="avatar-container">
                <div class="avatar-gradient">
                  <div class="avatar-inner">
                    A
                  </div>
                </div>
              </div>
              <div class="profile-info">
                <h2 class="profile-name">张三</h2>
                <p class="profile-title">高级前端开发工程师 | 5年经验</p>
                <div class="skill-tags">
                  <span class="skill-tag tag-blue">JavaScript</span>
                  <span class="skill-tag tag-purple">React</span>
                  <span class="skill-tag tag-pink">Vue.js</span>
                  <span class="skill-tag tag-cyan">TypeScript</span>
                </div>
              </div>
              <div class="profile-rating">
                <div class="rating-text">A+</div>
                <div class="rating-label">职业评级</div>
              </div>
            </div>
          </div>
        </div>

        <!-- 技能雷达图 -->
        <div class="charts-grid">
          <div class="radar-chart-card">
            <div class="radar-card-bg"></div>
            <div class="radar-card-content">
              <h3 class="chart-title title-purple">技能雷达图</h3>
              <div class="chart-container">
                <div class="radar-inner">
                  <!-- 雷达图背景 -->
                  <div class="radar-bg">
                    <div class="radar-level-1">
                      <div class="radar-level-2"></div>
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
              <div class="skills-grid">
                <div class="skill-item">
                  <span class="skill-label">前端开发</span>
                  <span class="skill-value value-blue">95%</span>
                </div>
                <div class="skill-item">
                  <span class="skill-label">后端能力</span>
                  <span class="skill-value value-purple">75%</span>
                </div>
                <div class="skill-item">
                  <span class="skill-label">设计思维</span>
                  <span class="skill-value value-pink">70%</span>
                </div>
                <div class="skill-item">
                  <span class="skill-label">项目管理</span>
                  <span class="skill-value value-cyan">65%</span>
                </div>
              </div>
            </div>
          </div>

          <!-- 职业发展路径 -->
          <div class="career-path-card">
            <div class="path-card-bg"></div>
            <div class="path-card-content">
              <h3 class="chart-title title-pink">职业发展路径</h3>
              <div class="career-steps">
                <div class="career-step">
                  <div class="step-badge badge-blue-purple">
                    1
                  </div>
                  <div class="step-info">
                    <h4 class="step-title">前端开发工程师</h4>
                    <p class="step-desc">1-3年经验 | 初级</p>
                  </div>
                </div>
                <div class="career-step">
                  <div class="step-badge badge-purple-pink">
                    2
                  </div>
                  <div class="step-info">
                    <h4 class="step-title">高级前端开发工程师</h4>
                    <p class="step-desc">3-5年经验 | 中级</p>
                  </div>
                </div>
                <div class="career-step">
                  <div class="step-badge badge-pink-orange">
                    3
                  </div>
                  <div class="step-info">
                    <h4 class="step-title">技术专家/架构师</h4>
                    <p class="step-desc">5-8年经验 | 高级</p>
                  </div>
                </div>
                <div class="career-step">
                  <div class="step-badge badge-orange-red">
                    4
                  </div>
                  <div class="step-info">
                    <h4 class="step-title">技术总监/CTO</h4>
                    <p class="step-desc">8+年经验 | 管理层</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 推荐学习路径 -->
        <div class="learning-card">
          <div class="learning-card-bg"></div>
          <div class="learning-card-content">
            <h3 class="chart-title title-cyan">推荐学习路径</h3>
            <div class="learning-items">
              <div class="learning-item">
                <div class="learning-header">
                  <h4 class="learning-title">微前端架构实践</h4>
                  <span class="learning-level">进阶课程</span>
                </div>
                <p class="learning-desc">学习现代微前端架构，提升大型项目管理能力</p>
                <div class="learning-tags">
                  <span class="learning-tag tag-blue">React</span>
                  <span class="learning-tag tag-purple">Vue</span>
                  <span class="learning-meta">预计 4周</span>
                </div>
              </div>
              <div class="learning-item">
                <div class="learning-header">
                  <h4 class="learning-title">TypeScript高级特性</h4>
                  <span class="learning-level">进阶课程</span>
                </div>
                <p class="learning-desc">深入掌握TypeScript高级类型系统和泛型编程</p>
                <div class="learning-tags">
                  <span class="learning-tag tag-pink">TypeScript</span>
                  <span class="learning-meta">预计 3周</span>
                </div>
              </div>
              <div class="learning-item">
                <div class="learning-header">
                  <h4 class="learning-title">性能优化实战</h4>
                  <span class="learning-level">进阶课程</span>
                </div>
                <p class="learning-desc">前端性能优化策略与实践，提升用户体验</p>
                <div class="learning-tags">
                  <span class="learning-tag tag-blue">Performance</span>
                  <span class="learning-meta">预计 3周</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 行动按钮 -->
        <div class="center-content">
          <button class="action-button">
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
.career-profile-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 3rem 1.5rem;
  position: relative;
}

.career-profile-content {
  max-width: 72rem;
  width: 100%;
}

.title-section {
  text-align: center;
  margin-bottom: 3rem;
}

.title-text {
  font-size: 3rem;
  font-weight: 900;
  margin-bottom: 1.5rem;
}

@media (min-width: 768px) {
  .title-text {
    font-size: 3.75rem;
  }
}

.title-gradient {
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background: linear-gradient(135deg, #60a5fa, #a855f7, #ec4899); /* from-blue-400 via-purple-500 to-pink-500 */
  font-size: 1em;
  display: block;
}

.title-subtext {
  font-size: 1.25rem;
  color: #9ca3af; /* text-gray-400 */
  max-width: 48rem;
  margin: 0 auto;
}

.profile-card {
  position: relative;
  padding: 2rem;
  border-radius: 1rem;
  background: linear-gradient(135deg, rgba(30, 58, 138, 0.3), rgba(147, 51, 234, 0.3)); /* from-blue-900/30 to-purple-900/30 */
  backdrop-filter: blur(12px);
  border: 1px solid rgba(59, 130, 246, 0.2); /* border-blue-500/20 */
  transition: all 0.5s cubic-bezier(0.4, 0, 0.2, 1);
  margin-bottom: 2rem;
}

.profile-card:hover {
  border-color: rgba(96, 165, 250, 0.5); /* border-blue-400/50 */
}

.profile-card-bg {
  position: absolute;
  inset: -1px;
  background: linear-gradient(90deg, #3b82f6, #8b5cf6); /* from-blue-500 to-purple-500 */
  border-radius: 1rem;
  opacity: 0;
  transition: opacity 0.5s ease;
}

.profile-card:hover .profile-card-bg {
  opacity: 0.3;
}

.profile-card-content {
  position: relative;
}

.profile-layout {
  display: flex;
  flex-direction: column;
}

@media (min-width: 768px) {
  .profile-layout {
    flex-direction: row;
  }
}

.profile-layout {
  align-items: center;
}

@media (min-width: 768px) {
  .profile-layout {
    align-items: flex-start;
  }
}

.profile-layout > *:last-child {
  margin-top: 1.5rem;
}

@media (min-width: 768px) {
  .profile-layout > *:last-child {
    margin-top: 0;
    margin-left: 2rem;
  }
}

.avatar-container {
  width: 8rem;
  height: 8rem;
  border-radius: 9999px;
  background: linear-gradient(135deg, #3b82f6, #8b5cf6); /* from-blue-500 to-purple-500 */
  padding: 0.25rem;
}

.avatar-inner {
  width: 100%;
  height: 100%;
  border-radius: 9999px;
  background-color: #000;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 2.5rem;
  font-weight: 700;
  color: white;
}

.profile-info {
  flex: 1;
  text-align: center;
}

@media (min-width: 768px) {
  .profile-info {
    text-align: left;
  }
}

.profile-name {
  font-size: 1.875rem;
  font-weight: 700;
  color: white;
  margin-bottom: 0.5rem;
}

.profile-title {
  color: #93c5fd; /* text-blue-300 */
  margin-bottom: 1rem;
}

.skill-tags {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 1rem;
  font-size: 0.875rem;
}

.skill-tag {
  padding: 0.25rem 0.75rem;
  border-radius: 9999px;
}

.tag-blue {
  background-color: rgba(59, 130, 246, 0.2); /* bg-blue-500/20 */
  color: #93c5fd; /* text-blue-300 */
}

.tag-purple {
  background-color: rgba(147, 51, 234, 0.2); /* bg-purple-500/20 */
  color: #d8b4fe; /* text-purple-300 */
}

.tag-pink {
  background-color: rgba(236, 72, 153, 0.2); /* bg-pink-500/20 */
  color: #f9a8d4; /* text-pink-300 */
}

.tag-cyan {
  background-color: rgba(6, 182, 212, 0.2); /* bg-cyan-500/20 */
  color: #67e8f9; /* text-cyan-300 */
}

.profile-rating {
  text-align: center;
}

.rating-text {
  font-size: 1.875rem;
  font-weight: 700;
  color: #fbbf24; /* text-yellow-400 */
}

.rating-label {
  color: #9ca3af; /* text-gray-400 */
}

.charts-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 2rem;
  margin-bottom: 2rem;
}

@media (min-width: 1024px) {
  .charts-grid {
    grid-template-columns: 1fr 1fr;
  }
}

.radar-chart-card, .career-path-card, .learning-card {
  position: relative;
  padding: 2rem;
  border-radius: 1rem;
  backdrop-filter: blur(12px);
  border: 1px solid;
  transition: all 0.5s cubic-bezier(0.4, 0, 0.2, 1);
}

.radar-chart-card {
  background: linear-gradient(135deg, rgba(147, 51, 234, 0.3), rgba(236, 72, 153, 0.3)); /* from-purple-900/30 to-pink-900/30 */
  border-color: rgba(168, 85, 247, 0.2); /* border-purple-500/20 */
}

.radar-chart-card:hover {
  border-color: rgba(192, 132, 252, 0.5); /* border-purple-400/50 */
}

.career-path-card {
  background: linear-gradient(135deg, rgba(236, 72, 153, 0.3), rgba(249, 115, 22, 0.3)); /* from-pink-900/30 to-orange-900/30 */
  border-color: rgba(236, 72, 153, 0.2); /* border-pink-500/20 */
}

.career-path-card:hover {
  border-color: rgba(249, 115, 22, 0.5); /* border-pink-400/50 */
}

.learning-card {
  background: linear-gradient(135deg, rgba(6, 182, 212, 0.3), rgba(30, 64, 175, 0.3)); /* from-cyan-900/30 to-blue-900/30 */
  border-color: rgba(6, 182, 212, 0.2); /* border-cyan-500/20 */
  margin-bottom: 2rem;
}

.learning-card:hover {
  border-color: rgba(6, 182, 212, 0.5); /* border-cyan-400/50 */
}

.radar-card-bg, .path-card-bg, .learning-card-bg {
  position: absolute;
  inset: -1px;
  border-radius: 1rem;
  opacity: 0;
  transition: opacity 0.5s ease;
}

.radar-card-bg {
  background: linear-gradient(90deg, #8b5cf6, #ec4899); /* from-purple-500 to-pink-500 */
}

.path-card-bg {
  background: linear-gradient(90deg, #ec4899, #f97316); /* from-pink-500 to-orange-500 */
}

.learning-card-bg {
  background: linear-gradient(90deg, #06b6d4, #3b82f6); /* from-cyan-500 to-blue-500 */
}

.radar-chart-card:hover .radar-card-bg,
.career-path-card:hover .path-card-bg,
.learning-card:hover .learning-card-bg {
  opacity: 0.3;
}

.radar-card-content, .path-card-content, .learning-card-content {
  position: relative;
}

.chart-title {
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 1.5rem;
}

.title-purple {
  color: #d8b4fe; /* text-purple-300 */
}

.title-pink {
  color: #f9a8d4; /* text-pink-300 */
}

.title-cyan {
  color: #67e8f9; /* text-cyan-300 */
}

.chart-container {
  height: 16rem;
  display: flex;
  align-items: center;
  justify-content: center;
}

.radar-inner {
  position: relative;
  width: 12rem;
  height: 12rem;
}

.radar-bg {
  position: absolute;
  inset: 0;
  border: 1px solid rgba(168, 85, 247, 0.3); /* border-purple-500/30 */
  border-radius: 9999px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.radar-level-1 {
  width: 75%;
  height: 75%;
  border: 1px solid rgba(168, 85, 247, 0.3); /* border-purple-500/30 */
  border-radius: 9999px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.radar-level-2 {
  width: 50%;
  height: 50%;
  border: 1px solid rgba(168, 85, 247, 0.3); /* border-purple-500/30 */
  border-radius: 9999px;
}

.radar-data {
  position: absolute;
  inset: 0;
}

.radar-axis-up {
  position: absolute;
  top: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 1px;
  height: 100%;
  background: linear-gradient(to top, transparent, #8b5cf6, transparent);
  opacity: 0.5;
}

.radar-axis-right {
  position: absolute;
  right: 0;
  top: 50%;
  transform: translateY(-50%);
  height: 1px;
  width: 100%;
  background: linear-gradient(to left, transparent, #8b5cf6, transparent);
  opacity: 0.5;
}

.radar-axis-down {
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 1px;
  height: 100%;
  background: linear-gradient(to bottom, transparent, #8b5cf6, transparent);
  opacity: 0.5;
}

.radar-axis-left {
  position: absolute;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  height: 1px;
  width: 100%;
  background: linear-gradient(to right, transparent, #8b5cf6, transparent);
  opacity: 0.5;
}

.radar-data-overlay {
  position: absolute;
  inset: 0;
}

.radar-data-fill {
  fill: url(#skillGradient);
  opacity: 0.6;
}

.radar-data-stroke {
  fill: transparent;
  stroke: url(#skillGradient);
  stroke-width: 2;
}

.skills-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
  margin-top: 1.5rem;
}

.skill-item {
  display: flex;
  justify-content: space-between;
}

.skill-label {
  color: #d1d5db; /* text-gray-300 */
}

.skill-value {
  font-weight: 700;
}

.value-blue {
  color: #60a5fa; /* text-blue-400 */
}

.value-purple {
  color: #a78bfa; /* text-purple-400 */
}

.value-pink {
  color: #f472b6; /* text-pink-400 */
}

.value-cyan {
  color: #22d3ee; /* text-cyan-400 */
}

.career-steps {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.career-step {
  display: flex;
  align-items: center;
  gap: 1rem;
}

@media (min-width: 640px) {
  .career-step {
    gap: 1rem;
  }
}

.step-badge {
  width: 3rem;
  height: 3rem;
  border-radius: 9999px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  color: white;
}

.badge-blue-purple {
  background: linear-gradient(90deg, #3b82f6, #8b5cf6); /* from-blue-500 to-purple-500 */
}

.badge-purple-pink {
  background: linear-gradient(90deg, #8b5cf6, #ec4899); /* from-purple-500 to-pink-500 */
}

.badge-pink-orange {
  background: linear-gradient(90deg, #ec4899, #f97316); /* from-pink-500 to-orange-500 */
}

.badge-orange-red {
  background: linear-gradient(90deg, #f97316, #ef4444); /* from-orange-500 to-red-500 */
}

.step-info {
  flex: 1;
}

.step-title {
  font-size: 1.125rem;
  font-weight: 700;
  color: white;
}

.step-desc {
  color: #9ca3af; /* text-gray-400 */
}

.learning-items {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.learning-item {
  padding: 1rem;
  background-color: rgba(0, 0, 0, 0.2); /* bg-black/20 */
  border-radius: 0.75rem;
}

.learning-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 0.5rem;
}

.learning-title {
  font-size: 1.125rem;
  font-weight: 700;
  color: white;
}

.learning-level {
  color: #22d3ee; /* text-cyan-400 */
  font-size: 0.875rem;
}

.learning-desc {
  color: #9ca3af; /* text-gray-400 */
  margin-bottom: 0.75rem;
}

.learning-tags {
  display: flex;
  align-items: center;
  gap: 1rem;
  font-size: 0.875rem;
}

.learning-tag {
  padding: 0.25rem 0.5rem;
  border-radius: 0.25rem;
  font-size: 0.875rem;
}

.learning-tag.tag-blue {
  background-color: rgba(59, 130, 246, 0.2); /* bg-blue-500/20 */
  color: #93c5fd; /* text-blue-300 */
}

.learning-tag.tag-purple {
  background-color: rgba(147, 51, 234, 0.2); /* bg-purple-500/20 */
  color: #d8b4fe; /* text-purple-300 */
}

.learning-tag.tag-pink {
  background-color: rgba(236, 72, 153, 0.2); /* bg-pink-500/20 */
  color: #f9a8d4; /* text-pink-300 */
}

.learning-meta {
  color: #6b7280; /* text-gray-500 */
}

.center-content {
  text-align: center;
}

.action-button {
  padding: 1rem 2rem;
  background: linear-gradient(90deg, #3b82f6, #8b5cf6); /* from-blue-500 to-purple-500 */
  border-radius: 0.75rem;
  font-weight: 700;
  font-size: 1.125rem;
  color: white;
  border: none;
  cursor: pointer;
  transition: all 0.3s ease;
}

.action-button:hover {
  background: linear-gradient(90deg, #2563eb, #7c3aed); /* hover:from-blue-600 hover:to-purple-600 */
  transform: scale(1.05);
  box-shadow: 0 25px 50px -12px rgba(59, 130, 246, 0.5); /* hover:shadow-2xl hover:shadow-blue-500/50 */
}

/* 亮色主题样式 */
[data-theme="light"] .profile-card,
[data-theme="light"] .radar-chart-card,
[data-theme="light"] .career-path-card,
[data-theme="light"] .learning-card {
  background: linear-gradient(135deg, var(--color-bg-card), var(--color-bg-card)) !important;
  border: 1px solid var(--color-border-light) !important;
}

[data-theme="light"] .profile-card,
[data-theme="light"] .radar-chart-card,
[data-theme="light"] .career-path-card,
[data-theme="light"] .learning-card {
  backdrop-filter: blur(20px) !important;
  -webkit-backdrop-filter: blur(20px) !important;
}

[data-theme="light"] .skill-label {
  color: var(--color-text-secondary) !important;
}

[data-theme="light"] .step-desc,
[data-theme="light"] .title-subtext,
[data-theme="light"] .learning-desc {
  color: var(--color-text-muted) !important;
}

[data-theme="light"] .profile-name,
[data-theme="light"] .step-title {
  color: var(--color-text-primary) !important;
}

[data-theme="light"] .learning-item {
  background: rgba(255, 255, 255, 0.5) !important;
}

[data-theme="light"] .action-button {
  background: linear-gradient(135deg, var(--color-stellar-blue), var(--color-nebula-purple)) !important;
  color: white !important;
  font-weight: 700 !important;
  box-shadow: 0 8px 25px rgba(30, 64, 175, 0.3) !important;
}

[data-theme="light"] .action-button:hover {
  background: linear-gradient(135deg, var(--color-stellar-blue), var(--color-nebula-purple)) !important;
  transform: translateY(-3px) scale(1.02) !important;
  box-shadow: 0 12px 35px rgba(30, 64, 175, 0.4) !important;
}

/* 背景渐变文字 */
.title-gradient {
  background-size: 200% 200%;
  animation: gradient 3s ease infinite;
}

@keyframes gradient {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

/* 悬停动效 */
.profile-card,
.radar-chart-card,
.career-path-card,
.learning-card {
  position: relative;
  transition: all 0.3s ease;
}

.profile-card:hover,
.radar-chart-card:hover,
.career-path-card:hover,
.learning-card:hover {
  transform: translateY(-5px);
}

/* 个人头像样式 */
.avatar-container {
  background: linear-gradient(135deg, var(--color-stellar-blue), var(--color-nebula-purple));
}

/* 亮色主题头像 */
[data-theme="light"] .avatar-container {
  background: linear-gradient(135deg, var(--color-stellar-blue), var(--color-nebula-purple));
}
</style>