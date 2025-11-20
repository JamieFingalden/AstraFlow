<template>
  <MainLayout>
    <div class="job-analysis-container">
      <div class="job-analysis-content">
        <!-- 标题区域 -->
        <div class="title-section">
          <h1 class="title-text">
            <span class="title-gradient">
              岗位分析
            </span>
          </h1>
          <p class="title-subtext">
            深度解析职位需求，智能匹配最适合的机会
          </p>
        </div>

        <!-- 分析工具卡片 -->
        <div class="analysis-grid">
          <!-- 职位匹配度分析 -->
          <div class="analysis-card card-blue">
            <div class="card-bg card-bg-blue"></div>
            <div class="card-content">
              <div class="card-emoji">🎯</div>
              <h3 class="card-title title-blue">职位匹配度分析</h3>
              <p class="card-description">AI智能分析你的简历与目标职位的匹配程度</p>
              <div class="progress-section">
                <div class="progress-item">
                  <div class="progress-label">
                    <span class="label-text">技能匹配度</span>
                    <span class="label-value value-blue">85%</span>
                  </div>
                  <div class="progress-bar">
                    <div class="progress-fill progress-fill-blue" style="width: 85%"></div>
                  </div>
                </div>
                <div class="progress-item">
                  <div class="progress-label">
                    <span class="label-text">经验匹配度</span>
                    <span class="label-value value-purple">78%</span>
                  </div>
                  <div class="progress-bar">
                    <div class="progress-fill progress-fill-purple" style="width: 78%"></div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- 技能差距分析 -->
          <div class="analysis-card card-purple">
            <div class="card-bg card-bg-purple"></div>
            <div class="card-content">
              <div class="card-emoji">🔍</div>
              <h3 class="card-title title-purple">技能差距分析</h3>
              <p class="card-description">识别与目标职位的关键技能差距</p>
              <div class="skills-gap-list">
                <div class="skill-item">
                  <div class="skill-status status-green"></div>
                  <span class="skill-text">JavaScript (已具备)</span>
                </div>
                <div class="skill-item">
                  <div class="skill-status status-yellow"></div>
                  <span class="skill-text">TypeScript (需提升)</span>
                </div>
                <div class="skill-item">
                  <div class="skill-status status-red"></div>
                  <span class="skill-text">React Native (需学习)</span>
                </div>
                <div class="skill-item">
                  <div class="skill-status status-red"></div>
                  <span class="skill-text">GraphQL (需学习)</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 市场分析 -->
        <div class="market-analysis-card">
          <div class="card-bg card-bg-pink"></div>
          <div class="card-content">
            <div class="card-emoji">📊</div>
            <h3 class="card-title title-pink">市场薪资分析</h3>
            <div class="salary-grid">
              <div class="salary-item">
                <div class="salary-value value-blue">¥25K</div>
                <div class="salary-label">平均薪资</div>
              </div>
              <div class="salary-item">
                <div class="salary-value value-purple">¥35K</div>
                <div class="salary-label">最高薪资</div>
              </div>
              <div class="salary-item">
                <div class="salary-value value-pink">¥18K</div>
                <div class="salary-label">最低薪资</div>
              </div>
            </div>
          </div>
        </div>

        <!-- 分析输入区域 -->
        <div class="input-card">
          <div class="input-header">
            <h3 class="input-title">开始岗位分析</h3>
            <p class="input-subtitle">输入职位信息，获取智能分析报告</p>
          </div>

          <div class="input-form">
            <div class="input-group">
              <label class="input-label">职位名称</label>
              <input
                type="text"
                placeholder="例如：高级前端开发工程师"
                class="input-field"
              />
            </div>

            <div class="input-group">
              <label class="input-label">公司要求</label>
              <textarea
                placeholder="粘贴职位描述或要求..."
                rows="6"
                class="input-textarea"
              ></textarea>
            </div>

            <button class="submit-button">
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
.job-analysis-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 3rem 1.5rem;
  position: relative;
}

.job-analysis-content {
  max-width: 56rem;
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

.analysis-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 2rem;
  margin-bottom: 3rem;
}

@media (min-width: 768px) {
  .analysis-grid {
    grid-template-columns: 1fr 1fr;
  }
}

.analysis-card, .market-analysis-card {
  position: relative;
  padding: 2rem;
  border-radius: 1rem;
  backdrop-filter: blur(12px);
  border: 1px solid;
  transition: all 0.5s cubic-bezier(0.4, 0, 0.2, 1);
  transform: scale(1);
}

