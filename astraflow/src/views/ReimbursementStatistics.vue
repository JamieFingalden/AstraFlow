<template>
  <div :class="'flex flex-col min-h-screen transition-colors duration-300 ' + (isDark ? 'bg-gray-900' : 'bg-gradient-to-br from-blue-50 via-cyan-50 to-indigo-50')">
    <!-- Particle background effect -->
    <div class="fixed inset-0 overflow-hidden pointer-events-none">
      <div :class="'absolute -top-40 -right-40 w-80 h-80 rounded-full mix-blend-multiply filter blur-xl animate-pulse ' + (isDark ? 'bg-cyan-900/30' : 'bg-cyan-500')"></div>
      <div :class="'absolute -bottom-40 -left-40 w-80 h-80 rounded-full mix-blend-multiply filter blur-xl animate-pulse ' + (isDark ? 'bg-blue-900/30' : 'bg-blue-500')"></div>
      <div :class="'absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-80 h-80 rounded-full mix-blend-multiply filter blur-xl animate-pulse ' + (isDark ? 'bg-purple-900/30' : 'bg-purple-500')"></div>
    </div>

    <!-- Header -->
    <header :class="'relative z-10 backdrop-blur-md border-b transition-all duration-300 ' + (isDark ? 'bg-gray-800/70 border-gray-700/50' : 'bg-white/70 border-gray-200/50')">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center h-16">
          <div class="flex items-center space-x-4">
            <!-- 返回按钮 -->
            <router-link
              to="/"
              :class="'flex items-center space-x-2 px-3 py-2 rounded-lg transition-all duration-200 ' + (isDark ? 'hover:bg-gray-700/50 text-gray-300 hover:text-white' : 'hover:bg-gray-100 text-gray-600 hover:text-gray-900')"
              title="返回首页"
            >
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="m15 18-6-6 6-6"/>
              </svg>
              <span class="hidden sm:inline">返回</span>
            </router-link>

            <div class="text-2xl font-bold bg-gradient-to-r from-cyan-400 to-blue-500 bg-clip-text text-transparent">
              AstraFlow
            </div>
            <h1 :class="'text-xl font-semibold transition-colors duration-300 ' + (isDark ? 'text-white' : 'text-gray-800')">
              报销统计
            </h1>
          </div>

          <div class="flex items-center space-x-4">
            <!-- 主题切换按钮 -->
            <button
              @click="toggleTheme"
              :class="'p-2 rounded-lg transition-all duration-200 ' + (isDark ? 'bg-gray-700/50 hover:bg-gray-600/50 text-gray-300' : 'bg-gray-100 hover:bg-gray-200 text-gray-700')"
            >
              <SunIcon v-if="isDark" :size="20" />
              <MoonIcon v-else :size="20" />
            </button>
          </div>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="relative z-10 flex-1 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- Data Overview Cards -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <!-- 本月总支出 -->
        <div :class="'rounded-2xl p-6 backdrop-blur-md border transition-all duration-300 hover:scale-105 hover:shadow-xl ' + (isDark ? 'bg-gradient-to-br from-red-900/30 to-orange-900/30 border-red-700/50' : 'bg-gradient-to-br from-red-50 to-orange-50 border-red-200/50')">
          <div class="flex items-center justify-between mb-4">
            <div :class="'p-3 rounded-xl ' + (isDark ? 'bg-red-900/50' : 'bg-red-100')">
              <WalletIcon :size="24" :class="isDark ? 'text-red-400' : 'text-red-600'" />
            </div>
            <span :class="'text-xs font-medium px-2 py-1 rounded-full ' + (isDark ? 'bg-red-900/50 text-red-300' : 'bg-red-100 text-red-700')">
              本月
            </span>
          </div>
          <h3 :class="'text-sm font-medium mb-2 ' + (isDark ? 'text-gray-400' : 'text-gray-600')">本月总支出</h3>
          <p :class="'text-2xl font-bold ' + (isDark ? 'text-white' : 'text-gray-900')">¥12,486.50</p>
          <p :class="'text-xs mt-2 ' + (isDark ? 'text-gray-500' : 'text-gray-500')">较上月 +8.5%</p>
        </div>

        <!-- 报销通过金额 -->
        <div :class="'rounded-2xl p-6 backdrop-blur-md border transition-all duration-300 hover:scale-105 hover:shadow-xl ' + (isDark ? 'bg-gradient-to-br from-green-900/30 to-emerald-900/30 border-green-700/50' : 'bg-gradient-to-br from-green-50 to-emerald-50 border-green-200/50')">
          <div class="flex items-center justify-between mb-4">
            <div :class="'p-3 rounded-xl ' + (isDark ? 'bg-green-900/50' : 'bg-green-100')">
              <CheckCircleIcon :size="24" :class="isDark ? 'text-green-400' : 'text-green-600'" />
            </div>
            <span :class="'text-xs font-medium px-2 py-1 rounded-full ' + (isDark ? 'bg-green-900/50 text-green-300' : 'bg-green-100 text-green-700')">
              已通过
            </span>
          </div>
          <h3 :class="'text-sm font-medium mb-2 ' + (isDark ? 'text-gray-400' : 'text-gray-600')">报销通过金额</h3>
          <p :class="'text-2xl font-bold ' + (isDark ? 'text-white' : 'text-gray-900')">¥8,234.00</p>
          <p :class="'text-xs mt-2 ' + (isDark ? 'text-gray-500' : 'text-gray-500')">通过率 66.0%</p>
        </div>

        <!-- 未报销金额 -->
        <div :class="'rounded-2xl p-6 backdrop-blur-md border transition-all duration-300 hover:scale-105 hover:shadow-xl ' + (isDark ? 'bg-gradient-to-br from-yellow-900/30 to-amber-900/30 border-yellow-700/50' : 'bg-gradient-to-br from-yellow-50 to-amber-50 border-yellow-200/50')">
          <div class="flex items-center justify-between mb-4">
            <div :class="'p-3 rounded-xl ' + (isDark ? 'bg-yellow-900/50' : 'bg-yellow-100')">
              <ClockIcon :size="24" :class="isDark ? 'text-yellow-400' : 'text-yellow-600'" />
            </div>
            <span :class="'text-xs font-medium px-2 py-1 rounded-full ' + (isDark ? 'bg-yellow-900/50 text-yellow-300' : 'bg-yellow-100 text-yellow-700')">
              待处理
            </span>
          </div>
          <h3 :class="'text-sm font-medium mb-2 ' + (isDark ? 'text-gray-400' : 'text-gray-600')">未报销金额</h3>
          <p :class="'text-2xl font-bold ' + (isDark ? 'text-white' : 'text-gray-900')">¥4,252.50</p>
          <p :class="'text-xs mt-2 ' + (isDark ? 'text-gray-500' : 'text-gray-500')">34.0% 待报销</p>
        </div>

        <!-- 报销通过率 -->
        <div :class="'rounded-2xl p-6 backdrop-blur-md border transition-all duration-300 hover:scale-105 hover:shadow-xl ' + (isDark ? 'bg-gradient-to-br from-purple-900/30 to-indigo-900/30 border-purple-700/50' : 'bg-gradient-to-br from-purple-50 to-indigo-50 border-purple-200/50')">
          <div class="flex items-center justify-between mb-4">
            <div :class="'p-3 rounded-xl ' + (isDark ? 'bg-purple-900/50' : 'bg-purple-100')">
              <TrendingUpIcon :size="24" :class="isDark ? 'text-purple-400' : 'text-purple-600'" />
            </div>
            <span :class="'text-xs font-medium px-2 py-1 rounded-full ' + (isDark ? 'bg-purple-900/50 text-purple-300' : 'bg-purple-100 text-purple-700')">
              统计
            </span>
          </div>
          <h3 :class="'text-sm font-medium mb-2 ' + (isDark ? 'text-gray-400' : 'text-gray-600')">报销通过率</h3>
          <p :class="'text-2xl font-bold ' + (isDark ? 'text-white' : 'text-gray-900')">66.0%</p>
          <p :class="'text-xs mt-2 ' + (isDark ? 'text-gray-500' : 'text-gray-500')">较上月 +2.3%</p>
        </div>
      </div>

      <!-- Charts Section -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- Monthly Trend Bar Chart -->
        <div :class="'rounded-2xl p-6 backdrop-blur-md border transition-all duration-300 ' + (isDark ? 'bg-gray-800/70 border-gray-700/50' : 'bg-white/70 border-gray-200/50')">
          <h3 :class="'text-lg font-semibold mb-6 ' + (isDark ? 'text-white' : 'text-gray-900')">最近6个月报销趋势</h3>
          <div ref="barChartRef" :style="{ height: '300px' }" class="w-full"></div>
        </div>

        <!-- Category Distribution Pie Chart -->
        <div :class="'rounded-2xl p-6 backdrop-blur-md border transition-all duration-300 ' + (isDark ? 'bg-gray-800/70 border-gray-700/50' : 'bg-white/70 border-gray-200/50')">
          <h3 :class="'text-lg font-semibold mb-6 ' + (isDark ? 'text-white' : 'text-gray-900')">支出分类分布</h3>
          <div ref="pieChartRef" :style="{ height: '300px' }" class="w-full"></div>
        </div>
      </div>
    </main>

    <!-- Footer -->
    <footer :class="'relative z-10 backdrop-blur-md border-t mt-12 transition-all duration-300 ' + (isDark ? 'bg-gray-800/70 border-gray-700/50' : 'bg-white/70 border-gray-200/50')">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
        <p :class="'text-center text-sm transition-colors duration-300 ' + (isDark ? 'text-gray-400' : 'text-gray-600')">
          © 2025 AstraFlow · Smart Expense Made Simple
        </p>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, nextTick, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useTheme } from '../composables/useTheme'
