# 贡献代码

[返回目录](README.md)

::: tip 提示
欢迎任何形式的贡献：提 issue、修 bug、补文档、加功能。请先阅读本页约定，让协作更顺畅。
:::

## 一、仓库结构须知

`CoolAdminGo` 是 **Go 多模块 monorepo + 前端(Vue3) + 文档(vuepress)** 的组合仓库：

| 路径 | 说明 |
|---|---|
| 根目录 | 聚合工程:`main.go`、`go.mod`、`internal/`、`modules/`、`contrib/`、`docs/`、`frontend/`、`RELEASE.md` |
| `cool/` | 核心库(独立 go.mod),Controller/Service/Model 引擎、DB、文件、分布式函数等 |
| `modules/base` 等 | 业务模块(各自独立 go.mod),内置模块是扩展的最佳范本 |
| `contrib/` | 驱动与文件驱动扩展(drivers/*、files/*) |
| `docs/` | 本文档站(vuepress) |
| `frontend/` | 前端管理端工程 |

改动**核心库 `cool/`** 或跨模块行为时影响面大，请在 PR 描述里说明并补测试/文档。

## 二、提 Issue

在 [issues](https://github.com/cool-team-official/cool-admin-go/issues) 提交前请先搜索是否已存在。一个高质量 issue 应包含：

- 复现步骤(越具体越好)
- 环境信息：OS / Go 版本 / 数据库 / 是否 Redis / 框架版本
- 相关日志与配置(注意**脱敏**,不要贴密钥)
- 若为界面问题,附截图

## 三、提交 PR 流程

```bash
# 1. Fork 并克隆到本地
git clone https://github.com/<你的用户名>/cool-admin-go
cd cool-admin-go
# 2. 新建功能分支(建议命名:fix/xxx、feat/xxx、docs/xxx)
git checkout -b feat/your-feature origin/master
# 3. 开发,本地自测
go build ./...            # 多模块请先按 RELEASE.md 用 go work 联调
# 4. 提交
git add .
git commit -m "feat: 增加 xxx 能力"
# 5. 推送并创建 PR(目标分支:上游 master)
git push origin feat/your-feature
```

PR 描述建议写清：**动机、改动点、自测方式、是否影响现有接口/数据结构**。

### 版本与分支

- 日常开发基于 `master`;提交信息建议遵循语义化(如 `feat:` / `fix:` / `docs:` / `chore:`);
- 涉及接口/数据结构的破坏性变更请注明,并在 `docs/changelog.md` 记录;
- 正式发布走仓库内 `RELEASE.md` 流程(版本号 + 子模块 tag + GitHub Release)。

## 四、文档贡献

文档站使用 **vuepress**(1.x) 构建,源文件在 `docs/`,站点本身发布到 `gh-pages`。

```bash
# 根目录安装依赖(首次)
yarn

# 本地预览(Node 17+ 需 legacy provider)
export NODE_OPTIONS=--openssl-legacy-provider
yarn docs:dev          # 打开 http://localhost:8080

# 构建产物校验
yarn docs:build        # 输出 docs/.vuepress/dist
```

文档规范（与既有页面保持一致）：

- 每页首行写 `[返回目录](README.md)`;
- 善用 `::: tip` / `::: warning` 提示容器与表格;
- 代码示例必须与当前代码一致,引用到具体文件时标注路径;
- 新增页面记得同步 `docs/README.md` 目录与 `docs/.vuepress/config.js` 的 sidebar;
- 涉及新版本特性时,同步更新 `docs/changelog.md`。

## 五、代码规范

- Go 代码使用 gofmt/goimports 格式化,`go vet` 无错误;
- 请求/响应结构体放在模块 `api/v1/`,控制器只做路由与转发,业务逻辑下沉 `service/`;
- 新表模型实现 `TableName()`(建议再加 `GroupName()`),并放在模块 `init()`/`init()` 里 `cool.CreateTable`;
- 不要 `go get` 直接往各 go.mod 塞大版本依赖;升级 gf 等核心依赖请提独立 PR;
- 敏感信息(密码、token)绝不硬编码、不打日志,参考[日志](log.md)的脱敏实现。

## 六、合并后

- CI 相关:仓库 CI 目前为手动触发(workflow_dispatch),合并 `master` 不会自动跑全量构建;
- 若你的改动需要发布新版本,请在 PR 中说明或与维护者沟通,发布流程见 `RELEASE.md`。

再次感谢你的贡献 🎉