.analysis-card {
  background: linear-gradient(135deg, rgba(30, 58, 138, 0.3), rgba(147, 51, 234, 0.3)); /* from-blue-900/30 to-purple-900/30 */
  border-color: rgba(59, 130, 246, 0.2); /* border-blue-500/20 */
}

.card-blue:hover {
  border-color: rgba(96, 165, 250, 0.5); /* hover:border-blue-400/50 */
}

.analysis-card.card-purple {
  background: linear-gradient(135deg, rgba(147, 51, 234, 0.3), rgba(236, 72, 153, 0.3)); /* from-purple-900/30 to-pink-900/30 */
  border-color: rgba(168, 85, 247, 0.2); /* border-purple-500/20 */
}

.card-purple:hover {
  border-color: rgba(192, 132, 252, 0.5); /* hover:border-purple-400/50 */
}

.market-analysis-card {
  background: linear-gradient(135deg, rgba(236, 72, 153, 0.3), rgba(249, 115, 22, 0.3)); /* from-pink-900/30 to-orange-900/30 */
  border-color: rgba(236, 72, 153, 0.2); /* border-pink-500/20 */
  margin-bottom: 3rem;
}

.market-analysis-card:hover {
  border-color: rgba(249, 115, 22, 0.5); /* hover:border-pink-400/50 */
}

.card-bg {
  position: absolute;
  inset: -1px;
  border-radius: 1rem;
  opacity: 0;
  transition: opacity 0.5s ease;
}

