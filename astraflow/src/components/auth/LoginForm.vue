<template>
  <div class="login-form">
    <div class="form-header">
      <h2 class="form-title">欢迎回来</h2>
      <p class="form-subtitle">登录到您的 AstraFlow 账户</p>
    </div>

    <!-- Email/Username Input -->
    <div class="form-field">
      <label for="username" class="field-label">
        用户名/邮箱
        <span class="required">*</span>
      </label>
      <div class="input-wrapper">
        <div class="input-icon">👤</div>
        <input
          id="username"
          v-model="form.username"
          type="text"
          placeholder="请输入用户名或邮箱"
          :disabled="loading"
          :class="{ 'has-error': errors.username }"
          class="form-input"
          @blur="validateUsername"
          autocomplete="username"
        />
      </div>
      <div v-if="errors.username" class="field-error">
        {{ errors.username }}
      </div>
    </div>

    <!-- Password Input -->
    <div class="form-field">
      <PasswordInput
        v-model="form.password"
        label="密码"
        placeholder="请输入您的密码"
        :required="true"
        :disabled="loading"
        :error-message="errors.password"
        :show-strength="false"
        :show-requirements="false"
        autocomplete="current-password"
        @blur="validatePassword"
      />
    </div>

    <!-- Remember Me & Forgot Password -->
    <div class="form-options">
      <label class="checkbox-wrapper">
        <input
          v-model="form.rememberMe"
          type="checkbox"
          :disabled="loading"
          class="checkbox-input"
        />
        <div class="checkbox-custom">
          <div class="checkbox-check">✓</div>
        </div>
        <span class="checkbox-label">记住我</span>
      </label>

      <button
        type="button"
        @click="handleForgotPassword"
        class="forgot-password-btn"
        :disabled="loading"
      >
        忘记密码？
      </button>
    </div>

    <!-- Login Button -->
    <button
      type="button"
      @click="handleLogin"
      :disabled="loading || !isFormValid"
      class="submit-btn gradient-stellar"
    >
      <div v-if="loading" class="loading-spinner"></div>
      <span v-else>登录</span>
    </button>

    <!-- Demo Account Notice -->
    <div class="demo-notice">
      <div class="notice-title">演示账户</div>
      <div class="notice-content">
        <p>邮箱: demo@astraflow.com</p>
        <p>密码: Demo@123</p>
        <button
          type="button"
          @click="fillDemoCredentials"
          class="demo-fill-btn"
          :disabled="loading"
        >
          填入演示账户
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/userStore'
import PasswordInput from './PasswordInput.vue'
import { login as loginApi } from '@/services/api/authApi'
import { ElMessage } from 'element-plus'

// Remove successMessage and generalError from reactive state since we're using ElMessage

const router = useRouter()
const userStore = useUserStore()

const loading = ref(false)

const form = reactive({
  username: '',
  password: '',
  rememberMe: false
})

const errors = reactive({
  username: '',
  password: ''
})

const isFormValid = computed(() => {
  return form.username &&
         form.password &&
         !errors.username &&
         !errors.password &&
         !loading.value
})

const validateUsername = () => {
  errors.username = ''

  if (!form.username) {
    errors.username = '请输入用户名或邮箱'
    return false
  }
  return true
}

const validatePassword = () => {
  errors.password = ''

  if (!form.password) {
    errors.password = '请输入密码'
    return false
  }

  if (form.password.length < 6) {
    errors.password = '密码长度至少为6位'
    return false
  }
  return true
}

const validateForm = () => {
  const isUsernameValid = validateUsername()
  const isPasswordValid = validatePassword()

  return isUsernameValid && isPasswordValid
}

const handleLogin = async () => {
  if (!validateForm()) {
    return
  }
  loading.value = true

  try {
    const response = await loginApi(form.username, form.password)
    if (!response?.status === 200) {
      throw new Error(response)
    }
    // 成功后更新 store 状态
    await userStore.loginWithResponse(response, form.rememberMe)
    
    ElMessage.success('登录成功！正在跳转...')

    // Redirect to dashboard after successful login
    setTimeout(() => {
      router.push('/visualization')
    }, 500)

  } catch (error) {
    // Handle different error scenarios with ElMessage
    if (error.response?.status === 401) {
      ElMessage.error('用户名或密码错误，请重试')
    } else if (error.response?.status === 429) {
      ElMessage.error('登录尝试次数过多，请稍后再试')
    } else if (error.message.includes('Network')) {
      ElMessage.error('网络连接失败，请检查您的网络连接')
    } else {
      ElMessage.error('登录失败，请稍后重试')
    }
  } finally {
    loading.value = false
  }
}

