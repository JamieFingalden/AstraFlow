import { createRouter, createWebHistory } from 'vue-router'
import MainLayout from '../layout/MainLayout.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'Login',
      component: () => import('../views/login/index.vue'),
      meta: { title: 'Login' }
    },
    {
      path: '/',
      component: MainLayout,
      redirect: '/dashboard',
      children: [
        {
          path: 'dashboard',
          name: 'Dashboard',
          component: () => import('../views/dashboard/index.vue'),
          meta: { title: 'Executive Dashboard' }
        },
        // Audit Module
        {
          path: 'audit',
          name: 'Audit',
          component: () => import('../views/audit/task-pool/index.vue'),
          meta: { title: 'Audit Task Pool' }
        },
        {
          path: 'audit/tasks/:id',
          name: 'AuditDetail',
          component: () => import('../views/audit/task-pool/detail.vue'),
          meta: { title: 'Invoice Audit', hideMenu: true }
        },
        {
          path: 'audit/settlement',
          name: 'Settlement',
          component: () => import('../views/audit/settlement/index.vue'),
          meta: { title: 'Finance Settlement' }
        },
        // Archive Module
        {
          path: 'archive',
          name: 'Archive',
          component: () => import('../views/archive/index.vue'),
          meta: { title: 'Invoice Archive' }
        },
        // System Management
        {
          path: 'system/employees',
          name: 'System',
          component: () => import('../views/system/employee/index.vue'),
          meta: { title: 'Team Management' }
        }
      ]
    }
  ]
})

// Simple Auth Guard (Optional logic)
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  if (to.path !== '/login' && !token) {
    next('/login')
  } else {
    next()
  }
})

export default router