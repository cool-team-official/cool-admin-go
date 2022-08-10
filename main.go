package main

import (
	_ "cool-admin-go/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"cool-admin-go/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
