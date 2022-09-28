# 文件上传驱动

## 本地文件上传(local)

配置文件中的配置项示例：

```yaml
cool:
  file:
    mode: "local"
    domain: "http://127.0.0.1:8002"
```
驱动引入：

```go
import _ "github.com/cool-team-official/cool-admin-go/contrib/files/local"
```