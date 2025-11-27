<template>
  <div :class="'stats-main-container ' + (isDark ? 'dark-theme' : 'light-theme')">
    <!-- Particle background effect -->
    <div class="fixed-container">
      <div :class="'particle-circle particle-1 ' + (isDark ? 'dark-particle' : 'light-particle')"></div>
      <div :class="'particle-circle particle-2 ' + (isDark ? 'dark-particle' : 'light-particle')"></div>
      <div :class="'particle-circle particle-3 ' + (isDark ? 'dark-particle' : 'light-particle')"></div>
    </div>

    <!-- Header -->
    <header :class="'stats-header ' + (isDark ? 'dark-theme' : 'light-theme')">
      <div class="container">
        <div class="header-content">
          <div class="header-left">
            <!-- 返回按钮 -->
            <router-link
              to="/"
              :class="'back-button ' + (isDark ? 'dark-theme' : 'light-theme')"
              title="返回首页"
            >
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="m15 18-6-6 6-6"/>
              </svg>
              <span class="back-button-text">返回</span>
            </router-link>

            <div class="brand-text">
              AstraFlow
            </div>
            <h1 :class="'page-title ' + (isDark ? 'dark-theme' : 'light-theme')">
              报销统计
            </h1>
          </div>

          <div class="header-right">
            <!-- 主题切换按钮 -->
            <button
              @click="toggleTheme"
              :class="'theme-toggle-btn ' + (isDark ? 'dark-theme' : 'light-theme')"
            >
              <SunIcon v-if="isDark" :size="20" />
              <MoonIcon v-else :size="20" />
            </button>
          </div>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="stats-main-content">
      <!-- Data Overview Cards -->
      <div class="stats-cards-grid">
        <!-- 总支出 -->
        <div :class="'stats-card ' + (isDark ? 'dark-theme red-theme' : 'light-theme red-theme')">
          <div class="card-header">
            <div :class="'card-icon-container ' + (isDark ? 'dark-theme red-theme' : 'light-theme red-theme')">
              <WalletIcon :size="24" :class="isDark ? 'card-icon-dark' : 'card-icon-light'" />
            </div>
          </div>
          <h3 :class="'card-subtitle ' + (isDark ? 'dark-theme' : 'light-theme')">总支出</h3>
          <p :class="'card-value ' + (isDark ? 'dark-theme' : 'light-theme')">¥{{ monthlyExpense ? monthlyExpense.toFixed(2) : '0.00' }}</p>
          <p :class="'card-change ' + (isDark ? 'dark-theme' : 'light-theme')">&nbsp;</p>
        </div>

        <!-- 报销通过金额 -->
        <div :class="'stats-card ' + (isDark ? 'dark-theme green-theme' : 'light-theme green-theme')">
          <div class="card-header">
            <div :class="'card-icon-container ' + (isDark ? 'dark-theme green-theme' : 'light-theme green-theme')">
              <CheckCircleIcon :size="24" :class="isDark ? 'card-icon-dark' : 'card-icon-light'" />
            </div>
            <span :class="'card-badge ' + (isDark ? 'dark-theme green-theme' : 'light-theme green-theme')">
              已通过
            </span>
          </div>
          <h3 :class="'card-subtitle ' + (isDark ? 'dark-theme' : 'light-theme')">报销通过金额</h3>
          <p :class="'card-value ' + (isDark ? 'dark-theme' : 'light-theme')">¥{{ approvedAmount ? approvedAmount.toFixed(2) : '0.00' }}</p>
          <p :class="'card-change ' + (isDark ? 'dark-theme' : 'light-theme')">通过率 {{ approvalRate?.toFixed(1) || '0.0' }}%</p>
        </div>

        <!-- 未报销金额 -->
        <div :class="'stats-card ' + (isDark ? 'dark-theme yellow-theme' : 'light-theme yellow-theme')">
          <div class="card-header">
            <div :class="'card-icon-container ' + (isDark ? 'dark-theme yellow-theme' : 'light-theme yellow-theme')">
              <ClockIcon :size="24" :class="isDark ? 'card-icon-dark' : 'card-icon-light'" />
            </div>
            <span :class="'card-badge ' + (isDark ? 'dark-theme yellow-theme' : 'light-theme yellow-theme')">
              待处理
            </span>
          </div>
          <h3 :class="'card-subtitle ' + (isDark ? 'dark-theme' : 'light-theme')">未报销金额</h3>
          <p :class="'card-value ' + (isDark ? 'dark-theme' : 'light-theme')">¥{{ pendingAmount ? pendingAmount.toFixed(2) : '0.00' }}</p>
          <p :class="'card-change ' + (isDark ? 'dark-theme' : 'light-theme')">34.0% 待报销</p>
        </div>

        <!-- 报销通过率 -->
        <div :class="'stats-card ' + (isDark ? 'dark-theme purple-theme' : 'light-theme purple-theme')">
          <div class="card-header">
            <div :class="'card-icon-container ' + (isDark ? 'dark-theme purple-theme' : 'light-theme purple-theme')">
              <TrendingUpIcon :size="24" :class="isDark ? 'card-icon-dark' : 'card-icon-light'" />
            </div>
            <span :class="'card-badge ' + (isDark ? 'dark-theme purple-theme' : 'light-theme purple-theme')">
              统计
            </span>
          </div>
          <h3 :class="'card-subtitle ' + (isDark ? 'dark-theme' : 'light-theme')">报销通过率</h3>
          <p :class="'card-value ' + (isDark ? 'dark-theme' : 'light-theme')">{{ approvalRate?.toFixed(1) || '0.0' }}%</p>
          <p :class="'card-change ' + (isDark ? 'dark-theme' : 'light-theme')">较上月 +2.3%</p>
        </div>
      </div>

      <!-- Charts Section -->
      <div class="charts-grid">
        <!-- Monthly Trend Bar Chart -->
        <div :class="'chart-card ' + (isDark ? 'dark-theme' : 'light-theme')">
          <h3 :class="'chart-title ' + (isDark ? 'dark-theme' : 'light-theme')">最近6个月报销趋势</h3>
          <div ref="barChartRef" class="chart-container"></div>
        </div>

        <!-- Category Distribution Pie Chart -->
        <div :class="'chart-card ' + (isDark ? 'dark-theme' : 'light-theme')">
          <h3 :class="'chart-title ' + (isDark ? 'dark-theme' : 'light-theme')">支出分类分布</h3>
          <div ref="pieChartRef" class="chart-container"></div>
        </div>
      </div>
    </main>

    <!-- Footer -->
    <footer :class="'stats-footer ' + (isDark ? 'dark-theme' : 'light-theme')">
      <div class="footer-container">
        <div class="footer-content">
          <p :class="'footer-text ' + (isDark ? 'dark-theme' : 'light-theme')">
            © 2025 AstraFlow · Smart Expense Made Simple
          </p>
        </div>
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
import { getReimbursementStatisticsData } from '../services/api/analyticsApi'

