<template>
  <div class="user-management">
    <!-- 顶部导航栏 -->
    <PageHeader title="用户管理" />

    <!-- 页面内容 -->
    <div class="page-content">
      <div class="page-subheader">
        <p>管理企业租户下的成员</p>
      </div>

      <!-- 添加用户按钮 -->
      <div class="header-actions">
        <button @click="showAddUserModal = true" class="btn btn-primary">
          添加用户
        </button>
      </div>

      <!-- 搜索和过滤 -->
      <div class="search-container">
        <div class="search-group">
          <input v-model="searchQuery" type="text" placeholder="搜索用户..." class="search-input" />
          <select v-model="filterRole" class="filter-select">
            <option value="">所有角色</option>
            <option value="admin">管理员</option>
            <option value="normal">普通用户</option>
          </select>
          <select v-model="filterStatus" class="filter-select">
            <option value="">所有状态</option>
            <option value="true">启用</option>
            <option value="false">禁用</option>
          </select>
        </div>
      </div>

      <!-- 用户列表 -->
      <div class="user-list-container">
        <div class="table-container">
          <table class="user-table">
            <thead>
              <tr>
                <th>账号</th>
                <th>邮箱</th>
                <th>手机号</th>
                <th>角色</th>
                <th>操作</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="user in filteredUsers" :key="user.id">
                <td>{{ user.username || user.name }}</td>
                <td>{{ user.email || '-' }}</td>
                <td>{{ user.phone || '-' }}</td>
                <td>
                  <span class="role-badge" :class="getRoleClass(user.role)">
                    {{ getRoleDisplayName(user.role) }}
                  </span>
                </td>
                <td class="actions">
                  <button @click="editUser(user)" class="btn btn-sm btn-outline">编辑</button>
                  <button @click="deleteUser(user)" class="btn btn-sm btn-danger"
                    :disabled="user.id === currentUser.id">
                    删除
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- 空状态 -->
        <div v-if="filteredUsers.length === 0 && users.length > 0" class="empty-state">
          <p>没有找到匹配的用户</p>
        </div>
        <div v-else-if="users.length === 0" class="empty-state">
          <p>暂无用户数据</p>
        </div>
      </div>

      <!-- 添加/编辑用户模态框 -->
      <div v-if="showAddUserModal" class="modal-overlay" @click="closeModal">
        <div class="modal-content" @click.stop>
          <div class="modal-header">
            <h3>{{ editingUser ? '编辑用户' : '添加用户' }}</h3>
            <button @click="closeModal" class="close-btn">&times;</button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="saveUser">
              <div class="form-group">
                <label for="userName">用户名</label>
                <input type="text" id="userName" v-model="form.username" :disabled="editingUser" required
                  class="form-input" />
              </div>
              <div class="form-group">
                <label for="userEmail">邮箱</label>
                <input type="email" id="userEmail" v-model="form.email" required
                  class="form-input" />
              </div>
              <div class="form-group">
                <label for="userPhone">手机号</label>
                <input type="text" id="userPhone" v-model="form.phone"
                  class="form-input" />
              </div>
              <div class="form-group">
                <label for="userRole">角色</label>
                <select id="userRole" v-model="form.role" :disabled="editingUser" required class="form-select">
                  <option value="normal">普通用户</option>
                  <option value="admin">管理员</option>
                </select>
              </div>
            </form>
          </div>
          <div class="modal-footer">
            <button @click="closeModal" class="btn btn-secondary">取消</button>
            <button @click="saveUser" class="btn btn-primary" :disabled="formLoading">
              {{ formLoading ? '保存中...' : (editingUser ? '更新' : '添加') }}
            </button>
          </div>
        </div>
      </div>

      <!-- 加载指示器 -->
      <div v-if="loading" class="loading-overlay">
        <div class="loading-spinner"></div>
        <p>加载中...</p>
      </div>

      <!-- 确认对话框 -->
      <div v-if="showConfirmDialog" class="modal-overlay" @click="closeConfirmDialog">
        <div class="modal-content" @click.stop>
          <div class="modal-body">
            <p>{{ confirmMessage }}</p>
          </div>
          <div class="modal-footer">
            <button @click="closeConfirmDialog" class="btn btn-secondary">取消</button>
            <button @click="confirmAction" class="btn btn-danger">确定</button>
          </div>
        </div>
      </div>

      <!-- 页脚 -->
      <PageFooter />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '../stores/userStore'
