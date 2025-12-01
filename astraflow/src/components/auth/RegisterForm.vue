<template>
  <form @submit.prevent="handleSubmit" class="register-form">
    <div class="form-header">
      <h2 class="form-title">创建企业账户</h2>
      <p class="form-subtitle">加入 AstraFlow，开启企业智能财务管理之旅</p>
    </div>

    <!-- Tenant Name (required for all users now) -->
    <div class="form-field">
      <label for="tenantName" class="field-label">
        企业名称
        <span class="required">*</span>
      </label>
      <div class="input-wrapper">
        <div class="input-icon">🏢</div>
        <input
          id="tenantName"
          v-model="form.tenantName"
          type="text"
          placeholder="请输入您的企业名称"
          :disabled="loading"
          :class="{ 'has-error': errors.tenantName }"
          class="form-input"
          @blur="validateTenantName"
          autocomplete="organization"
        />
      </div>
      <div v-if="errors.tenantName" class="field-error">
        {{ errors.tenantName }}
      </div>
    </div>

    <!-- Username -->
    <div class="form-field">
      <label for="username" class="field-label">
        用户名
        <span class="required">*</span>
      </label>
      <div class="input-wrapper">
        <div class="input-icon">👤</div>
        <input
          id="username"
          v-model="form.username"
          type="text"
          placeholder="请输入用户名"
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

    <!-- Email -->
    <div class="form-field">
      <label for="email" class="field-label">
        邮箱地址
        <span class="required">*</span>
      </label>
      <div class="input-wrapper">
        <div class="input-icon">📧</div>
        <input
          id="email"
          v-model="form.email"
          type="email"
          placeholder="请输入您的邮箱"
          :disabled="loading"
          :class="{ 'has-error': errors.email }"
          class="form-input"
          @blur="validateEmail"
          autocomplete="email"
        />
      </div>
      <div v-if="errors.email" class="field-error">
        {{ errors.email }}
      </div>
    </div>

    <!-- Password -->
    <div class="form-field">
      <PasswordInput
        v-model="form.password"
        label="密码"
        placeholder="请创建密码"
        :required="true"
        :disabled="loading"
        :error-message="errors.password"
        :show-strength="true"
        :show-requirements="true"
        autocomplete="new-password"
        @blur="validatePassword"
      />
    </div>

    <!-- Confirm Password -->
    <div class="form-field">
      <PasswordInput
        v-model="form.confirmPassword"
        label="确认密码"
        placeholder="请再次输入密码"
        :required="true"
        :disabled="loading"
        :error-message="errors.confirmPassword"
        :show-strength="false"
        :show-requirements="false"
        autocomplete="new-password"
        @blur="validateConfirmPassword"
      />
    </div>

    <!-- Terms Agreement -->
    <div class="form-field">
      <label class="checkbox-wrapper">
        <input
          v-model="form.agreeToTerms"
          type="checkbox"
          :disabled="loading"
          class="checkbox-input"
        />
        <div class="checkbox-custom">
          <div class="checkbox-check">✓</div>
        </div>
        <span class="checkbox-label">
          我已阅读并同意
          <a href="#" class="terms-link" @click.prevent="showTerms">《服务条款》</a>
          和
          <a href="#" class="terms-link" @click.prevent="showPrivacy">《隐私政策》</a>
        </span>
      </label>
      <div v-if="errors.agreeToTerms" class="field-error">
        {{ errors.agreeToTerms }}
      </div>
    </div>

    <!-- Submit Button -->
    <button
      type="submit"
      :disabled="loading || !isFormValid"
      class="submit-btn gradient-stellar"
    >
      <div v-if="loading" class="loading-spinner"></div>
      <span v-else>创建企业账户</span>
    </button>

    <!-- Login Link -->
    <div class="login-link">
      <span>已有账户？</span>
      <button
        type="button"
        @click="$emit('switchToLogin')"
        class="switch-form-btn"
        :disabled="loading"
      >
        立即登录
      </button>
    </div>
  </form>
</template>

