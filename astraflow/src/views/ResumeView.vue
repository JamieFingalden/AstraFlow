<template>
  <div class="resume-page">
    <!-- Stellar Background -->
    <StellarBackground />

    <!-- Page Header -->
    <header class="page-header">
      <div class="container">
        <div class="header-left">
          <button class="back-button" @click="goHome">
            <svg class="icon" viewBox="0 0 24 24" fill="none" stroke="currentColor">
              <path d="M19 12H5M12 19l-7-7 7-7"/>
            </svg>
            <span>返回首页</span>
          </button>
        </div>

        <div class="header-center">
          <h1 class="page-title text-stellar">AI智能简历生成</h1>
          <p class="page-subtitle">让AI为你打造专业简历，在星流中闪耀</p>
        </div>

        <div class="header-right">
          <ThemeToggle />
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="main-content">
      <div class="container">
        <div class="resume-layout">
          <!-- Left Side: Input Section -->
          <section class="input-section">
            <div class="section-card glass">
              <div class="section-header">
                <div class="section-icon glow-blue">
                  <svg class="icon" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                    <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                    <polyline points="14 2 14 8 20 8"></polyline>
                    <line x1="16" y1="13" x2="8" y2="13"></line>
                    <line x1="16" y1="17" x2="8" y2="17"></line>
                    <polyline points="10 9 9 9 8 9"></polyline>
                  </svg>
                </div>
                <h2 class="section-title">创建你的AI简历</h2>
              </div>

              <form class="resume-form" @submit.prevent="generateResume">
                <!-- Name Input -->
                <div class="form-group">
                  <label class="form-label">姓名</label>
                  <input
                    v-model="formData.name"
                    type="text"
                    class="form-input glass"
                    placeholder="请输入您的姓名"
                    required
                  />
                </div>

                <!-- Position Input -->
                <div class="form-group">
                  <label class="form-label">目标职位</label>
                  <input
                    v-model="formData.position"
                    type="text"
                    class="form-input glass"
                    placeholder="例如：Java后端开发工程师"
                    required
                  />
                </div>

                <!-- Skills Textarea -->
                <div class="form-group">
                  <label class="form-label">技能标签</label>
                  <textarea
                    v-model="formData.skills"
                    class="form-textarea glass"
                    rows="3"
                    placeholder="请输入您的技能，用逗号分隔&#10;例如：Spring Boot, MySQL, Redis, Docker"
                    required
                  ></textarea>
                </div>

                <!-- Experience Textarea -->
                <div class="form-group">
                  <label class="form-label">项目经验</label>
                  <textarea
                    v-model="formData.experience"
                    class="form-textarea glass"
                    rows="4"
                    placeholder="请描述您的项目经验&#10;例如：负责公司账单系统后端架构设计与接口实现"
                    required
                  ></textarea>
                </div>

                <!-- Upload Area -->
                <div class="upload-section">
                  <div class="upload-button glass" @click="handleMockUpload">
                    <svg class="icon" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                      <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
                      <polyline points="17 8 12 3 7 8"></polyline>
                      <line x1="12" y1="3" x2="12" y2="15"></line>
                    </svg>
                    <span>上传头像 / 简历文件</span>
                  </div>
                </div>

                <!-- Generate Button -->
                <button
                  type="submit"
                  class="generate-button"
                  :disabled="isGenerating"
                  :class="{ 'generating': isGenerating }"
                >
                  <span v-if="isGenerating" class="loading-content">
                    <div class="loading-spinner"></div>
                    <span>AI正在生成中...</span>
                  </span>
                  <span v-else class="button-content">
                    <span>生成简历</span>
                    <span class="button-star">✦</span>
                  </span>
                </button>
              </form>
            </div>
          </section>

          <!-- Right Side: Display Section -->
          <section class="display-section">
            <div class="section-card glass">
              <div class="section-header">
                <div class="section-icon glow-purple">
                  <svg class="icon" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                    <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"></path>
                  </svg>
                </div>
                <h2 class="section-title">AI生成结果</h2>
              </div>

              <!-- Resume Display Area -->
              <div class="resume-display">
                <div v-if="!isGenerating && !generatedResume" class="waiting-state">
                  <div class="waiting-icon animate-float">
                    <svg class="icon" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                      <circle cx="12" cy="12" r="10"></circle>
                      <polyline points="12 6 12 12 16 14"></polyline>
                    </svg>
                  </div>
                  <p class="waiting-text">等待生成中...</p>
                </div>

                <div v-else-if="isGenerating" class="generating-state">
                  <div class="ai-animation">
                    <div class="neural-network">
                      <div class="node" v-for="i in 8" :key="i" :style="getNodeStyle(i)"></div>
                      <div class="connection" v-for="i in 12" :key="`conn-${i}`" :style="getConnectionStyle(i)"></div>
                    </div>
                    <div class="generating-text">
                      <span class="typewriter-text">{{ typewriterText }}</span>
                      <span class="cursor">|</span>
                    </div>
                  </div>
                </div>

                <div v-else-if="generatedResume" class="result-state">
                  <div class="resume-card glass-dark">
                    <div class="resume-header">
                      <div class="resume-avatar">
                        <div class="avatar-placeholder glow-blue">
                          <svg class="icon" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                            <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                            <circle cx="12" cy="7" r="4"></circle>
                          </svg>
                        </div>
                      </div>
                      <div class="resume-title">
                        <h3 class="resume-name animate-fade-in-up">{{ generatedResume.name }}</h3>
                        <p class="resume-position animate-fade-in-up" style="animation-delay: 0.2s">
                          {{ generatedResume.position }}
                        </p>
                      </div>
                    </div>

                    <div class="resume-content">
                      <div class="resume-section">
                        <h4 class="section-label animate-fade-in-up" style="animation-delay: 0.4s">核心技能</h4>
                        <div class="skills-container animate-fade-in-up" style="animation-delay: 0.6s">
                          <span
                            v-for="(skill, index) in generatedResume.skills"
                            :key="skill"
                            class="skill-tag"
                            :style="{ animationDelay: `${0.8 + index * 0.1}s` }"
                          >
                            {{ skill }}
                          </span>
                        </div>
                      </div>

                      <div class="resume-section">
                        <h4 class="section-label animate-fade-in-up" style="animation-delay: 1.2s">项目经验</h4>
                        <div class="experience-list">
                          <div
                            v-for="(exp, index) in generatedResume.experience"
                            :key="index"
                            class="experience-item animate-fade-in-up"
                            :style="{ animationDelay: `${1.4 + index * 0.2}s` }"
                          >
                            <div class="exp-bullet"></div>
                            <span class="exp-text">{{ exp }}</span>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </section>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import StellarBackground from '../components/ui/StellarBackground.vue'
