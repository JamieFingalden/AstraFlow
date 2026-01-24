<template>
  <div class="h-full w-full flex overflow-hidden bg-white relative">
    
    <!-- LEFT PANEL: Image Viewer -->
    <div class="w-1/2 bg-slate-900 relative flex items-center justify-center overflow-hidden">
        <!-- Toolbar -->
        <div class="absolute bottom-6 left-1/2 transform -translate-x-1/2 z-20 bg-slate-800/80 backdrop-blur rounded-full px-4 py-2 flex items-center gap-4 border border-slate-700 shadow-xl">
            <el-button circle size="small" class="!bg-transparent !border-slate-500 !text-white hover:!bg-slate-700" @click="zoomOut">
                <el-icon><Minus /></el-icon>
            </el-button>
            <span class="text-xs font-mono text-slate-300 w-12 text-center">{{ Math.round(scale * 100) }}%</span>
            <el-button circle size="small" class="!bg-transparent !border-slate-500 !text-white hover:!bg-slate-700" @click="zoomIn">
                <el-icon><Plus /></el-icon>
            </el-button>
             <el-divider direction="vertical" class="!border-slate-600" />
            <el-button circle size="small" class="!bg-transparent !border-slate-500 !text-white hover:!bg-slate-700" @click="rotate">
                <el-icon><RefreshRight /></el-icon>
            </el-button>
        </div>

        <!-- Image Container -->
        <div 
            class="transition-transform duration-200 ease-out cursor-grab active:cursor-grabbing"
            :style="{ transform: `scale(${scale}) rotate(${rotation}deg)` }"
            @mousedown="startDrag"
        >
             <img 
                :src="invoiceData?.file_url || 'https://via.placeholder.com/600x800?text=Invoice+Preview'" 
                class="max-w-[80%] max-h-[90vh] shadow-2xl rounded-sm object-contain" 
                alt="Invoice"
                draggable="false"
             />
        </div>
        
        <div class="absolute top-4 left-4 z-20">
             <el-button circle class="!bg-slate-800/50 !border-0 !text-white" @click="$router.back()">
                <el-icon><Back /></el-icon>
             </el-button>
        </div>
    </div>

    <!-- RIGHT PANEL: Audit Form -->
    <div class="w-1/2 flex flex-col h-full bg-slate-50 border-l border-slate-200">
        <!-- Header -->
        <div class="px-8 py-6 bg-white border-b border-slate-200 flex justify-between items-start">
            <div>
                <div class="flex items-center gap-3 mb-1">
                    <h2 class="text-xl font-bold text-slate-900">Audit Invoice #{{ invoiceData?.id }}</h2>
                    <el-tag :type="invoiceData?.confidence >= 0.8 ? 'success' : 'warning'" effect="dark" size="small" round>
                         AI: {{ (invoiceData?.confidence * 100)?.toFixed(0) }}%
                    </el-tag>
                </div>
                <p class="text-sm text-slate-500">Review the extracted data against the original document.</p>
            </div>
            <div class="text-right">
                <div class="text-xs text-slate-400 uppercase tracking-wide font-semibold">Total Amount</div>
                <div class="text-2xl font-bold text-indigo-600">${{ invoiceData?.amount?.toFixed(2) }}</div>
            </div>
        </div>

        <!-- Scrollable Form Content -->
        <div class="flex-1 overflow-y-auto p-8">
            <el-form 
                ref="formRef"
                :model="formData" 
                :rules="rules"
                label-position="top" 
                class="max-w-xl mx-auto space-y-4"
            >
                <div class="grid grid-cols-2 gap-6">
                    <el-form-item label="Vendor Name" prop="vendor">
                        <el-input v-model="formData.vendor" />
                    </el-form-item>
                    <el-form-item label="Invoice Date" prop="date">
                        <el-date-picker v-model="formData.date" type="date" placeholder="Pick a date" class="!w-full" />
                    </el-form-item>
                </div>

                <div class="grid grid-cols-2 gap-6">
                    <el-form-item label="Total Amount" prop="amount">
                         <el-input-number v-model="formData.amount" :precision="2" :step="0.1" class="!w-full" />
                    </el-form-item>
                    <el-form-item label="Category" prop="category">
                        <el-select v-model="formData.category" placeholder="Select" class="!w-full">
                            <el-option label="Travel" value="Travel" />
                            <el-option label="Meals" value="Meals" />
                            <el-option label="Office Supplies" value="Office Supplies" />
                            <el-option label="Software" value="Software" />
                        </el-select>
                    </el-form-item>
                </div>

                <el-form-item label="Description">
                    <el-input v-model="formData.description" type="textarea" :rows="3" />
                </el-form-item>
                
                <el-divider content-position="left">Line Items ({{ formData.items?.length || 0 }})</el-divider>
                
                <div v-for="(item, index) in formData.items" :key="index" class="bg-white p-4 rounded-lg border border-slate-200 mb-3 text-sm">
                    <div class="flex justify-between font-medium text-slate-900 mb-2">
                        <span>{{ item.name }}</span>
                        <span>${{ item.amount }}</span>
                    </div>
                    <div class="flex gap-4 text-slate-500 text-xs">
                        <span>Qty: {{ item.count }}</span>
                        <span>Price: {{ item.unitPrice }}</span>
                    </div>
                </div>
                
            </el-form>
        </div>

        <!-- Sticky Footer -->
        <div class="bg-white border-t border-slate-200 p-6 flex items-center justify-between gap-4 z-10 shadow-[0_-4px_6px_-1px_rgba(0,0,0,0.05)]">
            <el-button 
                type="danger" 
                plain 
                class="flex-1 !h-12 !text-base !font-medium"
                @click="handleReject"
            >
                Reject / Flag
            </el-button>
            <el-button 
                type="success" 
                class="flex-1 !h-12 !text-base !font-medium !bg-emerald-600 !border-emerald-600 hover:!bg-emerald-700"
                :loading="submitting"
                @click="handleApprove"
            >
                Approve & Process
            </el-button>
        </div>

        <!-- Reject Dialog -->
         <el-dialog v-model="rejectDialogVisible" title="Reject Invoice" width="400px">
            <el-input v-model="rejectReason" type="textarea" placeholder="Please enter reason for rejection..." :rows="4" />
            <template #footer>
                <el-button @click="rejectDialogVisible = false">Cancel</el-button>
                <el-button type="danger" @click="confirmReject">Confirm Reject</el-button>
            </template>
         </el-dialog>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Plus, Minus, RefreshRight, Back } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import { getInvoiceDetail, processAudit } from '../../../api/audit'
