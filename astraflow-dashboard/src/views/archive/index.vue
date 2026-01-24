<template>
  <div class="h-full flex flex-col">
    <!-- Header / Filter Bar -->
    <div class="mb-6 flex flex-col gap-4">
      <div class="flex items-center justify-between">
          <h1 class="text-2xl font-bold text-slate-900 tracking-tight">Invoice Archive</h1>
           <el-button type="default" @click="handleExport">
            <el-icon class="mr-2"><Download /></el-icon>
            Export Excel
          </el-button>
      </div>
      
      <div class="bg-white p-4 rounded-xl shadow-sm border border-slate-200 flex flex-wrap gap-4 items-center">
         <!-- Search -->
         <el-input v-model="filters.keyword" placeholder="Search by Invoice No. or Vendor" class="!w-64" clearable>
            <template #prefix>
                <el-icon><Search /></el-icon>
            </template>
         </el-input>

         <!-- Status -->
         <el-select v-model="filters.status" placeholder="Status" class="!w-40" clearable>
            <el-option label="All" value="" />
            <el-option label="Paid" value="paid" />
            <el-option label="Approved" value="approved" />
            <el-option label="Pending" value="pending" />
            <el-option label="Rejected" value="rejected" />
         </el-select>

         <!-- Date Range -->
         <el-date-picker
            v-model="filters.dateRange"
            type="daterange"
            range-separator="To"
            start-placeholder="Start date"
            end-placeholder="End date"
            class="!w-80"
         />
         
         <el-button type="primary" class="!bg-indigo-600 !border-indigo-600" @click="fetchData">Search</el-button>
      </div>
    </div>

    <!-- Table Card -->
    <div class="bg-white rounded-xl shadow-sm border border-slate-200 flex-1 overflow-hidden flex flex-col">
      <el-table 
        v-loading="loading"
        :data="tableData" 
        style="width: 100%" 
        class="flex-1"
        header-cell-class-name="!bg-slate-50 !text-slate-600 !font-semibold !border-b-slate-100"
        cell-class-name="!text-slate-600 !border-b-slate-50 hover:!bg-slate-50"
      >
        <el-table-column prop="invoice_no" label="Invoice No." min-width="160" />
        <el-table-column prop="employee_name" label="Employee" min-width="140" />
        <el-table-column label="Amount" min-width="120">
             <template #default="{ row }">
                ${{ row.amount.toFixed(2) }}
             </template>
        </el-table-column>
        <el-table-column prop="date" label="Date" min-width="120" />
        
        <el-table-column label="Status" width="120">
          <template #default="{ row }">
             <el-tag :type="getStatusType(row.status)" round effect="light" class="capitalize border-0">
                 {{ row.status }}
             </el-tag>
          </template>
        </el-table-column>
        
        <el-table-column label="Actions" align="right" width="100">
             <template #default>
                <el-button link type="primary" size="small">View</el-button>
             </template>
        </el-table-column>
      </el-table>
      
      <!-- Pagination -->
      <div class="p-4 border-t border-slate-100 flex justify-end">
          <el-pagination background layout="prev, pager, next" :total="100" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { Search, Download } from '@element-plus/icons-vue'
import { getAllInvoices, exportInvoices } from '../../api/finance'
import type { Invoice } from '../../api/finance'
import { ElMessage } from 'element-plus'

const loading = ref(false)
const tableData = ref<Invoice[]>([])

const filters = reactive({
    keyword: '',
    status: '',
    dateRange: []
})

const getStatusType = (status: string) => {
    switch(status) {
        case 'paid': return 'info' // Gray/Success-ish depending on theme, 'info' is usually gray in Element. Let's use 'success' if we want green or 'info' for history. Prompt said Paid: Gray/Success.
        case 'approved': return 'primary' // Blue
        case 'pending': return 'warning' // Orange
        case 'rejected': return 'danger' // Red
        default: return 'info'
    }
}

// Mock Data
const mockData: Invoice[] = [
    { id: 1, invoice_no: 'INV-2023-999', employee_name: 'John Doe', amount: 120.00, date: '2023-10-01', status: 'paid' },
    { id: 2, invoice_no: 'INV-2023-888', employee_name: 'Jane Smith', amount: 450.00, date: '2023-10-05', status: 'paid' },
    { id: 3, invoice_no: 'INV-2023-777', employee_name: 'Bob Wilson', amount: 89.90, date: '2023-10-10', status: 'rejected' },
    { id: 4, invoice_no: 'INV-2023-666', employee_name: 'Alice Cooper', amount: 1200.00, date: '2023-10-12', status: 'approved' },
]

const fetchData = async () => {
    loading.value = true
    try {
        // const res = await getAllInvoices(filters)
        // tableData.value = res.data
        await new Promise(r => setTimeout(r, 600))
        tableData.value = mockData
    } catch (e) {
        ElMessage.error('Failed to load archive')
    } finally {
        loading.value = false
    }
}

const handleExport = async () => {
    try {
        ElMessage.success('Export started...')
        // await exportInvoices(filters)
    } catch (e) {
        ElMessage.error('Export failed')
    }
}

onMounted(() => {
    fetchData()
})
</script>
