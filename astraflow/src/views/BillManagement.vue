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
      title="账单管理"
      :show-back-button="true"
      back-to="/"
      :show-theme-toggle="true"
    />

    <!-- Action Button Container -->
    <div class="action-buttons-container">
      <button class="add-bill-button" @click="openCreateInvoiceModal">
        <PlusIcon :size="20" class="button-icon" />
        添加发票
      </button>
    </div>

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
              <option value="待定/pending">待定 (pending)</option>
              <option value="通过/recognized">通过 (recognized)</option>
              <option value="已报销/confirmed">已报销 (confirmed)</option>
              <option value="驳回/rejected">驳回 (rejected)</option>
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
                <th class="table-header-cell" style="text-align: center;">账单名称</th>
                <th class="table-header-cell" style="text-align: center;">分类</th>
                <th class="table-header-cell" style="text-align: center;">金额</th>
                <th class="table-header-cell" style="text-align: center;">日期</th>
                <th class="table-header-cell" style="text-align: center;">来源</th>
                <th class="table-header-cell" style="text-align: center;">状态</th>
                <th class="table-header-cell" style="text-align: center;">操作</th>
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
                      {{ bill.vendor }}
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
                    {{ formatDate(bill.date) }}
                  </p>
                </td>
                <td class="table-cell table-cell-centered">
                  <div class="source-container">
                    <component
                      :is="getSourceIcon(bill.source)"
                      :size="16"
                      class="source-icon"
                      :class="getSourceIconClass(bill.source)"
                    />
                    <span class="source-text">
                      {{ bill.vendor }}
                    </span>
                  </div>
                </td>
                <td class="table-cell">
                  <span class="status-tag" :class="getStatusClass(bill.status)">
                    {{ formatStatus(bill.status) }}
                  </span>
                </td>
                <td class="table-cell table-cell-centered">
                  <div class="action-buttons">
                    <button class="action-button" title="编辑" @click="openEditInvoiceModal(bill)">
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
                {{ formatStatus(bill.status) }}
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
                <p class="card-value">{{ formatDate(bill.date) }}</p>
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
                    {{ bill.vendor }}
                  </span>
                </div>
              </div>
            </div>
            <div class="card-actions">
              <button class="action-button" @click="openEditInvoiceModal(bill)">
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

    <!-- Page Footer -->
    <PageFooter />

    <!-- Invoice Modal -->
    <InvoiceModal
      v-model:visible="showInvoiceModal"
      :invoice="editingInvoice"
      :loading="modalLoading"
      @submit="handleInvoiceSubmit"
      @close="handleModalClose"
    />
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
import InvoiceModal from '../components/InvoiceModal.vue'
import PageHeader from '../components/ui/PageHeader.vue'
import PageFooter from '../components/ui/PageFooter.vue'
import { useMessage } from 'naive-ui'
import {
  getInvoicesByUser,
  getInvoicesByTenant,
  getInvoices,
  createInvoice,
  updateInvoice,
  deleteInvoice
} from '../services/api/invoiceApi'
import { useUserStore } from '../stores/userStore'

const router = useRouter()
const { theme, toggleTheme, isDark } = useTheme()
const message = useMessage()

// Local state instead of store
const bills = ref([])
const storeLoading = ref(false)
const error = ref(null)
const pagination = ref({
  page: 1,
  size: 10,
  total: 0
})

// Modal state
const showInvoiceModal = ref(false)
const editingInvoice = ref(null)
const modalLoading = ref(false)

// Search and filter state
const searchQuery = ref('')
const selectedCategory = ref('')
const selectedStatus = ref('')

// Pagination state - make it reactive to store
const currentPage = ref(1)
const itemsPerPage = 10

// Get user store for multi-tenant logic
const userStore = useUserStore()

// Computed properties for filtering and pagination
const filteredBills = computed(() => {
  return bills.value.filter(bill => {
    const name = bill.name || bill.title
    const category = bill.category
    const status = bill.status

    const matchesSearch = !searchQuery.value ||
      (name && name.toLowerCase().includes(searchQuery.value.toLowerCase())) ||
      (bill.description && bill.description.toLowerCase().includes(searchQuery.value.toLowerCase()))
    const matchesCategory = !selectedCategory.value || category === selectedCategory.value
    // Handle bilingual status filtering (e.g. "待定/pending", "通过/recognized", etc.)
    let matchesStatus = true
    if (selectedStatus.value) {
      if (selectedStatus.value.includes('/')) {
        // Extract both Chinese and English parts from the selected status
        const [chinesePart, englishPart] = selectedStatus.value.split('/')
        matchesStatus = status === englishPart.trim() || status === chinesePart.trim()
      } else {
        // Fallback for direct match
        matchesStatus = status === selectedStatus.value
      }
    }
    return matchesSearch && matchesCategory && matchesStatus
  })
})

