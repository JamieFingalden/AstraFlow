<template>
  <div class="h-full flex flex-col relative">
    <el-table
      v-loading="loading"
      :data="tableData"
      style="width: 100%"
      class="flex-1"
      header-cell-class-name="!bg-slate-50 !text-slate-600 !font-semibold !border-b-slate-100"
      cell-class-name="!text-slate-600 !border-b-slate-50 hover:!bg-slate-50"
      @selection-change="handleSelectionChange"
    >
      <el-table-column v-if="status === 'draft'" type="selection" width="55" />

      <el-table-column label="发票图片" width="100">
        <template #default="{ row }">
          <div class="w-16 h-16 rounded-lg bg-slate-100 flex items-center justify-center overflow-hidden">
            <el-image
              v-if="row.attachment?.file_url"
              :src="row.attachment.file_url"
              class="w-full h-full object-cover"
              :preview-src-list="[row.attachment.file_url]"
              preview-teleported
            />
            <el-icon v-else class="text-slate-400"><Picture /></el-icon>
          </div>
        </template>
      </el-table-column>

      <el-table-column label="金额" min-width="120">
        <template #default="{ row }">
          <span v-if="row.amount > 0" class="font-bold text-slate-800">¥{{ row.amount.toFixed(2) }}</span>
          <el-tag v-else type="info" effect="plain" round size="small">待识别</el-tag>
        </template>
      </el-table-column>

      <el-table-column label="发票日期" min-width="140">
         <template #default="{ row }">
            {{ row.invoice_date ? dayjs(row.invoice_date).format('YYYY-MM-DD') : '待识别' }}
         </template>
      </el-table-column>
      
      <el-table-column label="报销类别" min-width="100">
         <template #default="{ row }">
            {{ row.category || '-' }}
         </template>
      </el-table-column>

      <el-table-column label="当前状态" width="120">
        <template #default="{ row }">
          <el-tag :type="getStatusTagType(row.status)" effect="light" round class="capitalize !border-0 font-medium">
            {{ getStatusText(row.status) }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column label="操作" align="right" width="150" fixed="right">
        <template #default="{ row }">
          <el-button v-if="status === 'unconfirmed'" type="primary" link @click="handleEdit(row)">确认信息</el-button>
          <el-button v-if="status === 'draft'" type="primary" link @click="handleEdit(row)">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>

    <div class="p-4 border-t border-slate-100 flex justify-between items-center">
        <div class="flex items-center">
            <el-button 
                v-if="status === 'draft'" 
                type="primary" 
                class="!bg-indigo-600" 
                :disabled="selectedIds.length === 0"
                @click="handlePublish"
            >
                批量提交 ({{ selectedIds.length }})
            </el-button>
        </div>
      <el-pagination
        background
        layout="prev, pager, next"
        :total="total"
        :page-size="pageSize"
        :current-page="currentPage"
        @current-change="handlePageChange"
      />
    </div>

    <!-- 编辑/确认弹窗 -->
    <el-dialog
      v-model="dialogVisible"
      :title="status === 'unconfirmed' ? '确认发票信息' : '编辑发票'"
      width="500px"
    >
      <el-form :model="editForm" :rules="rules" ref="formRef" label-width="100px">
        <el-form-item label="发票金额" prop="amount">
          <el-input-number v-model="editForm.amount" :precision="2" :step="10" class="!w-full" />
        </el-form-item>
        <el-form-item label="发票日期" prop="invoice_date">
          <el-date-picker v-model="editForm.invoice_date" type="date" placeholder="选择日期" value-format="YYYY-MM-DD" class="!w-full" />
        </el-form-item>
        <el-form-item label="发票号码" prop="invoice_number">
          <el-input v-model="editForm.invoice_number" placeholder="请输入发票号码" />
        </el-form-item>
        <el-form-item label="商户名称" prop="vendor">
          <el-input v-model="editForm.vendor" placeholder="请输入商户名称" />
        </el-form-item>
        <el-form-item label="报销类别" prop="category">
          <el-select v-model="editForm.category" placeholder="请选择报销类别" class="!w-full">
            <el-option label="餐饮" value="餐饮" />
            <el-option label="交通" value="交通" />
            <el-option label="住宿" value="住宿" />
            <el-option label="办公用品" value="办公用品" />
            <el-option label="其他" value="其他" />
          </el-select>
        </el-form-item>
        <el-form-item label="备注说明" prop="description">
          <el-input v-model="editForm.description" type="textarea" placeholder="可选填" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" :loading="saving" @click="handleSave">
            {{ status === 'unconfirmed' ? '确认并移入草稿箱' : '保存' }}
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import { ElMessage, ElMessageBox } from 'element-plus';
import type { FormInstance, FormRules } from 'element-plus';
import { getMyInvoices, confirmInvoice, publishInvoices, updateInvoice } from '../../../api/invoice';
import dayjs from 'dayjs';
import { Picture } from '@element-plus/icons-vue';

const props = defineProps<{
  status: 'recognizing' | 'unconfirmed' | 'draft' | 'submitted';
}>();

const emit = defineEmits(['refresh-counts']);

const loading = ref(false);
const tableData = ref([]);
const total = ref(0);
const currentPage = ref(1);
const pageSize = 10;
const selectedIds = ref<number[]>([]);

// 弹窗与表单
const dialogVisible = ref(false);
const saving = ref(false);
const formRef = ref<FormInstance>();
const editForm = reactive({
  id: 0,
  amount: 0,
  invoice_date: '',
  invoice_number: '',
  vendor: '',
  category: '',
  description: '',
});

const rules = reactive<FormRules>({
  amount: [{ required: true, message: '请输入发票金额', trigger: 'blur' }],
  invoice_date: [{ required: true, message: '请选择发票日期', trigger: 'change' }],
  vendor: [{ required: true, message: '请输入商户名称', trigger: 'blur' }],
  category: [{ required: true, message: '请选择报销类别', trigger: 'change' }],
});

const fetchData = async () => {
  loading.value = true;
  try {
    const params = {
      status: props.status,
      page: currentPage.value,
      pageSize: pageSize,
    };
    const res = await getMyInvoices(params);
    tableData.value = res.data.items || [];
    total.value = res.data.total || 0;
  } catch (error) {
    ElMessage.error(`获取[${props.status}]列表失败`);
  } finally {
    loading.value = false;
  }
};

const handlePageChange = (newPage: number) => {
  currentPage.value = newPage;
  fetchData();
};

const handleSelectionChange = (selection: any[]) => {
  selectedIds.value = selection.map(item => item.id);
};

// 批量提交
const handlePublish = async () => {
  if (selectedIds.value.length === 0) return;
  try {
    await ElMessageBox.confirm(`确定要提交这 ${selectedIds.value.length} 张单据进行审核吗？`, '提示', {
      type: 'warning',
      confirmButtonText: '确定提交',
      cancelButtonText: '取消',
    });
    
    loading.value = true;
    await publishInvoices(selectedIds.value);
    ElMessage.success('提交成功');
    fetchData();
    emit('refresh-counts'); // 通知父组件更新角标数量
  } catch (e: any) {
    if (e !== 'cancel') {
      ElMessage.error(e.message || '提交失败');
    }
  } finally {
    loading.value = false;
  }
};

// 打开编辑/确认弹窗
const handleEdit = (row: any) => {
  editForm.id = row.id;
  editForm.amount = row.amount || 0;
  editForm.invoice_date = row.invoice_date ? dayjs(row.invoice_date).format('YYYY-MM-DD') : '';
  editForm.invoice_number = row.invoice_number || '';
  editForm.vendor = row.vendor || '';
  editForm.category = row.category || '';
  editForm.description = row.description || '';
  
  if (formRef.value) {
    formRef.value.clearValidate();
  }
  dialogVisible.value = true;
};

// 保存表单
const handleSave = async () => {
  if (!formRef.value) return;
  await formRef.value.validate(async (valid, fields) => {
    if (valid) {
      saving.value = true;
      try {
        const payload = {
            amount: editForm.amount,
            invoice_date: editForm.invoice_date ? new Date(editForm.invoice_date).toISOString() : undefined,
            invoice_number: editForm.invoice_number,
            vendor: editForm.vendor,
            category: editForm.category,
            description: editForm.description,
        };

        if (props.status === 'unconfirmed') {
          await confirmInvoice(editForm.id, payload);
          ElMessage.success('确认成功，单据已移入待发布(草稿箱)');
        } else {
          await updateInvoice(editForm.id, payload);
          ElMessage.success('更新成功');
        }
        dialogVisible.value = false;
        fetchData();
        emit('refresh-counts'); // 通知父组件更新角标数量
      } catch (error: any) {
        ElMessage.error(error.response?.data?.message || '操作失败');
      } finally {
        saving.value = false;
      }
    }
  });
};

const getStatusTagType = (status: string) => {
  const map: Record<string, string> = {
    recognizing: 'info',
    unconfirmed: 'warning',
    draft: 'primary',
    pending: 'warning',
    approved: 'success',
    rejected: 'danger',
    paid: 'info',
  };
  return map[status] || 'info';
};

const getStatusText = (status: string) => {
    const map: Record<string, string> = {
        recognizing: '待识别',
        unconfirmed: '待确认',
        draft: '待发布',
        pending: '审核中',
        approved: '已批准',
        rejected: '已驳回',
        paid: '已支付',
    };
    return map[status] || status;
}

onMounted(() => {
  fetchData();
});
</script>
