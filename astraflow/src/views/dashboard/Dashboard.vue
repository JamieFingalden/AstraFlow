<template>
  <div class="dashboard-container" :data-theme="isDark ? 'dark' : 'light'">
    <!-- Particle background effect -->
    <div class="particle-background">
      <div class="particle particle-cyan"></div>
      <div class="particle particle-blue"></div>
      <div class="particle particle-purple"></div>
    </div>

    <!-- Header -->
    <header class="dashboard-header">
      <div class="header-container">
        <div class="header-left">
          <!-- 返回主页按钮 -->
          <router-link to="/" class="back-button" title="返回主页">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="m15 18-6-6 6-6"/>
            </svg>
            <span class="back-text">返回</span>
          </router-link>

          <div class="brand-name">
            AstraFlow
          </div>
          <h1 class="page-title">账单可视化 Dashboard</h1>
        </div>

        <div class="header-right">
          <button @click="() => { toggleTheme(); handleThemeChange(); }" class="theme-toggle">
            <SunIcon v-if="isDark" :size="20" />
            <MoonIcon v-else :size="20" />
          </button>

          <div class="user-info">
            <div class="user-avatar">
              <UserCircle :size="20" class="user-icon" />
            </div>
            <span class="user-name">{{ userStore.user.name || '用户' }}</span>
          </div>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="main-content">
      <!-- Key Metrics Cards -->
      <div class="metrics-grid">
        <div
          v-for="(metric, index) in metrics"
          :key="metric.title"
          ref="el => setCardRef(el, index)"
          class="metric-card"
          @mouseenter="handleCardHover(index, true)"
          @mouseleave="handleCardHover(index, false)"
        >
          <div class="metric-content">
            <div class="metric-info">
              <p class="metric-title">
                {{ metric.title }}
              </p>
              <p class="metric-value">
                {{ metric.value }}
              </p>
              <p class="metric-change">{{ metric.change }}</p>
            </div>
            <div class="metric-icon" :class="metric.color.replace('from-', 'bg-').replace(' to-', '-to-')">
              <component :is="metric.icon" :size="24" />
            </div>
          </div>
        </div>
      </div>

      <!-- Charts Row -->
      <div class="charts-row">
        <!-- Expense Categories Pie Chart -->
        <div class="chart-card">
          <h3 class="chart-title">
            支出类别占比
          </h3>
          <div class="chart-container">
            <div ref="pieChartRef" class="echarts-container" v-if="expenseCategories && expenseCategories.length > 0"></div>
            <div v-else class="chart-placeholder">
              暂无支出类别数据
            </div>
          </div>
        </div>

        <!-- Monthly Expense Trend Line Chart -->
        <div class="chart-card">
          <h3 class="chart-title">
            每月支出趋势
          </h3>
          <div class="chart-container">
            <div ref="lineChartRef" class="echarts-container" v-if="monthlyExpenses && monthlyExpenses.length > 0"></div>
            <div v-else class="chart-placeholder">
              暂无每月支出数据
            </div>
          </div>
        </div>
      </div>

      <!-- Recent Bills and AI Insights -->
      <div class="content-row">
        <!-- Recent Bills Table -->
        <div ref="billsTableRef" class="bills-card">
          <h3 class="card-title">
            最近账单
          </h3>
          <div class="bills-table-container">
            <table class="bills-table">
              <thead>
                <tr class="table-header-row">
                  <th class="table-header-cell">日期</th>
                  <th class="table-header-cell">商户</th>
                  <th class="table-header-cell">类别</th>
                  <th class="table-header-cell">金额</th>
                  <th class="table-header-cell">状态</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="bill in recentBills"
                  :key="bill.id"
                  class="table-row"
                >
                  <td class="table-cell">
                    {{ bill.date }}
                  </td>
                  <td class="table-cell">
                    {{ bill.vendor }}
                  </td>
                  <td class="table-cell">
                    <span class="status-badge status-badge-blue">
                      {{ bill.category }}
                    </span>
                  </td>
                  <td class="table-cell table-cell-bold">
                    ¥{{ bill.amount }}
                  </td>
                  <td class="table-cell">
                    <span class="status-badge" :class="getStatusBadgeClass(bill.status)">
                      {{ bill.status }}
                    </span>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- AI Insights -->
        <div ref="aiInsightsRef" class="insights-card">
          <h3 class="card-title">
            AI 分析提示
          </h3>
          <div class="insights-container">
            <div
              v-for="insight in aiInsights"
              :key="insight.id"
              class="insight-item"
              :class="getInsightClass(insight.type)"
            >
              <p class="insight-text">{{ insight.message }}</p>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Footer -->
    <footer class="dashboard-footer">
      <div class="footer-container">
        <div class="footer-content">
          <p class="footer-text">
            © 2025 AstraFlow · Smart Expense Made Simple
          </p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import * as echarts from 'echarts'