const handleForgotPassword = () => {
  // TODO: Implement forgot password functionality
  console.log('Forgot password clicked')
  // Could open a modal or navigate to forgot password page
}

const fillDemoCredentials = () => {
  form.username = 'demo@astraflow.com'
  form.password = 'Demo@123'
  form.rememberMe = true

  // Clear any existing errors
  errors.username = ''
  errors.password = ''
}
</script>

<style scoped>
.login-form {
  width: 100%;
  max-width: 400px;
  margin: 0 auto;
}

/* Form Header */
.form-header {
  text-align: center;
  margin-bottom: var(--space-4xl);
}

.form-title {
  font-size: 1.875rem;
  font-weight: 700;
  color: var(--color-starlight);
  margin-bottom: var(--space-sm);
  font-family: var(--font-display);
}

.form-subtitle {
  color: var(--color-starlight);
  opacity: 0.7;
  font-size: 1rem;
  margin: 0;
}

/* Form Fields */
.form-field {
  margin-bottom: var(--space-lg);
}

.field-label {
  display: block;
  margin-bottom: var(--space-sm);
  font-weight: 500;
  color: var(--color-starlight);
  font-size: 0.875rem;
}

.required {
  color: var(--color-aurora-pink);
  margin-left: 2px;
}

.input-wrapper {
  position: relative;
}

.input-icon {
  position: absolute;
  left: var(--space-md);
  top: 50%;
  transform: translateY(-50%);
  font-size: 1rem;
  opacity: 0.7;
  z-index: 1;
}

.form-input {
  width: 100%;
  padding: var(--space-md) var(--space-md) var(--space-md) var(--space-2xl);
  border: 1px solid var(--glass-border);
  border-radius: var(--radius-md);
  background: rgba(255, 255, 255, 0.05);
  color: var(--color-starlight);
  font-size: 1rem;
  transition: all var(--duration-normal) ease;
  outline: none;
}

.form-input:focus {
  border-color: var(--color-stellar-blue);
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

.form-input.has-error {
  border-color: var(--color-aurora-pink);
  box-shadow: 0 0 0 3px rgba(236, 72, 153, 0.1);
}

.form-input:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.field-error {
  display: flex;
  align-items: center;
  gap: var(--space-xs);
  margin-top: var(--space-xs);
  color: var(--color-aurora-pink);
  font-size: 0.875rem;
}

/* Form Options */
.form-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--space-xl);
  flex-wrap: wrap;
  gap: var(--space-md);
}

.checkbox-wrapper {
  display: flex;
  align-items: center;
  gap: var(--space-sm);
  cursor: pointer;
  user-select: none;
}

.checkbox-input {
  display: none;
}

.checkbox-custom {
  width: 20px;
  height: 20px;
  border: 2px solid var(--glass-border);
  border-radius: var(--radius-sm);
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all var(--duration-normal) ease;
  background: rgba(255, 255, 255, 0.05);
}

.checkbox-input:checked + .checkbox-custom {
  background: var(--gradient-stellar);
  border-color: var(--color-stellar-blue);
}

.checkbox-check {
  color: white;
  font-size: 0.75rem;
  font-weight: bold;
  opacity: 0;
  transform: scale(0);
  transition: all var(--duration-normal) ease;
}

.checkbox-input:checked + .checkbox-custom .checkbox-check {
  opacity: 1;
  transform: scale(1);
}

.checkbox-label {
  color: var(--color-starlight);
  font-size: 0.875rem;
  opacity: 0.8;
}

.forgot-password-btn {
  background: none;
  border: none;
  color: var(--color-stellar-blue);
  font-size: 0.875rem;
  cursor: pointer;
  text-decoration: underline;
  transition: all var(--duration-normal) ease;
}

.forgot-password-btn:hover:not(:disabled) {
  color: var(--color-nebula-purple);
  text-shadow: 0 0 5px var(--color-nebula-glow);
}

.forgot-password-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

/* Submit Button */
.submit-btn {
  width: 100%;
  padding: var(--space-md) var(--space-lg);
  border: none;
  border-radius: var(--radius-md);
  color: white;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all var(--duration-normal) ease;
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 48px;
  position: relative;
  overflow: hidden;
}

.submit-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 10px 25px rgba(59, 130, 246, 0.3);
}

.submit-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
  transform: none;
}

