package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gres"
)

var (
	Dump = &gcmd.Command{
		Name:        "dump",
		Usage:       "cool-tools dump",
		Brief:       "查看打包的资源文件",
		Description: "查看打包的资源文件",
		Arguments:   []gcmd.Argument{},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			gres.Dump()
			return nil
		},
	}
)

// init
func init() {
	Main.AddCommand(Dump)
}