const router = useRouter()
const { isDark, toggleTheme } = useTheme()

// Chart refs
const barChartRef = ref(null)
const pieChartRef = ref(null)

// Chart instances
let barChart = null
let pieChart = null

// State for loading and data
const loading = ref(false)
const error = ref(null)

// Reactive variables for data
const monthlyExpense = ref(0)
const approvedAmount = ref(0)
const pendingAmount = ref(0)
const approvalRate = ref(0)

// For charts
const monthlyData = ref({
  months: ['1月', '2月', '3月', '4月', '5月', '6月', '7月', '8月', '9月', '10月', '11月', '12月'],
  total: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
  approved: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
  pending: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
})

const categoryData = ref([])

// Fetch reimbursement statistics data from API
const fetchStatisticsData = async () => {
  loading.value = true
  error.value = null

  try {
    const response = await getReimbursementStatisticsData()
    console.log('Reimbursement Stats API Response:', response) // Debug log

    if (response.code === 200 && response.data && response.data.reimbursement_statistics) {
      const reimbursementData = response.data.reimbursement_statistics
      const statistics = reimbursementData.statistics
      const monthlyTrends = reimbursementData.monthly_trends || []
      const categories = reimbursementData.categories || []

      console.log('Reimbursement data processed:', reimbursementData) // Debug log

      // Update statistics cards
      monthlyExpense.value = statistics.total_amount || 0
      approvedAmount.value = statistics.approved_amount || 0
      pendingAmount.value = statistics.pending_amount || 0
      approvalRate.value = statistics.approval_rate || 0

      console.log('Updated values:', { monthlyExpense: monthlyExpense.value, approvedAmount: approvedAmount.value, pendingAmount: pendingAmount.value, approvalRate: approvalRate.value }) // Debug log

      // Update category data for pie chart
      categoryData.value = categories.map(cat => ({
        name: cat.name,
        value: parseFloat(cat.value.toFixed(2)),
        percentage: monthlyExpense.value ? (cat.value / monthlyExpense.value * 100).toFixed(1) : 0
      }))

      console.log('Category data updated:', categoryData.value) // Debug log

      // Update monthly data for bar chart
      monthlyData.value = {
        months: monthlyTrends.map(trend => trend.month) || [],
        total: monthlyTrends.map(trend => trend.total) || [],
        approved: monthlyTrends.map(trend => trend.approved) || [],
        pending: monthlyTrends.map(trend => trend.pending) || []
      }

      console.log('Monthly data updated:', monthlyData.value) // Debug log
    } else {
      error.value = response.message || '获取报销统计信息失败'
      console.error('Reimbursement Stats API response error:', response) // Debug log
    }
  } catch (err) {
    console.error('Error fetching reimbursement statistics data:', err)
    error.value = '获取数据时发生错误，请稍后重试'
  } finally {
    loading.value = false
  }
}

