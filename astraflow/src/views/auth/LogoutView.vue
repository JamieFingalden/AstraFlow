<template>
  <div class="logout-container" v-if="showMessage">
    <div class="logout-content">
      <div class="logout-icon">👋</div>
      <h2 class="logout-title">{{ message.title }}</h2>
      <p class="logout-message">{{ message.content }}</p>
      <div class="logout-actions" v-if="message.showRedirect">
        <router-link to="/login" class="primary-btn">
          返回登录
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/userStore'

const router = useRouter()
const userStore = useUserStore()

// Messages for different logout scenarios
const messages = {
  success: {
    title: '已退出登录',
    content: '您已成功退出登录，欢迎下次使用 AstraFlow',
    showRedirect: true
  },
  error: {
    title: '退出登录失败',
    content: '退出登录时遇到问题，请稍后重试',
    showRedirect: true
  },
  processing: {
    title: '正在退出登录...',
    content: '请稍候，正在安全退出您的账户',
    showRedirect: false
  }
}

const showMessage = ref(false)
const message = ref(messages.processing)

onMounted(async () => {
  try {
    // Show processing message immediately
    message.value = messages.processing
    showMessage.value = true

    // Perform logout
    await userStore.logout()

    // Show success message
    message.value = messages.success

    // Redirect to login after delay
    setTimeout(() => {
      router.push('/login')
    }, 2000)
  } catch (error) {
    console.error('Logout error:', error)

    // Show error message
    message.value = messages.error
  }
})
</script>

<style scoped>
.logout-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #0f172a, #1e293b);
  padding: 1rem;
}

.logout-content {
  text-align: center;
  max-width: 400px;
  padding: 2rem;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 1rem;
  backdrop-filter: blur(10px);
}

.logout-icon {
  font-size: 4rem;
  margin-bottom: 1rem;
  animation: wave 1s ease-in-out;
}

@keyframes wave {
  0%, 100% { transform: rotate(0deg); }
  25% { transform: rotate(10deg); }
  75% { transform: rotate(-10deg); }
}

.logout-title {
  font-size: 1.5rem;
  font-weight: 600;
  color: white;
  margin-bottom: 0.5rem;
}

.logout-message {
  color: rgba(255, 255, 255, 0.7);
  margin-bottom: 1.5rem;
  line-height: 1.5;
}

.logout-actions {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.primary-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0.75rem 1.5rem;
  background: linear-gradient(135deg, #3b82f6, #06b6d4);
  color: white;
  text-decoration: none;
  border-radius: 0.5rem;
  font-weight: 500;
  transition: all 0.2s ease;
  border: none;
  cursor: pointer;
}

.primary-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 25px rgba(59, 130, 246, 0.3);
}

.primary-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}
</style>