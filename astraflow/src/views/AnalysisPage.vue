<template>
  <div :class="`app-container ${isDark ? 'dark-theme' : 'light-theme'}`">
    <!-- Particle background effect -->
    <div class="fixed inset-0 overflow-hidden pointer-events-none">
      <div :class="`absolute -top-40 -right-40 particle-cyan ${isDark ? 'dark-particle' : 'light-particle'}`"></div>
      <div :class="`absolute -bottom-40 -left-40 particle-blue ${isDark ? 'dark-particle' : 'light-particle'}`"></div>
      <div :class="`absolute top-1/2 left-1/2 particle-purple ${isDark ? 'dark-particle' : 'light-particle'}`"></div>
    </div>

    <!-- Page Header -->
    <PageHeader
      title="AI识别结果 Analysis"
      :show-back-button="true"
      back-to="/upload"
      :show-theme-toggle="true"
    />

    <!-- Main Content -->
    <main class="relative z-10 main-content">
      <!-- Loading State -->
      <div v-if="isLoading" class="loading-container">
        <div class="spinner-container">
          <div class="spinner-border"></div>
          <div :class="`spinner-inner ${isDark ? 'dark-spinner' : 'light-spinner'}`"></div>
        </div>
        <h3 :class="`loading-title ${isDark ? 'dark-text' : 'light-text'}`">
          AI 正在分析中...
        </h3>
        <p :class="`loading-subtitle ${isDark ? 'dark-muted-text' : 'light-muted-text'}`">
          请稍候，我们正在深入分析您的账单信息
        </p>
      </div>

      <!-- Analysis Results -->
      <div v-else class="results-container">
        <!-- Left Side: Recognition Details -->
        <div class="results-left">
          <!-- Main Result Card -->
          <div
            :class="`result-card ${isDark ? 'dark-result-card' : 'light-result-card'}`"
          >
            <h2 :class="`card-title ${isDark ? 'dark-text' : 'light-text'}`">
              识别结果摘要
            </h2>

            <div class="details-grid">
              <!-- Amount -->
              <div
                v-for="(detail, index) in recognitionDetails"
                :key="detail.label"
                :class="`detail-item ${isDark ? 'dark-detail-item' : 'light-detail-item'}`"
                :style="{ animationDelay: `${index * 100}ms` }"
              >
                <div :class="`detail-icon ${detail.iconBg}`">
                  <component :is="detail.icon" :size="20" :class="detail.iconColor" />
                </div>
                <div class="detail-content">
                  <p :class="`detail-label ${isDark ? 'dark-muted-text' : 'light-muted-text'}`">
                    {{ detail.label }}
                  </p>
                  <p :class="`detail-value ${isDark ? 'dark-text' : 'light-text'}`">
                    {{ detail.value }}
                  </p>
                </div>
              </div>
            </div>
          </div>

          <!-- AI Insights Card -->
          <div
            :class="`insights-card ${isDark ? 'dark-insights-card' : 'light-insights-card'}`"
          >
            <div class="insights-header">
              <div :class="`insights-icon ${isDark ? 'dark-insights-icon' : 'light-insights-icon'}`">
                <BrainCircuitIcon :size="20" class="insights-icon-svg" />
              </div>
              <h3 :class="`insights-title ${isDark ? 'dark-text' : 'light-text'}`">
                AI 智能总结
              </h3>
            </div>

            <div class="insights-list">
              <div
                v-for="insight in aiInsights"
                :key="insight.id"
                :class="`insight-item ${isDark ? 'dark-insight-' + insight.type : 'light-insight-' + insight.type}`"
              >
                <p class="insight-text">{{ insight.message }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Right Side: Data Analysis -->
        <div class="results-right">
          <!-- Category Distribution Pie Chart -->
          <div
            :class="`chart-card ${isDark ? 'dark-chart-card' : 'light-chart-card'}`"
          >
            <h3 :class="`chart-title ${isDark ? 'dark-text' : 'light-text'}`">
              支出类别占比
            </h3>
            <div class="chart-container">
              <v-chart
                :option="categoryChartOption"
                :theme="isDark() ? 'dark' : 'light'"
                style="height: 100%; width: 100%;"
              />
            </div>
          </div>

          <!-- 7-Day Trend Line Chart -->
          <div
            :class="`chart-card ${isDark ? 'dark-chart-card' : 'light-chart-card'}`"
          >
            <h3 :class="`chart-title ${isDark ? 'dark-text' : 'light-text'}`">
              近7天消费趋势
            </h3>
            <div class="chart-container">
              <v-chart
                :option="trendChartOption"
                :theme="isDark() ? 'dark' : 'light'"
                style="height: 100%; width: 100%;"
              />
            </div>
          </div>

          <!-- Payment Source Bar Chart -->
          <div
            :class="`chart-card ${isDark ? 'dark-chart-card' : 'light-chart-card'}`"
          >
            <h3 :class="`chart-title ${isDark ? 'dark-text' : 'light-text'}`">
              支付来源金额统计
            </h3>
            <div class="chart-container">
              <v-chart
                :option="sourceChartOption"
                :theme="isDark() ? 'dark' : 'light'"
                style="height: 100%; width: 100%;"
              />
            </div>
          </div>
        </div>
      </div>

      <!-- Action Buttons -->
      <div v-if="!isLoading" class="action-buttons">
        <button
          @click="exportExpenseReport"
          :class="`primary-btn ${isDark ? 'dark-primary-btn' : 'light-primary-btn'}`"
        >
          <DownloadIcon :size="20" class="btn-icon" />
          导出为报销单
        </button>
        <button
          @click="goToUpload"
          :class="`secondary-btn ${isDark ? 'dark-secondary-btn' : 'light-secondary-btn'}`"
        >
          <UploadIcon :size="20" class="btn-icon" />
          返回上传页
        </button>
      </div>
    </main>

    <!-- Page Footer -->
    <PageFooter />
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { PieChart, LineChart, BarChart } from 'echarts/charts'
import {
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent
} from 'echarts/components'
import VChart from 'vue-echarts'
import {
  DollarSignIcon,
  CalendarIcon,
  TagIcon,
  CreditCardIcon,
  TrendingUpIcon,
  BrainCircuitIcon,
  SunIcon,
  MoonIcon,
  DownloadIcon,
  UploadIcon
} from 'lucide-vue-next'
import { useTheme } from '../composables/useTheme'
import PageHeader from '../components/ui/PageHeader.vue'
import PageFooter from '../components/ui/PageFooter.vue'

const router = useRouter()
const { theme, toggleTheme, isDark } = useTheme()

// Register ECharts components
use([
  CanvasRenderer,
  PieChart,
  LineChart,
  BarChart,
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent
])

// Loading state
const isLoading = ref(true)

// API Configuration
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:5000'
const RESULT_ENDPOINT = '/api/ocr/result'

// Recognition result data (mock data for now)
const recognitionResult = ref({
  amount: 128.50,
  date: '2025-11-08',
  category: '餐饮',
  source: '微信支付',
  confidence: 0.98
})

// Recognition details for display
const recognitionDetails = computed(() => [
  {
    label: '发票金额',
    value: `¥${recognitionResult.value.amount.toFixed(2)}`,
    icon: DollarSignIcon,
    iconBg: 'bg-green-100 dark:bg-green-900/50',
    iconColor: 'text-green-600 dark:text-green-400'
  },
  {
    label: '消费日期',
    value: recognitionResult.value.date,
    icon: CalendarIcon,
    iconBg: 'bg-blue-100 dark:bg-blue-900/50',
    iconColor: 'text-blue-600 dark:text-blue-400'
  },
  {
    label: '消费类别',
    value: recognitionResult.value.category,
    icon: TagIcon,
    iconBg: 'bg-purple-100 dark:bg-purple-900/50',
    iconColor: 'text-purple-600 dark:text-purple-400'
  },
  {
    label: '支付来源',
    value: recognitionResult.value.source,
    icon: CreditCardIcon,
    iconBg: 'bg-orange-100 dark:bg-orange-900/50',
    iconColor: 'text-orange-600 dark:text-orange-400'
  },
  {
    label: '识别置信度',
    value: `${Math.round(recognitionResult.value.confidence * 100)}%`,
    icon: TrendingUpIcon,
    iconBg: 'bg-cyan-100 dark:bg-cyan-900/50',
    iconColor: 'text-cyan-600 dark:text-cyan-400'
  }
])

// AI insights
const aiInsights = ref([
  {
    id: 1,
    message: '系统检测到本月餐饮支出较上月增长25%，建议控制预算',
    type: 'warning'
  },
  {
    id: 2,
    message: 'AI识别到3笔相似金额的报销，可能存在重复提交',
    type: 'info'
  },
  {
    id: 3,
    message: '该笔消费符合正常模式，建议直接提交报销',
    type: 'success'
  }
])

// Chart options
const categoryChartOption = computed(() => ({
  tooltip: {
    trigger: 'item',
    formatter: '{a} <br/>{b}: {c} ({d}%)'
  },
  legend: {
    orient: 'vertical',
    left: 'left',
    textStyle: {
      color: isDark() ? '#9ca3af' : '#6b7280'
    }
  },
  series: [
    {
      name: '支出类别',
      type: 'pie',
      radius: '50%',
      data: [
        { value: 35, name: '餐饮', itemStyle: { color: '#3B82F6' } },
        { value: 25, name: '交通', itemStyle: { color: '#10B981' } },
        { value: 20, name: '办公', itemStyle: { color: '#8B5CF6' } },
        { value: 15, name: '住宿', itemStyle: { color: '#EF4444' } },
        { value: 5, name: '其他', itemStyle: { color: '#F59E0B' } }
      ],
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      }
    }
  ]
}))

