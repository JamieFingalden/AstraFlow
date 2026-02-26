<template>
  <div class="h-full flex flex-col relative pb-28">
    <div class="mb-6 flex items-center justify-between">
      <div>
        <h1 class="text-2xl font-bold text-slate-900 tracking-tight">结算中心</h1>
        <p class="text-slate-500 mt-1 text-sm">展示所有已审核通过的待打款单据，支持批量确认打款。</p>
      </div>
      <div class="flex items-center gap-3">
        <el-button type="primary" class="!bg-indigo-600 !border-indigo-600 hover:!bg-indigo-700" @click="fetchData">
          <el-icon class="mr-2"><Refresh /></el-icon>
          刷新
        </el-button>
      </div>
    </div>

    <div class="bg-white rounded-xl shadow-sm border border-slate-200 flex-1 overflow-hidden flex flex-col">
      <div class="px-6 py-4 border-b border-slate-100">
        <el-input
          v-model="keyword"
          clearable
          placeholder="可按收款人/分类/发票号筛选"
          class="!w-[320px]"
        />
      </div>

      <el-table
        ref="tableRef"
        v-loading="loading"
        :data="filteredData"
        row-key="id"
        style="width: 100%"
        class="flex-1"
        header-cell-class-name="!bg-slate-50 !text-slate-600 !font-semibold !border-b-slate-100"
        cell-class-name="!text-slate-600 !border-b-slate-50 hover:!bg-slate-50"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />

        <el-table-column label="单据编号(ID)" min-width="170">
          <template #default="{ row }">
            <span class="font-mono text-slate-700">{{ row.invoice_number || `#${row.id}` }}</span>
          </template>
        </el-table-column>

        <el-table-column label="收款人" min-width="140">
          <template #default="{ row }">
            <span>{{ row.user?.name || row.user?.username || '-' }}</span>
          </template>
        </el-table-column>

        <el-table-column prop="category" label="分类" min-width="130" />

        <el-table-column label="核定金额" min-width="150" align="right">
          <template #default="{ row }">
            <span class="font-bold text-slate-900">¥{{ Number(row.amount || 0).toFixed(2) }}</span>
          </template>
        </el-table-column>

        <el-table-column label="审核通过时间" min-width="180">
          <template #default="{ row }">
            {{ formatDateTime(row.updated_at) }}
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

    <div class="fixed bottom-0 left-[240px] right-0 z-40 bg-white/95 backdrop-blur border-t border-slate-200 px-8 py-4">
      <div class="max-w-[1400px] mx-auto flex items-center justify-between gap-4">
        <div class="text-slate-700">
          已选 <span class="font-bold text-slate-900">{{ selectedCount }}</span> 项，合计打款金额：
          <span class="font-bold text-emerald-600">¥ {{ selectedTotalAmount.toFixed(2) }}</span>
        </div>

        <el-button
          type="success"
          class="!px-8 !font-semibold !bg-emerald-600 !border-emerald-600 hover:!bg-emerald-700"
          :disabled="selectedCount === 0"
          :loading="submitting"
          @click="handleBatchPay"
        >
          批量打款
        </el-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { Refresh } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import type { ElTable } from 'element-plus'
import { batchPayInvoices, getApprovedInvoices } from '../../../api/settlement'
import type { SettlementInvoiceItem } from '../../../api/settlement'
import { formatDateTime } from '../../../utils/datetime'

const loading = ref(false)
const submitting = ref(false)
const keyword = ref('')

const tableRef = ref<InstanceType<typeof ElTable>>()

const page = ref(1)
const size = ref(10)
const total = ref(0)

const tableData = ref<SettlementInvoiceItem[]>([])
const selectedRows = ref<SettlementInvoiceItem[]>([])

/**
 * 根据关键字对当前页数据进行前端筛选。
 * @returns 过滤后的单据列表。
 */
const filteredData = computed(() => {
  const search = keyword.value.trim().toLowerCase()
  if (!search) {
    return tableData.value
  }

  return tableData.value.filter((item) => {
    const payee = item.user?.name || item.user?.username || ''
    const invoiceNo = item.invoice_number || `${item.id}`
    const category = item.category || ''
    return `${payee} ${invoiceNo} ${category}`.toLowerCase().includes(search)
  })
})

/**
 * 计算当前已勾选单据数量。
 * @returns 勾选数量。
 */
const selectedCount = computed(() => selectedRows.value.length)

/**
 * 计算当前已勾选单据总金额。
 * @returns 打款总金额。
 */
const selectedTotalAmount = computed(() => {
  return selectedRows.value.reduce((sum, item) => sum + Number(item.amount || 0), 0)
})

/**
 * 拉取待打款列表。
 * 入参：无（使用当前分页状态）。
 * 出参：无（结果写入响应式状态）。
 */
const fetchData = async () => {
  loading.value = true
  try {
    const data = await getApprovedInvoices({ page: page.value, size: size.value })
    tableData.value = data.items || []
    total.value = Number(data.total || 0)
    selectedRows.value = []
    tableRef.value?.clearSelection()
  } catch (e: any) {
    ElMessage.error(e?.response?.data?.message || '加载待打款单据失败')
  } finally {
    loading.value = false
  }
}

/**
 * 处理表格勾选变化。
 * @param rows 当前所有勾选行。
 */
const handleSelectionChange = (rows: SettlementInvoiceItem[]) => {
  selectedRows.value = rows
}

/**
 * 处理页码变化并重新加载数据。
 * @param val 新页码。
 */
const handlePageChange = (val: number) => {
  page.value = val
  fetchData()
}

/**
 * 处理每页条数变化并重新加载数据。
 * @param val 新每页条数。
 */
const handleSizeChange = (val: number) => {
  size.value = val
  page.value = 1
  fetchData()
}

/**
 * 执行批量打款动作。
 * 业务流程：二次确认 -> 调用打款接口 -> 提示成功 -> 清空勾选并刷新列表。
 */
const handleBatchPay = async () => {
  if (selectedCount.value === 0) {
    return
  }

  const confirmText = `确认对选中的 ${selectedCount.value} 笔单据进行打款，共计 ¥ ${selectedTotalAmount.value.toFixed(2)} 吗？此操作不可逆。`

  try {
    await ElMessageBox.confirm(confirmText, '确认打款', {
      confirmButtonText: '确认打款',
      cancelButtonText: '取消',
      type: 'warning',
    })

    submitting.value = true
    const invoiceIds = selectedRows.value.map((item) => item.id)
    await batchPayInvoices({ invoice_ids: invoiceIds })
    ElMessage.success('批量打款成功')
    await fetchData()
  } catch (e: any) {
    if (e !== 'cancel' && e !== 'close') {
      ElMessage.error(e?.response?.data?.message || '批量打款失败')
    }
  } finally {
    submitting.value = false
  }
}

onMounted(() => {
  fetchData()
})
</script>
