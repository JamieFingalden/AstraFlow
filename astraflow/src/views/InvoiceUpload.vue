<template>
  <div class="app-container" :data-theme="isDark ? 'dark' : 'light'">
    <!-- Particle background effect -->
    <div class="particle-background">
      <div class="particle particle-cyan"></div>
      <div class="particle particle-blue"></div>
      <div class="particle particle-purple"></div>
    </div>

    <!-- Page Header -->
    <PageHeader
      title="发票上传 Upload Invoice"
      :show-back-button="true"
      back-to="/"
      :show-theme-toggle="true"
    />

    <!-- Main Content -->
    <main class="main-content">
      <div class="content-wrapper">
        <!-- Upload Area -->
        <transition name="upload-fade" mode="out-in">
          <div
            v-if="!isUploading && !uploadResult"
            @click="triggerFileUpload"
            @dragover.prevent="handleDragOver"
            @dragleave.prevent="handleDragLeave"
            @drop.prevent="handleDrop"
            class="upload-area"
            :class="{ 'dragging': isDragging }"
          >
            <input
              ref="fileInput"
              type="file"
              accept="image/*,.pdf"
              @change="handleFileSelect"
              class="file-input"
            />

            <!-- Upload Icon -->
            <div class="upload-icon-container">
              <div class="upload-icon" :class="{ 'dragging': isDragging }">
                <UploadIcon :size="48" class="upload-icon-svg" />
              </div>
            </div>

            <!-- Upload Text -->
            <div class="upload-text">
              <h2 class="upload-title">
                上传你的发票或账单图片
              </h2>
              <p class="upload-subtitle">
                AstraFlow 将自动识别并分类
              </p>

              <!-- Upload Button -->
              <button class="upload-button">
                选择文件上传
              </button>

              <!-- File Info -->
              <p class="upload-info">
                支持 JPG、PNG、PDF 格式，包括发票、微信支付、支付宝账单等，最大 10MB
              </p>
            </div>
          </div>

          <!-- Upload Progress -->
          <div v-else-if="isUploading" class="upload-progress">
            <div class="progress-content">
              <!-- Loading Animation -->
              <div class="loading-container">
                <div class="loading-spinner">
                  <div class="spinner-bg"></div>
                  <div class="spinner-active"></div>
                </div>
              </div>

              <!-- Progress Text -->
              <h3 class="progress-title">
                AI 正在识别发票信息，请稍候...
              </h3>

              <p class="progress-description">
                {{ progressText }}
              </p>

              <!-- Progress Bar -->
              <div class="progress-bar-container">
                <div class="progress-bar-bg">
                  <div
                    class="progress-bar-fill"
                    :style="{ width: `${uploadProgress}%` }"
                  ></div>
                </div>
                <p class="progress-percentage">
                  {{ uploadProgress }}%
                </p>
              </div>
            </div>
          </div>

          <!-- Upload Result -->
          <div v-else-if="uploadResult" class="upload-result">
            <div class="result-content">
              <!-- Success Icon -->
              <div class="success-icon-container">
                <div class="success-icon">
                  <CheckCircleIcon :size="48" class="success-icon-svg" />
                </div>
              </div>

              <!-- Success Message -->
              <h3 class="result-title">
                识别完成！
              </h3>

              <!-- Recognition Results -->
              <div class="result-grid">
                <div class="result-item">
                  <p class="result-label">发票金额</p>
                  <p class="result-value">{{ uploadResult.amount }}</p>
                </div>
                <div class="result-item">
                  <p class="result-label">发票日期</p>
                  <p class="result-value">{{ uploadResult.date }}</p>
                </div>
                <div class="result-item">
                  <p class="result-label">发票类别</p>
                  <p class="result-value">{{ uploadResult.category }}</p>
                </div>
                <div class="result-item">
                  <p class="result-label">支付来源</p>
                  <p class="result-value">{{ uploadResult.source }}</p>
                </div>
              </div>

              <!-- Confidence Score -->
              <div class="confidence-section">
                <p class="confidence-label">AI 识别置信度</p>
                <div class="confidence-bar-container">
                  <div class="confidence-bar-bg">
                    <div
                      class="confidence-bar-fill"
                      :class="getConfidenceClass(uploadResult.confidence)"
                      :style="{ width: `${uploadResult.confidence}%` }"
                    ></div>
                  </div>
                  <span class="confidence-text" :class="getConfidenceClass(uploadResult.confidence)">
                    {{ uploadResult.confidence }}%
                  </span>
                </div>
              </div>

              <!-- Action Buttons -->
              <div class="action-buttons">
                <button @click="viewAnalysis" class="btn btn-primary">
                  查看分析结果
                </button>
                <button @click="resetUpload" class="btn btn-secondary">
                  继续上传
                </button>
                <button @click="viewDashboard" class="btn btn-outline">
                  查看账单
                </button>
              </div>
            </div>
          </div>
        </transition>

        <!-- Upload History -->
        <div v-if="uploadHistory.length > 0" class="upload-history">
          <h3 class="history-title">
            最近上传记录
          </h3>
          <div class="history-grid">
            <div
              v-for="record in uploadHistory"
              :key="record.id"
              class="history-item"
            >
              <div class="history-item-header">
                <div class="status-icon" :class="getStatusClass(record.status)">
                  <component
                    :is="getStatusIcon(record.status)"
                    :size="20"
                  />
                </div>
                <span class="history-time">
                  {{ record.time }}
                </span>
              </div>
              <h4 class="history-item-title">
                {{ record.name }}
              </h4>
              <p class="history-item-description">
                {{ record.category }} · {{ record.amount }} · {{ record.source }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Page Footer -->
    <PageFooter />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import {
  UploadIcon,
  CheckCircleIcon,
  ClockIcon,
  XCircleIcon,
  SunIcon,
  MoonIcon
} from 'lucide-vue-next'
import { useTheme } from '../composables/useTheme'
import PageHeader from '../components/ui/PageHeader.vue'
import PageFooter from '../components/ui/PageFooter.vue'

const router = useRouter()
const { theme, toggleTheme, isDark } = useTheme()


// Upload state
const isUploading = ref(false)
const isDragging = ref(false)
const uploadProgress = ref(0)
const progressText = ref('')
const uploadResult = ref(null)
const fileInput = ref(null)

// API Configuration
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:5000'
const OCR_ENDPOINT = '/api/ocr/upload'

// Supported file types
const SUPPORTED_TYPES = {
  images: ['image/jpeg', 'image/jpg', 'image/png', 'image/webp'],
  documents: ['application/pdf'],
  maxSize: 10 * 1024 * 1024 // 10MB
}

// File type detection
const getFileType = (file) => {
  if (SUPPORTED_TYPES.images.includes(file.type)) return 'image'
  if (SUPPORTED_TYPES.documents.includes(file.type)) return 'document'
  return 'unknown'
}

// Validate file
const validateFile = (file) => {
  const allSupportedTypes = [...SUPPORTED_TYPES.images, ...SUPPORTED_TYPES.documents]

  if (!allSupportedTypes.includes(file.type)) {
    alert('不支持的文件格式。请上传 JPG、PNG、PDF 格式的文件。')
    return false
  }

  if (file.size > SUPPORTED_TYPES.maxSize) {
    alert('文件大小超过限制。请上传小于 10MB 的文件。')
    return false
  }

  return true
}

// Upload history
const uploadHistory = ref([
  {
    id: 1,
    name: '星巴克咖啡发票',
    category: '餐饮',
    amount: '¥128.50',
    source: '微信支付',
    status: 'success',
    time: '2小时前'
  },
  {
    id: 2,
    name: '滴滴出行账单',
    category: '交通',
    amount: '¥45.00',
    source: '支付宝',
    status: 'success',
    time: '1天前'
  },
  {
    id: 3,
    name: '京东办公用品',
    category: '办公',
    amount: '¥299.90',
    source: '银联支付',
    status: 'processing',
    time: '2天前'
  }
])


// File upload functions
const triggerFileUpload = () => {
  fileInput.value?.click()
}

const handleFileSelect = (event) => {
  const file = event.target.files[0]
  if (file) {
    processFile(file)
  }
}

const handleDragOver = (event) => {
  isDragging.value = true
}

const handleDragLeave = (event) => {
  isDragging.value = false
}

const handleDrop = (event) => {
  isDragging.value = false
  const file = event.dataTransfer.files[0]
  if (file) {
    processFile(file)
  }
}

// API Upload Function (for future Flask OCR integration)
const uploadToOCR = async (file) => {
  try {
    const formData = new FormData()
    formData.append('file', file)
    formData.append('fileType', getFileType(file))

    // Future: Replace with actual API call
    // const response = await axios.post(`${API_BASE_URL}${OCR_ENDPOINT}`, formData, {
    //   headers: {
    //     'Content-Type': 'multipart/form-data',
    //   },
    //   onUploadProgress: (progressEvent) => {
    //     const progress = Math.round((progressEvent.loaded * 100) / progressEvent.total)
    //     uploadProgress.value = Math.min(progress, 90) // Keep some room for processing
    //   }
    // })

    // Mock API response for development
    await mockAPICall()

    // Mock response structure that Flask OCR should return
    const mockApiResponse = {
      amount: (Math.random() * 500 + 20).toFixed(2), // Random amount between 20-520
      date: new Date().toISOString().split('T')[0], // Current date in YYYY-MM-DD format
      category: ['餐饮', '交通', '办公', '住宿', '其他'][Math.floor(Math.random() * 5)],
      source: detectPaymentSource(file.name), // Detect payment source
      confidence: Math.floor(Math.random() * 15) + 85 // 85-99%
    }

    return { success: true, data: mockApiResponse }

  } catch (error) {
    console.error('OCR Upload Error:', error)
    return {
      success: false,
      error: error.response?.data?.message || '上传失败，请重试'
    }
  }
}

// Mock API call function (simulates network delay)
const mockAPICall = () => {
  return new Promise((resolve) => {
    setTimeout(resolve, 2000 + Math.random() * 2000) // 2-4 seconds delay
  })
}

// Detect payment source from filename (mock implementation)
const detectPaymentSource = (filename) => {
  const name = filename.toLowerCase()

  if (name.includes('wechat') || name.includes('微信')) return '微信支付'
  if (name.includes('alipay') || name.includes('支付宝')) return '支付宝'
  if (name.includes('didi') || name.includes('滴滴')) return '滴滴出行'
  if (name.includes('meituan') || name.includes('美团')) return '美团'
  if (name.includes('starbucks') || name.includes('星巴克')) return '星巴克'

  const sources = ['微信支付', '支付宝', '银联支付', '现金', '其他']
  return sources[Math.floor(Math.random() * sources.length)]
}

const processFile = async (file) => {
  // Validate file first
  if (!validateFile(file)) {
    return
  }

  // Start upload
  isUploading.value = true
  uploadProgress.value = 0
  progressText.value = '正在上传文件...'

  // Simulate upload progress
  const progressSteps = [
    { progress: 20, text: '正在上传文件...' },
    { progress: 40, text: '文件上传完成，开始AI识别...' },
    { progress: 60, text: '正在分析发票内容...' },
    { progress: 80, text: '正在提取关键信息...' },
    { progress: 95, text: '正在验证识别结果...' }
  ]

  let currentStep = 0
  const progressInterval = setInterval(() => {
    if (currentStep < progressSteps.length) {
      uploadProgress.value = progressSteps[currentStep].progress
      progressText.value = progressSteps[currentStep].text
      currentStep++
    } else {
      clearInterval(progressInterval)
    }
  }, 600)

  try {
    // Call OCR API (currently mocked)
    const result = await uploadToOCR(file)

    // Clear any remaining progress interval
    clearInterval(progressInterval)

    if (result.success) {
      uploadProgress.value = 100
      progressText.value = '识别完成！'
      completeUpload(file, result.data)
    } else {
      throw new Error(result.error)
    }
  } catch (error) {
    clearInterval(progressInterval)
    alert(error.message || '上传失败，请重试')
    resetUpload()
  }
}

const completeUpload = (file, apiResult) => {
  // Use API result or fallback to mock data
  uploadResult.value = {
    amount: `¥${apiResult.amount}`,
    date: apiResult.date,
    category: apiResult.category,
    source: apiResult.source,
    confidence: Math.round(apiResult.confidence * 100) // Convert decimal to percentage
  }

  // Add to history
  uploadHistory.value.unshift({
    id: Date.now(),
    name: file.name.replace(/\.[^/.]+$/, ""),
    category: uploadResult.value.category,
    amount: uploadResult.value.amount,
    source: uploadResult.value.source,
    status: 'success',
    time: '刚刚'
  })

  // Keep only last 3 records
  if (uploadHistory.value.length > 3) {
    uploadHistory.value = uploadHistory.value.slice(0, 3)
  }

  isUploading.value = false
}

const resetUpload = () => {
  uploadResult.value = null
  uploadProgress.value = 0
  progressText.value = ''
}

const viewDashboard = () => {
  router.push('/dashboard')
}

const viewAnalysis = () => {
  router.push('/analysis')
}

// Helper methods for dynamic classes
const getConfidenceClass = (confidence) => {
  if (confidence >= 90) return 'confidence-high'
  if (confidence >= 70) return 'confidence-medium'
  return 'confidence-low'
}

const getStatusClass = (status) => {
  if (status === 'success') return 'status-success'
  if (status === 'processing') return 'status-processing'
  return 'status-error'
}

const getStatusIcon = (status) => {
  if (status === 'success') return CheckCircleIcon
  if (status === 'processing') return ClockIcon
  return XCircleIcon
}
</script>

<style scoped>
/* Main Container */
.app-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  transition: color 0.3s ease, background-color 0.3s ease;
}