const trendChartOption = computed(() => ({
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'cross'
    }
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
  xAxis: {
    type: 'category',
    boundaryGap: false,
    data: ['11-02', '11-03', '11-04', '11-05', '11-06', '11-07', '11-08'],
    axisLine: {
      lineStyle: {
        color: isDark() ? '#374151' : '#e5e7eb'
      }
    },
    axisLabel: {
      color: isDark() ? '#9ca3af' : '#6b7280'
    }
  },
  yAxis: {
    type: 'value',
    axisLine: {
      lineStyle: {
        color: isDark() ? '#374151' : '#e5e7eb'
      }
    },
    axisLabel: {
      color: isDark() ? '#9ca3af' : '#6b7280',
      formatter: '¥{value}'
    },
    splitLine: {
      lineStyle: {
        color: isDark() ? '#374151' : '#e5e7eb'
      }
    }
  },
  series: [
    {
      name: '消费金额',
      type: 'line',
      smooth: true,
      data: [120, 132, 101, 134, 90, 230, 210],
      itemStyle: {
        color: '#3B82F6'
      },
      areaStyle: {
        color: {
          type: 'linear',
          x: 0,
          y: 0,
          x2: 0,
          y2: 1,
          colorStops: [
            {
              offset: 0,
              color: isDark() ? 'rgba(59, 130, 246, 0.3)' : 'rgba(59, 130, 246, 0.2)'
            },
            {
              offset: 1,
              color: isDark() ? 'rgba(59, 130, 246, 0.1)' : 'rgba(59, 130, 246, 0.05)'
            }
          ]
        }
      }
    }
  ]
}))