import * as echarts from 'echarts'
import {
  WalletIcon,
  CheckCircleIcon,
  ClockIcon,
  TrendingUpIcon,
  SunIcon,
  MoonIcon
} from 'lucide-vue-next'

const router = useRouter()
const { isDark, toggleTheme } = useTheme()

// Chart refs
const barChartRef = ref(null)
const pieChartRef = ref(null)

// Chart instances
let barChart = null
let pieChart = null

// Mock data for charts
const monthlyData = {
  months: ['6月', '7月', '8月', '9月', '10月', '11月'],
  total: [8500, 9200, 7800, 10500, 11500, 12486.50],
  approved: [5200, 6100, 4800, 6800, 7500, 8234.00],
  pending: [3300, 3100, 3000, 3700, 4000, 4252.50]
}

const categoryData = [
  { name: '餐饮', value: 3856.20, percentage: 30.9 },
  { name: '交通', value: 2847.80, percentage: 22.8 },
  { name: '住宿', value: 2346.50, percentage: 18.8 },
  { name: '办公', value: 1872.30, percentage: 15.0 },
  { name: '其他', value: 1563.70, percentage: 12.5 }
]

// Initialize bar chart
const initBarChart = () => {
  if (!barChartRef.value) return

  barChart = echarts.init(barChartRef.value)

  const option = {
    tooltip: {
      trigger: 'axis',
      axisPointer: {
        type: 'shadow'
      },
      backgroundColor: isDark.value ? '#1f2937' : '#ffffff',
      borderColor: isDark.value ? '#374151' : '#e5e7eb',
      textStyle: {
        color: isDark.value ? '#f3f4f6' : '#111827'
      }
    },
    legend: {
      data: ['总支出', '已报销', '待报销'],
      textStyle: {
        color: isDark.value ? '#f3f4f6' : '#111827'
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
      data: monthlyData.months,
      axisLabel: {
        color: isDark.value ? '#9ca3af' : '#6b7280'
      },
      axisLine: {
        lineStyle: {
          color: isDark.value ? '#374151' : '#e5e7eb'
        }
      }
    },
    yAxis: {
      type: 'value',
      axisLabel: {
        color: isDark.value ? '#9ca3af' : '#6b7280',
        formatter: '¥{value}'
      },
      axisLine: {
        lineStyle: {
          color: isDark.value ? '#374151' : '#e5e7eb'
        }
      },
      splitLine: {
        lineStyle: {
          color: isDark.value ? '#374151' : '#f3f4f6'
        }
      }
    },
    series: [
      {
        name: '总支出',
        type: 'bar',
        data: monthlyData.total,
        itemStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            { offset: 0, color: '#3b82f6' },
            { offset: 1, color: '#1d4ed8' }
          ])
        },
        emphasis: {
          itemStyle: {
            shadowBlur: 10,
            shadowOffsetX: 0,
            shadowColor: 'rgba(59, 130, 246, 0.5)'
          }
        }
      },
      {
        name: '已报销',
        type: 'bar',
        data: monthlyData.approved,
        itemStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            { offset: 0, color: '#10b981' },
            { offset: 1, color: '#059669' }
          ])
        },
        emphasis: {
          itemStyle: {
            shadowBlur: 10,
            shadowOffsetX: 0,
            shadowColor: 'rgba(16, 185, 129, 0.5)'
          }
        }
      },
      {
        name: '待报销',
        type: 'bar',
        data: monthlyData.pending,
        itemStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            { offset: 0, color: '#f59e0b' },
            { offset: 1, color: '#d97706' }
          ])
        },
        emphasis: {
          itemStyle: {
            shadowBlur: 10,
            shadowOffsetX: 0,
            shadowColor: 'rgba(245, 158, 11, 0.5)'
          }
        }
      }
    ]
  }

  barChart.setOption(option)
}

