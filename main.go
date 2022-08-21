package main

import (
	_ "cool-admin-go-simple/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"

	_ "cool-admin-go-simple/modules"

	"github.com/gogf/gf/v2/os/gctx"

	"cool-admin-go-simple/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
