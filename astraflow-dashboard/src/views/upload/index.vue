<template>
  <div class="h-full flex flex-col items-center justify-center bg-slate-50 p-4">
    <div class="w-full max-w-2xl">
      <!-- Header -->
      <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-slate-900 tracking-tight">上传新的发票</h1>
        <p class="text-slate-500 mt-2">选择一种方式开始你的报销流程。</p>
      </div>

      <!-- Upload Card -->
      <div class="bg-white rounded-2xl shadow-lg border border-slate-200 p-8">
        <el-upload
          ref="uploadRef"
          drag
          :auto-upload="false"
          :on-change="handleFileChange"
          :limit="1"
          :on-exceed="handleExceed"
          class="w-full"
        >
          <el-icon class="el-icon--upload text-slate-400" :size="60"><upload-filled /></el-icon>
          <div class="el-upload__text">
            将文件拖到此处, 或 <em>点击上传</em>
          </div>
          <template #tip>
            <div class="el-upload__tip text-center mt-4">
              支持 JPG, PNG, PDF 等格式, 文件大小不超过 10MB。
            </div>
          </template>
        </el-upload>
        
        <!-- Form for manual mode -->
        <div v-if="selectedFile && uploadMode === 'manual'" class="mt-6 space-y-4 pt-4 border-t border-slate-100">
           <div class="grid grid-cols-2 gap-4">
              <el-form-item label="发票金额">
                <el-input-number v-model="manualForm.amount" :precision="2" :step="1" class="!w-full" />
              </el-form-item>
              <el-form-item label="发票日期">
                <el-date-picker v-model="manualForm.invoice_date" type="date" placeholder="选择日期" class="!w-full" format="YYYY-MM-DD" value-format="YYYY-MM-DD" />
              </el-form-item>
           </div>
           <el-form-item label="消费类别">
               <el-select v-model="manualForm.category" placeholder="请选择类别" class="!w-full">
                  <el-option label="办公用品" value="Office Supplies" />
                  <el-option label="差旅交通" value="Travel" />
                  <el-option label="餐饮" value="Meals" />
                  <el-option label="软件服务" value="Software" />
                  <el-option label="其他" value="Other" />
               </el-select>
            </el-form-item>
            <el-form-item label="费用描述 (选填)">
              <el-input v-model="manualForm.description" type="textarea" :rows="2" placeholder="简单描述一下费用" />
            </el-form-item>
        </div>
      </div>

      <!-- Action Buttons -->
      <div v-if="selectedFile" class="mt-6 flex gap-4">
        <el-button 
          type="primary" 
          class="!flex-1 !h-12 !text-base !font-medium !bg-indigo-600 !border-indigo-600" 
          :loading="isOcrLoading"
          @click="handleOcrUpload"
        >
          <el-icon class="mr-2"><MagicStick /></el-icon>
          智能识别
        </el-button>
        <el-button 
          type="success"
          class="!flex-1 !h-12 !text-base !font-medium"
          :loading="isManualLoading"
          @click="handleManualUpload"
        >
          <el-icon class="mr-2"><DocumentAdd /></el-icon>
          存入草稿箱
        </el-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { UploadFilled, MagicStick, DocumentAdd } from '@element-plus/icons-vue'
import { ElMessage, genFileId } from 'element-plus'
import type { UploadInstance, UploadProps, UploadRawFile } from 'element-plus'
// 注意：我们需要在 astraflow-dashboard 中创建这个 api/invoice.ts 文件
import { uploadInvoiceForOCR, uploadInvoiceManually } from '../../api/invoice' 

const uploadRef = ref<UploadInstance>()
const selectedFile = ref<UploadRawFile | null>(null)
const uploadMode = ref<'ocr' | 'manual'>('ocr') // Default to ocr, changes when manual form is shown

const isOcrLoading = ref(false)
const isManualLoading = ref(false)

const manualForm = reactive({
  amount: 0,
  invoice_date: '',
  category: '',
  description: ''
})

const handleFileChange: UploadProps['onChange'] = (uploadFile) => {
  selectedFile.value = uploadFile.raw!
  uploadMode.value = 'manual'
}

const handleExceed: UploadProps['onExceed'] = (files) => {
  uploadRef.value!.clearFiles()
  const file = files[0] as UploadRawFile
  file.uid = genFileId()
  uploadRef.value!.handleStart(file)
  selectedFile.value = file
  uploadMode.value = 'manual'
}

const resetForm = () => {
    uploadRef.value?.clearFiles()
    selectedFile.value = null
    manualForm.amount = 0
    manualForm.invoice_date = ''
    manualForm.category = ''
    manualForm.description = ''
}

const handleOcrUpload = async () => {
  if (!selectedFile.value) {
    ElMessage.warning('请先选择一个文件')
    return
  }
  isOcrLoading.value = true
  try {
    const res = await uploadInvoiceForOCR(selectedFile.value)
    ElMessage.success('上传成功，已加入 AI 识别队列！')
    resetForm()
  } catch (error) {
    ElMessage.error('智能识别请求失败')
  } finally {
    isOcrLoading.value = false
  }
}

const handleManualUpload = async () => {
  if (!selectedFile.value) {
    ElMessage.warning('请先选择一个文件')
    return
  }
  if (!manualForm.amount || !manualForm.invoice_date || !manualForm.category) {
      ElMessage.warning('请填写所有必填项')
      return
  }
  isManualLoading.value = true
  try {
    await uploadInvoiceManually({
      file: selectedFile.value,
      ...manualForm
    })
    ElMessage.success('发票已成功存入草稿箱！')
    resetForm()
  } catch (error) {
    ElMessage.error('存入草稿箱失败')
  } finally {
    isManualLoading.value = false
  }
}
</script>

<style scoped>
:deep(.el-upload-dragger) {
  padding: 40px 20px;
  border-radius: 1rem;
  transition: all 0.2s;
}
:deep(.el-upload-dragger:hover) {
    border-color: #4f46e5; /* Indigo-600 */
}
</style>
