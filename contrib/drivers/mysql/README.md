# cool-admin-go mysql 驱动包

扩展了 GoFrame 的 mysql 包,集成了 gorm相关功能.

## 使用方法

引入规则应早于 `modules`相关引入,建议在 main.go 中进行引入。

```go
import (
    _ "github.com/cool-team-official/cool-admin-go/contrib/drivers/mysql"

    // 换行然后再入模块包,防止编辑器自动排序导致引入顺序错乱
    _ "github.com/cool-team-official/cool-admin-go/modules/base"

)
```
## 配置