// Initialize pie chart
const initPieChart = () => {
  if (!pieChartRef.value) return

  pieChart = echarts.init(pieChartRef.value)

  const option = {
    tooltip: {
      trigger: 'item',
      backgroundColor: isDark.value ? '#1f2937' : '#ffffff',
      borderColor: isDark.value ? '#374151' : '#e5e7eb',
      textStyle: {
        color: isDark.value ? '#f3f4f6' : '#111827'
      },
      formatter: '{a} <br/>{b}: ¥{c} ({d}%)'
    },
    legend: {
      orient: 'vertical',
      left: 'left',
      textStyle: {
        color: isDark.value ? '#f3f4f6' : '#111827'
      }
    },
    series: [
      {
        name: '支出分布',
        type: 'pie',
        radius: ['40%', '70%'],
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 10,
          borderColor: isDark.value ? '#1f2937' : '#ffffff',
          borderWidth: 2
        },
        label: {
          show: false,
          position: 'center'
        },
        emphasis: {
          label: {
            show: true,
            fontSize: 20,
            fontWeight: 'bold',
            color: isDark.value ? '#f3f4f6' : '#111827'
          }
        },
        labelLine: {
          show: false
        },
        data: categoryData.map((item, index) => ({
          value: item.value,
          name: item.name,
          itemStyle: {
            color: [
              '#ef4444', '#f97316', '#eab308', '#22c55e', '#3b82f6'
            ][index]
          }
        }))
      }
    ]
  }

  pieChart.setOption(option)
}

// Update charts when theme changes
const updateChartsTheme = () => {
  if (barChart) {
    initBarChart()
  }
  if (pieChart) {
    initPieChart()
  }
}

// Handle window resize
const handleResize = () => {
  if (barChart) {
    barChart.resize()
  }
  if (pieChart) {
    pieChart.resize()
  }
}

onMounted(async () => {
  await nextTick()
  initBarChart()
  initPieChart()
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  if (barChart) {
    barChart.dispose()
  }
  if (pieChart) {
    pieChart.dispose()
  }
  window.removeEventListener('resize', handleResize)
})

// Watch for theme changes
const unwatch = watch(isDark, () => {
  updateChartsTheme()
})

// Clean up watcher
onUnmounted(() => {
  if (unwatch) {
    unwatch()
  }
})
</script>

<style scoped>
/* Custom animations */
@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-10px); }
}

.animate-float {
  animation: float 3s ease-in-out infinite;
}

/* Chart container responsive */
@media (max-width: 768px) {
  .grid > div {
    min-height: 250px;
  }
}
</style>