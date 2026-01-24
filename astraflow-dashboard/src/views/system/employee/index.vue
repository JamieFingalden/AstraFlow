<template>
  <div class="h-full flex flex-col">
    <!-- Header -->
    <div class="mb-6 flex items-center justify-between">
      <div>
        <h1 class="text-2xl font-bold text-slate-900 tracking-tight">Team Members</h1>
        <p class="text-slate-500 mt-1 text-sm">Manage access, roles, and status for your organization.</p>
      </div>
      <el-button type="primary" class="!bg-indigo-600 !border-indigo-600 hover:!bg-indigo-700" @click="handleInvite">
        <el-icon class="mr-2"><Plus /></el-icon>
        Create User
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
        <!-- Member Column -->
        <el-table-column label="Member" min-width="240">
          <template #default="{ row }">
            <div class="flex items-center gap-3">
              <el-avatar :size="40" :src="row.avatar || 'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png'" class="border border-slate-200" />
              <div class="flex flex-col">
                <span class="text-slate-900 font-medium leading-tight">{{ row.name }}</span>
                <span class="text-slate-400 text-xs mt-0.5">{{ row.username }}</span>
              </div>
            </div>
          </template>
        </el-table-column>

        <!-- Role Column -->
        <el-table-column label="Role" width="150">
          <template #default="{ row }">
            <el-tag 
              :type="getRoleTagType(row.role_key)" 
              effect="light" 
              round
              class="capitalize !border-0 font-medium"
              :class="getRoleTagClass(row.role_key)"
            >
              {{ row.role_key }}
            </el-tag>
          </template>
        </el-table-column>

        <!-- Status Column -->
        <el-table-column label="Status" width="120">
          <template #default="{ row }">
            <div class="flex items-center gap-2">
              <span class="w-2 h-2 rounded-full" :class="row.status === 'active' ? 'bg-emerald-500' : 'bg-red-500'"></span>
              <span class="text-sm font-medium" :class="row.status === 'active' ? 'text-slate-700' : 'text-slate-400'">
                {{ row.status === 'active' ? 'Active' : 'Disabled' }}
              </span>
            </div>
          </template>
        </el-table-column>

        <!-- Actions -->
        <el-table-column label="Actions" align="right" width="200">
          <template #default="{ row }">
             <el-button link type="primary" size="small" @click="handleEdit(row)">Edit</el-button>
             <el-divider direction="vertical" />
             <el-button 
              link 
              :type="row.status === 'active' ? 'danger' : 'success'" 
              size="small"
              @click="toggleStatus(row)"
            >
              {{ row.status === 'active' ? 'Disable' : 'Enable' }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <!-- User Dialog -->
    <el-dialog
      v-model="dialogVisible"
      :title="isEditMode ? 'Edit User' : 'Create New User'"
      width="500px"
      class="!rounded-xl"
      destroy-on-close
    >
      <el-form ref="formRef" :model="formData" :rules="rules" label-position="top">
        <el-form-item label="Full Name" prop="name">
          <el-input v-model="formData.name" placeholder="e.g. John Doe" />
        </el-form-item>
        
        <el-form-item label="Username" prop="username">
          <el-input v-model="formData.username" placeholder="e.g. john.doe" :disabled="isEditMode" />
        </el-form-item>

        <el-form-item label="Password" prop="password">
          <el-input v-model="formData.password" placeholder="Leave blank to keep unchanged" :show-password="true" />
          <div v-if="!isEditMode" class="text-xs text-slate-400 mt-1">Default: 123456</div>
        </el-form-item>

        <el-form-item label="Role" prop="role_key">
          <el-select v-model="formData.role_key" placeholder="Select Role" class="w-full">
            <el-option label="Administrator" value="admin">
                <span class="flex items-center gap-2">
                    <span class="w-2 h-2 rounded-full bg-purple-500"></span>
                    Administrator
                </span>
            </el-option>
            <el-option label="Finance Auditor" value="auditor">
                 <span class="flex items-center gap-2">
                    <span class="w-2 h-2 rounded-full bg-blue-500"></span>
                    Finance Auditor
                </span>
            </el-option>
            <el-option label="Regular Employee" value="employee">
                 <span class="flex items-center gap-2">
                    <span class="w-2 h-2 rounded-full bg-slate-400"></span>
                    Regular Employee
                </span>
            </el-option>
          </el-select>
        </el-form-item>
      </el-form>
      
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">Cancel</el-button>
          <el-button type="primary" class="!bg-indigo-600 !border-indigo-600" :loading="submitting" @click="submitForm">
            {{ isEditMode ? 'Save Changes' : 'Create User' }}
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
import type { User, CreateUserParams } from '../../../api/user'

// -- Table Data --
const loading = ref(false)
const tableData = ref<User[]>([])

// Mock Data
const mockData: User[] = [
    { id: 1, name: 'Administrator', username: 'admin', role_key: 'admin', status: 'active', avatar: 'https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png' },
    { id: 2, name: 'Alice Auditor', username: 'alice', role_key: 'auditor', status: 'active' },
    { id: 3, name: 'Bob Employee', username: 'bob', role_key: 'employee', status: 'disabled' },
]

const fetchData = async () => {
  loading.value = true
  try {
    await new Promise(r => setTimeout(r, 600))
    tableData.value = mockData
  } catch (err) {
    ElMessage.error('Failed to load team members')
  } finally {
    loading.value = false
  }
}

// -- Styling --
const getRoleTagType = (role: string) => {
    switch (role) {
        case 'admin': return 'danger'
        case 'auditor': return 'primary'
        default: return 'info'
    }
}

const getRoleTagClass = (role: string) => {
     switch (role) {
        case 'admin': return '!bg-purple-100 !text-purple-700'
        case 'auditor': return '!bg-blue-100 !text-blue-700'
        default: return '!bg-slate-100 !text-slate-600'
    }
}

// -- Actions --
const toggleStatus = async (row: User) => {
    const action = row.status === 'active' ? 'disable' : 'enable'
    try {
        await ElMessageBox.confirm(
            `Are you sure you want to ${action} ${row.name}?`,
            'Warning',
            { confirmButtonText: 'Yes', cancelButtonText: 'Cancel', type: 'warning' }
        )
        row.status = action === 'disable' ? 'disabled' : 'active'
        ElMessage.success(`User ${action}d successfully`)
    } catch (e) { }
}

// -- Dialog Logic --
const dialogVisible = ref(false)
const submitting = ref(false)
const isEditMode = ref(false)
const formRef = ref<FormInstance>()
const editingId = ref<number | null>(null)

const formData = reactive<CreateUserParams>({
    name: '',
    username: '',
    password: '',
    role_key: 'employee'
})

const rules = reactive<FormRules>({
    name: [{ required: true, message: 'Name is required', trigger: 'blur' }],
    username: [{ required: true, message: 'Username is required', trigger: 'blur' }],
    role_key: [{ required: true, message: 'Role is required', trigger: 'change' }]
})

const handleInvite = () => {
    isEditMode.value = false
    editingId.value = null
    formData.name = ''
    formData.username = ''
    formData.password = '123456'
    formData.role_key = 'employee'
    dialogVisible.value = true
}

const handleEdit = (row: User) => {
    isEditMode.value = true
    editingId.value = row.id
    formData.name = row.name
    formData.username = row.username
    formData.password = '' // Don't show existing password
    formData.role_key = row.role_key
    dialogVisible.value = true
}

const submitForm = async () => {
    if (!formRef.value) return
    
    await formRef.value.validate(async (valid) => {
        if (valid) {
            submitting.value = true
            try {
                await new Promise(r => setTimeout(r, 500)) // Mock API
                
                if (isEditMode.value && editingId.value) {
                    // Update
                    const index = tableData.value.findIndex(u => u.id === editingId.value)
                    if (index !== -1) {
                        tableData.value[index] = { 
                            ...tableData.value[index], 
                            name: formData.name, 
                            role_key: formData.role_key as any 
                        }
                    }
                    ElMessage.success('User updated successfully')
                } else {
                    // Create
                    tableData.value.push({
                        id: Date.now(),
                        name: formData.name,
                        username: formData.username,
                        role_key: formData.role_key as any,
                        status: 'active'
                    })
                    ElMessage.success('User created successfully')
                }
                dialogVisible.value = false
            } catch (e) {
                ElMessage.error('Operation failed')
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