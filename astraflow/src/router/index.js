import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Dashboard from '../views/dashboard/Dashboard.vue'
import InvoiceUpload from '../views/InvoiceUpload.vue'
import AnalysisPage from '../views/AnalysisPage.vue'
import BillManagement from '../views/BillManagement.vue'
import ReimbursementStatistics from '../views/ReimbursementStatistics.vue'
import SettingsCenter from '../views/SettingsCenter.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: HomeView
  },
  {
    path: '/upload',
    name: 'InvoiceUpload',
    component: InvoiceUpload
  },
  {
    path: '/bills',
    name: 'BillManagement',
    component: BillManagement
  },
  {
    path: '/visualization',
    name: 'Visualization',
    component: Dashboard
  },
  {
    path: '/ai-result',
    name: 'AIResult',
    component: AnalysisPage
  },
  {
    path: '/statistics',
    name: 'ReimbursementStatistics',
    component: ReimbursementStatistics
  },
  {
    path: '/settings',
    name: 'SettingsCenter',
    component: SettingsCenter
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router