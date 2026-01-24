<template>
  <div class="h-screen w-full flex bg-slate-50 overflow-hidden">
    
    <!-- Sidebar -->
    <aside class="w-[240px] bg-slate-900 text-slate-300 flex flex-col flex-shrink-0 transition-all duration-300 z-20">
      <!-- Logo Area -->
      <div class="h-[60px] flex items-center px-6 border-b border-slate-800 bg-slate-900">
        <div class="flex items-center gap-3 text-white">
             <div class="w-8 h-8 rounded bg-indigo-600 flex items-center justify-center shadow-lg shadow-indigo-900/50">
                <el-icon :size="18"><Monitor /></el-icon>
             </div>
             <span class="font-bold text-lg tracking-tight">AstraFlow</span>
        </div>
      </div>

      <!-- Navigation -->
      <el-scrollbar class="flex-1">
        <el-menu
          :default-active="activeRoute"
          class="!border-none !bg-transparent w-full mt-2"
          text-color="#94a3b8"
          active-text-color="#ffffff"
          background-color="transparent"
          router
        >
          <div class="px-4 mb-2 mt-4 text-xs font-semibold uppercase tracking-wider text-slate-500">
            Overview
          </div>

          <el-menu-item index="/dashboard" class="menu-item-override">
            <el-icon><Odometer /></el-icon>
            <span>Dashboard</span>
          </el-menu-item>

          <!-- Audit Section -->
          <template v-if="isAuditor || isAdmin">
             <div class="px-4 mb-2 mt-6 text-xs font-semibold uppercase tracking-wider text-slate-500">
                Finance Operations
             </div>
             <el-menu-item index="/audit" class="menu-item-override">
                <el-icon><DocumentChecked /></el-icon>
                <span>Task Pool</span>
             </el-menu-item>
             <el-menu-item index="/audit/settlement" class="menu-item-override">
                <el-icon><Wallet /></el-icon>
                <span>Settlement</span>
             </el-menu-item>
          </template>

          <!-- Data & Archive -->
          <div class="px-4 mb-2 mt-6 text-xs font-semibold uppercase tracking-wider text-slate-500">
            Records
          </div>
          <el-menu-item index="/archive" class="menu-item-override">
            <el-icon><Collection /></el-icon>
            <span>Archive</span>
          </el-menu-item>

          <!-- Admin Section -->
          <template v-if="isAdmin">
             <div class="px-4 mb-2 mt-6 text-xs font-semibold uppercase tracking-wider text-slate-500">
                Administration
             </div>
             <el-menu-item index="/system/employees" class="menu-item-override">
                <el-icon><User /></el-icon>
                <span>Team Management</span>
             </el-menu-item>
          </template>
        </el-menu>
      </el-scrollbar>

      <!-- User Profile (Bottom Sidebar) -->
      <div class="p-4 border-t border-slate-800 bg-slate-900">
          <div class="flex items-center gap-3">
              <el-avatar :size="36" :src="userStore.userInfo?.avatar" class="border-2 border-slate-700" />
              <div class="flex flex-col overflow-hidden text-left">
                  <span class="text-sm font-medium text-white truncate">{{ userStore.userInfo?.name || 'User' }}</span>
                  <span class="text-xs text-slate-500 truncate capitalize">{{ userStore.role_key }}</span>
              </div>
          </div>
      </div>
    </aside>

    <!-- Main Content Area -->
    <div class="flex-1 flex flex-col min-w-0 bg-slate-50">
      
      <!-- Header -->
      <header class="h-[60px] bg-white border-b border-slate-200 flex items-center justify-between px-6 shadow-sm z-10">
        <!-- Page Title -->
        <div class="flex items-center text-slate-800 font-semibold">
             {{ routeTitle }}
        </div>

        <!-- Right Actions -->
        <div class="flex items-center gap-4">
             <el-button circle plain size="small">
                <el-icon><Bell /></el-icon>
             </el-button>
             
             <el-dropdown trigger="click" @command="handleCommand">
                <span class="flex items-center gap-2 cursor-pointer outline-none text-slate-600 hover:text-slate-900 transition-colors">
                    <span class="text-sm font-medium">{{ userStore.userInfo?.name }}</span>
                    <el-icon><ArrowDown /></el-icon>
                </span>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item command="profile">Profile</el-dropdown-item>
                    <el-dropdown-item divided command="logout" class="text-rose-500">Logout</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
             </el-dropdown>
        </div>
      </header>

      <!-- Scrollable View -->
      <main class="flex-1 overflow-auto p-6 relative">
         <router-view v-slot="{ Component }">
            <transition name="fade" mode="out-in">
                <component :is="Component" />
            </transition>
         </router-view>
      </main>

    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '../stores/userStore'
import { 
    Odometer, 
    DocumentChecked, 
    User, 
    Monitor, 
    Bell, 
    ArrowDown,
    Wallet,
    Collection
} from '@element-plus/icons-vue'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

// Route States
const activeRoute = computed(() => {
    // Handle nested routes like /audit/tasks/:id highlighting /audit
    if (route.path.startsWith('/audit')) return '/audit'
    return route.path
})

const routeTitle = computed(() => (route.meta.title as string) || 'Dashboard')

// Role Helpers
const isAdmin = computed(() => userStore.role_key === 'admin')
const isAuditor = computed(() => userStore.role_key === 'auditor')

// Handlers
const handleCommand = (command: string) => {
    if (command === 'logout') {
        userStore.logout()
        router.push('/login')
    }
}
</script>

<style scoped>
.menu-item-override {
    margin: 4px 16px;
    height: 40px;
    line-height: 40px;
    border-radius: 8px;
}

.menu-item-override:hover {
    background-color: #1e293b !important;
    color: #ffffff !important;
}

.menu-item-override.is-active {
    background-color: #4f46e5 !important;
    color: #ffffff !important;
    font-weight: 500;
    box-shadow: 0 4px 6px -1px rgba(79, 70, 229, 0.3);
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
