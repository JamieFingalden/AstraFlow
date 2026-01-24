<template>
  <div class="h-full flex flex-col">
    <!-- Header -->
    <div class="mb-6 flex items-center justify-between">
      <div>
        <h1 class="text-2xl font-bold text-slate-900 tracking-tight">Audit Task Pool</h1>
        <p class="text-slate-500 mt-1 text-sm">Review pending invoices processed by AI.</p>
      </div>
      <el-button type="primary" class="!bg-indigo-600 !border-indigo-600 hover:!bg-indigo-700" @click="fetchData">
        <el-icon class="mr-2"><Refresh /></el-icon>
        Refresh
      </el-button>
    </div>

    <!-- Table Card -->
    <div class="bg-white rounded-xl shadow-sm border border-slate-200 flex-1 overflow-hidden flex flex-col">
      <el-table 
        v-loading="loading"
        :data="tableData" 
        style="width: 100%" 
        class="flex-1"
        header-cell-class-name="!bg-slate-50 !text-slate-600 !font-semibold !border-b-slate-100"
        cell-class-name="!text-slate-600 !border-b-slate-50 hover:!bg-slate-50 cursor-pointer"
        @row-click="handleRowClick"
      >
        <!-- Trace ID -->
        <el-table-column label="Trace ID" width="180">
            <template #default="{ row }">
                <span class="font-mono text-xs text-slate-500 bg-slate-100 px-2 py-1 rounded">{{ row.trace_id }}</span>
            </template>
        </el-table-column>

        <!-- User -->
        <el-table-column label="Submitted By" width="200">
          <template #default="{ row }">
            <div class="flex items-center gap-3">
              <el-avatar :size="32" :src="row.user?.avatar" class="border border-slate-200">
                 {{ row.user?.name?.charAt(0) }}
              </el-avatar>
              <span class="text-slate-900 font-medium text-sm">{{ row.user?.name }}</span>
            </div>
          </template>
        </el-table-column>

        <!-- Vendor -->
        <el-table-column prop="vendor" label="Vendor" min-width="180" />

        <!-- Amount -->
        <el-table-column label="Amount" width="140">
            <template #default="{ row }">
                <span class="font-bold text-slate-900">${{ row.amount?.toFixed(2) }}</span>
            </template>
        </el-table-column>

        <!-- Date -->
        <el-table-column prop="date" label="Date" width="140" />

        <!-- AI Confidence -->
        <el-table-column label="AI Confidence" width="140">
          <template #default="{ row }">
            <el-tag 
              :type="row.confidence >= 0.8 ? 'success' : 'warning'" 
              effect="light" 
              round
              class="!border-0 font-medium"
            >
              {{ (row.confidence * 100).toFixed(0) }}% {{ row.confidence >= 0.8 ? 'High' : 'Low' }}
            </el-tag>
          </template>
        </el-table-column>

        <!-- Actions -->
        <el-table-column label="Action" align="right" width="100">
          <template #default="{ row }">
            <el-button 
                link 
                type="primary" 
                class="!font-medium"
                @click.stop="handleRowClick(row)"
            >
                Review
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Refresh } from '@element-plus/icons-vue'
import { getPendingInvoices } from '../../../api/audit'
import type { Invoice } from '../../../api/audit'
import { ElMessage } from 'element-plus'

const router = useRouter()
const loading = ref(false)
const tableData = ref<Invoice[]>([])

// Mock Data
const mockData: Invoice[] = [
    { 
        id: 1, 
        trace_id: 'TRC-928374', 
        user: { name: 'John Doe', avatar: '' }, 
        vendor: 'Amazon Web Services', 
        amount: 124.50, 
        date: '2023-10-24', 
        confidence: 0.95, 
        status: 'pending',
        file_url: '' 
    },
    { 
        id: 2, 
        trace_id: 'TRC-192834', 
        user: { name: 'Sarah Smith', avatar: '' }, 
        vendor: 'Uber Trip', 
        amount: 34.20, 
        date: '2023-10-23', 
        confidence: 0.72, 
        status: 'pending',
        file_url: '' 
    }
]

const fetchData = async () => {
    loading.value = true
    try {
        // const res = await getPendingInvoices()
        // tableData.value = res.data
        
        // Mock
        await new Promise(r => setTimeout(r, 600))
        tableData.value = mockData
    } catch (e) {
        ElMessage.error('Failed to load tasks')
    } finally {
        loading.value = false
    }
}

const handleRowClick = (row: Invoice) => {
    router.push(`/audit/tasks/${row.id}`)
}

onMounted(() => {
    fetchData()
})
</script>
