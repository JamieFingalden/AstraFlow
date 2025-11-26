import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import { NButton, NInput, NSelect, NModal, NCard, NForm, NFormItem, NDatePicker, NInputNumber, NRadio, NRadioGroup, NSpace, NLayout, NLayoutHeader, NLayoutSider, NLayoutContent, NLayoutFooter, NMenu, NIcon, NAlert, NMessageProvider, NDialogProvider, NNotificationProvider, NConfigProvider, zhCN, dateZhCN } from 'naive-ui'
import './style.css'
import App from './App.vue'
import router from './router'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)
app.use(ElementPlus)

// Naive UI global configuration
app.use(NMessageProvider)
app.use(NDialogProvider)
app.use(NNotificationProvider)

// Global scrollbar behavior
let scrollTimeout

const handleScroll = () => {
  // Add scrolling class to show scrollbar
  document.body.classList.add('scrolling')

  // Clear existing timeout
  clearTimeout(scrollTimeout)

  // Remove scrolling class after scrolling stops
  scrollTimeout = setTimeout(() => {
    document.body.classList.remove('scrolling')
  }, 1000) // Hide scrollbar after 1 second of inactivity
}

// Listen for scroll events globally
window.addEventListener('scroll', handleScroll, { passive: true })

app.mount('#app')
