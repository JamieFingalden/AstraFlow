<template>
  <div :class="'min-h-screen transition-colors duration-300 ' + (isDark ? 'bg-gray-900' : 'bg-gradient-to-br from-blue-50 via-cyan-50 to-indigo-50')">
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
              账单管理
            </h1>
          </div>

          <div class="flex items-center space-x-4">
            <!-- 添加账单按钮 -->
            <button
              :class="'px-6 py-2 rounded-lg font-semibold transition-all duration-300 transform hover:scale-105 ' + (isDark ? 'bg-gradient-to-r from-blue-600 to-cyan-600 hover:from-blue-500 hover:to-cyan-500 text-white shadow-lg shadow-blue-500/25' : 'bg-gradient-to-r from-blue-500 to-cyan-500 hover:from-blue-600 hover:to-cyan-600 text-white shadow-lg shadow-blue-500/30')"
            >
              <PlusIcon :size="20" class="inline mr-2" />
              添加账单
            </button>

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
    <main class="relative z-10 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <!-- Search and Filter Section -->
      <div :class="`rounded-2xl p-6 mb-6 backdrop-blur-md border transition-all duration-300 ${
        isDark
          ? 'bg-gray-800/70 border-gray-700/50'
          : 'bg-white/70 border-gray-200/50'
      }`">
        <div class="flex flex-col md:flex-row gap-4 items-center justify-between">
          <!-- Search Bar -->
          <div class="flex-1 w-full md:w-auto">
            <div class="relative">
              <SearchIcon
                :class="'absolute left-3 top-1/2 transform -translate-y-1/2 ' + (isDark ? 'text-gray-400' : 'text-gray-500')"
                :size="20"
              />
              <input
                v-model="searchQuery"
                type="text"
                placeholder="搜索账单名称..."
                :class="'w-full pl-10 pr-4 py-3 rounded-lg border transition-all duration-300 focus:outline-none focus:ring-2 focus:ring-blue-500/50 ' + (isDark ? 'bg-gray-700/50 border-gray-600 text-white placeholder-gray-400' : 'bg-white border-gray-300 text-gray-900 placeholder-gray-500')"
              />
            </div>
          </div>

          <!-- Category Filter -->
          <div class="flex items-center space-x-2">
            <span :class="'text-sm font-medium ' + (isDark ? 'text-gray-300' : 'text-gray-700')">
              分类筛选：
            </span>
            <select
              v-model="selectedCategory"
              :class="'px-4 py-2 rounded-lg border transition-all duration-300 focus:outline-none focus:ring-2 focus:ring-blue-500/50 ' + (isDark ? 'bg-gray-700/50 border-gray-600 text-white' : 'bg-white border-gray-300 text-gray-900')"
            >
              <option value="">全部分类</option>
              <option value="交通">交通</option>
              <option value="餐饮">餐饮</option>
              <option value="住宿">住宿</option>
              <option value="办公">办公</option>
              <option value="其他">其他</option>
            </select>
          </div>

          <!-- Status Filter -->
          <div class="flex items-center space-x-2">
            <span :class="'text-sm font-medium ' + (isDark ? 'text-gray-300' : 'text-gray-700')">
              状态筛选：
            </span>
            <select
              v-model="selectedStatus"
              :class="'px-4 py-2 rounded-lg border transition-all duration-300 focus:outline-none focus:ring-2 focus:ring-blue-500/50 ' + (isDark ? 'bg-gray-700/50 border-gray-600 text-white' : 'bg-white border-gray-300 text-gray-900')"
            >
              <option value="">全部状态</option>
              <option value="已报销">已报销</option>
              <option value="未报销">未报销</option>
            </select>
          </div>
        </div>
      </div>

      <!-- Bills Table -->
      <div :class="'rounded-2xl overflow-hidden backdrop-blur-md border transition-all duration-300 ' + (isDark ? 'bg-gray-800/70 border-gray-700/50' : 'bg-white/70 border-gray-200/50')">
        <!-- Desktop Table -->
        <div class="hidden overflow-x-auto md:block">
          <table class="w-full">
            <thead>
              <tr :class="'border-b transition-colors duration-300 ' + (isDark ? 'border-gray-700' : 'border-gray-200')">
                <th :class="'text-left py-4 px-6 font-semibold transition-colors duration-300 ' + (isDark ? 'text-gray-300' : 'text-gray-700')">
                  账单名称
                </th>
                <th :class="'text-left py-4 px-6 font-semibold transition-colors duration-300 ' + (isDark ? 'text-gray-300' : 'text-gray-700')">
                  分类
                </th>
                <th :class="'text-left py-4 px-6 font-semibold transition-colors duration-300 ' + (isDark ? 'text-gray-300' : 'text-gray-700')">
                  金额
                </th>
                <th :class="'text-left py-4 px-6 font-semibold transition-colors duration-300 ' + (isDark ? 'text-gray-300' : 'text-gray-700')">
                  日期
                </th>
                <th :class="'text-left py-4 px-6 font-semibold transition-colors duration-300 ' + (isDark ? 'text-gray-300' : 'text-gray-700')">
                  来源
                </th>
                <th :class="'text-left py-4 px-6 font-semibold transition-colors duration-300 ' + (isDark ? 'text-gray-300' : 'text-gray-700')">
                  状态
                </th>
                <th :class="'text-center py-4 px-6 font-semibold transition-colors duration-300 ' + (isDark ? 'text-gray-300' : 'text-gray-700')">
                  操作
                </th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="bill in paginatedBills"
                :key="bill.id"
                :class="'border-b transition-all duration-300 hover:bg-opacity-50 ' + (isDark ? 'border-gray-700 hover:bg-gray-700/30' : 'border-gray-200 hover:bg-gray-50')"
              >
                <td class="py-4 px-6">
                  <div>
                    <p :class="'font-medium transition-colors duration-300 ' + (isDark ? 'text-white' : 'text-gray-900')">
                      {{ bill.name }}
                    </p>
                    <p :class="'text-sm mt-1 transition-colors duration-300 ' + (isDark ? 'text-gray-400' : 'text-gray-500')">
                      {{ bill.description }}
                    </p>
                  </div>
                </td>
                <td class="py-4 px-6">
                  <span :class="`px-3 py-1 rounded-full text-xs font-medium transition-colors duration-300 ${
                    getCategoryStyle(bill.category)
                  }`">
                    {{ bill.category }}
                  </span>
                </td>
                <td class="py-4 px-6">
                  <p :class="'font-semibold transition-colors duration-300 ' + (isDark ? 'text-white' : 'text-gray-900')">
                    ¥{{ bill.amount.toFixed(2) }}
                  </p>
                </td>
                <td class="py-4 px-6">
                  <p :class="'text-sm transition-colors duration-300 ' + (isDark ? 'text-gray-300' : 'text-gray-600')">
                    {{ bill.date }}
                  </p>
                </td>
                <td class="py-4 px-6">
                  <div class="flex items-center space-x-2">
                    <component
                      :is="getSourceIcon(bill.source)"
                      :size="16"
                      :class="getSourceIconColor(bill.source)"
                    />
                    <span :class="'text-sm transition-colors duration-300 ' + (isDark ? 'text-gray-300' : 'text-gray-600')">
                      {{ bill.source }}
                    </span>
                  </div>
                </td>
                <td class="py-4 px-6">
                  <span :class="`px-3 py-1 rounded-full text-xs font-medium transition-colors duration-300 ${
                    bill.status === '已报销'
                      ? 'bg-green-100 text-green-800 dark:bg-green-900/50 dark:text-green-300'
                      : 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900/50 dark:text-yellow-300'
                  }`">
                    {{ bill.status }}
                  </span>
                </td>
                <td class="py-4 px-6">
                  <div class="flex items-center justify-center space-x-2">
                    <button
                      :class="'p-2 rounded-lg transition-all duration-200 ' + (isDark ? 'hover:bg-gray-700 text-gray-300' : 'hover:bg-gray-100 text-gray-600')"
                      title="编辑"
                    >
                      <EditIcon :size="16" />
                    </button>
                    <button
                      :class="'p-2 rounded-lg transition-all duration-200 ' + (isDark ? 'hover:bg-gray-700 text-gray-300' : 'hover:bg-gray-100 text-gray-600')"
                      title="删除"
                      @click="confirmDeleteBill(bill)"
                    >
                      <Trash2Icon :size="16" />
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Mobile Card View -->
        <div class="md:hidden space-y-4">
          <div
            v-for="bill in paginatedBills"
            :key="bill.id"
            :class="'rounded-xl p-4 border transition-all duration-300 ' + (isDark ? 'bg-gray-800/50 border-gray-700/50' : 'bg-white/50 border-gray-200/50')"
          >
            <div class="flex justify-between items-start mb-3">
              <div>
                <h3 :class="'font-semibold transition-colors duration-300 ' + (isDark ? 'text-white' : 'text-gray-900')">
                  {{ bill.name }}
                </h3>
                <p :class="'text-sm mt-1 transition-colors duration-300 ' + (isDark ? 'text-gray-400' : 'text-gray-500')">
                  {{ bill.description }}
                </p>
              </div>
              <span :class="'px-2 py-1 rounded-full text-xs font-medium transition-colors duration-300 ' + (bill.status === '已报销' ? 'bg-green-100 text-green-800 dark:bg-green-900/50 dark:text-green-300' : 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900/50 dark:text-yellow-300')">
                {{ bill.status }}
              </span>
            </div>
            <div class="grid grid-cols-2 gap-2 text-sm">
              <div>
                <p class="text-gray-500 dark:text-gray-400">分类</p>
                <p :class="'font-medium transition-colors duration-300 ' + (isDark ? 'text-white' : 'text-gray-900')">
                  {{ bill.category }}
                </p>
              </div>
              <div>
                <p class="text-gray-500 dark:text-gray-400">金额</p>
                <p :class="'font-semibold transition-colors duration-300 ' + (isDark ? 'text-white' : 'text-gray-900')">
                  ¥{{ bill.amount.toFixed(2) }}
                </p>
              </div>
              <div>
                <p class="text-gray-500 dark:text-gray-400">日期</p>
                <p :class="'transition-colors duration-300 ' + (isDark ? 'text-gray-300' : 'text-gray-600')">
                  {{ bill.date }}
                </p>
              </div>
              <div>
                <p class="text-gray-500 dark:text-gray-400">来源</p>
                <div class="flex items-center space-x-1">
                  <component
                    :is="getSourceIcon(bill.source)"
                    :size="14"
                    :class="getSourceIconColor(bill.source)"
                  />
                  <span :class="'text-xs transition-colors duration-300 ' + (isDark ? 'text-gray-300' : 'text-gray-600')">
                    {{ bill.source }}
                  </span>
                </div>
              </div>
            </div>
            <div class="flex justify-end space-x-2 mt-3">
              <button
                :class="`p-2 rounded-lg transition-all duration-200 ${
                  isDark
                    ? 'hover:bg-gray-700 text-gray-300'
                    : 'hover:bg-gray-100 text-gray-600'
                }`"
              >
                <EditIcon :size="16" />
              </button>
              <button
                :class="`p-2 rounded-lg transition-all duration-200 ${
                  isDark
                    ? 'hover:bg-gray-700 text-gray-300'
                    : 'hover:bg-gray-100 text-gray-600'
                }`"
                @click="confirmDeleteBill(bill)"
              >
                <Trash2Icon :size="16" />
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Pagination -->
      <div :class="'flex justify-center items-center space-x-4 mt-6 ' + (isDark ? 'text-gray-300' : 'text-gray-600')">
        <button
          @click="goToPage(currentPage - 1)"
          :disabled="currentPage === 1"
          :class="'px-3 py-2 rounded-lg border transition-all duration-200 ' + (currentPage === 1 ? 'opacity-50 cursor-not-allowed' : isDark ? 'hover:bg-gray-700 border-gray-600 text-gray-300' : 'hover:bg-white border-gray-300 text-gray-700')"
        >
          <ChevronLeft :size="20" />
        </button>

        <span :class="text-sm">
          第 {{ currentPage }} 页，共 {{ totalPages }} 页
        </span>

        <button
          @click="goToPage(currentPage + 1)"
          :disabled="currentPage === totalPages"
          :class="`px-3 py-2 rounded-lg border transition-all duration-200 ${
            currentPage === totalPages
              ? 'opacity-50 cursor-not-allowed'
              : isDark
                ? 'hover:bg-gray-700 border-gray-600 text-gray-300'
                : 'hover:bg-white border-gray-300 text-gray-700'
          }`"
        >
          <ChevronRight :size="20" />
        </button>
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
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import {
  PlusIcon,
  SearchIcon,
  EditIcon,
  Trash2Icon,
  ChevronLeft,
  ChevronRight,
  SunIcon,
  MoonIcon,
  CreditCardIcon,
  SmartphoneIcon,
  DollarSignIcon,
  FileTextIcon
} from 'lucide-vue-next'
import { useTheme } from '../composables/useTheme'

const router = useRouter()
const { theme, toggleTheme, isDark } = useTheme()

// Search and filter state
const searchQuery = ref('')
const selectedCategory = ref('')
const selectedStatus = ref('')

// Pagination state
const currentPage = ref(1)
const itemsPerPage = 8

// Mock bills data
const allBills = ref([
  {
    id: 1,
    name: '星巴克咖啡',
    description: '工作日早晨咖啡',
    category: '餐饮',
    amount: 35.50,
    date: '2025-11-08',
    source: '微信支付',
    status: '已报销'
  },
  {
    id: 2,
    name: '滴滴出行',
    description: '从公司到客户处',
    category: '交通',
    amount: 45.00,
    date: '2025-11-08',
    source: '支付宝',
    status: '已报销'
  },
  {
    id: 3,
    name: '京东文具',
    description: '办公用品采购',
    category: '办公',
    amount: 128.90,
    date: '2025-11-07',
    source: '银行卡',
    status: '未报销'
  },
  {
    id: 4,
    name: '如家酒店',
    description: '出差住宿费用',
    category: '住宿',
    amount: 280.00,
    date: '2025-11-06',
    source: '手动添加',
    status: '已报销'
  },
  {
    id: 5,
    name: '美团外卖',
    description: '团队午餐',
    category: '餐饮',
    amount: 89.60,
    date: '2025-11-05',
    source: '支付宝',
    status: '未报销'
  },
  {
    id: 6,
    name: '地铁月卡',
    description: '11月地铁充值',
    category: '交通',
    amount: 200.00,
    date: '2025-11-01',
    source: '手动添加',
    status: '已报销'
  },
  {
    id: 7,
    name: '肯德基',
    description: '快餐消费',
    category: '餐饮',
    amount: 52.00,
    date: '2025-10-30',
    source: '微信支付',
    status: '未报销'
  },
  {
    id: 8,
    name: '办公用品',
    description: '打印纸和笔',
    category: '办公',
    amount: 67.30,
    date: '2025-10-28',
    source: '银行卡',
    status: '已报销'
  },
  {
    id: 9,
    name: '出租车',
    description: '夜间出行',
    category: '交通',
    amount: 78.50,
    date: '2025-10-25',
    source: '现金',
    status: '已报销'
  },
  {
    id: 10,
    name: '便利店',
    description: '日常用品',
    category: '其他',
    amount: 25.80,
    date: '2025-10-24',
    source: '手动添加',
    status: '未报销'
  },
  {
    id: 11,
    name: '海底捞',
    description: '团队聚餐',
    category: '餐饮',
    amount: 168.00,
    date: '2025-10-22',
    source: '支付宝',
    status: '已报销'
  },
  {
    id: 12,
    name: '滴滴专车',
    description: '商务出行',
    category: '交通',
    amount: 95.00,
    date: '2025-10-20',
    source: '微信支付',
    status: '已报销'
  },
  {
    id: 13,
    name: '超市购物',
    description: '生活用品',
    category: '其他',
    amount: 45.60,
    date: '2025-10-18',
    source: '银行卡',
    status: '未报销'
  },
  {
    id: 14,
    name: '打印服务',
    category: '办公',
    amount: 15.00,
    date: '2025-10-15',
    source: '手动添加',
    status: '已报销'
  },
  {
    id: 15,
    name: '麦当劳',
    description: '快餐消费',
    category: '餐饮',
    amount: 38.00,
    date: '2025-10-12',
    source: '微信支付',
    status: '未报销'
  }
])

// Computed properties for filtering and pagination
const filteredBills = computed(() => {
  return allBills.value.filter(bill => {
    const matchesSearch = bill.name.toLowerCase().includes(searchQuery.value.toLowerCase())
    const matchesCategory = !selectedCategory.value || bill.category === selectedCategory.value
    const matchesStatus = !selectedStatus.value || bill.status === selectedStatus.value
    return matchesSearch && matchesCategory && matchesStatus
  })
})

const totalPages = computed(() => Math.ceil(filteredBills.value.length / itemsPerPage))
const paginatedBills = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage
  const end = start + itemsPerPage
  return filteredBills.value.slice(start, end)
})

// Methods
const getCategoryStyle = (category) => {
  const styles = {
    '交通': 'bg-blue-100 text-blue-800 dark:bg-blue-900/50 dark:text-blue-300',
    '餐饮': 'bg-green-100 text-green-800 dark:bg-green-900/50 dark:text-green-300',
    '住宿': 'bg-purple-100 text-purple-800 dark:bg-purple-900/50 dark:text-purple-300',
    '办公': 'bg-orange-100 text-orange-800 dark:bg-orange-900/50 dark:text-orange-300',
    '其他': 'bg-gray-100 text-gray-800 dark:bg-gray-900/50 dark:text-gray-300'
  }
  return styles[category] || 'bg-gray-100 text-gray-800'
}

const getSourceIcon = (source) => {
  const icons = {
    '微信支付': CreditCardIcon,
    '支付宝': SmartphoneIcon,
    '银行卡': DollarSignIcon,
    '手动添加': FileTextIcon,
    '现金': DollarSignIcon
  }
  return icons[source] || FileTextIcon
}

const getSourceIconColor = (source) => {
  const colors = {
    '微信支付': 'text-green-500',
    '支付宝': 'text-blue-500',
    '银行卡': 'text-orange-500',
    '手动添加': 'text-purple-500',
    '现金': 'text-gray-500'
  }
  return colors[source] || 'text-gray-500'
}

const goToPage = (page) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
  }
}

const confirmDeleteBill = (bill) => {
  if (confirm(`确定要删除账单 "${bill.name}" 吗？`)) {
    const index = allBills.value.findIndex(b => b.id === bill.id)
    if (index > -1) {
      allBills.value.splice(index, 1)
    }
  }
}

onMounted(() => {
  // 初始化时如果有查询参数，应用过滤条件
  // 这里可以添加从路由参数读取的过滤条件
})
</script>

<style scoped>
/* Smooth transitions */
.transition-all {
  transition-property: all;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
}

/* Particle background optimization */
.dark .animate-pulse {
  opacity: 0.3;
}

/* Table hover effects */
tr:hover td {
  transform: translateY(-1px);
}

/* Mobile card hover effects */
.rounded-xl:hover {
  transform: translateY(-2px);
}

/* Button group spacing */
.flex.space-x-2 > * + * {
  margin-left: 0.5rem;
}

/* Focus states */
.focus\:ring-2:focus {
  outline: none;
}

/* Disabled states */
.opacity-50 {
  opacity: 0.5;
  cursor: not-allowed;
}
</style>