import ThemeToggle from '../components/ui/ThemeToggle.vue'

const router = useRouter()
const isGenerating = ref(false)
const generatedResume = ref(null)
const typewriterText = ref('')

const formData = reactive({
  name: '',
  position: '',
  skills: '',
  experience: ''
})

const goHome = () => {
  router.push('/')
}

const handleMockUpload = () => {
  // Mock upload functionality
  console.log('Upload clicked - mock functionality')
}

const typewriterPhrases = [
  'AI正在分析您的信息...',
  '正在优化简历结构...',
  '匹配最佳词汇表达...',
  '生成专业简历格式...',
  '即将完成...'
]

const generateResume = async () => {
  isGenerating.value = true
  generatedResume.value = null
  typewriterText.value = ''

  // Typewriter effect
  let phraseIndex = 0
  let charIndex = 0
  let currentPhrase = ''

  const typewriterInterval = setInterval(() => {
    if (phraseIndex < typewriterPhrases.length) {
      currentPhrase = typewriterPhrases[phraseIndex]

      if (charIndex < currentPhrase.length) {
        typewriterText.value += currentPhrase[charIndex]
        charIndex++
      } else {
        typewriterText.value = ''
        phraseIndex++
        charIndex = 0
      }
    } else {
      clearInterval(typewriterInterval)
    }
  }, 50)

  // Simulate AI generation delay
  setTimeout(async () => {
    try {
      // Fetch mock data
      const response = await fetch('/mock/resumeData.json')
      const mockData = await response.json()

      // Use mock data or form data
      generatedResume.value = {
        name: formData.name || mockData.name,
        position: formData.position || mockData.position,
        skills: formData.skills ? formData.skills.split(',').map(s => s.trim()) : mockData.skills,
        experience: formData.experience
          ? formData.experience.split('\n').filter(line => line.trim())
          : mockData.experience
      }
    } catch (error) {
      console.error('Error generating resume:', error)
    } finally {
      isGenerating.value = false
      clearInterval(typewriterInterval)
    }
  }, 3000)
}

