import {useMemo} from 'react'
import {useRequest} from 'ahooks'
import {fetchUserRoutes} from '@/service/api'
import {useDispatch, useSelector} from 'react-redux'
import {componentMount, getFlattenRoutes} from '@/routes/helpers'
import {useHistory} from 'react-router'
import {isArray, registerGlobalFunction} from '@/utils/common'

export type IRoute = {
    name: string;
    breadcrumb?: boolean;
    children?: IRoute[];
    // 当前路由是否渲染菜单项，为 true 的话不会在菜单中显示，但可通过路由地址访问。
    ignore?: boolean;
    // 路由地址
    path?: string;
    is_link?: boolean;
    // 路由组件
    component?: string;
    meta?: {
        // 菜单图标
        icon?: string;
        // 菜单标题
        title?: string;
        // 菜单是否隐藏
        hide?: boolean;
    }
};

// 静态路由
export const staticRoutes: IRoute[] = [
    // {
    //     name: "menu.dashboard",
    //     meta: {
    //         icon: "dashboard",
    //         title: "控制台",
    //     },
    //     children: [
    //         {
    //             name: "menu.dashboard.workplace",
    //             path: "/dashboard/workplace",
    //             component: "dashboard/workplace",
    //             meta: {
    //                 title: "工作台",
    //             }
    //         },
    //     ],
    // },
]

// 动态路由
const useRoute = () => {
    const {routes} = useSelector((state: GlobalState) => state)
    const dispatch = useDispatch()
    const history = useHistory()

    // 获取路由数据
    const dynamicRoutes = useRequest(fetchUserRoutes, {
        manual: true,
        cacheKey: 'app-dynamic-routes',
        onSuccess: async ({data}) => {
            if (!isArray(data)) return
            dispatch({
                type: 'update-routes',
                payload: {
                    routes: await componentMount([...staticRoutes, ...data])
                },
            })
        }
    })

    // 刷新路由全局方法
    registerGlobalFunction('refreshRoutes', () => dynamicRoutes.runAsync())

    // 默认路由
    const defaultRoute = useMemo(() => {
        const first = routes.find(r => r.isHome == 1) || routes[0]
        if (first) {
            const _path = first?.children?.[0]?.path || first.path

            return _path?.replace(/^\//, '')
        }
        return ''
    }, [routes])

    // 获取当前路由
    const getCurrentRoute = () => getFlattenRoutes(routes).find((tab) => tab.path.split('?')[0] === history.location.pathname.replace(/\/\d+/g, '/:id'))

    return {routes, defaultRoute, getCurrentRoute}
}

export default useRoute
