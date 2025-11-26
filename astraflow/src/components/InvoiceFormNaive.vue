<template>
  <n-form
    :model="formData"
    :rules="rules"
    ref="formRef"
    class="invoice-form"
    :label-placement="labelPlacement"
    :label-width="100"
  >
    <n-grid :cols="2" :x-gap="16" :y-gap="16">
      <!-- 发票号码 -->
      <n-form-item-gi :span="1" label="发票号码" path="invoice_number">
        <n-input
          v-model:value="formData.invoice_number"
          placeholder="请输入发票号码"
          :input-props="{ id: 'invoiceNumber' }"
        />
      </n-form-item-gi>

      <!-- 发票日期 -->
      <n-form-item-gi :span="1" label="发票日期" path="invoice_date">
        <n-date-picker
          v-model:value="formData.invoice_date"
          type="date"
          placeholder="请选择发票日期"
          :input-props="{ id: 'invoiceDate' }"
          style="width: 100%"
        />
      </n-form-item-gi>

      <!-- 金额 -->
      <n-form-item-gi :span="1" label="金额" path="amount">
        <n-input-number
          v-model:value="formData.amount"
          placeholder="0.00"
          :precision="2"
          :min="0"
          :max="999999999"
          :input-props="{ id: 'amount' }"
          style="width: 100%"
        />
      </n-form-item-gi>

      <!-- 供应商 -->
      <n-form-item-gi :span="1" label="供应商/商户" path="vendor">
        <n-input
          v-model:value="formData.vendor"
          placeholder="请输入供应商名称"
          :input-props="{ id: 'vendor' }"
        />
      </n-form-item-gi>

      <!-- 分类 -->
      <n-form-item-gi :span="1" label="分类" path="category">
        <n-select
          v-model:value="formData.category"
          :options="categoryOptions"
          placeholder="请选择分类"
          :input-props="{ id: 'category' }"
        />
      </n-form-item-gi>

      <!-- 税号 -->
      <n-form-item-gi :span="1" label="税号" path="taxId">
        <n-input
          v-model:value="formData.taxId"
          placeholder="请输入税号"
          :input-props="{ id: 'taxId' }"
        />
      </n-form-item-gi>

      <!-- 支付来源 -->
      <n-form-item-gi :span="1" label="支付来源" path="payment_source">
        <n-select
          v-model:value="formData.payment_source"
          :options="paymentSourceOptions"
          placeholder="请选择支付来源"
          :input-props="{ id: 'paymentSource' }"
        />
      </n-form-item-gi>

      <!-- 状态 -->
      <n-form-item-gi :span="1" label="状态" path="status">
        <n-select
          v-model:value="formData.status"
          :options="statusOptions"
          placeholder="请选择状态"
          :input-props="{ id: 'status' }"
        />
      </n-form-item-gi>

      <!-- 图片URL -->
      <n-form-item-gi :span="1" label="发票图片URL" path="image_url">
        <n-input
          v-model:value="formData.image_url"
          placeholder="请输入发票图片链接"
          type="textarea"
          :input-props="{ id: 'imageUrl' }"
        />
      </n-form-item-gi>
    </n-grid>

    <!-- 描述 -->
    <n-form-item label="描述" path="description">
      <n-input
        v-model:value="formData.description"
        placeholder="请输入发票描述或备注"
        type="textarea"
        :autosize="{ minRows: 3, maxRows: 5 }"
        :input-props="{ id: 'description' }"
      />
    </n-form-item>

    <!-- 表单操作按钮 -->
    <n-space justify="end" class="form-actions">
      <n-button @click="onCancel" :disabled="loading">
        取消
      </n-button>
      <n-button
        type="primary"
        @click="submitForm"
        :loading="loading"
      >
        {{ isEditMode ? '更新发票' : '创建发票' }}
      </n-button>
    </n-space>
  </n-form>

  <!-- 错误提示 -->
  <n-alert
    v-if="error"
    :content="error"
    type="error"
    :style="{ marginTop: '16px' }"
  />
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
  invoice_number: '',
  invoice_date: null,
  amount: null,
  vendor: '',
  category: '',
  taxId: '',
  payment_source: '',
  status: 'pending',
  image_url: '',
  description: ''
})