<script setup>
import { ref, computed, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/userStore'
import { register as registerApi } from '@/services/api/authApi'
import { getRoleDisplayName, getTenantRoleFromRole } from '@/utils/userUtils' // 假设我们创建一个工具函数
import PasswordInput from './PasswordInput.vue'
import { ElMessage } from 'element-plus'

const emit = defineEmits(['switchToLogin'])

const router = useRouter()
const userStore = useUserStore()

const loading = ref(false)
const successMessage = ref('')

const form = reactive({
  tenantName: '',
  username: '',
  email: '',
  password: '',
  confirmPassword: '',
  agreeToTerms: false
})

const errors = reactive({
  tenantName: '',
  username: '',
  email: '',
  password: '',
  confirmPassword: '',
  agreeToTerms: ''
})

const isFormValid = computed(() => {
  return form.username &&
         form.email &&
         form.password &&
         form.confirmPassword &&
         form.tenantName && // Now tenantName is required
         form.agreeToTerms &&
         !Object.values(errors).some(error => error) &&
         !loading.value
})

const validateTenantName = () => {
  errors.tenantName = ''

  if (!form.tenantName) {
    errors.tenantName = '请输入企业名称'
    return false
  }

  if (form.tenantName && form.tenantName.length < 2) {
    errors.tenantName = '企业名称至少需要2个字符'
    return false
  }

  return true
}

const validateUsername = () => {
  errors.username = ''

  if (!form.username) {
    errors.username = '请输入用户名'
    return false
  }

  if (form.username.length < 3) {
    errors.username = '用户名至少需要3个字符'
    return false
  }

  if (form.username.length > 20) {
    errors.username = '用户名不能超过20个字符'
    return false
  }

  if (!/^[a-zA-Z0-9_\u4e00-\u9fa5]+$/.test(form.username)) {
    errors.username = '用户名只能包含字母、数字、下划线和中文'
    return false
  }

  return true
}

const validateEmail = () => {
  errors.email = ''

  if (!form.email) {
    errors.email = '请输入邮箱地址'
    return false
  }

  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  if (!emailRegex.test(form.email)) {
    errors.email = '请输入有效的邮箱地址'
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

  if (form.password.length < 8) {
    errors.password = '密码长度至少为8位'
    return false
  }

  // Check for at least 3 of the 4 requirements
  let score = 0
  if (/[A-Z]/.test(form.password)) score++
  if (/[a-z]/.test(form.password)) score++
  if (/\d/.test(form.password)) score++
  if (/[!@#$%^&*(),.?":{}|<>]/.test(form.password)) score++

  if (score < 3) {
    errors.password = '密码强度不足，请满足至少3项密码要求'
    return false
  }

  return true
}

const validateConfirmPassword = () => {
  errors.confirmPassword = ''

  if (!form.confirmPassword) {
    errors.confirmPassword = '请确认密码'
    return false
  }

  if (form.password !== form.confirmPassword) {
    errors.confirmPassword = '两次输入的密码不一致'
    return false
  }

  return true
}

const validateTerms = () => {
  errors.agreeToTerms = ''

  if (!form.agreeToTerms) {
    errors.agreeToTerms = '请同意服务条款和隐私政策'
    return false
  }

  return true
}

const validateForm = () => {
  const validations = [
    validateTenantName(),
    validateUsername(),
    validateEmail(),
    validatePassword(),
    validateConfirmPassword(),
    validateTerms()
  ]

  return validations.every(valid => valid)
}

const handleSubmit = async () => {
  if (!validateForm()) {
    return
  }

  loading.value = true
  successMessage.value = ''

  try {
    let userData = {
      username: form.username,
      email: form.email,
      password: form.password,
      phone: '', // 前端暂无电话输入
      tenant_name: form.tenantName // Add tenant name to registration data
    }

    // Now all users will be registered as tenant users
    // The backend should create a new tenant for this user
    userData.tenant_id = null // This will be set by the backend when creating the new tenant

    // 直接调用API
    const response = await registerApi(userData)
    const { user: backendUser, token: access_token, refresh_token } = response.data

    // 转换后端用户数据格式到前端格式
    const frontendUserData = {
      id: backendUser.id,
      name: backendUser.username,
      email: backendUser.email,
      avatar: '',
      permissions: [],
      role: backendUser.role,
      roleId: backendUser.role_id || null,
      roleName: getRoleDisplayName(backendUser.role),
      tenantId: backendUser.tenant_id,
      tenantName: form.tenantName, // Set the tenant name from form
      tenantRole: getTenantRoleFromRole(backendUser.role), // User will have tenant role
      isAuthenticated: true
    }

    // 使用userStore管理状态
    userStore.setUser(frontendUserData)
    userStore.setTokens(access_token, refresh_token, false)

    successMessage.value = '账户创建成功！正在跳转到登录页面...'

    // Redirect to login page after successful registration
    setTimeout(() => {
      emit('switchToLogin')
    }, 2000)

  } catch (error) {
    // Handle different error scenarios
    if (error.response?.status === 409) {
      ElMessage.error('该用户名或邮箱已被注册，请使用其他用户名或邮箱')
    } else if (error.response?.status === 400) {
      ElMessage.error('注册信息有误，请检查后重试')
    } else if (error.message.includes('Network')) {
      ElMessage.error('网络连接失败，请检查您的网络连接')
    } else {
      ElMessage.error('注册失败，请稍后重试')
    }
  } finally {
    loading.value = false
  }
}

const showTerms = () => {
  // TODO: Implement terms modal or navigation
  console.log('Show terms of service')
}

const showPrivacy = () => {
  // TODO: Implement privacy policy modal or navigation
  console.log('Show privacy policy')
}
</script>

<style scoped>
.register-form {
  width: 100%;
  max-width: 450px;
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

/* User Type Options */
.user-type-options {
  display: flex;
  flex-direction: column;
  gap: var(--space-sm);
}

.user-type-option {
  cursor: pointer;
  padding: var(--space-md);
  border: 2px solid var(--glass-border);
  border-radius: var(--radius-md);
  background: rgba(255, 255, 255, 0.05);
  transition: all var(--duration-normal) ease;
}

.user-type-option:hover {
  border-color: var(--color-stellar-blue);
  background: rgba(59, 130, 246, 0.1);
}

.user-type-option.selected {
  border-color: var(--color-stellar-blue);
  background: rgba(59, 130, 246, 0.2);
  box-shadow: 0 0 20px rgba(59, 130, 246, 0.3);
}

.user-type-radio {
  display: none;
}

.option-content {
  display: flex;
  align-items: center;
  gap: var(--space-md);
}

.option-icon {
  font-size: 1.5rem;
  opacity: 0.8;
}

.option-details {
  flex: 1;
}

.option-title {
  font-weight: 600;
  color: var(--color-starlight);
  margin-bottom: 2px;
}

.option-description {
  font-size: 0.875rem;
  color: var(--color-starlight);
  opacity: 0.7;
}

/* Input Wrapper */
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

/* Checkbox Styles */
.checkbox-wrapper {
  display: flex;
  align-items: flex-start;
  gap: var(--space-sm);
  cursor: pointer;
  user-select: none;
  line-height: 1.5;
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
  flex-shrink: 0;
  margin-top: 2px;
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
  flex: 1;
}

.terms-link {
  color: var(--color-stellar-blue);
  text-decoration: underline;
  transition: all var(--duration-normal) ease;
}

.terms-link:hover {
  color: var(--color-nebula-purple);
  text-shadow: 0 0 5px var(--color-nebula-glow);
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

/* Login Link */
.login-link {
  text-align: center;
  margin-top: var(--space-xl);
  color: var(--color-starlight);
  opacity: 0.7;
  font-size: 0.875rem;
}

.switch-form-btn {
  background: none;
  border: none;
  color: var(--color-stellar-blue);
  font-size: 0.875rem;
  cursor: pointer;
  text-decoration: underline;
  margin-left: var(--space-xs);
  transition: all var(--duration-normal) ease;
}

.switch-form-btn:hover:not(:disabled) {
  color: var(--color-nebula-purple);
  text-shadow: 0 0 5px var(--color-nebula-glow);
}

.switch-form-btn:disabled {
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

[data-theme="light"] .user-type-option {
  background: rgba(0, 0, 0, 0.05);
  border-color: rgba(0, 0, 0, 0.2);
}

[data-theme="light"] .user-type-option:hover {
  border-color: var(--color-stellar-blue);
  background: rgba(59, 130, 246, 0.1);
}

[data-theme="light"] .user-type-option.selected {
  border-color: var(--color-stellar-blue);
  background: rgba(59, 130, 246, 0.2);
  box-shadow: 0 0 20px rgba(56, 189, 248, 0.3);
}

[data-theme="light"] .option-title {
  color: var(--color-text-primary);
}

[data-theme="light"] .option-description {
  color: var(--color-text-secondary);
  opacity: 0.7;
}

[data-theme="light"] .checkbox-custom {
  background: rgba(0, 0, 0, 0.05);
  border-color: rgba(0, 0, 0, 0.2);
}

[data-theme="light"] .checkbox-label {
  color: var(--color-text-secondary);
  opacity: 0.8;
}

/* 修复浅色主题下复选框勾选标记颜色 */
[data-theme="light"] .checkbox-check {
  color: white;
}

[data-theme="light"] .checkbox-input:checked + .checkbox-custom {
  background: linear-gradient(135deg, #3b82f6 0%, #8b5cf6 100%);
  border-color: #3b82f6;
}

/* 确保浅色主题下复选框勾选标记始终可见 */
[data-theme="light"] .checkbox-input:checked + .checkbox-custom .checkbox-check {
  color: white !important;
  opacity: 1 !important;
}

[data-theme="light"] .terms-link {
  color: var(--color-stellar-blue);
}

[data-theme="light"] .terms-link:hover {
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

[data-theme="light"] .switch-form-btn {
  color: var(--color-stellar-blue);
}

[data-theme="light"] .switch-form-btn:hover:not(:disabled) {
  color: var(--color-nebula-purple);
  text-shadow: 0 0 5px rgba(124, 58, 237, 0.4);
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
  .form-title {
    font-size: 1.5rem;
  }

  .form-input {
    font-size: 16px; /* Prevents zoom on iOS */
  }

  .option-content {
    gap: var(--space-sm);
  }

  .option-icon {
    font-size: 1.25rem;
  }
}

@media (max-width: 380px) {
  .user-type-options {
    gap: var(--space-xs);
  }

  .user-type-option {
    padding: var(--space-sm);
  }
}
</style>