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
              <h3 class="result-title" :class="{ 'processing-title': uploadResult.amount === '处理中...' }">
                {{ uploadResult.amount === '处理中...' ? '文件已上传' : '识别完成！' }}
              </h3>

              <!-- Recognition Results -->
              <div class="result-notification">
                <p>发票识别后将会在下方列表展示</p>
              </div>

              <!-- Processing notification -->
              <div v-if="uploadResult.amount === '处理中...'" class="processing-notification">
                <p>发票识别正在进行中，结果将在处理完成后自动更新。</p>
                <p>您也可以前往账单管理页面查看最新识别结果。</p>
              </div>

              <!-- Action Buttons -->
              <div class="action-buttons">

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
        <div class="upload-history">
          <h3 class="history-title">
            最近上传记录
          </h3>
          
          <div v-if="historyLoading" class="loading-history">
             <div class="spinner-bg" style="width: 2rem; height: 2rem; margin: 2rem auto;"></div>
          </div>
          
          <div v-else-if="uploadHistory.length === 0" class="no-history">
            暂无上传记录
          </div>

          <div v-else class="history-list">
            <div
              v-for="record in uploadHistory"
              :key="record.id"
              class="history-list-item"
            >
              <div class="item-left">
                <div class="status-icon" :class="getStatusClass(record.status)">
                  <component
                    :is="getStatusIcon(record.status)"
                    :size="24"
                  />
                </div>
                <div class="item-info">
                  <h4 class="history-item-title">
                    {{ record.name }}
                  </h4>
                  <div class="history-item-meta">
                    <span class="meta-type">{{ record.type }}</span>
                    <span class="meta-dot">·</span>
                    <span class="meta-size">{{ record.size }}</span>
                    <span class="meta-dot">·</span>
                    <span class="meta-time">{{ record.time }}</span>
                  </div>
                </div>
              </div>
              
              <div class="item-right">
                <div class="history-status-badge" :class="getStatusClass(record.status, true)">
                    {{ record.statusText }}
                </div>
              </div>
            </div>
          </div>

          <!-- Pagination -->
          <div v-if="pagination.itemCount > 0" class="pagination-container">
            <n-pagination
              v-model:page="pagination.page"
              v-model:page-size="pagination.pageSize"
              :item-count="pagination.itemCount"
              @update:page="handlePageChange"
            />
          </div>
        </div>
      </div>
    </main>

    <!-- Page Footer -->
    <PageFooter />
  </div>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue'
import { useRouter } from 'vue-router'
import {
  UploadIcon,
  CheckCircleIcon,
  ClockIcon,
  XCircleIcon,
  AlertCircleIcon,
  SunIcon,
  MoonIcon
} from 'lucide-vue-next'
import { useTheme } from '../composables/useTheme'
import PageHeader from '../components/ui/PageHeader.vue'
import PageFooter from '../components/ui/PageFooter.vue'
import { attachmentApi } from '../services/api/attachmentApi'
import { NPagination } from 'naive-ui'

const router = useRouter()
const { theme, toggleTheme, isDark } = useTheme()


// Upload state
const isUploading = ref(false)
const isDragging = ref(false)
const uploadProgress = ref(0)
const progressText = ref('')
const uploadResult = ref(null)
const fileInput = ref(null)

// History state
const uploadHistory = ref([])
const historyLoading = ref(false)
const pagination = reactive({
  page: 1,
  pageSize: 10, // List layout
  itemCount: 0
})

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

