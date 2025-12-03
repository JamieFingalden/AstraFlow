<template>
  <div :class="'settings-main-container ' + (isDark ? 'dark-theme' : 'light-theme')">
    <!-- Particle background effect -->
    <div class="fixed-container">
      <div :class="'particle-circle particle-1 ' + (isDark ? 'dark-particle' : 'light-particle')"></div>
      <div :class="'particle-circle particle-2 ' + (isDark ? 'dark-particle' : 'light-particle')"></div>
      <div :class="'particle-circle particle-3 ' + (isDark ? 'dark-particle' : 'light-particle')"></div>
    </div>

    <!-- Page Header -->
    <PageHeader
      title="设置中心"
      :show-back-button="true"
      back-to="/"
      :show-theme-toggle="true"
    />

    <!-- Main Content -->
    <main class="settings-main-content">
      <div class="settings-content-wrapper">
        <!-- 左侧导航菜单 -->
        <aside class="settings-sidebar">
          <!-- 移动端下拉菜单 -->
          <div class="mobile-nav-select">
            <select
              v-model="settingsStore.activeTab"
              :class="'nav-select ' + (isDark ? 'dark-theme' : 'light-theme')"
              @change="handleTabChange"
            >
              <option value="account">账户信息</option>
              <option value="preferences">报销偏好</option>
              <option value="ai">AI 接口配置</option>
              <option value="appearance">系统外观</option>
            </select>
          </div>

          <!-- 桌面端垂直菜单 -->
          <nav :class="'settings-nav ' + (isDark ? 'dark-theme' : 'light-theme')">
            <ul class="nav-list">
              <li>
                <button
                  @click="setActiveTab('account')"
                  :class="'nav-item ' + (settingsStore.activeTab === 'account' ? 'active' : '') + ' ' + (isDark ? 'dark-theme' : 'light-theme')"
                >
                  <UserIcon :size="20" />
                  <span>账户信息</span>
                </button>
              </li>
              <li>
                <button
                  @click="setActiveTab('preferences')"
                  :class="'nav-item ' + (settingsStore.activeTab === 'preferences' ? 'active' : '') + ' ' + (isDark ? 'dark-theme' : 'light-theme')"
                >
                  <SettingsIcon :size="20" />
                  <span>报销偏好</span>
                </button>
              </li>
              <li>
                <button
                  @click="setActiveTab('ai')"
                  :class="'nav-item ' + (settingsStore.activeTab === 'ai' ? 'active' : '') + ' ' + (isDark ? 'dark-theme' : 'light-theme')"
                >
                  <CpuIcon :size="20" />
                  <span>AI 接口配置</span>
                </button>
              </li>
              <li>
                <button
                  @click="setActiveTab('appearance')"
                  :class="'nav-item ' + (settingsStore.activeTab === 'appearance' ? 'active' : '') + ' ' + (isDark ? 'dark-theme' : 'light-theme')"
                >
                  <PaletteIcon :size="20" />
                  <span>系统外观</span>
                </button>
              </li>
            </ul>
          </nav>
        </aside>

        <!-- 右侧内容区域 -->
        <section class="settings-content">
          <!-- 账户信息面板 -->
          <div v-if="settingsStore.activeTab === 'account'" :class="'settings-panel ' + (isDark ? 'dark-theme' : 'light-theme')">
            <h2 :class="'panel-title ' + (isDark ? 'dark-theme' : 'light-theme')">账户信息</h2>

            <div class="panel-content">
              <!-- 头像 -->
              <div class="avatar-section">
                <div :class="'avatar-container ' + (isDark ? 'dark-theme' : 'light-theme')">
                  <img :src="settingsStore.accountInfo.avatar" alt="头像" class="avatar-image" />
                </div>
                <button :class="'avatar-upload-btn ' + (isDark ? 'dark-theme' : 'light-theme')">
                  更换头像
                </button>
              </div>

              <!-- 用户信息表单 -->
              <div class="form-grid">
                <div class="form-field">
                  <label :class="'form-label ' + (isDark ? 'dark-theme' : 'light-theme')">用户名</label>
                  <input
                    v-model="tempAccountInfo.username"
                    type="text"
                    :class="'form-input ' + (isDark ? 'dark-theme' : 'light-theme')"
                  />
                </div>
                <div class="form-field">
                  <label :class="'form-label ' + (isDark ? 'dark-theme' : 'light-theme')">邮箱</label>
                  <input
                    v-model="tempAccountInfo.email"
                    type="email"
                    :class="'form-input ' + (isDark ? 'dark-theme' : 'light-theme')"
                  />
                </div>
                <div class="form-field full-width">
                  <label :class="'form-label ' + (isDark ? 'dark-theme' : 'light-theme')">公司名称</label>
                  <input
                    v-model="tempAccountInfo.company"
                    type="text"
                    :class="'form-input ' + (isDark ? 'dark-theme' : 'light-theme')"
                  />
                </div>
              </div>

              <!-- 保存按钮 -->
              <div class="button-row">
                <button
                  @click="saveAccountInfo"
                  :class="'primary-btn ' + (isDark ? 'dark-theme' : 'light-theme')"
                >
                  保存修改
                </button>
              </div>
            </div>
          </div>

          <!-- 报销偏好面板 -->
          <div v-if="settingsStore.activeTab === 'preferences'" :class="'settings-panel ' + (isDark ? 'dark-theme' : 'light-theme')">
            <h2 :class="'panel-title ' + (isDark ? 'dark-theme' : 'light-theme')">报销偏好</h2>

            <div class="panel-content">
              <!-- 自动分类开关 -->
              <div class="toggle-section">
                <div class="toggle-info">
                  <h3 :class="'toggle-title ' + (isDark ? 'dark-theme' : 'light-theme')">自动分类</h3>
                  <p :class="'toggle-description ' + (isDark ? 'dark-theme' : 'light-theme')">自动识别账单分类</p>
                </div>
                <label class="switch">
                  <input
                    type="checkbox"
                    v-model="tempPreferences.autoCategory"
                    class="switch-input"
                  />
                  <span class="switch-slider"></span>
                </label>
              </div>

              <!-- 报销上限提醒 -->
              <div class="form-field">
                <label :class="'form-label ' + (isDark ? 'dark-theme' : 'light-theme')">报销上限提醒阈值</label>
                <div class="input-with-unit">
                  <input
                    v-model.number="tempPreferences.spendingLimit"
                    type="number"
                    :class="'form-input ' + (isDark ? 'dark-theme' : 'light-theme')"
                  />
                  <span :class="'unit-label ' + (isDark ? 'dark-theme' : 'light-theme')">元</span>
                </div>
              </div>

              <!-- 货币单位 -->
              <div class="form-field">
                <label :class="'form-label ' + (isDark ? 'dark-theme' : 'light-theme')">货币单位</label>
                <select
                  v-model="tempPreferences.currency"
                  :class="'form-select ' + (isDark ? 'dark-theme' : 'light-theme')"
                >
                  <option value="CNY">人民币 (¥)</option>
                  <option value="USD">美元 ($)</option>
                </select>
              </div>

              <!-- 保存按钮 -->
              <div class="button-row">
                <button
                  @click="savePreferences"
                  :class="'primary-btn ' + (isDark ? 'dark-theme' : 'light-theme')"
                >
                  保存偏好设置
                </button>
              </div>
            </div>
          </div>

          <!-- AI接口配置面板 -->
          <div v-if="settingsStore.activeTab === 'ai'" :class="'settings-panel ' + (isDark ? 'dark-theme' : 'light-theme')">
            <h2 :class="'panel-title ' + (isDark ? 'dark-theme' : 'light-theme')">AI 接口配置</h2>

            <div class="panel-content">
              <!-- Flask接口URL -->
              <div class="form-field">
                <label :class="'form-label ' + (isDark ? 'dark-theme' : 'light-theme')">Flask 识别接口 URL</label>
                <input
                  v-model="tempAiConfig.flaskUrl"
                  type="url"
                  placeholder="http://localhost:5000/api/ocr"
                  :class="'form-input ' + (isDark ? 'dark-theme' : 'light-theme')"
                />
              </div>

              <!-- API密钥 -->
              <div class="form-field">
                <label :class="'form-label ' + (isDark ? 'dark-theme' : 'light-theme')">API 密钥</label>
                <input
                  v-model="tempAiConfig.apiKey"
                  type="password"
                  placeholder="输入您的API密钥"
                  :class="'form-input ' + (isDark ? 'dark-theme' : 'light-theme')"
                />
              </div>

              <!-- 操作按钮 -->
              <div class="button-row">
                <button
                  @click="testConnection"
                  :disabled="testing"
                  :class="'secondary-btn ' + (testing ? 'disabled' : '') + ' ' + (isDark ? 'dark-theme' : 'light-theme')"
                >
                  {{ testing ? '测试中...' : '测试连接' }}
                </button>
                <button
                  @click="saveAiConfig"
                  :class="'primary-btn ' + (isDark ? 'dark-theme' : 'light-theme')"
                >
                  保存配置
                </button>
              </div>

              <!-- 连接状态提示 -->
              <div v-if="connectionStatus" :class="'status-message ' + (connectionStatus.success ? 'success' : 'error') + ' ' + (isDark ? 'dark-theme' : 'light-theme')">
                {{ connectionStatus.message }}
              </div>
            </div>
          </div>

          <!-- 系统外观面板 -->
          <div v-if="settingsStore.activeTab === 'appearance'" :class="'settings-panel ' + (isDark ? 'dark-theme' : 'light-theme')">
            <h2 :class="'panel-title ' + (isDark ? 'dark-theme' : 'light-theme')">系统外观</h2>

            <div class="panel-content">
              <!-- 主题切换 -->
              <div class="theme-section">
                <label :class="'form-label ' + (isDark ? 'dark-theme' : 'light-theme')">主题模式</label>
                <div class="theme-options">
                  <button
                    v-for="theme in themeOptions"
                    :key="theme.value"
                    @click="tempAppearance.theme = theme.value"
                    :class="'theme-option ' + (tempAppearance.theme === theme.value ? 'active' : '') + ' ' + (isDark ? 'dark-theme' : 'light-theme')"
                  >
                    <component :is="theme.icon" :size="20" :class="'theme-option-icon ' + (isDark ? 'dark-theme' : 'light-theme')" />
                    <span :class="'theme-option-label ' + (isDark ? 'dark-theme' : 'light-theme')">{{ theme.label }}</span>
                  </button>
                </div>
              </div>

              <!-- 主色调选择 -->
              <div class="color-section">
                <label :class="'form-label ' + (isDark ? 'dark-theme' : 'light-theme')">主色调</label>
                <div class="color-options">
                  <button
                    v-for="color in colorOptions"
                    :key="color.value"
                    @click="tempAppearance.primaryColor = color.value"
                    :class="'color-option ' + (tempAppearance.primaryColor === color.value ? 'active' : '') + ' ' + (isDark ? 'dark-theme' : 'light-theme')"
                  >
                    <div :class="'color-swatch ' + color.value"></div>
                    <span :class="'color-label ' + (isDark ? 'dark-theme' : 'light-theme')">{{ color.label }}</span>
                  </button>
                </div>
              </div>

              <!-- 保存按钮 -->
              <div class="button-row">
                <button
                  @click="saveAppearance"
                  :class="'primary-btn ' + (isDark ? 'dark-theme' : 'light-theme')"
                >
                  保存外观设置
                </button>
              </div>
            </div>
          </div>
        </section>
      </div>
    </main>

    <!-- Page Footer -->
    <PageFooter />
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useTheme } from '../composables/useTheme'
import { useSettingsStore } from '../stores/settings'
import PageHeader from '../components/ui/PageHeader.vue'
import PageFooter from '../components/ui/PageFooter.vue'
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
/* 主容器 */
.settings-main-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  transition: background-color 0.3s ease, color 0.3s ease;
}

