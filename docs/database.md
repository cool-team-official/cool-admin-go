# 数据库

[返回目录](README.md)

::: tip 提示
`CoolAdminGo` 采用 **双 ORM 架构**：

- **GoFrame gdb**：负责所有业务读写（`g.DB(group).Model(table)`），并提供软删除、时间字段自动维护、多数据源(分组)等能力；
- **Gorm**：仅用于**自动建表**（`AutoMigrate`）与初始数据填充。

因此你会在 Model 上同时看到 `gorm:"..."` 建表标签，业务代码里却用 `cool.DBM(...)`（gdb）查询——两者是分工关系，不是重复。
:::

## 一、多数据源(分组)

配置文件 `manifest/config/config.yaml` 中，`database` 下**每个 key 即一个数据源（GoFrame 中的分组 group）**：

```yaml
database:
  default:      # 数据源名称,不指定时 default 为默认数据源
    type: "mysql"
    host: "127.0.0.1"
    port: "3306"
    user: "root"
    pass: "123456"
    name: "cooltest"
    charset: "utf8mb4"
    timezone: "Asia/Shanghai"
    debug: true
    createdAt: "createTime"   # gdb 逻辑时间字段 -> 物理列
    updatedAt: "updateTime"
    deletedAt: "deleteTime"   # 软删除列
  # baseConfig:               # 第二个数据源示例(sqlite)
  #   type: "sqlite"
  #   link: "base-config.sqlite"
  #   extra: busy_timeout=5000
  #   createdAt: "createTime"
  #   updatedAt: "updateTime"
  #   debug: true
```

每种数据库的连接配置说明与示例，见[配置](config.md)页的“数据库配置”章节（支持 SQLite / MySQL / PostgreSQL）。

## 二、数据模型与表

每个 Model 文件（如 `modules/base/model/base_sys_user.go`）承担两件事：

1. **定义结构**：内嵌 `*cool.Model` 获得基础字段，业务字段用 `gorm` 标签声明列信息（供建表），`json` 标签声明接口字段；
2. **自动建表**：`init()` 中调用 `cool.CreateTable(...)`，当配置 `cool.autoMigrate: true`（默认）时首次运行会自动 `AutoMigrate`。

```go
package model

import "github.com/cool-team-official/cool-admin-go/cool"

type BaseSysUser struct {
    *cool.Model
    DepartmentID uint    `gorm:"column:departmentId;type:bigint;index" json:"departmentId"` // 部门ID
    Username     string  `gorm:"column:username;type:varchar(100);not null;index" json:"username"`
    Password     string  `gorm:"column:password;type:varchar(255);not null" json:"password"`
    Status       *int32  `gorm:"column:status;not null;default:1" json:"status"` // 状态 0:禁用 1:启用
}
```

### 2.1 基础字段(内嵌 `cool.Model`)

见 `cool/model.go`，包含 `ID/CreateTime/UpdateTime/DeletedAt`（详见[CRUD](crud.md)页表格）。

### 2.2 时间与软删除的维护方

| 能力 | 维护方 | 机制 |
|---|---|---|
| `createTime/updateTime` 自动写入 | Gorm(写)+gdb(读) | gorm 标签 `autoCreateTime:nano / autoUpdateTime:nano` 建表时写值;gdb 查询时也据此过滤/排序 |
| 软删除 | **gdb** | 依据 `database.*.deletedAt: "deleteTime"` 配置;查询自动加 `deleteTime IS NULL`,`Delete` 自动转 `UPDATE deleteTime=now` |
| 建表 | Gorm | `cool.CreateTable → InitDB(group) → AutoMigrate` |

::: warning 注意
`cool.Model` 中声明的逻辑字段名是 `DeletedAt`，但配置映射的是物理列 `deleteTime`。若自行实现 Model 而非内嵌 `cool.Model`，请保证物理列与配置一致。
:::

## 三、代码中的查询方式

