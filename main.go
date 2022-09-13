package main

import (
	_ "github.com/cool-team-official/cool-admin-go/internal/packed"

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
