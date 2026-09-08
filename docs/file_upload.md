# 文件上传

[返回目录](README.md)

::: tip 提示
`CoolAdminGo` 内置了统一的文件上传能力，通过配置即可在 **本地存储 / MinIO / 阿里云 OSS** 之间切换，业务代码无需改动。
:::

## 一、内置上传接口

框架已提供两个可直接调用的接口（登录即可访问，见 `modules/base/controller/admin/base_comm.go`）：

| 方法 | 路径 | 说明 |
|---|---|---|
| `GET` | `/admin/base/comm/uploadMode` | 获取当前上传模式 |
| `POST` | `/admin/base/comm/upload` | 上传文件(form-data,字段名 `file`) |

前端拿到返回的地址后即可回显/入库。

## 二、工作原理

统一抽象在 `cool/coolfile/coolfile.go`：

```go
type Driver interface {
    New() Driver
    GetMode() (data interface{}, err error)
    Upload(ctx g.Ctx) (string, error)   // 返回可访问的文件地址
}
```

- 工厂 `cool.File()` 依据配置 `cool.file.mode` 从注册表选择驱动；
- 各驱动在包级 `init()` 中通过 `coolfile.Register(name, driver)` 自注册。

```go
// cool/file.go
var File = coolfile.NewFile   // 使用:cool.File()
```

因此**必须 import 对应驱动包才会注册**，否则配置了未注册的模式会在调用时 panic。

## 三、本地存储(Local)

### 配置

```yaml
cool:
  file:
    mode: "local"                 # local | minio | oss
    domain: "http://127.0.0.1:8002"  # 访问域名
```

### 启用

```go
// main.go
import (
    _ "github.com/cool-team-official/cool-admin-go/contrib/files/local"
)
```

### 行为

- 文件保存到 `./public/uploads/YYYYMMDD/`（按日分目录），文件名自动生成;
- 自动创建 `./public` 静态目录并注册 `/public` 静态路由（`s.AddStaticPath("/public", "./public")`），所以返回地址形如 `http://127.0.0.1:8002/public/uploads/20260908/xxx.png`，可直接访问；
- 上传表单字段为 `file`，为空时报"上传文件为空"。

## 四、MinIO

### 配置

MinIO 复用 `cool.file.oss` 配置节（注释即标明“oss 配置项兼容 minio”）：

```yaml
cool:
  file:
    mode: "minio"
    domain: "http://127.0.0.1:8002"  # 业务访问域名(与 mode 无关时可不关注)
    oss:                              # minio 同样读取该节
      endpoint: "192.168.192.110:9000"  # MinIO 地址(不带 http)
      accessKeyID: "accessKeyID"
      secretAccessKey: "secretAccessKey"
      bucketName: "cool-admin-go"
      useSSL: false                    # minio 使用
      location: "us-east-1"            # minio 使用
```

### 启用

```go
import (
    _ "github.com/cool-team-official/cool-admin-go/contrib/files/minio"
)
```

### 行为

- `New()` 仅在 `cool.file.mode == "minio"` 时创建客户端并 `MakeBucket`（不存在时自动创建）；
- 文件 `PutObject` 到 `uploads/YYYYMMDD/<key|随机16位>`，返回对象地址 `info.Location`。

## 五、阿里云 OSS

### 配置

```yaml
cool:
  file:
    mode: "oss"
    oss:
      endpoint: "oss-cn-hangzhou.aliyuncs.com"
      accessKeyID: "accessKeyID"
      secretAccessKey: "secretAccessKey"
      bucketName: "cool-admin-go"   # 需要 bucket 公开读
```

### 启用

```go
import (
    _ "github.com/cool-team-official/cool-admin-go/contrib/files/oss"
)
```

### 行为

- 自动检查/创建 bucket（公开读）；
- `PutObject` 到 `uploads/YYYYMMDD/<key|随机16位>`，返回公网地址 `https://<bucket>.<endpoint>/<path>`。

## 六、在代码中直接调用

除内置接口外，业务代码也可直接调用：

```go
import "github.com/cool-team-official/cool-admin-go/cool"

// 在任意请求处理方法内
url, err := cool.File().Upload(ctx)   // 读取当前请求中的 file 字段并上传
if err != nil {
    return cool.Fail(err.Error()), err
}
res = cool.Ok(url)
return
```

::: warning 注意
- 三个驱动包的 import 必须与 `cool.file.mode` 匹配：模式为 `local` 却没 import local 包、或 import 了但模式不是对应值时，都会导致调用失败/panic。
- 官方脚手架模板在 `main.go` 中默认启用 **local**，MinIO/OSS 按需打开注释即可。
:::

## 七、上传模式与前端约定

前端可通过 `GET /admin/base/comm/uploadMode` 获取当前上传模式（返回值含 `mode`/`type`），据此决定回显拼接等行为。

::: warning 已知遗留
`contrib/files/minio` 与 `contrib/files/oss` 的 `GetMode()` 返回值中 `mode` 恒为字符串 `"local"`（实现沿袭本地驱动），实际使用以你配置的 `cool.file.mode` 为准。若前端依赖该字段精确判断，请注意此差异。
:::
