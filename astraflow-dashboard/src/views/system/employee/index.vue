<template>
  <div class="h-full flex flex-col">
    <!-- Header -->
    <div class="mb-6 flex items-center justify-between">
      <div>
        <h1 class="text-2xl font-bold text-slate-900 tracking-tight">Team Members</h1>
        <p class="text-slate-500 mt-1 text-sm">Manage access, roles, and status for your organization.</p>
      </div>
      <el-button type="primary" class="!bg-indigo-600 !border-indigo-600 hover:!bg-indigo-700" @click="handleAdd">
        <el-icon class="mr-2"><Plus /></el-icon>
        Add User
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
        <el-table-column label="Name" min-width="180">
          <template #default="{ row }">
            <div class="flex items-center gap-3">
              <el-avatar :size="36" :src="row.avatar || 'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png'" class="border border-slate-200" />
              <span class="text-slate-900 font-medium">{{ row.name }}</span>
            </div>
          </template>
        </el-table-column>

        <!-- Username -->
        <el-table-column prop="username" label="Username" min-width="150" />

        <!-- Role -->
        <el-table-column label="Role" width="150">
          <template #default="{ row }">
            <el-tag 
              :type="getRoleTagType(row.role_key)" 
              effect="light" 
              round
              class="capitalize !border-0 font-medium"
            >
              {{ row.role_key }}
            </el-tag>
          </template>
        </el-table-column>

        <!-- Created At -->
        <el-table-column label="Created At" width="180">
          <template #default="{ row }">
            {{ formatDate(row.created_at) }}
          </template>
        </el-table-column>

        <!-- Actions -->
        <el-table-column label="Operations" align="right" width="220">
          <template #default="{ row }">
             <el-button link type="primary" size="small" @click="handleResetPwd(row)">Reset Pwd</el-button>
             <el-divider direction="vertical" />
             <el-button 
              link 
              type="danger"
              size="small"
              @click="handleDelete(row)"
            >
              Delete
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- User Dialog -->
    <el-dialog
      v-model="dialogVisible"
      title="Add User"
      width="500px"
      class="!rounded-xl"
      destroy-on-close
    >
      <el-form ref="formRef" :model="formData" :rules="rules" label-position="top">
        <el-form-item label="Username" prop="username">
          <el-input v-model="formData.username" placeholder="e.g. john.doe" />
        </el-form-item>

        <el-form-item label="Full Name" prop="name">
          <el-input v-model="formData.name" placeholder="e.g. John Doe" />
        </el-form-item>
        
        <el-form-item label="Password" prop="password">
          <el-input v-model="formData.password" placeholder="Set initial password" show-password />
        </el-form-item>

        <el-form-item label="Role" prop="role_key">
          <el-select v-model="formData.role_key" placeholder="Select Role" class="w-full">
            <el-option label="Administrator" value="admin" />
            <el-option label="Finance Auditor" value="auditor" />
            <el-option label="Regular Employee" value="employee" />
          </el-select>
        </el-form-item>
      </el-form>
      
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">Cancel</el-button>
          <el-button type="primary" class="!bg-indigo-600 !border-indigo-600" :loading="submitting" @click="submitForm">
            Create
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
    ElMessage.error('Failed to load users')
  } finally {
    loading.value = false
  }
}

const formatDate = (date?: string) => {
    if (!date) return '-'
    return dayjs(date).format('YYYY-MM-DD HH:mm')
}

// -- Styling --
const getRoleTagType = (role: string) => {
    switch (role) {
        case 'admin': return 'danger'
        case 'auditor': return 'primary'
        case 'employee': return 'success'
        default: return 'info'
    }
}

// -- Actions --
const handleDelete = async (row: User) => {
    try {
        await ElMessageBox.confirm(
            `Are you sure you want to delete user "${row.name}"? This action cannot be undone.`,
            'Delete Confirmation',
            { confirmButtonText: 'Delete', cancelButtonText: 'Cancel', type: 'warning', confirmButtonClass: 'el-button--danger' }
        )
        
        await deleteUser(row.id)
        ElMessage.success('User deleted successfully')
        fetchData()
    } catch (e) {
        if (e !== 'cancel') ElMessage.error('Delete failed')
    }
}

const handleResetPwd = async (row: User) => {
    try {
        const { value } = await ElMessageBox.prompt('Please enter the new password', 'Reset Password', {
            confirmButtonText: 'Reset',
            cancelButtonText: 'Cancel',
            inputType: 'password',
            inputPattern: /.{6,}/,
            inputErrorMessage: 'Password must be at least 6 characters'
        })
        
        if (value) {
            await resetUserPassword(row.id, value)
            ElMessage.success('Password reset successfully')
        }
    } catch (e) {
         if (e !== 'cancel') ElMessage.error('Reset failed')
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
    username: [{ required: true, message: 'Username is required', trigger: 'blur' }],
    name: [{ required: true, message: 'Name is required', trigger: 'blur' }],
    password: [{ required: true, message: 'Password is required', trigger: 'blur' }],
    role_key: [{ required: true, message: 'Role is required', trigger: 'change' }]
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
                ElMessage.success('User created successfully')
                dialogVisible.value = false
                fetchData()
            } catch (e) {
                ElMessage.error('Failed to create user')
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
