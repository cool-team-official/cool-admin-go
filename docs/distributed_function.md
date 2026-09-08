# 分布式函数

[返回目录](README.md)

::: tip 提示
`cool/func.go` 提供了一套 **"函数级分布式调度"** 机制：把一个可复用业务封装成 `CoolFunc` 并注册后，即可被定时任务、REST 接口或集群消息统一驱动执行——单机直接执行，Redis 集群模式下通过订阅发布保证只在一个节点执行。
:::

## 一、核心接口

```go
// cool/func.go
type CoolFunc interface {
    Func(ctx g.Ctx, param string) (err error) // 业务函数体,param 为字符串参数
    IsSingleton() bool   // 单例:同一时刻只允许一个在执行(计划任务中使用)
    IsAllWorker() bool   // true:所有 worker 节点都执行;false:仅主节点执行
}
```

注册与执行：

```go
var FuncMap = make(map[string]CoolFunc)   // 全局函数表
func RegisterFunc(name string, f CoolFunc) // 注册
func GetFunc(name string) CoolFunc          // 查询
func RunFunc(ctx g.Ctx, funcstring string) (err error)      // 执行
func ClusterRunFunc(ctx g.Ctx, funcstring string) (err error) // 集群执行(发布)
func ListenFunc(ctx g.Ctx)                  // 订阅循环(redis 模式)
```

## 二、RunFunc 的函数串语法

`RunFunc` 接收形如 `名称(参数)` 的字符串：

```text
TaskStart(1)          # 名称不带空格,参数可为任意字符串(自行解析)
TaskAddTask(2)
```

- 解析逻辑:取第一个 `(` 前的作为函数名,括号内为 `param`;
- 函数不存在时返回错误 `函数不存在:<名称>`;
- **非 `IsAllWorker()` 的函数**:会先竞争/读取缓存键 `cool:masterflag`(缓存 60s),仅当本进程 `ProcessFlag` 等于该值时真正执行,其余节点跳过——即**单点执行(主进程)**;
- `IsAllWorker()` 返回 `true` 的函数则每个节点都执行(如广播类任务)。

## 三、单机模式 vs Redis 集群模式

由 `cool.CacheManager` 是否使用 Redis 决定(`cool.IsRedisMode`):

| 模式 | `ClusterRunFunc` 行为 | `ListenFunc` |
|---|---|---|
| 单机(默认 gcache) | 直接 `RunFunc` 本地执行 | 未开启;若调用会 panic("集群模式下, 请使用Redis作为缓存") |
| Redis(`redis.cool` 配置生效) | `PUBLISH cool:func <funcstring>` 到 Redis | `main` 启动时 `go cool.ListenFunc(ctx)`,`SUBSCRIBE cool:func` 死循环,收到消息即 `RunFunc` |

因此:线上多副本部署时,开启 Redis 后,调用 `ClusterRunFunc("Xxx(1)")` 只会有一个实例真正跑业务(未设 AllWorker 时),天然避免重复执行。

## 四、真实注册示例

### 1. base 模块:清空日志函数

```go
// modules/base/funcs/base_func.go
type BaseFuncClearLog struct{}
func (f *BaseFuncClearLog) Func(ctx g.Ctx, param string) (err error) {
    // 清空 base_sys_log
    ...
    return
}
func (f *BaseFuncClearLog) IsSingleton() bool { return true }
func (f *BaseFuncClearLog) IsAllWorker() bool { return false }

func init() {
    cool.RegisterFunc("BaseFuncClearLog", &BaseFuncClearLog{})
}
```

### 2. task 模块:任务调度函数(随 task 模块启动注册)

```go
// modules/task/funcs/task_funcs.go(示意)
cool.RegisterFunc("TaskStartFunc",  &taskStartFunc{})   // 启动任务
cool.RegisterFunc("TaskStopFunc",   &taskStopFunc{})    // 停止任务
cool.RegisterFunc("TaskTest",       &taskTest{})        // 测试任务
cool.RegisterFunc("TaskAddTask",    &taskAddTask{})     // 恢复调度(传任务 id)
```

框架在**进程启动时恢复开启中的任务**:`task/task.go` 中扫描 `status=1` 的任务,逐个 `RunFunc("TaskAddTask(id)")` 把调度重新挂到 gcron。

## 五、与定时任务的配合

`task` 模块底层即调用分布式函数执行任务体(见[定时任务](cron.md))。开启任务等于:

```text
(管理端 Start 接口) → service.EnableTask
                        ├─ gcron.AddSingleton/Add(依据 cron/间隔表达式)
                        └─ cron 到点 → cool.RunFunc("TaskStart(id)") → 执行 Task 里配置的函数体
```

若任务函数体需在所有节点执行则实现 `IsAllWorker()=true`,否则默认仅主节点执行。

## 六、注意事项

::: warning
- `ListenFunc` 只在 **Redis 模式**下由 `internal/cmd/cmd.go` 的 `Main` 启动;未配置 Redis 时别调用它。
- 单机模式(无 Redis)下 `ClusterRunFunc` 等同 `RunFunc`,功能不受影响,只是没有跨节点能力。
- 函数名在全局唯一(`RegisterFunc` 直接覆盖),业务模块请加前缀避免与系统函数冲突。
- `param` 是裸字符串,复杂参数请约定分隔符/JSON,在 `Func` 内自行解析。
:::
