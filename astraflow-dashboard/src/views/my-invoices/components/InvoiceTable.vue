<template>
  <div class="h-full flex flex-col">
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
          <!-- Placeholder until attachment relation is loaded -->
          <div class="w-16 h-16 rounded-lg bg-slate-100 flex items-center justify-center">
            <el-icon class="text-slate-400"><Picture /></el-icon>
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

      <el-table-column label="当前状态" width="120">
        <template #default="{ row }">
          <el-tag :type="getStatusTagType(row.status)" effect="light" round class="capitalize !border-0 font-medium">
            {{ getStatusText(row.status) }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column label="操作" align="right" width="150">
        <template #default="{ row }">
          <el-button v-if="status === 'unconfirmed'" type="primary" link @click="handleEdit(row)">确认/修改</el-button>
          <el-button v-if="status === 'draft'" type="primary" link @click="handleEdit(row)">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>

    <div class="p-4 border-t border-slate-100 flex justify-between items-center">
        <el-button v-if="status === 'draft' && selectedIds.length > 0" type="primary" class="!bg-indigo-600" @click="handlePublish">
            批量提交 ({{ selectedIds.length }})
        </el-button>
        <div v-else></div> <!-- Placeholder to keep pagination to the right -->
      <el-pagination
        background
        layout="prev, pager, next"
        :total="total"
        :page-size="pageSize"
        :current-page="currentPage"
        @current-change="handlePageChange"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { ElMessage } from 'element-plus';
import { getMyInvoices } from '../../../api/invoice';
import dayjs from 'dayjs';
import { Picture } from '@element-plus/icons-vue';

const props = defineProps<{
  status: 'recognizing' | 'unconfirmed' | 'draft' | 'submitted';
}>();

const loading = ref(false);
const tableData = ref([]);
const total = ref(0);
const currentPage = ref(1);
const pageSize = 10;
const selectedIds = ref<number[]>([]);

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

const handlePublish = async () => {
    // Placeholder for publish logic
    ElMessage.info('提交功能待实现');
};

const handleEdit = (row: any) => {
    // Placeholder for edit logic
    ElMessage.info(`编辑功能待实现 (ID: ${row.id})`);
}

const getStatusTagType = (status: string) => {
  const map = {
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
    const map = {
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
