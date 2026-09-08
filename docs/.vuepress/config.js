module.exports = {
    title: 'CoolAdminGo文档',
    description: 'CoolAdminGo文档',
    base: '/cool-admin-go/',
    locales: {
        '/': {
            lang: 'zh-CN',
            title: 'CoolAdminGo文档',
            description: 'CoolAdminGo文档',
        },
    },
    themeConfig: {
      logo: '/logo-admin-new.png',
      repo: 'https://github.com/cool-team-official/cool-admin-go',
      docsDir: 'docs',
      editLinks: true,
      editLinkText: '在 GitHub 上编辑此页',
      nav: [
        { text: '首页', link: '/' },
        { text: 'CoolAdmin官网', link: 'https://cool-js.com' },
        { text: 'GoFrame官网', link: 'https://goframe.org' },
      ],
      lastUpdated: '上次更新',
      sidebar: [
        "/",
        {
          title: '基础',
          collapsable: false,
          children: [
            '/introduction',
            '/development',
            '/cli',
            '/quick_start',
            '/config',
          ],
        },
        {
          title: '开发指南',
          collapsable: false,
          children: [
            '/crud',
            '/database',
            '/file_upload',
            '/form_validate',
            '/system_module',
            '/distributed_function',
            '/cron',
          ],
        },
        {
          title: '运维部署',
          collapsable: false,
          children: [
            '/deploy',
            '/log',
          ],
        },
        {
          title: '参与贡献',
          collapsable: false,
          children: [
            '/faq',
            '/contributing',
            '/feedback',
          ],
        },
        {
          title: '版本记录',
          collapsable: false,
          children: [
            '/changelog',
            '/known_issues',
          ],
        },
      ],
    },
  }
  