import { getTenantUsers, createTenantUser, updateTenantUser, deleteTenantUser } from '../services/api/userManagementApi'
import PageHeader from '../components/ui/PageHeader.vue'
import PageFooter from '../components/ui/PageFooter.vue'

// State
const users = ref([])
const loading = ref(false)
const formLoading = ref(false)
const showAddUserModal = ref(false)
const showConfirmDialog = ref(false)
const editingUser = ref(null)
const confirmMessage = ref('')
const confirmCallback = ref(null)

// Search and filter state
const searchQuery = ref('')
const filterRole = ref('')
const filterStatus = ref('')

// Form data
const form = ref({
  username: '',
  email: '',
  phone: '',
  role: 'normal'
})

const userStore = useUserStore()

// Current logged in user
const currentUser = computed(() => userStore.user)

// Filtered users based on search and filters
const filteredUsers = computed(() => {
  let filtered = users.value

  // Search filter
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(user =>
      (user.username || user.name).toLowerCase().includes(query) ||
      (user.email || '').toLowerCase().includes(query) ||
      (user.phone || '').toLowerCase().includes(query)
    )
  }

  // Role filter
  if (filterRole.value) {
    filtered = filtered.filter(user => (user.role || 'normal') === filterRole.value)
  }

  // Status filter
  if (filterStatus.value !== '') {
    const isActive = filterStatus.value === 'true'
    filtered = filtered.filter(user => (user.deleted_at ? false : true) === isActive)
  }

  return filtered
})

// Check if current user is tenant admin or system admin
// Note: This check is also handled by the router guard
if (!userStore.loading && !userStore.isTenantAdmin() && !userStore.isSystemAdmin()) {
  // Redirect or show error - this should be handled by router guard
}

// Get role display name
const getRoleDisplayName = (role) => {
  const roleMap = {
    'admin': '管理员',
    'normal': '普通用户'
  }
  return roleMap[role] || role
}

// Get role CSS class
const getRoleClass = (role) => {
  const roleClassMap = {
    'admin': 'role-admin',
    'normal': 'role-normal'
  }
  return roleClassMap[role] || 'role-normal'
}

// Fetch users
const fetchUsers = async () => {
  loading.value = true
  try {
    const response = await getTenantUsers()
    users.value = response.data || []
  } catch (error) {
    console.error('Failed to fetch users:', error)
    // Handle error appropriately
    // For development, we'll use mock data
    users.value = [
      {
        id: 1,
        name: '张三',
        email: 'zhangsan@example.com',
        role: 'admin',
        is_active: true
      },
      {
        id: 2,
        name: '李四',
        email: 'lisi@example.com',
        role: 'normal',
        is_active: true
      },
      {
        id: 3,
        name: '王五',
        email: 'wangwu@example.com',
        role: 'normal',
        is_active: false
      }
    ]
  } finally {
    loading.value = false
  }
}

// Open edit user modal
const editUser = (user) => {
  editingUser.value = user
  form.value = {
    username: user.username || user.name,
    email: user.email,
    phone: user.phone || '',
    role: user.role
  }
  showAddUserModal.value = true
}

// Delete user
const deleteUser = (user) => {
  if (user.id === currentUser.value.id) {
    alert('不能删除自己')
    return
  }

  confirmMessage.value = `确定要删除用户 "${user.name}" 吗？此操作不可撤销。`
  confirmCallback.value = async () => {
    try {
      // Call API to delete user
      await deleteTenantUser(user.id)
      // Remove from local list
      users.value = users.value.filter(u => u.id !== user.id)
    } catch (error) {
      console.error('Failed to delete user:', error)
      // Handle error
      alert('删除用户失败: ' + (error.message || '未知错误'))
    }
    closeConfirmDialog()
  }
  showConfirmDialog.value = true
}

