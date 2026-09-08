# 增删改查(CRUD)

[返回目录](README.md)

::: tip 提示
`CoolAdminGo` 内置了一套**通用 CRUD 引擎**：只要定义好 `Model`、`Service`、`Controller`，即可自动获得 `Add / Delete / Update / Info / List / Page` 六个标准接口，无需手写 SQL 与路由。
:::

## 一、整体结构

一个业务功能由三部分组成，以 `demo` 模块为例：

```bash
modules/demo/
├── controller/           # 控制器:负责路由与参数接收
│   └── admin/
│       └── demo_sample.go
├── model/                # 数据模型:对应一张数据表
│   └── demo_sample.go
└── service/              # 服务层:业务逻辑(可配置查询方式)
    └── demo_sample.go
```

- **Model** 内嵌 `*cool.Model`，自带 `id / createTime / updateTime / deletedAt` 基础字段；
- **Service** 内嵌 `*cool.Service`，声明模型后自动获得增删改查实现；
- **Controller** 内嵌 `*cool.Controller`，配置接口前缀与开放的能力（`Api` 白名单），由框架自动注册路由。

## 二、定义 Model

```go
package model

import "github.com/cool-team-official/cool-admin-go/cool"

const TableNameDemoSample = "demo_sample"

// DemoSample 数据表 <demo_sample>
type DemoSample struct {
    *cool.Model
    // 在此声明业务字段,例如:
    Name string `gorm:"column:name;type:varchar(255);not null;comment:名称" json:"name"`
}

// TableName 返回表名(必须实现)
func (*DemoSample) TableName() string { return TableNameDemoSample }

// GroupName 返回数据源分组,默认 "default",可不实现(继承默认)
// func (*DemoSample) GroupName() string { return "default" }

// NewDemoSample 构造函数
func NewDemoSample() *DemoSample {
    return &DemoSample{ Model: cool.NewModel() }
}

// init 首次运行时自动建表(autoMigrate 开启时)
func init() {
    cool.CreateTable(&DemoSample{})
}
```

::: tip 基础字段说明
内嵌的 `*cool.Model` 提供（见 `cool/model.go`）：

| 字段 | 类型 | 说明 |
|---|---|---|
| `ID` | `uint` | 主键,对应 `id` |
| `CreateTime` | `time.Time` | 创建时间,对应 `createTime`,自动写入纳秒时间戳 |
| `UpdateTime` | `time.Time` | 更新时间,对应 `updateTime`,自动更新 |
| `DeletedAt` | `time.Time` | 软删除标记;由 GoFrame gdb 依据配置自动维护,删除自动转为更新该列 |
:::

## 三、定义 Service

```go
package service

import (
    "github.com/cool-team-official/cool-admin-go/cool"
    "github.com/cool-team-official/cool-admin-go/modules/demo/model"
)

type DemoSampleService struct {
    *cool.Service
}

func NewDemoSampleService() *DemoSampleService {
    return &DemoSampleService{
        &cool.Service{
            Model: model.NewDemoSample(),
        },
    }
}
```

