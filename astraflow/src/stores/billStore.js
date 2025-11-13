import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useBillStore = defineStore('bill', () => {
  // State
  const bills = ref([])
  const loading = ref(false)
  const error = ref(null)

  // Mock bill data
  const mockBills = [
    {
      id: 1,
      title: '办公用品采购',
      amount: 299.50,
      date: '2025-01-15',
      category: '办公用品',
      status: '已批准',
      vendor: '文具供应商',
      invoiceNumber: 'INV-001',
      description: '采购办公用品',
      receiptUrl: '/receipts/receipt1.jpg'
    },
    {
      id: 2,
      title: '差旅费',
      amount: 1250.00,
      date: '2025-01-12',
      category: '差旅',
      status: '待审批',
      vendor: '航空公司',
      invoiceNumber: 'INV-002',
      description: '出差机票',
      receiptUrl: '/receipts/receipt2.jpg'
    },
    {
      id: 3,
      title: '餐饮费',
      amount: 185.30,
      date: '2025-01-10',
      category: '餐饮',
      status: '已批准',
      vendor: '餐厅',
      invoiceNumber: 'INV-003',
      description: '商务餐费',
      receiptUrl: '/receipts/receipt3.jpg'
    },
    {
      id: 4,
      title: '培训费',
      amount: 2500.00,
      date: '2025-01-08',
      category: '培训',
      status: '已拒绝',
      vendor: '培训机构',
      invoiceNumber: 'INV-004',
      description: '技能培训课程',
      receiptUrl: '/receipts/receipt4.jpg'
    }
  ]

  // Initialize with mock data
  bills.value = [...mockBills]

  // Getters
  const getAllBills = computed(() => bills.value)
  const getBillById = computed(() => (id) => bills.value.find(bill => bill.id === id))
  const getBillsByStatus = computed(() => (status) => bills.value.filter(bill => bill.status === status))
  const getBillsByCategory = computed(() => (category) => bills.value.filter(bill => bill.category === category))
  const getTotalAmount = computed(() => bills.value.reduce((sum, bill) => sum + bill.amount, 0))
  const getBillsCount = computed(() => bills.value.length)

  // Actions
  const fetchBills = async () => {
    loading.value = true
    error.value = null
    try {
      // Simulate API call
      await new Promise(resolve => setTimeout(resolve, 1000))
      // In real app, this would fetch from API
      bills.value = [...mockBills]
    } catch (err) {
      error.value = err.message
    } finally {
      loading.value = false
    }
  }

  const addBill = async (billData) => {
    loading.value = true
    error.value = null
    try {
      // Simulate API call
      await new Promise(resolve => setTimeout(resolve, 500))
      const newBill = {
        id: bills.value.length + 1,
        ...billData,
        date: new Date().toISOString().split('T')[0]
      }
      bills.value.push(newBill)
    } catch (err) {
      error.value = err.message
    } finally {
      loading.value = false
    }
  }

  const updateBill = async (id, billData) => {
    loading.value = true
    error.value = null
    try {
      // Simulate API call
      await new Promise(resolve => setTimeout(resolve, 500))
      const index = bills.value.findIndex(bill => bill.id === id)
      if (index !== -1) {
        bills.value[index] = { ...bills.value[index], ...billData }
      }
    } catch (err) {
      error.value = err.message
    } finally {
      loading.value = false
    }
  }

  const deleteBill = async (id) => {
    loading.value = true
    error.value = null
    try {
      // Simulate API call
      await new Promise(resolve => setTimeout(resolve, 500))
      bills.value = bills.value.filter(bill => bill.id !== id)
    } catch (err) {
      error.value = err.message
    } finally {
      loading.value = false
    }
  }

  const approveBill = async (id) => {
    loading.value = true
    error.value = null
    try {
      // Simulate API call
      await new Promise(resolve => setTimeout(resolve, 300))
      const index = bills.value.findIndex(bill => bill.id === id)
      if (index !== -1) {
        bills.value[index].status = '已批准'
      }
    } catch (err) {
      error.value = err.message
    } finally {
      loading.value = false
    }
  }

  const rejectBill = async (id) => {
    loading.value = true
    error.value = null
    try {
      // Simulate API call
      await new Promise(resolve => setTimeout(resolve, 300))
      const index = bills.value.findIndex(bill => bill.id === id)
      if (index !== -1) {
        bills.value[index].status = '已拒绝'
      }
    } catch (err) {
      error.value = err.message
    } finally {
      loading.value = false
    }
  }

  return {
    // State
    bills,
    loading,
    error,

    // Getters
    getAllBills,
    getBillById,
    getBillsByStatus,
    getBillsByCategory,
    getTotalAmount,
    getBillsCount,

    // Actions
    fetchBills,
    addBill,
    updateBill,
    deleteBill,
    approveBill,
    rejectBill
  }
})