// Initialize bar chart
const initBarChart = () => {
  if (!barChartRef.value) return

  barChart = echarts.init(barChartRef.value, isDark.value ? 'dark' : 'default')

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
      },
      formatter: (params) => {
        const date = params[0].axisValue
        let tooltip = `${date}<br/>`
        params.forEach(param => {
          const value = parseFloat(param.value).toFixed(2)
          tooltip += `${param.seriesName}: ¥${value}<br/>`
        })
        return tooltip
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
      data: monthlyData.value.months,
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
        formatter: (value) => `¥${value.toFixed(2)}`
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
        data: monthlyData.value.total,
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
        data: monthlyData.value.approved,
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
        data: monthlyData.value.pending,
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

  pieChart = echarts.init(pieChartRef.value, isDark.value ? 'dark' : 'default')

  const option = {
    tooltip: {
      trigger: 'item',
      backgroundColor: isDark.value ? '#1f2937' : '#ffffff',
      borderColor: isDark.value ? '#374151' : '#e5e7eb',
      textStyle: {
        color: isDark.value ? '#f3f4f6' : '#111827'
      },
      formatter: (params) => {
        const value = parseFloat(params.value).toFixed(2)
        return `${params.seriesName} <br/>${params.name}: ¥${value} (${params.percent}%)`
      }
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
        data: categoryData.value.map((item, index) => ({
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
const updateChartsTheme = async () => {
  if (barChart) {
    barChart.dispose()
  }
  if (pieChart) {
    pieChart.dispose()
  }
  await nextTick()
  initBarChart()
  initPieChart()
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
  await fetchStatisticsData() // Fetch data first
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
/* 主容器 */
.stats-main-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  transition: background-color 0.3s ease, color 0.3s ease;
}

.stats-main-container.dark-theme {
  background-color: #111827;
}

.stats-main-container.light-theme {
  background: linear-gradient(135deg, #eff6ff, #f0f9ff, #e0e7ff);
}

/* 固定容器 */
.fixed-container {
  position: fixed;
  inset: 0;
  overflow: hidden;
  pointer-events: none;
}

/* 粒子圆圈 */
.particle-circle {
  position: absolute;
  width: 20rem;
  height: 20rem;
  border-radius: 50%;
  mix-blend-mode: multiply;
  filter: blur(6rem);
  animation: pulse 4s ease-in-out infinite;
}

.particle-1 {
  top: -10rem;
  right: -10rem;
}

.particle-2 {
  bottom: -10rem;
  left: -10rem;
}

.particle-3 {
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.particle-circle.dark-particle {
  background-color: rgba(6, 182, 212, 0.3);
}

.particle-circle.light-particle {
  background-color: #06b6d4;
}

@keyframes pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.8;
  }
}

/* 头部 */
.stats-header {
  position: relative;
  z-index: 10;
  backdrop-filter: blur(12px);
  border-bottom: 1px solid;
  transition: all 0.3s ease;
}

.stats-header.dark-theme {
  background-color: rgba(31, 41, 55, 0.7);
  border-color: rgba(55, 65, 81, 0.5);
}

.stats-header.light-theme {
  background-color: rgba(255, 255, 255, 0.7);
  border-color: rgba(229, 231, 235, 0.5);
}

.container {
  max-width: 80rem;
  margin: 0 auto;
  padding: 0 1rem;
}

@media (min-width: 640px) {
  .container {
    padding: 0 1.5rem;
  }
}

@media (min-width: 1024px) {
  .container {
    padding: 0 2rem;
  }
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 4rem;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 1rem;
}

/* 返回按钮 */
.back-button {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 0.75rem;
  border-radius: 0.5rem;
  transition: all 0.2s ease;
  text-decoration: none;
}

.back-button.dark-theme {
  color: #d1d5db;
}

.back-button.light-theme {
  color: #4b5563;
}

.back-button:hover.dark-theme {
  background-color: rgba(55, 65, 81, 0.5);
  color: #ffffff;
}

.back-button:hover.light-theme {
  background-color: #f3f4f6;
  color: #1f2937;
}

.back-button-text {
  display: none;
}

@media (min-width: 640px) {
  .back-button-text {
    display: inline;
  }
}

/* 品牌文字 */
.brand-text {
  font-size: 1.5rem;
  font-weight: 700;
  background: linear-gradient(135deg, #06b6d4, #3b82f6);
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
}

/* 页面标题 */
.page-title {
  font-size: 1.25rem;
  font-weight: 600;
  transition: color 0.3s ease;
}

.page-title.dark-theme {
  color: #ffffff;
}

.page-title.light-theme {
  color: #1f2937;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 1rem;
}

/* 主题切换按钮 */
.theme-toggle-btn {
  padding: 0.5rem;
  border-radius: 0.5rem;
  transition: all 0.2s ease;
  border: none;
  cursor: pointer;
}

.theme-toggle-btn.dark-theme {
  background-color: rgba(55, 65, 81, 0.5);
  color: #d1d5db;
}

.theme-toggle-btn.light-theme {
  background-color: #f3f4f6;
  color: #374151;
}

.theme-toggle-btn:hover.dark-theme {
  background-color: rgba(75, 85, 99, 0.5);
}

.theme-toggle-btn:hover.light-theme {
  background-color: #e5e7eb;
}

/* 主内容 */
.stats-main-content {
  position: relative;
  z-index: 10;
  flex: 1;
  width: 80rem;
  margin: 0 auto;
  padding: 2rem 1rem;
}

@media (min-width: 640px) {
  .stats-main-content {
    padding: 2rem 1.5rem;
  }
}

@media (min-width: 1024px) {
  .stats-main-content {
    padding: 2rem 2rem;
  }
}

/* 统计卡片网格 */
.stats-cards-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1.5rem;
  margin-bottom: 2rem;
}

@media (min-width: 768px) {
  .stats-cards-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (min-width: 1024px) {
  .stats-cards-grid {
    grid-template-columns: repeat(4, 1fr);
  }
}

/* 统计卡片 */
.stats-card {
  border-radius: 1rem;
  padding: 1.5rem;
  border: 1px solid;
  transition: all 0.3s ease;
  backdrop-filter: blur(12px);
  cursor: pointer;
}

.stats-card:hover {
  transform: scale(1.05);
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
}

.stats-card.dark-theme {
  background-color: rgba(31, 41, 55, 0.7);
  border-color: rgba(55, 65, 81, 0.5);
}

.stats-card.light-theme {
  background-color: rgba(255, 255, 255, 0.7);
  border-color: rgba(229, 231, 235, 0.5);
}

/* 红色主题 */
.stats-card.red-theme.dark-theme {
  background: linear-gradient(135deg, rgba(127, 29, 29, 0.3), rgba(146, 39, 64, 0.3));
  border-color: rgba(209, 55, 78, 0.5);
}

.stats-card.red-theme.light-theme {
  background: linear-gradient(135deg, #fef2f2, #fff7f7);
  border-color: rgba(254, 202, 202, 0.5);
}

/* 绿色主题 */
.stats-card.green-theme.dark-theme {
  background: linear-gradient(135deg, rgba(21, 128, 61, 0.3), rgba(22, 101, 52, 0.3));
  border-color: rgba(74, 222, 128, 0.5);
}

.stats-card.green-theme.light-theme {
  background: linear-gradient(135deg, #f0fdf4, #ecfdf5);
  border-color: rgba(187, 247, 208, 0.5);
}

/* 黄色主题 */
.stats-card.yellow-theme.dark-theme {
  background: linear-gradient(135deg, rgba(161, 98, 7, 0.3), rgba(180, 83, 9, 0.3));
  border-color: rgba(253, 230, 138, 0.5);
}

.stats-card.yellow-theme.light-theme {
  background: linear-gradient(135deg, #fefce8, #fef9c3);
  border-color: rgba(254, 240, 138, 0.5);
}

/* 紫色主题 */
.stats-card.purple-theme.dark-theme {
  background: linear-gradient(135deg, rgba(88, 28, 135, 0.3), rgba(76, 29, 149, 0.3));
  border-color: rgba(165, 180, 252, 0.5);
}

.stats-card.purple-theme.light-theme {
  background: linear-gradient(135deg, #f5f3ff, #faf5ff);
  border-color: rgba(224, 231, 255, 0.5);
}

/* 卡片头部 */
.card-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 1rem;
}

/* 卡片图标容器 */
.card-icon-container {
  padding: 0.75rem;
  border-radius: 0.75rem;
}

.card-icon-container.dark-theme.red-theme {
  background-color: rgba(127, 29, 29, 0.5);
}

.card-icon-container.light-theme.red-theme {
  background-color: #fee2e2;
}

.card-icon-container.dark-theme.green-theme {
  background-color: rgba(21, 128, 61, 0.5);
}

.card-icon-container.light-theme.green-theme {
  background-color: #dcfce7;
}

.card-icon-container.dark-theme.yellow-theme {
  background-color: rgba(161, 98, 7, 0.5);
}

.card-icon-container.light-theme.yellow-theme {
  background-color: #fef9c3;
}

.card-icon-container.dark-theme.purple-theme {
  background-color: rgba(88, 28, 135, 0.5);
}

.card-icon-container.light-theme.purple-theme {
  background-color: #e0e7ff;
}

/* 卡片徽章 */
.card-badge {
  font-size: 0.75rem;
  font-weight: 500;
  padding: 0.25rem 0.5rem;
  border-radius: 9999px;
}

.card-badge.dark-theme.red-theme {
  background-color: rgba(127, 29, 29, 0.5);
  color: #fcd34d;
}

.card-badge.light-theme.red-theme {
  background-color: #fee2e2;
  color: #b91c1c;
}

.card-badge.dark-theme.green-theme {
  background-color: rgba(21, 128, 61, 0.5);
  color: #86efac;
}

.card-badge.light-theme.green-theme {
  background-color: #dcfce7;
  color: #166534;
}

.card-badge.dark-theme.yellow-theme {
  background-color: rgba(161, 98, 7, 0.5);
  color: #fde68a;
}

.card-badge.light-theme.yellow-theme {
  background-color: #fef9c3;
  color: #854d0e;
}

.card-badge.dark-theme.purple-theme {
  background-color: rgba(88, 28, 135, 0.5);
  color: #c7d2fe;
}

.card-badge.light-theme.purple-theme {
  background-color: #e0e7ff;
  color: #4f46e5;
}

/* 卡片子标题 */
.card-subtitle {
  font-size: 0.875rem;
  font-weight: 500;
  margin-bottom: 0.5rem;
}

.card-subtitle.dark-theme {
  color: #9ca3af;
}

.card-subtitle.light-theme {
  color: #4b5563;
}

/* 卡片值 */
.card-value {
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 0.25rem;
}

.card-value.dark-theme {
  color: #ffffff;
}

.card-value.light-theme {
  color: #1f2937;
}

/* 卡片变化 */
.card-change {
  font-size: 0.75rem;
  margin-top: 0.5rem;
}

.card-change.dark-theme {
  color: #6b7280;
}

.card-change.light-theme {
  color: #6b7280;
}

/* 图标类 */
.card-icon-dark {
  color: #fbbf24;
}

.card-icon-light {
  color: #dc2626;
}

/* 图表网格 */
.charts-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1.5rem;
}

@media (min-width: 1024px) {
  .charts-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

/* 图表卡片 */
.chart-card {
  border-radius: 1rem;
  padding: 1.5rem;
  border: 1px solid;
  transition: all 0.3s ease;
  backdrop-filter: blur(12px);
}

.chart-card.dark-theme {
  background-color: rgba(31, 41, 55, 0.7);
  border-color: rgba(55, 65, 81, 0.5);
}

.chart-card.light-theme {
  background-color: rgba(255, 255, 255, 0.7);
  border-color: rgba(229, 231, 235, 0.5);
}

/* 图表标题 */
.chart-title {
  font-size: 1.125rem;
  font-weight: 600;
  margin-bottom: 1.5rem;
}

.chart-title.dark-theme {
  color: #ffffff;
}

.chart-title.light-theme {
  color: #1f2937;
}

/* 图表容器 */
.chart-container {
  width: 100%;
  height: 300px;
}

/* 页脚 */
.stats-footer {
  position: relative;
  z-index: 10;
  backdrop-filter: blur(12px);
  border-top: 1px solid;
  margin-top: 3rem;
  transition: all 0.3s ease;
}

.footer-container {
  max-width: 80rem;
  margin: 0 auto;
  padding: 1.5rem 1rem;
}

.footer-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
}

.stats-footer.dark-theme {
  background-color: rgba(31, 41, 55, 0.7);
  border-color: rgba(55, 65, 81, 0.5);
}

.stats-footer.light-theme {
  background-color: rgba(255, 255, 255, 0.7);
  border-color: rgba(229, 231, 235, 0.5);
}

.footer-text {
  text-align: center;
  font-size: 0.875rem;
  transition: color 0.3s ease;
}

.footer-text.dark-theme {
  color: #9ca3af;
}

.footer-text.light-theme {
  color: #6b7280;
}

/* 响应式图表容器 */
@media (max-width: 768px) {
  .chart-container {
    height: 250px;
    min-height: 250px;
  }
}
</style>