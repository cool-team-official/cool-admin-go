# 系统模块

[返回目录](README.md)

::: tip 提示
系统管理能力全部位于 `modules/base` 模块（包含登录鉴权、用户、角色、菜单、部门、参数、日志、通用接口等）。它是框架内置的唯一"业务型"模块，也是学习整套框架的最佳范本。
:::

## 一、base 模块总览

```
modules/base/
├── base.go                 # 模块入口:注册 controller/service,填充初始数据
├── api/v1/                 # 请求/响应结构体(每个接口一个 Req)
├── controller/
│   └── admin/              # 管理端控制器(内嵌 cool.Controller / ControllerSimple)
├── middleware/             # authority 鉴权、i18n、日志中间件
├── model/                  # base_sys_* 系列数据模型
├── service/                # 服务层(登录/权限/用户/菜单/…)
├── resource/initjson/      # 初始数据 JSON(打包进二进制)
└── packed/                 # gres 打包后的资源(勿手工编辑)
```

## 二、开放接口(登录前)

见 `controller/admin/base_open.go`（`cool.ControllerSimple`,前缀 `/admin/base/open`）：

| 方法 | 路径 | 说明 |
|---|---|---|
| `GET` | `/admin/base/open/captcha` | 图形验证码(参数 `height`/`width`,默认 40/150) |
| `GET` | `/admin/base/open/eps` | 前端路由/权限点信息(供前端初始化) |
| `POST` | `/admin/base/open/login` | 登录(用户名/密码/验证码) |
| `POST` | `/admin/base/open/refreshToken` | 刷新 access token |

## 三、通用接口(登录后)

见 `controller/admin/base_comm.go`（前缀 `/admin/base/comm`）：

| 方法 | 路径 | 说明 |
|---|---|---|
| `GET` | `/admin/base/comm/person` | 当前登录人信息 |
| `GET` | `/admin/base/comm/permmenu` | 当前用户菜单与按钮权限 |
| `POST` | `/admin/base/comm/logout` | 退出登录 |
| `POST` | `/admin/base/comm/upload` | 上传文件(见[文件上传](file_upload.md)) |
| `GET` | `/admin/base/comm/uploadMode` | 上传模式 |
| `POST` | `/admin/base/comm/personUpdate` | 更新个人信息(头像/昵称等) |

## 四、鉴权与权限范围(三档路由)

`modules/base/middleware/authority.go` 实现 **JWT + SSO + 数据权限** 三合一的拦截器：

| 路由规则 | 鉴权要求 |
|---|---|
| `/admin/*/open/*` | 无需登录(白名单,含验证码/登录/eps) |
| `/admin/*/comm/*` | 仅需有效 token(上传/个人信息/权限菜单等) |
| `/admin/base/sys/*` 等其余 | 有效 token + **权限点校验**,URL 转 `a:b:c`(如 `/admin/base/sys/user/page` → `base:sys:user:page`)与用户 perms 比对;超管(`userId==1`)全部放行 |

关键点(加深理解):

- **token 结构**:JWT(HS256),payload 含 `userId/passwordVersion` 等;`passwordVersion` 用于改密/重置密码后**使旧 token 立即失效**(用户表每次改密会更新该字段)。
- **权限缓存**:登录/刷新后写 `admin:perms:<userId>`、`admin:department:<userId>`(见 `service/base_sys_perms.go`),每次请求无需查库。
- **数据范围**:`admin:department:<userId>` 缓存了该用户可见的部门树,用户管理/日志等列表据此过滤(`departmentId` 关联)。
- **弱默认密钥**:内置默认 JWT secret 存在时会警告并**自动替换为随机值**(1.5.x 安全加固)。

## 五、标准管理接口(前缀 `/admin/base/sys`)

以下控制器均内嵌 `cool.Controller`、开放标准六动作(Add/Delete/Update/Info/List/Page),并追加少量自定义接口:

| 控制器 | 前缀 | 表/模型 | 说明与追加接口 |
|---|---|---|---|
| 用户 | `/admin/base/sys/user` | `base_sys_user` | 增删改查/详情;**`POST /move`** 调整部门;**密码 md5 存储**;Info 忽略 `password`;`InsertParam` 注入操作人 |
| 角色 | `/admin/base/sys/role` | `base_sys_role` | 增删改查;**关联菜单/部门**,`userId`(创建人)自动注入 |
| 菜单 | `/admin/base/sys/menu` | `base_sys_menu` | 增删改查;类型 0目录/1菜单/2按钮;`orderNum` 排序 |
| 部门 | `/admin/base/sys/department` | `base_sys_department` | 增删改查 |
| 参数 | `/admin/base/sys/param` | `base_sys_param` | 增删改查(key-value 配置) |
| 日志 | `/admin/base/sys/log` | `base_sys_log` | 列表/分页/清空(**无 Add/Delete/Update**),见[日志](log.md) |

### 关键表结构

| 表 | 主要字段 |
|---|---|
| `base_sys_user` | `departmentId,name,username,password,passwordVersion,nickName,headImg,phone,email,status(0禁用1启用),remark,socketId` |
| `base_sys_role` | `name,label,remark,userId,relevance`(数据权限是否关联上下级) |
| `base_sys_menu` | `parentId,name,router,perms,type,icon,orderNum,viewPath,keepAlive,isShow` |
| `base_sys_department` | `name,parentId,orderNum` |
| `base_sys_param` | `name,key,data(JSON 字符串),remark` |
| `base_sys_log` | `userId,action,ip,ipAddr,params,createTime` |
| `base_sys_user_role` | `userId,roleId` |
| `base_sys_role_menu` | `roleId,menuId` |
| `base_sys_role_department` | `roleId,departmentId` |

## 六、登录流程

1. 前端先 `GET captcha` 取得 `captchaId` 与图片;
2. `POST login` 提交 `username/password/captchaId/verifyCode`;验证码从缓存读取并**一次性消费**,5 分钟有效;
3. 校验密码(md5),失败连续 5 次则锁定 10 分钟(`admin:loginFail:<user>:<ip>`);
4. 成功后返回 `{ expire, token, refreshExpire, refreshToken }`;服务端写 `admin:token:<userId>`、`admin:token:refresh:<userId>` 缓存;
5. 前端每请求带 `Authorization: Bearer <token>`,接近过期用 `refreshToken` 换新;
6. 退出(`logout`)清除上述缓存与 `admin:perms/admin:department`。

## 七、前端配合(electron 系 / Vue3)

- 菜单渲染与按钮权限数据来自 `GET /admin/base/comm/permmenu`,返回 `{ perms:[...], menus:[...] }`;
- 前端路由表依据 `menus.viewPath` 映射到页面组件,按钮 `v-permission` 指令校验 `perms`。
- 详情见[快速开始](quick_start.md)与前端工程 `src/cool` 目录。
