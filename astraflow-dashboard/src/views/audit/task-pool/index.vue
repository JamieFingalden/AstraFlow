<template>
  <div class="h-full flex flex-col gap-4">
    <div class="page-header">
      <div>
        <h1 class="af-title text-2xl font-semibold text-slate-900 tracking-tight">审核任务池</h1>
        <p class="page-subtitle">查看并处理待审核（pending）的单据。</p>
      </div>
      <el-button type="primary" class="af-button-primary" @click="fetchData">
        <el-icon class="mr-2"><Refresh /></el-icon>
        刷新
      </el-button>
    </div>

    <div class="page-shell flex-1 overflow-hidden flex flex-col">
      <el-table
        v-loading="loading"
        :data="tableData"
        style="width: 100%"
        class="flex-1"
        header-cell-class-name="!bg-slate-50 !text-slate-600 !font-semibold !border-b-slate-100"
        cell-class-name="!text-slate-600 !border-b-slate-50 hover:!bg-slate-50 cursor-pointer"
        @row-click="handleRowClick"
      >
        <el-table-column label="单据ID" width="120">
          <template #default="{ row }">
            <span class="font-mono text-xs text-slate-500 bg-slate-100 px-2 py-1 rounded">#{{ row.id }}</span>
          </template>
        </el-table-column>

        <el-table-column label="提交人" min-width="160">
          <template #default="{ row }">
            <span class="text-slate-900 font-medium text-sm">{{ row.user_name || '-' }}</span>
          </template>
        </el-table-column>

        <el-table-column prop="invoice_number" label="发票号" min-width="170" />

        <el-table-column prop="vendor" label="商家" min-width="180" />

        <el-table-column label="金额" width="140">
          <template #default="{ row }">
            <span class="font-bold text-slate-900">¥{{ Number(row.amount || 0).toFixed(2) }}</span>
          </template>
        </el-table-column>

        <el-table-column label="发票日期" width="140">
          <template #default="{ row }">
            {{ formatDate(row.invoice_date) }}
          </template>
        </el-table-column>

        <el-table-column label="提交时间" width="180">
          <template #default="{ row }">
            {{ formatDateTime(row.created_at) }}
          </template>
        </el-table-column>

        <el-table-column label="操作" align="right" width="100">
          <template #default="{ row }">
            <el-button
                link
                type="primary"
                class="!font-medium"
                @click.stop="handleRowClick(row as PendingInvoiceItem)"
            >
                审核
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="px-6 py-4 border-t border-slate-100 flex items-center justify-end">
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
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Refresh } from '@element-plus/icons-vue'
import { getPendingInvoices } from '../../../api/audit'
import type { PendingInvoiceItem } from '../../../api/audit'
import { ElMessage } from 'element-plus'
import { formatDate, formatDateTime } from '../../../utils/datetime'

const router = useRouter()
const loading = ref(false)
const tableData = ref<PendingInvoiceItem[]>([])
const page = ref(1)
const size = ref(10)
const total = ref(0)

const fetchData = async () => {
  loading.value = true
  try {
    const data = await getPendingInvoices({ page: page.value, size: size.value })
    tableData.value = data.items || []
    total.value = Number(data.total || 0)
  } catch (e: any) {
    ElMessage.error(e?.response?.data?.message || '加载审核任务失败')
  } finally {
    loading.value = false
  }
}

const handleRowClick = (row: PendingInvoiceItem) => {
  router.push(`/audit/tasks/${row.id}`)
}

const handlePageChange = (val: number) => {
  page.value = val
  fetchData()
}

const handleSizeChange = (val: number) => {
  size.value = val
  page.value = 1
  fetchData()
}

onMounted(() => {
  fetchData()
})
</script>
