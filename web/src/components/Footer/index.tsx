import {GithubOutlined} from '@ant-design/icons'
import {DefaultFooter} from '@ant-design/pro-components'
import React from 'react'

const Footer: React.FC = () => {
    return (
        <DefaultFooter
            style={{background: 'none'}}
            copyright={false}
            links={[
                {
                    key: 'github',
                    title: <GithubOutlined/>,
                    href: 'https://github.com/slowlyo/amis-base',
                    blankTarget: true,
                },
                {
                    key: 'amis Base',
                    title: 'amis Base',
                    href: 'https://github.com/slowlyo/amis-base',
                    blankTarget: true,
                },
            ]}
        />
    )
}

export default Footer