import {
  CreditCard,
  DollarSign,
  FileText,
  CheckCircle,
  SunIcon,
  MoonIcon,
  UserCircle,
  LogOut
} from 'lucide-vue-next'
import { useTheme } from '../../composables/useTheme'
import { useUserStore } from '../../stores/userStore'
import { getDashboardData } from '../../services/api/analyticsApi'

const router = useRouter()
const { theme, toggleTheme, isDark } = useTheme()
const userStore = useUserStore()

// Refs for animation elements
const cardRefs = ref([])
const pieChartRef = ref(null)
const lineChartRef = ref(null)
const billsTableRef = ref(null)
const aiInsightsRef = ref(null)

// State for loading and error
const loading = ref(false)
const error = ref(null)

// Set card ref function
const setCardRef = (el, index) => {
  if (el) {
    cardRefs.value[index] = el
  }
}

// Handle card hover for animation
const handleCardHover = (index, isHover) => {
  const card = cardRefs.value[index]
  if (card) {
    if (isHover) {
      card.style.transform = 'translateY(-5px) scale(1.02)'
      card.style.boxShadow = isDark.value
        ? '0 20px 25px -5px rgba(56, 189, 248, 0.15), 0 10px 10px -5px rgba(56, 189, 248, 0.08)'
        : '0 20px 25px -5px rgba(56, 189, 248, 0.25), 0 10px 10px -5px rgba(56, 189, 248, 0.15)'
    } else {
      card.style.transform = 'translateY(0) scale(1)'
      card.style.boxShadow = isDark.value
        ? '0 10px 15px -3px rgba(0, 0, 0, 0.3), 0 4px 6px -2px rgba(0, 0, 0, 0.2)'
        : '0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)'
    }
  }
}

// Update charts when theme changes
const handleThemeChange = () => {
  nextTick(() => {
    updateChartTheme()
  })
}

// Animate elements on mount
const animateElements = () => {
  // Animate key metrics cards
  cardRefs.value.forEach((card, index) => {
    if (card) {
      card.style.opacity = '0'
      card.style.transform = 'translateY(20px)'

      setTimeout(() => {
        card.style.transition = 'opacity 0.5s ease, transform 0.5s ease'
        card.style.opacity = '1'
        card.style.transform = 'translateY(0)'
      }, index * 100)
    }
  })

  // Initialize charts after a short delay to ensure DOM is ready
  setTimeout(() => {
    initCharts()
  }, 100)

  // Animate tables and insights
  const otherElements = [billsTableRef.value, aiInsightsRef.value]
  otherElements.forEach((element, index) => {
    if (element) {
      element.style.opacity = '0'
      element.style.transform = 'translateY(20px)'

      setTimeout(() => {
        element.style.transition = 'opacity 0.5s ease, transform 0.5s ease'
        element.style.opacity = '1'
        element.style.transform = 'translateY(0)'
      }, 500 + index * 100)
    }
  })
}

