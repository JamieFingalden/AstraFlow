<template>
  <form @submit.prevent="submitForm" class="invoice-form">
    <div class="form-grid">
      <!-- 发票号码 -->
      <div class="form-group">
        <label for="invoiceNumber" class="form-label">
          发票号码 <span class="required">*</span>
        </label>
        <input
          id="invoiceNumber"
          v-model="formData.invoiceNumber"
          type="text"
          class="form-input"
          placeholder="请输入发票号码"
          required
        />
      </div>

      <!-- 发票日期 -->
      <div class="form-group">
        <label for="invoiceDate" class="form-label">
          发票日期 <span class="required">*</span>
        </label>
        <input
          id="invoiceDate"
          v-model="formData.date"
          type="date"
          class="form-input"
          required
        />
      </div>

      <!-- 金额 -->
      <div class="form-group">
        <label for="amount" class="form-label">
          金额 <span class="required">*</span>
        </label>
        <input
          id="amount"
          v-model="formData.amount"
          type="number"
          step="0.01"
          class="form-input"
          placeholder="0.00"
          required
        />
      </div>

      <!-- 供应商 -->
      <div class="form-group">
        <label for="vendor" class="form-label">
          供应商/商户 <span class="required">*</span>
        </label>
        <input
          id="vendor"
          v-model="formData.vendor"
          type="text"
          class="form-input"
          placeholder="请输入供应商名称"
          required
        />
      </div>

      <!-- 分类 -->
      <div class="form-group">
        <label for="category" class="form-label">
          分类 <span class="required">*</span>
        </label>
        <select
          id="category"
          v-model="formData.category"
          class="form-select"
          required
        >
          <option value="">请选择分类</option>
          <option value="餐饮">餐饮</option>
          <option value="交通">交通</option>
          <option value="住宿">住宿</option>
          <option value="办公">办公</option>
          <option value="其他">其他</option>
        </select>
      </div>

      <!-- 税号 -->
      <div class="form-group">
        <label for="taxId" class="form-label">税号</label>
        <input
          id="taxId"
          v-model="formData.taxId"
          type="text"
          class="form-input"
          placeholder="请输入税号"
        />
      </div>

      <!-- 支付来源 -->
      <div class="form-group">
        <label for="paymentSource" class="form-label">支付来源</label>
        <select
          id="paymentSource"
          v-model="formData.paymentSource"
          class="form-select"
        >
          <option value="">请选择支付来源</option>
          <option value="微信支付">微信支付</option>
          <option value="支付宝">支付宝</option>
          <option value="银行卡">银行卡</option>
          <option value="手动添加">手动添加</option>
          <option value="现金">现金</option>
        </select>
      </div>

      <!-- 状态 -->
      <div class="form-group">
        <label for="status" class="form-label">状态</label>
        <select
          id="status"
          v-model="formData.status"
          class="form-select"
        >
          <option value="pending">待处理</option>
          <option value="confirmed">已确认</option>
          <option value="rejected">已拒绝</option>
        </select>
      </div>

      <!-- 图片URL -->
      <div class="form-group">
        <label for="imageUrl" class="form-label">发票图片URL</label>
        <input
          id="imageUrl"
          v-model="formData.imageUrl"
          type="url"
          class="form-input"
          placeholder="请输入发票图片链接"
        />
      </div>
    </div>

    <!-- 描述 -->
    <div class="form-group">
      <label for="description" class="form-label">描述</label>
      <textarea
        id="description"
        v-model="formData.description"
        class="form-textarea"
        placeholder="请输入发票描述或备注"
        rows="3"
      ></textarea>
    </div>

    <!-- 表单操作按钮 -->
    <div class="form-actions">
      <button
        type="button"
        class="btn btn-secondary"
        @click="onCancel"
      >
        取消
      </button>
      <button
        type="submit"
        class="btn btn-primary"
        :disabled="loading"
      >
        <span v-if="loading" class="btn-loading">处理中...</span>
        <span v-else>{{ isEditMode ? '更新发票' : '创建发票' }}</span>
      </button>
    </div>

    <!-- 错误提示 -->
    <div
      v-if="error"
      class="error-message"
      :data-theme="isDark ? 'dark' : 'light'"
    >
      {{ error }}
    </div>
  </form>
</template>

<script setup>
import { ref, watch, defineProps, defineEmits, computed } from 'vue'
import { useTheme } from '../composables/useTheme'

// Props
const props = defineProps({
  invoice: {
    type: Object,
    default: null
  },
  loading: {
    type: Boolean,
    default: false
  }
})

// Emits
const emit = defineEmits(['submit', 'cancel'])

// Theme
const { isDark } = useTheme()

// Form data
const formData = ref({
  invoiceNumber: '',
  date: '',
  amount: '',
  vendor: '',
  category: '',
  taxId: '',
  paymentSource: '',
  status: 'pending',
  imageUrl: '',
  description: ''
})

// Error state
const error = ref('')

// Check if in edit mode
const isEditMode = computed(() => !!props.invoice)

