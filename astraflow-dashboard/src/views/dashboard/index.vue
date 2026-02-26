<template>
  <div class="h-full flex flex-col gap-6">
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-2xl font-bold text-slate-900 tracking-tight">数据仪表盘</h1>
        <p class="text-slate-500 mt-1 text-sm">根据当前角色展示报销统计数据。</p>
      </div>
      <el-button type="primary" class="!bg-indigo-600 !border-indigo-600 hover:!bg-indigo-700" @click="loadStats">
        刷新
      </el-button>
    </div>

    <el-row :gutter="16">
      <el-col :xs="24" :sm="12" :lg="8">
        <el-card shadow="hover" class="h-full">
          <el-statistic :title="cardOne.title" :value="cardOne.value">
            <template #prefix>
              <el-icon><Document /></el-icon>
            </template>
          </el-statistic>
        </el-card>
      </el-col>

      <el-col :xs="24" :sm="12" :lg="8">
        <el-card shadow="hover" class="h-full">
          <el-statistic :title="cardTwo.title" :value="cardTwo.value" :precision="2">
            <template #prefix>
              <span>¥</span>
            </template>
          </el-statistic>
        </el-card>
      </el-col>

      <el-col :xs="24" :sm="12" :lg="8">
        <el-card shadow="hover" class="h-full">
          <el-statistic :title="cardThree.title" :value="cardThree.value" :precision="cardThree.precision">
            <template #prefix>
              <span>{{ cardThree.prefix }}</span>
            </template>
          </el-statistic>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { Document } from '@element-plus/icons-vue'
import { getDashboardStats } from '../../api/dashboard'
import type { CompanyDashboardStats, DashboardStats, EmployeeDashboardStats } from '../../api/dashboard'
import { useUserStore } from '../../stores/userStore'

const userStore = useUserStore()
const isEmployee = computed(() => userStore.role_key === 'employee')

const employeeStats = ref<EmployeeDashboardStats>({
  unconfirmed: 0,
  processing_amount: 0,
  total_reimbursed_amount: 0,
})

const companyStats = ref<CompanyDashboardStats>({
  company_pending_count: 0,
  monthly_paid_amount: 0,
  company_to_pay_count: 0,
})

const cardOne = computed(() => {
  if (isEmployee.value) {
    return { title: '待确认单据数', value: employeeStats.value.unconfirmed }
  }
  return { title: '全公司待审核数', value: companyStats.value.company_pending_count }
})

const cardTwo = computed(() => {
  if (isEmployee.value) {
    return { title: '处理中金额', value: employeeStats.value.processing_amount }
  }
  return { title: '本月已打款总额', value: companyStats.value.monthly_paid_amount }
})

const cardThree = computed(() => {
  if (isEmployee.value) {
    return {
      title: '累计报销金额',
      value: employeeStats.value.total_reimbursed_amount,
      prefix: '¥',
      precision: 2,
    }
  }
  return {
    title: '全公司待打款数',
    value: companyStats.value.company_to_pay_count,
    prefix: '',
    precision: 0,
  }
})

/**
 * 加载仪表盘统计数据。
 * 入参：无（由后端根据当前登录用户角色自动计算口径）。
 * 出参：无（结果写入页面响应式状态）。
 */
const loadStats = async () => {
  try {
    const data: DashboardStats = await getDashboardStats()
    if ('unconfirmed' in data) {
      employeeStats.value = {
        unconfirmed: Number(data.unconfirmed || 0),
        processing_amount: Number(data.processing_amount || 0),
        total_reimbursed_amount: Number(data.total_reimbursed_amount || 0),
      }
      return
    }

    companyStats.value = {
      company_pending_count: Number(data.company_pending_count || 0),
      monthly_paid_amount: Number(data.monthly_paid_amount || 0),
      company_to_pay_count: Number(data.company_to_pay_count || 0),
    }
  } catch (e: any) {
    ElMessage.error(e?.response?.data?.message || '加载仪表盘统计失败')
  }
}

onMounted(() => {
  loadStats()
})
</script>