// Fetch dashboard data from API
const fetchDashboardData = async () => {
  loading.value = true
  error.value = null

  try {
    const response = await getDashboardData()
    console.log('Dashboard API Response:', response) // Debug log

    if (response.code === 200 && response.data && response.data.dashboard) {
      const dashboardData = response.data.dashboard
      console.log('Dashboard data processed:', dashboardData) // Debug log

      // Update metrics
      metrics.value = [
        {
          title: '本月总支出',
          value: `¥${dashboardData.metrics.monthly_expense?.toLocaleString() || '0'}`,
          icon: DollarSign,
          color: 'from-blue-500 to-cyan-500',
          change: '+0.0%' // Placeholder - will calculate from API when available
        },
        {
          title: '自动报销笔数',
          value: `${dashboardData.metrics.auto_processed || 0}笔`,
          icon: FileText,
          color: 'from-green-500 to-emerald-500',
          change: '+0.0%' // Placeholder - will calculate from API when available
        },
        {
          title: 'AI识别准确率',
          value: `${dashboardData.metrics.ai_accuracy?.toFixed(1) || '0.0'}%`,
          icon: CheckCircle,
          color: 'from-purple-500 to-pink-500',
          change: '+0.0%' // Placeholder - will calculate from API when available
        },
        {
          title: '待审核报销',
          value: `${dashboardData.metrics.pending_reviews || 0}条`,
          icon: CreditCard,
          color: 'from-amber-500 to-orange-500',
          change: '+0.0%' // Placeholder - will calculate from API when available
        }
      ]
      console.log('Metrics updated:', metrics.value) // Debug log

      // Update expense categories
      expenseCategories.value = dashboardData.expense_categories || []
      console.log('Expense categories updated:', expenseCategories.value) // Debug log

      // Update monthly expenses
      monthlyExpenses.value = dashboardData.monthly_expenses || []
      console.log('Monthly expenses updated:', monthlyExpenses.value) // Debug log

      // Update recent bills and format dates for display
      recentBills.value = (dashboardData.recent_bills || []).map(bill => ({
        ...bill,
        date: new Date(bill.date).toLocaleDateString('zh-CN')
      }))
      console.log('Recent bills updated:', recentBills.value) // Debug log

      // Update AI insights (keeping empty as requested by user)
      // aiInsights.value = dashboardData.ai_insights || []
      // For now, we'll use empty array as AI insights will be implemented later
      aiInsights.value = []
      console.log('AI Insights updated:', aiInsights.value) // Debug log
    } else {
      error.value = response.message || '获取仪表板数据失败'
      console.error('Dashboard API response error:', response) // Debug log
    }
  } catch (err) {
    console.error('Error fetching dashboard data:', err)
    error.value = '获取数据时发生错误，请稍后重试'
  } finally {
    loading.value = false
    // Force Vue to re-render the charts by making the data arrays reactive again
    if (metrics.value && metrics.value.length > 0) {
      metrics.value = [...metrics.value]
    }
    if (expenseCategories.value && expenseCategories.value.length > 0) {
      expenseCategories.value = [...expenseCategories.value]
    }
    if (monthlyExpenses.value && monthlyExpenses.value.length > 0) {
      monthlyExpenses.value = [...monthlyExpenses.value]
    }
    if (recentBills.value && recentBills.value.length > 0) {
      recentBills.value = [...recentBills.value]
    }
    if (aiInsights.value && aiInsights.value.length > 0) {
      aiInsights.value = [...aiInsights.value]
    }

    // Trigger chart re-rendering after a small delay
    setTimeout(() => {
      // Force a Vue update by touching the reactive variables again
      expenseCategories.value = [...expenseCategories.value]
      monthlyExpenses.value = [...monthlyExpenses.value]
    }, 50)

    // Animate elements after data is loaded
    setTimeout(animateElements, 100)
  }
}

onMounted(() => {
  fetchDashboardData()
})

