<template>
  <AuthLayout>
    <div class="login-view">
      <!-- Tab Navigation -->
      <div class="tab-navigation">
        <button
          :class="['tab-btn', { active: currentTab === 'login' }]"
          @click="switchTab('login')"
          :disabled="isTransitioning"
        >
          <span class="tab-icon">🔐</span>
          <span class="tab-text">登录</span>
        </button>
        <button
          :class="['tab-btn', { active: currentTab === 'register' }]"
          @click="switchTab('register')"
          :disabled="isTransitioning"
        >
          <span class="tab-icon">✨</span>
          <span class="tab-text">注册</span>
        </button>
      </div>

      <!-- Form Container -->
      <div class="form-container">
        <Transition
          name="form-slide"
          mode="out-in"
          @before-enter="onBeforeEnter"
          @after-enter="onAfterEnter"
        >
          <!-- Login Form -->
          <div v-if="currentTab === 'login'" key="login" class="form-panel">
            <LoginForm @switch-to-register="switchTab('register')" />
          </div>

          <!-- Register Form -->
          <div v-else key="register" class="form-panel">
            <RegisterForm @switch-to-login="switchTab('login')" />
          </div>
        </Transition>
      </div>
    </div>
  </AuthLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/userStore'
import AuthLayout from './AuthLayout.vue'
import LoginForm from '@/components/auth/LoginForm.vue'
import RegisterForm from '@/components/auth/RegisterForm.vue'

const router = useRouter()
const userStore = useUserStore()

const currentTab = ref('login')
const isTransitioning = ref(false)

const switchTab = (tab) => {
  if (isTransitioning.value || currentTab.value === tab) return

  isTransitioning.value = true
  currentTab.value = tab

  // Reset transitioning state after animation
  setTimeout(() => {
    isTransitioning.value = false
  }, 300)
}

const onBeforeEnter = () => {
  isTransitioning.value = true
}

const onAfterEnter = () => {
  isTransitioning.value = false
}

const handleDemoLogin = async () => {
  try {
    // Auto-login with demo credentials
    await userStore.login('demo@astraflow.com', 'Demo@123', true)
    router.push('/visualization')
  } catch (error) {
    console.error('Demo login error:', error)
  }
}

// Check if user is already authenticated and redirect
onMounted(() => {
  if (userStore.isAuthenticated) {
    router.push('/visualization')
  }
})
</script>

<style scoped>
.login-view {
  width: 100%;
  max-width: 500px;
  margin: 0 auto;
}

/* Tab Navigation */
.tab-navigation {
  display: flex;
  background: rgba(255, 255, 255, 0.05);
  border-radius: var(--radius-xl);
  padding: 4px;
  margin-bottom: var(--space-5xl);
  position: relative;
}

.tab-btn {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: var(--space-sm);
  padding: var(--space-md) var(--space-lg);
  background: transparent;
  border: none;
  border-radius: var(--radius-lg);
  color: var(--color-starlight);
  opacity: 0.7;
  cursor: pointer;
  transition: all var(--duration-normal) ease;
  font-size: 1rem;
  font-weight: 500;
  position: relative;
  z-index: 2;
}

.tab-btn:hover:not(:disabled) {
  opacity: 1;
}

.tab-btn.active {
  background: var(--gradient-stellar);
  opacity: 1;
  box-shadow: 0 4px 15px rgba(59, 130, 246, 0.3);
}

.tab-btn:disabled {
  cursor: not-allowed;
}

.tab-icon {
  font-size: 1.1rem;
}

.tab-text {
  font-weight: 600;
}

/* Form Container */
.form-container {
  position: relative;
  min-height: 500px; /* Adjust based on your form heights */
}

.form-panel {
  width: 100%;
}

