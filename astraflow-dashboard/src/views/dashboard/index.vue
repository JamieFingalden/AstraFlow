<template>
  <div class="h-full flex flex-col gap-6 page-shell p-6 md:p-7 dashboard-shell">
    <div class="page-header">
      <div>
        <h1 class="af-title text-2xl md:text-3xl font-semibold text-slate-900">数据仪表盘</h1>
        <p class="page-subtitle">根据当前角色展示报销统计数据。</p>
      </div>
      <el-button type="primary" class="af-button-primary" @click="loadStats">
        <el-icon class="mr-2"><RefreshRight /></el-icon>
        刷新
      </el-button>
    </div>

    <div class="dashboard-banner">
      <div>
        <div class="banner-title">{{ bannerTitle }}</div>
        <div class="banner-subtitle">{{ bannerSubtitle }}</div>
      </div>
      <div class="banner-meta">
        <div class="meta-label">更新时间</div>
        <div class="meta-value">{{ lastUpdatedText }}</div>
      </div>
    </div>

    <el-row :gutter="16">
      <el-col :xs="24" :sm="12" :lg="8">
        <el-card shadow="never" class="h-full data-card stat-card">
          <div class="stat-head">
            <span class="stat-label">{{ cardOne.title }}</span>
            <el-icon class="stat-icon stat-icon-blue"><Document /></el-icon>
          </div>
          <div class="stat-value">{{ numberDisplay(cardOne.value) }}</div>
          <div class="stat-foot">当前关键待处理任务</div>
        </el-card>
      </el-col>

      <el-col :xs="24" :sm="12" :lg="8">
        <el-card shadow="never" class="h-full data-card stat-card">
          <div class="stat-head">
            <span class="stat-label">{{ cardTwo.title }}</span>
            <el-icon class="stat-icon stat-icon-green"><TrendCharts /></el-icon>
          </div>
          <div class="stat-value">¥{{ amountDisplay(cardTwo.value) }}</div>
          <div class="stat-foot">资金处理规模</div>
        </el-card>
      </el-col>

      <el-col :xs="24" :sm="12" :lg="8">
        <el-card shadow="never" class="h-full data-card stat-card">
          <div class="stat-head">
            <span class="stat-label">{{ cardThree.title }}</span>
            <el-icon class="stat-icon stat-icon-amber"><Wallet /></el-icon>
          </div>
          <div class="stat-value">{{ cardThree.prefix }}{{ Number(cardThree.value || 0).toFixed(cardThree.precision) }}</div>
          <div class="stat-foot">累计沉淀数据</div>
        </el-card>
      </el-col>
    </el-row>

    <div class="insight-grid">
      <div class="insight-panel">
        <div class="insight-title">流程健康度</div>
        <div class="insight-value" :class="healthToneClass">{{ healthText }}</div>
        <div class="insight-desc">{{ healthDesc }}</div>
      </div>
      <div class="insight-panel">
        <div class="insight-title">建议动作</div>
        <div class="insight-desc">{{ recommendation }}</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { Document, TrendCharts, Wallet, RefreshRight } from '@element-plus/icons-vue'
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

const lastUpdatedAt = ref<Date | null>(null)

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

const bannerTitle = computed(() => (isEmployee.value ? '个人报销态势总览' : '企业审核与结算总览'))
const bannerSubtitle = computed(() => (isEmployee.value ? '聚焦待确认与处理中单据，缩短提交流程时长。' : '平衡审核效率与打款节奏，保障财务流转稳定。'))

const lastUpdatedText = computed(() => {
  if (!lastUpdatedAt.value) {
    return '刚刚'
  }
  return lastUpdatedAt.value.toLocaleTimeString('zh-CN', { hour12: false })
})

const healthScore = computed(() => {
  if (isEmployee.value) {
    const pendingCount = Number(employeeStats.value.unconfirmed || 0)
    if (pendingCount <= 2) return 90
    if (pendingCount <= 6) return 70
    return 45
  }

  const toAudit = Number(companyStats.value.company_pending_count || 0)
  const toPay = Number(companyStats.value.company_to_pay_count || 0)
  const pressure = toAudit + toPay
  if (pressure <= 10) return 88
  if (pressure <= 30) return 68
  return 42
})

const healthText = computed(() => {
  if (healthScore.value >= 80) return '健康'
  if (healthScore.value >= 60) return '可控'
  return '需关注'
})

const healthToneClass = computed(() => {
  if (healthScore.value >= 80) return 'tone-good'
  if (healthScore.value >= 60) return 'tone-mid'
  return 'tone-risk'
})

