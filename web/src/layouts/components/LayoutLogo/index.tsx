import {Image} from 'antd'
import useSettings from '@/hooks/useSetting'

// Logo
const LayoutLogo = ({onlyLogo = false}) => {
    const {getSetting} = useSettings()

    // Logo 宽度
    const width = onlyLogo ? ' w-full' : ' w-[220px]'

    // 文字颜色
    const textColor = () => {
        if (getSetting('theme.darkTheme')) {
            return 'text-[var(--colors-neutral-text-1)]'
        }

        if (getSetting('theme.topTheme') === 'dark' && getSetting('theme.layoutMode') == 'top-mix') {
            return 'text-white'
        }

        if (getSetting('theme.topTheme') === 'dark' && getSetting('theme.layoutMode') == 'top') {
            return 'text-white'
        }

        if (getSetting('theme.siderTheme') === 'dark' && getSetting('theme.layoutMode') == 'default') {
            return 'text-white'
        }

        return 'text-[var(--colors-neutral-text-1)]'
    }

    return (
        <div className={'h-[65px] flex justify-center items-center' + width}>
            {getSetting('logo') &&
                <div className="w-[35px]" title={getSetting('appName')}>
                    <Image width={35} preview={false} src={getSetting('logo')}/>
                </div>
            }
            {onlyLogo ||
                <div className={'text-xl font-medium ml-2 truncate ' + textColor()}>{getSetting('appName')}</div>}
        </div>
    )
}

export default LayoutLogo