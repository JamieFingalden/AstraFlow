<template>
  <div class="password-input-wrapper">
    <label v-if="label" :for="inputId" class="input-label">
      {{ label }}
      <span v-if="required" class="required-indicator">*</span>
    </label>

    <div class="password-input-container" :class="{ 'has-error': errorMessage }">
      <input
        :id="inputId"
        :type="showPassword ? 'text' : 'password'"
        :value="modelValue"
        :placeholder="placeholder"
        :disabled="disabled"
        :required="required"
        :autocomplete="autocomplete"
        class="password-input"
        @input="handleInput"
        @blur="handleBlur"
        @focus="handleFocus"
      />

      <button
        type="button"
        class="password-toggle-btn"
        :disabled="disabled"
        @click="togglePasswordVisibility"
        :title="showPassword ? '隐藏密码' : '显示密码'"
      >
        <div class="toggle-icon">
          <div v-if="showPassword" class="eye-open">👁️</div>
          <div v-else class="eye-closed">👁️‍🗨️</div>
        </div>
      </button>
    </div>

    <!-- Password strength indicator -->
    <div v-if="showStrength && modelValue" class="password-strength">
      <div class="strength-bar">
        <div
          class="strength-fill"
          :class="strengthClass"
          :style="{ width: `${strengthPercentage}%` }"
        ></div>
      </div>
      <span class="strength-text" :class="strengthClass">
        {{ strengthText }}
      </span>
    </div>

    <!-- Error message -->
    <div v-if="errorMessage" class="error-message">
      <div class="error-icon">⚠️</div>
      <span>{{ errorMessage }}</span>
    </div>

    <!-- Password requirements hint -->
    <div v-if="showRequirements && focused" class="password-requirements">
      <div class="requirements-title">密码要求：</div>
      <ul class="requirements-list">
        <li
          v-for="requirement in passwordRequirements"
          :key="requirement.key"
          class="requirement-item"
          :class="{ 'met': requirement.met }"
        >
          <div class="requirement-icon">
            <div v-if="requirement.met" class="check-icon">✓</div>
            <div v-else class="circle-icon">○</div>
          </div>
          <span>{{ requirement.text }}</span>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  },
  label: {
    type: String,
    default: ''
  },
  placeholder: {
    type: String,
    default: '请输入密码'
  },
  required: {
    type: Boolean,
    default: false
  },
  disabled: {
    type: Boolean,
    default: false
  },
  autocomplete: {
    type: String,
    default: 'current-password'
  },
  showStrength: {
    type: Boolean,
    default: true
  },
  showRequirements: {
    type: Boolean,
    default: false
  },
  errorMessage: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update:modelValue', 'blur', 'focus'])

const showPassword = ref(false)
const focused = ref(false)
const inputId = computed(() => `password-${Math.random().toString(36).substr(2, 9)}`)

const togglePasswordVisibility = () => {
  showPassword.value = !showPassword.value
}

const handleInput = (event) => {
  emit('update:modelValue', event.target.value)
}

const handleBlur = (event) => {
  focused.value = false
  emit('blur', event)
}

const handleFocus = (event) => {
  focused.value = true
  emit('focus', event)
}