.loading-spinner {
  width: 20px;
  height: 20px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-top: 2px solid white;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

/* Messages */
.success-message {
  display: flex;
  align-items: center;
  gap: var(--space-sm);
  margin-top: var(--space-lg);
  padding: var(--space-md);
  background: rgba(34, 197, 94, 0.1);
  border: 1px solid rgba(34, 197, 94, 0.3);
  border-radius: var(--radius-md);
  color: #22c55e;
  font-size: 0.875rem;
}

.success-icon {
  font-size: 1rem;
}

.general-error {
  display: flex;
  align-items: center;
  gap: var(--space-sm);
  margin-top: var(--space-lg);
  padding: var(--space-md);
  background: rgba(236, 72, 153, 0.1);
  border: 1px solid rgba(236, 72, 153, 0.3);
  border-radius: var(--radius-md);
  color: var(--color-aurora-pink);
  font-size: 0.875rem;
}

.error-icon {
  font-size: 1rem;
}

/* Demo Notice */
.demo-notice {
  margin-top: var(--space-2xl);
  padding: var(--space-lg);
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: var(--radius-md);
  text-align: center;
}

.notice-title {
  font-weight: 600;
  color: var(--color-starlight);
  margin-bottom: var(--space-sm);
  font-size: 0.875rem;
}

.notice-content p {
  color: var(--color-starlight);
  opacity: 0.7;
  font-size: 0.875rem;
  margin: var(--space-xs) 0;
}

.demo-fill-btn {
  margin-top: var(--space-sm);
  padding: var(--space-xs) var(--space-md);
  background: rgba(59, 130, 246, 0.2);
  border: 1px solid var(--color-stellar-blue);
  border-radius: var(--radius-sm);
  color: var(--color-stellar-blue);
  font-size: 0.875rem;
  cursor: pointer;
  transition: all var(--duration-normal) ease;
}

.demo-fill-btn:hover:not(:disabled) {
  background: rgba(59, 130, 246, 0.3);
  transform: translateY(-1px);
}

.demo-fill-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

/* Animations */
@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

/* Light theme adjustments */
[data-theme="light"] .form-input {
  background: rgba(0, 0, 0, 0.05);
  border-color: rgba(0, 0, 0, 0.2);
  color: var(--color-text-primary);
}

[data-theme="light"] .form-input:focus {
  border-color: var(--color-stellar-blue);
  box-shadow: 0 0 0 3px rgba(37, 99, 235, 0.1);
}

[data-theme="light"] .form-input.has-error {
  border-color: var(--color-aurora-pink);
  box-shadow: 0 0 0 3px rgba(236, 72, 153, 0.1);
}

[data-theme="light"] .field-label {
  color: var(--color-text-primary);
}

[data-theme="light"] .checkbox-custom {
  background: rgba(0, 0, 0, 0.05);
  border-color: rgba(0, 0, 0, 0.2);
}

[data-theme="light"] .checkbox-label {
  color: var(--color-text-secondary);
}

[data-theme="light"] .forgot-password-btn {
  color: var(--color-stellar-blue);
}

[data-theme="light"] .forgot-password-btn:hover:not(:disabled) {
  color: var(--color-nebula-purple);
  text-shadow: 0 0 5px rgba(124, 58, 237, 0.4);
}

[data-theme="light"] .success-message {
  background: rgba(34, 197, 94, 0.1);
  border-color: rgba(34, 197, 94, 0.3);
  color: #16a34a;
}

[data-theme="light"] .general-error {
  background: rgba(239, 68, 68, 0.1);
  border-color: rgba(239, 68, 68, 0.3);
  color: #dc2626;
}

[data-theme="light"] .demo-notice {
  background: rgba(0, 0, 0, 0.05);
  border-color: rgba(0, 0, 0, 0.1);
}

[data-theme="light"] .demo-fill-btn {
  background: rgba(59, 130, 246, 0.1);
  border-color: var(--color-stellar-blue);
  color: var(--color-stellar-blue);
}

[data-theme="light"] .demo-fill-btn:hover:not(:disabled) {
  background: rgba(59, 130, 246, 0.2);
}

/* Form input focus and selection */
.form-input::selection {
  background: var(--color-stellar-blue);
  color: white;
}

.form-input::-moz-selection {
  background: var(--color-stellar-blue);
  color: white;
}

/* Light theme text selection in form */
[data-theme="light"] .form-input::selection {
  background: var(--color-stellar-blue);
  color: white;
}

[data-theme="light"] .form-input::-moz-selection {
  background: var(--color-stellar-blue);
  color: white;
}

/* Responsive Design */
@media (max-width: 480px) {
  .form-options {
    flex-direction: column;
    align-items: flex-start;
    gap: var(--space-sm);
  }

  .form-title {
    font-size: 1.5rem;
  }

  .form-input {
    font-size: 16px; /* Prevents zoom on iOS */
  }
}
</style>