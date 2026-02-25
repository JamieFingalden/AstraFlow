import { createRouter, createWebHistory } from 'vue-router'
import MainLayout from '../layout/MainLayout.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'Login',
      component: () => import('../views/login/index.vue'),
      meta: { title: '登录' }
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
          meta: { title: '仪表盘' }
        },
        {
          path: 'upload',
          name: 'Upload',
          component: () => import('../views/upload/index.vue'),
          meta: { title: '上传单据' }
        },
        {
          path: 'my-invoices',
          name: 'MyInvoices',
          component: () => import('../views/my-invoices/index.vue'),
          meta: { title: '我的票夹' }
        },
        // Audit Module
        {
          path: 'audit/tasks',
          name: 'Audit',
          component: () => import('../views/audit/task-pool/index.vue'),
          meta: { title: '审核任务池' }
        },
        {
          path: 'audit/tasks/:id',
          name: 'AuditDetail',
          component: () => import('../views/audit/task-pool/detail.vue'),
          meta: { title: '发票审核', hideMenu: true }
        },
        {
          path: 'audit/settlement',
          name: 'Settlement',
          component: () => import('../views/audit/settlement/index.vue'),
          meta: { title: '结算中心' }
        },
        // Archive Module
        {
          path: 'archive',
          name: 'Archive',
          component: () => import('../views/archive/index.vue'),
          meta: { title: '历史归档' }
        },
        // System Management
        {
          path: 'system/employees',
          name: 'System',
          component: () => import('../views/system/employee/index.vue'),
          meta: { title: '员工管理' }
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
