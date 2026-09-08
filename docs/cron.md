# 定时任务

[返回目录](README.md)

::: tip 提示
`task` 模块（`modules/task`）提供可视化的定时任务管理：支持 **cron 表达式** 与 **时间间隔** 两种任务类型，任务可指定要执行的[分布式函数](distributed_function.md)，执行结果自动记录日志，并支持"单例任务"防止并发重复执行。
:::

## 一、核心数据表

### 1. `task_info`(任务定义)

| 字段 | 类型 | 说明 |
|---|---|---|
| `jobId` | varchar | 任务 ID(gcron 用) |
| `name` | varchar | 任务名称(**唯一**,重复会被 `UniqueKey` 拦截) |
| `type` | int | 0:系统 1:用户 |
| `taskType` | int | **0:cron 表达式 1:时间间隔** |
| `cron` | varchar | cron 表达式(robfig/cron 格式,支持 `@every 10s`) |
| `every` | int | 间隔时间(毫秒;taskType=1 时生效,换算为秒后转 `@every Ns`) |
| `limit` | int | 限制次数,空为不限制 |
| `service` | varchar | 执行的服务函数串,如 `TaskTest(hello)` —— 对应[分布式函数](distributed_function.md)名 |
| `status` | int | 0:关闭 1:开启 |
| `startDate` / `endDate` | datetime | 生效时间段(早于 `startDate` 不执行) |
| `data` | varchar | 附带数据(透传给函数) |
| `remark` | varchar | 备注 |

### 2. `task_log`(执行日志)

| 字段 | 说明 |
|---|---|
| `taskId` | 所属任务 id |
| `status` | 0:失败 1:成功 |
| `detail` | 明细(失败为错误信息,成功为"任务执行成功") |

::: tip 日志保留策略
见 `TaskInfoService.Record`：**成功日志每个任务只保留最新 20 条**（超出自动删除），失败日志不删除，方便排查。
:::

## 二、管理接口

控制器位于 `modules/task/controller/admin/task_info.go`，标准 CRUD 之外提供：

| 方法 | 路径 | 说明 |
|---|---|---|
| `POST` | `/admin/task/info/once` | **立即执行一次**(`id`) |
| `POST` | `/admin/task/info/start` | 开启任务(`id`,见下) |
| `POST` | `/admin/task/info/stop` | 停止任务(`id`) |
| `GET` | `/admin/task/info/log` | 查询任务执行日志(可带 `id`/`status`,分页) |

新增任务时 `status=1` 即会通过 `ModifyAfter`(Add) 自动启用调度。

## 三、启用/停止的原理

### 启用 `EnableTask(ctx, cronId, funcstring, cron, startDate)`

```text
检查 FuncMap 中存在 funcstring 对应函数
├─ 是单例函数(IsSingleton=true)
│    → gcron.AddSingleton(按 cron 到点触发,若正在执行则跳过本次)
├─ 否则
│    → gcron.Add(普通调度)
每次到点:若当前时间 < startDate 则跳过;
        否则 cool.RunFunc(funcstring),并按结果写 task_log
最后 SetNextRunTime 刷新 nextRunTime
```

### 停止 `DisableTask(ctx, cronId)`

```go
gcron.Remove(cronId) // 移除该任务的全部 cron 调度
```

## 四、增删改的联动(ModifyAfter)

`TaskInfoService.ModifyAfter` 在框架对任务做增删改后自动调度分布式函数，保证"改配置即生效"：

| 动作 | 触发 |
|---|---|
| Add 且 `status=1` | `ClusterRunFunc("TaskAddTask(id)")` —— 恢复/挂起调度 |
| Update 且 `status=1` | `ClusterRunFunc("TaskStartFunc(id)")` —— 按新 cron 重启 |
| Update 且 `status=0` | `ClusterRunFunc("TaskStopFunc(id)")` —— 停掉调度 |
| Delete | `ClusterRunFunc("TaskStopFunc(id)")` —— 先停调度再删 |

## 五、内置调度函数

均在 `modules/task/funcs/` 中注册（`TaskStartFunc`/`TaskStopFunc`/`TaskTest`/`TaskAddTask`）：

- `TaskStartFunc(id)`：将任务 `status` 置 1 并按 `taskType` 组装 cron(`@every Ns`)后 `EnableTask`;
- `TaskStopFunc(id)`：`DisableTask` 并置 `status` 0;
- `TaskAddTask(id)`：进程启动时用于恢复开启中的任务调度;
- `TaskTest(param)`：测试函数,可立即验证调度链路。

> 注意：这些函数 `IsAllWorker()==true`,因此**多副本下每个节点都会各自挂 cron**（保证某个节点挂了任务仍能执行），真正执行任务体时再按函数本身是否单例收敛。若你的任务函数不想被多副本重复执行，实现 `IsSingleton()==true` 并在函数内做好幂等即可。

## 六、进程启动自动恢复

`modules/task/task.go` 在模块 `init()` 时：

1. 填充初始数据（`FillInitData`）；
2. 查询所有 `status=1` 的任务；
3. 逐个 `cool.RunFunc(ctx, "TaskAddTask("+id+")")` 把调度恢复到 gcron。

因此重启服务后**开启中的任务会自动恢复**，无需人工干预。

## 七、写一个自定义定时任务

```go
// 1. 定义函数(任意位置,建议放业务模块 funcs 子包)
type MyJob struct{}
func (f *MyJob) Func(ctx g.Ctx, param string) error {
    g.Log().Info(ctx, "我的定时任务开始", param)
    // 业务逻辑...
    return nil
}
func (f *MyJob) IsSingleton() bool { return true } // 单例
func (f *MyJob) IsAllWorker() bool { return false } // 仅主节点执行
func init() { cool.RegisterFunc("MyJob", &MyJob{}) }

// 2. 管理端新增一条任务记录:
//    name=MyJob、taskType=0、cron=0 */5 * * * *、service=MyJob(参数写在括号里)
//    status 设为 1 即自动开始调度
```

::: warning 注意
- `service` 字段的写法必须与函数注册名一致且带括号,如 `MyJob()`、`TaskTest(abc)`;括号内参数传给 `Func` 的 `param`。
- 任务模块依赖分布式函数机制,多机部署请确认缓存为 Redis(否则 `ClusterRunFunc` 退化为本机执行)。
:::
