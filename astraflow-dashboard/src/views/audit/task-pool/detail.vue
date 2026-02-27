<template>
  <div class="h-full w-full flex overflow-hidden bg-white relative" v-loading="loading">
    <div class="w-1/2 bg-slate-900 relative flex items-center justify-center overflow-hidden">
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

      <div class="transition-transform duration-200 ease-out" :style="{ transform: `scale(${scale}) rotate(${rotation}deg)` }">
        <img
          :src="displayImageUrl"
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

    <div class="w-1/2 flex flex-col h-full bg-slate-50 border-l border-slate-200">
      <div class="px-8 py-6 bg-white border-b border-slate-200 flex justify-between items-start">
        <div>
          <div class="flex items-center gap-3 mb-1">
            <h2 class="text-xl font-bold text-slate-900">审核单据 #{{ invoiceData?.id }}</h2>
            <el-tag type="warning" effect="light" round>{{ invoiceData?.status }}</el-tag>
          </div>
          <p class="text-sm text-slate-500">核对单据信息并执行通过或驳回。</p>
        </div>
        <div class="text-right">
          <div class="text-xs text-slate-400 uppercase tracking-wide font-semibold">当前金额</div>
          <div class="text-2xl font-bold text-indigo-600">¥{{ Number(invoiceData?.amount || 0).toFixed(2) }}</div>
        </div>
      </div>

      <div class="flex-1 overflow-y-auto p-8">
        <el-descriptions :column="2" border class="bg-white rounded-lg">
          <el-descriptions-item label="提交人">{{ invoiceData?.user_name || '-' }}</el-descriptions-item>
          <el-descriptions-item label="用户名">{{ invoiceData?.user_username || '-' }}</el-descriptions-item>
          <el-descriptions-item label="发票号">{{ invoiceData?.invoice_number || '-' }}</el-descriptions-item>
          <el-descriptions-item label="发票日期">{{ formatDate(invoiceData?.invoice_date) }}</el-descriptions-item>
          <el-descriptions-item label="商家">{{ invoiceData?.vendor || '-' }}</el-descriptions-item>
          <el-descriptions-item label="分类">{{ invoiceData?.category || '-' }}</el-descriptions-item>
          <el-descriptions-item label="描述" :span="2">{{ invoiceData?.description || '-' }}</el-descriptions-item>
        </el-descriptions>

        <el-form ref="formRef" :model="formData" :rules="rules" label-position="top" class="max-w-xl mt-6">
          <el-form-item label="审批通过金额" prop="final_amount">
            <el-input-number v-model="formData.final_amount" :precision="2" :step="1" :min="0" class="!w-full" />
          </el-form-item>

          <el-form-item label="审核备注（通过可选，驳回必填）" prop="remarks">
            <el-input v-model="formData.remarks" type="textarea" :rows="4" placeholder="请输入审核说明" />
          </el-form-item>
        </el-form>
      </div>

      <div class="bg-white border-t border-slate-200 p-6 flex items-center justify-between gap-4 z-10 shadow-[0_-4px_6px_-1px_rgba(0,0,0,0.05)]">
        <el-button type="danger" plain class="flex-1 !h-12 !text-base !font-medium" :loading="submitting" @click="handleReject">
          驳回
        </el-button>
        <el-button
          type="success"
          class="flex-1 !h-12 !text-base !font-medium !bg-emerald-600 !border-emerald-600 hover:!bg-emerald-700"
          :loading="submitting"
          @click="handleApprove"
        >
          通过
        </el-button>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Plus, Minus, RefreshRight, Back } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getInvoiceDetail, processAudit } from '../../../api/audit'
import type { FormInstance, FormRules } from 'element-plus'
import type { InvoiceDetail, ReviewInvoicePayload } from '../../../api/audit'
import { formatDate } from '../../../utils/datetime'

const route = useRoute()
const router = useRouter()

const loading = ref(false)
const submitting = ref(false)
const invoiceData = ref<InvoiceDetail | null>(null)
const formData = reactive<ReviewInvoicePayload>({
  action: 'approve',
  final_amount: undefined,
  remarks: '',
})

/**
 * 计算图片可访问地址：
 * - 后端返回绝对地址时直接使用
 * - 后端返回 /uploads/... 相对地址时，自动补齐后端域名
 */
const displayImageUrl = computed(() => {
  const rawUrl = invoiceData.value?.image_url || ''
  if (!rawUrl) {
    return 'https://via.placeholder.com/600x800?text=Invoice+Preview'
  }

  if (/^https?:\/\//i.test(rawUrl)) {
    return rawUrl
  }

  const apiBase = import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1'
  const backendOrigin = apiBase.replace(/\/api\/v1\/?$/, '').replace(/\/$/, '')

  if (rawUrl.startsWith('/')) {
    return `${backendOrigin}${rawUrl}`
  }

  if (rawUrl.startsWith('./')) {
    return `${backendOrigin}/${rawUrl.slice(2)}`
  }

  return `${backendOrigin}/${rawUrl}`
})

const scale = ref(1)
const rotation = ref(0)
const zoomIn = () => scale.value = Math.min(scale.value + 0.2, 3)
const zoomOut = () => scale.value = Math.max(scale.value - 0.2, 0.5)
const rotate = () => rotation.value += 90

const fetchData = async () => {
  const id = route.params.id
  if (!id) return

  loading.value = true
  try {
    const data = await getInvoiceDetail(id as string)
    invoiceData.value = data
    formData.final_amount = Number(data.amount || 0)
    formData.remarks = data.review_remarks || ''
  } catch (e: any) {
    ElMessage.error(e?.response?.data?.message || '加载单据详情失败')
  } finally {
    loading.value = false
  }
}

const formRef = ref<FormInstance>()
const rules: FormRules = {
  final_amount: [
    {
      validator: (_rule, value, callback) => {
        if (value === undefined || value === null || Number(value) < 0) {
          callback(new Error('请输入合法金额'))
          return
        }
        callback()
      },
      trigger: 'blur',
    }
  ],
}

/**
 * 提交审核结果。
 * @param action 审核动作：approve 表示通过，reject 表示驳回。
 */
const submitReview = async (action: 'approve' | 'reject') => {
  const id = route.params.id
  if (!id) return

  const payload: ReviewInvoicePayload = { action }
  if (action === 'approve') {
    payload.final_amount = formData.final_amount
    payload.remarks = formData.remarks?.trim() || ''
  } else {
    payload.remarks = formData.remarks?.trim() || ''
  }

  submitting.value = true
  try {
    await processAudit(id as string, payload)
    ElMessage.success(action === 'approve' ? '审核通过成功' : '驳回成功')
    router.push('/audit/tasks')
  } catch (e: any) {
    ElMessage.error(e?.response?.data?.message || '审核处理失败')
  } finally {
    submitting.value = false
  }
}

const handleApprove = async () => {
  if (!formRef.value) return
  await formRef.value.validate(async (valid: boolean) => {
    if (valid) {
      await submitReview('approve')
    }
  })
}

/**
 * 处理驳回动作。
 * 驳回时要求审核备注非空，审核备注与驳回原因共用同一字段。
 */
const handleReject = () => {
  if (!formData.remarks?.trim()) {
    ElMessageBox.alert('驳回前请先填写审核备注。', '提示', {
      type: 'warning',
      confirmButtonText: '我知道了',
    })
    return
  }
  submitReview('reject')
}

onMounted(() => {
  fetchData()
})
</script>