仅配置 `Model` 即拥有完整 CRUD。`cool.Service` 还支持注入各种查询配置（详见下文 [五、查询配置](#五查询配置service)）。

## 四、定义 Controller 并注册路由

```go
package admin

import (
    "context"

    "github.com/cool-team-official/cool-admin-go/cool"
    "github.com/cool-team-official/cool-admin-go/modules/demo/service"
    "github.com/gogf/gf/v2/frame/g"
)

type DemoSampleController struct {
    *cool.Controller
}

func init() {
    var demo_sample_controller = &DemoSampleController{
        &cool.Controller{
            Prefix:  "/admin/demo/demo_sample",                     // 接口前缀
            Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"}, // 开放的内置接口
            Service: service.NewDemoSampleService(),
        },
    }
    // 注册路由(自动挂响应封装中间件并绑定)
    cool.RegisterController(demo_sample_controller)
}
```

启动后即可获得以下接口（`Prefix` 与请求结构体 `g.Meta` 中 `path` 拼接）：

| 方法 | 路径 | 说明 | 请求参数 |
|---|---|---|---|
| `POST` | `/admin/demo/demo_sample/add` | 新增 | 业务字段 JSON |
| `POST` | `/admin/demo/demo_sample/delete` | 删除(可批量) | `ids`: 数组 |
| `POST` | `/admin/demo/demo_sample/update` | 修改 | `id` + 业务字段 |
| `GET` | `/admin/demo/demo_sample/info` | 详情 | `id` |
| `POST` | `/admin/demo/demo_sample/list` | 列表 | `order`/`sort`,见下 |
| `POST` | `/admin/demo/demo_sample/page` | 分页 | `page`/`size`/`order`/`sort`,见下 |

::: tip `Api` 白名单
`Api` 中列出的动作才会被注册；不写即默认全部不开放（访问返回 404）。`cool.Controller` 内置 6 个动作的名称即 `Add/Delete/Update/Info/List/Page`（见 `cool/controller.go`）。
:::

## 五、在 Controller 上追加自定义接口

标准 CRUD 之外，可像普通方法一样定义任意接口。方法签名遵循 `(ctx, req) (res, err)`，请求结构体用 `g.Meta` 标注 `path` 与 `method`，框架会自动将其绑定到同一前缀下。

```go
// 自定义接口:GET /admin/demo/demo_sample/welcome
type DemoSampleWelcomeReq struct {
    g.Meta `path:"/welcome" method:"GET"`
}
type DemoSampleWelcomeRes struct {
    *cool.BaseRes
    Data interface{} `json:"data"`
}

func (c *DemoSampleController) Welcome(ctx context.Context, req *DemoSampleWelcomeReq) (res *DemoSampleWelcomeRes, err error) {
    res = &DemoSampleWelcomeRes{
        BaseRes: cool.Ok("Welcome to Cool Admin Go"),
        Data:    map[string]interface{}{"name": "Cool Admin Go"},
    }
    return
}
```

## 六、查询配置(Service)

在 `cool.Service` 中可通过以下字段精细控制查询行为（见 `cool/service.go`）：

### 6.1 `QueryOp` 查询条件

| 字段 | 说明 |
|---|---|
| `FieldEQ []string` | 等值筛选字段（参数名与数据库列名一致时自动 `where`） |
| `KeyWordField []string` | 模糊搜索字段,前端传 `keyWord` 即在这几个字段上 `like` |
| `AddOrderby g.MapStrStr` | 追加默认排序,如 `{"orderNum": "asc"}` |
| `Where func(ctx) []g.Array` | 自定义条件,返回 `{"列 操作符 ?", 值, true}` 三元组 |
| `Select string` | 指定查询字段,如 `"id,name"` 或 `"a.id,a.name,b.name AS bname"` |
| `Join []*JoinOp` | 关联查询配置 |
| `Extend func(ctx, m) *gdb.Model` | 追加任意查询条件 |
| `ModifyResult func(ctx, data)` | 返回前修改结果 |

### 6.2 `JoinOp` 关联查询

```go
type JoinOp struct {
    Model     IModel   // 关联的 Model
    Alias     string   // 别名
    Condition string   // 关联条件,如 "base_sys_log.userId = u.id"
    Type      JoinType // LeftJoin / RightJoin / InnerJoin
}
```

实际范例（日志列表关联用户表取姓名,`modules/base/service/base_sys_log.go`）：

```go
&cool.Service{
    Model: model.NewBaseSysLog(),
    PageQueryOp: &cool.QueryOp{
        FieldEQ: []string{"action", "ip", "userId"},
        KeyWordField: []string{"action", "ip", "ipAddr"},
        Select: "l.*, u.name AS userName",
        Join: []*cool.JoinOp{
            {
                Model:     model.NewBaseSysUser(),
                Alias:     "u",
                Condition: "l.userId = u.id",
                Type:      "leftJoin",
            },
        },
        AddOrderby: map[string]string{"l.id": "desc"},
    },
}
```

### 6.3 其它常用字段

| 字段 | 说明 | 范例 |
|---|---|---|
| `InsertParam func(ctx) g.MapStrAny` | `Add` 时自动注入参数(如当前登录用户) | `modules/base/service/base_sys_role.go` |
| `Before func(ctx) error` | CRUD 执行前的钩子 | — |
| `InfoIgnoreProperty string` | `Info` 忽略的字段(逗号分隔,如密码) | `modules/base/service/base_sys_user.go` 的 `"password"` |
| `UniqueKey g.MapStrStr` | 唯一键校验,`key:列名 value:错误提示` | `{"name": "名称已存在"}` |
| `NotNullKey g.MapStrStr` | 非空键校验 | `{"name": "名称不能为空"}` |

### 6.4 请求后/前钩子 `ModifyBefore` / `ModifyAfter`

`Controller` 在 `Add/Delete/Update` 前后调用对应 Service 方法（若被实现）：

```go
// 例如:新增/修改后刷新关联数据(见 task_info / base_sys_role service)
func (s *XxxService) ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error) {
    return
}
```

### 6.5 覆盖默认实现

需要完全自定义某个动作时,在 Service 上实现同名方法即可覆盖内置逻辑（接口见 `cool.IService`）：`ServiceAdd / ServiceDelete / ServiceUpdate / ServiceInfo / ServiceList / ServicePage`。

例如 `modules/base/service/base_sys_user.go` 覆盖 `ServiceInfo` 追加角色信息、覆盖 `ServiceAdd` 对密码做 `md5`：

```go
func (s *BaseSysUserService) ServiceInfo(ctx context.Context, req *cool.InfoReq) (data interface{}, err error) {
    // ...先取默认详情,再附加 roleIdList...
}
func (s *BaseSysUserService) ServiceAdd(ctx context.Context, req *cool.AddReq) (data interface{}, err error) {
    // rmap["password"] = md5(rmap["password"]) 后交给默认逻辑
    return s.Service.ServiceAdd(ctx, req)
}
```

## 七、纯接口控制器(ControllerSimple)

若某个控制器只需要自定义接口、不需要内置 CRUD，可内嵌 `cool.ControllerSimple`（见 `cool/controller-simple.go`）：

```go
type BaseOpen struct {
    *cool.ControllerSimple
}

func init() {
    var open = &BaseOpen{
        ControllerSimple: &cool.ControllerSimple{Prefix: "/admin/base/open"},
    }
    cool.RegisterControllerSimple(open)
}
```

范例：`modules/base/controller/admin/base_comm.go`（登录后通用接口组）、`base_open.go`（登录/验证码/eps 开放接口组）。

## 八、完整最小示例(demo 模块)

直接阅读内置 `modules/demo` 模块即可获得一个开箱即用的完整 CRUD 示例：

```bash
modules/demo/
├── demo.go                        # 模块入口(import controller 等触发注册)
├── controller/
│   ├── controller.go              # import admin 子包
│   └── admin/demo_sample.go       # Controller + 注册 + Welcome 自定义接口
├── model/demo_sample.go           # Model
└── service/demo_sample.go         # Service
```

把 `modules/demo` 整个复制一份并替换表名/字段/前缀，即可快速开始新业务模块的开发。

::: warning 注意
- 业务表请**显式实现 `TableName()`**；`GroupName()` 缺省为 `default`，多数据源场景按需实现。
- 涉及二进制/大字段时注意接口返回值结构；自定义接口响应推荐直接返回 `cool.Ok(data)` / `cool.Fail(msg)`（见[响应规范](config.md)相关章节与 `cool/cool.go`）。
:::
