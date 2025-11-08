import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ResumeView from '../views/ResumeView.vue'
import JobAnalysisView from '../views/JobAnalysisView.vue'
import CareerProfileView from '../views/CareerProfileView.vue'
import Dashboard from '../views/dashboard/Dashboard.vue'
import InvoiceUpload from '../views/InvoiceUpload.vue'
import AnalysisPage from '../views/AnalysisPage.vue'
import BillManagement from '../views/BillManagement.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: HomeView
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard
  },
  {
    path: '/upload',
    name: 'InvoiceUpload',
    component: InvoiceUpload
  },
  {
    path: '/analysis',
    name: 'AnalysisPage',
    component: AnalysisPage
  },
  {
    path: '/bills',
    name: 'BillManagement',
    component: BillManagement
  },
  {
    path: '/resume',
    name: 'Resume',
    component: ResumeView
  },
  {
    path: '/job-analysis',
    name: 'JobAnalysis',
    component: JobAnalysisView
  },
  {
    path: '/profile',
    name: 'CareerProfile',
    component: CareerProfileView
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router