统一通过 `cool.DBM(model)` 获取对应分组的 gdb Model：

```go
import "github.com/cool-team-official/cool-admin-go/cool"

// DBM 依据 model.GroupName() 选中数据源,表名为 model.TableName()
m := cool.DBM(model.NewBaseSysUser())

count, err := m.Where("username", "admin").Count()
var user *model.BaseSysUser
err = m.Where("id", 1).Scan(&user)
```

- 内部服务层已封装好常规 CRUD（见[CRUD](crud.md)），业务代码多数时候只需写 `QueryOp` 配置，无需手写 SQL。
- 复杂查询可随时用 `Extend` 或在 Service 方法内直接使用 gdb 链式调用。

::: tip 已废弃
`cool.GDBM` 已标记 `Deprecated`,请使用 `cool.DBM` 替代（`cool/db.go`）。
:::

## 四、连接初始化原理

1. 每个 Model 的 `init()` 调用 `cool.CreateTable`；
2. `cool.CreateTable` 首次遇到某数据源时调用 `InitDB(group)`（`cool/initdb.go`）——读取 **gdb 同一份配置**，再通过 `cool/cooldb.GetConn(config)` 建立 Gorm 连接并缓存到 `cool.GormDBS[group]`；
3. 连接**按需创建、按分组缓存**，因此直到真正用到某数据源才会建连。

## 五、驱动注册(选择数据库)

数据库驱动位于 `contrib/drivers/`，通过**空导入**在 `main.go` 中启用（对应关系）：

| 数据库 | import | 注册名 |
|---|---|---|
| MySQL | `_ "github.com/cool-team-official/cool-admin-go/contrib/drivers/mysql"` | `mysql`/`mariadb`/`tidb` |
| PostgreSQL | `_ "github.com/cool-team-official/cool-admin-go/contrib/drivers/pgsql"` | `pgsql` |
| SQLite | `_ "github.com/cool-team-official/cool-admin-go/contrib/drivers/sqlite"` | `sqlite` |

```go
// main.go
import (
    _ "github.com/cool-team-official/cool-admin-go/contrib/drivers/mysql" // 按需选择其一
    // _ "github.com/cool-team-official/cool-admin-go/contrib/drivers/pgsql"
    // _ "github.com/cool-team-official/cool-admin-go/contrib/drivers/sqlite"

    _ "github.com/cool-team-official/cool-admin-go/modules"
)
```

::: warning 注意
驱动 import 应**早于使用数据库的包**；为防止编辑器自动排序改变注册顺序，可在其下方保留一个空行。
:::

驱动通过 `cooldb.Register(name, driver)` 自注册（`cool/cooldb/cooldb.go`），你也可以按同样接口编写自己的驱动：

```go
type Driver interface {
    GetConn(node *gdb.ConfigNode) (db *gorm.DB, err error)
}
func Register(name string, driver Driver) error
```

## 六、初始数据填充

模块首次运行时可将内置 JSON 数据整包写入表（见 `cool/initdb.go` 的 `FillInitData`），通过 `base_sys_init` 表记录初始化进度避免重复。模块入口中调用，例如 `modules/base/base.go`：

```go
cool.FillInitData(ctx, "base", &model.BaseSysMenu{})
cool.FillInitData(ctx, "base", &model.BaseSysUser{})
// ...
```

数据文件打包在模块 `resource/` 中（见模块 `packed` 目录），如需定制初始菜单/角色/参数，修改对应 JSON 资源后重新 `make pack` 即可。

## 七、常见问题

::: details 修改表结构不生效?
`cool.autoMigrate` 开启时 gorm `AutoMigrate` 只增列/建新表,**不会删除或改类型**。破坏性变更请手工执行 DDL。
:::

::: details 业务查询走哪个库?
Model 实现 `GroupName()` 决定；缺省 `default`。多个库各写一个 group 配置即可。
:::
