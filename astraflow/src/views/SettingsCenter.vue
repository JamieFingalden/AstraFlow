<template>
  <div :class="'min-h-screen transition-colors duration-300 ' + (isDark ? 'bg-gray-900' : 'bg-gradient-to-br from-blue-50 via-cyan-50 to-indigo-50')">
    <!-- Particle background effect -->
    <div class="fixed inset-0 overflow-hidden pointer-events-none">
      <div :class="'absolute -top-40 -right-40 w-80 h-80 rounded-full mix-blend-multiply filter blur-xl animate-pulse ' + (isDark ? 'bg-cyan-900/30' : 'bg-cyan-500')"></div>
      <div :class="'absolute -bottom-40 -left-40 w-80 h-80 rounded-full mix-blend-multiply filter blur-xl animate-pulse ' + (isDark ? 'bg-blue-900/30' : 'bg-blue-500')"></div>
      <div :class="'absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-80 h-80 rounded-full mix-blend-multiply filter blur-xl animate-pulse ' + (isDark ? 'bg-purple-900/30' : 'bg-purple-500')"></div>
    </div>

    <!-- Header -->
    <header :class="'relative z-10 backdrop-blur-md border-b transition-all duration-300 ' + (isDark ? 'bg-gray-800/70 border-gray-700/50' : 'bg-white/70 border-gray-200/50')">
      <div class="px-4 mx-auto max-w-7xl sm:px-6 lg:px-8">
        <div class="flex items-center justify-between h-16">
          <div class="flex items-center space-x-4">
            <!-- 返回按钮 -->
            <router-link
              to="/"
              :class="'flex items-center space-x-2 px-3 py-2 rounded-lg transition-all duration-200 ' + (isDark ? 'hover:bg-gray-700/50 text-gray-300 hover:text-white' : 'hover:bg-gray-100 text-gray-600 hover:text-gray-900')"
              title="返回首页"
            >
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="m15 18-6-6 6-6"/>
              </svg>
              <span class="hidden sm:inline">返回</span>
            </router-link>

            <div class="text-2xl font-bold text-transparent bg-gradient-to-r from-cyan-400 to-blue-500 bg-clip-text">
              AstraFlow
            </div>
            <h1 :class="'text-xl font-semibold transition-colors duration-300 ' + (isDark ? 'text-white' : 'text-gray-800')">
              设置中心
            </h1>
          </div>

          <div class="flex items-center space-x-4">
            <!-- 主题切换按钮 -->
            <button
              @click="toggleTheme"
              :class="'p-2 rounded-lg transition-all duration-200 ' + (isDark ? 'bg-gray-700/50 hover:bg-gray-600/50 text-gray-300' : 'bg-gray-100 hover:bg-gray-200 text-gray-700')"
            >
              <SunIcon v-if="isDark" :size="20" />
              <MoonIcon v-else :size="20" />
            </button>
          </div>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="relative z-10 px-4 py-8 mx-auto max-w-7xl sm:px-6 lg:px-8">
      <div class="flex flex-col gap-6 lg:flex-row">
        <!-- 左侧导航菜单 -->
        <aside :class="'lg:w-64 flex-shrink-0 '">
          <!-- 移动端下拉菜单 -->
          <div class="mb-6 lg:hidden">
            <select
              v-model="settingsStore.activeTab"
              :class="'w-full px-4 py-3 rounded-lg border transition-all duration-300 focus:outline-none focus:ring-2 focus:ring-blue-500/50 ' + (isDark ? 'bg-gray-800/70 border-gray-700 text-white' : 'bg-white/70 border-gray-300 text-gray-900')"
              @change="handleTabChange"
            >
              <option value="account">账户信息</option>
              <option value="preferences">报销偏好</option>
              <option value="ai">AI 接口配置</option>
              <option value="appearance">系统外观</option>
            </select>
          </div>

          <!-- 桌面端垂直菜单 -->
          <nav :class="'hidden lg:block rounded-2xl p-4 backdrop-blur-md border transition-all duration-300 ' + (isDark ? 'bg-gray-800/70 border-gray-700/50' : 'bg-white/70 border-gray-200/50')">
            <ul class="space-y-2">
              <li>
                <button
                  @click="setActiveTab('account')"
                  :class="'w-full text-left px-4 py-3 rounded-lg transition-all duration-200 flex items-center space-x-3 ' +
                    (settingsStore.activeTab === 'account'
                      ? (isDark ? 'bg-blue-900/50 text-blue-300 border-l-4 border-blue-400' : 'bg-blue-50 text-blue-700 border-l-4 border-blue-500')
                      : (isDark ? 'hover:bg-gray-700/50 text-gray-300' : 'hover:bg-gray-100 text-gray-700'))"
                >
                  <UserIcon :size="20" />
                  <span>账户信息</span>
                </button>
              </li>
              <li>
                <button
                  @click="setActiveTab('preferences')"
                  :class="'w-full text-left px-4 py-3 rounded-lg transition-all duration-200 flex items-center space-x-3 ' +
                    (settingsStore.activeTab === 'preferences'
                      ? (isDark ? 'bg-blue-900/50 text-blue-300 border-l-4 border-blue-400' : 'bg-blue-50 text-blue-700 border-l-4 border-blue-500')
                      : (isDark ? 'hover:bg-gray-700/50 text-gray-300' : 'hover:bg-gray-100 text-gray-700'))"
                >
                  <SettingsIcon :size="20" />
                  <span>报销偏好</span>
                </button>
              </li>
              <li>
                <button
                  @click="setActiveTab('ai')"
                  :class="'w-full text-left px-4 py-3 rounded-lg transition-all duration-200 flex items-center space-x-3 ' +
                    (settingsStore.activeTab === 'ai'
                      ? (isDark ? 'bg-blue-900/50 text-blue-300 border-l-4 border-blue-400' : 'bg-blue-50 text-blue-700 border-l-4 border-blue-500')
                      : (isDark ? 'hover:bg-gray-700/50 text-gray-300' : 'hover:bg-gray-100 text-gray-700'))"
                >
                  <CpuIcon :size="20" />
                  <span>AI 接口配置</span>
                </button>
              </li>
              <li>
                <button
                  @click="setActiveTab('appearance')"
                  :class="'w-full text-left px-4 py-3 rounded-lg transition-all duration-200 flex items-center space-x-3 ' +
                    (settingsStore.activeTab === 'appearance'
                      ? (isDark ? 'bg-blue-900/50 text-blue-300 border-l-4 border-blue-400' : 'bg-blue-50 text-blue-700 border-l-4 border-blue-500')
                      : (isDark ? 'hover:bg-gray-700/50 text-gray-300' : 'hover:bg-gray-100 text-gray-700'))"
                >
                  <PaletteIcon :size="20" />
                  <span>系统外观</span>
                </button>
              </li>
            </ul>
          </nav>
        </aside>

        <!-- 右侧内容区域 -->
        <section class="flex-1">
          <!-- 账户信息面板 -->
          <div v-if="settingsStore.activeTab === 'account'" :class="'rounded-2xl p-6 backdrop-blur-md border transition-all duration-300 ' + (isDark ? 'bg-gray-800/70 border-gray-700/50' : 'bg-white/70 border-gray-200/50')">
            <h2 :class="'text-2xl font-bold mb-6 ' + (isDark ? 'text-white' : 'text-gray-900')">账户信息</h2>

            <div class="space-y-6">
              <!-- 头像 -->
              <div class="flex items-center space-x-4">
                <div :class="'w-20 h-20 rounded-full overflow-hidden border-4 ' + (isDark ? 'border-gray-600' : 'border-gray-200')">
                  <img :src="settingsStore.accountInfo.avatar" alt="头像" class="object-cover w-full h-full" />
                </div>
                <button :class="'px-4 py-2 rounded-lg transition-all duration-200 ' + (isDark ? 'bg-gray-700 hover:bg-gray-600 text-gray-300' : 'bg-gray-100 hover:bg-gray-200 text-gray-700')">
                  更换头像
                </button>
              </div>

              <!-- 用户信息表单 -->
              <div class="grid grid-cols-1 gap-6 md:grid-cols-2">
                <div>
                  <label :class="'block text-sm font-medium mb-2 ' + (isDark ? 'text-gray-300' : 'text-gray-700')">用户名</label>
                  <input
                    v-model="tempAccountInfo.username"
                    type="text"
                    :class="'w-full px-4 py-2 rounded-lg border transition-all duration-300 focus:outline-none focus:ring-2 focus:ring-blue-500/50 ' + (isDark ? 'bg-gray-700/50 border-gray-600 text-white' : 'bg-white/50 border-gray-300 text-gray-900')"
                  />
                </div>
                <div>
                  <label :class="'block text-sm font-medium mb-2 ' + (isDark ? 'text-gray-300' : 'text-gray-700')">邮箱</label>
                  <input
                    v-model="tempAccountInfo.email"
                    type="email"
                    :class="'w-full px-4 py-2 rounded-lg border transition-all duration-300 focus:outline-none focus:ring-2 focus:ring-blue-500/50 ' + (isDark ? 'bg-gray-700/50 border-gray-600 text-white' : 'bg-white/50 border-gray-300 text-gray-900')"
                  />
                </div>
                <div class="md:col-span-2">
                  <label :class="'block text-sm font-medium mb-2 ' + (isDark ? 'text-gray-300' : 'text-gray-700')">公司名称</label>
                  <input
                    v-model="tempAccountInfo.company"
                    type="text"
                    :class="'w-full px-4 py-2 rounded-lg border transition-all duration-300 focus:outline-none focus:ring-2 focus:ring-blue-500/50 ' + (isDark ? 'bg-gray-700/50 border-gray-600 text-white' : 'bg-white/50 border-gray-300 text-gray-900')"
                  />
                </div>
              </div>

              <!-- 保存按钮 -->
              <div class="flex justify-end">
                <button
                  @click="saveAccountInfo"
                  :class="'px-6 py-2 rounded-lg font-semibold transition-all duration-300 transform hover:scale-105 ' + (isDark ? 'bg-blue-600 hover:bg-blue-500 text-white' : 'bg-blue-500 hover:bg-blue-600 text-white')"
                >
                  保存修改
                </button>
              </div>
            </div>
          </div>

          <!-- 报销偏好面板 -->
          <div v-if="settingsStore.activeTab === 'preferences'" :class="'rounded-2xl p-6 backdrop-blur-md border transition-all duration-300 ' + (isDark ? 'bg-gray-800/70 border-gray-700/50' : 'bg-white/70 border-gray-200/50')">
            <h2 :class="'text-2xl font-bold mb-6 ' + (isDark ? 'text-white' : 'text-gray-900')">报销偏好</h2>

            <div class="space-y-6">
              <!-- 自动分类开关 -->
              <div class="flex items-center justify-between">
                <div>
                  <h3 :class="'font-medium ' + (isDark ? 'text-white' : 'text-gray-900')">自动分类</h3>
                  <p :class="'text-sm mt-1 ' + (isDark ? 'text-gray-400' : 'text-gray-600')">自动识别账单分类</p>
                </div>
                <label class="relative inline-flex items-center cursor-pointer">
                  <input
                    type="checkbox"
                    v-model="tempPreferences.autoCategory"
                    :class="'sr-only peer'"
                  />
                  <div :class="'w-11 h-6 bg-gray-200 peer-focus:outline-none rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[\'\'] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-600'"></div>
                </label>
              </div>

              <!-- 报销上限提醒 -->
              <div>
                <label :class="'block text-sm font-medium mb-2 ' + (isDark ? 'text-gray-300' : 'text-gray-700')">报销上限提醒阈值</label>
                <div class="flex items-center space-x-2">
                  <input
                    v-model.number="tempPreferences.spendingLimit"
                    type="number"
                    :class="'flex-1 px-4 py-2 rounded-lg border transition-all duration-300 focus:outline-none focus:ring-2 focus:ring-blue-500/50 ' + (isDark ? 'bg-gray-700/50 border-gray-600 text-white' : 'bg-white/50 border-gray-300 text-gray-900')"
                  />
                  <span :class="'text-sm ' + (isDark ? 'text-gray-400' : 'text-gray-600')">元</span>
                </div>
              </div>

              <!-- 货币单位 -->
              <div>
                <label :class="'block text-sm font-medium mb-2 ' + (isDark ? 'text-gray-300' : 'text-gray-700')">货币单位</label>
                <select
                  v-model="tempPreferences.currency"
                  :class="'w-full px-4 py-2 rounded-lg border transition-all duration-300 focus:outline-none focus:ring-2 focus:ring-blue-500/50 ' + (isDark ? 'bg-gray-700/50 border-gray-600 text-white' : 'bg-white/50 border-gray-300 text-gray-900')"
                >
                  <option value="CNY">人民币 (¥)</option>
                  <option value="USD">美元 ($)</option>
                </select>
              </div>

              <!-- 保存按钮 -->
              <div class="flex justify-end">
                <button
                  @click="savePreferences"
                  :class="'px-6 py-2 rounded-lg font-semibold transition-all duration-300 transform hover:scale-105 ' + (isDark ? 'bg-blue-600 hover:bg-blue-500 text-white' : 'bg-blue-500 hover:bg-blue-600 text-white')"
                >
                  保存偏好设置
                </button>
              </div>
            </div>
          </div>

          <!-- AI接口配置面板 -->
          <div v-if="settingsStore.activeTab === 'ai'" :class="'rounded-2xl p-6 backdrop-blur-md border transition-all duration-300 ' + (isDark ? 'bg-gray-800/70 border-gray-700/50' : 'bg-white/70 border-gray-200/50')">
            <h2 :class="'text-2xl font-bold mb-6 ' + (isDark ? 'text-white' : 'text-gray-900')">AI 接口配置</h2>

            <div class="space-y-6">
              <!-- Flask接口URL -->
              <div>
                <label :class="'block text-sm font-medium mb-2 ' + (isDark ? 'text-gray-300' : 'text-gray-700')">Flask 识别接口 URL</label>
                <input
                  v-model="tempAiConfig.flaskUrl"
                  type="url"
                  placeholder="http://localhost:5000/api/ocr"
                  :class="'w-full px-4 py-2 rounded-lg border transition-all duration-300 focus:outline-none focus:ring-2 focus:ring-blue-500/50 ' + (isDark ? 'bg-gray-700/50 border-gray-600 text-white' : 'bg-white/50 border-gray-300 text-gray-900')"
                />
              </div>

              <!-- API密钥 -->
              <div>
                <label :class="'block text-sm font-medium mb-2 ' + (isDark ? 'text-gray-300' : 'text-gray-700')">API 密钥</label>
                <input
                  v-model="tempAiConfig.apiKey"
                  type="password"
                  placeholder="输入您的API密钥"
                  :class="'w-full px-4 py-2 rounded-lg border transition-all duration-300 focus:outline-none focus:ring-2 focus:ring-blue-500/50 ' + (isDark ? 'bg-gray-700/50 border-gray-600 text-white' : 'bg-white/50 border-gray-300 text-gray-900')"
                />
              </div>

              <!-- 操作按钮 -->
              <div class="flex space-x-4">
                <button
                  @click="testConnection"
                  :disabled="testing"
                  :class="'px-6 py-2 rounded-lg font-semibold transition-all duration-300 ' + (testing ? 'opacity-50 cursor-not-allowed' : 'transform hover:scale-105') + ' ' + (isDark ? 'bg-green-600 hover:bg-green-500 text-white' : 'bg-green-500 hover:bg-green-600 text-white')"
                >
                  {{ testing ? '测试中...' : '测试连接' }}
                </button>
                <button
                  @click="saveAiConfig"
                  :class="'px-6 py-2 rounded-lg font-semibold transition-all duration-300 transform hover:scale-105 ' + (isDark ? 'bg-blue-600 hover:bg-blue-500 text-white' : 'bg-blue-500 hover:bg-blue-600 text-white')"
                >
                  保存配置
                </button>
              </div>

              <!-- 连接状态提示 -->
              <div v-if="connectionStatus" :class="'p-4 rounded-lg ' + (connectionStatus.success ? 'bg-green-100 text-green-800 dark:bg-green-900/50 dark:text-green-300' : 'bg-red-100 text-red-800 dark:bg-red-900/50 dark:text-red-300')">
                {{ connectionStatus.message }}
              </div>
            </div>
          </div>

          <!-- 系统外观面板 -->
          <div v-if="settingsStore.activeTab === 'appearance'" :class="'rounded-2xl p-6 backdrop-blur-md border transition-all duration-300 ' + (isDark ? 'bg-gray-800/70 border-gray-700/50' : 'bg-white/70 border-gray-200/50')">
            <h2 :class="'text-2xl font-bold mb-6 ' + (isDark ? 'text-white' : 'text-gray-900')">系统外观</h2>

            <div class="space-y-6">
              <!-- 主题切换 -->
              <div>
                <label :class="'block text-sm font-medium mb-3 ' + (isDark ? 'text-gray-300' : 'text-gray-700')">主题模式</label>
                <div class="grid grid-cols-3 gap-3">
                  <button
                    v-for="theme in themeOptions"
                    :key="theme.value"
                    @click="tempAppearance.theme = theme.value"
                    :class="'p-3 rounded-lg border-2 transition-all duration-200 ' +
                      (tempAppearance.theme === theme.value
                        ? 'border-blue-500 bg-blue-50 dark:bg-blue-900/30'
                        : 'border-gray-300 dark:border-gray-600')"
                  >
                    <component :is="theme.icon" :size="20" :class="'mx-auto mb-1 ' + (isDark ? 'text-gray-300' : 'text-gray-700')" />
                    <span :class="'text-xs ' + (isDark ? 'text-gray-300' : 'text-gray-700')">{{ theme.label }}</span>
                  </button>
                </div>
              </div>

              <!-- 主色调选择 -->
              <div>
                <label :class="'block text-sm font-medium mb-3 ' + (isDark ? 'text-gray-300' : 'text-gray-700')">主色调</label>
                <div class="grid grid-cols-4 gap-3">
                  <button
                    v-for="color in colorOptions"
                    :key="color.value"
                    @click="tempAppearance.primaryColor = color.value"
                    :class="'p-3 rounded-lg border-2 transition-all duration-200 ' +
                      (tempAppearance.primaryColor === color.value
                        ? 'border-gray-900 dark:border-white'
                        : 'border-gray-300 dark:border-gray-600')"
                  >
                    <div :class="'w-6 h-6 rounded-full mx-auto mb-1 ' + color.class"></div>
                    <span :class="'text-xs ' + (isDark ? 'text-gray-300' : 'text-gray-700')">{{ color.label }}</span>
                  </button>
                </div>
              </div>

              <!-- 保存按钮 -->
              <div class="flex justify-end">
                <button
                  @click="saveAppearance"
                  :class="'px-6 py-2 rounded-lg font-semibold transition-all duration-300 transform hover:scale-105 ' + (isDark ? 'bg-blue-600 hover:bg-blue-500 text-white' : 'bg-blue-500 hover:bg-blue-600 text-white')"
                >
                  保存外观设置
                </button>
              </div>
            </div>
          </div>
        </section>
      </div>
    </main>

    <!-- Footer -->
    <footer :class="'relative z-10 backdrop-blur-md border-t mt-12 transition-all duration-300 ' + (isDark ? 'bg-gray-800/70 border-gray-700/50' : 'bg-white/70 border-gray-200/50')">
      <div class="px-4 py-6 mx-auto max-w-7xl sm:px-6 lg:px-8">
        <p :class="'text-center text-sm transition-colors duration-300 ' + (isDark ? 'text-gray-400' : 'text-gray-600')">
          © 2025 AstraFlow · Smart Expense Made Simple
        </p>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useTheme } from '../composables/useTheme'
