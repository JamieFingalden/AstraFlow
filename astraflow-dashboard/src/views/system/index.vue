<template>
  <el-card>
    <template #header>
        <div class="flex justify-between items-center">
            <span>Employee Management</span>
            <el-button type="primary">Add Employee</el-button>
        </div>
    </template>
    <el-table :data="[]" style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="Name" width="180" />
        <el-table-column prop="role" label="Role" />
    </el-table>
    <el-empty description="No employees found" />
  </el-card>
</template>
