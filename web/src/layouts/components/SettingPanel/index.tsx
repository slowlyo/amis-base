import {
    Badge,
    Button,
    Checkbox,
    ColorPicker,
    Drawer,
    Form,
    InputNumber,
    message,
    Radio,
    Select,
    Space,
    Switch,
    Tag
} from 'antd'
import useStore from '@/hooks/useStore'
import {useLang} from '@/hooks/useLang'
import SelectLayout from '@/layouts/components/SettingPanel/components/SelectLayout'
import useTheme from '@/hooks/useTheme'
import useSetting from '@/hooks/useSetting'
import {useHistory} from 'react-router-dom'
import {appLoaded, getCacheKey} from '@/utils/common'
import {useRequest} from 'ahooks'
import {saveSettings} from '@/service/api'
import {useState} from 'react'

// 设置面板
const SettingPanel = () => {
    const history = useHistory()
    const pathname = history.location.pathname
    const {state, dispatch} = useStore()
    const {setThemeColor, setDarkTheme} = useTheme()
    const {settings} = useSetting()
    const {t} = useLang()

    const [cachedSettings, setCachedSettings] = useState(localStorage.getItem(getCacheKey('settings')))

    // 关闭设置面板
    const closeSetting = () => {
        dispatch({
            type: 'update-open-setting',
            payload: {openSetting: false}
        })
    }

    // 保存设置
    const handleChange = (key, value) => {
        if (key === 'themeColor') setThemeColor(value)
        if (key === 'darkTheme') setDarkTheme(value)

        const theme = Object.assign({}, settings.theme, {[key]: value})
        dispatch({
            type: 'update-settings',
            payload: {settings: {...settings, theme}}
        })

        if (key === 'layoutMode') {
            // 解决刷新后页面不显示的问题
            window.$owl.appLoader()
            history.push('/')
            setTimeout(() => {
                history.push(pathname)
                setTimeout(() => {
                    appLoaded()
                }, 500)
            }, 200)
        }
    }

    // 保存设置
    const save = useRequest(saveSettings, {
        manual: true,
        onSuccess: () => {
            message.success(t('theme_setting.save_success'))

            // 保存成功后，更新缓存
            localStorage.setItem(getCacheKey('settings'), JSON.stringify(settings))
            setCachedSettings(JSON.stringify(settings))
        }
    })

    // 提交设置
    const submit = () => {
        save.run({theme: settings.theme})
    }

    // 获取动画选项
    const getAnimateOptions = (type) => {
        const animateOptions = ['alpha', 'left', 'right', 'top', 'bottom', 'scale', 'scaleBig', 'scaleX', 'scaleY']

        return animateOptions.map((item) => {
            return {
                label: t(`theme_setting.animate_${type}_${item}`),
                value: item
            }
        })
    }

    return (
        <Drawer open={state.openSetting}
                onClose={closeSetting}
                closeIcon={false}
                title={t('theme_setting.title')}
                footer={(
                    <>
                        <Space>
                            <Badge dot
                                   count={cachedSettings != JSON.stringify(state.settings) ? 1 : 0}
                                   title={t('theme_setting.need_save')}>
                                <Button type="primary"
                                        onClick={submit}
                                        disabled={save.loading}>{t('theme_setting.save_btn')}</Button>
                            </Badge>
                            <Button onClick={closeSetting}
                                    disabled={save.loading}>{t('theme_setting.cancel_btn')}</Button>

                        </Space>
                    </>
                )}>
            <Form labelAlign="left" labelCol={{span: 8}} wrapperCol={{span: 16}}>
                {/*主题色*/}
                <Form.Item colon={false} label={t('theme_setting.theme_color')}>
                    <ColorPicker showText
                                 disabledAlpha
                                 disabledFormat
                                 onChange={(_, v) => handleChange('themeColor', v)}
                                 value={settings.theme.themeColor}
                                 presets={[{
                                     label: t('theme_setting.preinstall'),
                                     colors: ['#1677FF', '#F5222D', '#FA8C16', '#52C41A', '#13A8A8', '#2F54EB', '#722ED1', '#EB2F96'],
                                 }]}/>
                </Form.Item>

                {/*暗黑模式*/}
                {!settings.theme.followSystemTheme && (
                    <Form.Item colon={false} label={t('theme_setting.dark_theme')}>
                        <Switch checked={settings.theme.darkTheme}
                                onChange={(value) => handleChange('darkTheme', value)}/>
                    </Form.Item>
                )}

                {/*布局模式*/}
                <Form.Item colon={false} label={t('theme_setting.layout_mode')}>
                    <SelectLayout current={settings.theme.layoutMode}
                                  change={(value) => handleChange('layoutMode', value)}/>
                </Form.Item>

                {/*登录页模板*/}
                <Form.Item colon={false} label={t('theme_setting.login_template')}>
                    <Radio.Group
                        onChange={(e) => handleChange('loginTemplate', e.target.value)}
                        value={settings.theme.loginTemplate}
                        optionType="button"
                        options={[
                            {label: t('theme_setting.default'), value: 'default'},
                            {label: 'amis', value: 'amis'}
                        ]}/>
                </Form.Item>

                {/*顶部菜单主题*/}
                {settings.theme.layoutMode != 'double' && (
                    <Form.Item colon={false} label={t('theme_setting.top_theme')}>
                        <Radio.Group
                            onChange={(e) => handleChange('topTheme', e.target.value)}
                            value={settings.theme.topTheme}
                            optionType="button"
                            options={[
                                {label: t('theme_setting.light'), value: 'light'},
                                {label: t('theme_setting.dark'), value: 'dark'}
                            ]}/>
                    </Form.Item>
                )}

                {/*侧边栏主题*/}
                {(settings.theme.layoutMode != 'double' && settings.theme.layoutMode != 'top') && (
                    <Form.Item colon={false} label={t('theme_setting.sider_theme')}>
                        <Radio.Group
                            onChange={(e) => handleChange('siderTheme', e.target.value)}
                            value={settings.theme.siderTheme}
                            optionType="button"
                            options={[
                                {label: t('theme_setting.light'), value: 'light'},
                                {label: t('theme_setting.dark'), value: 'dark'}
                            ]}/>
                    </Form.Item>
                )}

                {/*页面内容*/}
                <Form.Item colon={false} label={t('theme_setting.page_content')}>
                    <Space direction="vertical">
                        <Checkbox onChange={(e) => handleChange('footer', e.target.checked)}
                                  checked={settings.theme.footer}>{t('theme_setting.footer')}</Checkbox>
                        <Checkbox onChange={(e) => handleChange('breadcrumb', e.target.checked)}
                                  checked={settings.theme.breadcrumb}>{t('theme_setting.breadcrumb')}</Checkbox>
                        <Checkbox onChange={(e) => handleChange('enableTab', e.target.checked)}
                                  checked={settings.theme.enableTab}>{t('theme_setting.tab')}</Checkbox>
                        <Checkbox onChange={(e) => handleChange('tabIcon', e.target.checked)}
                                  checked={settings.theme.tabIcon}>{t('theme_setting.tab_icon')}</Checkbox>
                    </Space>
                </Form.Item>

                {/*进场动画*/}
                <Form.Item colon={false} label={t('theme_setting.animate_in')}>
                    <InputNumber addonAfter="ms"
                                 onChange={(value) => handleChange('animateInDuration', value)}
                                 value={settings.theme.animateInDuration}
                                 addonBefore={(
                                     <Select options={getAnimateOptions('in')}
                                             dropdownStyle={{width: 100}}
                                             onChange={(value) => handleChange('animateInType', value)}
                                             value={settings.theme.animateInType}/>
                                 )}/>
                </Form.Item>

                {/*离场动画*/}
                <Form.Item colon={false} label={t('theme_setting.animate_out')}>
                    <InputNumber addonAfter="ms"
                                 onChange={(value) => handleChange('animateOutDuration', value)}
                                 value={settings.theme.animateOutDuration}
                                 addonBefore={(
                                     <Select options={getAnimateOptions('out')}
                                             dropdownStyle={{width: 100}}
                                             onChange={(value) => handleChange('animateOutType', value)}
                                             value={settings.theme.animateOutType}/>
                                 )}/>
                </Form.Item>

                <Form.Item colon={false} label={t('theme_setting.keep_alive')}>
                    <Space size="large">
                        <Switch checked={settings.theme.keepAlive}
                                onChange={(value) => handleChange('keepAlive', value)}/>
                        {/*<Tag color="warning">Alpha</Tag>*/}
                    </Space>
                </Form.Item>

                <Form.Item colon={false} label={t('theme_setting.accordion_menu')}>
                    <Space size="large">
                        <Switch checked={settings.theme.accordionMenu}
                                onChange={(value) => handleChange('accordionMenu', value)}/>
                    </Space>
                </Form.Item>
            </Form>
        </Drawer>
    )
}
export default SettingPanel
