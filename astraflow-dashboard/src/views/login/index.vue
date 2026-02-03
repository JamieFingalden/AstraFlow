<template>
  <div class="min-h-screen bg-slate-50 flex items-center justify-center p-4 relative overflow-hidden">
    <!-- Abstract Background Pattern -->
    <div class="absolute inset-0 z-0 opacity-40">
        <div class="absolute top-0 -left-4 w-72 h-72 bg-purple-300 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob"></div>
        <div class="absolute top-0 -right-4 w-72 h-72 bg-indigo-300 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob animation-delay-2000"></div>
        <div class="absolute -bottom-8 left-20 w-72 h-72 bg-blue-300 rounded-full mix-blend-multiply filter blur-xl opacity-70 animate-blob animation-delay-4000"></div>
    </div>

    <div class="w-full max-w-md bg-white rounded-2xl shadow-xl z-10 p-8 md:p-10 border border-slate-100">
      <div class="mb-10 text-center">
        <div class="inline-flex items-center justify-center w-12 h-12 rounded-lg bg-indigo-600 text-white mb-4 shadow-lg shadow-indigo-200">
           <el-icon :size="24"><Monitor /></el-icon>
        </div>
        <h1 class="text-2xl font-bold text-slate-900 tracking-tight">AstraFlow</h1>
        <p class="text-slate-500 mt-2 text-sm">Sign in to your dashboard</p>
      </div>

      <el-form 
        ref="loginFormRef"
        :model="loginForm"
        :rules="rules"
        label-position="top"
        class="space-y-4"
        @keyup.enter="handleLogin"
      >
        <el-form-item prop="username" class="!mb-4">
          <label class="block text-sm font-medium text-slate-700 mb-1.5">Username</label>
          <el-input 
            v-model="loginForm.username" 
            placeholder="Enter your username"
            class="!h-10 custom-input"
          >
             <template #prefix>
                <el-icon class="text-slate-400"><User /></el-icon>
             </template>
          </el-input>
        </el-form-item>

        <el-form-item prop="password" class="!mb-6">
          <label class="block text-sm font-medium text-slate-700 mb-1.5">Password</label>
          <el-input 
            v-model="loginForm.password" 
            type="password" 
            placeholder="••••••••"
            show-password
            class="!h-10 custom-input"
          >
            <template #prefix>
                <el-icon class="text-slate-400"><Lock /></el-icon>
             </template>
          </el-input>
        </el-form-item>

        <el-button 
            type="primary" 
            :loading="loading"
            class="!w-full !h-11 !bg-indigo-600 !border-indigo-600 hover:!bg-indigo-700 hover:!border-indigo-700 !text-white !font-medium !rounded-lg !text-base transition-colors duration-200 shadow-md shadow-indigo-100"
            @click="handleLogin"
        >
          Sign In
        </el-button>
        
        <div class="mt-6 text-center text-xs text-slate-400">
            &copy; 2026 AstraFlow Systems. All rights reserved.
        </div>
      </el-form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { User, Lock, Monitor } from '@element-plus/icons-vue'
import { useUserStore } from '../../stores/userStore'
import { ElMessage } from 'element-plus'

const router = useRouter()
const userStore = useUserStore()

const loginFormRef = ref()
const loading = ref(false)

const loginForm = reactive({
  username: '',
  password: ''
})

const rules = {
  username: [{ required: true, message: 'Please enter your username', trigger: 'blur' }],
  password: [{ required: true, message: 'Please enter your password', trigger: 'blur' }]
}

const handleLogin = async () => {
  if (!loginFormRef.value) return
  
  await loginFormRef.value.validate(async (valid: boolean) => {
    if (valid) {
      loading.value = true
      try {
        await userStore.login(loginForm)
        ElMessage.success('Welcome back!')
        router.push('/dashboard')
      } catch (error: any) {
        console.error('Login error:', error)
        ElMessage.error(error.message || 'Login failed. Please check your credentials.')
      } finally {
        loading.value = false
      }
    }
  })
}
</script>

<style scoped>
/* Custom overrides for Element Plus inputs to match Tailwind/Linear style */
:deep(.el-input__wrapper) {
    box-shadow: 0 0 0 1px #e2e8f0 inset;
    border-radius: 0.5rem;
    padding: 1px 11px;
    background-color: #ffffff;
    transition: all 0.2s;
}

:deep(.el-input__wrapper.is-focus) {
    box-shadow: 0 0 0 2px #4f46e5 inset !important; /* Indigo-600 */
}

:deep(.el-input__wrapper:hover) {
    box-shadow: 0 0 0 1px #cbd5e1 inset;
}

.animate-blob {
    animation: blob 7s infinite;
}
.animation-delay-2000 {
    animation-delay: 2s;
}
.animation-delay-4000 {
    animation-delay: 4s;
}

@keyframes blob {
    0% { transform: translate(0px, 0px) scale(1); }
    33% { transform: translate(30px, -50px) scale(1.1); }
    66% { transform: translate(-20px, 20px) scale(0.9); }
    100% { transform: translate(0px, 0px) scale(1); }
}
</style>