import { useSettingsStore } from '../stores/settings'
import {
  UserIcon,
  SettingsIcon,
  CpuIcon,
  PaletteIcon,
  SunIcon,
  MoonIcon,
  MonitorIcon
} from 'lucide-vue-next'

const { isDark, toggleTheme } = useTheme()
const settingsStore = useSettingsStore()

// 临时数据
const tempAccountInfo = ref({ ...settingsStore.accountInfo })
const tempPreferences = ref({ ...settingsStore.reimbursementPreferences })
const tempAiConfig = ref({ ...settingsStore.aiConfig })
const tempAppearance = ref({ ...settingsStore.appearance })

// 状态
const testing = ref(false)
const connectionStatus = ref(null)

// 主题选项
const themeOptions = [
  { value: 'light', label: '浅色', icon: SunIcon },
  { value: 'dark', label: '深色', icon: MoonIcon },
  { value: 'system', label: '跟随系统', icon: MonitorIcon }
]

// 颜色选项
const colorOptions = [
  { value: 'blue', label: '蓝色', class: 'bg-blue-500' },
  { value: 'green', label: '绿色', class: 'bg-green-500' },
  { value: 'purple', label: '紫色', class: 'bg-purple-500' },
  { value: 'gray', label: '灰色', class: 'bg-gray-500' }
]

