<template>
  <div :class="`flex flex-col min-h-screen transition-colors duration-300 ${isDark ? 'bg-gray-900' : 'bg-gradient-to-br from-blue-50 via-cyan-50 to-indigo-50'}`">
    <!-- Particle background effect -->
    <div class="fixed inset-0 overflow-hidden pointer-events-none">
      <div :class="`absolute -top-40 -right-40 w-80 h-80 rounded-full mix-blend-multiply filter blur-xl animate-pulse ${isDark ? 'bg-cyan-900/30' : 'bg-cyan-500'}`"></div>
      <div :class="`absolute -bottom-40 -left-40 w-80 h-80 rounded-full mix-blend-multiply filter blur-xl animate-pulse ${isDark ? 'bg-blue-900/30' : 'bg-blue-500'}`"></div>
      <div :class="`absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-80 h-80 rounded-full mix-blend-multiply filter blur-xl animate-pulse ${isDark ? 'bg-purple-900/30' : 'bg-purple-500'}`"></div>
    </div>

    <!-- Header -->
    <header :class="`relative z-10 backdrop-blur-md border-b transition-all duration-300 ${
      isDark
        ? 'bg-gray-800/70 border-gray-700/50'
        : 'bg-white/70 border-gray-200/50'
    }`">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center h-16">
          <div class="flex items-center space-x-4">
            <!-- 返回主页按钮 -->
            <router-link
              to="/"
              :class="`flex items-center space-x-2 px-3 py-2 rounded-lg transition-all duration-200 ${
                isDark
                  ? 'hover:bg-gray-700/50 text-gray-300 hover:text-white'
                  : 'hover:bg-gray-100 text-gray-600 hover:text-gray-900'
              }`"
              title="返回主页"
            >
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="m15 18-6-6 6-6"/>
              </svg>
              <span class="hidden sm:inline">返回</span>
            </router-link>

            <div class="text-2xl font-bold bg-gradient-to-r from-cyan-400 to-blue-500 bg-clip-text text-transparent">
              AstraFlow
            </div>
            <h1 :class="`text-xl font-semibold transition-colors duration-300 ${isDark ? 'text-white' : 'text-gray-800'}`">
              发票上传 Upload Invoice
            </h1>
          </div>

          <div class="flex items-center space-x-4">
            <button
              @click="toggleTheme"
              :class="`p-2 rounded-lg transition-all duration-200 ${
                isDark
                  ? 'bg-gray-700/50 hover:bg-gray-600/50 text-gray-300'
                  : 'bg-gray-100 hover:bg-gray-200 text-gray-700'
              }`"
            >
              <SunIcon v-if="isDark" :size="20" />
              <MoonIcon v-else :size="20" />
            </button>
          </div>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="relative z-10 flex-grow flex items-center justify-center  px-4 py-12">
      <div class="max-w-4xl w-full">
        <!-- Upload Area -->
        <transition name="upload-fade" mode="out-in">
          <div
            v-if="!isUploading && !uploadResult"
            @click="triggerFileUpload"
            @dragover.prevent="handleDragOver"
            @dragleave.prevent="handleDragLeave"
            @drop.prevent="handleDrop"
            :class="`relative rounded-2xl p-12 border-2 border-dashed backdrop-blur-md transition-all duration-300 cursor-pointer ${
              isDragging
                ? 'border-cyan-400 bg-cyan-50/20 dark:bg-cyan-900/20 shadow-2xl shadow-cyan-400/20'
                : isDark
                  ? 'border-gray-600 bg-gray-800/50 hover:border-cyan-500 hover:shadow-cyan-500/10'
                  : 'border-gray-300 bg-white/50 hover:border-cyan-400 hover:shadow-cyan-400/10'
            }`"
          >
            <input
              ref="fileInput"
              type="file"
              accept="image/*,.pdf"
              @change="handleFileSelect"
              class="hidden"
            />

            <!-- Upload Icon -->
            <div class="flex justify-center mb-6">
              <div :class="`p-6 rounded-full transition-all duration-300 ${
                isDragging
                  ? 'bg-cyan-500 scale-110'
                  : isDark
                    ? 'bg-blue-600 hover:bg-blue-500'
                    : 'bg-blue-500 hover:bg-blue-600'
              }`">
                <UploadIcon :size="48" class="text-white" />
              </div>
            </div>

            <!-- Upload Text -->
            <div class="text-center">
              <h2 :class="`text-2xl font-bold mb-3 transition-colors duration-300 ${
                isDark ? 'text-white' : 'text-gray-800'
              }`">
                上传你的发票或账单图片
              </h2>
              <p :class="`text-lg mb-6 transition-colors duration-300 ${
                isDark ? 'text-gray-300' : 'text-gray-600'
              }`">
                AstraFlow 将自动识别并分类
              </p>

              <!-- Upload Button -->
              <button
                :class="`px-8 py-3 rounded-lg font-semibold transition-all duration-300 transform hover:scale-105 ${
                  isDark
                    ? 'bg-gradient-to-r from-blue-600 to-cyan-600 hover:from-blue-500 hover:to-cyan-500 text-white shadow-lg shadow-blue-500/25 hover:shadow-xl hover:shadow-blue-500/30'
                    : 'bg-gradient-to-r from-blue-500 to-cyan-500 hover:from-blue-600 hover:to-cyan-600 text-white shadow-lg shadow-blue-500/30 hover:shadow-xl hover:shadow-blue-500/40'
                }`"
              >
                选择文件上传
              </button>

              <!-- File Info -->
              <p :class="`text-sm mt-4 transition-colors duration-300 ${
                isDark ? 'text-gray-400' : 'text-gray-500'
              }`">
                支持 JPG、PNG、PDF 格式，包括发票、微信支付、支付宝账单等，最大 10MB
              </p>
            </div>
          </div>

          <!-- Upload Progress -->
          <div
            v-else-if="isUploading"
            :class="`rounded-2xl p-12 backdrop-blur-md transition-all duration-300 ${
              isDark
                ? 'bg-gray-800/70 border border-gray-700/50'
                : 'bg-white/70 border border-gray-200/50'
            }`"
          >
            <div class="text-center">
              <!-- Loading Animation -->
              <div class="flex justify-center mb-6">
                <div class="relative">
                  <div class="animate-spin rounded-full h-16 w-16 border-4 border-blue-200"></div>
                  <div :class="`absolute top-0 left-0 animate-spin rounded-full h-16 w-16 border-4 border-transparent border-t-blue-500`"></div>
                </div>
              </div>

              <!-- Progress Text -->
              <h3 :class="`text-xl font-semibold mb-2 transition-colors duration-300 ${
                isDark ? 'text-white' : 'text-gray-800'
              }`">
                AI 正在识别发票信息，请稍候...
              </h3>

              <p :class="`text-sm mb-6 transition-colors duration-300 ${
                isDark ? 'text-gray-400' : 'text-gray-600'
              }`">
                {{ progressText }}
              </p>

              <!-- Progress Bar -->
              <div class="max-w-md mx-auto">
                <div :class="`h-2 rounded-full overflow-hidden ${
                  isDark ? 'bg-gray-700' : 'bg-gray-200'
                }`">
                  <div
                    :class="`h-full transition-all duration-500 ${
                      isDark ? 'bg-gradient-to-r from-blue-500 to-cyan-500' : 'bg-gradient-to-r from-blue-500 to-cyan-500'
                    }`"
                    :style="{ width: `${uploadProgress}%` }"
                  ></div>
                </div>
                <p :class="`text-sm mt-2 transition-colors duration-300 ${
                  isDark ? 'text-gray-400' : 'text-gray-600'
                }`">
                  {{ uploadProgress }}%
                </p>
              </div>
            </div>
          </div>

          <!-- Upload Result -->
          <div
            v-else-if="uploadResult"
            :class="`rounded-2xl p-12 backdrop-blur-md transition-all duration-300 ${
              isDark
                ? 'bg-gray-800/70 border border-gray-700/50'
                : 'bg-white/70 border border-gray-200/50'
            }`"
          >
            <div class="text-center">
              <!-- Success Icon -->
              <div class="flex justify-center mb-6">
                <div class="p-6 rounded-full bg-green-500">
                  <CheckCircleIcon :size="48" class="text-white" />
                </div>
              </div>

              <!-- Success Message -->
              <h3 :class="`text-2xl font-bold mb-6 transition-colors duration-300 ${
                isDark ? 'text-white' : 'text-gray-800'
              }`">
                识别完成！
              </h3>

              <!-- Recognition Results -->
              <div :class="`grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-8`">
                <div :class="`p-4 rounded-lg ${
                  isDark ? 'bg-gray-700/50' : 'bg-gray-100'
                }`">
                  <p :class="`text-sm mb-1 transition-colors duration-300 ${
                    isDark ? 'text-gray-400' : 'text-gray-600'
                  }`">发票金额</p>
                  <p :class="`text-xl font-bold transition-colors duration-300 ${
                    isDark ? 'text-white' : 'text-gray-800'
                  }`">{{ uploadResult.amount }}</p>
                </div>
                <div :class="`p-4 rounded-lg ${
                  isDark ? 'bg-gray-700/50' : 'bg-gray-100'
                }`">
                  <p :class="`text-sm mb-1 transition-colors duration-300 ${
                    isDark ? 'text-gray-400' : 'text-gray-600'
                  }`">发票日期</p>
                  <p :class="`text-xl font-bold transition-colors duration-300 ${
                    isDark ? 'text-white' : 'text-gray-800'
                  }`">{{ uploadResult.date }}</p>
                </div>
                <div :class="`p-4 rounded-lg ${
                  isDark ? 'bg-gray-700/50' : 'bg-gray-100'
                }`">
                  <p :class="`text-sm mb-1 transition-colors duration-300 ${
                    isDark ? 'text-gray-400' : 'text-gray-600'
                  }`">发票类别</p>
                  <p :class="`text-xl font-bold transition-colors duration-300 ${
                    isDark ? 'text-white' : 'text-gray-800'
                  }`">{{ uploadResult.category }}</p>
                </div>
                <div :class="`p-4 rounded-lg ${
                  isDark ? 'bg-gray-700/50' : 'bg-gray-100'
                }`">
                  <p :class="`text-sm mb-1 transition-colors duration-300 ${
                    isDark ? 'text-gray-400' : 'text-gray-600'
                  }`">支付来源</p>
                  <p :class="`text-xl font-bold transition-colors duration-300 ${
                    isDark ? 'text-white' : 'text-gray-800'
                  }`">{{ uploadResult.source }}</p>
                </div>
              </div>

              <!-- Confidence Score -->
              <div class="mb-8">
                <p :class="`text-sm mb-2 transition-colors duration-300 ${
                  isDark ? 'text-gray-400' : 'text-gray-600'
                }`">AI 识别置信度</p>
                <div class="flex items-center justify-center space-x-2">
                  <div :class="`h-3 rounded-full overflow-hidden max-w-xs ${
                    isDark ? 'bg-gray-700' : 'bg-gray-200'
                  }`">
                    <div
                      :class="`h-full transition-all duration-1000 ${
                        uploadResult.confidence >= 90
                          ? 'bg-green-500'
                          : uploadResult.confidence >= 70
                          ? 'bg-yellow-500'
                          : 'bg-red-500'
                      }`"
                      :style="{ width: `${uploadResult.confidence}%` }"
                    ></div>
                  </div>
                  <span :class="`font-bold ${
                    uploadResult.confidence >= 90
                      ? 'text-green-500'
                      : uploadResult.confidence >= 70
                      ? 'text-yellow-500'
                      : 'text-red-500'
                  }`">
                    {{ uploadResult.confidence }}%
                  </span>
                </div>
              </div>

              <!-- Action Buttons -->
              <div class="flex flex-col sm:flex-row gap-4 justify-center">
                <button
                  @click="viewAnalysis"
                  :class="`px-6 py-3 rounded-lg font-semibold transition-all duration-300 transform hover:scale-105 ${
                    isDark
                      ? 'bg-gradient-to-r from-purple-600 to-pink-600 hover:from-purple-500 hover:to-pink-500 text-white shadow-lg shadow-purple-500/25'
                      : 'bg-gradient-to-r from-purple-500 to-pink-500 hover:from-purple-600 hover:to-pink-600 text-white shadow-lg shadow-purple-500/30'
                  }`"
                >
                  查看分析结果
                </button>
                <button
                  @click="resetUpload"
                  :class="`px-6 py-3 rounded-lg font-semibold transition-all duration-300 transform hover:scale-105 ${
                    isDark
                      ? 'bg-gradient-to-r from-blue-600 to-cyan-600 hover:from-blue-500 hover:to-cyan-500 text-white shadow-lg shadow-blue-500/25'
                      : 'bg-gradient-to-r from-blue-500 to-cyan-500 hover:from-blue-600 hover:to-cyan-600 text-white shadow-lg shadow-blue-500/30'
                  }`"
                >
                  继续上传
                </button>
                <button
                  @click="viewDashboard"
                  :class="`px-6 py-3 rounded-lg font-semibold transition-all duration-300 ${
                    isDark
                      ? 'border border-gray-600 hover:bg-gray-700 text-gray-300'
                      : 'border border-gray-300 hover:bg-gray-100 text-gray-700'
                  }`"
                >
                  查看账单
                </button>
              </div>
            </div>
          </div>
        </transition>

        <!-- Upload History -->
        <div v-if="uploadHistory.length > 0" class="mt-12">
          <h3 :class="`text-lg font-semibold mb-6 transition-colors duration-300 ${
            isDark ? 'text-white' : 'text-gray-800'
          }`">
            最近上传记录
          </h3>
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div
              v-for="record in uploadHistory"
              :key="record.id"
              :class="`p-6 rounded-xl border backdrop-blur-md transition-all duration-300 hover:shadow-lg ${
                isDark
                  ? 'bg-gray-800/50 border-gray-700/50 hover:border-cyan-500/30'
                  : 'bg-white/50 border-gray-200/50 hover:border-cyan-400/30'
              }`"
            >
              <div class="flex items-start justify-between mb-4">
                <div :class="`p-2 rounded-lg ${
                  record.status === 'success'
                    ? 'bg-green-100 text-green-600'
                    : record.status === 'processing'
                    ? 'bg-yellow-100 text-yellow-600'
                    : 'bg-red-100 text-red-600'
                }`">
                  <component
                    :is="record.status === 'success' ? CheckCircleIcon : record.status === 'processing' ? ClockIcon : XCircleIcon"
                    :size="20"
                  />
                </div>
                <span :class="`text-xs transition-colors duration-300 ${
                  isDark ? 'text-gray-400' : 'text-gray-500'
                }`">
                  {{ record.time }}
                </span>
              </div>
              <h4 :class="`font-medium mb-2 transition-colors duration-300 ${
                isDark ? 'text-white' : 'text-gray-800'
              }`">
                {{ record.name }}
              </h4>
              <p :class="`text-sm transition-colors duration-300 ${
                isDark ? 'text-gray-400' : 'text-gray-600'
              }`">
                {{ record.category }} · {{ record.amount }} · {{ record.source }}
              </p>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Footer -->
    <footer :class="`relative z-10 backdrop-blur-md border-t transition-all duration-300 ${
      isDark
        ? 'bg-gray-800/70 border-gray-700/50'
        : 'bg-white/70 border-gray-200/50'
    }`">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
        <p :class="`text-center text-sm transition-colors duration-300 ${isDark ? 'text-gray-400' : 'text-gray-600'}`">
          © 2025 AstraFlow · Smart Expense Made Simple
        </p>
      </div>
    </footer>
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
</script>

<style scoped>
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

/* Particle background optimization */
.dark .animate-pulse {
  opacity: 0.3;
}

/* Glow effects */
.hover\:shadow-cyan-500\/10:hover {
  box-shadow: 0 20px 25px -5px rgba(56, 189, 248, 0.15), 0 10px 10px -5px rgba(56, 189, 248, 0.08);
}

.hover\:shadow-cyan-400\/10:hover {
  box-shadow: 0 20px 25px -5px rgba(56, 189, 248, 0.25), 0 10px 10px -5px rgba(56, 189, 248, 0.15);
}

.shadow-cyan-400\/20 {
  box-shadow: 0 25px 50px -12px rgba(56, 189, 248, 0.25);
}

/* Smooth transitions */
.transition-all {
  transition-property: all;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
}
</style>