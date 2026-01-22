<template>
  <div class="invoice-form">
    <n-form
      :model="formData"
      :rules="rules"
      ref="formRef"
      :label-placement="labelPlacement"
      :label-width="100"
      @submit.prevent
    >
      <n-grid :cols="2" :x-gap="16" :y-gap="16">
        <!-- 发票号码 -->
        <n-form-item-gi :span="1" label="发票号码" path="invoiceNumber">
          <n-input
            v-model:value="formData.invoiceNumber"
            placeholder="请输入发票号码"
            :input-props="{ id: 'invoiceNumber' }"
          />
        </n-form-item-gi>

        <!-- 发票日期 -->
        <n-form-item-gi :span="1" label="发票日期" path="date">
          <n-date-picker
            v-model:value="formData.date"
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
        <n-form-item-gi :span="1" label="支付来源" path="paymentSource">
          <n-select
            v-model:value="formData.paymentSource"
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
        <n-form-item-gi :span="1" label="发票图片URL" path="imageUrl">
          <n-input
            v-model:value="formData.imageUrl"
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
        <n-button @click="onCancel" :disabled="loading" @keydown.enter.prevent>
          取消
        </n-button>
        <n-button
          type="primary"
          @click="submitForm"
          :loading="loading"
          @keydown.enter.prevent
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
  </div>
</template>

<script setup>
import { ref, watch, defineProps, defineEmits, computed } from 'vue'
import { useTheme } from '../composables/useTheme'
import { NForm, NFormItemGi, NGrid, NInput, NDatePicker, NInputNumber, NSelect, NButton, NSpace, NAlert } from 'naive-ui'

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
  date: null,
  amount: null,
  vendor: '',
  category: '',
  taxId: '',
  paymentSource: '',
  status: 'pending',
  imageUrl: '',
  description: ''
})

// Form reference for validation
const formRef = ref()

// Check if in edit mode
const isEditMode = computed(() => !!props.invoice)

// Form rules for validation
const rules = {
  invoiceNumber: {
    required: true,
    message: '请输入发票号码',
    trigger: 'blur'
  },
  date: {
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
  { label: '通过', value: 'approved' },
  { label: '已报销', value: 'paid' },
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
      if (newInvoice.date) {
        dateValue = new Date(newInvoice.date).getTime()
      } else if (newInvoice.invoice_date) {
        dateValue = new Date(newInvoice.invoice_date).getTime()
      }

      formData.value = {
        invoiceNumber: newInvoice.invoiceNumber || newInvoice.invoice_number || newInvoice.title || '',
        date: dateValue,
        amount: newInvoice.amount || null,
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
        date: null,
        amount: null,
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

// Handle form submission
const submitForm = async () => {
  try {
    await formRef.value.validate()
    error.value = ''

    // Format data for submission with backend-compatible field names
    const submitData = {
      invoice_number: formData.value.invoiceNumber,
      invoice_date: formData.value.date ? new Date(formData.value.date).toISOString() : null,
      amount: formData.value.amount,
      vendor: formData.value.vendor,
      category: formData.value.category,
      taxId: formData.value.taxId,
      payment_source: formData.value.paymentSource,
      status: formData.value.status,
      image_url: formData.value.imageUrl || null,
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