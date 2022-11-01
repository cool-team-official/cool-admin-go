package main

import (
	_ "github.com/cool-team-official/cool-admin-go/internal/packed"

	_ "github.com/cool-team-official/cool-admin-go/contrib/files/local"

	// Minio，按需启用
	_ "github.com/cool-team-official/cool-admin-go/contrib/files/minio"

	// 阿里云OSS，按需启用
	_ "github.com/cool-team-official/cool-admin-go/contrib/files/oss"

	_ "github.com/cool-team-official/cool-admin-go/contrib/drivers/sqlite"

	_ "github.com/cool-team-official/cool-admin-go/contrib/drivers/mysql"

	_ "github.com/cool-team-official/cool-admin-go/modules"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/cool-team-official/cool-admin-go/internal/cmd"
)

func main() {
	// gres.Dump()
	cmd.Main.Run(gctx.New())
}
