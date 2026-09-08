# 常见问题

[返回目录](README.md)

::: tip 提示
汇总开发与部署中最常遇到的问题。若这里没有你的问题，欢迎到 [issue](https://github.com/cool-team-official/cool-admin-go/issues) 反馈，或在[问题反馈](feedback.md)了解渠道。
:::

## 一、框架/架构类

### Q:为什么既有 Gorm 又有 GoFrame gdb？会不会混乱？

不会。它们是**分工**关系（详见[数据库](database.md)）：

- **Gorm** 只负责 `AutoMigrate` 自动建表（`cool.CreateTable`）与初始数据填充；
- 业务查询统一走 **gdb**：`cool.DBM(model)` / 框架封装好的 Service。

绝大部分场景你只需定义 Model + QueryOp，两个 ORM 的细节框架都已隔离。

### Q:单模块(如 `modules/demo`)单独 `go build` 报错,为什么？

**设计如此。** 本仓库是多模块 monorepo，`cool` 库与各模块 go.mod 彼此独立；`cool/initdb.go` 不内置数据库驱动（驱动由宿主 `main.go` 空导入注册）。所以：

- 在仓库根目录(聚合了所有模块与驱动的工程)构建才完整；
- 发布/联调用 `go work`(见 `RELEASE.md`)；单独模块只做库级编译验证。

### Q:数据库驱动需要 import 才能生效？

对。MySQL/PgSQL/SQLite 驱动位于 `contrib/drivers/*`，必须 `_ "..."` 空导入才会注册。若配置了某个库却没导入对应驱动，建表/连接时会报错。见[数据库](database.md)驱动章节。

## 二、文件上传类

### Q:上传接口报错 / panic？

大概率是 **`cool.file.mode` 与 import 的驱动不匹配**：

- `mode: "local"` 却没导入 `contrib/files/local` → 调用 `cool.File()` 时注册表为空；
- 导入了 minio/oss 但 `mode` 不是对应值 → 该驱动 `New()` 返回 nil。

请核对 `main.go` import 与 `config.yaml` 保持一致（见[文件上传](file_upload.md)）。

### Q:MinIO/OSS 的 `uploadMode` 返回的 mode 是 `local`？

这是历史遗留：`contrib/files/minio` 与 `oss` 的 `GetMode()` 返回 `mode:"local"`（代码沿袭）。**实际存储模式以 `cool.file.mode` 配置为准**，前端如需精确判断请按配置读取。

### Q:上传成功但浏览器访问 404？

本地驱动文件在 `./public/uploads/...`,需要后端启动时注册了 `/public` 静态路由且 `cool.file.domain` 可访问；生产用 Nginx 时要放行 `/public/`,或改用 MinIO/OSS。

## 三、登录/权限类

### Q:验证码总是失败？

- 验证码缓存于 `cool:captcha:*`,**一次性**,5 分钟内有效;
- 请确认前后端时钟一致、`captchaId` 与图片为同一次获取;
- Redis 模式下确认 `redis.cool` 可用(否则验证码写在内存,重启即失效)。

### Q:账号被锁定？

连续登录失败 5 次会锁定 10 分钟(键 `admin:loginFail:<用户名>:<IP>`)。等待或由管理员清理 Redis 中该键。

### Q:改密/重置密码后旧 token 还有效？

框架通过用户表 `passwordVersion` 字段使**旧 token 立即失效**：改密会更新该字段，JWT 校验时比对版本号。若仍能访问,请检查是否多个 base 实例/缓存不一致。

### Q:接口 401 / 权限不足?

- 确认请求头 `Authorization: Bearer <token>` 正确、token 未过期;
- 新加接口要在对应菜单/角色上配置权限点(URL 会转 `base:sys:xxx:xxx` 比对);
- 超管(`userId==1`)默认全放行,可先用超管验证业务本身。

## 四、定时任务/分布式类

### Q:任务不执行?

- 确认任务 `status=1` 且 `service` 函数名与注册名一致(带括号,如 `MyJob()`)；
- 看 `task_log` 是否有失败记录;
- 单机部署确认没配 Redis 也正常(退化为本地执行);多副本必须配 Redis,且任务函数体按需实现幂等;
- cron 时间注意 `server.timezone` 与容器/服务器时区一致。

### Q:多副本任务重复执行?

多副本时调度函数 `TaskAddTask` 等 `IsAllWorker()==true`,**每个节点都会挂 cron**;真正执行任务体时,若函数 `IsSingleton()==true` 会由 `cool:masterflag` 选举的单点执行,但**到点瞬间多个节点可能同时触发**,因此关键任务函数请自行保证幂等/加锁(见[定时任务](cron.md)与[分布式函数](distributed_function.md))。

## 五、其他

### Q:配置文件名打错?

文档/示例曾出现 `mainfest` 笔误,正确目录是 `manifest/config/config.yaml`(`manifest` 无 `in`)。老教程里的 `pkg/dao` 也已不存在,统一用 `cool.DBM`。

### Q:文档站本地预览?

```bash
# 仓库根目录
yarn            # 安装 vuepress(首次)
export NODE_OPTIONS=--openssl-legacy-provider   # Node 17+ 需要
yarn docs:dev   # http://localhost:8080
```

详见[贡献代码](contributing.md)的文档部分。

### Q:升级 GoFrame 大版本后编译报错?

`cool` 库与各模块对 gf 版本敏感,请以根 `go.mod` 锁定版本为准,多模块同时升级,勿只改其中一个模块的 go.mod。