.settings-main-container.dark-theme {
  background-color: #111827;
}

.settings-main-container.light-theme {
  background: linear-gradient(135deg, #eff6ff, #f0f9ff, #e0e7ff);
}

/* 主内容 */
.settings-main-content {
  position: relative;
  z-index: 10;
  flex: 1;
  max-width: 80rem;
  width: 100%;
  margin: 0 auto;
  padding: 2rem 1rem;
}

@media (min-width: 640px) {
  .settings-main-content {
    padding: 2rem 1.5rem;
  }
}

@media (min-width: 1024px) {
  .settings-main-content {
    padding: 2rem 2rem;
  }
}

/* 固定容器 */
.fixed-container {
  position: fixed;
  inset: 0;
  overflow: hidden;
  pointer-events: none;
}

/* 粒子圆圈 */
.particle-circle {
  position: absolute;
  width: 20rem;
  height: 20rem;
  border-radius: 50%;
  mix-blend-mode: multiply;
  filter: blur(6rem);
  animation: pulse 4s ease-in-out infinite;
}

.particle-1 {
  top: -10rem;
  right: -10rem;
}

.particle-2 {
  bottom: -10rem;
  left: -10rem;
}

.particle-3 {
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.particle-circle.dark-particle {
  background-color: rgba(6, 182, 212, 0.3);
}

.particle-circle.light-particle {
  background-color: #06b6d4;
}

@keyframes pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.8;
  }
}