const sourceChartOption = computed(() => ({
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'shadow'
    }
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
  xAxis: {
    type: 'category',
    data: ['微信支付', '支付宝', '银行卡', '现金'],
    axisLine: {
      lineStyle: {
        color: isDark() ? '#374151' : '#e5e7eb'
      }
    },
    axisLabel: {
      color: isDark() ? '#9ca3af' : '#6b7280'
    }
  },
  yAxis: {
    type: 'value',
    axisLine: {
      lineStyle: {
        color: isDark() ? '#374151' : '#e5e7eb'
      }
    },
    axisLabel: {
      color: isDark() ? '#9ca3af' : '#6b7280',
      formatter: '¥{value}'
    },
    splitLine: {
      lineStyle: {
        color: isDark() ? '#374151' : '#e5e7eb'
      }
    }
  },
  series: [
    {
      name: '消费金额',
      type: 'bar',
      data: [450, 320, 180, 90],
      itemStyle: {
        color: {
          type: 'linear',
          x: 0,
          y: 0,
          x2: 0,
          y2: 1,
          colorStops: [
            {
              offset: 0,
              color: '#06b6d4'
            },
            {
              offset: 1,
              color: '#2563eb'
            }
          ]
        },
        borderRadius: [4, 4, 0, 0]
      },
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowColor: 'rgba(6, 182, 212, 0.5)'
        }
      }
    }
  ]
}))


