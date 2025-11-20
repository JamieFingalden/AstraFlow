<template>
  <div class="app-container" :data-theme="isDark ? 'dark' : 'light'">
    <!-- Particle background effect -->
    <div class="particle-background">
      <div class="particle particle-cyan"></div>
      <div class="particle particle-blue"></div>
      <div class="particle particle-purple"></div>
    </div>

    <!-- Header -->
    <header class="app-header">
      <div class="header-container">
        <div class="header-left">
          <!-- 返回按钮 -->
          <router-link to="/" class="back-button" title="返回首页">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="m15 18-6-6 6-6"/>
            </svg>
            <span class="back-text">返回</span>
          </router-link>

          <div class="brand-name">
            AstraFlow
          </div>
          <h1 class="page-title">账单管理</h1>
        </div>

        <div class="header-right">
          <!-- 添加账单按钮 -->
          <button class="add-bill-button">
            <PlusIcon :size="20" class="button-icon" />
            添加账单
          </button>

          <!-- 主题切换按钮 -->
          <button @click="toggleTheme" class="theme-toggle">
            <SunIcon v-if="isDark" :size="20" />
            <MoonIcon v-else :size="20" />
          </button>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="main-content">
      <!-- Search and Filter Section -->
      <div class="search-filter-container">
        <div class="search-filter-content">
          <!-- Search Bar -->
          <div class="search-container">
            <div class="search-wrapper">
              <SearchIcon class="search-icon" :size="20" />
              <input
                v-model="searchQuery"
                type="text"
                placeholder="搜索账单名称..."
                class="search-input"
              />
            </div>
          </div>

          <!-- Category Filter -->
          <div class="filter-group">
            <span class="filter-label">
              分类筛选：
            </span>
            <select v-model="selectedCategory" class="filter-select">
              <option value="">全部分类</option>
              <option value="交通">交通</option>
              <option value="餐饮">餐饮</option>
              <option value="住宿">住宿</option>
              <option value="办公">办公</option>
              <option value="其他">其他</option>
            </select>
          </div>

          <!-- Status Filter -->
          <div class="filter-group">
            <span class="filter-label">
              状态筛选：
            </span>
            <select v-model="selectedStatus" class="filter-select">
              <option value="">全部状态</option>
              <option value="已报销">已报销</option>
              <option value="未报销">未报销</option>
            </select>
          </div>
        </div>
      </div>

      <!-- Bills Table -->
      <div class="bills-table-container">
        <!-- Desktop Table -->
        <div class="desktop-table">
          <table class="bills-table">
            <thead>
              <tr class="table-header">
                <th class="table-header-cell">账单名称</th>
                <th class="table-header-cell">分类</th>
                <th class="table-header-cell">金额</th>
                <th class="table-header-cell">日期</th>
                <th class="table-header-cell">来源</th>
                <th class="table-header-cell">状态</th>
                <th class="table-header-cell">操作</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="bill in paginatedBills"
                :key="bill.id"
                class="table-row"
              >
                <td class="table-cell">
                  <div>
                    <p class="bill-name">
                      {{ bill.name }}
                    </p>
                    <p class="bill-description">
                      {{ bill.description }}
                    </p>
                  </div>
                </td>
                <td class="table-cell">
                  <span class="category-tag" :class="getCategoryClass(bill.category)">
                    {{ bill.category }}
                  </span>
                </td>
                <td class="table-cell">
                  <p class="bill-amount">
                    ¥{{ bill.amount.toFixed(2) }}
                  </p>
                </td>
                <td class="table-cell">
                  <p class="bill-date">
                    {{ bill.date }}
                  </p>
                </td>
                <td class="table-cell">
                  <div class="source-container">
                    <component
                      :is="getSourceIcon(bill.source)"
                      :size="16"
                      class="source-icon"
                      :class="getSourceIconClass(bill.source)"
                    />
                    <span class="source-text">
                      {{ bill.source }}
                    </span>
                  </div>
                </td>
                <td class="table-cell">
                  <span class="status-tag" :class="getStatusClass(bill.status)">
                    {{ bill.status }}
                  </span>
                </td>
                <td class="table-cell">
                  <div class="action-buttons">
                    <button class="action-button" title="编辑">
                      <EditIcon :size="16" />
                    </button>
                    <button class="action-button" title="删除" @click="confirmDeleteBill(bill)">
                      <Trash2Icon :size="16" />
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Mobile Card View -->
        <div class="mobile-cards">
          <div
            v-for="bill in paginatedBills"
            :key="bill.id"
            class="bill-card"
          >
            <div class="card-header">
              <div>
                <h3 class="card-title">
                  {{ bill.name }}
                </h3>
                <p class="card-description">
                  {{ bill.description }}
                </p>
              </div>
              <span class="status-tag" :class="getStatusClass(bill.status)">
                {{ bill.status }}
              </span>
            </div>
            <div class="card-grid">
              <div>
                <p class="card-label">分类</p>
                <p class="card-value">{{ bill.category }}</p>
              </div>
              <div>
                <p class="card-label">金额</p>
                <p class="card-value">{{ bill.amount.toFixed(2) }}</p>
              </div>
              <div>
                <p class="card-label">日期</p>
                <p class="card-value">{{ bill.date }}</p>
              </div>
              <div>
                <p class="card-label">来源</p>
                <div class="source-container">
                  <component
                    :is="getSourceIcon(bill.source)"
                    :size="14"
                    class="source-icon"
                    :class="getSourceIconClass(bill.source)"
                  />
                  <span class="source-text">
                    {{ bill.source }}
                  </span>
                </div>
              </div>
            </div>
            <div class="card-actions">
              <button class="action-button">
                <EditIcon :size="16" />
              </button>
              <button class="action-button" @click="confirmDeleteBill(bill)">
                <Trash2Icon :size="16" />
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Pagination -->
      <div class="pagination-container">
        <button
          @click="goToPage(currentPage - 1)"
          :disabled="currentPage === 1"
          class="pagination-button"
          :class="{ 'disabled': currentPage === 1 }"
        >
          <ChevronLeft :size="20" />
        </button>

        <span class="pagination-info">
          第 {{ currentPage }} 页，共 {{ totalPages }} 页
        </span>

        <button
          @click="goToPage(currentPage + 1)"
          :disabled="currentPage === totalPages"
          class="pagination-button"
          :class="{ 'disabled': currentPage === totalPages }"
        >
          <ChevronRight :size="20" />
        </button>
      </div>
    </main>

    <!-- Footer -->
    <footer class="app-footer">
      <div class="footer-container">
        <p class="footer-text">
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