.app-container[data-theme="dark"] {
  background-color: #111827;
}

.app-container[data-theme="light"] {
  background: linear-gradient(135deg, #eff6ff, #ecfeff, #e0f2fe);
}

/* Particle Background */
.particle-background {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  overflow: hidden;
  pointer-events: none;
}

.particle {
  position: absolute;
  width: 20rem;
  height: 20rem;
  border-radius: 50%;
  mix-blend-mode: multiply;
  filter: blur(5rem);
  animation: pulse 4s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

.particle-cyan {
  top: -10rem;
  right: -10rem;
}

.app-container[data-theme="dark"] .particle-cyan {
  background-color: rgba(34, 211, 238, 0.3);
}

.app-container[data-theme="light"] .particle-cyan {
  background-color: #06b6d4;
}

.particle-blue {
  bottom: -10rem;
  left: -10rem;
}

.app-container[data-theme="dark"] .particle-blue {
  background-color: rgba(59, 130, 246, 0.3);
}

.app-container[data-theme="light"] .particle-blue {
  background-color: #3b82f6;
}

.particle-purple {
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.app-container[data-theme="dark"] .particle-purple {
  background-color: rgba(147, 51, 234, 0.3);
}

.app-container[data-theme="light"] .particle-purple {
  background-color: #a855f7;
}

/* Header */
.app-header {
  position: relative;
  z-index: 10;
  backdrop-filter: blur(12px);
  border-bottom: 1px solid;
  transition: all 0.3s ease;
}

.app-container[data-theme="dark"] .app-header {
  background-color: rgba(31, 41, 55, 0.7);
  border-color: rgba(55, 65, 81, 0.5);
}

.app-container[data-theme="light"] .app-header {
  background-color: rgba(255, 255, 255, 0.7);
  border-color: rgba(229, 231, 235, 0.5);
}

.header-content {
  max-width: 80rem;
  margin: 0 auto;
  padding: 0 1rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 4rem;
  flex-wrap: nowrap;
}

@media (min-width: 640px) {
  .header-content {
    padding: 0 1.5rem;
  }
}

@media (min-width: 1024px) {
  .header-content {
    padding: 0 2rem;
  }
}

.header-left {
  display: flex;
  align-items: center;
  gap: 1rem;
  flex: 1;
  min-width: 0; /* Allow flex item to shrink */
  overflow: hidden;
}

.back-button {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 0.75rem;
  border-radius: 0.5rem;
  text-decoration: none;
  transition: all 0.2s ease;
}

.app-container[data-theme="dark"] .back-button {
  color: #d1d5db;
}

.app-container[data-theme="dark"] .back-button:hover {
  background-color: rgba(55, 65, 81, 0.5);
  color: white;
}

.app-container[data-theme="light"] .back-button {
  color: #4b5563;
}

.app-container[data-theme="light"] .back-button:hover {
  background-color: #f3f4f6;
  color: #111827;
}

.back-text {
  display: none;
}

@media (min-width: 640px) {
  .back-text {
    display: inline;
  }
}

.brand-name {
  font-size: 1.5rem;
  font-weight: 700;
  background: linear-gradient(135deg, #22d3ee, #3b82f6);
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
}

.page-title {
  font-size: 1.25rem;
  font-weight: 600;
  transition: color 0.3s ease;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  flex-shrink: 0;
}

.app-container[data-theme="dark"] .page-title {
  color: white;
}

.app-container[data-theme="light"] .page-title {
  color: #1f2937;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 1rem;
  flex: 0; /* Don't grow */
}

.theme-toggle {
  padding: 0.5rem;
  border-radius: 0.5rem;
  border: none;
  transition: all 0.2s ease;
  cursor: pointer;
}

.app-container[data-theme="dark"] .theme-toggle {
  background-color: rgba(55, 65, 81, 0.5);
  color: #d1d5db;
}

.app-container[data-theme="dark"] .theme-toggle:hover {
  background-color: rgba(75, 85, 99, 0.5);
}

.app-container[data-theme="light"] .theme-toggle {
  background-color: #f3f4f6;
  color: #374151;
}

.app-container[data-theme="light"] .theme-toggle:hover {
  background-color: #e5e7eb;
}

/* Main Content */
.main-content {
  position: relative;
  z-index: 10;
  flex-grow: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 3rem 1rem;
}

.content-wrapper {
  max-width: 56rem;
  width: 100%;
}

/* Upload Area */
.upload-area {
  position: relative;
  border-radius: 1rem;
  padding: 3rem;
  border: 2px dashed;
  backdrop-filter: blur(12px);
  transition: all 0.3s ease;
  cursor: pointer;
}

.app-container[data-theme="dark"] .upload-area {
  border-color: #4b5563;
  background-color: rgba(31, 41, 55, 0.5);
}

.app-container[data-theme="dark"] .upload-area:hover {
  border-color: #06b6d4;
  box-shadow: 0 20px 25px -5px rgba(6, 189, 212, 0.1), 0 10px 10px -5px rgba(6, 189, 212, 0.04);
}

.app-container[data-theme="light"] .upload-area {
  border-color: #d1d5db;
  background-color: rgba(255, 255, 255, 0.5);
}

.app-container[data-theme="light"] .upload-area:hover {
  border-color: #22d3ee;
  box-shadow: 0 20px 25px -5px rgba(34, 211, 238, 0.15), 0 10px 10px -5px rgba(34, 211, 238, 0.08);
}

.upload-area.dragging {
  border-color: #22d3ee;
  background-color: rgba(34, 211, 238, 0.1);
  box-shadow: 0 25px 50px -12px rgba(34, 211, 238, 0.25);
}

.file-input {
  display: none;
}

/* Upload Icon */
.upload-icon-container {
  display: flex;
  justify-content: center;
  margin-bottom: 1.5rem;
}

.upload-icon {
  padding: 1.5rem;
  border-radius: 50%;
  transition: all 0.3s ease;
}

.app-container[data-theme="dark"] .upload-icon {
  background-color: #2563eb;
}

.app-container[data-theme="dark"] .upload-icon:hover {
  background-color: #3b82f6;
}

.app-container[data-theme="light"] .upload-icon {
  background-color: #3b82f6;
}

.app-container[data-theme="light"] .upload-icon:hover {
  background-color: #2563eb;
}

.upload-icon.dragging {
  background-color: #06b6d4;
  transform: scale(1.1);
}

.upload-icon-svg {
  color: white;
}

/* Upload Text */
.upload-text {
  text-align: center;
}

.upload-title {
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 0.75rem;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .upload-title {
  color: white;
}

.app-container[data-theme="light"] .upload-title {
  color: #1f2937;
}

.upload-subtitle {
  font-size: 1.125rem;
  margin-bottom: 1.5rem;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .upload-subtitle {
  color: #d1d5db;
}

.app-container[data-theme="light"] .upload-subtitle {
  color: #4b5563;
}

.upload-button {
  padding: 0.75rem 2rem;
  border-radius: 0.5rem;
  font-weight: 600;
  border: none;
  cursor: pointer;
  transition: all 0.3s ease;
  transform: scale(1);
}

.upload-button:hover {
  transform: scale(1.05);
}

.app-container[data-theme="dark"] .upload-button {
  background: linear-gradient(135deg, #2563eb, #06b6d4);
  color: white;
  box-shadow: 0 10px 15px -3px rgba(37, 99, 235, 0.25);
}

.app-container[data-theme="dark"] .upload-button:hover {
  background: linear-gradient(135deg, #3b82f6, #22d3ee);
  box-shadow: 0 20px 25px -5px rgba(37, 99, 235, 0.3);
}

.app-container[data-theme="light"] .upload-button {
  background: linear-gradient(135deg, #3b82f6, #06b6d4);
  color: white;
  box-shadow: 0 10px 15px -3px rgba(37, 99, 235, 0.3);
}

.app-container[data-theme="light"] .upload-button:hover {
  background: linear-gradient(135deg, #2563eb, #0891b2);
  box-shadow: 0 20px 25px -5px rgba(37, 99, 235, 0.4);
}

.upload-info {
  font-size: 0.875rem;
  margin-top: 1rem;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .upload-info {
  color: #9ca3af;
}

.app-container[data-theme="light"] .upload-info {
  color: #6b7280;
}

/* Upload Progress */
.upload-progress {
  border-radius: 1rem;
  padding: 3rem;
  backdrop-filter: blur(12px);
  transition: all 0.3s ease;
}

.app-container[data-theme="dark"] .upload-progress {
  background-color: rgba(31, 41, 55, 0.7);
  border: 1px solid rgba(55, 65, 81, 0.5);
}

.app-container[data-theme="light"] .upload-progress {
  background-color: rgba(255, 255, 255, 0.7);
  border: 1px solid rgba(229, 231, 235, 0.5);
}

.progress-content {
  text-align: center;
}

.loading-container {
  display: flex;
  justify-content: center;
  margin-bottom: 1.5rem;
}

.loading-spinner {
  position: relative;
}

.spinner-bg {
  width: 4rem;
  height: 4rem;
  border-radius: 50%;
  border: 4px solid #dbeafe;
  animation: spin 1s linear infinite;
}

.spinner-active {
  position: absolute;
  top: 0;
  left: 0;
  width: 4rem;
  height: 4rem;
  border-radius: 50%;
  border: 4px solid transparent;
  border-top-color: #3b82f6;
  animation: spin 1s linear infinite;
}

.progress-title {
  font-size: 1.25rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .progress-title {
  color: white;
}

.app-container[data-theme="light"] .progress-title {
  color: #1f2937;
}

.progress-description {
  font-size: 0.875rem;
  margin-bottom: 1.5rem;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .progress-description {
  color: #9ca3af;
}

.app-container[data-theme="light"] .progress-description {
  color: #4b5563;
}

.progress-bar-container {
  max-width: 16rem;
  margin: 0 auto;
}

.progress-bar-bg {
  height: 0.5rem;
  border-radius: 9999px;
  overflow: hidden;
}

.app-container[data-theme="dark"] .progress-bar-bg {
  background-color: #374151;
}

.app-container[data-theme="light"] .progress-bar-bg {
  background-color: #e5e7eb;
}

.progress-bar-fill {
  height: 100%;
  background: linear-gradient(135deg, #3b82f6, #06b6d4);
  transition: width 0.5s ease;
}

.progress-percentage {
  font-size: 0.875rem;
  margin-top: 0.5rem;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .progress-percentage {
  color: #9ca3af;
}

.app-container[data-theme="light"] .progress-percentage {
  color: #4b5563;
}

/* Upload Result */
.upload-result {
  border-radius: 1rem;
  padding: 3rem;
  backdrop-filter: blur(12px);
  transition: all 0.3s ease;
}

.app-container[data-theme="dark"] .upload-result {
  background-color: rgba(31, 41, 55, 0.7);
  border: 1px solid rgba(55, 65, 81, 0.5);
}

.app-container[data-theme="light"] .upload-result {
  background-color: rgba(255, 255, 255, 0.7);
  border: 1px solid rgba(229, 231, 235, 0.5);
}

.result-content {
  text-align: center;
}

.success-icon-container {
  display: flex;
  justify-content: center;
  margin-bottom: 1.5rem;
}

.success-icon {
  padding: 1.5rem;
  border-radius: 50%;
  background-color: #10b981;
}

.success-icon-svg {
  color: white;
}

.result-title {
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 1.5rem;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .result-title {
  color: white;
}

.app-container[data-theme="light"] .result-title {
  color: #1f2937;
}

.result-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1rem;
  margin-bottom: 2rem;
}

@media (min-width: 768px) {
  .result-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (min-width: 1024px) {
  .result-grid {
    grid-template-columns: repeat(4, 1fr);
  }
}

.result-item {
  padding: 1rem;
  border-radius: 0.5rem;
}

.app-container[data-theme="dark"] .result-item {
  background-color: rgba(55, 65, 81, 0.5);
}

.app-container[data-theme="light"] .result-item {
  background-color: #f3f4f6;
}

.result-label {
  font-size: 0.875rem;
  margin-bottom: 0.25rem;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .result-label {
  color: #9ca3af;
}

.app-container[data-theme="light"] .result-label {
  color: #4b5563;
}

.result-value {
  font-size: 1.25rem;
  font-weight: 700;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .result-value {
  color: white;
}

.app-container[data-theme="light"] .result-value {
  color: #1f2937;
}

/* Confidence Section */
.confidence-section {
  margin-bottom: 2rem;
}

.confidence-label {
  font-size: 0.875rem;
  margin-bottom: 0.5rem;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .confidence-label {
  color: #9ca3af;
}

.app-container[data-theme="light"] .confidence-label {
  color: #4b5563;
}

.confidence-bar-container {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
}

.confidence-bar-bg {
  height: 0.75rem;
  border-radius: 9999px;
  overflow: hidden;
  max-width: 12rem;
}

.app-container[data-theme="dark"] .confidence-bar-bg {
  background-color: #374151;
}

.app-container[data-theme="light"] .confidence-bar-bg {
  background-color: #e5e7eb;
}

.confidence-bar-fill {
  height: 100%;
  transition: width 1s ease;
}

.confidence-high {
  background-color: #10b981;
}

.confidence-medium {
  background-color: #eab308;
}

.confidence-low {
  background-color: #ef4444;
}

.confidence-text {
  font-weight: 700;
}

.confidence-high {
  color: #10b981;
}

.confidence-medium {
  color: #eab308;
}

.confidence-low {
  color: #ef4444;
}

/* Action Buttons */
.action-buttons {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  justify-content: center;
}

@media (min-width: 640px) {
  .action-buttons {
    flex-direction: row;
  }
}

.btn {
  padding: 0.75rem 1.5rem;
  border-radius: 0.5rem;
  font-weight: 600;
  border: none;
  cursor: pointer;
  transition: all 0.3s ease;
  transform: scale(1);
}

.btn:hover {
  transform: scale(1.05);
}

.btn-primary {
  background: linear-gradient(135deg, #9333ea, #ec4899);
  color: white;
}

.app-container[data-theme="dark"] .btn-primary {
  box-shadow: 0 10px 15px -3px rgba(147, 51, 234, 0.25);
}

.app-container[data-theme="light"] .btn-primary {
  box-shadow: 0 10px 15px -3px rgba(147, 51, 234, 0.3);
}

.btn-secondary {
  background: linear-gradient(135deg, #2563eb, #06b6d4);
  color: white;
}

.app-container[data-theme="dark"] .btn-secondary {
  box-shadow: 0 10px 15px -3px rgba(37, 99, 235, 0.25);
}

.app-container[data-theme="light"] .btn-secondary {
  box-shadow: 0 10px 15px -3px rgba(37, 99, 235, 0.3);
}

.btn-outline {
  border: 1px solid;
  background: transparent;
}

.app-container[data-theme="dark"] .btn-outline {
  border-color: #4b5563;
  color: #d1d5db;
}

.app-container[data-theme="dark"] .btn-outline:hover {
  background-color: #374151;
}

.app-container[data-theme="light"] .btn-outline {
  border-color: #d1d5db;
  color: #374151;
}

.app-container[data-theme="light"] .btn-outline:hover {
  background-color: #f3f4f6;
}

/* Upload History */
.upload-history {
  margin-top: 3rem;
}

.history-title {
  font-size: 1.125rem;
  font-weight: 600;
  margin-bottom: 1.5rem;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .history-title {
  color: white;
}

.app-container[data-theme="light"] .history-title {
  color: #1f2937;
}

.history-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1.5rem;
}

@media (min-width: 768px) {
  .history-grid {
    grid-template-columns: repeat(3, 1fr);
  }
}

.history-item {
  padding: 1.5rem;
  border-radius: 0.75rem;
  border: 1px solid;
  backdrop-filter: blur(12px);
  transition: all 0.3s ease;
}

.app-container[data-theme="dark"] .history-item {
  background-color: rgba(31, 41, 55, 0.5);
  border-color: rgba(55, 65, 81, 0.5);
}

.app-container[data-theme="dark"] .history-item:hover {
  border-color: rgba(6, 189, 212, 0.3);
  box-shadow: 0 10px 15px -3px rgba(6, 189, 212, 0.1);
}

.app-container[data-theme="light"] .history-item {
  background-color: rgba(255, 255, 255, 0.5);
  border-color: rgba(229, 231, 235, 0.5);
}

.app-container[data-theme="light"] .history-item:hover {
  border-color: rgba(34, 211, 238, 0.3);
  box-shadow: 0 10px 15px -3px rgba(34, 211, 238, 0.1);
}

.history-item-header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  margin-bottom: 1rem;
}

.status-icon {
  padding: 0.5rem;
  border-radius: 0.5rem;
}

.status-success {
  background-color: #dcfce7;
  color: #16a34a;
}

.status-processing {
  background-color: #fef3c7;
  color: #d97706;
}

.status-error {
  background-color: #fee2e2;
  color: #dc2626;
}

.history-time {
  font-size: 0.75rem;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .history-time {
  color: #9ca3af;
}

.app-container[data-theme="light"] .history-time {
  color: #6b7280;
}

.history-item-title {
  font-weight: 500;
  margin-bottom: 0.5rem;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .history-item-title {
  color: white;
}

.app-container[data-theme="light"] .history-item-title {
  color: #1f2937;
}

.history-item-description {
  font-size: 0.875rem;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .history-item-description {
  color: #9ca3af;
}

.app-container[data-theme="light"] .history-item-description {
  color: #4b5563;
}


/* Upload area transitions */
.upload-fade-enter-active,
.upload-fade-leave-active {
  transition: opacity 0.5s ease, transform 0.5s ease;
}

.upload-fade-enter-from {
  opacity: 0;
  transform: translateY(20px);
}

.upload-fade-leave-to {
  opacity: 0;
  transform: translateY(-20px);
}

/* Animations */
@keyframes pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.5;
  }
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
</style>