// Form reference for validation
const formRef = ref()

// Check if in edit mode
const isEditMode = computed(() => !!props.invoice)

// Form rules for validation
const rules = {
  invoice_number: {
    required: true,
    message: '请输入发票号码',
    trigger: 'blur'
  },
  invoice_date: {
    type: 'number',
    required: true,
    message: '请选择发票日期',
    trigger: 'change'
  },
  amount: {
    type: 'number',
    required: true,
    message: '请输入金额',
    trigger: 'blur'
  },
  vendor: {
    required: true,
    message: '请输入供应商/商户名称',
    trigger: 'blur'
  },
  category: {
    required: true,
    message: '请选择分类',
    trigger: 'change'
  }
}

// Options for selects
const categoryOptions = [
  { label: '餐饮', value: '餐饮' },
  { label: '交通', value: '交通' },
  { label: '住宿', value: '住宿' },
  { label: '办公', value: '办公' },
  { label: '其他', value: '其他' }
]

const paymentSourceOptions = [
  { label: '微信支付', value: '微信支付' },
  { label: '支付宝', value: '支付宝' },
  { label: '银行卡', value: '银行卡' },
  { label: '手动添加', value: '手动添加' },
  { label: '现金', value: '现金' }
]

const statusOptions = [
  { label: '待处理', value: 'pending' },
  { label: '已确认', value: 'confirmed' },
  { label: '已拒绝', value: 'rejected' }
]

// Label placement
const labelPlacement = ref('top')

// Error state
const error = ref('')

// Initialize form data when invoice prop changes
watch(
  () => props.invoice,
  (newInvoice) => {
    if (newInvoice) {
      // Format date from ISO string to timestamp for Naive UI date picker
      let dateValue = null
      if (newInvoice.invoice_date) {
        dateValue = new Date(newInvoice.invoice_date).getTime()
      } else if (newInvoice.date) {
        dateValue = new Date(newInvoice.date).getTime()
      }

      formData.value = {
        invoice_number: newInvoice.invoiceNumber || newInvoice.invoice_number || newInvoice.title || '',
        invoice_date: dateValue,
        amount: newInvoice.amount || null,
        vendor: newInvoice.vendor || '',
        category: newInvoice.category || '',
        taxId: newInvoice.taxId || newInvoice.tax_id || '',
        payment_source: newInvoice.paymentSource || newInvoice.payment_source || '',
        status: newInvoice.status || 'pending',
        image_url: newInvoice.imageUrl || newInvoice.image_url || newInvoice.imageURL || '',
        description: newInvoice.description || ''
      }
    } else {
      // Reset form for new invoice
      formData.value = {
        invoice_number: '',
        invoice_date: null,
        amount: null,
        vendor: '',
        category: '',
        taxId: '',
        payment_source: '',
        status: 'pending',
        image_url: '',
        description: ''
      }
    }
  },
  { immediate: true }
)

// Handle form submission
const submitForm = async () => {
  try {
    await formRef.value.validate()
    error.value = ''

    // Format data for submission with backend-compatible field names
    const submitData = {
      invoice_number: formData.value.invoice_number,
      invoice_date: formData.value.invoice_date ? new Date(formData.value.invoice_date).toISOString() : null,
      amount: formData.value.amount,
      vendor: formData.value.vendor,
      category: formData.value.category,
      taxId: formData.value.taxId,
      payment_source: formData.value.payment_source,
      status: formData.value.status,
      image_url: formData.value.image_url || null,
      description: formData.value.description || null
    }

    emit('submit', submitData)
  } catch (error) {
    console.error('Form validation failed:', error)
    error.value = '请检查表单信息并重新提交'
  }
}

// Handle cancel
const onCancel = () => {
  emit('cancel')
}
</script>

<style scoped>
.invoice-form {
  padding: 20px;
}

.form-actions {
  margin-top: 24px;
  width: 100%;
}
</style>