onUnmounted(() => {
  // Clean up any animation timeouts if needed
  disposeCharts()
})


// Initialize with empty arrays/objects to be populated by API
const expenseCategories = ref([])
const monthlyExpenses = ref([])
const recentBills = ref([])
const aiInsights = ref([])
const metrics = ref([])

// Initialize and update charts
const initCharts = () => {
  // Initialize pie chart
  if (expenseCategories.value && expenseCategories.value.length > 0) {
    const pieChartDom = pieChartRef.value
    if (pieChartDom) {
      const pieChart = echarts.getInstanceByDom(pieChartDom) || echarts.init(pieChartDom, isDark.value ? 'dark' : 'light')

      const pieOption = {
        tooltip: {
          trigger: 'item',
          formatter: '{a} <br/>{b}: ¥{c} ({d}%)'
        },
        legend: {
          orient: 'vertical',
          left: 'left',
          textStyle: {
            color: isDark.value ? '#fff' : '#000'
          }
        },
        series: [
          {
            name: '支出类别',
            type: 'pie',
            radius: '50%',
            data: expenseCategories.value.map(item => ({
              value: item.value,
              name: item.name,
              itemStyle: { color: item.color }
            })),
            emphasis: {
              itemStyle: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: 'rgba(0, 0, 0, 0.5)'
              }
            }
          }
        ]
      }

      pieChart.setOption(pieOption, true)
    }
  }

  // Initialize line chart
  if (monthlyExpenses.value && monthlyExpenses.value.length > 0) {
    const lineChartDom = lineChartRef.value
    if (lineChartDom) {
      const lineChart = echarts.getInstanceByDom(lineChartDom) || echarts.init(lineChartDom, isDark.value ? 'dark' : 'light')

      const lineOption = {
        tooltip: {
          trigger: 'axis',
          formatter: (params) => {
            const param = params[0]
            return `${param.name}<br/>${param.seriesName}: ¥${param.value}`
          }
        },
        legend: {
          data: ['月支出'],
          textStyle: {
            color: isDark.value ? '#fff' : '#000'
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
          data: monthlyExpenses.value.map(item => item.month),
          axisLabel: {
            color: isDark.value ? '#fff' : '#000'
          }
        },
        yAxis: {
          type: 'value',
          axisLabel: {
            color: isDark.value ? '#fff' : '#000',
            formatter: '¥{value}'
          }
        },
        series: [
          {
            name: '月支出',
            type: 'line',
            smooth: true, // This makes it a curved line chart
            data: monthlyExpenses.value.map(item => item.expense),
            itemStyle: {
              color: '#3B82F6'
            },
            lineStyle: {
              color: '#3B82F6'
            }
          }
        ]
      }

      lineChart.setOption(lineOption, true)
    }
  }
}

// Dispose charts when component unmounts
const disposeCharts = () => {
  const pieChartDom = pieChartRef.value
  const lineChartDom = lineChartRef.value

  if (pieChartDom) {
    const pieChart = echarts.getInstanceByDom(pieChartDom)
    if (pieChart) {
      pieChart.dispose()
    }
  }

  if (lineChartDom) {
    const lineChart = echarts.getInstanceByDom(lineChartDom)
    if (lineChart) {
      lineChart.dispose()
    }
  }
}

// Update chart when theme changes
const updateChartTheme = () => {
  const pieChartDom = pieChartRef.value
  const lineChartDom = lineChartRef.value

  if (pieChartDom) {
    const pieChart = echarts.getInstanceByDom(pieChartDom)
    if (pieChart) {
      pieChart.dispose()
      initCharts()
    }
  }

  if (lineChartDom) {
    const lineChart = echarts.getInstanceByDom(lineChartDom)
    if (lineChart) {
      lineChart.dispose()
      initCharts()
    }
  }
}

