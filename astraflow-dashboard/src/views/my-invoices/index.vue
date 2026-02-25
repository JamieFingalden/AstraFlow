<template>
  <div class="h-full flex flex-col p-6 bg-slate-50">
    <div class="bg-white rounded-xl shadow-sm border border-slate-200 flex-1 overflow-hidden flex flex-col">
      <div class="p-6 border-b border-slate-100">
        <h1 class="text-2xl font-bold text-slate-900 tracking-tight">我的票夹</h1>
        <p class="text-slate-500 mt-1 text-sm">管理我上传的发票，提交或修改。</p>
      </div>
      
      <el-tabs v-model="activeTab" class="flex-1 flex flex-col -mb-px">
        <el-tab-pane :label="`待识别 (${counts.recognizing})`" name="recognizing" class="h-full" lazy>
           <InvoiceTable v-if="activeTab === 'recognizing'" :status="activeTab" @refresh-counts="fetchCounts" />
        </el-tab-pane>
        <el-tab-pane :label="`待确认 (${counts.unconfirmed})`" name="unconfirmed" class="h-full" lazy>
            <InvoiceTable v-if="activeTab === 'unconfirmed'" :status="activeTab" @refresh-counts="fetchCounts" />
        </el-tab-pane>
        <el-tab-pane :label="`待发布 (${counts.draft})`" name="draft" class="h-full" lazy>
            <InvoiceTable v-if="activeTab === 'draft'" :status="activeTab" @refresh-counts="fetchCounts" />
        </el-tab-pane>
         <el-tab-pane :label="`已提交 (${counts.submitted})`" name="submitted" class="h-full" lazy>
            <InvoiceTable v-if="activeTab === 'submitted'" :status="activeTab" @refresh-counts="fetchCounts" />
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import InvoiceTable from './components/InvoiceTable.vue';
import { getMyInvoices } from '../../api/invoice';

const activeTab = ref('recognizing');

const counts = reactive({
  recognizing: 0,
  unconfirmed: 0,
  draft: 0,
  submitted: 0,
});

// Fetch counts for all tabs to display in badges
const fetchCounts = async () => {
    try {
        const statuses = ['recognizing', 'unconfirmed', 'draft', 'submitted'];
        const requests = statuses.map(status => getMyInvoices({ status, page: 1, pageSize: 1 }));
        const results = await Promise.all(requests);
        
        counts.recognizing = results[0].data.total || 0;
        counts.unconfirmed = results[1].data.total || 0;
        counts.draft = results[2].data.total || 0;
        counts.submitted = results[3].data.total || 0;
    } catch (e) {
        console.error("Failed to fetch invoice counts", e);
    }
};

onMounted(() => {
    // Initial counts fetch on page load
    fetchCounts();
});
</script>

<style scoped>
:deep(.el-tabs__header) {
    padding: 0 24px;
    margin-bottom: 0;
}
:deep(.el-tabs__content) {
    flex: 1;
    overflow-y: auto;
    padding: 24px;
}
:deep(.el-tab-pane) {
    height: 100%;
}
</style>
