<template>
  <div :class="`min-h-screen transition-colors duration-300 ${isDark ? 'bg-gray-900' : 'bg-gradient-to-br from-blue-50 via-cyan-50 to-indigo-50'}`">
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
            <!-- 返回上传页按钮 -->
            <router-link
              to="/upload"
              :class="`flex items-center space-x-2 px-3 py-2 rounded-lg transition-all duration-200 ${
                isDark
                  ? 'hover:bg-gray-700/50 text-gray-300 hover:text-white'
                  : 'hover:bg-gray-100 text-gray-600 hover:text-gray-900'
              }`"
              title="返回上传页"
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
              AI识别结果 Analysis
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
    <main class="relative z-10 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- Loading State -->
      <div v-if="isLoading" class="flex flex-col items-center justify-center min-h-[60vh]">
        <div class="relative mb-6">
          <div class="animate-spin rounded-full h-16 w-16 border-4 border-blue-200"></div>
          <div :class="`absolute top-0 left-0 animate-spin rounded-full h-16 w-16 border-4 border-transparent border-t-blue-500`"></div>
        </div>
        <h3 :class="`text-xl font-semibold mb-2 transition-colors duration-300 ${
          isDark ? 'text-white' : 'text-gray-800'
        }`">
          AI 正在分析中...
        </h3>
        <p :class="`text-sm transition-colors duration-300 ${
          isDark ? 'text-gray-400' : 'text-gray-600'
        }`">
          请稍候，我们正在深入分析您的账单信息
        </p>
      </div>

      <!-- Analysis Results -->
      <div v-else class="grid grid-cols-1 lg:grid-cols-2 gap-8">
        <!-- Left Side: Recognition Details -->
        <div class="space-y-6">
          <!-- Main Result Card -->
          <div
            :class="`rounded-2xl p-8 backdrop-blur-md border transition-all duration-300 ${
              isDark
                ? 'bg-gray-800/70 border-gray-700/50'
                : 'bg-white/70 border-gray-200/50'
            }`"
          >
            <h2 :class="`text-2xl font-bold mb-6 transition-colors duration-300 ${
              isDark ? 'text-white' : 'text-gray-800'
            }`">
              识别结果摘要
            </h2>

            <div class="grid grid-cols-1 gap-4">
              <!-- Amount -->
              <div
                v-for="(detail, index) in recognitionDetails"
                :key="detail.label"
                :class="`flex items-center p-4 rounded-lg transition-all duration-300 transform hover:scale-[1.02] ${
                  isDark ? 'bg-gray-700/50' : 'bg-gray-50'
                }`"
                :style="{ animationDelay: `${index * 100}ms` }"
              >
                <div :class="`p-3 rounded-lg mr-4 ${
                  detail.iconBg
                }`">
                  <component :is="detail.icon" :size="20" :class="detail.iconColor" />
                </div>
                <div class="flex-grow">
                  <p :class="`text-sm mb-1 transition-colors duration-300 ${
                    isDark ? 'text-gray-400' : 'text-gray-600'
                  }`">
                    {{ detail.label }}
                  </p>
                  <p :class="`text-lg font-semibold transition-colors duration-300 ${
                    isDark ? 'text-white' : 'text-gray-800'
                  }`">
                    {{ detail.value }}
                  </p>
                </div>
              </div>
            </div>
          </div>

          <!-- AI Insights Card -->
          <div
            :class="`rounded-2xl p-6 backdrop-blur-md border transition-all duration-300 ${
              isDark
                ? 'bg-gradient-to-r from-blue-900/30 to-cyan-900/30 border-blue-700/50'
                : 'bg-gradient-to-r from-blue-50/70 to-cyan-50/70 border-blue-200/50'
            }`"
          >
            <div class="flex items-center mb-4">
              <div :class="`p-2 rounded-lg mr-3 ${
                isDark ? 'bg-blue-900/50' : 'bg-blue-100'
              }`">
                <BrainCircuitIcon :size="20" class="text-blue-500" />
              </div>
              <h3 :class="`text-lg font-semibold transition-colors duration-300 ${
                isDark ? 'text-white' : 'text-gray-800'
              }`">
                AI 智能总结
              </h3>
            </div>

            <div class="space-y-3">
              <div
                v-for="insight in aiInsights"
                :key="insight.id"
                :class="`p-3 rounded-lg border-l-4 transition-all duration-300 ${
                  insight.type === 'warning'
                    ? isDark
                      ? 'bg-yellow-900/20 border-yellow-600 text-yellow-300'
                      : 'bg-yellow-50 border-yellow-500 text-yellow-800'
                    : insight.type === 'info'
                    ? isDark
                      ? 'bg-blue-900/20 border-blue-600 text-blue-300'
                      : 'bg-blue-50 border-blue-500 text-blue-800'
                    : isDark
                      ? 'bg-green-900/20 border-green-600 text-green-300'
                      : 'bg-green-50 border-green-500 text-green-800'
                }`"
              >
                <p class="text-sm">{{ insight.message }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Right Side: Data Analysis -->
        <div class="space-y-6">
          <!-- Category Distribution Pie Chart -->
          <div
            :class="`rounded-2xl p-6 backdrop-blur-md border transition-all duration-300 ${
              isDark
                ? 'bg-gray-800/70 border-gray-700/50'
                : 'bg-white/70 border-gray-200/50'
            }`"
          >
            <h3 :class="`text-lg font-semibold mb-4 transition-colors duration-300 ${
              isDark ? 'text-white' : 'text-gray-800'
            }`">
              支出类别占比
            </h3>
            <div class="h-64">
              <v-chart
                :option="categoryChartOption"
                :theme="isDark ? 'dark' : 'light'"
                style="height: 100%; width: 100%;"
              />
            </div>
          </div>

          <!-- 7-Day Trend Line Chart -->
          <div
            :class="`rounded-2xl p-6 backdrop-blur-md border transition-all duration-300 ${
              isDark
                ? 'bg-gray-800/70 border-gray-700/50'
                : 'bg-white/70 border-gray-200/50'
            }`"
          >
            <h3 :class="`text-lg font-semibold mb-4 transition-colors duration-300 ${
              isDark ? 'text-white' : 'text-gray-800'
            }`">
              近7天消费趋势
            </h3>
            <div class="h-64">
              <v-chart
                :option="trendChartOption"
                :theme="isDark ? 'dark' : 'light'"
                style="height: 100%; width: 100%;"
              />
            </div>
          </div>

          <!-- Payment Source Bar Chart -->
          <div
            :class="`rounded-2xl p-6 backdrop-blur-md border transition-all duration-300 ${
              isDark
                ? 'bg-gray-800/70 border-gray-700/50'
                : 'bg-white/70 border-gray-200/50'
            }`"
          >
            <h3 :class="`text-lg font-semibold mb-4 transition-colors duration-300 ${
              isDark ? 'text-white' : 'text-gray-800'
            }`">
              支付来源金额统计
            </h3>
            <div class="h-64">
              <v-chart
                :option="sourceChartOption"
                :theme="isDark ? 'dark' : 'light'"
                style="height: 100%; width: 100%;"
              />
            </div>
          </div>
        </div>
      </div>

      <!-- Action Buttons -->
      <div v-if="!isLoading" class="mt-12 flex flex-col sm:flex-row gap-4 justify-center">
        <button
          @click="exportExpenseReport"
          :class="`px-8 py-4 rounded-lg font-semibold transition-all duration-300 transform hover:scale-105 ${
            isDark
              ? 'bg-gradient-to-r from-blue-600 to-cyan-600 hover:from-blue-500 hover:to-cyan-500 text-white shadow-lg shadow-blue-500/25 hover:shadow-xl hover:shadow-blue-500/30'
              : 'bg-gradient-to-r from-blue-500 to-cyan-500 hover:from-blue-600 hover:to-cyan-600 text-white shadow-lg shadow-blue-500/30 hover:shadow-xl hover:shadow-blue-500/40'
          }`"
        >
          <DownloadIcon :size="20" class="inline mr-2" />
          导出为报销单
        </button>
        <button
          @click="goToUpload"
          :class="`px-8 py-4 rounded-lg font-semibold transition-all duration-300 transform hover:scale-105 ${
            isDark
              ? 'border border-gray-600 hover:bg-gray-700 text-gray-300'
              : 'border border-gray-300 hover:bg-gray-100 text-gray-700'
          }`"
        >
          <UploadIcon :size="20" class="inline mr-2" />
          返回上传页
        </button>
      </div>
    </main>

    <!-- Footer -->
    <footer :class="`relative z-10 backdrop-blur-md border-t mt-12 transition-all duration-300 ${
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

const router = useRouter()

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

// Theme state
const isDark = ref(false)
const isMounted = ref(false)

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
      color: isDark.value ? '#9ca3af' : '#6b7280'
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
        color: isDark.value ? '#374151' : '#e5e7eb'
      }
    },
    axisLabel: {
      color: isDark.value ? '#9ca3af' : '#6b7280'
    }
  },
  yAxis: {
    type: 'value',
    axisLine: {
      lineStyle: {
        color: isDark.value ? '#374151' : '#e5e7eb'
      }
    },
    axisLabel: {
      color: isDark.value ? '#9ca3af' : '#6b7280',
      formatter: '¥{value}'
    },
    splitLine: {
      lineStyle: {
        color: isDark.value ? '#374151' : '#e5e7eb'
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
              color: isDark.value ? 'rgba(59, 130, 246, 0.3)' : 'rgba(59, 130, 246, 0.2)'
            },
            {
              offset: 1,
              color: isDark.value ? 'rgba(59, 130, 246, 0.1)' : 'rgba(59, 130, 246, 0.05)'
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
        color: isDark.value ? '#374151' : '#e5e7eb'
      }
    },
    axisLabel: {
      color: isDark.value ? '#9ca3af' : '#6b7280'
    }
  },
  yAxis: {
    type: 'value',
    axisLine: {
      lineStyle: {
        color: isDark.value ? '#374151' : '#e5e7eb'
      }
    },
    axisLabel: {
      color: isDark.value ? '#9ca3af' : '#6b7280',
      formatter: '¥{value}'
    },
    splitLine: {
      lineStyle: {
        color: isDark.value ? '#374151' : '#e5e7eb'
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

// Theme functions
const toggleTheme = () => {
  isDark.value = !isDark.value
}

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
  isMounted.value = true

  // Check system preference for theme
  const systemPrefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches
  isDark.value = localStorage.getItem('theme') === 'dark' || (systemPrefersDark && !localStorage.getItem('theme'))

  // Apply theme class to document
  document.documentElement.classList.toggle('dark', isDark.value)

  // Fetch analysis data
  await fetchAnalysisResult()
})

// Watch for theme changes
import { watch } from 'vue'

watch(isDark, (newVal) => {
  if (isMounted.value) {
    document.documentElement.classList.toggle('dark', newVal)
    localStorage.setItem('theme', newVal ? 'dark' : 'light')
  }
}, { immediate: true })
</script>

<style scoped>
/* Fade in animation for detail cards */
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

.grid > div {
  animation: fadeInUp 0.6s ease-out forwards;
  opacity: 0;
}

/* Particle background optimization */
.dark .animate-pulse {
  opacity: 0.3;
}

/* Smooth transitions */
.transition-all {
  transition-property: all;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
}

/* Chart container glow effects */
.backdrop-blur-md:hover {
  box-shadow: 0 20px 25px -5px rgba(56, 189, 248, 0.25), 0 10px 10px -5px rgba(56, 189, 248, 0.15);
}

.dark .backdrop-blur-md:hover {
  box-shadow: 0 20px 25px -5px rgba(56, 189, 248, 0.15), 0 10px 10px -5px rgba(56, 189, 248, 0.08);
}
</style>