.card-bg-blue {
  background: linear-gradient(90deg, #3b82f6, #8b5cf6); /* from-blue-500 to-purple-500 */
}

.card-bg-purple {
  background: linear-gradient(90deg, #8b5cf6, #ec4899); /* from-purple-500 to-pink-500 */
}

.card-bg-pink {
  background: linear-gradient(90deg, #ec4899, #f97316); /* from-pink-500 to-orange-500 */
}

.analysis-card:hover .card-bg,
.market-analysis-card:hover .card-bg {
  opacity: 0.3;
}

.card-content {
  position: relative;
}

.card-emoji {
  font-size: 2.5rem;
  margin-bottom: 1rem;
}

.card-title {
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 1rem;
}

.title-blue {
  color: #93c5fd; /* text-blue-300 */
}

.title-purple {
  color: #d8b4fe; /* text-purple-300 */
}

.title-pink {
  color: #f9a8d4; /* text-pink-300 */
}

.card-description {
  color: #9ca3af; /* text-gray-400 */
  margin-bottom: 1.5rem;
}

.progress-section {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.progress-item {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.progress-label {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.label-text {
  color: #d1d5db; /* text-gray-300 */
}

.label-value {
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

.progress-bar {
  width: 100%;
  background-color: #374151; /* bg-gray-700 */
  border-radius: 0.5rem;
  height: 0.5rem;
}

.progress-fill {
  height: 100%;
  border-radius: 0.5rem;
}

.progress-fill-blue {
  background: linear-gradient(90deg, #3b82f6, #8b5cf6); /* from-blue-500 to-purple-500 */
}

.progress-fill-purple {
  background: linear-gradient(90deg, #8b5cf6, #ec4899); /* from-purple-500 to-pink-500 */
}

.skills-gap-list {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.skill-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.skill-status {
  width: 0.5rem;
  height: 0.5rem;
  border-radius: 50%;
}

.status-green {
  background-color: #22c55e; /* bg-green-500 */
}

.status-yellow {
  background-color: #eab308; /* bg-yellow-500 */
}

.status-red {
  background-color: #ef4444; /* bg-red-500 */
}

.skill-text {
  color: #d1d5db; /* text-gray-300 */
}

.salary-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1.5rem;
}

@media (min-width: 768px) {
  .salary-grid {
    grid-template-columns: 1fr 1fr 1fr;
  }
}

.salary-item {
  text-align: center;
  padding: 1rem;
  border-radius: 0.75rem;
  background-color: rgba(0, 0, 0, 0.2); /* bg-black/20 */
}

.salary-value {
  font-size: 1.875rem;
  font-weight: 700;
  margin-bottom: 0.25rem;
}

.salary-label {
  color: #9ca3af; /* text-gray-400 */
}

.input-card {
  background: linear-gradient(135deg, rgba(30, 58, 138, 0.2), rgba(147, 51, 234, 0.2)); /* from-blue-900/20 to-purple-900/20 */
  backdrop-filter: blur(12px);
  border: 1px solid rgba(59, 130, 246, 0.2); /* border-blue-500/20 */
  border-radius: 1rem;
  padding: 2rem;
}

.input-header {
  text-align: center;
  margin-bottom: 1.5rem;
}

.input-title {
  font-size: 1.5rem;
  font-weight: 700;
  color: #e5e7eb; /* text-gray-200 */
  margin-bottom: 0.5rem;
}

.input-subtitle {
  color: #9ca3af; /* text-gray-400 */
}

.input-form {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.input-group {
  display: flex;
  flex-direction: column;
}

.input-label {
  display: block;
  color: #d1d5db; /* text-gray-300 */
  margin-bottom: 0.75rem;
}

.input-field, .input-textarea {
  width: 100%;
  padding: 0.75rem 1rem;
  background-color: rgba(0, 0, 0, 0.3); /* bg-black/30 */
  border: 1px solid rgba(59, 130, 246, 0.3); /* border-blue-500/30 */
  border-radius: 0.75rem;
  color: white;
  font-size: 1rem;
}

.input-field::placeholder, .input-textarea::placeholder {
  color: #6b7280; /* placeholder-gray-500 */
}

.input-field:focus, .input-textarea:focus {
  outline: none;
  border-color: #60a5fa; /* focus:border-blue-400 */
}

.input-textarea {
  resize: none;
}

.submit-button {
  width: 100%;
  padding: 1rem 0;
  background: linear-gradient(90deg, #3b82f6, #8b5cf6); /* from-blue-500 to-purple-500 */
  border-radius: 0.75rem;
  font-weight: 700;
  font-size: 1.125rem;
  color: white;
  border: none;
  cursor: pointer;
  transition: all 0.3s ease;
}

.submit-button:hover {
  background: linear-gradient(90deg, #2563eb, #7c3aed); /* hover:from-blue-600 hover:to-purple-600 */
  transform: scale(1.05);
  box-shadow: 0 25px 50px -12px rgba(59, 130, 246, 0.5); /* hover:shadow-2xl hover:shadow-blue-500/50 */
}

/* 亮色主题样式 */
[data-theme="light"] .analysis-card,
[data-theme="light"] .market-analysis-card {
  background: linear-gradient(135deg, var(--color-bg-card), var(--color-bg-card)) !important;
  border: 1px solid var(--color-border-light) !important;
}

[data-theme="light"] .analysis-card,
[data-theme="light"] .market-analysis-card,
[data-theme="light"] .input-card {
  backdrop-filter: blur(20px) !important;
  -webkit-backdrop-filter: blur(20px) !important;
}

[data-theme="light"] .input-label,
[data-theme="light"] .label-text,
[data-theme="light"] .skill-text {
  color: var(--color-text-secondary) !important;
}

[data-theme="light"] .title-subtext,
[data-theme="light"] .card-description,
[data-theme="light"] .salary-label,
[data-theme="light"] .input-subtitle {
  color: var(--color-text-muted) !important;
}

[data-theme="light"] .input-field,
[data-theme="light"] .input-textarea {
  background: rgba(255, 255, 255, 0.7) !important; /* bg-black/30 */
  border: 1px solid var(--color-border-light) !important; /* border-blue-500/30 */
  color: var(--color-text-primary) !important;
}

[data-theme="light"] .input-field::placeholder,
[data-theme="light"] .input-textarea::placeholder {
  color: var(--color-text-muted) !important; /* placeholder-gray-500 */
}

[data-theme="light"] .input-field:focus,
[data-theme="light"] .input-textarea:focus {
  border-color: var(--color-stellar-blue) !important; /* focus:border-blue-400 */
}

[data-theme="light"] .salary-item {
  background: rgba(255, 255, 255, 0.5) !important; /* bg-black/20 */
}

[data-theme="light"] .submit-button {
  background: linear-gradient(135deg, var(--color-stellar-blue), var(--color-nebula-purple)) !important;
  color: white !important;
  font-weight: 700 !important;
  box-shadow: 0 8px 25px rgba(30, 64, 175, 0.3) !important;
}

[data-theme="light"] .submit-button:hover {
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
.analysis-card,
.market-analysis-card {
  position: relative;
  transition: all 0.3s ease;
}

.analysis-card:hover,
.market-analysis-card:hover {
  transform: translateY(-5px) scale(1.05);
}
</style>