<template>
  <div class="h-full flex flex-col">
    <!-- Header -->
    <div class="mb-6 flex items-center justify-between">
      <div>
        <h1 class="text-2xl font-bold text-slate-900 tracking-tight">员工管理</h1>
        <p class="text-slate-500 mt-1 text-sm">管理企业成员的访问权限、角色和状态。</p>
      </div>
      <el-button type="primary" class="!bg-indigo-600 !border-indigo-600 hover:!bg-indigo-700" @click="handleAdd">
        <el-icon class="mr-2"><Plus /></el-icon>
        添加成员
      </el-button>
    </div>

    <!-- Table Card -->
    <div class="bg-white rounded-xl shadow-sm border border-slate-200 flex-1 overflow-hidden flex flex-col">
      <el-table 
        v-loading="loading"
        :data="tableData" 
        style="width: 100%" 
        class="flex-1"
        header-cell-class-name="!bg-slate-50 !text-slate-600 !font-semibold !border-b-slate-100"
        cell-class-name="!text-slate-600 !border-b-slate-50 hover:!bg-slate-50"
      >
        <!-- Avatar & Name -->
        <el-table-column label="姓名" min-width="180">
          <template #default="{ row }">
            <div class="flex items-center gap-3">
              <el-avatar :size="36" :src="row.avatar || 'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png'" class="border border-slate-200" />
              <span class="text-slate-900 font-medium">{{ row.name }}</span>
            </div>
          </template>
        </el-table-column>

        <!-- Username -->
        <el-table-column prop="username" label="账号" min-width="150" />

        <!-- Role -->
        <el-table-column label="角色" width="150">
          <template #default="{ row }">
            <el-tag 
              :type="getRoleTagType(getUserRoleKey(row))" 
              effect="light" 
              round
              class="!border-0 !font-semibold role-tag"
            >
              {{ getRoleName(row) }}
            </el-tag>
          </template>
        </el-table-column>

        <!-- Created At -->
        <el-table-column label="创建时间" width="180">
          <template #default="{ row }">
            {{ formatDate(row.created_at) }}
          </template>
        </el-table-column>

        <!-- Actions -->
        <el-table-column label="操作" align="right" width="220">
          <template #default="{ row }">
             <el-button link type="primary" size="small" @click="handleResetPwd(row)">重置密码</el-button>
             <el-divider direction="vertical" />
             <el-button 
              link 
              type="danger"
              size="small"
              @click="handleDelete(row)"
            >
              移除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- User Dialog -->
    <el-dialog
      v-model="dialogVisible"
      title="添加成员"
      width="500px"
      class="!rounded-xl"
      destroy-on-close
    >
      <el-form ref="formRef" :model="formData" :rules="rules" label-position="top">
        <el-form-item label="账号" prop="username">
          <el-input v-model="formData.username" placeholder="请输入账号" />
        </el-form-item>

        <el-form-item label="姓名" prop="name">
          <el-input v-model="formData.name" placeholder="请输入姓名" />
        </el-form-item>
        
        <el-form-item label="密码" prop="password">
          <el-input v-model="formData.password" placeholder="设置初始密码" show-password />
        </el-form-item>

        <el-form-item label="角色" prop="role_key">
          <el-select v-model="formData.role_key" placeholder="选择角色" class="w-full">
            <el-option label="人事管理员" value="admin" />
            <el-option label="财务专员" value="auditor" />
            <el-option label="普通员工" value="employee" />
          </el-select>
        </el-form-item>
      </el-form>
      
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" class="!bg-indigo-600 !border-indigo-600" :loading="submitting" @click="submitForm">
            确认
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { Plus } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'
import { getEmployeeList, createUser, deleteUser, resetUserPassword } from '../../../api/user'
import type { User, CreateUserParams } from '../../../api/user'
import dayjs from 'dayjs'

// -- Table Data --
const loading = ref(false)
const tableData = ref<User[]>([])

const fetchData = async () => {
  loading.value = true
  try {
    const res = await getEmployeeList()
    // Assuming API returns standard wrapper or data directly based on request.ts
    tableData.value = Array.isArray(res) ? res : (res.data || [])
  } catch (err) {
    ElMessage.error('获取用户列表失败')
  } finally {
    loading.value = false
  }
}

const formatDate = (date?: string) => {
    if (!date) return '-'
    return dayjs(date).format('YYYY-MM-DD HH:mm')
}

// -- Styling --
const getUserRoleKey = (user: User) => user.role_key || user.role?.key || ''

const getRoleTagType = (role: string) => {
    switch (role) {
        case 'admin': return 'danger'
        case 'auditor': return 'primary'
        case 'employee': return 'success'
        default: return 'info'
    }
}

const getRoleName = (user: User) => {
    const role = getUserRoleKey(user)
    switch (role) {
        case 'admin': return '人事管理员'
        case 'auditor': return '财务专员'
        case 'employee': return '普通员工'
        default: return user.role?.name || role || '未分配'
    }
}

// -- Actions --
const handleDelete = async (row: User) => {
    try {
        await ElMessageBox.confirm(
            `确认要移除用户 "${row.name}" 吗? 此操作无法撤销。`,
            '删除确认',
            { confirmButtonText: '移除', cancelButtonText: '取消', type: 'warning', confirmButtonClass: 'el-button--danger' }
        )
        
        await deleteUser(row.id)
        ElMessage.success('用户已移除')
        fetchData()
    } catch (e) {
        if (e !== 'cancel') ElMessage.error('删除失败')
    }
}

const handleResetPwd = async (row: User) => {
    try {
        const { value } = await ElMessageBox.prompt('请输入新密码', '重置密码', {
            confirmButtonText: '重置',
            cancelButtonText: '取消',
            inputType: 'password',
            inputPattern: /.{6,}/,
            inputErrorMessage: '密码至少需要6个字符'
        })
        
        if (value) {
            await resetUserPassword(row.id, value)
            ElMessage.success('密码重置成功')
        }
    } catch (e) {
         if (e !== 'cancel') ElMessage.error('重置失败')
    }
}

// -- Dialog Logic --
const dialogVisible = ref(false)
const submitting = ref(false)
const formRef = ref<FormInstance>()

const formData = reactive<CreateUserParams>({
    username: '',
    name: '',
    password: '',
    role_key: 'employee'
})

const rules = reactive<FormRules>({
    username: [{ required: true, message: '请输入账号', trigger: 'blur' }],
    name: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
    password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
    role_key: [{ required: true, message: '请选择角色', trigger: 'change' }]
})

const handleAdd = () => {
    formData.username = ''
    formData.name = ''
    formData.password = ''
    formData.role_key = 'employee'
    dialogVisible.value = true
}

const submitForm = async () => {
    if (!formRef.value) return
    
    await formRef.value.validate(async (valid: boolean) => {
        if (valid) {
            submitting.value = true
            try {
                await createUser(formData)
                ElMessage.success('用户创建成功')
                dialogVisible.value = false
                fetchData()
            } catch (e) {
                ElMessage.error('创建用户失败')
            } finally {
                submitting.value = false
            }
        }
    })
}

onMounted(() => {
    fetchData()
})
</script>

<style scoped>
.role-tag {
  min-width: 72px;
  justify-content: center;
}
</style>