const healthDesc = computed(() => {
  if (isEmployee.value) {
    if (healthScore.value >= 80) return '待确认单据数量较低，提交流程顺畅。'
    if (healthScore.value >= 60) return '建议尽快确认单据，避免堆积影响报销时效。'
    return '当前待确认单据较多，建议优先处理高金额单据。'
  }

  if (healthScore.value >= 80) return '审核与打款压力较低，整体节奏稳定。'
  if (healthScore.value >= 60) return '待审核与待打款有积压，建议分批处理。'
  return '流程压力较高，建议优先清理高风险与高金额单据。'
})

const recommendation = computed(() => {
  if (isEmployee.value) {
    return '先处理待确认，再批量发布草稿；金额异常单据优先核验。'
  }
  return '建议上午清理审核池，下午集中打款，减少跨天遗留。'
})

const numberDisplay = (val: number) => Number(val || 0).toLocaleString('zh-CN')
const amountDisplay = (val: number) => Number(val || 0).toLocaleString('zh-CN', { minimumFractionDigits: 2, maximumFractionDigits: 2 })

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
      lastUpdatedAt.value = new Date()
      return
    }

    companyStats.value = {
      company_pending_count: Number(data.company_pending_count || 0),
      monthly_paid_amount: Number(data.monthly_paid_amount || 0),
      company_to_pay_count: Number(data.company_to_pay_count || 0),
    }
    lastUpdatedAt.value = new Date()
  } catch (e: any) {
    ElMessage.error(e?.response?.data?.message || '加载仪表盘统计失败')
  }
}

onMounted(() => {
  loadStats()
})
</script>

<style scoped>
.stat-card {
  padding: 18px;
  transition: transform 0.22s ease, box-shadow 0.22s ease;
  position: relative;
  overflow: hidden;
}

.stat-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 16px 32px rgba(18, 39, 67, 0.12);
}

.stat-card::before {
  content: '';
  position: absolute;
  inset: 0;
  background: radial-gradient(circle at 92% 8%, rgba(31, 111, 235, 0.12), transparent 40%);
  pointer-events: none;
}

.stat-head {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 10px;
}

.stat-label {
  color: #5d728a;
  font-size: 14px;
}

.stat-value {
  font-family: 'Lexend', sans-serif;
  font-size: 34px;
  font-weight: 600;
  color: #0e2948;
  line-height: 1.2;
}

.stat-foot {
  margin-top: 8px;
  color: #7186a0;
  font-size: 12px;
}

.stat-icon {
  width: 30px;
  height: 30px;
  border-radius: 9px;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
}

.stat-icon-blue {
  background: rgba(31, 111, 235, 0.15);
  color: #185ec6;
}

.stat-icon-green {
  background: rgba(22, 163, 74, 0.16);
  color: #167a3a;
}

.stat-icon-amber {
  background: rgba(217, 119, 6, 0.18);
  color: #9a5a08;
}

.dashboard-shell {
  position: relative;
}

.dashboard-banner {
  border: 1px solid #d6e4f3;
  background: linear-gradient(120deg, #eaf4ff 0%, #f4fbff 55%, #ecfbf1 100%);
  border-radius: 16px;
  padding: 16px 18px;
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 12px;
}

.banner-title {
  font-family: 'Lexend', sans-serif;
  font-size: 18px;
  font-weight: 600;
  color: #0f2b4b;
}

.banner-subtitle {
  margin-top: 4px;
  color: #5f7692;
  font-size: 13px;
}

.banner-meta {
  background: rgba(255, 255, 255, 0.72);
  border: 1px solid #dbe8f6;
  border-radius: 12px;
  padding: 10px 12px;
  min-width: 120px;
}

.meta-label {
  font-size: 12px;
  color: #65819f;
}

.meta-value {
  margin-top: 3px;
  font-family: 'Lexend', sans-serif;
  color: #123559;
}

.insight-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 14px;
}

.insight-panel {
  border: 1px solid #d7e4f1;
  border-radius: 14px;
  background: #fbfdff;
  padding: 14px 16px;
}

.insight-title {
  font-size: 13px;
  color: #64809e;
}

.insight-value {
  margin-top: 6px;
  font-family: 'Lexend', sans-serif;
  font-size: 24px;
  font-weight: 600;
}

.insight-desc {
  margin-top: 6px;
  font-size: 13px;
  color: #556d88;
  line-height: 1.5;
}

.tone-good {
  color: #128044;
}

.tone-mid {
  color: #96600a;
}

.tone-risk {
  color: #bb2f2f;
}

@media (max-width: 768px) {
  .dashboard-banner {
    flex-direction: column;
  }

  .insight-grid {
    grid-template-columns: 1fr;
  }
}
</style>