const totalPages = computed(() => {
  // Use pagination from store if available, otherwise fallback to local calculation
  if (pagination.value && pagination.value.total > 0) {
    return Math.ceil(pagination.value.total / itemsPerPage)
  }
  return Math.ceil(filteredBills.value.length / itemsPerPage)
})

const paginatedBills = computed(() => {
  // When using API pagination, we rely on the store data which is already paginated
  // But we still apply client-side filtering if needed
  return filteredBills.value
})

// Methods
const formatStatus = (status) => {
  // Map backend status values to bilingual display
  const statusMap = {
    'pending': '待定 (pending)',
    '待定': '待定 (pending)',
    'recognized': '通过 (recognized)',
    '通过': '通过 (recognized)',
    'confirmed': '已报销 (confirmed)',
    '已报销': '已报销 (confirmed)',
    'rejected': '驳回 (rejected)',
    '驳回': '驳回 (rejected)',
    '待审批': '待定 (pending)',
    '已拒绝': '驳回 (rejected)'
  }
  return statusMap[status] || status
}

const formatDate = (dateString) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  if (isNaN(date.getTime())) return dateString // 如果日期无效，返回原始值
  return date.toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit'
  })
}

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
  // Check if the source field is paymentSource for invoices
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

const goToPage = async (page) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
    // Fetch bills for the new page
    await loadBills()
  }
}

