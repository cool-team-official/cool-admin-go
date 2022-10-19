# cool-admin-go sqlite 驱动包

扩展了 GoFrame 的 sqlite 包,集成了 gorm 相关功能.

## 使用方法

引入规则应早于 `modules`相关引入,建议在 main.go 中进行引入。

```go
import (
    _ "github.com/cool-team-official/cool-admin-go/contrib/drivers/sqlite"

    // 换行然后再入模块包,防止编辑器自动排序导致引入顺序错乱
    _ "github.com/cool-team-official/cool-admin-go/modules/base"

)
```

## 配置

```yaml
database:
  default:
    type: "sqlite" # 数据库类型
    name: "cool.sqlite" # 数据库名称,对于sqlite来说就是数据库文件名
    extra: busy_timeout=5000 # 扩展参数 如 busy_timeout=5000&journal_mode=ALL
    createdAt: "createTime" # 创建时间字段名称
    updatedAt: "updateTime" # 更新时间字段名称
    debug: true # 开启调试模式,启用后将在控制台打印相关sql语句
```
