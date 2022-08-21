package main

import (
	_ "github.com/cool-team-official/cool-admin-go/cool-tools/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gres"

	"github.com/cool-team-official/cool-admin-go/cool-tools/internal/cmd"
)

func main() {
	gres.Dump()
	cmd.Main.Run(gctx.New())
}