/* Form Slide Transition */
.form-slide-enter-active,
.form-slide-leave-active {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.form-slide-enter-from {
  opacity: 0;
  transform: translateX(30px);
}

.form-slide-leave-to {
  opacity: 0;
  transform: translateX(-30px);
}

.form-slide-enter-to,
.form-slide-leave-from {
  opacity: 1;
  transform: translateX(0);
}

/* Social Login Divider */
.social-divider {
  display: flex;
  align-items: center;
  gap: var(--space-md);
  margin: var(--space-2xl) 0 var(--space-lg);
}

.divider-line {
  flex: 1;
  height: 1px;
  background: linear-gradient(
    to right,
    transparent,
    var(--color-starlight),
    transparent
  );
  opacity: 0.3;
}

.divider-text {
  color: var(--color-starlight);
  opacity: 0.6;
  font-size: 0.875rem;
  font-weight: 500;
}

/* Social Login Buttons */
.social-login {
  display: flex;
  flex-direction: column;
  gap: var(--space-sm);
  margin-bottom: var(--space-2xl);
}

.social-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: var(--space-md);
  padding: var(--space-md);
  border: 1px solid var(--glass-border);
  border-radius: var(--radius-md);
  background: rgba(255, 255, 255, 0.05);
  color: var(--color-starlight);
  cursor: pointer;
  transition: all var(--duration-normal) ease;
  text-decoration: none;
  font-size: 0.875rem;
}

.social-btn:hover:not(:disabled) {
  background: rgba(255, 255, 255, 0.1);
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.3);
}

.social-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.social-icon {
  font-size: 1.2rem;
}

.social-text {
  font-weight: 500;
}

/* Quick Access */
.quick-access {
  text-align: center;
  margin-top: var(--space-2xl);
  padding-top: var(--space-xl);
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.quick-access-title {
  color: var(--color-starlight);
  opacity: 0.6;
  font-size: 0.875rem;
  margin-bottom: var(--space-md);
  font-weight: 500;
}

.quick-access-links {
  display: flex;
  justify-content: center;
  gap: var(--space-lg);
  flex-wrap: wrap;
}

.quick-link {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: var(--space-xs);
  padding: var(--space-sm);
  text-decoration: none;
  color: var(--color-starlight);
  opacity: 0.6;
  transition: all var(--duration-normal) ease;
  border-radius: var(--radius-md);
  min-width: 80px;
}

.quick-link:hover {
  opacity: 1;
  background: rgba(255, 255, 255, 0.05);
  transform: translateY(-2px);
}

.quick-link-icon {
  font-size: 1.5rem;
  margin-bottom: 2px;
}

.quick-link-text {
  font-size: 0.75rem;
  font-weight: 500;
}

/* Light theme adjustments */
[data-theme="light"] .tab-navigation {
  background: rgba(0, 0, 0, 0.05);
}

[data-theme="light"] .tab-btn {
  color: var(--color-text-secondary);
  opacity: 0.7;
}

[data-theme="light"] .tab-btn.active {
  color: white;
  opacity: 1;
  box-shadow: 0 4px 15px rgba(56, 189, 248, 0.3);
}

[data-theme="light"] .tab-btn:hover:not(:disabled) {
  opacity: 1;
}

[data-theme="light"] .demo-notice {
  background: var(--glass-bg);
  border: 1px solid var(--glass-border);
}

[data-theme="light"] .demo-title {
  color: var(--color-text-primary);
}

[data-theme="light"] .demo-content p {
  color: var(--color-text-secondary);
  opacity: 0.8;
}

[data-theme="light"] .demo-login-btn {
  background: var(--gradient-stellar);
  color: white;
}

/* Responsive Design */
@media (max-width: 768px) {
  .login-view {
    max-width: 100%;
  }

  .tab-btn {
    padding: var(--space-sm) var(--space-md);
    font-size: 0.875rem;
  }

  .tab-icon {
    font-size: 1rem;
  }

  .quick-access-links {
    gap: var(--space-md);
  }

  .quick-link {
    min-width: 70px;
  }
}

@media (max-width: 480px) {
  .social-login {
    gap: var(--space-xs);
  }

  .social-btn {
    padding: var(--space-sm);
    font-size: 0.8rem;
  }

  .quick-access-links {
    gap: var(--space-sm);
  }

  .quick-link {
    min-width: 60px;
    padding: var(--space-xs);
  }

  .quick-link-icon {
    font-size: 1.25rem;
  }

  .quick-link-text {
    font-size: 0.7rem;
  }
}

@media (max-width: 380px) {
  .tab-navigation {
    margin-bottom: var(--space-xl);
  }

  .tab-btn {
    gap: var(--space-xs);
    padding: var(--space-xs) var(--space-sm);
  }

  .tab-text {
    font-size: 0.8rem;
  }
}
</style>