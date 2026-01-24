<template>
  <el-card>
    <template #header>
        <div class="flex justify-between items-center">
            <span>Audit Task Pool</span>
            <el-button type="primary">Refresh</el-button>
        </div>
    </template>
    <el-table :data="[]" style="width: 100%">
        <el-table-column prop="date" label="Date" width="180" />
        <el-table-column prop="name" label="Name" width="180" />
        <el-table-column prop="status" label="Status" />
    </el-table>
    <el-empty description="No tasks found" />
  </el-card>
</template>