// 方法
const setActiveTab = (tab) => {
  settingsStore.setActiveTab(tab)
}

const handleTabChange = (event) => {
  settingsStore.setActiveTab(event.target.value)
}

const saveAccountInfo = () => {
  settingsStore.updateAccountInfo(tempAccountInfo.value)
  showSuccessMessage('账户信息已保存')
}

const savePreferences = () => {
  settingsStore.updateReimbursementPreferences(tempPreferences.value)
  showSuccessMessage('报销偏好已保存')
}

const saveAiConfig = () => {
  settingsStore.updateAiConfig(tempAiConfig.value)
  showSuccessMessage('AI配置已保存')
}

const saveAppearance = () => {
  settingsStore.updateAppearance(tempAppearance.value)
  showSuccessMessage('外观设置已保存')
}

const testConnection = async () => {
  testing.value = true
  connectionStatus.value = null

  try {
    const success = await settingsStore.testAiConnection()
    connectionStatus.value = {
      success,
      message: success ? '连接测试成功！' : '连接测试失败，请检查配置'
    }
  } catch (error) {
    connectionStatus.value = {
      success: false,
      message: '连接测试出错：' + error.message
    }
  } finally {
    testing.value = false
  }
}

const showSuccessMessage = (message) => {
  // 这里可以集成一个toast组件，暂时用简单的提示
  console.log(message)
}

// 监听store变化，更新临时数据
watch(() => settingsStore.accountInfo, (newVal) => {
  tempAccountInfo.value = { ...newVal }
}, { deep: true })

watch(() => settingsStore.reimbursementPreferences, (newVal) => {
  tempPreferences.value = { ...newVal }
}, { deep: true })

watch(() => settingsStore.aiConfig, (newVal) => {
  tempAiConfig.value = { ...newVal }
}, { deep: true })

watch(() => settingsStore.appearance, (newVal) => {
  tempAppearance.value = { ...newVal }
}, { deep: true })
</script>

<style scoped>
/* 自定义动画 */
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}

.rounded-2xl {
  animation: fadeIn 0.3s ease-out;
}

/* 响应式网格 */
@media (max-width: 1024px) {
  .grid-cols-3 {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 640px) {
  .grid-cols-3,
  .grid-cols-4 {
    grid-template-columns: 1fr;
  }
}
</style>