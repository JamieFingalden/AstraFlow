<template>
  <div class="h-screen w-full flex bg-transparent overflow-hidden p-3 md:p-4 gap-3 md:gap-4">
    
    <!-- Sidebar -->
    <aside class="w-[250px] sidebar-panel text-slate-200 flex flex-col flex-shrink-0 transition-all duration-300 z-20 rounded-2xl overflow-hidden">
      <!-- Logo Area -->
      <div class="h-[64px] flex items-center px-6 border-b border-slate-800/60">
        <div class="flex items-center gap-3 text-white">
             <div class="w-8 h-8 rounded-lg bg-sky-500 flex items-center justify-center shadow-lg shadow-sky-900/40">
                <el-icon :size="18"><Monitor /></el-icon>
             </div>
             <span class="font-bold text-lg tracking-tight">AstraFlow</span>
        </div>
      </div>

      <!-- Navigation -->
      <el-scrollbar class="flex-1">
        <el-menu
          :default-active="activeRoute"
          class="!border-none !bg-transparent w-full mt-3"
          text-color="#94a3b8"
          active-text-color="#ffffff"
          background-color="transparent"
          unique-opened
          router
        >
          <!-- Dashboard -->
          <el-menu-item index="/dashboard" class="menu-item-override">
            <el-icon><Odometer /></el-icon>
            <span>仪表盘</span>
          </el-menu-item>
          
          <!-- Upload -->
          <el-menu-item index="/upload" class="menu-item-override">
            <el-icon><Upload /></el-icon>
            <span>上传单据</span>
          </el-menu-item>

          <!-- My Invoices -->
          <el-menu-item index="/my-invoices" class="menu-item-override">
            <el-icon><FolderOpened /></el-icon>
            <span>我的票夹</span>
          </el-menu-item>

          <!-- Audit Center (Auditor Only) -->
          <el-sub-menu index="audit-center" v-if="isAuditor">
            <template #title>
              <el-icon><DocumentChecked /></el-icon>
              <span>审核中心</span>
            </template>
            <el-menu-item index="/audit/tasks" class="menu-item-override">
              <el-icon><Monitor /></el-icon>
              <span>审核任务池</span>
            </el-menu-item>
            <el-menu-item index="/audit/settlement" class="menu-item-override">
              <el-icon><Wallet /></el-icon>
              <span>结算中心</span>
            </el-menu-item>
          </el-sub-menu>

          <!-- System Management (Admin Only) -->
          <el-sub-menu index="system-mgmt" v-if="isAdmin">
            <template #title>
              <el-icon><User /></el-icon>
              <span>系统管理</span>
            </template>
            <el-menu-item index="/system/employees" class="menu-item-override">
              <el-icon><User /></el-icon>
              <span>员工管理</span>
            </el-menu-item>
          </el-sub-menu>

          <!-- Archive -->
          <el-menu-item index="/archive" class="menu-item-override">
            <el-icon><Collection /></el-icon>
            <span>历史归档</span>
          </el-menu-item>
        </el-menu>
      </el-scrollbar>

      <!-- User Profile (Bottom Sidebar) -->
      <div class="p-4 border-t border-slate-700/50 bg-slate-900/40 backdrop-blur-sm">
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
    <div class="flex-1 flex flex-col min-w-0 bg-transparent">
      
      <!-- Header -->
      <header class="h-[64px] page-shell !rounded-2xl !shadow-none flex items-center justify-between px-5 md:px-6 z-10">
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
                    <el-dropdown-item command="profile">个人资料</el-dropdown-item>
                    <el-dropdown-item divided command="logout" class="text-rose-500">退出登录</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
             </el-dropdown>
        </div>
      </header>

      <!-- Scrollable View -->
      <main class="flex-1 overflow-auto p-3 md:p-4 relative">
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
    Collection,
    Upload,
    FolderOpened
} from '@element-plus/icons-vue'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

// Route States
const activeRoute = computed(() => route.path)

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

.sidebar-panel {
    background:
      radial-gradient(circle at 20% 0%, rgba(56, 189, 248, 0.2), transparent 38%),
      radial-gradient(circle at 110% 120%, rgba(34, 197, 94, 0.15), transparent 45%),
      linear-gradient(180deg, #0b1b30 0%, #0f2742 55%, #12314f 100%);
    border: 1px solid rgba(148, 163, 184, 0.2);
    box-shadow: 0 16px 36px rgba(12, 31, 54, 0.35);
}

.menu-item-override:hover {
    background-color: rgba(56, 103, 152, 0.45) !important;
    color: #ffffff !important;
}

.menu-item-override.is-active {
    background: linear-gradient(135deg, #1d7cf2 0%, #16a34a 120%) !important;
    color: #ffffff !important;
    font-weight: 500;
    box-shadow: 0 10px 20px rgba(20, 87, 196, 0.4);
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