// Format helpers
const formatDate = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  const now = new Date()
  const diff = now - date
  
  // If less than 24 hours, show relative time
  if (diff < 24 * 60 * 60 * 1000) {
    if (diff < 60 * 60 * 1000) {
      const minutes = Math.max(1, Math.floor(diff / (60 * 1000)))
      return `${minutes}分钟前`
    }
    const hours = Math.floor(diff / (60 * 60 * 1000))
    return `${hours}小时前`
  }
  
  return date.toLocaleDateString('zh-CN', {
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const formatFileSize = (bytes) => {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

// Fetch history
const fetchHistory = async () => {
  historyLoading.value = true
  try {
    const response = await attachmentApi.getAttachmentsByTenant({
      page: pagination.page,
      page_size: pagination.pageSize
    })
    
    if (response.data && response.code === 200) {
      const { attachments, total } = response.data
      pagination.itemCount = total
      
      uploadHistory.value = (attachments || []).map(item => {
        // Map status int to string for UI logic
        let status = 'pending'
        let statusText = '识别中'
        
        // Backend status: 0=pending, 1=success, 2=failed, 3=duplicate
        if (item.status === 1) {
            status = 'success'
            statusText = '识别成功'
        } else if (item.status === 2) {
            status = 'failed'
            statusText = '识别失败'
        } else if (item.status === 3) {
            status = 'duplicate'
            statusText = '重复上传'
        } else {
            status = 'processing' // 0 or default
            statusText = '识别中'
        }
        
        return {
          id: item.id,
          name: item.file_name || '未命名文件', // Prioritize file_name, fallback to URL parsing
          type: item.file_type || 'Unknown',
          size: formatFileSize(item.file_size || 0),
          status: status,
          statusText: statusText,
          time: formatDate(item.created_at)
        }
      })
    }
  } catch (error) {
    console.error('Failed to fetch upload history:', error)
  } finally {
    historyLoading.value = false
  }
}

const handlePageChange = (page) => {
  pagination.page = page
  fetchHistory()
}

onMounted(() => {
  fetchHistory()
})


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

// API Upload Function (actual Go backend API)
const uploadToOCR = async (file) => {
  try {
    const uploadData = {
      file: file
    }

    // Call the actual backend API
    const response = await attachmentApi.uploadFile(uploadData)

    // Check if the API response is successful
    if (response.code === 200) {
      // Return a structure that matches what the frontend expects
      // For now, we return a placeholder since the actual OCR result will come later via callback
      const result = {
        amount: '待OCR处理...', // Will be filled in later when callback updates the invoice
        date: '待OCR处理...',
        category: '待OCR处理...',
        source: '待OCR处理...'
      }

      return { success: true, data: result }
    } else {
      throw new Error(response.message || '上传失败')
    }
  } catch (error) {
    console.error('OCR Upload Error:', error)
    return {
      success: false,
      error: error.response?.data?.message || error.message || '上传失败，请重试'
    }
  }
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
    // Call actual OCR API
    const result = await uploadToOCR(file)

    // Clear any remaining progress interval
    clearInterval(progressInterval)

    if (result.success) {
      uploadProgress.value = 100
      progressText.value = '文件上传成功，OCR识别中...结果将在完成后显示'
      completeUpload(file, result.data)
      // Refresh history after upload
      setTimeout(() => fetchHistory(), 1000)
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
  // Use API result with proper handling for pending OCR results
  uploadResult.value = {
    amount: apiResult.amount !== '待OCR处理...' ? `¥${apiResult.amount}` : '处理中...',
    date: apiResult.date !== '待OCR处理...' ? apiResult.date : '处理中...',
    category: apiResult.category !== '待OCR处理...' ? apiResult.category : '处理中...',
    source: apiResult.source !== '待OCR处理...' ? apiResult.source : '处理中...'
  }
  
  // Note: Local history update removed in favor of fetchHistory() refresh

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




const getStatusClass = (status, isText = false) => {
  if (isText) {
      if (status === 'success') return 'text-green-600 dark:text-green-400'
      if (status === 'processing') return 'text-yellow-600 dark:text-yellow-400'
      if (status === 'duplicate') return 'text-purple-600 dark:text-purple-400'
      return 'text-red-600 dark:text-red-400'
  }
  if (status === 'success') return 'status-success'
  if (status === 'processing') return 'status-processing'
  if (status === 'duplicate') return 'status-duplicate'
  return 'status-error'
}

const getStatusIcon = (status) => {
  if (status === 'success') return CheckCircleIcon
  if (status === 'processing') return ClockIcon
  if (status === 'duplicate') return AlertCircleIcon
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

.history-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.history-list-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 1.5rem;
  border-radius: 0.75rem;
  border: 1px solid;
  backdrop-filter: blur(12px);
  transition: all 0.3s ease;
}

.app-container[data-theme="dark"] .history-list-item {
  background-color: rgba(31, 41, 55, 0.5);
  border-color: rgba(55, 65, 81, 0.5);
}

.app-container[data-theme="dark"] .history-list-item:hover {
  border-color: rgba(6, 189, 212, 0.3);
  box-shadow: 0 10px 15px -3px rgba(6, 189, 212, 0.1);
}

.app-container[data-theme="light"] .history-list-item {
  background-color: rgba(255, 255, 255, 0.5);
  border-color: rgba(229, 231, 235, 0.5);
}

.app-container[data-theme="light"] .history-list-item:hover {
  border-color: rgba(34, 211, 238, 0.3);
  box-shadow: 0 10px 15px -3px rgba(34, 211, 238, 0.1);
}

.item-left {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.item-info {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.status-icon {
  padding: 0.25rem;
  border-radius: 0.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
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

.status-duplicate {
  background-color: #f3e8ff;
  color: #9333ea;
}

.history-item-title {
  font-size: 1rem;
  font-weight: 600;
  margin: 0;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .history-item-title {
  color: white;
}

.app-container[data-theme="light"] .history-item-title {
  color: #1f2937;
}

.history-item-meta {
  display: flex;
  align-items: center;
  font-size: 0.875rem;
  transition: color 0.3s ease;
}

.meta-dot {
  margin: 0 0.5rem;
}

.app-container[data-theme="dark"] .history-item-meta {
  color: #9ca3af;
}

.app-container[data-theme="light"] .history-item-meta {
  color: #6b7280;
}

.history-status-badge {
  padding: 0.25rem 0.75rem;
  border-radius: 9999px;
  font-size: 0.75rem;
  font-weight: 500;
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

/* Processing title styles */
.result-title.processing-title {
  color: #f59e0b !important; /* 橙色表示处理中 */
}

/* Result notification styles */
.result-notification {
  margin: 1.5rem 0;
  padding: 1rem;
  border-radius: 0.5rem;
  background-color: #e0f2fe; /* 浅蓝色背景 */
  border: 1px solid #7dd3fc;
  text-align: center;
}

.result-notification p {
  margin: 0;
  font-size: 1rem;
  color: #0369a1; /* 深蓝色文字 */
  font-weight: 500;
}

.processing-notification {
  margin: 1.5rem 0;
  padding: 1rem;
  border-radius: 0.5rem;
  background-color: #fef3c7; /* 黄色背景提示 */
  border: 1px solid #fbbf24;
}

.processing-notification p {
  margin: 0.25rem 0;
  font-size: 0.875rem;
  color: #92400e; /* 深黄色文字 */
  text-align: center;
}

.pagination-container {
    display: flex;
    justify-content: center;
    margin-top: 2rem;
}

.no-history {
    text-align: center;
    padding: 3rem;
    color: #9ca3af;
}
</style>