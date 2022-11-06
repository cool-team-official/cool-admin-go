# 配置

[返回目录](README.md)

## 配置文件

`CoolAdminGo`配置文件默认位于`mainfest/config/config.yaml`.继承自`GoFrame`的配置文件，支持多种格式，包括`yaml`、`toml`、`json`、`ini`、`xml`等。可访问[GoFrame 配置文件](https://goframe.org/pages/viewpage.action?pageId=1114668)查看更多配置文件格式。

## 数据库配置

`CoolAdminGo`支持多种数据库，可通过引入不同的依赖包进行切换，同时您也可以开发自己的数据库驱动。

### SQLite 配置

使用`sqllite`数据库时，需在`main.go`中引入`_ "github.com/cool-team-official/cool-admin-go/contrib/drivers/sqlite"`包，然后在`config.yaml`中配置`sqlite`数据库。

::: warning 注意
`sqlite`引入应早于使用数据库的包， 为防止编辑器自动排序，可在数据包引入下方加一个空行。
:::

```go
// main.go
import (
    // 引入sqlite驱动
    _ "github.com/cool-team-official/cool-admin-go/contrib/drivers/sqlite"

    // 引入其他包
    "github.com/cool-team-official/cool-admin-go/pkg/dao"

)
```

配置文件中相关配置如下：

```yaml
database:
  default: # 数据源名称,当不指定数据源时 default 为默认数据源
    type: "sqlite" # 数据库类型
    link: "cool.sqlite" # 数据库文件名称，可以带路径，如：/tmp/cool.sqlite
    extra: busy_timeout=5000 # 数据库连接扩展参数
    createdAt: "createTime" # 创建时间字段
    updatedAt: "updateTime" # 更新时间字段
    debug: true # 是否开启调试模式，开启后会打印SQL日志
```

### MySQL 配置

使用`mysql`数据库时，需在`main.go`中引入`_ "github.com/cool-team-official/cool-admin-go/contrib/drivers/mysql"`包，然后在`config.yaml`中配置`mysql`数据库。

::: warning 注意
`mysql`引入应早于使用数据库的包， 为防止编辑器自动排序，可在数据包引入下方加一个空行。
:::

```go
// main.go
import (
    // 引入mysql驱动
    _ "github.com/cool-team-official/cool-admin-go/contrib/drivers/mysql"

    // 引入其他包
    "github.com/cool-team-official/cool-admin-go/pkg/dao"

)
```

配置文件中相关配置如下：

```yaml
database:
  default: # 数据源名称,当不指定数据源时 default 为默认数据源
    type: "mysql" # 数据库类型
    host: "127.0.0.1" # 数据库地址
    port: "3306" # 数据库端口
    user: "root" # 数据库用户名
    pass: "123456" # 数据库密码
    name: "cooltest" # 数据库名称
    charset: "utf8mb4" # 数据库编码
    timezone: "Asia/Shanghai" # 数据库时区
    debug: true # 是否开启调试模式，开启后会打印SQL日志
    createdAt: "createTime" # 创建时间字段
    updatedAt: "updateTime" # 更新时间字段
```

可以使用`docker-compose`快速启动一个`mysql`数据库，配置如下：

```yaml
# docker-compose.yaml
version: "3"
services:
  # mysql8 数据库
  mysql8:
    image: mysql:8
    container_name: mysql8
    command: --default-authentication-plugin=mysql_native_password --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
    # restart: always # 重启策略
    environment:
      TZ: Asia/Shanghai # 指定时区
      MYSQL_ROOT_PASSWORD: "123456" # 配置root用户密码
      MYSQL_DATABASE: "cooltest" # 业务库名
      MYSQL_USER: "cooltest" # 业务库用户名
      MTSQL_PASSWORD: "123123" # 业务库密码
    ports:
      - 3306:3306
    volumes:
      - ./data/mysql/:/var/lib/mysql/
```

启动`mysql`数据库：

```bash
docker compose  -f "docker-compose.yml" up -d --build mysql8
```

关闭`mysql`数据库：

```bash
docker compose  -f "docker-compose.yml" down mysql8
```

