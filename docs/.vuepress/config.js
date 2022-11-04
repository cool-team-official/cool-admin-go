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
        "/introduction",
        "/feedback",
        "/development",
        "/cli",
        "/quick_start",
        "/known_issues",
      ],
    },
  }
  