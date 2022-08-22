package cmd

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool-tools/internal/service"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Module = &gcmd.Command{
		Name:        "module",
		Usage:       "cool-tools module moduleName",
		Brief:       "在modules目录下创建模块",
		Description: "在modules目录下创建模块, 并且创建相应的目录结构,注意: 如果模块已经存在, 则会覆盖原有的模块,本命令需在项目根目录下执行.",
		Arguments: []gcmd.Argument{
			{
				Name:   "moduleName",
				IsArg:  true,
				Orphan: false,
			},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			moduleName := parser.GetArg(2).String()
			if moduleName == "" {
				println("moduleName is empty")
				return nil
			}
			err = service.CreatModule(ctx, moduleName)
			return
		},
	}
	M = &gcmd.Command{
		Name:        "m",
		Usage:       "cool-tools module moduleName",
		Brief:       "在modules目录下创建模块,为module的简写模式",
		Description: "在modules目录下创建模块, 并且创建相应的目录结构,注意: 如果模块已经存在, 则会覆盖原有的模块,本命令需在项目根目录下执行.",
		Arguments: []gcmd.Argument{
			{
				Name:   "moduleName",
				IsArg:  true,
				Orphan: false,
			},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			moduleName := parser.GetArg(2).String()
			if moduleName == "" {
				println("moduleName is empty")
				return nil
			}
			err = service.CreatModule(ctx, moduleName)
			return
		},
	}
)

func init() {
	Main.AddCommand(Module)
	Main.AddCommand(M)
}
