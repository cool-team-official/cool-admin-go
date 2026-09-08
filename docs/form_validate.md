# 表单验证

[返回目录](README.md)

::: tip 提示
`CoolAdminGo` 基于 **GoFrame 的请求校验**（`v` 标签）与 **响应封装**（`code:1001` 等）实现统一入参校验：非法参数会返回错误码 `51`，提示语可由 `#` 自定义。
:::

## 一、校验标签速查

在请求结构体字段上打 `v` 标签即可，常用规则：

| 标签 | 说明 | 示例 |
|---|---|---|
| `required` | 必填 | `v:"required"` |
| `integer` | 整数 | `v:"integer"` |
| `min/max` | 数值范围 | `v:"min:1"` |
| `length` | 字符串长度 | `v:"length:2,12"` |
| `email` | 邮箱格式 | `v:"email"` |
| `phone` | 手机号 | `v:"phone"` |
| `date` | 日期格式 | `v:"date"` |
| `in` | 枚举白名单 | `v:"in:1,2"` |
| `required-if` | 条件必填 | `v:"required-if:type,1"` |

可同时用 `|` 组合多条；**每条规则后可用 `#` 自定义错误提示**（不写则用默认提示）。

## 二、真实示例

### 1. 登录接口(开放接口,字段全部必填)

`modules/base/api/v1/base_open.go`：

```go
type BaseOpenLoginReq struct {
    g.Meta     `path:"/login" method:"POST"`   // 路由元数据
    Username   string `json:"username" p:"username" v:"required"`     // 用户名
    Password   string `json:"password" p:"password" v:"required"`     // 密码
    CaptchaId  string `json:"captchaId" p:"captchaId" v:"required"`   // 验证码ID
    VerifyCode string `json:"verifyCode" p:"verifyCode" v:"required"` // 验证码
}
```

### 2. 批量删除(自定义错误提示)

`modules/base/api/v1/base_sys_user.go`：

```go
type BaseSysUserDeleteReq struct {
    g.Meta `path:"/delete" method:"POST"`
    Ids    []uint `json:"ids" v:"required#请选择要删除的数据"` // 自定义提示
}
```

### 3. Info 详情(整数必填)

```go
type BaseSysUserInfoReq struct {
    g.Meta `path:"/info" method:"GET"`
    Id     uint `json:"id" v:"integer|required"` // 整数且必填
}
```

### 4. 刷新 token

```go
type BaseOpenRefreshTokenReq struct {
    g.Meta     `path:"/refreshToken" method:"POST"`
    RefreshToken string `json:"refreshToken" v:"required#refreshToken不能为空"`
}
```

### 5. 定时任务启停(整数必填,带自定义提示)

`modules/task/api/v1/task_info.go`：

```go
type TaskInfoStartReq struct {
    g.Meta `path:"/start" method:"POST"`
    Id     uint `json:"id" v:"required#请输入id"`
}
```

## 三、校验失败的响应

校验失败由响应中间件统一转为如下结构(错误码 `51`,语义"参数错误")：

```json
{
  "code": 51,
  "message": "请选择要删除的数据",
  "data": null
}
```

`code` 含义（见 `cool/middleware_handler_response.go` 与 i18n 资源）：

| code | 含义 |
|---|---|
| `1000` | 成功(`BaseResMessage`) |
| `1001` | 业务失败(`Fail`) |
| `51` | 参数错误 |
| `1003` | 服务内部错误 |
| `404`/`403`/`500` | 资源不存在/无权限/服务器错误 |

前端根据 `code === 51` 直接提示 `message` 即可，无需在业务代码里重复书写。

## 四、自定义校验规则

GoFrame 校验器支持在 `main` 中注册自定义规则（`gvalid.RegisterRule`），或在 `init` 中注册 i18n 消息。业务中最常见做法：

- 简单约束直接用 `v` 标签；
- 数据库唯一性/存在性校验建议放在 **Service 层**（`UniqueKey`/`NotNullKey` 或 `ServiceAdd/ServiceUpdate` 中），见[CRUD](crud.md)；
- 涉及用户态（登录/权限）的校验一律在 Service/中间件完成，不要依赖前端。

::: warning 注意
- 请求结构体字段的 `p` 标签用于映射参数名,与 `json` 不同;`v` 校验读取的是 `p`(默认小写字段名)。
- 自定义错误提示用 `#` 拼接在规则后,如 `v:"required#请选择要删除的数据"`。
:::
