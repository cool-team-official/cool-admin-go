package main

import (
	_ "github.com/cool-team-official/cool-admin-go/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"

	_ "github.com/cool-team-official/cool-admin-go/modules"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gres"

	"github.com/cool-team-official/cool-admin-go/internal/cmd"
)

func main() {
	gres.Dump()
	cmd.Main.Run(gctx.New())
}