const confirmDeleteBill = async (bill) => {
  if (confirm(`确定要删除发票 "${bill.title || bill.name}" 吗？`)) {
    try {
      await deleteInvoice(bill.id)
      // Remove the deleted invoice from local state
      bills.value = bills.value.filter(item => item.id !== bill.id)
      // Show success message
      message.success('发票删除成功')
    } catch (err) {
      console.error('删除发票失败:', err)
      message.error('删除发票失败: ' + (err.message || '未知错误'))
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
  // Map backend statuses to UI statuses
  if (status === 'confirmed' || status === '已报销') {
    return 'status-approved'
  } else if (status === 'recognized' || status === '通过') {
    return 'status-recognized'
  } else if (status === 'pending' || status === '待审批' || status === '待定') {
    return 'status-pending'
  } else if (status === 'rejected' || status === '已拒绝' || status === '驳回') {
    return 'status-rejected'
  }
  return 'status-pending'
}

// Computed property for loading state
const loading = computed(() => {
  return storeLoading.value || modalLoading.value
})

// Method to open invoice modal for creating new invoice
const openCreateInvoiceModal = () => {
  editingInvoice.value = null
  showInvoiceModal.value = true
}

// Method to open invoice modal for editing existing invoice
const openEditInvoiceModal = (bill) => {
  editingInvoice.value = bill
  showInvoiceModal.value = true
}

// Method to handle invoice form submission
const handleInvoiceSubmit = async (formData) => {
  modalLoading.value = true
  try {
    if (editingInvoice.value) {
      // Update existing invoice
      await updateInvoice(editingInvoice.value.id, formData)
      message.success('发票更新成功')

      // Update the invoice in local state
      const index = bills.value.findIndex(bill => bill.id === editingInvoice.value.id)
      if (index !== -1) {
        bills.value[index] = {
          ...bills.value[index],
          ...formData,
          id: editingInvoice.value.id,
          name: formData.invoiceNumber || formData.title || '未命名发票',
          invoiceNumber: formData.invoice_number || formData.invoiceNumber
        }
      }
    } else {
      // Create new invoice
      const response = await createInvoice(formData)
      if (response?.data?.invoice) {
        const newInvoice = response.data.invoice
        // Add the new invoice to local state
        bills.value.unshift({
          id: newInvoice.id,
          name: newInvoice.invoice_number || newInvoice.title || '未命名发票',
          amount: newInvoice.amount,
          date: newInvoice.invoice_date || newInvoice.date,
          category: newInvoice.category,
          status: newInvoice.status || 'pending',
          vendor: newInvoice.vendor,
          imageUrl: newInvoice.image_url,
          description: newInvoice.description || '',
          invoiceNumber: newInvoice.invoice_number,
          taxId: newInvoice.taxId,
          paymentSource: newInvoice.payment_source,
          receiptUrl: newInvoice.receipt_url || newInvoice.receiptUrl || ''
        })
      }
      message.success('发票创建成功')
    }
    // Close modal and refresh list
    showInvoiceModal.value = false
    editingInvoice.value = null
  } catch (err) {
    console.error('操作发票失败:', err)
    message.error('操作发票失败: ' + (err.message || '未知错误'))
  } finally {
    modalLoading.value = false
  }
}

// Method to handle modal close
const handleModalClose = () => {
  showInvoiceModal.value = false
  editingInvoice.value = null
}

const loadBills = async () => {
  storeLoading.value = true
  error.value = null
  try {
    let response

    // Determine which endpoint to use based on user type
      response = await getInvoicesByUser({
        page: currentPage.value,
        page_size: itemsPerPage
      })
    // if (userStore.isPersonalUser()) {
    //   // Personal user - get invoices by user
    //   response = await getInvoicesByUser({
    //     page: currentPage.value,
    //     page_size: itemsPerPage
    //   })
    // } else if (userStore.user?.tenantId) {
    //   // Tenant user - get invoices by tenan
    //   response = await getInvoicesByTenant({
    //     page: currentPage.value,
    //     page_size: itemsPerPage
    //   })
    // } else {
    //   // Default - get all invoices (for admin users)
    //   response = await getInvoices({
    //     page: currentPage.value,
    //     page_size: itemsPerPage
    //   })
    // }

    if (response?.data) {
      // Map invoice fields to bill fields for compatibility
      bills.value = response.data.invoices?.map(invoice => ({
        id: invoice.id,
        name: invoice.invoice_number || invoice.title || '未命名发票',
        amount: invoice.amount,
        date: invoice.invoice_date || invoice.date,
        category: invoice.category,
        status: invoice.status || invoice.status,
        vendor: invoice.vendor,
        imageUrl: invoice.image_url,
        description: invoice.description || '',
        invoiceNumber: invoice.invoice_number,
        taxId: invoice.taxId,
        paymentSource: invoice.payment_source,
        receiptUrl: invoice.receipt_url || invoice.receiptUrl || ''
      })) || []

      // Update pagination info with new fields from backend
      if (response.data.page !== undefined) {
        pagination.value.page = response.data.page
      }
      if (response.data.size !== undefined) {
        pagination.value.size = response.data.size
      }
      if (response.data.total !== undefined) {
        pagination.value.total = response.data.total
      }
      if (response.data.totalPages !== undefined) {
        // Update totalPages using the new field from backend
        totalPages.value = response.data.totalPages
      }
    }
  } catch (err) {
    console.error('加载发票失败:', err)
    message.error('加载发票失败: ' + (err.message || '未知错误'))
  } finally {
    storeLoading.value = false
  }
}

onMounted(async () => {
  // Load bills when component mounts
  await loadBills()
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
  /* min-height: 85vh; */
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
  min-height: 70vh;
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
  text-align: center;
  vertical-align: middle;
  height: 100%;
}

.table-cell-centered {
  display: table-cell !important;
  text-align: center !important;
  vertical-align: middle !important;
  height: 100%;
}

.table-cell-centered .source-container,
.table-cell-centered .action-buttons {
  justify-content: center;
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

.status-rejected {
  background-color: #fee2e2;
  color: #dc2626;
}

.app-container[data-theme="dark"] .status-rejected {
  background-color: #b91c1c;
  color: #fecaca;
}

.status-recognized {
  background-color: #dbeafe;
  color: #2563eb;
}

.app-container[data-theme="dark"] .status-recognized {
  background-color: #1e40af;
  color: #bfdbfe;
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

.action-buttons-container {
  max-width: 80rem;
  margin: 0 auto;
  padding: 0 1rem;
  display: flex;
  justify-content: flex-end;
  padding-top: 1rem;
}

@media (min-width: 640px) {
  .action-buttons-container {
    padding: 0 1.5rem;
    padding-top: 1rem;
  }
}

@media (min-width: 1024px) {
  .action-buttons-container {
    padding: 0 2rem;
    padding-top: 1rem;
  }
}
</style>