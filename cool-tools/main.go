package main

import (
	// _ "github.com/cool-team-official/cool-admin-go/contrib/drivers/mssql"
	_ "github.com/cool-team-official/cool-admin-go/contrib/drivers/mysql"
	_ "github.com/cool-team-official/cool-admin-go/contrib/drivers/sqlite"
	"github.com/cool-team-official/cool-admin-go/cool-tools/internal/cmd"
	_ "github.com/cool-team-official/cool-admin-go/cool-tools/internal/packed"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	// gres.Dump()

	err := cmd.Main.RunWithError(gctx.New())
	if err != nil {
		println(err.Error())
	}
}
