import request from '../request';

/**
 * 获取当前用户的发票列表
 * @param {object} params 查询参数，例如 { status: 'recognizing', page: 1, pageSize: 10 }
 */
export const getMyInvoices = (params) => {
  return request({
    url: '/invoices/my-invoices',
    method: 'get',
    params,
  });
};

/**
 * 员工确认 AI 识别结果
 * @param {number} id 发票ID
 * @param {object} data 确认后的数据
 */
export const confirmInvoice = (id, data) => {
    return request({
        url: `/invoices/${id}/confirm`, // 假设的确认接口，后端需要实现
        method: 'put',
        data,
    });
};

/**
 * 员工将草稿状态的发票提交审核
 * @param {number[]} ids 发票ID列表
 */
export const publishInvoices = (ids) => {
    return request({
        url: '/invoices/publish', // 假设的发布接口，后端需要实现
        method: 'post',
        data: { ids },
    });
};

/**
 * 更新发票信息
 * @param {number} id 发票ID
 * @param {object} data 更新的数据
 */
export const updateInvoice = (id, data) => {
  return request({
    url: `/invoices/${id}`,
    method: 'put',
    data,
  });
};

/**
 * 删除发票
 * @param {number} id 发票ID
 */
export const deleteInvoice = (id) => {
  return request({
    url: `/invoices/${id}`,
    method: 'delete',
  });
};
