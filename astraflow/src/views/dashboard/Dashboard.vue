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
              账单可视化 Dashboard
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

            <div class="flex items-center space-x-2">
              <div class="w-8 h-8 bg-gradient-to-r from-cyan-400 to-blue-500 rounded-full flex items-center justify-center">
                <UserCircle :size="20" class="text-white" />
              </div>
              <span :class="`text-sm transition-colors duration-300 ${isDark ? 'text-gray-300' : 'text-gray-700'}`">用户</span>
            </div>
          </div>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="relative z-10 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- Key Metrics Cards -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        <div
          v-for="(metric, index) in metrics"
          :key="metric.title"
          :ref="el => setCardRef(el, index)"
          :class="`rounded-xl p-6 border backdrop-blur-md transition-all duration-300 ${
            isDark
              ? 'bg-gray-800/70 border-gray-700/50 hover:shadow-cyan-500/10'
              : 'bg-white/70 border-gray-200/50 hover:shadow-cyan-500/20'
          }`"
          @mouseenter="handleCardHover(index, true)"
          @mouseleave="handleCardHover(index, false)"
        >
          <div class="flex items-center justify-between">
            <div>
              <p :class="`text-sm mb-1 transition-colors duration-300 ${isDark ? 'text-gray-400' : 'text-gray-600'}`">
                {{ metric.title }}
              </p>
              <p :class="`text-2xl font-bold transition-colors duration-300 ${isDark ? 'text-white' : 'text-gray-800'}`">
                {{ metric.value }}
              </p>
              <p class="text-xs mt-1 text-green-500">{{ metric.change }}</p>
            </div>
            <div :class="`p-3 rounded-lg bg-gradient-to-r ${metric.color} text-white`">
              <component :is="metric.icon" :size="24" />
            </div>
          </div>
        </div>
      </div>

      <!-- Charts Row -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-8">
        <!-- Expense Categories Pie Chart -->
        <div
          ref="pieChartRef"
          :class="`rounded-xl p-6 border backdrop-blur-md transition-all duration-300 ${
            isDark
              ? 'bg-gray-800/70 border-gray-700/50'
              : 'bg-white/70 border-gray-200/50'
          }`"
        >
          <h3 :class="`text-lg font-semibold mb-4 transition-colors duration-300 ${isDark ? 'text-white' : 'text-gray-800'}`">
            支出类别占比
          </h3>
          <div class="h-80">
            <PieChart>
              <Pie
                :data="expenseCategories"
                cx="50%"
                cy="50%"
                :labelLine="false"
                :outerRadius="80"
                fill="#8884d8"
                dataKey="value"
                :label="renderCustomizedLabel"
              >
                <Cell v-for="(entry, index) in expenseCategories" :key="`cell-${index}`" :fill="entry.color" />
              </Pie>
              <Tooltip
                :contentStyle="{
                  backgroundColor: isDark ? '#1f2937' : '#ffffff',
                  border: isDark ? '1px solid #374151' : '1px solid #e5e7eb',
                  borderRadius: '8px',
                  color: isDark ? '#ffffff' : '#000000'
                }"
              />
            </PieChart>
          </div>
        </div>

        <!-- Weekly Expense Trend Line Chart -->
        <div
          ref="lineChartRef"
          :class="`rounded-xl p-6 border backdrop-blur-md transition-all duration-300 ${
            isDark
              ? 'bg-gray-800/70 border-gray-700/50'
              : 'bg-white/70 border-gray-200/50'
          }`"
        >
          <h3 :class="`text-lg font-semibold mb-4 transition-colors duration-300 ${isDark ? 'text-white' : 'text-gray-800'}`">
            每周支出趋势
          </h3>
          <div class="h-80">
            <LineChart :data="weeklyExpenses">
              <CartesianGrid
                strokeDasharray="3 3"
                :stroke="isDark ? '#374151' : '#e5e7eb'"
              />
              <XAxis
                dataKey="week"
                :stroke="isDark ? '#9ca3af' : '#6b7280'"
              />
              <YAxis
                :stroke="isDark ? '#9ca3af' : '#6b7280'"
                :tickFormatter="formatYAxis"
              />
              <Tooltip
                :contentStyle="{
                  backgroundColor: isDark ? '#1f2937' : '#ffffff',
                  border: isDark ? '1px solid #374151' : '1px solid #e5e7eb',
                  borderRadius: '8px',
                  color: isDark ? '#ffffff' : '#000000'
                }"
                :formatter="formatTooltip"
              />
              <Legend />
              <Line
                type="monotone"
                dataKey="expense"
                stroke="#3B82F6"
                :strokeWidth="3"
                :activeDot="{ r: 8 }"
                name="周支出"
              />
            </LineChart>
          </div>
        </div>
      </div>

      <!-- Recent Bills and AI Insights -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
        <!-- Recent Bills Table -->
        <div
          ref="billsTableRef"
          :class="`rounded-xl p-6 border backdrop-blur-md transition-all duration-300 ${
            isDark
              ? 'bg-gray-800/70 border-gray-700/50'
              : 'bg-white/70 border-gray-200/50'
          }`"
        >
          <h3 :class="`text-lg font-semibold mb-4 transition-colors duration-300 ${isDark ? 'text-white' : 'text-gray-800'}`">
            最近账单
          </h3>
          <div class="overflow-x-auto">
            <table class="w-full text-sm">
              <thead>
                <tr :class="`border-b transition-colors duration-300 ${isDark ? 'border-gray-700' : 'border-gray-200'}`">
                  <th :class="`text-left py-3 px-4 transition-colors duration-300 ${isDark ? 'text-gray-400' : 'text-gray-600'}`">日期</th>
                  <th :class="`text-left py-3 px-4 transition-colors duration-300 ${isDark ? 'text-gray-400' : 'text-gray-600'}`">商户</th>
                  <th :class="`text-left py-3 px-4 transition-colors duration-300 ${isDark ? 'text-gray-400' : 'text-gray-600'}`">类别</th>
                  <th :class="`text-left py-3 px-4 transition-colors duration-300 ${isDark ? 'text-gray-400' : 'text-gray-600'}`">金额</th>
                  <th :class="`text-left py-3 px-4 transition-colors duration-300 ${isDark ? 'text-gray-400' : 'text-gray-600'}`">状态</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="bill in recentBills"
                  :key="bill.id"
                  :class="`border-b transition-all duration-300 ${
                    isDark
                      ? 'border-gray-800 hover:bg-gray-700/30'
                      : 'border-gray-100 hover:bg-gray-50'
                  }`"
                >
                  <td :class="`py-3 px-4 transition-colors duration-300 ${isDark ? 'text-gray-300' : 'text-gray-700'}`">
                    {{ bill.date }}
                  </td>
                  <td :class="`py-3 px-4 transition-colors duration-300 ${isDark ? 'text-gray-300' : 'text-gray-700'}`">
                    {{ bill.vendor }}
                  </td>
                  <td class="py-3 px-4">
                    <span :class="`px-2 py-1 rounded-full text-xs transition-colors duration-300 ${
                      isDark
                        ? 'bg-blue-900/50 text-blue-300'
                        : 'bg-blue-100 text-blue-800'
                    }`">
                      {{ bill.category }}
                    </span>
                  </td>
                  <td :class="`py-3 px-4 font-medium transition-colors duration-300 ${isDark ? 'text-white' : 'text-gray-800'}`">
                    ¥{{ bill.amount }}
                  </td>
                  <td class="py-3 px-4">
                    <span :class="`px-2 py-1 rounded-full text-xs transition-colors duration-300 ${
                      bill.status === '已通过'
                        ? isDark
                          ? 'bg-green-900/50 text-green-300'
                          : 'bg-green-100 text-green-800'
                        : bill.status === '待审核'
                        ? isDark
                          ? 'bg-yellow-900/50 text-yellow-300'
                          : 'bg-yellow-100 text-yellow-800'
                        : isDark
                          ? 'bg-red-900/50 text-red-300'
                          : 'bg-red-100 text-red-800'
                    }`">
                      {{ bill.status }}
                    </span>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- AI Insights -->
        <div
          ref="aiInsightsRef"
          :class="`rounded-xl p-6 border backdrop-blur-md transition-all duration-300 ${
            isDark
              ? 'bg-gray-800/70 border-gray-700/50'
              : 'bg-white/70 border-gray-200/50'
          }`"
        >
          <h3 :class="`text-lg font-semibold mb-4 transition-colors duration-300 ${isDark ? 'text-white' : 'text-gray-800'}`">
            AI 分析提示
          </h3>
          <div class="space-y-4">
            <div
              v-for="insight in aiInsights"
              :key="insight.id"
              :class="`p-4 rounded-lg border-l-4 transition-colors duration-300 ${
                insight.type === 'warning'
                  ? isDark
                    ? 'bg-yellow-900/20 border-yellow-600 text-yellow-300'
                    : 'bg-yellow-50 border-yellow-500 text-yellow-800'
                  : insight.type === 'info'
                  ? isDark
                    ? 'bg-blue-900/20 border-blue-600 text-blue-300'
                    : 'bg-blue-50 border-blue-500 text-blue-800'
                  : isDark
                    ? 'bg-red-900/20 border-red-600 text-red-300'
                    : 'bg-red-50 border-red-500 text-red-800'
              }`"
            >
              <p class="text-sm">{{ insight.message }}</p>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Footer -->
    <footer :class="`relative z-10 backdrop-blur-md border-t mt-12 transition-all duration-300 ${
      isDark
        ? 'bg-gray-800/70 border-gray-700/50'
        : 'bg-white/70 border-gray-200/50'
    }`">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
        <!-- 返回主页按钮 -->
        <div class="flex flex-col items-center space-y-4">
          <router-link
            to="/"
            :class="`inline-flex items-center space-x-2 px-6 py-3 rounded-lg font-medium transition-all duration-300 transform hover:scale-105 ${
              isDark
                ? 'bg-gradient-to-r from-cyan-600 to-blue-600 hover:from-cyan-500 hover:to-blue-500 text-white shadow-lg shadow-cyan-500/25'
                : 'bg-gradient-to-r from-cyan-500 to-blue-500 hover:from-cyan-600 hover:to-blue-600 text-white shadow-lg shadow-cyan-500/30'
            }`"
          >
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="m15 18-6-6 6-6"/>
            </svg>
            <span>返回主页</span>
          </router-link>

          <p :class="`text-center text-sm transition-colors duration-300 ${isDark ? 'text-gray-400' : 'text-gray-600'}`">
            © 2025 AstraFlow · Smart Expense Made Simple
          </p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { PieChart, Pie, Cell, LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts'
import {
  CreditCard,
  DollarSign,
  FileText,
  CheckCircle,
  SunIcon,
  MoonIcon,
  UserCircle
} from 'lucide-vue-next'
import { useTheme } from '../../composables/useTheme'

const { theme, toggleTheme, isDark } = useTheme()

// Refs for animation elements
const cardRefs = ref([])
const pieChartRef = ref(null)
const lineChartRef = ref(null)
const billsTableRef = ref(null)
const aiInsightsRef = ref(null)

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

  // Animate charts
  const charts = [pieChartRef.value, lineChartRef.value]
  charts.forEach((chart, index) => {
    if (chart) {
      chart.style.opacity = '0'
      chart.style.transform = index === 0 ? 'translateX(-20px)' : 'translateX(20px)'

      setTimeout(() => {
        chart.style.transition = 'opacity 0.5s ease, transform 0.5s ease'
        chart.style.opacity = '1'
        chart.style.transform = 'translateX(0)'
      }, 300)
    }
  })

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

onMounted(() => {
  // Animate elements after a short delay
  setTimeout(animateElements, 100)
})

onUnmounted(() => {
  // Clean up any animation timeouts if needed
})

// Custom label for pie chart
const renderCustomizedLabel = ({ name, percent }) => {
  return `${name} ${(percent * 100).toFixed(0)}%`
}

// Format Y axis for line chart
const formatYAxis = (value) => {
  return `¥${value}`
}

// Format tooltip for line chart
const formatTooltip = (value) => {
  return [`¥${value}`, '支出']
}

// Mock data for the charts
const expenseCategories = ref([
  { name: '餐饮', value: 4500, color: '#3B82F6' },
  { name: '交通', value: 3200, color: '#10B981' },
  { name: '办公', value: 2800, color: '#8B5CF6' },
  { name: '住宿', value: 6500, color: '#EF4444' },
  { name: '其他', value: 1200, color: '#F59E0B' }
])

const weeklyExpenses = ref([
  { week: '第1周', expense: 3200 },
  { week: '第2周', expense: 4500 },
  { week: '第3周', expense: 2800 },
  { week: '第4周', expense: 5800 }
])

const recentBills = ref([
  { id: 1, date: '2025-11-01', category: '餐饮', amount: 128.50, status: '已通过', vendor: '星巴克' },
  { id: 2, date: '2025-10-31', category: '交通', amount: 45.00, status: '待审核', vendor: '滴滴出行' },
  { id: 3, date: '2025-10-30', category: '办公', amount: 299.90, status: '已通过', vendor: '京东' },
  { id: 4, date: '2025-10-29', category: '餐饮', amount: 89.60, status: '已拒绝', vendor: '麦当劳' },
  { id: 5, date: '2025-10-28', category: '住宿', amount: 450.00, status: '已通过', vendor: '汉庭酒店' }
])

const aiInsights = ref([
  { id: 1, message: '系统检测到本月餐饮支出较上月增长25%，建议控制预算', type: 'warning' },
  { id: 2, message: 'AI识别到3笔相似金额的报销，可能存在重复提交', type: 'info' },
  { id: 3, message: '交通费用报销已达到月度上限，剩余预算不足', type: 'alert' }
])

// Mock metrics data
const metrics = ref([
  {
    title: '本月总支出',
    value: '¥12,380',
    icon: DollarSign,
    color: 'from-blue-500 to-cyan-500',
    change: '+12.5%'
  },
  {
    title: '自动报销笔数',
    value: '48笔',
    icon: FileText,
    color: 'from-green-500 to-emerald-500',
    change: '+8.2%'
  },
  {
    title: 'AI识别准确率',
    value: '98.7%',
    icon: CheckCircle,
    color: 'from-purple-500 to-pink-500',
    change: '+1.2%'
  },
  {
    title: '待审核报销',
    value: '5条',
    icon: CreditCard,
    color: 'from-amber-500 to-orange-500',
    change: '-2.1%'
  }
])
</script>

<style scoped>
/* 暗色模式下的粒子背景优化 */
.dark .animate-pulse {
  opacity: 0.3;
}

/* 确保所有过渡效果流畅 */
.transition-all {
  transition-property: all;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
}

/* 暗色模式下的卡片阴影增强 */
.dark .hover\:shadow-cyan-500\/10:hover {
  box-shadow: 0 20px 25px -5px rgba(56, 189, 248, 0.15), 0 10px 10px -5px rgba(56, 189, 248, 0.08);
}
</style>