import type { Invoice, AuditResult } from '../../../api/audit'

const route = useRoute()
const router = useRouter()

// -- State --
const loading = ref(false)
const submitting = ref(false)
const invoiceData = ref<Invoice | null>(null)
const formData = reactive<AuditResult>({
    status: 'approved',
    vendor: '',
    date: '',
    amount: 0,
    category: '',
    items: [],
    description: ''
})

// -- Image Controls --
const scale = ref(1)
const rotation = ref(0)
const zoomIn = () => scale.value = Math.min(scale.value + 0.2, 3)
const zoomOut = () => scale.value = Math.max(scale.value - 0.2, 0.5)
const rotate = () => rotation.value += 90

// Drag Logic (Simple Placeholder)
const startDrag = (e: MouseEvent) => {
    // Implement pan logic here if needed, for now just placeholder
    e.preventDefault()
}

// -- Fetch Data --
const fetchData = async () => {
    const id = route.params.id
    if (!id) return
    
    loading.value = true
    try {
        // const res = await getInvoiceDetail(id as string)
        // invoiceData.value = res.data
        
        // Mock
        await new Promise(r => setTimeout(r, 500))
        invoiceData.value = {
            id: id,
            trace_id: 'TRC-' + id,
            user: { name: 'John Doe' },
            vendor: 'Amazon Web Services',
            amount: 124.50,
            date: '2023-10-24',
            confidence: 0.92,
            status: 'pending',
            file_url: 'https://images.unsplash.com/photo-1634733988138-bf2c3a2a13fa?q=80&w=1000&auto=format&fit=crop', // Placeholder Invoice Image
            category: 'Software',
            items: [
                { name: 'Hosting Fee', count: 1, amount: 100.00, unitPrice: 100.00 },
                { name: 'Tax', count: 1, amount: 24.50, unitPrice: 24.50 }
            ]
        } as any

        // Fill Form
        formData.vendor = invoiceData.value?.vendor || ''
        formData.date = invoiceData.value?.date || ''
        formData.amount = invoiceData.value?.amount || 0
        formData.items = invoiceData.value?.items || []
        formData.category = invoiceData.value?.category || ''
        
    } catch (e) {
        ElMessage.error('Failed to load invoice details')
    } finally {
        loading.value = false
    }
}

// -- Actions --
const formRef = ref()
const rules = {
    vendor: [{ required: true, message: 'Required', trigger: 'blur' }],
    amount: [{ required: true, message: 'Required', trigger: 'blur' }],
    category: [{ required: true, message: 'Required', trigger: 'change' }]
}

const handleApprove = async () => {
    if (!formRef.value) return
    await formRef.value.validate(async (valid: boolean) => {
        if (valid) {
            submitting.value = true
            try {
                // await processAudit(route.params.id as string, { ...formData, status: 'approved' })
                await new Promise(r => setTimeout(r, 800))
                ElMessage.success('Invoice Approved')
                router.push('/audit')
            } catch (e) {
                ElMessage.error('Failed')
            } finally {
                submitting.value = false
            }
        }
    })
}

const rejectDialogVisible = ref(false)
const rejectReason = ref('')

const handleReject = () => {
    rejectReason.value = ''
    rejectDialogVisible.value = true
}

const confirmReject = async () => {
    if (!rejectReason.value) return ElMessage.warning('Please provide a reason')
    
    try {
        // await processAudit(route.params.id as string, { status: 'rejected', remarks: rejectReason.value })
         await new Promise(r => setTimeout(r, 800))
        ElMessage.info('Invoice Rejected')
        router.push('/audit')
    } catch (e) {
        ElMessage.error('Failed')
    }
}

onMounted(() => {
    fetchData()
})
</script>
