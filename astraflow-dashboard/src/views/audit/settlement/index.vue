<template>
  <div class="h-full flex flex-col relative pb-20"> <!-- Added padding-bottom for floating bar -->
    <!-- Header -->
    <div class="mb-6 flex items-center justify-between">
      <div>
        <h1 class="text-2xl font-bold text-slate-900 tracking-tight">Finance Settlement</h1>
        <p class="text-slate-500 mt-1 text-sm">Review and process payments for approved invoices.</p>
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
        cell-class-name="!text-slate-600 !border-b-slate-50 hover:!bg-slate-50"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        
        <el-table-column label="Invoice No." min-width="180">
            <template #default="{ row }">
                <span class="font-mono text-slate-700">{{ row.invoice_no }}</span>
            </template>
        </el-table-column>

        <el-table-column prop="employee_name" label="Employee" min-width="150" />

        <el-table-column label="Amount" min-width="150">
            <template #default="{ row }">
                <span class="font-bold text-slate-900">${{ row.amount.toFixed(2) }}</span>
            </template>
        </el-table-column>

        <el-table-column prop="approved_date" label="Approved Date" min-width="150" />
      </el-table>
    </div>

    <!-- Floating Action Bar -->
    <transition name="slide-up">
        <div v-if="selectedRows.length > 0" class="fixed bottom-8 left-1/2 transform -translate-x-1/2 bg-slate-900 text-white rounded-full shadow-2xl shadow-slate-900/40 px-6 py-3 flex items-center gap-8 z-50 min-w-[400px]">
            <div class="flex flex-col">
                <span class="text-xs text-slate-400 font-medium uppercase tracking-wider">Selected</span>
                <span class="font-bold text-lg">{{ selectedRows.length }} items</span>
            </div>
            
            <div class="h-8 w-px bg-slate-700"></div>
            
            <div class="flex flex-col flex-1">
                <span class="text-xs text-slate-400 font-medium uppercase tracking-wider">Total Payable</span>
                <span class="font-bold text-lg text-emerald-400">${{ totalAmount.toFixed(2) }}</span>
            </div>

            <el-button type="success" :loading="submitting" class="!rounded-full !px-6 !font-bold !bg-emerald-500 !border-emerald-500 hover:!bg-emerald-600 !text-white" @click="handleBatchPay">
                Mark as Paid
            </el-button>
        </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { Refresh } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getApprovedInvoices, markBatchPaid } from '../../../api/finance'
import type { Invoice } from '../../../api/finance'

const loading = ref(false)
const submitting = ref(false)
const tableData = ref<Invoice[]>([])
const selectedRows = ref<Invoice[]>([])

// Computed Total
const totalAmount = computed(() => {
    return selectedRows.value.reduce((sum, item) => sum + item.amount, 0)
})

// Mock Data
const mockData: Invoice[] = [
    { id: 101, invoice_no: 'INV-2023-001', employee_name: 'John Doe', amount: 150.00, approved_date: '2023-10-25', status: 'approved' },
    { id: 102, invoice_no: 'INV-2023-002', employee_name: 'Jane Smith', amount: 42.50, approved_date: '2023-10-25', status: 'approved' },
    { id: 103, invoice_no: 'INV-2023-003', employee_name: 'Mike Johnson', amount: 890.00, approved_date: '2023-10-24', status: 'approved' },
    { id: 104, invoice_no: 'INV-2023-004', employee_name: 'Sarah Connor', amount: 1200.00, approved_date: '2023-10-23', status: 'approved' },
]

const fetchData = async () => {
    loading.value = true
    try {
        // const res = await getApprovedInvoices()
        // tableData.value = res.data
        await new Promise(r => setTimeout(r, 600))
        tableData.value = mockData
    } catch (e) {
        ElMessage.error('Failed to load data')
    } finally {
        loading.value = false
    }
}

const handleSelectionChange = (val: Invoice[]) => {
    selectedRows.value = val
}

const handleBatchPay = async () => {
    try {
        await ElMessageBox.confirm(
            `Mark ${selectedRows.value.length} invoices as PAID? Total: $${totalAmount.value.toFixed(2)}`,
            'Confirm Payment',
            { confirmButtonText: 'Confirm', cancelButtonText: 'Cancel', type: 'warning' }
        )
        
        submitting.value = true
        const ids = selectedRows.value.map(i => i.id)
        
        // await markBatchPaid(ids)
        await new Promise(r => setTimeout(r, 1000))
        
        ElMessage.success('Payment processed successfully')
        
        // Remove processed items from list
        tableData.value = tableData.value.filter(item => !ids.includes(item.id))
        selectedRows.value = []
        
    } catch (e) {
        // Cancelled or Error
    } finally {
        submitting.value = false
    }
}

onMounted(() => {
    fetchData()
})
</script>

<style scoped>
.slide-up-enter-active,
.slide-up-leave-active {
  transition: all 0.3s cubic-bezier(0.19, 1, 0.22, 1);
}

.slide-up-enter-from,
.slide-up-leave-to {
  transform: translate(-50%, 100%);
  opacity: 0;
}
</style>
