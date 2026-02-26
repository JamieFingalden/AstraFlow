/**
 * 将时间值格式化为本地可读时间（YYYY-MM-DD HH:mm:ss）。
 * @param value 原始时间值，支持 ISO 字符串、Date 或空值。
 * @returns 本地化时间字符串；无法解析时返回 "-"。
 */
export const formatDateTime = (value?: string | Date | null): string => {
  if (!value) {
    return '-'
  }

  const date = value instanceof Date ? value : new Date(value)
  if (Number.isNaN(date.getTime())) {
    return '-'
  }

  const year = date.getFullYear()
  const month = `${date.getMonth() + 1}`.padStart(2, '0')
  const day = `${date.getDate()}`.padStart(2, '0')
  const hour = `${date.getHours()}`.padStart(2, '0')
  const minute = `${date.getMinutes()}`.padStart(2, '0')
  const second = `${date.getSeconds()}`.padStart(2, '0')

  return `${year}-${month}-${day} ${hour}:${minute}:${second}`
}

/**
 * 将时间值格式化为日期（YYYY-MM-DD）。
 * @param value 原始时间值，支持 ISO 字符串、Date 或空值。
 * @returns 日期字符串；无法解析时返回 "-"。
 */
export const formatDate = (value?: string | Date | null): string => {
  if (!value) {
    return '-'
  }

  const date = value instanceof Date ? value : new Date(value)
  if (Number.isNaN(date.getTime())) {
    return '-'
  }

  const year = date.getFullYear()
  const month = `${date.getMonth() + 1}`.padStart(2, '0')
  const day = `${date.getDate()}`.padStart(2, '0')

  return `${year}-${month}-${day}`
}