// Helper methods for dynamic classes
const getStatusBadgeClass = (status) => {
  if (status === '已通过') return 'status-badge-green'
  if (status === '待审核') return 'status-badge-yellow'
  return 'status-badge-red'
}

const getInsightClass = (type) => {
  if (type === 'warning') return 'insight-warning'
  if (type === 'info') return 'insight-info'
  return 'insight-alert'
}
</script>

<style scoped>
/* Main Container */
.dashboard-container {
  min-height: 100vh;
  transition: color 0.3s ease, background-color 0.3s ease;
}

.dashboard-container[data-theme="dark"] {
  background-color: #111827;
}

.dashboard-container[data-theme="light"] {
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

.dashboard-container[data-theme="dark"] .particle-cyan {
  background-color: rgba(34, 211, 238, 0.3);
}

.dashboard-container[data-theme="light"] .particle-cyan {
  background-color: #06b6d4;
}

.particle-blue {
  bottom: -10rem;
  left: -10rem;
}

.dashboard-container[data-theme="dark"] .particle-blue {
  background-color: rgba(59, 130, 246, 0.3);
}

.dashboard-container[data-theme="light"] .particle-blue {
  background-color: #3b82f6;
}

.particle-purple {
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.dashboard-container[data-theme="dark"] .particle-purple {
  background-color: rgba(147, 51, 234, 0.3);
}

.dashboard-container[data-theme="light"] .particle-purple {
  background-color: #a855f7;
}

/* Header */
.dashboard-header {
  position: relative;
  z-index: 10;
  backdrop-filter: blur(12px);
  border-bottom: 1px solid;
  transition: all 0.3s ease;
}

.dashboard-container[data-theme="dark"] .dashboard-header {
  background-color: rgba(31, 41, 55, 0.7);
  border-color: rgba(55, 65, 81, 0.5);
}

.dashboard-container[data-theme="light"] .dashboard-header {
  background-color: rgba(255, 255, 255, 0.7);
  border-color: rgba(229, 231, 235, 0.5);
}

.header-container {
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
  .header-container {
    padding: 0 1.5rem;
  }
}

@media (min-width: 1024px) {
  .header-container {
    padding: 0 2rem;
  }
}

.header-left {
  display: flex;
  align-items: center;
  gap: 1rem;
  flex: 15;
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
  color: inherit;
}

.dashboard-container[data-theme="dark"] .back-button {
  color: #d1d5db;
}

.dashboard-container[data-theme="dark"] .back-button:hover {
  background-color: rgba(55, 65, 81, 0.5);
  color: white;
}

.dashboard-container[data-theme="light"] .back-button {
  color: #4b5563;
}

.dashboard-container[data-theme="light"] .back-button:hover {
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
  line-height: 1.2;
}

.dashboard-container[data-theme="dark"] .page-title {
  color: white;
}

.dashboard-container[data-theme="light"] .page-title {
  color: #1f2937;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 1rem;
  flex: 2; /* Don't grow */
}

.theme-toggle {
  padding: 0.5rem;
  border-radius: 0.5rem;
  border: none;
  transition: all 0.2s ease;
  cursor: pointer;
}

.dashboard-container[data-theme="dark"] .theme-toggle {
  background-color: rgba(55, 65, 81, 0.5);
  color: #d1d5db;
}

.dashboard-container[data-theme="dark"] .theme-toggle:hover {
  background-color: rgba(75, 85, 99, 0.5);
}

.dashboard-container[data-theme="light"] .theme-toggle {
  background-color: #f3f4f6;
  color: #374151;
}

.dashboard-container[data-theme="light"] .theme-toggle:hover {
  background-color: #e5e7eb;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.dashboard-container[data-theme="dark"] .logout-btn {
  background-color: rgba(55, 65, 81, 0.5);
  color: #d1d5db;
}

.dashboard-container[data-theme="dark"] .logout-btn:hover {
  background-color: rgba(75, 85, 99, 0.5);
}

.dashboard-container[data-theme="light"] .logout-btn {
  background-color: #f3f4f6;
  color: #374151;
}

.dashboard-container[data-theme="light"] .logout-btn:hover {
  background-color: #e5e7eb;
}

.user-avatar {
  width: 2rem;
  height: 2rem;
  background: linear-gradient(135deg, #22d3ee, #3b82f6);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.user-icon {
  color: white;
}

.user-name {
  font-size: 0.875rem;
  transition: color 0.3s ease;
}

.dashboard-container[data-theme="dark"] .user-name {
  color: #d1d5db;
}

.dashboard-container[data-theme="light"] .user-name {
  color: #374151;
}

/* Main Content */
.main-content {
  position: relative;
  z-index: 10;
  max-width: 80rem;
  margin: 0 auto;
  padding: 2rem 1rem;
}

@media (min-width: 640px) {
  .main-content {
    padding: 2rem 1.5rem;
  }
}

@media (min-width: 1024px) {
  .main-content {
    padding: 2rem 2rem;
  }
}

/* Metrics Grid */
.metrics-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1.5rem;
  margin-bottom: 2rem;
}

@media (min-width: 768px) {
  .metrics-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (min-width: 1024px) {
  .metrics-grid {
    grid-template-columns: repeat(4, 1fr);
  }
}

.metric-card {
  border-radius: 0.75rem;
  padding: 1.5rem;
  border: 1px solid;
  backdrop-filter: blur(12px);
  transition: all 0.3s ease;
}

.dashboard-container[data-theme="dark"] .metric-card {
  background-color: rgba(31, 41, 55, 0.7);
  border-color: rgba(55, 65, 81, 0.5);
}

.dashboard-container[data-theme="light"] .metric-card {
  background-color: rgba(255, 255, 255, 0.7);
  border-color: rgba(229, 231, 235, 0.5);
}

.metric-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.metric-info {
  flex: 1;
}

.metric-title {
  font-size: 0.875rem;
  margin-bottom: 0.25rem;
  transition: color 0.3s ease;
}

.dashboard-container[data-theme="dark"] .metric-title {
  color: #9ca3af;
}

.dashboard-container[data-theme="light"] .metric-title {
  color: #4b5363;
}

.metric-value {
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 0.25rem;
  transition: color 0.3s ease;
}

.dashboard-container[data-theme="dark"] .metric-value {
  color: white;
}

.dashboard-container[data-theme="light"] .metric-value {
  color: #1f2937;
}

.metric-change {
  font-size: 0.75rem;
  color: #10b981;
}

.metric-icon {
  padding: 0.75rem;
  border-radius: 0.5rem;
  color: white;
}

.bg-blue-500-to-cyan-500 {
  background: linear-gradient(135deg, #3b82f6, #06b6d4);
}

.bg-green-500-to-emerald-500 {
  background: linear-gradient(135deg, #10b981, #047857);
}

.bg-purple-500-to-pink-500 {
  background: linear-gradient(135deg, #8b5cf6, #ec4899);
}

.bg-amber-500-to-orange-500 {
  background: linear-gradient(135deg, #f59e0b, #f97316);
}

/* Charts Row */
.charts-row {
  display: grid;
  grid-template-columns: 1fr;
  gap: 2rem;
  margin-bottom: 2rem;
}

@media (min-width: 1024px) {
  .charts-row {
    grid-template-columns: repeat(2, 1fr);
  }
}

.chart-card {
  border-radius: 0.75rem;
  padding: 1.5rem;
  border: 1px solid;
  backdrop-filter: blur(12px);
  transition: all 0.3s ease;
}

.dashboard-container[data-theme="dark"] .chart-card {
  background-color: rgba(31, 41, 55, 0.7);
  border-color: rgba(55, 65, 81, 0.5);
}

.dashboard-container[data-theme="light"] .chart-card {
  background-color: rgba(255, 255, 255, 0.7);
  border-color: rgba(229, 231, 235, 0.5);
}

.chart-title {
  font-size: 1.125rem;
  font-weight: 600;
  margin-bottom: 1rem;
  transition: color 0.3s ease;
}

.dashboard-container[data-theme="dark"] .chart-title {
  color: white;
}

.dashboard-container[data-theme="light"] .chart-title {
  color: #1f2937;
}

.chart-container {
  height: 20rem;
}

.chart-placeholder {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #9ca3af;
  font-size: 0.875rem;
}

.echarts-container {
  width: 100%;
  height: 100%;
}

/* Content Row */
.content-row {
  display: grid;
  grid-template-columns: 1fr;
  gap: 2rem;
}

@media (min-width: 1024px) {
  .content-row {
    grid-template-columns: repeat(2, 1fr);
  }
}

.bills-card,
.insights-card {
  border-radius: 0.75rem;
  padding: 1.5rem;
  border: 1px solid;
  backdrop-filter: blur(12px);
  transition: all 0.3s ease;
}

.dashboard-container[data-theme="dark"] .bills-card,
.dashboard-container[data-theme="dark"] .insights-card {
  background-color: rgba(31, 41, 55, 0.7);
  border-color: rgba(55, 65, 81, 0.5);
}

.dashboard-container[data-theme="light"] .bills-card,
.dashboard-container[data-theme="light"] .insights-card {
  background-color: rgba(255, 255, 255, 0.7);
  border-color: rgba(229, 231, 235, 0.5);
}

.card-title {
  font-size: 1.125rem;
  font-weight: 600;
  margin-bottom: 1rem;
  transition: color 0.3s ease;
}

.dashboard-container[data-theme="dark"] .card-title {
  color: white;
}

.dashboard-container[data-theme="light"] .card-title {
  color: #1f2937;
}

.bills-table-container {
  overflow-x: auto;
}

.bills-table {
  width: 100%;
  font-size: 0.875rem;
}

.table-header-row {
  border-bottom: 1px solid;
  transition: border-color 0.3s ease;
}

.dashboard-container[data-theme="dark"] .table-header-row {
  border-color: #374151;
}

.dashboard-container[data-theme="light"] .table-header-row {
  border-color: #e5e7eb;
}

.table-header-cell {
  text-align: left;
  padding: 0.75rem 1rem;
  font-weight: 500;
  transition: color 0.3s ease;
}

.dashboard-container[data-theme="dark"] .table-header-cell {
  color: #9ca3af;
}

.dashboard-container[data-theme="light"] .table-header-cell {
  color: #4b5563;
}

.table-row {
  border-bottom: 1px solid;
  transition: all 0.3s ease;
}

.dashboard-container[data-theme="dark"] .table-row {
  border-color: #1f2937;
}

.dashboard-container[data-theme="light"] .table-row {
  border-color: #f3f4f6;
}

.dashboard-container[data-theme="dark"] .table-row:hover {
  background-color: rgba(55, 65, 81, 0.3);
}

.dashboard-container[data-theme="light"] .table-row:hover {
  background-color: rgba(249, 250, 251, 0.5);
}

.table-cell {
  padding: 0.75rem 1rem;
  transition: color 0.3s ease;
}

.dashboard-container[data-theme="dark"] .table-cell {
  color: #d1d5db;
}

.dashboard-container[data-theme="light"] .table-cell {
  color: #374151;
}

.table-cell-bold {
  font-weight: 500;
  transition: color 0.3s ease;
}

.dashboard-container[data-theme="dark"] .table-cell-bold {
  color: white;
}

.dashboard-container[data-theme="light"] .table-cell-bold {
  color: #111827;
}

.status-badge {
  padding: 0.25rem 0.5rem;
  border-radius: 9999px;
  font-size: 0.75rem;
  font-weight: 500;
}

.status-badge-blue {
  background-color: #dbeafe;
  color: #1d4ed8;
}

.dashboard-container[data-theme="dark"] .status-badge-blue {
  background-color: #1e3a8a;
  color: #bfdbfe;
}

.status-badge-green {
  background-color: #dcfce7;
  color: #166534;
}

.dashboard-container[data-theme="dark"] .status-badge-green {
  background-color: #14532d;
  color: #bbf7d0;
}

.status-badge-yellow {
  background-color: #fef3c7;
  color: #a16207;
}

.dashboard-container[data-theme="dark"] .status-badge-yellow {
  background-color: #7c2d12;
  color: #fde68a;
}

.status-badge-red {
  background-color: #fee2e2;
  color: #b91c1c;
}

.dashboard-container[data-theme="dark"] .status-badge-red {
  background-color: #7f1d1d;
  color: #fecaca;
}

.insights-container {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.insight-item {
  padding: 1rem;
  border-radius: 0.5rem;
  border-left-width: 4px;
}

.insight-warning {
  background-color: #fef3c7;
  border-left-color: #f59e0b;
  color: #a16207;
}

.dashboard-container[data-theme="dark"] .insight-warning {
  background-color: #7c2d12;
  color: #fde68a;
}

.insight-info {
  background-color: #dbeafe;
  border-left-color: #3b82f6;
  color: #1d4ed8;
}

.dashboard-container[data-theme="dark"] .insight-info {
  background-color: #1e3a8a;
  color: #bfdbfe;
}

.insight-alert {
  background-color: #fee2e2;
  border-left-color: #ef4444;
  color: #b91c1c;
}

.dashboard-container[data-theme="dark"] .insight-alert {
  background-color: #7f1d1d;
  color: #fecaca;
}

/* Footer */
.dashboard-footer {
  position: relative;
  z-index: 10;
  backdrop-filter: blur(12px);
  border-top: 1px solid;
  margin-top: 1rem;
  transition: all 0.3s ease;
}

.dashboard-container[data-theme="dark"] .dashboard-footer {
  background-color: rgba(31, 41, 55, 0.7);
  border-color: rgba(55, 65, 81, 0.5);
}

.dashboard-container[data-theme="light"] .dashboard-footer {
  background-color: rgba(255, 255, 255, 0.7);
  border-color: rgba(229, 231, 235, 0.5);
}

.footer-container {
  max-width: 80rem;
  margin: 0 auto;
  padding: 1.5rem 1rem;
}

@media (min-width: 640px) {
  .footer-container {
    padding: 1.5rem 1.5rem;
  }
}

@media (min-width: 1024px) {
  .footer-container {
    padding: 1.5rem 2rem;
  }
}

.footer-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
}

.home-button {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1.5rem;
  border-radius: 0.5rem;
  font-weight: 500;
  text-decoration: none;
  transition: all 0.3s ease;
  transform: scale(1);
  background: linear-gradient(135deg, #06b6d4, #3b82f6);
  color: white;
}

.home-button:hover {
  transform: scale(1.05);
  background: linear-gradient(135deg, #0891b2, #2563eb);
}

.dashboard-container[data-theme="dark"] .home-button {
  box-shadow: 0 10px 15px -3px rgba(6, 182, 212, 0.25);
}

.dashboard-container[data-theme="light"] .home-button {
  box-shadow: 0 10px 15px -3px rgba(6, 182, 212, 0.3);
}

.dashboard-container[data-theme="dark"] .home-button:hover {
  box-shadow: 0 20px 25px -5px rgba(6, 182, 212, 0.3);
}

.dashboard-container[data-theme="light"] .home-button:hover {
  box-shadow: 0 20px 25px -5px rgba(6, 182, 212, 0.4);
}

.footer-text {
  text-align: center;
  font-size: 0.875rem;
  transition: color 0.3s ease;
}

.dashboard-container[data-theme="dark"] .footer-text {
  color: #9ca3af;
}

.dashboard-container[data-theme="light"] .footer-text {
  color: #6b7280;
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
</style>