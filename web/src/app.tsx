import {AvatarDropdown, AvatarName, Footer} from '@/components'
import type {Settings as LayoutSettings} from '@ant-design/pro-components'
import {SettingDrawer} from '@ant-design/pro-components'
// @ts-ignore
import type {RunTimeLayoutConfig} from '@umijs/max'
import {history} from '@umijs/max'
import defaultSettings from '../config/defaultSettings'
import {errorConfig} from './requestErrorConfig'
import {currentUser as queryCurrentUser} from '@/services/ant-design-pro/api'
import React from 'react'
import {featSettings, saveSettings} from '@/services/system'

// @ts-ignore
const loginPath = '/user/login'

/**
 * @see  https://umijs.org/zh-CN/plugins/plugin-initial-state
 * */
export async function getInitialState(): Promise<{
    settings?: Partial<LayoutSettings>;
    currentUser?: API.CurrentUser;
    loading?: boolean;
    fetchUserInfo?: () => Promise<API.CurrentUser | undefined>;
}> {
    const systemSettings = await featSettings()

    const settings = systemSettings?.data || {theme: defaultSettings, dev: false}

    const fetchUserInfo = async () => {
        return {
            name: 'Serati Ma',
            avatar: 'https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png',
            userid: '00000001',
            email: 'antdesign@alipay.com',
            signature: '海纳百川，有容乃大',
            title: '交互专家',
            group: '蚂蚁金服－某某某事业群－某某平台部－某某技术部－UED',
            tags: [
                {key: '0', label: '很有想法的',},
                {key: '1', label: '专注设计',},
                {key: '2', label: '辣~',},
                {key: '3', label: '大长腿',},
                {key: '4', label: '川妹子',},
                {key: '5', label: '海纳百川',},
            ],
            notifyCount: 12,
            unreadCount: 11,
            country: 'China',
            geographic: {province: {label: '浙江省', key: '330000',}, city: {label: '杭州市', key: '330100',},},
            address: '西湖区工专路 77 号',
            phone: '0752-268888888',
        }
    }
    // 如果不是登录页面，执行
    const {location} = history
    if (location.pathname !== loginPath) {
        const currentUser = await fetchUserInfo()

        return {fetchUserInfo, currentUser, settings}
    }

    return {fetchUserInfo, settings}
}

// ProLayout 支持的api https://procomponents.ant.design/components/layout
// @ts-ignore
export const layout: RunTimeLayoutConfig = ({initialState, setInitialState}) => {
    // 处理色弱模式会显w
    let dom = document.querySelector('body')
    if (dom) {
        if (initialState.settings.theme.colorWeak) {
            dom.dataset.prosettingdrawer = dom.style.filter
            dom.style.filter = 'invert(80%)'
        } else {
            dom.style.filter = dom.dataset.prosettingdrawer || 'none'
            delete dom.dataset.prosettingdrawer
        }
    }

    return {
        actionsRender: () => [],
        avatarProps: {
            src: initialState?.currentUser?.avatar,
            title: <AvatarName/>,
            render: (_, avatarChildren) => {
                return <AvatarDropdown menu>{avatarChildren}</AvatarDropdown>
            },
        },
        waterMarkProps: false,
        footerRender: () => <Footer/>,
        onPageChange: () => {
            const {location} = history
            // 如果没有登录，重定向到 login
            if (!initialState?.currentUser && location.pathname !== loginPath) {
                history.push(loginPath)
            }
        },
        bgLayoutImgList: [
            {
                src: 'https://mdn.alipayobjects.com/yuyan_qk0oxh/afts/img/D2LWSqNny4sAAAAAAAAAAAAAFl94AQBr',
                left: 85,
                bottom: 100,
                height: '303px',
            },
            {
                src: 'https://mdn.alipayobjects.com/yuyan_qk0oxh/afts/img/C2TWRpJpiC0AAAAAAAAAAAAAFl94AQBr',
                bottom: -68,
                right: -45,
                height: '303px',
            },
            {
                src: 'https://mdn.alipayobjects.com/yuyan_qk0oxh/afts/img/F6vSTbj8KpYAAAAAAAAAAAAAFl94AQBr',
                bottom: 0,
                left: 0,
                width: '331px',
            },
        ],
        links: [],
        menuHeaderRender: undefined,
        // 自定义 403 页面
        // unAccessible: <div>unAccessible</div>,
        title: initialState.settings.appName,
        // 增加一个 loading 的状态
        childrenRender: (children) => {
            // if (initialState?.loading) return <PageLoading />;
            return (
                <>
                    {children}
                    {initialState?.settings?.dev && (
                        <SettingDrawer
                            disableUrlParams
                            enableDarkTheme={false}
                            hideCopyButton
                            hideHintAlert
                            settings={initialState?.settings?.theme}
                            onSettingChange={(theme) => {
                                setInitialState((preInitialState) => ({
                                    ...preInitialState,
                                    ...{settings: {...initialState?.settings, theme}},
                                }))

                                // 保存设置
                                saveSettings({
                                    key: 'system.theme',
                                    value: theme,
                                }).then()
                            }}
                        />
                    )}
                </>
            )
        },
        ...initialState?.settings?.theme,
    }
}

/**
 * @name request 配置，可以配置错误处理
 * 它基于 axios 和 ahooks 的 useRequest 提供了一套统一的网络请求和错误处理方案。
 * @doc https://umijs.org/docs/max/request#配置
 */
export const request = {
    ...errorConfig,
}
