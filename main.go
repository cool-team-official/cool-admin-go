package main

import (
	_ "cool-admin-go-simple/internal/packed"

	_ "github.com/cool-team-official/cool-admin-go/contrib/drivers/sqlite"

	_ "cool-admin-go-simple/modules"

	"github.com/gogf/gf/v2/os/gctx"

	"cool-admin-go-simple/internal/cmd"
)

func main() {
	// gres.Dump()
	cmd.Main.Run(gctx.New())
}
