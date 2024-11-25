// @ts-ignore
/* eslint-disable */
import {request} from '@umijs/max'

const prefix = '/admin-api/'

// 获取系统配置
export async function featSettings() {
    return request<{ data: any; }>(prefix + 'settings', {method: 'GET'})
}

// 保存系统配置
export async function saveSettings(params: any) {
    return request<{ data: any; }>(prefix + 'settings', {method: 'POST', data: params})
}