// Save user (add or update)
const saveUser = async () => {
  formLoading.value = true
  try {
    if (editingUser.value) {
      // Update existing user
      // Only update email and phone when editing user (as per requirements)
      const userData = {
        email: form.value.email,
        phone: form.value.phone // Add phone field as expected by backend
      };
      const response = await updateTenantUser(editingUser.value.id, userData)
      const index = users.value.findIndex(u => u.id === editingUser.value.id)
      if (index !== -1) {
        // Map backend response to frontend format
        const updatedUser = {
          ...editingUser.value,
          username: response.data.username || response.data.name || form.value.username,
          email: response.data.email || form.value.email,
          phone: response.data.phone || form.value.phone,
          role: response.data.role || form.value.role,
          ...response.data
        };
        users.value[index] = updatedUser;
      }
    } else {
      // Add new user
      // Map frontend fields to backend expectations
      const userData = {
        username: form.value.username,
        password: 'default123', // Default password for new users, should be changed by user later
        email: form.value.email,
        role: form.value.role === 'admin', // Map string to boolean: 'admin' becomes true, 'normal' becomes false
        phone: form.value.phone // Add phone field as expected by backend
      };
      const response = await createTenantUser(userData)
      // Map backend response to frontend format
      const newUser = {
        id: response.data.id,
        username: response.data.username || userData.username,
        email: response.data.email || userData.email,
        phone: response.data.phone || userData.phone,
        role: response.data.role || userData.role,
        created_at: response.data.created_at
      };
      users.value.push(newUser)
    }

    closeModal()
    // Refresh the user list to ensure it's up to date
    await fetchUsers()
  } catch (error) {
    console.error('Failed to save user:', error)
    // Handle error
    alert('保存用户失败: ' + (error.message || '未知错误'))
  } finally {
    formLoading.value = false
  }
}

// Close modal
const closeModal = () => {
  showAddUserModal.value = false
  editingUser.value = null
  form.value = {
    username: '',
    email: '',
    phone: '',
    role: 'normal'
  }
}

// Close confirm dialog
const closeConfirmDialog = () => {
  showConfirmDialog.value = false
  confirmMessage.value = ''
  confirmCallback.value = null
}

// Confirm action
const confirmAction = () => {
  if (confirmCallback.value) {
    confirmCallback.value()
  }
}


// Format date
const formatDate = (dateString) => {
  if (!dateString) return '-'
  try {
    const date = new Date(dateString)
    return date.toLocaleDateString('zh-CN', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit'
    })
  } catch (error) {
    return '-'
  }
}

// Initialize
onMounted(() => {
  fetchUsers()
})
</script>

<style scoped>
.user-management {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  color: #ffffff;
}

.page-content {
  width: 80rem;
  margin: 0 auto;
  flex: 1;
  padding: 0 3rem;
}

.page-subheader {
  padding: 20px 0;
  margin-bottom: 24px;
}

.page-subheader p {
  margin: 0;
  color: #9ca3af;
  font-size: 16px;
}

.header-actions {
  margin-bottom: 20px;
  display: flex;
  justify-content: flex-end;
}

/* 搜索和过滤 */
.search-container {
  margin-bottom: 20px;
}

.search-group {
  display: flex;
  gap: 12px;
  flex-wrap: wrap;
}

.search-input,
.filter-select {
  padding: 8px 12px;
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 8px;
  background: rgba(255, 255, 255, 0.1);
  color: #ffffff;
  font-size: 14px;
  min-width: 200px;
  backdrop-filter: blur(10px);
}

.search-input::placeholder {
  color: #9ca3af;
}

.search-input:focus,
.filter-select:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.25);
}

.filter-select {
  cursor: pointer;
}

.filter-select option {
  background: #1f2937;
  color: #ffffff;
}

/* 按钮样式 */
.btn {
  padding: 8px 16px;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
  transition: all 0.3s ease;
  backdrop-filter: blur(10px);
}

