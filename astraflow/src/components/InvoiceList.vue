<template>
  <div>
    <el-table :data="invoices" v-loading="loading" @selection-change="handleSelectionChange">
      <!-- Checkbox for 'draft' status -->
      <el-table-column v-if="status === 'draft'" type="selection" width="55" />

      <el-table-column label="发票图片" width="120">
        <template #default="{ row }">
          <el-image 
            :src="row.attachment.file_url" 
            :preview-src-list="[row.attachment.file_url]"
            class="w-16 h-16 rounded-md object-cover"
            fit="cover"
          />
        </template>
      </el-table-column>
      
      <el-table-column label="金额" prop="amount">
         <template #default="{ row }">
            <span v-if="row.amount > 0">￥{{ row.amount.toFixed(2) }}</span>
            <el-tag v-else type="info" size="small">待识别</el-tag>
         </template>
      </el-table-column>
      
      <el-table-column label="发票日期" prop="invoice_date">
        <template #default="{ row }">
            {{ row.invoice_date ? formatDate(row.invoice_date) : '待识别' }}
        </template>
      </el-table-column>

      <el-table-column label="状态" prop="status">
        <template #default="{ row }">
          <el-tag v-if="isDuplicateInvoice(row)" type="danger" effect="light">已存在</el-tag>
          <el-tag v-else :type="getStatusTagType(row.status)" effect="light">{{ getStatusText(row.status) }}</el-tag>
        </template>
      </el-table-column>

      <el-table-column label="操作" align="right">
        <template #default="{ row }">
          <div v-if="status === 'unconfirmed'">
            <el-button v-if="!isDuplicateInvoice(row)" type="primary" link @click="$emit('confirm', row)">确认/修改</el-button>
            <el-button v-else type="danger" link @click="$emit('delete', row)">删除</el-button>
          </div>
           <div v-if="status === 'draft'">
            <el-button type="primary" link @click="$emit('edit', row)">编辑</el-button>
            <el-button type="danger" link>删除</el-button>
          </div>
          <div v-if="status === 'recognizing'">
             <el-icon class="is-loading"><Loading /></el-icon>
             <span class="ml-2 text-gray-400 text-xs">识别中...</span>
          </div>
        </template>
      </el-table-column>
    </el-table>
    
    <div v-if="status === 'draft'" class="mt-4 flex justify-start">
        <el-button type="primary" @click="publishSelected">批量提交审核</el-button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { Loading } from '@element-plus/icons-vue';

const props = defineProps({
  invoices: {
    type: Array,
    required: true,
  },
  loading: {
    type: Boolean,
    default: false,
  },
  status: {
    type: String,
    required: true,
  }
});

const emit = defineEmits(['confirm', 'publish', 'delete']);

const selectedInvoices = ref([]);

const handleSelectionChange = (val) => {
  selectedInvoices.value = val;
};

const publishSelected = () => {
    const ids = selectedInvoices.value.map(inv => inv.id);
    emit('publish', ids);
};

const formatDate = (dateString) => {
  const date = new Date(dateString);
  return date.toLocaleDateString();
};

const getStatusTagType = (status) => {
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

const getStatusText = (status) => {
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

const isDuplicateInvoice = (invoice) => {
  if (!invoice?.pre_audit_reasons) return false;

  try {
    const reasons = JSON.parse(invoice.pre_audit_reasons);
    return Array.isArray(reasons) && reasons.includes('发票号已存在');
  } catch (error) {
    return String(invoice.pre_audit_reasons).includes('发票号已存在');
  }
};
</script>
