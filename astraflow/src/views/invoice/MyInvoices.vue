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
           <InvoiceList :invoices="invoiceData.items" :loading="loading" status="unconfirmed" @confirm="handleConfirm" />
        </el-tab-pane>
        <el-tab-pane label="待发布" name="draft">
           <InvoiceList :invoices="invoiceData.items" :loading="loading" status="draft" @publish="handlePublish" />
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
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import { getMyInvoices, publishInvoices } from '../../services/api/invoice.js';
import { ElMessage } from 'element-plus';
import InvoiceList from '../../components/InvoiceList.vue'; // We will create this component next

const activeTab = ref('recognizing');
const loading = ref(false);
const invoiceData = reactive({
  items: [],
  total: 0,
  page: 1,
  pageSize: 10,
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

const handleConfirm = (invoice) => {
    console.log('Confirming invoice:', invoice);
    // TODO: Implement confirmation logic (e.g., show a dialog, call confirm API)
    ElMessage.success(`发票 #${invoice.id} 已确认并移至待发布`);
    fetchData(); // Refresh list
};

const handlePublish = async (selectedIds) => {
    if (selectedIds.length === 0) {
        ElMessage.warning('请至少选择一张发票进行提交');
        return;
    }
    try {
        await publishInvoices(selectedIds);
        ElMessage.success('成功提交，请等待财务审核');
        fetchData(); // Refresh list
    } catch(e) {
        ElMessage.error('提交失败');
    }
};

onMounted(() => {
  fetchData();
});
</script>
