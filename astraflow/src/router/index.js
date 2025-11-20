import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/auth/LoginView.vue'
import Dashboard from '../views/dashboard/Dashboard.vue'
import InvoiceUpload from '../views/InvoiceUpload.vue'
import AnalysisPage from '../views/AnalysisPage.vue'
import BillManagement from '../views/BillManagement.vue'
import ReimbursementStatistics from '../views/ReimbursementStatistics.vue'
import SettingsCenter from '../views/SettingsCenter.vue'
import { useUserStore } from '../stores/userStore'

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: LoginView,
    meta: {
      requiresAuth: false,
      hideForAuth: true // Redirect authenticated users away from login
    }
  },
  {
    path: '/',
    name: 'Home',
    component: HomeView,
    meta: { requiresAuth: false }
  },
  {
    path: '/upload',
    name: 'InvoiceUpload',
    component: InvoiceUpload,
    meta: { requiresAuth: true }
  },
  {
    path: '/bills',
    name: 'BillManagement',
    component: BillManagement,
    meta: { requiresAuth: true }
  },
  {
    path: '/visualization',
    name: 'Visualization',
    component: Dashboard,
    meta: { requiresAuth: true }
  },
  {
    path: '/ai-result',
    name: 'AIResult',
    component: AnalysisPage,
    meta: { requiresAuth: true }
  },
  {
    path: '/statistics',
    name: 'ReimbursementStatistics',
    component: ReimbursementStatistics,
    meta: { requiresAuth: true }
  },
  {
    path: '/settings',
    name: 'SettingsCenter',
    component: SettingsCenter,
    meta: { requiresAuth: true }
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

// Navigation guards for authentication
router.beforeEach((to, from, next) => {
  const userStore = useUserStore()
  const isAuthenticated = userStore.isAuthenticated

  // Handle routes that require authentication
  if (to.meta.requiresAuth && !isAuthenticated) {
    // Redirect to login page with redirect query
    next({
      name: 'Login',
      query: { redirect: to.fullPath }
    })
    return
  }

  // Handle routes that should be hidden for authenticated users
  if (to.meta.hideForAuth && isAuthenticated) {
    // Redirect to dashboard or home page
    next({ name: 'Visualization' })
    return
  }

  // Handle redirect query after successful login
  if (to.name === 'Login' && isAuthenticated && to.query.redirect) {
    next(to.query.redirect)
    return
  }

  next()
})

export default router