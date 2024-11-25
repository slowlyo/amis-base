import {useMount, useRequest} from 'ahooks'
import {fetchLogout, fetchSettings, fetchUserInfo} from '@/service/api'
import {
    appLoaded,
    clearMsgSign,
    getCacheKey,
    goToLoginPage,
    inLoginPage,
    registerGlobalFunction,
    Token
} from '@/utils/common'
import {dynamicAssetsHandler} from '@/utils/dynamicAssets'
import {registerCustomComponents} from '@/components/AmisRender/CustomComponents'

// 应用初始化
const useSetup = (store) => {

    // 初始化配置信息
    const initSettings = useRequest(fetchSettings, {
        manual: true,
        retryCount: 3,
        cacheKey: 'app-settings',
        onBefore() {
            store.dispatch({
                type: 'update-userInfo',
                payload: {userLoading: true},
            })
        },
        onSuccess(res) {
            store.dispatch({
                type: 'update-settings',
                payload: {settings: res.data},
            })
        },
        onFinally() {
            store.dispatch({
                type: 'update-inited',
                payload: {inited: true},
            })
        }
    })

    // 初始化用户信息
    const initUserInfo = useRequest(fetchUserInfo, {
        manual: true,
        onSuccess(res) {
            store.dispatch({
                type: 'update-userInfo',
                payload: {userInfo: res.data},
            })
        }
    })

    // 退出登录
    const logout = useRequest(fetchLogout, {
        manual: true,
        onFinally() {
            Token().clear()
            goToLoginPage()
        }
    })

    // 注册全局函数
    const registerFunctions = () => {
        registerGlobalFunction('logout', () => logout.run())
        registerGlobalFunction('getCacheKey', (val: string) => getCacheKey(val))
    }

    // 初始化
    const init = async () => {
        clearMsgSign()
        registerFunctions()
        registerCustomComponents()

        await initSettings.runAsync()

        if (Token().value) {
            let res = await initUserInfo.runAsync()
            if (res.data?.code != 401) {
                await window.$owl.refreshRoutes()
            }
        } else if (!inLoginPage()) {
            goToLoginPage()
        }

        appLoaded()
    }

    // 初始化
    useMount(() => {
        init().then()
    })
}

export default useSetup