/* 头部 */
.settings-header {
  position: relative;
  z-index: 10;
  backdrop-filter: blur(12px);
  border-bottom: 1px solid;
  transition: all 0.3s ease;
  height: 4rem;
}

.settings-header.dark-theme {
  background-color: rgba(31, 41, 55, 0.7);
  border-color: rgba(55, 65, 81, 0.5);
}

.settings-header.light-theme {
  background-color: rgba(255, 255, 255, 0.7);
  border-color: rgba(229, 231, 235, 0.5);
}

.container {
  padding: 0 1rem;
  max-width: 80rem;
  margin: 0 auto;
}

@media (min-width: 640px) {
  .container {
    padding: 0 1.5rem;
  }
}

@media (min-width: 1024px) {
  .container {
    padding: 0 2rem;
  }
}

.header-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 4rem;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 1rem;
}

/* 返回按钮 */
.back-button {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 0.75rem;
  border-radius: 0.5rem;
  transition: all 0.2s ease;
  text-decoration: none;
}

.back-button.dark-theme {
  color: #d1d5db;
}

.back-button.light-theme {
  color: #4b5563;
}

.back-button:hover.dark-theme {
  background-color: rgba(55, 65, 81, 0.5);
  color: #ffffff;
}

.back-button:hover.light-theme {
  background-color: #f3f4f6;
  color: #1f2937;
}