// Helper methods for dynamic classes
const getCategoryClass = (category) => {
  const classes = {
    '交通': 'category-transport',
    '餐饮': 'category-food',
    '住宿': 'category-lodging',
    '办公': 'category-office',
    '其他': 'category-other'
  }
  return classes[category] || 'category-other'
}

const getSourceIconClass = (source) => {
  const classes = {
    '微信支付': 'source-wechat',
    '支付宝': 'source-alipay',
    '银行卡': 'source-bank',
    '手动添加': 'source-manual',
    '现金': 'source-cash'
  }
  return classes[source] || 'source-manual'
}

const getStatusClass = (status) => {
  return status === '已报销' ? 'status-approved' : 'status-pending'
}

onMounted(() => {
  // 初始化时如果有查询参数，应用过滤条件
  // 这里可以添加从路由参数读取的过滤条件
})
</script>

<style scoped>
/* Main Container */
.app-container {
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
  color: inherit;
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
  line-height: 1.2;
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
  flex: 0.2; /* Don't grow */
}

.add-bill-button {
  padding: 0.5rem 1.5rem;
  border-radius: 0.5rem;
  font-weight: 600;
  border: none;
  cursor: pointer;
  transition: all 0.3s ease;
  transform: scale(1);
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.add-bill-button:hover {
  transform: scale(1.05);
}

.app-container[data-theme="dark"] .add-bill-button {
  background: linear-gradient(135deg, #2563eb, #06b6d4);
  color: white;
  box-shadow: 0 10px 15px -3px rgba(37, 99, 235, 0.25);
}

.app-container[data-theme="dark"] .add-bill-button:hover {
  background: linear-gradient(135deg, #3b82f6, #22d3ee);
  box-shadow: 0 20px 25px -5px rgba(37, 99, 235, 0.3);
}

.app-container[data-theme="light"] .add-bill-button {
  background: linear-gradient(135deg, #3b82f6, #06b6d4);
  color: white;
  box-shadow: 0 10px 15px -3px rgba(37, 99, 235, 0.3);
}

.app-container[data-theme="light"] .add-bill-button:hover {
  background: linear-gradient(135deg, #2563eb, #0891b2);
  box-shadow: 0 20px 25px -5px rgba(37, 99, 235, 0.4);
}

.button-icon {
  margin-right: 0.5rem;
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

/* Search and Filter Section */
.search-filter-container {
  border-radius: 1rem;
  padding: 1.5rem;
  margin-bottom: 1.5rem;
  backdrop-filter: blur(12px);
  border: 1px solid;
  transition: all 0.3s ease;
}

.app-container[data-theme="dark"] .search-filter-container {
  background-color: rgba(31, 41, 55, 0.7);
  border-color: rgba(55, 65, 81, 0.5);
}

.app-container[data-theme="light"] .search-filter-container {
  background-color: rgba(255, 255, 255, 0.7);
  border-color: rgba(229, 231, 235, 0.5);
}

.search-filter-content {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

@media (min-width: 768px) {
  .search-filter-content {
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
  }
}

/* Search Container */
.search-container {
  flex: 1;
  width: 100%;
}

@media (min-width: 768px) {
  .search-container {
    width: auto;
  }
}

.search-wrapper {
  position: relative;
}

.search-icon {
  position: absolute;
  left: 1rem;
  top: 50%;
  transform: translateY(-50%);
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .search-icon {
  color: #9ca3af;
}

.app-container[data-theme="light"] .search-icon {
  color: #6b7280;
}

.search-input {
  width: 100%;
  padding: 0.75rem 1rem 0.75rem 2.5rem;
  border-radius: 0.5rem;
  border: 1px solid;
  transition: all 0.3s ease;
  outline: none;
}

.app-container[data-theme="dark"] .search-input {
  background-color: rgba(55, 65, 81, 0.5);
  border-color: #4b5563;
  color: white;
  placeholder-color: #9ca3af;
}

.app-container[data-theme="light"] .search-input {
  background-color: white;
  border-color: #d1d5db;
  color: #111827;
  placeholder-color: #6b7280;
}

.search-input:focus {
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.5);
}

/* Filter Group */
.filter-group {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.filter-label {
  font-size: 0.875rem;
  font-weight: 500;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .filter-label {
  color: #d1d5db;
}

.app-container[data-theme="light"] .filter-label {
  color: #374151;
}

.filter-select {
  padding: 0.5rem 1rem;
  border-radius: 0.5rem;
  border: 1px solid;
  transition: all 0.3s ease;
  outline: none;
  cursor: pointer;
}

.app-container[data-theme="dark"] .filter-select {
  background-color: rgba(55, 65, 81, 0.5);
  border-color: #4b5563;
  color: white;
}

.app-container[data-theme="light"] .filter-select {
  background-color: white;
  border-color: #d1d5db;
  color: #111827;
}

.filter-select:focus {
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.5);
}

/* Bills Table Container */
.bills-table-container {
  border-radius: 1rem;
  overflow: hidden;
  backdrop-filter: blur(12px);
  border: 1px solid;
  transition: all 0.3s ease;
}

.app-container[data-theme="dark"] .bills-table-container {
  background-color: rgba(31, 41, 55, 0.7);
  border-color: rgba(55, 65, 81, 0.5);
}

.app-container[data-theme="light"] .bills-table-container {
  background-color: rgba(255, 255, 255, 0.7);
  border-color: rgba(229, 231, 235, 0.5);
}

/* Desktop Table */
.desktop-table {
  display: none;
}

@media (min-width: 768px) {
  .desktop-table {
    display: block;
  }
}

.bills-table {
  width: 100%;
  border-collapse: collapse;
}

.table-header {
  border-bottom: 1px solid;
  transition: border-color 0.3s ease;
}

.app-container[data-theme="dark"] .table-header {
  border-color: #374151;
}

.app-container[data-theme="light"] .table-header {
  border-color: #e5e7eb;
}

.table-header-cell {
  text-align: left;
  padding: 1rem 1.5rem;
  font-weight: 600;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .table-header-cell {
  color: #d1d5db;
}

.app-container[data-theme="light"] .table-header-cell {
  color: #374151;
}

.table-row {
  border-bottom: 1px solid;
  transition: all 0.3s ease;
}

.app-container[data-theme="dark"] .table-row {
  border-color: #374151;
}

.app-container[data-theme="light"] .table-row {
  border-color: #e5e7eb;
}

.table-row:last-child {
  border-bottom: none;
}

.table-row:hover {
  background-color: rgba(55, 65, 81, 0.15);
}

.app-container[data-theme="light"] .table-row:hover {
  background-color: rgba(249, 250, 251, 0.5);
}

.table-cell {
  padding: 1rem 1.5rem;
}

.bill-name {
  font-weight: 500;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .bill-name {
  color: white;
}

.app-container[data-theme="light"] .bill-name {
  color: #111827;
}

.bill-description {
  font-size: 0.875rem;
  margin-top: 0.25rem;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .bill-description {
  color: #9ca3af;
}

.app-container[data-theme="light"] .bill-description {
  color: #6b7280;
}

.category-tag {
  padding: 0.25rem 0.75rem;
  border-radius: 9999px;
  font-size: 0.75rem;
  font-weight: 500;
}

.category-transport {
  background-color: #dbeafe;
  color: #1d4ed8;
}

.app-container[data-theme="dark"] .category-transport {
  background-color: #1e3a8a;
  color: #bfdbfe;
}

.category-food {
  background-color: #dcfce7;
  color: #166534;
}

.app-container[data-theme="dark"] .category-food {
  background-color: #14532d;
  color: #bbf7d0;
}

.category-lodging {
  background-color: #f3e8ff;
  color: #6b21a8;
}

.app-container[data-theme="dark"] .category-lodging {
  background-color: #4c1d95;
  color: #e9d5ff;
}

.category-office {
  background-color: #ffedd5;
  color: #9a3412;
}

.app-container[data-theme="dark"] .category-office {
  background-color: #7c2d12;
  color: #fed7aa;
}

.category-other {
  background-color: #e5e7eb;
  color: #374151;
}

.app-container[data-theme="dark"] .category-other {
  background-color: #374151;
  color: #d1d5db;
}

.bill-amount {
  font-weight: 600;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .bill-amount {
  color: white;
}

.app-container[data-theme="light"] .bill-amount {
  color: #111827;
}

.bill-date {
  font-size: 0.875rem;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .bill-date {
  color: #9ca3af;
}

.app-container[data-theme="light"] .bill-date {
  color: #4b5563;
}

.source-container {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.source-icon {
  width: 1rem;
  height: 1rem;
}

.source-wechat {
  color: #10b981;
}

.source-alipay {
  color: #3b82f6;
}

.source-bank {
  color: #f59e0b;
}

.source-manual {
  color: #8b5cf6;
}

.source-cash {
  color: #6b7280;
}

.status-tag {
  padding: 0.25rem 0.75rem;
  border-radius: 9999px;
  font-size: 0.75rem;
  font-weight: 500;
}

.status-approved {
  background-color: #dcfce7;
  color: #16a34a;
}

.app-container[data-theme="dark"] .status-approved {
  background-color: #15803d;
  color: #bbf7d0;
}

.status-pending {
  background-color: #fef3c7;
  color: #d97706;
}

.app-container[data-theme="dark"] .status-pending {
  background-color: #a16207;
  color: #fde68a;
}

.action-buttons {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
}

.action-button {
  padding: 0.5rem;
  border-radius: 0.5rem;
  border: none;
  background: none;
  cursor: pointer;
  transition: all 0.2s ease;
}

.app-container[data-theme="dark"] .action-button {
  color: #d1d5db;
}

.app-container[data-theme="dark"] .action-button:hover {
  background-color: #374151;
}

.app-container[data-theme="light"] .action-button {
  color: #4b5563;
}

.app-container[data-theme="light"] .action-button:hover {
  background-color: #f3f4f6;
}

/* Mobile Card View */
.mobile-cards {
  display: block;
}

@media (min-width: 768px) {
  .mobile-cards {
    display: none;
  }
}

.bill-card {
  border-radius: 0.75rem;
  padding: 1rem;
  border: 1px solid;
  transition: all 0.3s ease;
}

.app-container[data-theme="dark"] .bill-card {
  background-color: rgba(31, 41, 55, 0.5);
  border-color: rgba(55, 65, 81, 0.5);
}

.app-container[data-theme="light"] .bill-card {
  background-color: rgba(255, 255, 255, 0.5);
  border-color: rgba(229, 231, 235, 0.5);
}

.bill-card:hover {
  transform: translateY(-2px);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 0.75rem;
}

.card-title {
  font-weight: 500;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .card-title {
  color: white;
}

.app-container[data-theme="light"] .card-title {
  color: #111827;
}

.card-description {
  font-size: 0.875rem;
  margin-top: 0.25rem;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .card-description {
  color: #9ca3af;
}

.app-container[data-theme="light"] .card-description {
  color: #6b7280;
}

.card-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 0.5rem;
  font-size: 0.875rem;
}

.card-label {
  color: #9ca3af;
}

.app-container[data-theme="dark"] .card-label {
  color: #6b7280;
}

.card-value {
  font-weight: 500;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .card-value {
  color: white;
}

.app-container[data-theme="light"] .card-value {
  color: #111827;
}

.card-actions {
  display: flex;
  justify-content: flex-end;
  gap: 0.5rem;
  margin-top: 0.75rem;
}

/* Pagination */
.pagination-container {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1rem;
  margin-top: 1.5rem;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .pagination-container {
  color: #d1d5db;
}

.app-container[data-theme="light"] .pagination-container {
  color: #4b5563;
}

.pagination-button {
  padding: 0.5rem 0.75rem;
  border-radius: 0.5rem;
  border: 1px solid;
  background: none;
  cursor: pointer;
  transition: all 0.2s ease;
}

.app-container[data-theme="dark"] .pagination-button {
  border-color: #4b5563;
  color: #d1d5db;
}

.app-container[data-theme="dark"] .pagination-button:hover:not(.disabled) {
  background-color: #374151;
}

.app-container[data-theme="light"] .pagination-button {
  border-color: #d1d5db;
  color: #374151;
}

.app-container[data-theme="light"] .pagination-button:hover:not(.disabled) {
  background-color: #f9fafb;
}

.pagination-button.disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.pagination-info {
  font-size: 0.875rem;
}

/* Footer */
.app-footer {
  position: relative;
  z-index: 10;
  backdrop-filter: blur(12px);
  border-top: 1px solid;
  transition: all 0.3s ease;
}

.app-container[data-theme="dark"] .app-footer {
  background-color: rgba(31, 41, 55, 0.7);
  border-color: rgba(55, 65, 81, 0.5);
}

.app-container[data-theme="light"] .app-footer {
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

.footer-text {
  text-align: center;
  font-size: 0.875rem;
  transition: color 0.3s ease;
}

.app-container[data-theme="dark"] .footer-text {
  color: #9ca3af;
}

.app-container[data-theme="light"] .footer-text {
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

/* Responsive adjustments */
@media (max-width: 767px) {
  .header-left,
  .header-right {
    flex: 1;
  }

  .add-bill-button {
    padding: 0.5rem 1rem;
    font-size: 0.875rem;
  }
}
</style>