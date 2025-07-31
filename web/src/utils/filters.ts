import dayjs from 'dayjs'

export const formatDate = (d: Date | string, dateFormat = 'YYYY/MM/DD HH:mm:ss'): string => {
  return dayjs(d).format(dateFormat)
}

export const formatTimeOnly = (d: Date | string, timeFormat = 'HH:mm') => {
  return formatDate(d, timeFormat)
}

export const $filters = {
  formatDate,
  formatTimeOnly,
}