.back-button-text {
  display: none;
}

@media (min-width: 640px) {
  .back-button-text {
    display: inline;
  }
}

/* 品牌文字 */
.brand-text {
  font-size: 1.5rem;
  font-weight: 700;
  background: linear-gradient(135deg, #06b6d4, #3b82f6);
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
}

/* 页面标题 */
.page-title {
  font-size: 1.25rem;
  font-weight: 600;
  transition: color 0.3s ease;
}

.page-title.dark-theme {
  color: #ffffff;
}

.page-title.light-theme {
  color: #1f2937;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 1rem;
}

/* 主题切换按钮 */
.theme-toggle-btn {
  padding: 0.5rem;
  border-radius: 0.5rem;
  transition: all 0.2s ease;
  border: none;
  cursor: pointer;
}

.theme-toggle-btn.dark-theme {
  background-color: rgba(55, 65, 81, 0.5);
  color: #d1d5db;
}

.theme-toggle-btn.light-theme {
  background-color: #f3f4f6;
  color: #374151;
}

.theme-toggle-btn:hover.dark-theme {
  background-color: rgba(75, 85, 99, 0.5);
}

.theme-toggle-btn:hover.light-theme {
  background-color: #e5e7eb;
}

.settings-content-wrapper {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  width: 100%;
  max-width: 100%;
}

@media (min-width: 1024px) {
  .settings-content-wrapper {
    flex-direction: row;
  }
}

/* 侧边栏 */
.settings-sidebar {
  flex-shrink: 0;
  width: 100%;
}

@media (min-width: 1024px) {
  .settings-sidebar {
    width: 16rem;
    flex-shrink: 0;
  }
}

/* 移动端导航选择 */
.mobile-nav-select {
  margin-bottom: 1.5rem;
}

@media (min-width: 1024px) {
  .mobile-nav-select {
    display: none;
  }
}

.nav-select {
  width: 100%;
  padding: 0.75rem 1rem;
  border-radius: 0.5rem;
  border: 1px solid;
  transition: all 0.3s ease;
  background-color: transparent;
  font-size: 1rem;
}

.nav-select.dark-theme {
  background-color: rgba(31, 41, 55, 0.7);
  border-color: #374151;
  color: #ffffff;
}

.nav-select.light-theme {
  background-color: rgba(255, 255, 255, 0.7);
  border-color: #d1d5db;
  color: #1f2937;
}

/* 设置导航 */
.settings-nav {
  display: none;
  border-radius: 1rem;
  padding: 1rem;
  border: 1px solid;
  transition: all 0.3s ease;
}

@media (min-width: 1024px) {
  .settings-nav {
    display: block;
  }
}

.settings-nav.dark-theme {
  background-color: rgba(31, 41, 55, 0.7);
  border-color: rgba(55, 65, 81, 0.5);
}

.settings-nav.light-theme {
  background-color: rgba(255, 255, 255, 0.7);
  border-color: rgba(229, 231, 235, 0.5);
}

.nav-list {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  list-style: none;
  margin: 0;
  padding: 0;
}

.nav-item {
  width: 100%;
  text-align: left;
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.75rem 1rem;
  border-radius: 0.5rem;
  transition: all 0.2s ease;
  border: none;
  cursor: pointer;
}

.nav-item.dark-theme {
  color: #d1d5db;
}

.nav-item.light-theme {
  color: #4b5563;
}

.nav-item:hover.dark-theme {
  background-color: rgba(55, 65, 81, 0.5);
}

.nav-item:hover.light-theme {
  background-color: #f3f4f6;
}

.nav-item.active {
  border-left: 4px solid;
  font-weight: 500;
}

.nav-item.active.dark-theme {
  background-color: rgba(30, 58, 138, 0.5);
  color: #93c5fd;
  border-color: #60a5fa;
}

.nav-item.active.light-theme {
  background-color: #eff6ff;
  color: #1d4ed8;
  border-color: #3b82f6;
}

/* 设置内容 */
.settings-content {
  flex: 1;
  max-width: 60rem;
  width: 100%;
  margin: 0 auto;
}

@media (max-width: 1024px) {
  .settings-content {
    max-width: 100%;
    padding: 0 1rem;
  }
}

@media (max-width: 768px) {
  .settings-content {
    padding: 0 0.5rem;
  }
}

/* 设置面板 */
.settings-panel {
  border-radius: 1rem;
  padding: 1.5rem;
  border: 1px solid;
  transition: all 0.3s ease;
  backdrop-filter: blur(12px);
  width: 100%;
  max-width: 100%;
}

.settings-panel.dark-theme {
  background-color: rgba(31, 41, 55, 0.7);
  border-color: rgba(55, 65, 81, 0.5);
}

.settings-panel.light-theme {
  background-color: rgba(255, 255, 255, 0.7);
  border-color: rgba(229, 231, 235, 0.5);
}

/* 面板标题 */
.panel-title {
  font-size: 1.5rem;
  font-weight: 700;
  margin-bottom: 1.5rem;
}

.panel-title.dark-theme {
  color: #ffffff;
}

.panel-title.light-theme {
  color: #1f2937;
}

/* 面板内容 */
.panel-content {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

/* 头像区域 */
.avatar-section {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.avatar-container {
  width: 5rem;
  height: 5rem;
  border-radius: 50%;
  overflow: hidden;
  border: 4px solid;
}

.avatar-container.dark-theme {
  border-color: #4b5563;
}

.avatar-container.light-theme {
  border-color: #e5e7eb;
}

.avatar-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.avatar-upload-btn {
  padding: 0.5rem 1rem;
  border-radius: 0.5rem;
  transition: all 0.2s ease;
  border: none;
  cursor: pointer;
}

.avatar-upload-btn.dark-theme {
  background-color: #374151;
  color: #d1d5db;
}

.avatar-upload-btn.light-theme {
  background-color: #f3f4f6;
  color: #374151;
}

.avatar-upload-btn:hover.dark-theme {
  background-color: #4b5563;
}

.avatar-upload-btn:hover.light-theme {
  background-color: #e5e7eb;
}

/* 表单网格 */
.form-grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1.5rem;
}

@media (min-width: 768px) {
  .form-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

/* 表单字段 */
.form-field {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.form-field.full-width {
  grid-column: span 2;
}

/* 表单标签 */
.form-label {
  font-size: 0.875rem;
  font-weight: 500;
  display: block;
}

.form-label.dark-theme {
  color: #d1d5db;
}

.form-label.light-theme {
  color: #4b5563;
}

/* 表单输入 */
.form-input {
  width: 100%;
  padding: 0.5rem 1rem;
  border-radius: 0.5rem;
  border: 1px solid;
  transition: all 0.3s ease;
  background-color: transparent;
}

.form-input.dark-theme {
  background-color: rgba(55, 65, 81, 0.5);
  border-color: #4b5563;
  color: #ffffff;
}

.form-input.light-theme {
  background-color: rgba(255, 255, 255, 0.5);
  border-color: #d1d5db;
  color: #1f2937;
}

.form-input:focus {
  outline: none;
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.5);
}

/* 表单选择 */
.form-select {
  width: 100%;
  padding: 0.5rem 1rem;
  border-radius: 0.5rem;
  border: 1px solid;
  transition: all 0.3s ease;
  background-color: transparent;
  font-size: 1rem;
}

.form-select.dark-theme {
  background-color: rgba(55, 65, 81, 0.5);
  border-color: #4b5563;
  color: #ffffff;
}

.form-select.light-theme {
  background-color: rgba(255, 255, 255, 0.5);
  border-color: #d1d5db;
  color: #1f2937;
}

/* 按钮行 */
.button-row {
  display: flex;
  justify-content: flex-end;
}

/* 主要按钮 */
.primary-btn {
  padding: 0.5rem 1.5rem;
  border-radius: 0.5rem;
  font-weight: 600;
  transition: all 0.3s ease;
  transform: scale(1);
  border: none;
  cursor: pointer;
}

.primary-btn.dark-theme {
  background-color: #2563eb;
  color: #ffffff;
}

.primary-btn.light-theme {
  background-color: #3b82f6;
  color: #ffffff;
}

.primary-btn:hover {
  transform: scale(1.05);
}

.primary-btn:hover.dark-theme {
  background-color: #3b82f6;
}

.primary-btn:hover.light-theme {
  background-color: #2563eb;
}

/* 次要按钮 */
.secondary-btn {
  padding: 0.5rem 1.5rem;
  border-radius: 0.5rem;
  font-weight: 600;
  transition: all 0.3s ease;
  transform: scale(1);
  border: none;
  cursor: pointer;
}

.secondary-btn.dark-theme {
  background-color: #16a34a;
  color: #ffffff;
}

.secondary-btn.light-theme {
  background-color: #22c55e;
  color: #ffffff;
}

.secondary-btn:hover:not(.disabled) {
  transform: scale(1.05);
}

.secondary-btn:hover:not(.disabled).dark-theme {
  background-color: #22c55e;
}

.secondary-btn:hover:not(.disabled).light-theme {
  background-color: #16a34a;
}

.secondary-btn.disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

/* 切换部分 */
.toggle-section {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

@media (max-width: 640px) {
  .toggle-section {
    flex-direction: column;
    align-items: stretch;
    gap: 1rem;
  }
}

.toggle-info {
  flex: 1;
}

.toggle-title {
  font-weight: 500;
}

.toggle-title.dark-theme {
  color: #ffffff;
}

.toggle-title.light-theme {
  color: #1f2937;
}

.toggle-description {
  font-size: 0.875rem;
  margin-top: 0.25rem;
}

.toggle-description.dark-theme {
  color: #9ca3af;
}

.toggle-description.light-theme {
  color: #6b7280;
}

/* 开关组件 */
.switch {
  position: relative;
  display: inline-block;
  width: 2.75rem;
  height: 1.5rem;
}

.switch-input {
  opacity: 0;
  width: 0;
  height: 0;
}

.switch-slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #e5e7eb;
  transition: .4s;
  border-radius: 9999px;
}

[data-theme="dark"] .switch-slider {
  background-color: #374151;
}

.switch-slider::after {
  position: absolute;
  content: "";
  height: 1.25rem;
  width: 1.25rem;
  left: 2px;
  bottom: 2px;
  background-color: white;
  transition: .4s;
  border-radius: 9999px;
}

[data-theme="dark"] .switch-slider::after {
  background-color: #f9fafb;
}

.switch-input:checked + .switch-slider {
  background-color: #3b82f6;
}

.switch-input:checked + .switch-slider::after {
  transform: translateX(1.125rem);
}

/* 带单位的输入 */
.input-with-unit {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.unit-label {
  font-size: 0.875rem;
}

.unit-label.dark-theme {
  color: #9ca3af;
}

.unit-label.light-theme {
  color: #6b7280;
}

/* 状态消息 */
.status-message {
  padding: 1rem;
  border-radius: 0.5rem;
}

.status-message.success {
  background-color: rgba(22, 163, 74, 0.1);
  color: #16a34a;
}

.status-message.error {
  background-color: rgba(239, 68, 68, 0.1);
  color: #ef4444;
}

.status-message.success.dark-theme {
  background-color: rgba(22, 163, 74, 0.3);
  color: #86efac;
}

.status-message.error.dark-theme {
  background-color: rgba(239, 68, 68, 0.3);
  color: #fca5a5;
}

/* 主题部分 */
.theme-section,
.color-section {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

/* 主题选项 */
.theme-options {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 0.75rem;
}

/* 主题选项按钮 */
.theme-option {
  padding: 0.75rem;
  border-radius: 0.5rem;
  border: 2px solid;
  transition: all 0.2s ease;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.25rem;
  background: none;
  cursor: pointer;
}

.theme-option.dark-theme {
  border-color: #374151;
}

.theme-option.light-theme {
  border-color: #d1d5db;
}

.theme-option.active {
  border-color: #3b82f6;
}

.theme-option.active.dark-theme {
  background-color: rgba(30, 64, 175, 0.1);
}

.theme-option.active.light-theme {
  background-color: #eff6ff;
}

.theme-option-icon {
  margin: 0 auto 0.25rem auto;
}

.theme-option-icon.dark-theme {
  color: #d1d5db;
}

.theme-option-icon.light-theme {
  color: #4b5563;
}

.theme-option-label {
  font-size: 0.75rem;
}

.theme-option-label.dark-theme {
  color: #d1d5db;
}

.theme-option-label.light-theme {
  color: #4b5563;
}

/* 颜色选项 */
.color-options {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 0.75rem;
}

/* 颜色选项按钮 */
.color-option {
  padding: 0.75rem;
  border-radius: 0.5rem;
  border: 2px solid;
  transition: all 0.2s ease;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.25rem;
  background: none;
  cursor: pointer;
}

.color-option.dark-theme {
  border-color: #374151;
}

.color-option.light-theme {
  border-color: #d1d5db;
}

.color-option.active {
  border-color: #1f2937;
}

[data-theme="dark"] .color-option.active {
  border-color: #ffffff;
}

/* 颜色色块 */
.color-swatch {
  width: 1.5rem;
  height: 1.5rem;
  border-radius: 50%;
}

.color-swatch.blue { background-color: #3b82f6; }
.color-swatch.green { background-color: #22c55e; }
.color-swatch.purple { background-color: #8b5cf6; }
.color-swatch.gray { background-color: #6b7280; }

.color-label {
  font-size: 0.75rem;
}

.color-label.dark-theme {
  color: #d1d5db;
}

.color-label.light-theme {
  color: #4b5563;
}

/* 页脚 */

/* 响应式网格 */
@media (max-width: 1024px) {
  .theme-options {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 640px) {
  .theme-options,
  .color-options {
    grid-template-columns: 1fr;
  }
}
</style>