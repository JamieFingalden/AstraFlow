<template>
  <n-modal
    v-model:show="showModal"
    :mask-closable="false"
    preset="card"
    :style="{ width: '600px', maxWidth: '90vw' }"
    :title="isEditMode ? '编辑发票' : '添加发票'"
    :bordered="false"
    size="huge"
    role="dialog"
    aria-modal="true"
  >
    <invoice-form-naive
      :invoice="editingInvoice"
      :loading="loading"
      @submit="handleFormSubmit"
      @cancel="closeModal"
    />
  </n-modal>
</template>

<script setup>
import { ref, watch, computed } from 'vue'
import InvoiceFormNaive from './InvoiceFormNaive.vue'

// Props
const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
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
const emit = defineEmits(['update:modelValue', 'submit', 'cancel'])

// Local state
const showModal = ref(false)
const editingInvoice = ref(null)

// Check if in edit mode
const isEditMode = computed(() => !!editingInvoice.value)

// Watch for changes to the modelValue prop
watch(
  () => props.modelValue,
  (newValue) => {
    showModal.value = newValue
  }
)

// Watch for changes to the invoice prop
watch(
  () => props.invoice,
  (newInvoice) => {
    editingInvoice.value = newInvoice
  }
)

// Watch for changes to the local showModal state
watch(
  () => showModal.value,
  (newValue) => {
    if (!newValue) {
      editingInvoice.value = null
    }
    emit('update:modelValue', newValue)
  }
)

// Handle form submission
const handleFormSubmit = (formData) => {
  emit('submit', formData)
}

// Close modal
const closeModal = () => {
  emit('update:modelValue', false)
}
</script>

<style scoped>
/* No additional styles needed for Naive UI modal */
</style>