.btn-primary {
  background: linear-gradient(to right, #3b82f6, #8b5cf6);
  color: white;
  border: 1px solid rgba(59, 130, 246, 0.3);
}

.btn-primary:hover:not(:disabled) {
  background: linear-gradient(to right, #2563eb, #7c3aed);
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(59, 130, 246, 0.3);
}

.btn-secondary {
  background: rgba(107, 114, 128, 0.2);
  color: #e5e7eb;
  border: 1px solid rgba(107, 114, 128, 0.3);
}

.btn-secondary:hover {
  background: rgba(107, 114, 128, 0.3);
  border-color: rgba(107, 114, 128, 0.5);
}

.btn-outline {
  background: transparent;
  border: 1px solid rgba(59, 130, 246, 0.5);
  color: #60a5fa;
}

.btn-outline:hover {
  background: rgba(59, 130, 246, 0.2);
  border-color: #60a5fa;
  color: #93c5fd;
}

.btn-danger {
  background: rgba(239, 68, 68, 0.2);
  color: #f87171;
  border: 1px solid rgba(239, 68, 68, 0.3);
}

.btn-danger:hover:not(:disabled) {
  background: rgba(239, 68, 68, 0.3);
  border-color: rgba(239, 68, 68, 0.5);
  color: #fca5a5;
}

.btn-warning {
  background: rgba(245, 158, 11, 0.2);
  color: #fbbf24;
  border: 1px solid rgba(245, 158, 11, 0.3);
}

.btn-warning:hover {
  background: rgba(245, 158, 11, 0.3);
  border-color: rgba(245, 158, 11, 0.5);
  color: #fcd34d;
}

.btn-success {
  background: rgba(34, 197, 94, 0.2);
  color: #4ade80;
  border: 1px solid rgba(34, 197, 94, 0.3);
}

.btn-success:hover {
  background: rgba(34, 197, 94, 0.3);
  border-color: rgba(34, 197, 94, 0.5);
  color: #86efac;
}

.btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.btn-sm {
  padding: 6px 12px;
  font-size: 12px;
}

/* 用户列表容器 */
.user-list-container {
  margin-bottom: 20px;
  background: rgba(255, 255, 255, 0.05);
  border-radius: 16px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  overflow: hidden;
  backdrop-filter: blur(20px);
}

.table-container {
  overflow-x: auto;
}

.user-table {
  width: 100%;
  border-collapse: collapse;
}

.user-table th,
.user-table td {
  padding: 16px;
  text-align: left;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.user-table th {
  background: rgba(255, 255, 255, 0.05);
  font-weight: 600;
  color: #e5e7eb;
  backdrop-filter: blur(10px);
}

.user-table tbody tr:hover {
  background: rgba(255, 255, 255, 0.05);
}

/* 角色和状态徽章 */
.role-badge,
.status-badge {
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 500;
}

.role-admin {
  background: rgba(34, 197, 94, 0.2);
  color: #4ade80;
  border: 1px solid rgba(34, 197, 94, 0.3);
}

.role-normal {
  background: rgba(59, 130, 246, 0.2);
  color: #60a5fa;
  border: 1px solid rgba(59, 130, 246, 0.3);
}

.status-active {
  background: rgba(34, 197, 94, 0.2);
  color: #4ade80;
  border: 1px solid rgba(34, 197, 94, 0.3);
}

.status-inactive {
  background: rgba(239, 68, 68, 0.2);
  color: #f87171;
  border: 1px solid rgba(239, 68, 68, 0.3);
}

.actions {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

/* 空状态 */
.empty-state {
  padding: 48px;
  text-align: center;
  color: #9ca3af;
}

/* 模态框 */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.8);
  backdrop-filter: blur(10px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  background: rgba(24, 24, 27, 0.95);
  border-radius: 16px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  width: 100%;
  max-width: 500px;
  max-height: 90vh;
  overflow-y: auto;
  backdrop-filter: blur(20px);
}

.modal-header {
  padding: 20px 24px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-header h3 {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
  color: #ffffff;
}

.close-btn {
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
  color: #9ca3af;
  padding: 0;
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  transition: all 0.2s ease;
}

.close-btn:hover {
  color: #ffffff;
  background: rgba(255, 255, 255, 0.1);
}

.modal-body {
  padding: 24px;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  margin-bottom: 8px;
  font-weight: 500;
  color: #e5e7eb;
}

.form-input,
.form-select {
  width: 100%;
  padding: 12px;
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 8px;
  font-size: 14px;
  background: rgba(255, 255, 255, 0.05);
  color: #ffffff;
  backdrop-filter: blur(10px);
}

.form-input:focus,
.form-select:focus {
  outline: none;
  border-color: #3b82f6;
  box-shadow: 0 0 0 2px rgba(59, 130, 246, 0.25);
}

.form-input:disabled {
  background: rgba(255, 255, 255, 0.1);
  cursor: not-allowed;
  opacity: 0.6;
}

.form-select {
  cursor: pointer;
}

.form-select option {
  background: #1f2937;
  color: #ffffff;
}

.checkbox-label {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #e5e7eb;
  cursor: pointer;
}

.form-checkbox {
  width: 16px;
  height: 16px;
  accent-color: #3b82f6;
}

.modal-footer {
  padding: 20px 24px;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

/* 加载指示器 */
.loading-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.8);
  backdrop-filter: blur(10px);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  z-index: 2000;
}

.loading-spinner {
  width: 40px;
  height: 40px;
  border: 4px solid rgba(255, 255, 255, 0.1);
  border-top: 4px solid #3b82f6;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(360deg);
  }
}

.loading-overlay p {
  margin-top: 16px;
  font-size: 16px;
  color: #e5e7eb;
}


/* 亮色主题适配 */
[data-theme="light"] .user-management {
  color: #1f2937;
}

[data-theme="light"] .page-header h1 {
  background: linear-gradient(to right, #3b82f6, #8b5cf6);
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
}

[data-theme="light"] .page-header p {
  color: #6b7280;
}

[data-theme="light"] .search-input,
[data-theme="light"] .filter-select,
[data-theme="light"] .form-input,
[data-theme="light"] .form-select {
  background: rgba(255, 255, 255, 0.8);
  border-color: rgba(0, 0, 0, 0.1);
  color: #1f2937;
}

[data-theme="light"] .search-input::placeholder {
  color: #9ca3af;
}

[data-theme="light"] .user-list-container {
  background: rgba(255, 255, 255, 0.8);
  border-color: rgba(0, 0, 0, 0.1);
}

[data-theme="light"] .user-table th,
[data-theme="light"] .user-table td {
  border-color: rgba(0, 0, 0, 0.1);
}

[data-theme="light"] .user-table th {
  background: rgba(255, 255, 255, 0.8);
  color: #374151;
}

[data-theme="light"] .user-table tbody tr:hover {
  background: rgba(255, 255, 255, 0.8);
}

[data-theme="light"] .btn-secondary {
  background: rgba(107, 114, 128, 0.1);
  color: #374151;
  border-color: rgba(107, 114, 128, 0.2);
}

[data-theme="light"] .btn-secondary:hover {
  background: rgba(107, 114, 128, 0.2);
}

[data-theme="light"] .modal-content {
  background: rgba(255, 255, 255, 0.95);
  border-color: rgba(0, 0, 0, 0.1);
}

[data-theme="light"] .modal-header {
  border-color: rgba(0, 0, 0, 0.1);
}

[data-theme="light"] .modal-header h3 {
  color: #1f2937;
}

[data-theme="light"] .close-btn {
  color: #6b7280;
}

[data-theme="light"] .close-btn:hover {
  color: #1f2937;
  background: rgba(0, 0, 0, 0.05);
}

[data-theme="light"] .form-group label {
  color: #374151;
}

[data-theme="light"] .checkbox-label {
  color: #374151;
}

[data-theme="light"] .modal-footer {
  border-color: rgba(0, 0, 0, 0.1);
}

[data-theme="light"] .empty-state {
  color: #6b7280;
}

[data-theme="light"] .loading-overlay {
  background-color: rgba(255, 255, 255, 0.9);
}

[data-theme="light"] .loading-overlay p {
  color: #374151;
}

[data-theme="light"] .loading-spinner {
  border-color: rgba(0, 0, 0, 0.1);
  border-top-color: #3b82f6;
}



/* 响应式设计 */
@media (max-width: 768px) {
  .page-content {
    padding: 0 16px;
  }

  .search-group {
    flex-direction: column;
  }

  .search-input,
  .filter-select {
    min-width: 100%;
  }

  .table-container {
    overflow-x: auto;
  }

  .user-table {
    min-width: 600px;
  }

  .actions {
    flex-direction: column;
    gap: 4px;
  }

  .btn-sm {
    font-size: 12px;
    padding: 4px 8px;
  }
}
</style>