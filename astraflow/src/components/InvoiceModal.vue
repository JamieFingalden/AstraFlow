<template>
  <div
    v-if="visible"
    class="modal-overlay"
    @click="handleOverlayClick"
    :data-theme="isDark ? 'dark' : 'light'"
  >
    <div class="modal-container" ref="modalRef">
      <div class="modal-header">
        <h3 class="modal-title">
          {{ invoice ? '编辑发票' : '新增发票' }}
        </h3>
        <button
          class="modal-close"
          @click="closeModal"
          aria-label="关闭模态框"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="20"
            height="20"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path d="M18 6 6 18" />
            <path d="m6 6 12 12" />
          </svg>
        </button>
      </div>

      <div class="modal-body">
        <InvoiceForm
          :invoice="invoice"
          :loading="loading"
          @submit="handleFormSubmit"
          @cancel="closeModal"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, defineProps, defineEmits, onMounted, onUnmounted } from 'vue'
import InvoiceForm from './InvoiceForm.vue'
import { useTheme } from '../composables/useTheme'
import { onClickOutside } from '@vueuse/core'

// Props
const props = defineProps({
  visible: {
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
const emit = defineEmits(['update:visible', 'submit', 'close'])

// Theme
const { isDark } = useTheme()

// Refs
const modalRef = ref(null)

// Handle overlay click (close modal if clicking outside content)
const handleOverlayClick = (event) => {
  if (event.target === event.currentTarget) {
    closeModal()
  }
}

// Close modal
const closeModal = () => {
  emit('update:visible', false)
  emit('close')
}

// Handle form submission
const handleFormSubmit = (formData) => {
  emit('submit', formData)
}

// Close modal on Escape key
const handleEscapeKey = (event) => {
  if (event.key === 'Escape') {
    closeModal()
  }
}

// Set up event listeners
onMounted(() => {
  document.addEventListener('keydown', handleEscapeKey)
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleEscapeKey)
})

// Use VueUse onClickOutside to handle clicks outside the modal content
onClickOutside(modalRef, () => {
  if (props.visible) {
    closeModal()
  }
})
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  padding: 1rem;
}

.modal-container {
  background: transparent;
  width: 100%;
  max-width: 600px;
  border-radius: 0.75rem;
  overflow: hidden;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
  animation: modalSlideIn 0.3s ease-out;
  backdrop-filter: blur(12px);
  border: 1px solid;
}

[data-theme="dark"] .modal-container {
  background-color: rgba(31, 41, 55, 0.9);
  border-color: rgba(55, 65, 81, 0.5);
}

[data-theme="light"] .modal-container {
  background-color: rgba(255, 255, 255, 0.9);
  border-color: rgba(229, 231, 235, 0.5);
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1.5rem 1.5rem 0 1.5rem;
  border-bottom: 1px solid transparent;
}

[data-theme="dark"] .modal-header {
  border-color: rgba(55, 65, 81, 0.5);
}

[data-theme="light"] .modal-header {
  border-color: rgba(229, 231, 235, 0.5);
}

.modal-title {
  font-size: 1.25rem;
  font-weight: 600;
  margin: 0;
  transition: color 0.3s ease;
}

[data-theme="dark"] .modal-title {
  color: white;
}

[data-theme="light"] .modal-title {
  color: #111827;
}

.modal-close {
  background: none;
  border: none;
  cursor: pointer;
  padding: 0.5rem;
  border-radius: 0.25rem;
  transition: all 0.2s ease;
}

[data-theme="dark"] .modal-close {
  color: #9ca3af;
}

[data-theme="light"] .modal-close {
  color: #6b7280;
}

.modal-close:hover {
  background-color: rgba(0, 0, 0, 0.1);
}

[data-theme="dark"] .modal-close:hover {
  color: white;
  background-color: rgba(55, 65, 81, 0.5);
}

[data-theme="light"] .modal-close:hover {
  color: #111827;
  background-color: #f3f4f6;
}

.modal-body {
  padding: 0 1.5rem 1.5rem 1.5rem;
}

@keyframes modalSlideIn {
  from {
    opacity: 0;
    transform: translateY(-20px) scale(0.95);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

/* Responsive design */
@media (max-width: 768px) {
  .modal-overlay {
    padding: 0.5rem;
  }

  .modal-container {
    max-width: calc(100vw - 1rem);
    margin: 0.5rem;
  }
}
</style>