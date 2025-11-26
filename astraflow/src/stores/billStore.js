import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { getInvoices, getInvoicesByUser, getInvoicesByTenant, getInvoiceById, createInvoice, updateInvoice, deleteInvoice } from '@/services/api/invoiceApi'
import { useUserStore } from '@/stores/userStore'

export const useBillStore = defineStore('bill', () => {
  // State
  const bills = ref([])
  const loading = ref(false)
  const error = ref(null)
  const pagination = ref({
    page: 1,
    size: 10,
    total: 0
  })

  // Getters
  const getAllBills = computed(() => bills.value)
  const getBillById = computed(() => (id) => bills.value.find(bill => bill.id === id))
  const getBillsByStatus = computed(() => (status) => bills.value.filter(bill => bill.status === status))
  const getBillsByCategory = computed(() => (category) => bills.value.filter(bill => bill.category === category))
  const getTotalAmount = computed(() => bills.value.reduce((sum, bill) => sum + (bill.amount || 0), 0))
  const getBillsCount = computed(() => bills.value.length)

  // Actions
  const fetchBills = async (params = {}) => {
    loading.value = true
    error.value = null
    try {
      const userStore = useUserStore()
      let response

      // Determine which endpoint to use based on user type
      if (userStore.isPersonalUser()) {
        // Personal user - get invoices by user
        response = await getInvoicesByUser({
          page: params.page || pagination.value.page,
          page_size: params.page_size || pagination.value.size,
          ...params
        })
      } else if (userStore.user.tenantId) {
        // Tenant user - get invoices by tenant
        response = await getInvoicesByTenant({
          page: params.page || pagination.value.page,
          page_size: params.page_size || pagination.value.size,
          ...params
        })
      } else {
        // Default - get all invoices (for admin users)
        response = await getInvoices({
          page: params.page || pagination.value.page,
          page_size: params.page_size || pagination.value.size,
          ...params
        })
      }

      if (response?.data) {
        // Map invoice fields to bill fields for compatibility
        bills.value = response.data.tenants?.map(invoice => ({
          id: invoice.id,
          title: invoice.invoice_number || invoice.title || '未命名发票',
          amount: invoice.amount,
          date: invoice.invoice_date || invoice.date,
          category: invoice.category,
          status: invoice.status || invoice.status,
          vendor: invoice.vendor,
          invoiceNumber: invoice.invoice_number,
          taxId: invoice.taxId,
          paymentSource: invoice.payment_source,
          description: invoice.description || invoice.description || '',
          receiptUrl: invoice.receipt_url || invoice.receiptUrl || ''
        })) || []

        // Update pagination info
        if (response.data.page !== undefined) {
          pagination.value.page = response.data.page
        }
        if (response.data.size !== undefined) {
          pagination.value.size = response.data.size
        }
        if (response.data.total !== undefined) {
          pagination.value.total = response.data.total
        }
      }
    } catch (err) {
      error.value = err.message || '获取发票列表失败'
      console.error('Error fetching bills:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  const fetchBillById = async (id) => {
    loading.value = true
    error.value = null
    try {
      const response = await getInvoiceById(id)

      if (response?.data?.invoice) {
        const invoice = response.data.invoice
        // Find and update the specific invoice in the store
        const index = bills.value.findIndex(bill => bill.id === id)
        if (index !== -1) {
          bills.value[index] = {
            id: invoice.id,
            title: invoice.invoice_number || invoice.title || '未命名发票',
            amount: invoice.amount,
            date: invoice.invoice_date || invoice.date,
            category: invoice.category,
            status: invoice.status || invoice.status,
            vendor: invoice.vendor,
            invoiceNumber: invoice.invoice_number,
            taxId: invoice.taxId,
            paymentSource: invoice.payment_source,
            description: invoice.description || invoice.description || '',
            receiptUrl: invoice.receipt_url || invoice.receiptUrl || ''
          }
        } else {
          // If not in the list, add it
          bills.value.push({
            id: invoice.id,
            title: invoice.invoice_number || invoice.title || '未命名发票',
            amount: invoice.amount,
            date: invoice.invoice_date || invoice.date,
            category: invoice.category,
            status: invoice.status || invoice.status,
            vendor: invoice.vendor,
            invoiceNumber: invoice.invoice_number,
            taxId: invoice.taxId,
            paymentSource: invoice.payment_source,
            description: invoice.description || invoice.description || '',
            receiptUrl: invoice.receipt_url || invoice.receiptUrl || ''
          })
        }
        return bills.value[index >= 0 ? index : bills.value.length - 1]
      }
    } catch (err) {
      error.value = err.message || '获取发票详情失败'
      console.error('Error fetching bill by id:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  const addBill = async (billData) => {
    loading.value = true
    error.value = null
    try {
      const response = await createInvoice({
        invoice_number: billData.invoiceNumber || billData.title,
        invoice_date: billData.date ? new Date(billData.date).toISOString() : null,
        amount: billData.amount,
        vendor: billData.vendor,
        taxId: billData.taxId,
        category: billData.category,
        payment_source: billData.paymentSource,
        status: billData.status || 'pending',
        description: billData.description
      })

      if (response?.data?.invoice) {
        const invoice = response.data.invoice
        const newBill = {
          id: invoice.id,
          title: invoice.invoice_number || invoice.title || '未命名发票',
          amount: invoice.amount,
          date: invoice.invoice_date || invoice.date,
          category: invoice.category,
          status: invoice.status || 'pending',
          vendor: invoice.vendor,
          invoiceNumber: invoice.invoice_number,
          taxId: invoice.taxId,
          paymentSource: invoice.payment_source,
          description: invoice.description || '',
          receiptUrl: invoice.receipt_url || invoice.receiptUrl || ''
        }
        bills.value.unshift(newBill) // Add to the beginning of the list
        return newBill
      }
    } catch (err) {
      error.value = err.message || '创建发票失败'
      console.error('Error adding bill:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  const updateBill = async (id, billData) => {
    loading.value = true
    error.value = null
    try {
      const response = await updateInvoice(id, {
        invoice_number: billData.invoiceNumber || billData.title,
        invoice_date: billData.date ? new Date(billData.date).toISOString() : null,
        amount: billData.amount,
        vendor: billData.vendor,
        taxId: billData.taxId,
        category: billData.category,
        payment_source: billData.paymentSource,
        status: billData.status,
        description: billData.description
      })

      if (response?.data?.invoice) {
        const invoice = response.data.invoice
        const index = bills.value.findIndex(bill => bill.id === id)
        if (index !== -1) {
          bills.value[index] = {
            id: invoice.id,
            title: invoice.invoice_number || invoice.title || '未命名发票',
            amount: invoice.amount,
            date: invoice.invoice_date || invoice.date,
            category: invoice.category,
            status: invoice.status || invoice.status,
            vendor: invoice.vendor,
            invoiceNumber: invoice.invoice_number,
            taxId: invoice.taxId,
            paymentSource: invoice.payment_source,
            description: invoice.description || invoice.description || '',
            receiptUrl: invoice.receipt_url || invoice.receiptUrl || ''
          }
        }
        return bills.value[index]
      }
    } catch (err) {
      error.value = err.message || '更新发票失败'
      console.error('Error updating bill:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  const deleteBill = async (id) => {
    loading.value = true
    error.value = null
    try {
      await deleteInvoice(id)
      bills.value = bills.value.filter(bill => bill.id !== id)
      return true
    } catch (err) {
      error.value = err.message || '删除发票失败'
      console.error('Error deleting bill:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  const approveBill = async (id) => {
    loading.value = true
    error.value = null
    try {
      const response = await updateInvoice(id, { status: 'confirmed' })

      if (response?.data?.invoice) {
        const invoice = response.data.invoice
        const index = bills.value.findIndex(bill => bill.id === id)
        if (index !== -1) {
          bills.value[index] = {
            ...bills.value[index],
            status: invoice.status || 'confirmed'
          }
        }
        return bills.value[index]
      }
    } catch (err) {
      error.value = err.message || '审批发票失败'
      console.error('Error approving bill:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  const rejectBill = async (id) => {
    loading.value = true
    error.value = null
    try {
      const response = await updateInvoice(id, { status: 'rejected' })

      if (response?.data?.invoice) {
        const invoice = response.data.invoice
        const index = bills.value.findIndex(bill => bill.id === id)
        if (index !== -1) {
          bills.value[index] = {
            ...bills.value[index],
            status: invoice.status || 'rejected'
          }
        }
        return bills.value[index]
      }
    } catch (err) {
      error.value = err.message || '拒绝发票失败'
      console.error('Error rejecting bill:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  return {
    // State
    bills,
    loading,
    error,
    pagination,

    // Getters
    getAllBills,
    getBillById,
    getBillsByStatus,
    getBillsByCategory,
    getTotalAmount,
    getBillsCount,

    // Actions
    fetchBills,
    fetchBillById,
    addBill,
    updateBill,
    deleteBill,
    approveBill,
    rejectBill
  }
})