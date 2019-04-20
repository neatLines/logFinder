import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/table/list',
    method: 'get',
    params
  })
}


export function getHosts() {
  return request({
    url: '/v1/hosts',
    method: 'get'
  })
}