// Animation helpers
const getNodeStyle = (index) => {
  const positions = [
    { top: '20%', left: '20%' },
    { top: '20%', right: '20%' },
    { top: '50%', left: '10%' },
    { top: '50%', right: '10%' },
    { top: '80%', left: '20%' },
    { top: '80%', right: '20%' },
    { top: '40%', left: '50%' },
    { top: '60%', left: '50%' }
  ]
  return positions[index - 1] || {}
}

const getConnectionStyle = (index) => {
  const connections = [
    { width: '30%', top: '25%', left: '25%', transform: 'rotate(45deg)' },
    { width: '30%', top: '25%', right: '25%', transform: 'rotate(-45deg)' },
    { width: '40%', top: '50%', left: '15%', transform: 'rotate(0deg)' },
    { width: '40%', top: '50%', right: '15%', transform: 'rotate(0deg)' },
    { width: '30%', top: '75%', left: '25%', transform: 'rotate(-45deg)' },
    { width: '30%', top: '75%', right: '25%', transform: 'rotate(45deg)' }
  ]
  return connections[index % connections.length] || {}
}
</script>

<style scoped>
.resume-page {
  min-height: 100vh;
  position: relative;
  color: var(--color-starlight);
}

/* Page Header */
.page-header {
  padding: var(--space-2xl) 0;
  position: relative;
}

.page-header .container {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: var(--space-lg);
}

.header-left {
  flex-shrink: 0;
}

.header-center {
  flex: 1;
  text-align: center;
}

.header-right {
  flex-shrink: 0;
}

.back-button {
  display: flex;
  align-items: center;
  gap: var(--space-sm);
  padding: var(--space-sm) var(--space-md);
  background: rgba(59, 130, 246, 0.1);
  border: 1px solid rgba(59, 130, 246, 0.3);
  border-radius: var(--radius-lg);
  color: var(--color-stellar-blue);
  cursor: pointer;
  transition: all var(--duration-normal) ease;
}

.back-button:hover {
  background: rgba(59, 130, 246, 0.2);
  transform: translateX(-2px);
}

.back-button .icon {
  width: 20px;
  height: 20px;
}

.page-title {
  font-size: 2.5rem;
  font-weight: 800;
  margin-bottom: var(--space-md);
}

.page-subtitle {
  font-size: 1.1rem;
  opacity: 0.7;
  font-weight: 300;
}

/* Light Theme Header Styles */
[data-theme="light"] .back-button {
  background: rgba(59, 130, 246, 0.08);
  border-color: rgba(59, 130, 246, 0.2);
  color: var(--color-stellar-blue);
}

[data-theme="light"] .back-button:hover {
  background: rgba(59, 130, 246, 0.15);
}

[data-theme="light"] .page-subtitle {
  opacity: 0.6;
  color: var(--color-starlight);
}

/* Main Content */
.main-content {
  padding-bottom: var(--space-3xl);
}

.resume-layout {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: var(--space-2xl);
  align-items: start;
}

/* Section Cards */
.section-card {
  padding: var(--space-2xl);
  border-radius: var(--radius-2xl);
  transition: all var(--duration-normal) ease;
}

.section-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 20px 40px rgba(59, 130, 246, 0.15);
}

.section-header {
  display: flex;
  align-items: center;
  gap: var(--space-md);
  margin-bottom: var(--space-xl);
}

.section-icon {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(59, 130, 246, 0.1);
  border: 1px solid rgba(59, 130, 246, 0.3);
}

.section-icon .icon {
  width: 24px;
  height: 24px;
  stroke: currentColor;
}

.section-title {
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--color-starlight);
}

/* Form Styles */
.resume-form {
  display: flex;
  flex-direction: column;
  gap: var(--space-lg);
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: var(--space-sm);
}

.form-label {
  font-weight: 600;
  color: var(--color-stellar-blue);
  font-size: 0.9rem;
}