// Initialize form data when invoice prop changes
watch(
  () => props.invoice,
  (newInvoice) => {
    if (newInvoice) {
      formData.value = {
        invoiceNumber: newInvoice.invoiceNumber || newInvoice.invoice_number || newInvoice.title || '',
        date: newInvoice.date || newInvoice.invoice_date || '',
        amount: newInvoice.amount || '',
        vendor: newInvoice.vendor || '',
        category: newInvoice.category || '',
        taxId: newInvoice.taxId || newInvoice.tax_id || '',
        paymentSource: newInvoice.paymentSource || newInvoice.payment_source || '',
        status: newInvoice.status || 'pending',
        imageUrl: newInvoice.imageUrl || newInvoice.image_url || newInvoice.imageURL || '',
        description: newInvoice.description || ''
      }
    } else {
      // Reset form for new invoice
      formData.value = {
        invoiceNumber: '',
        date: '',
        amount: '',
        vendor: '',
        category: '',
        taxId: '',
        paymentSource: '',
        status: 'pending',
        imageUrl: '',
        description: ''
      }
    }
  },
  { immediate: true }
)

// Validate form
const validateForm = () => {
  error.value = ''

  if (!formData.value.invoiceNumber.trim()) {
    error.value = '请输入发票号码'
    return false
  }

  if (!formData.value.date) {
    error.value = '请选择发票日期'
    return false
  }

  if (!formData.value.amount || parseFloat(formData.value.amount) <= 0) {
    error.value = '请输入有效的金额'
    return false
  }

  if (!formData.value.vendor.trim()) {
    error.value = '请输入供应商/商户名称'
    return false
  }

  if (!formData.value.category) {
    error.value = '请选择分类'
    return false
  }

  return true
}

// Handle form submission
const submitForm = () => {
  if (!validateForm()) {
    return
  }

  // Format data for submission with field names that match backend JSON tags
  const submitData = {
    invoice_number: formData.value.invoiceNumber,
    invoice_date: formData.value.date ? new Date(formData.value.date).toISOString() : '',
    amount: parseFloat(formData.value.amount),
    vendor: formData.value.vendor,
    category: formData.value.category,
    taxId: formData.value.taxId,
    payment_source: formData.value.paymentSource,
    status: formData.value.status || 'pending',
    image_url: formData.value.imageUrl || '',
    description: formData.value.description || ''
  }

  emit('submit', submitData)
}

// Handle cancel
const onCancel = () => {
  emit('cancel')
}
</script>

<style scoped>
.invoice-form {
  width: 100%;
}

.form-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;
  margin-bottom: 1rem;
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-label {
  margin-bottom: 0.5rem;
  font-weight: 500;
  transition: color 0.3s ease;
}

[data-theme="dark"] .form-label {
  color: #d1d5db;
}

[data-theme="light"] .form-label {
  color: #374151;
}

.required {
  color: #ef4444;
  margin-left: 0.125rem;
}

.form-input,
.form-select,
.form-textarea {
  padding: 0.75rem;
  border: 1px solid;
  border-radius: 0.5rem;
  transition: all 0.3s ease;
  outline: none;
}

[data-theme="dark"] .form-input,
[data-theme="dark"] .form-select,
[data-theme="dark"] .form-textarea {
  background-color: rgba(55, 65, 81, 0.5);
  border-color: #4b5563;
  color: white;
}

[data-theme="light"] .form-input,
[data-theme="light"] .form-select,
[data-theme="light"] .form-textarea {
  background-color: white;
  border-color: #d1d5db;
  color: #111827;
}

.form-input:focus,
.form-select:focus,
.form-textarea:focus {
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.5);
}

.form-textarea {
  resize: vertical;
  min-height: 80px;
}

.form-actions {
  display: flex;
  gap: 1rem;
  justify-content: flex-end;
  margin-top: 1.5rem;
}

.btn {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 0.5rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  display: inline-flex;
  align-items: center;
  justify-content: center;
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.btn-secondary {
  background: transparent;
  border: 1px solid;
}

[data-theme="dark"] .btn-secondary {
  border-color: #4b5563;
  color: #d1d5db;
}

[data-theme="light"] .btn-secondary {
  border-color: #d1d5db;
  color: #4b5563;
}

.btn-secondary:hover:not(:disabled) {
  transform: translateY(-1px);
}

[data-theme="dark"] .btn-secondary:hover:not(:disabled) {
  background-color: #374151;
}

[data-theme="light"] .btn-secondary:hover:not(:disabled) {
  background-color: #f3f4f6;
}

.btn-primary {
  color: white;
}

[data-theme="dark"] .btn-primary {
  background: linear-gradient(135deg, #2563eb, #06b6d4);
}

[data-theme="light"] .btn-primary {
  background: linear-gradient(135deg, #3b82f6, #06b6d4);
}

.btn-primary:hover:not(:disabled) {
  transform: translateY(-1px);
  box-shadow: 0 10px 15px -3px rgba(37, 99, 235, 0.3);
}

[data-theme="dark"] .btn-primary:hover:not(:disabled) {
  background: linear-gradient(135deg, #3b82f6, #22d3ee);
}

[data-theme="light"] .btn-primary:hover:not(:disabled) {
  background: linear-gradient(135deg, #2563eb, #0891b2);
}

.btn-loading {
  display: inline-flex;
  align-items: center;
}

.error-message {
  margin-top: 1rem;
  padding: 0.75rem;
  border-radius: 0.5rem;
  font-size: 0.875rem;
  border: 1px solid;
}

[data-theme="dark"] .error-message {
  background-color: rgba(153, 27, 27, 0.2);
  border-color: #7f1d1d;
  color: #fca5a5;
}

[data-theme="light"] .error-message {
  background-color: #fef2f2;
  border-color: #fecaca;
  color: #dc2626;
}

/* Responsive design */
@media (max-width: 768px) {
  .form-grid {
    grid-template-columns: 1fr;
  }

  .form-actions {
    flex-direction: column;
  }

  .btn {
    width: 100%;
  }
}
</style>