// Fetch analysis result from API (mock implementation)
const fetchAnalysisResult = async () => {
  try {
    // Future: Replace with actual API call
    // const response = await axios.get(`${API_BASE_URL}${RESULT_ENDPOINT}`)
    // recognitionResult.value = response.data

    // Mock API call for development
    await new Promise(resolve => setTimeout(resolve, 2000))

    // Mock data
    recognitionResult.value = {
      amount: 128.50,
      date: '2025-11-08',
      category: '餐饮',
      source: '微信支付',
      confidence: 0.98
    }

  } catch (error) {
    console.error('Failed to fetch analysis result:', error)
    // Handle error appropriately
  } finally {
    isLoading.value = false
  }
}

// Action functions
const exportExpenseReport = () => {
  // Future: Implement export functionality
  alert('导出功能开发中...')
}

const goToUpload = () => {
  router.push('/upload')
}

// Lifecycle
onMounted(async () => {
  // Fetch analysis data
  await fetchAnalysisResult()
})
</script>

<style scoped>
.app-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  transition: color 0.3s ease, background-color 0.3s ease;
}

/* Dark and light theme backgrounds */
.dark-theme {
  background-color: #111827; /* gray-900 */
}

.light-theme {
  background: linear-gradient(to bottom right, #eff6ff, #ecfeff, #e0e7ff); /* from-blue-50 via-cyan-50 to-indigo-50 */
}

/* Particle background effect */
.fixed {
  position: fixed;
}

.inset-0 {
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
}

.overflow-hidden {
  overflow: hidden;
}

.pointer-events-none {
  pointer-events: none;
}

.absolute {
  position: absolute;
}

.-top-40 {
  top: -10rem;
}

.-right-40 {
  right: -10rem;
}

.-bottom-40 {
  bottom: -10rem;
}

.-left-40 {
  left: -10rem;
}

.top-1\/2 {
  top: 50%;
}

.left-1\/2 {
  left: 50%;
}

.transform {
  transform: translate(0, 0);
}

.-translate-x-1\/2 {
  transform: translateX(-50%);
}

.-translate-y-1\/2 {
  transform: translateY(-50%);
}

.w-80 {
  width: 20rem;
}

.h-80 {
  height: 20rem;
}

.rounded-full {
  border-radius: 9999px;
}

.mix-blend-multiply {
  mix-blend-mode: multiply;
}

.filter {
  filter: blur(0);
}

.blur-xl {
  filter: blur(24px);
}

.animate-pulse {
  animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

.particle-cyan {
  background-color: #0ea5e9; /* cyan-500 */
}

.particle-blue {
  background-color: #3b82f6; /* blue-500 */
}

.particle-purple {
  background-color: #8b5cf6; /* purple-500 */
}

.dark-particle {
  background-color: rgba(14, 159, 233, 0.3); /* cyan-500/30 */
}

.light-particle {
  background-color: #0ea5e9; /* cyan-500 */
}

/* Header styles */
.header {
  position: relative;
  z-index: 10;
  backdrop-filter: blur(12px);
  border-bottom-width: 1px;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.dark-header {
  background-color: rgba(55, 65, 81, 0.7); /* gray-800/70 */
  border-color: rgba(55, 65, 81, 0.5); /* gray-700/50 */
}

.light-header {
  background-color: rgba(255, 255, 255, 0.7); /* white/70 */
  border-color: rgba(229, 231, 235, 0.5); /* gray-200/50 */
}

.container {
  padding-left: 1rem;
  padding-right: 1rem;
  max-width: 89.6rem;
  margin-left: auto;
  margin-right: auto;
  padding-left: 1.5rem;
  padding-right: 1.5rem;
}

@media (min-width: 640px) {
  .container {
    padding-left: 1.5rem;
    padding-right: 1.5rem;
  }
}

@media (min-width: 1024px) {
  .container {
    padding-left: 2rem;
    padding-right: 2rem;
  }
}

.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 4rem;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 1rem;
}

/* Back button */
.back-button {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 0.75rem;
  border-radius: 0.5rem;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
}

.dark-back-button {
  color: #d1d5db; /* gray-300 */
  background-color: transparent;
}

.dark-back-button:hover {
  background-color: rgba(55, 65, 81, 0.5); /* gray-700/50 */
  color: #ffffff; /* white */
}

.light-back-button {
  color: #4b5563; /* gray-600 */
  background-color: transparent;
}

.light-back-button:hover {
  background-color: #f3f4f6; /* gray-100 */
  color: #1f2937; /* gray-900 */
}

.hidden-sm {
  display: none;
}

@media (min-width: 640px) {
  .hidden-sm {
    display: inline;
  }
}

/* Brand gradient */
.brand-gradient {
  font-size: 1.5rem;
  font-weight: 700;
  background: linear-gradient(to right, #22d3ee, #3b82f6); /* from-cyan-400 to-blue-500 */
  background-clip: text;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

/* Page title */
.page-title {
  font-size: 1.25rem;
  font-weight: 600;
  transition: color 0.3s ease;
}

.dark-text {
  color: #ffffff; /* white */
}

.light-text {
  color: #1f2937; /* gray-800 */
}

/* Theme toggle button */
.theme-toggle-btn {
  padding: 0.5rem;
  border-radius: 0.5rem;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
}

.dark-theme-toggle {
  background-color: rgba(55, 65, 81, 0.5); /* gray-700/50 */
}

.dark-theme-toggle:hover {
  background-color: rgba(75, 85, 99, 0.5); /* gray-600/50 */
  color: #d1d5db; /* gray-300 */
}

.light-theme-toggle {
  background-color: #f3f4f6; /* gray-100 */
}

.light-theme-toggle:hover {
  background-color: #e5e7eb; /* gray-200 */
  color: #374151; /* gray-700 */
}

/* Main content */
.main-content {
  position: relative;
  z-index: 10;
  flex: 1;
  padding: 2rem 1rem;
  max-width: 89.6rem;
  margin-left: auto;
  margin-right: auto;
  padding-left: 1.5rem;
  padding-right: 1.5rem;
}

@media (min-width: 640px) {
  .main-content {
    padding-left: 1.5rem;
    padding-right: 1.5rem;
  }
}

@media (min-width: 1024px) {
  .main-content {
    padding-left: 2rem;
    padding-right: 2rem;
  }
}

/* Loading styles */
.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 60vh;
}

.spinner-container {
  position: relative;
  margin-bottom: 1.5rem;
}

.spinner-border {
  width: 4rem;
  height: 4rem;
  border: 4px solid #e5e7eb; /* blue-200 */
  border-radius: 9999px;
  animation: spin 1s linear infinite;
}

.spinner-inner {
  position: absolute;
  top: 0;
  left: 0;
  width: 4rem;
  height: 4rem;
  border: 4px solid transparent;
  border-top-color: #3b82f6; /* blue-500 */
  border-radius: 9999px;
  animation: spin 1s linear infinite;
}

.dark-spinner {
  border-top-color: #3b82f6; /* blue-500 */
}

.light-spinner {
  border-top-color: #3b82f6; /* blue-500 */
}

.loading-title {
  font-size: 1.25rem;
  font-weight: 600;
  margin-bottom: 0.5rem;
  transition: color 0.3s ease;
}

.loading-subtitle {
  font-size: 0.875rem;
  transition: color 0.3s ease;
}

.dark-muted-text {
  color: #9ca3af; /* gray-400 */
}

.light-muted-text {
  color: #4b5563; /* gray-600 */
}

/* Results container */
.results-container {
  display: grid;
  grid-template-columns: 1fr;
  gap: 2rem;
}

@media (min-width: 1024px) {
  .results-container {
    grid-template-columns: 1fr 1fr;
  }
}

.results-left {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.results-right {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

/* Result card */
.result-card {
  border-radius: 1rem;
  padding: 2rem;
  backdrop-filter: blur(12px);
  border: 1px solid;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.dark-result-card {
  background-color: rgba(55, 65, 81, 0.7); /* gray-800/70 */
  border-color: rgba(55, 65, 81, 0.5); /* gray-700/50 */
}

.light-result-card {
  background-color: rgba(255, 255, 255, 0.7); /* white/70 */
  border-color: rgba(229, 231, 235, 0.5); /* gray-200/50 */
}

.card-title {
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 1.5rem;
  transition: color 0.3s ease;
}

/* Details grid */
.details-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1rem;
}

/* Detail item */
.detail-item {
  display: flex;
  align-items: center;
  padding: 1rem;
  border-radius: 0.5rem;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  transform: scale(1);
  animation: fadeInUp 0.6s ease-out;
  opacity: 0;
}

.detail-item:hover {
  transform: scale(1.02);
}

.dark-detail-item {
  background-color: rgba(55, 65, 81, 0.5); /* gray-700/50 */
}

.light-detail-item {
  background-color: #f9fafb; /* gray-50 */
}

.detail-icon {
  padding: 0.75rem;
  border-radius: 0.5rem;
  margin-right: 1rem;
}

.detail-content {
  flex: 1;
}

.detail-label {
  font-size: 0.875rem;
  margin-bottom: 0.25rem;
  transition: color 0.3s ease;
}

.detail-value {
  font-size: 1.125rem;
  font-weight: 600;
  transition: color 0.3s ease;
}

/* Insights card */
.insights-card {
  border-radius: 1rem;
  padding: 1.5rem;
  backdrop-filter: blur(12px);
  border: 1px solid;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.dark-insights-card {
  background: linear-gradient(90deg, rgba(30, 58, 138, 0.3), rgba(6, 182, 212, 0.3)); /* from-blue-900/30 to-cyan-900/30 */
  border-color: rgba(37, 99, 235, 0.5); /* blue-700/50 */
}

.light-insights-card {
  background: linear-gradient(90deg, rgba(239, 246, 255, 0.7), rgba(224, 231, 255, 0.7)); /* from-blue-50/70 to-cyan-50/70 */
  border-color: rgba(209, 213, 219, 0.5); /* blue-200/50 */
}

.insights-header {
  display: flex;
  align-items: center;
  margin-bottom: 1rem;
}

.insights-icon {
  padding: 0.5rem;
  border-radius: 0.5rem;
  margin-right: 0.75rem;
}

.dark-insights-icon {
  background-color: rgba(30, 58, 138, 0.5); /* blue-900/50 */
}

.light-insights-icon {
  background-color: #dbeafe; /* blue-100 */
}

.insights-icon-svg {
  color: #3b82f6; /* blue-500 */
}

.insights-title {
  font-size: 1.125rem;
  font-weight: 600;
  transition: color 0.3s ease;
}

.insights-list {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.insight-item {
  padding: 0.75rem;
  border-radius: 0.5rem;
  border-left-width: 4px;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.dark-insight-warning {
  background-color: rgba(146, 64, 14, 0.2); /* yellow-900/20 */
  border-color: #d97706; /* yellow-600 */
  color: #fbbf24; /* yellow-300 */
}

.light-insight-warning {
  background-color: #fefce8; /* yellow-50 */
  border-color: #eab308; /* yellow-500 */
  color: #92400e; /* yellow-800 */
}

.dark-insight-info {
  background-color: rgba(30, 58, 138, 0.2); /* blue-900/20 */
  border-color: #3730a3; /* blue-600 */
  color: #93c5fd; /* blue-300 */
}

.light-insight-info {
  background-color: #eff6ff; /* blue-50 */
  border-color: #3b82f6; /* blue-500 */
  color: #1e40af; /* blue-800 */
}

.dark-insight-success {
  background-color: rgba(21, 128, 61, 0.2); /* green-900/20 */
  border-color: #16a34a; /* green-600 */
  color: #4ade80; /* green-300 */
}

.light-insight-success {
  background-color: #f0fdf4; /* green-50 */
  border-color: #22c55e; /* green-500 */
  color: #166534; /* green-800 */
}

/* Chart card */
.chart-card {
  border-radius: 1rem;
  padding: 1.5rem;
  backdrop-filter: blur(12px);
  border: 1px solid;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.dark-chart-card {
  background-color: rgba(55, 65, 81, 0.7); /* gray-800/70 */
  border-color: rgba(55, 65, 81, 0.5); /* gray-700/50 */
}

.light-chart-card {
  background-color: rgba(255, 255, 255, 0.7); /* white/70 */
  border-color: rgba(229, 231, 235, 0.5); /* gray-200/50 */
}

.chart-title {
  font-size: 1.125rem;
  font-weight: 600;
  margin-bottom: 1rem;
  transition: color 0.3s ease;
}

.chart-container {
  height: 16rem;
}

/* Action buttons */
.action-buttons {
  display: flex;
  flex-direction: column;
  justify-content: center;
  gap: 1rem;
  margin-top: 3rem;
}

@media (min-width: 640px) {
  .action-buttons {
    flex-direction: row;
  }
}

.primary-btn {
  padding: 2rem 2rem;
  border-radius: 0.5rem;
  font-weight: 600;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  transform: scale(1);
}

.primary-btn:hover {
  transform: scale(1.05);
}

.dark-primary-btn {
  background: linear-gradient(90deg, #2563eb, #0891b2); /* from-blue-600 to-cyan-600 */
  color: #ffffff; /* white */
  box-shadow: 0 8px 20px rgba(59, 130, 246, 0.25);
}

.dark-primary-btn:hover {
  background: linear-gradient(90deg, #3b82f6, #0ea5e9); /* from-blue-500 to-cyan-500 */
  box-shadow: 0 12px 25px rgba(59, 130, 246, 0.3);
}

.light-primary-btn {
  background: linear-gradient(90deg, #3b82f6, #06b6d4); /* from-blue-500 to-cyan-500 */
  color: #ffffff; /* white */
  box-shadow: 0 8px 20px rgba(59, 130, 246, 0.3);
}

.light-primary-btn:hover {
  background: linear-gradient(90deg, #2563eb, #0891b2); /* from-blue-600 to-cyan-600 */
  box-shadow: 0 12px 25px rgba(59, 130, 246, 0.4);
}

.secondary-btn {
  padding: 2rem 2rem;
  border-radius: 0.5rem;
  font-weight: 600;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  transform: scale(1);
}

.secondary-btn:hover {
  transform: scale(1.05);
}

.dark-secondary-btn {
  border: 1px solid #4b5563; /* border-gray-600 */
  background-color: transparent;
  color: #d1d5db; /* gray-300 */
}

.dark-secondary-btn:hover {
  background-color: #374151; /* gray-700 */
}

.light-secondary-btn {
  border: 1px solid #d1d5db; /* border-gray-300 */
  background-color: transparent;
  color: #374151; /* gray-700 */
}

.light-secondary-btn:hover {
  background-color: #f3f4f6; /* gray-100 */
}

.btn-icon {
  display: inline;
  margin-right: 0.5rem;
}


/* Animations */
@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

@keyframes pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.5;
  }
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Smooth transitions */
.transition-all {
  transition-property: all;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
}

/* Chart container glow effects on hover */
.backdrop-blur-md:hover {
  box-shadow: 0 20px 25px -5px rgba(56, 189, 248, 0.25), 0 10px 10px -5px rgba(56, 189, 248, 0.15);
}

.dark .backdrop-blur-md:hover {
  box-shadow: 0 20px 25px -5px rgba(56, 189, 248, 0.15), 0 10px 10px -5px rgba(56, 189, 248, 0.08);
}
</style>