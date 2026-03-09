<template>
  <div class="h-full flex flex-col gap-4">
    <div>
      <h1 class="af-title text-2xl font-semibold text-slate-900 tracking-tight">历史归档查询</h1>
      <p class="page-subtitle">按关键字、状态和时间范围检索历史单据。</p>
    </div>

    <el-card shadow="never" class="page-shell !rounded-2xl">
      <el-form inline :model="filters" @submit.prevent>
        <el-form-item label="关键字">
          <el-input
            v-model="filters.keyword"
            clearable
            placeholder="搜索事由/分类"
            style="width: 220px"
            @keyup.enter="handleSearch"
          />
        </el-form-item>

        <el-form-item label="状态">
          <el-select v-model="filters.status" clearable placeholder="请选择状态" style="width: 180px">
            <el-option label="待审核" value="pending" />
            <el-option label="待打款" value="approved" />
            <el-option label="已打款" value="paid" />
            <el-option label="已驳回" value="rejected" />
          </el-select>
        </el-form-item>

        <el-form-item label="时间范围">
          <el-date-picker
            v-model="filters.dateRange"
            type="daterange"
            value-format="YYYY-MM-DD"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            style="width: 300px"
          />
        </el-form-item>

        <el-form-item>
          <el-button type="primary" class="af-button-primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card shadow="never" class="page-shell !rounded-2xl flex-1 flex flex-col">
      <el-table v-loading="loading" :data="tableData" style="width: 100%" class="flex-1">
        <el-table-column prop="invoice_number" label="单据编号" min-width="170" />

        <el-table-column label="提交人" min-width="140">
          <template #default="{ row }">
            {{ row.user?.name || row.user?.username || '-' }}
          </template>
        </el-table-column>

        <el-table-column prop="category" label="分类" min-width="120" />

        <el-table-column label="金额" min-width="120" align="right">
          <template #default="{ row }">
            ¥{{ Number(row.amount || 0).toFixed(2) }}
          </template>
        </el-table-column>

        <el-table-column label="状态" min-width="120">
          <template #default="{ row }">
            <el-tag :type="statusTagType(row.status)" round>
              {{ statusText(row.status) }}
            </el-tag>
          </template>
        </el-table-column>

        <el-table-column label="提交时间" min-width="180">
          <template #default="{ row }">
            {{ formatDateTime(row.created_at) }}
          </template>
        </el-table-column>
      </el-table>

      <div class="pt-4 flex justify-end">
        <el-pagination
          background
          layout="total, prev, pager, next, sizes"
          :current-page="page"
          :page-size="size"
          :page-sizes="[10, 20, 50, 100]"
          :total="total"
          @current-change="handlePageChange"
          @size-change="handleSizeChange"
        />
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { getArchiveInvoices } from '../../api/dashboard'
import type { ArchiveInvoiceItem } from '../../api/dashboard'
import { formatDateTime } from '../../utils/datetime'

const loading = ref(false)
const tableData = ref<ArchiveInvoiceItem[]>([])
const page = ref(1)
const size = ref(10)
const total = ref(0)

const filters = reactive({
  keyword: '',
  status: '',
  dateRange: [] as string[],
})

/**
 * 获取归档列表。
 * 入参：无（使用当前筛选与分页状态组装请求参数）。
 * 出参：无（结果写入 tableData/total）。
 */
const fetchData = async () => {
  loading.value = true
  try {
    const [startDate, endDate] = filters.dateRange || []
    const data = await getArchiveInvoices({
      page: page.value,
      size: size.value,
      status: filters.status || undefined,
      keyword: filters.keyword || undefined,
      start_date: startDate || undefined,
      end_date: endDate || undefined,
    })

    tableData.value = data.items || []
    total.value = Number(data.total || 0)
  } catch (e: any) {
    ElMessage.error(e?.response?.data?.message || '加载归档列表失败')
  } finally {
    loading.value = false
  }
}

/**
 * 处理搜索动作。
 * 入参：无。
 * 出参：无（重置到第一页并重新查询）。
 */
const handleSearch = () => {
  page.value = 1
  fetchData()
}

/**
 * 处理重置动作。
 * 入参：无。
 * 出参：无（清空所有筛选并重新查询）。
 */
const handleReset = () => {
  filters.keyword = ''
  filters.status = ''
  filters.dateRange = []
  page.value = 1
  fetchData()
}

/**
 * 处理分页页码变化。
 * @param val 新页码。
 */
const handlePageChange = (val: number) => {
  page.value = val
  fetchData()
}

/**
 * 处理分页每页条数变化。
 * @param val 新每页条数。
 */
const handleSizeChange = (val: number) => {
  size.value = val
  page.value = 1
  fetchData()
}

/**
 * 将状态值映射为中文文案。
 * @param status 单据状态。
 * @returns 中文状态文案。
 */
const statusText = (status: string) => {
  switch (status) {
    case 'pending':
      return '待审核'
    case 'approved':
      return '待打款'
    case 'paid':
      return '已打款'
    case 'rejected':
      return '已驳回'
    case 'unconfirmed':
      return '待确认'
    case 'draft':
      return '草稿'
    case 'recognizing':
      return '识别中'
    default:
      return status
  }
}

/**
 * 将状态值映射为标签颜色。
 * @param status 单据状态。
 * @returns Element Plus 标签类型。
 */
const statusTagType = (status: string): 'success' | 'warning' | 'info' | 'danger' | 'primary' => {
  switch (status) {
    case 'pending':
      return 'warning'
    case 'approved':
      return 'primary'
    case 'paid':
      return 'success'
    case 'rejected':
      return 'danger'
    default:
      return 'info'
  }
}

onMounted(() => {
  fetchData()
})
</script>
