import {useSelector} from 'react-redux'
import RefreshButton from '@/layouts/components/LayoutTopBar/components/RefreshButton'
import FullscreenButton from '@/layouts/components/LayoutTopBar/components/FullscreenButton'
import SettingButton from '@/layouts/components/LayoutTopBar/components/SettingButton'
import useSetting from '@/hooks/useSetting'
import {Avatar, Dropdown, Space} from 'antd'
import {useHistory} from 'react-router'

// 顶部导航栏
const LayoutTopBar = () => {
    const {userInfo} = useSelector((state: GlobalState) => state)
    const {getSetting} = useSetting()
    const history = useHistory()
    // const darkTheme = () => getSetting('theme.topTheme') == 'dark' && getSetting('theme.layoutMode') != 'double'

    const items = [
        {key: 'user', label: '个人中心'},
        {key: 'logout', label: '退出登录'}
    ]

    const onClick = ({key}) => {
        if (key == 'logout') {
            window.$owl.logout()
            return
        }

        if (key == 'user') {
            history.push('/user')
            return
        }
    }

    return (
        <div className="h-full flex justify-around items-center">
            {/* prependNav */}
            {/*{getSetting('nav.prependNav') && <AmisRender schema={getSetting('nav.prependNav')}/>}*/}

            <RefreshButton/>
            {/*<DarkThemeButton/>*/}
            <FullscreenButton/>
            {/*<LocaleButton/>*/}
            {getSetting('dev') && <SettingButton/>}

            {/* appendNav */}
            {/*{getSetting('nav.appendNav') && <AmisRender schema={getSetting('nav.appendNav')}/>}*/}

            <div className="min-w-[120px] pl-3 pr-5">
                <Dropdown menu={{items, onClick}}>
                    <Space>
                        <Avatar size={36} src={<img src={userInfo.avatar} alt="avatar"/>}/>
                        <span>{userInfo.name}</span>
                    </Space>
                </Dropdown>
            </div>
            {/*<div className="user-navbar bg-transparent h-full cursor-pointer min-w-[120px]" style={darkTheme() ? {*/}
            {/*    // @ts-ignore*/}
            {/*    '--colors-neutral-text-1': '#fff',*/}
            {/*} : {}}>*/}
            {/*    <AmisRender className="h-full !text-white min-w-[120px]" schema={userInfo?.menus}/>*/}
            {/*</div>*/}
        </div>
    )
}

export default LayoutTopBar
