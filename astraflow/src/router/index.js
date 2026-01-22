import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/auth/LoginView.vue'
import LogoutView from '../views/auth/LogoutView.vue'
import Dashboard from '../views/dashboard/Dashboard.vue'
import InvoiceUpload from '../views/InvoiceUpload.vue'

import BillManagement from '../views/BillManagement.vue'
import ReimbursementStatistics from '../views/ReimbursementStatistics.vue'
import SettingsCenter from '../views/SettingsCenter.vue'
import UserManagement from '../views/UserManagement.vue'
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
    path: '/logout',
    name: 'Logout',
    component: LogoutView,
    meta: {
      requiresAuth: true // Only authenticated users can logout
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
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard,
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
  },
  {
    path: '/users',
    name: 'UserManagement',
    component: UserManagement,
    meta: {
      requiresAuth: true,
      requiresTenantAdmin: true
    }
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

// Navigation guards for authentication
router.beforeEach(async (to, from, next) => {
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

  // Handle routes that require tenant admin access
  if (to.meta.requiresTenantAdmin && isAuthenticated) {
    if (!userStore.isTenantAdmin() && !userStore.isSystemAdmin()) {
      // Redirect to dashboard or show error page
      next({ name: 'Dashboard' }) // Or you could redirect to an error page
      return
    }
  }

  // Handle routes that should be hidden for authenticated users
  if (to.meta.hideForAuth && isAuthenticated) {
    // Redirect to dashboard or home page
    next({ name: 'Dashboard' })
    return
  }

  // Note: Redirect query is handled by LoginForm component after successful login
  // We don't automatically redirect here to allow users to see the login page

  next()
})

export default router