.form-input,
.form-textarea {
  padding: var(--space-md);
  border: 1px solid rgba(59, 130, 246, 0.3);
  border-radius: var(--radius-lg);
  background: rgba(15, 23, 42, 0.6);
  color: var(--color-starlight);
  font-size: 1rem;
  transition: all var(--duration-normal) ease;
}

.form-input:focus,
.form-textarea:focus {
  outline: none;
  border-color: var(--color-stellar-blue);
  box-shadow: 0 0 20px rgba(59, 130, 246, 0.2);
}

.form-input::placeholder,
.form-textarea::placeholder {
  color: var(--color-starlight);
  opacity: 0.4;
}

/* Upload Section */
.upload-section {
  margin: var(--space-lg) 0;
}

.upload-button {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: var(--space-sm);
  padding: var(--space-lg);
  border: 2px dashed rgba(59, 130, 246, 0.4);
  border-radius: var(--radius-lg);
  cursor: pointer;
  transition: all var(--duration-normal) ease;
}

.upload-button:hover {
  border-color: var(--color-stellar-blue);
  background: rgba(59, 130, 246, 0.1);
}

.upload-button .icon {
  width: 24px;
  height: 24px;
  opacity: 0.7;
}

/* Generate Button */
.generate-button {
  padding: var(--space-lg) var(--space-xl);
  background: var(--gradient-stellar);
  border: none;
  border-radius: var(--radius-xl);
  color: white;
  font-size: 1.1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all var(--duration-normal) ease;
  position: relative;
  overflow: hidden;
}

.generate-button:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 10px 30px rgba(59, 130, 246, 0.4);
}

.generate-button:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.button-content {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: var(--space-sm);
}

.button-star {
  font-size: 1.2rem;
  animation: pulseGlow 2s ease-in-out infinite;
}

.loading-content {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: var(--space-md);
}

