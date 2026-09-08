# 操作日志

[返回目录](README.md)

::: tip 提示
`base` 模块内置了一套**请求操作日志**能力：所有 `/admin/*` 请求自动记录操作人、动作、IP 与请求参数(脱敏)到 `base_sys_log` 表，并支持保留天数自动清理。开关在 `config.yaml` 的 `modules.base.middleware.log.enable`。
:::

## 一、开启/关闭

```yaml
modules:
  base:
    middleware:
      log:
        enable: true   # 是否记录操作日志
```

## 二、记录时机(中间件链)

`modules/base/middleware/middleware.go` 中根据配置注册拦截器：

```go
g.Server().BindMiddleware("/admin/*", BaseLog) // 仅 /admin 前缀请求会被记录
```

`BaseLog`(见 `middleware/log.go`)在进入业务前调用 `BaseSysLogService.Record(ctx)` 落库，再 `Next()` 放行。因此**记录的是"请求快照"，不包含响应结果**。

## 三、落库字段

见 `service/base_sys_log.go` 的 `Record`：

| 字段 | 内容 |
|---|---|
| `userId` | 当前登录人 id(未登录/公共接口为 0) |
| `action` | `HTTP方法 + ":" + 请求路径`，如 `POST:/admin/base/sys/user/delete` |
| `ip` | 客户端 IP |
| `ipAddr` | 客户端 IP(占位,同 ip;如需归属地可接入第三方 IP 库) |
| `params` | 请求体(已脱敏,见下) |

## 四、请求参数脱敏规则

为避免明文密码/验证码/令牌写入数据库，`params` 落库前经过处理：

1. **文件上传等 `multipart/` 请求**：不读取 body，写占位 `[文件上传内容省略]`；
2. **其它非 JSON/文本/表单编码(如二进制)请求**：写 `[二进制内容省略]`；
3. **JSON/表单请求**：
   - 移除敏感键 `password / oldPassword / newPassword / verifyCode / refreshToken / accessToken / secret`（大小写不敏感，非 JSON 则整段隐藏）；
   - 长度超过 **4000 字符**截断为 `...[内容过长已截断]`。

这样即使登录/改密/刷新 token 的请求也不会把敏感信息泄入日志。

## 五、查看与维护接口

`modules/base/controller/admin/base_sys_log.go`，前缀 `/admin/base/sys/log`：

| 方法 | 路径 | 说明 |
|---|---|---|
| 标准六动作 | `/admin/base/sys/log/list` 等 | 列表/分页(支持 `action`/`ip`/`userId` 等值筛选与关键字搜索) |
| `POST` | `/admin/base/sys/log/setKeep` | 设置日志保留天数(写 `base_sys_conf.logKeep`) |
| `GET` | `/admin/base/sys/log/getKeep` | 读取当前保留天数 |
| `POST` | `/admin/base/sys/log/clear` | 立即清空(`isAll=true` 全清)或按保留天数清理 |

日志列表通过 `LeftJoin base_sys_user` 带出操作人姓名(`userName`)。

## 六、清理逻辑

`Clear(isAll bool)`：

```text
isAll=true : DELETE 全表
isAll=false: 读取 base_sys_conf.logKeep(默认天数),删除 createTime < now-keepDays 的记录
```

框架同时注册了一个**清理函数** `BaseFuncClearLog`（`modules/base/funcs/base_sys_log_clear.go`），可在[定时任务](cron.md)里配置周期清理：

```go
// param: "true" 全清,否则按保留天数
cool.RegisterFunc("BaseFuncClearLog", &BaseFuncClearLog{})
```

建议：在管理端「系统设置 → 定时任务」新建一条 cron 任务(如每天凌晨执行 `BaseFuncClearLog()`)，再配合 `logKeep` 保留天数即可无人值守。

::: warning 注意
- `base_sys_log` 表**不应**被业务直接修改；清理请走 `clear` 接口或 `BaseFuncClearLog` 函数。
- 若对写入性能敏感（高并发 /admin 请求），可在 `config.yaml` 关闭 `log.enable`，或自行扩展为异步/消息队列落库。
:::
