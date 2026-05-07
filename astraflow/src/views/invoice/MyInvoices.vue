<template>
  <div class="p-6 bg-gray-50 min-h-full">
    <el-card class="rounded-xl">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-xl font-bold">我的票夹</span>
        </div>
      </template>

      <el-tabs v-model="activeTab" @tab-click="handleTabClick">
        <el-tab-pane label="待识别" name="recognizing">
          <InvoiceList :invoices="invoiceData.items" :loading="loading" status="recognizing" />
        </el-tab-pane>
        <el-tab-pane label="待确认" name="unconfirmed">
           <InvoiceList :invoices="invoiceData.items" :loading="loading" status="unconfirmed" @confirm="handleConfirm" @delete="handleDelete" />
        </el-tab-pane>
        <el-tab-pane label="待发布" name="draft">
           <InvoiceList :invoices="invoiceData.items" :loading="loading" status="draft" @publish="handlePublish" @edit="handleEdit" />
        </el-tab-pane>
        <el-tab-pane label="已提交" name="submitted">
           <InvoiceList :invoices="invoiceData.items" :loading="loading" status="submitted" />
        </el-tab-pane>
      </el-tabs>

      <div class="flex justify-end mt-4">
        <el-pagination
          background
          layout="prev, pager, next"
          :total="invoiceData.total"
          :page-size="invoiceData.pageSize"
          :current-page="invoiceData.page"
          @current-change="handlePageChange"
        />
      </div>
    </el-card>

    <!-- 编辑/确认弹窗 -->
    <el-dialog
      v-model="dialogVisible"
      :title="activeTab === 'unconfirmed' ? '确认发票信息' : '编辑发票'"
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
            {{ activeTab === 'unconfirmed' ? '确认并移入草稿箱' : '保存' }}
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import { getMyInvoices, publishInvoices, confirmInvoice, updateInvoice, deleteInvoice } from '../../services/api/invoice.js';
import { ElMessage, ElMessageBox } from 'element-plus';
import InvoiceList from '../../components/InvoiceList.vue';

const activeTab = ref('recognizing');
const loading = ref(false);
const invoiceData = reactive({
  items: [],
  total: 0,
  page: 1,
  pageSize: 10,
});

const dialogVisible = ref(false);
const saving = ref(false);
const formRef = ref(null);
const editForm = reactive({
  id: 0,
  amount: 0,
  invoice_date: '',
  invoice_number: '',
  vendor: '',
  category: '',
  description: '',
});

const rules = reactive({
  amount: [{ required: true, message: '请输入发票金额', trigger: 'blur' }],
  invoice_date: [{ required: true, message: '请选择发票日期', trigger: 'change' }],
  vendor: [{ required: true, message: '请输入商户名称', trigger: 'blur' }],
  category: [{ required: true, message: '请选择报销类别', trigger: 'change' }],
});

const fetchData = async () => {
  loading.value = true;
  try {
    const params = {
      status: activeTab.value,
      page: invoiceData.page,
      pageSize: invoiceData.pageSize,
    };
    const res = await getMyInvoices(params);
    invoiceData.items = res.data.items;
    invoiceData.total = res.data.total;
  } catch (error) {
    ElMessage.error('获取发票列表失败');
    console.error(error);
  } finally {
    loading.value = false;
  }
};

const handleTabClick = (tab) => {
  invoiceData.page = 1; // Reset page on tab change
  fetchData();
};

const handlePageChange = (newPage) => {
  invoiceData.page = newPage;
  fetchData();
};

const openDialog = (invoice) => {
  editForm.id = invoice.id;
  editForm.amount = invoice.amount || 0;
  // Format date correctly to YYYY-MM-DD
  if (invoice.invoice_date) {
      const date = new Date(invoice.invoice_date);
      const year = date.getFullYear();
      const month = String(date.getMonth() + 1).padStart(2, '0');
      const day = String(date.getDate()).padStart(2, '0');
      editForm.invoice_date = `${year}-${month}-${day}`;
  } else {
      editForm.invoice_date = '';
  }
  editForm.invoice_number = invoice.invoice_number || '';
  editForm.vendor = invoice.vendor || '';
  editForm.category = invoice.category || '';
  editForm.description = invoice.description || '';
  
  if (formRef.value) {
    formRef.value.clearValidate();
  }
  dialogVisible.value = true;
};

const handleConfirm = (invoice) => {
    openDialog(invoice);
};

const handleEdit = (invoice) => {
    openDialog(invoice);
};

const handleDelete = async (invoice) => {
  try {
    await ElMessageBox.confirm('该发票号已存在，只能删除当前识别记录。确定删除吗？', '删除发票', {
      type: 'warning',
      confirmButtonText: '确定删除',
      cancelButtonText: '取消',
    });

    loading.value = true;
    await deleteInvoice(invoice.id);
    ElMessage.success('删除成功');
    fetchData();
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error(error.response?.data?.message || '删除失败');
    }
  } finally {
    loading.value = false;
  }
};

const handleSave = async () => {
  if (!formRef.value) return;
  await formRef.value.validate(async (valid) => {
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

        if (activeTab.value === 'unconfirmed') {
          await confirmInvoice(editForm.id, payload);
          ElMessage.success('确认成功，单据已移入待发布(草稿箱)');
        } else {
          await updateInvoice(editForm.id, payload);
          ElMessage.success('更新成功');
        }
        dialogVisible.value = false;
        fetchData();
      } catch (error) {
        ElMessage.error(error.response?.data?.message || '操作失败');
      } finally {
        saving.value = false;
      }
    }
  });
};

const handlePublish = async (selectedIds) => {
    if (selectedIds.length === 0) {
        ElMessage.warning('请至少选择一张发票进行提交');
        return;
    }
    
    try {
        await ElMessageBox.confirm(`确定要提交这 ${selectedIds.length} 张单据进行审核吗？`, '提示', {
            type: 'warning',
            confirmButtonText: '确定提交',
            cancelButtonText: '取消',
        });
        
        loading.value = true;
        await publishInvoices(selectedIds);
        ElMessage.success('成功提交，请等待财务审核');
        fetchData(); // Refresh list
    } catch(e) {
        if (e !== 'cancel') {
            ElMessage.error(e.message || '提交失败');
        }
    } finally {
        loading.value = false;
    }
};

onMounted(() => {
  fetchData();
});
</script>

<style scoped>
/* Scoped styles here if needed */
:deep(.el-date-editor.el-input), :deep(.el-date-editor.el-input__wrapper) {
    width: 100% !important;
}
</style>
