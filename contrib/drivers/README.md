# drivers

数据库驱动包,集合了goframe的gdb驱动 和gorm的驱动

# 安装

以mysql为例.

```bash
go get -u github.com/cool-team-official/cool-admin-go/contrib/drivers/mysql
```

在你的项目中选择并导入:

```go
import _ "github.com/cool-team-official/cool-admin-go/contrib/drivers/mysql"
```

在 main.go 的顶部导入:

```go
package main

import (
	_ "github.com/cool-team-official/cool-admin-go/contrib/drivers/mysql"

	// Other imported packages.
)

func main() {
	// Main logics.
}
```

# 支持的数据库

## MySQL/MariaDB/TiDB

```go
import _ "github.com/cool-team-official/cool-admin-go/contrib/drivers/mysql"
```

## SQLite

```go
import _ "github.com/cool-team-official/cool-admin-go/contrib/drivers/sqlite"
```

Note:

- It does not support `Save` features.

## PostgreSQL

```go
import _ "github.com/cool-team-official/cool-admin-go/contrib/drivers/pgsql"
```

Note:

- It does not support `Save/Replace` features.
- It does not support `LastInsertId`.

## SQL Server

```go
import _ "github.com/cool-team-official/cool-admin-go/contrib/drivers/mssql"
```

Note:

- It does not support `Save/Replace` features.
- It does not support `LastInsertId`.
- It supports server version >= `SQL Server2005`

## Oracle

```go
import _ "github.com/cool-team-official/cool-admin-go/contrib/drivers/oracle"
```

Note:

- It does not support `Save/Replace` features.
- It does not support `LastInsertId`.

## ClickHouse

```go
import _ "github.com/cool-team-official/cool-admin-go/contrib/drivers/clickhouse"
```

Note:

- It does not support `InsertIgnore/InsertGetId` features.
- It does not support `Save/Replace` features.
- It does not support `Transaction` feature.
- It does not support `RowsAffected` feature.

# Custom Drivers

It's quick and easy, please refer to current driver source.
It's quite appreciated if any PR for new drivers support into current repo.