.loading-spinner {
  width: 20px;
  height: 20px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-top: 2px solid white;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

/* Resume Display */
.resume-display {
  min-height: 400px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.waiting-state {
  text-align: center;
}

.waiting-icon {
  margin-bottom: var(--space-lg);
  opacity: 0.3;
}

.waiting-icon .icon {
  width: 48px;
  height: 48px;
}

.waiting-text {
  font-size: 1.1rem;
  opacity: 0.5;
}

/* AI Animation */
.generating-state {
  text-align: center;
  width: 100%;
}

.neural-network {
  position: relative;
  width: 200px;
  height: 150px;
  margin: 0 auto var(--space-lg);
}

.node {
  position: absolute;
  width: 12px;
  height: 12px;
  background: var(--color-stellar-blue);
  border-radius: 50%;
  animation: pulseGlow 2s ease-in-out infinite;
}

.connection {
  position: absolute;
  height: 1px;
  background: linear-gradient(90deg, transparent, var(--color-nebula-purple), transparent);
  animation: shimmer 2s linear infinite;
}

.generating-text {
  font-family: 'Courier New', monospace;
  font-size: 0.9rem;
  color: var(--color-nebula-purple);
}

.cursor {
  animation: blink 1s infinite;
}

@keyframes blink {
  0%, 50% { opacity: 1; }
  51%, 100% { opacity: 0; }
}

/* Resume Result */
.result-state {
  width: 100%;
}

.resume-card {
  padding: var(--space-xl);
  border-radius: var(--radius-xl);
}

.resume-header {
  display: flex;
  align-items: center;
  gap: var(--space-lg);
  margin-bottom: var(--space-xl);
  padding-bottom: var(--space-lg);
  border-bottom: 1px solid rgba(59, 130, 246, 0.2);
}

.resume-avatar {
  flex-shrink: 0;
}

.avatar-placeholder {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  background: rgba(59, 130, 246, 0.1);
  border: 2px solid rgba(59, 130, 246, 0.3);
}

.avatar-placeholder .icon {
  width: 40px;
  height: 40px;
  stroke: var(--color-stellar-blue);
}

.resume-title {
  flex: 1;
}

.resume-name {
  font-size: 1.8rem;
  font-weight: 700;
  margin-bottom: var(--space-sm);
}

.resume-position {
  font-size: 1.1rem;
  color: var(--color-nebula-purple);
  opacity: 0.8;
}

.resume-section {
  margin-bottom: var(--space-xl);
}

.section-label {
  font-size: 1.1rem;
  font-weight: 600;
  margin-bottom: var(--space-md);
  color: var(--color-stellar-blue);
}

.skills-container {
  display: flex;
  flex-wrap: wrap;
  gap: var(--space-sm);
}

.skill-tag {
  padding: var(--space-sm) var(--space-md);
  background: rgba(59, 130, 246, 0.1);
  border: 1px solid rgba(59, 130, 246, 0.3);
  border-radius: var(--radius-lg);
  font-size: 0.9rem;
  animation: fadeInUp 0.6s ease-out forwards;
  opacity: 0;
}

.experience-list {
  display: flex;
  flex-direction: column;
  gap: var(--space-md);
}

.experience-item {
  display: flex;
  align-items: flex-start;
  gap: var(--space-md);
  animation: fadeInUp 0.6s ease-out forwards;
  opacity: 0;
}

.exp-bullet {
  width: 8px;
  height: 8px;
  background: var(--color-nebula-purple);
  border-radius: 50%;
  margin-top: 8px;
  flex-shrink: 0;
}

.exp-text {
  flex: 1;
  line-height: 1.6;
  opacity: 0.8;
}

/* Light Theme Styles */
[data-theme="light"] .section-title {
  color: var(--color-starlight);
}

[data-theme="light"] .form-label {
  color: var(--color-stellar-blue);
  opacity: 0.8;
}

[data-theme="light"] .form-input,
[data-theme="light"] .form-textarea {
  background: rgba(255, 255, 255, 0.7);
  border-color: rgba(59, 130, 246, 0.2);
  color: var(--color-starlight);
}

[data-theme="light"] .form-input:focus,
[data-theme="light"] .form-textarea:focus {
  border-color: var(--color-stellar-blue);
  background: rgba(255, 255, 255, 0.9);
  box-shadow: 0 0 20px rgba(59, 130, 246, 0.15);
}

[data-theme="light"] .form-input::placeholder,
[data-theme="light"] .form-textarea::placeholder {
  color: var(--color-starlight);
  opacity: 0.4;
}

[data-theme="light"] .upload-button {
  border-color: rgba(59, 130, 246, 0.3);
  background: rgba(255, 255, 255, 0.5);
}

[data-theme="light"] .upload-button:hover {
  border-color: var(--color-stellar-blue);
  background: rgba(59, 130, 246, 0.08);
}

[data-theme="light"] .generate-button {
  box-shadow: 0 10px 30px rgba(59, 130, 246, 0.25);
}

[data-theme="light"] .generate-button:hover:not(:disabled) {
  box-shadow: 0 15px 40px rgba(59, 130, 246, 0.35);
}

[data-theme="light"] .waiting-text {
  opacity: 0.4;
  color: var(--color-starlight);
}

[data-theme="light"] .generating-text {
  color: var(--color-nebula-purple);
  opacity: 0.8;
}

[data-theme="light"] .resume-card {
  background: rgba(255, 255, 255, 0.8);
  border-color: rgba(59, 130, 246, 0.15);
}

[data-theme="light"] .resume-name {
  color: var(--color-starlight);
}

[data-theme="light"] .resume-position {
  color: var(--color-nebula-purple);
  opacity: 0.7;
}

[data-theme="light"] .section-label {
  color: var(--color-stellar-blue);
  opacity: 0.8;
}

[data-theme="light"] .skill-tag {
  background: rgba(59, 130, 246, 0.08);
  border-color: rgba(59, 130, 246, 0.2);
  color: var(--color-stellar-blue);
}

[data-theme="light"] .exp-bullet {
  background: var(--color-nebula-purple);
  opacity: 0.7;
}

[data-theme="light"] .exp-text {
  color: var(--color-starlight);
  opacity: 0.7;
}

/* Responsive Design */
@media (max-width: 968px) {
  .resume-layout {
    grid-template-columns: 1fr;
    gap: var(--space-xl);
  }

  .page-title {
    font-size: 2rem;
  }

  .back-button {
    margin-bottom: var(--space-lg);
  }

  .page-header .container {
    flex-direction: column;
    gap: var(--space-md);
  }

  .header-center {
    order: -1;
  }
}

/* Animation Classes */
.animate-fade-in-up {
  animation: fadeInUp 0.6s ease-out forwards;
  opacity: 0;
}
</style>