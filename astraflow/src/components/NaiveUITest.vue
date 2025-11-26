<template>
  <div class="naive-test-container">
    <h2>Naive UI 测试组件</h2>

    <!-- 测试按钮 -->
    <n-space>
      <n-button @click="showMessage" type="primary">测试消息</n-button>
      <n-button @click="showDialog" type="info">测试对话框</n-button>
      <n-button @click="showNotification" type="success">测试通知</n-button>
    </n-space>

    <!-- 测试表单 -->
    <n-form
      :model="formData"
      :rules="rules"
      ref="formRef"
      class="test-form"
    >
      <n-form-item label="用户名" path="username">
        <n-input
          v-model:value="formData.username"
          placeholder="请输入用户名"
        />
      </n-form-item>

      <n-form-item label="邮箱" path="email">
        <n-input
          v-model:value="formData.email"
          placeholder="请输入邮箱"
        />
      </n-form-item>

      <n-form-item label="选择分类" path="category">
        <n-select
          v-model:value="formData.category"
          :options="categoryOptions"
          placeholder="请选择分类"
        />
      </n-form-item>

      <n-form-item label="日期选择">
        <n-date-picker
          v-model:value="formData.date"
          type="datetime"
          clearable
        />
      </n-form-item>

      <n-form-item label="金额">
        <n-input-number
          v-model:value="formData.amount"
          :min="0"
          :max="10000"
          :precision="2"
        />
      </n-form-item>

      <n-form-item label="状态">
        <n-radio-group v-model:value="formData.status" name="status">
          <n-space>
            <n-radio value="pending">待处理</n-radio>
            <n-radio value="confirmed">已确认</n-radio>
            <n-radio value="rejected">已拒绝</n-radio>
          </n-space>
        </n-radio-group>
      </n-form-item>

      <n-button
        @click="submitForm"
        type="primary"
        :loading="submitting"
      >
        提交
      </n-button>
    </n-form>

    <!-- 测试模态框 -->
    <n-modal v-model:show="showModal" :mask-closable="false">
      <n-card
        style="width: 600px;"
        title="测试模态框"
        :bordered="false"
        size="huge"
        role="dialog"
        aria-modal="true"
      >
        <p>这是一个使用 Naive UI 创建的模态框</p>
        <n-button @click="closeModal" type="primary">关闭</n-button>
      </n-card>
    </n-modal>
  </div>
</template>

<script setup>
import { ref, inject } from 'vue'
import { useMessage, useDialog, useNotification } from 'naive-ui'

const message = useMessage()
const dialog = useDialog()
const notification = useNotification()

const formData = ref({
  username: '',
  email: '',
  category: null,
  date: null,
  amount: null,
  status: 'pending'
})

const rules = {
  username: { required: true, message: '请输入用户名', trigger: 'blur' },
  email: {
    required: true,
    message: '请输入邮箱',
    trigger: 'blur',
    pattern: /^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/
  }
}

const categoryOptions = [
  { label: '餐饮', value: '餐饮' },
  { label: '交通', value: '交通' },
  { label: '住宿', value: '住宿' },
  { label: '办公', value: '办公' },
  { label: '其他', value: '其他' }
]

const showModal = ref(false)
const submitting = ref(false)

const formRef = ref()

const showMessage = () => {
  message.success('这是来自 Naive UI 的成功消息！')
}

const showDialog = () => {
  dialog.warning({
    title: '测试对话框',
    content: '这是一个使用 Naive UI 的对话框',
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: () => {
      message.info('点击了确定')
    },
    onNegativeClick: () => {
      message.info('点击了取消')
    }
  })
}

const showNotification = () => {
  notification.success({
    title: '测试通知',
    content: '这是一个使用 Naive UI 的通知消息',
    duration: 3000
  })
}

const submitForm = async () => {
  const { validate } = formRef.value
  try {
    await validate()
    submitting.value = true

    // 模拟提交
    setTimeout(() => {
      message.success('表单提交成功！')
      submitting.value = false
    }, 1000)
  } catch (error) {
    message.error('表单验证失败')
  }
}

const openModal = () => {
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
}
</script>

<style scoped>
.naive-test-container {
  padding: 20px;
  max-width: 800px;
  margin: 0 auto;
}

.test-form {
  margin-top: 20px;
  padding: 20px;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
}
</style>