// Password validation requirements
const passwordRequirements = computed(() => [
  {
    key: 'length',
    text: '至少8个字符',
    met: props.modelValue.length >= 8
  },
  {
    key: 'uppercase',
    text: '包含大写字母',
    met: /[A-Z]/.test(props.modelValue)
  },
  {
    key: 'lowercase',
    text: '包含小写字母',
    met: /[a-z]/.test(props.modelValue)
  },
  {
    key: 'number',
    text: '包含数字',
    met: /\d/.test(props.modelValue)
  },
  {
    key: 'special',
    text: '包含特殊字符 (!@#$%^&*)',
    met: /[!@#$%^&*(),.?":{}|<>]/.test(props.modelValue)
  }
])

// Password strength calculation
const strengthScore = computed(() => {
  let score = 0
  const password = props.modelValue

  if (password.length >= 8) score += 1
  if (password.length >= 12) score += 1
  if (/[A-Z]/.test(password)) score += 1
  if (/[a-z]/.test(password)) score += 1
  if (/\d/.test(password)) score += 1
  if (/[!@#$%^&*(),.?":{}|<>]/.test(password)) score += 1

  return score
})

const strengthPercentage = computed(() => {
  return Math.min((strengthScore.value / 6) * 100, 100)
})

const strengthClass = computed(() => {
  if (strengthScore.value <= 2) return 'weak'
  if (strengthScore.value <= 4) return 'medium'
  return 'strong'
})

const strengthText = computed(() => {
  if (strengthScore.value <= 2) return '弱'
  if (strengthScore.value <= 4) return '中等'
  return '强'
})
</script>

<style scoped>
.password-input-wrapper {
  width: 100%;
}

.input-label {
  display: block;
  margin-bottom: var(--space-sm);
  font-weight: 500;
  color: var(--color-starlight);
  font-size: 0.875rem;
}

.required-indicator {
  color: var(--color-aurora-pink);
  margin-left: 2px;
}

.password-input-container {
  position: relative;
  display: flex;
  align-items: center;
}

.password-input {
  width: 100%;
  padding: var(--space-md) var(--space-2xl) var(--space-md) var(--space-md);
  border: 1px solid var(--glass-border);
  border-radius: var(--radius-md);
  background: rgba(255, 255, 255, 0.05);
  color: var(--color-starlight);
  font-size: 1rem;
  transition: all var(--duration-normal) ease;
  outline: none;
}

.password-input:focus {
  border-color: var(--color-stellar-blue);
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

.password-input:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.password-input-container.has-error .password-input {
  border-color: var(--color-aurora-pink);
  box-shadow: 0 0 0 3px rgba(236, 72, 153, 0.1);
}

.password-toggle-btn {
  position: absolute;
  right: var(--space-sm);
  background: none;
  border: none;
  cursor: pointer;
  padding: var(--space-sm);
  border-radius: var(--radius-sm);
  transition: all var(--duration-normal) ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.password-toggle-btn:hover:not(:disabled) {
  background: rgba(255, 255, 255, 0.1);
}

.password-toggle-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.toggle-icon {
  font-size: 1.1rem;
  opacity: 0.7;
  transition: opacity var(--duration-normal) ease;
}

.password-toggle-btn:hover:not(:disabled) .toggle-icon {
  opacity: 1;
}

/* Password Strength Indicator */
.password-strength {
  margin-top: var(--space-sm);
  display: flex;
  align-items: center;
  gap: var(--space-sm);
}

.strength-bar {
  flex: 1;
  height: 4px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 2px;
  overflow: hidden;
}

.strength-fill {
  height: 100%;
  transition: all var(--duration-normal) ease;
  border-radius: 2px;
}

.strength-fill.weak {
  background: var(--color-aurora-pink);
}

.strength-fill.medium {
  background: var(--color-cosmic-cyan);
}

.strength-fill.strong {
  background: var(--color-stellar-blue);
}

.strength-text {
  font-size: 0.75rem;
  font-weight: 500;
  min-width: 30px;
}

.strength-text.weak {
  color: var(--color-aurora-pink);
}

.strength-text.medium {
  color: var(--color-cosmic-cyan);
}

.strength-text.strong {
  color: var(--color-stellar-blue);
}

/* Error Message */
.error-message {
  display: flex;
  align-items: center;
  gap: var(--space-xs);
  margin-top: var(--space-sm);
  color: var(--color-aurora-pink);
  font-size: 0.875rem;
}

.error-icon {
  font-size: 0.875rem;
}

/* Password Requirements */
.password-requirements {
  margin-top: var(--space-md);
  padding: var(--space-md);
  background: rgba(255, 255, 255, 0.05);
  border-radius: var(--radius-md);
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.requirements-title {
  font-size: 0.875rem;
  font-weight: 500;
  color: var(--color-starlight);
  margin-bottom: var(--space-sm);
}

.requirements-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.requirement-item {
  display: flex;
  align-items: center;
  gap: var(--space-sm);
  font-size: 0.875rem;
  color: var(--color-starlight);
  opacity: 0.7;
  transition: all var(--duration-normal) ease;
  margin-bottom: var(--space-xs);
}

.requirement-item:last-child {
  margin-bottom: 0;
}

.requirement-item.met {
  opacity: 1;
  color: var(--color-stellar-blue);
}

.requirement-icon {
  font-size: 0.875rem;
  width: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.check-icon {
  color: var(--color-stellar-blue);
}

.circle-icon {
  color: var(--color-starlight);
  opacity: 0.5;
}

/* Light theme adjustments */
[data-theme="light"] .password-input {
  background: var(--color-form-bg);
  border-color: var(--color-form-border);
  color: var(--color-text-primary);
}

[data-theme="light"] .password-input:focus {
  border-color: var(--color-form-focus-border);
  box-shadow: 0 0 0 3px var(--color-form-focus-shadow);
}

[data-theme="light"] .password-input-container.has-error .password-input {
  border-color: var(--color-form-error-border);
  box-shadow: 0 0 0 3px var(--color-form-error-shadow);
}

[data-theme="light"] .password-requirements {
  background: var(--color-form-bg);
  border: 1px solid var(--color-form-border);
}

/* Password input text selection */
.password-input::selection {
  background: var(--color-stellar-blue);
  color: white;
}

.password-input::-moz-selection {
  background: var(--color-stellar-blue);
  color: white;
}

/* Light theme text selection in password input */
[data-theme="light"] .password-input::selection {
  background: var(--color-stellar-blue);
  color: white;
}

[data-theme="light"] .password-input::-moz-selection {
  background: var(--color-stellar-blue);
  color: white;
}

/* Responsive Design */
@media (max-width: 480px) {
  .password-input {
    font-size: 16px; /* Prevents zoom on iOS */
  }
}
</style>