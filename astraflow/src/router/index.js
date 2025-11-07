import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import ResumeView from '../views/ResumeView.vue'
import JobAnalysisView from '../views/JobAnalysisView.vue'
import CareerProfileView from '../views/CareerProfileView.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: HomeView
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