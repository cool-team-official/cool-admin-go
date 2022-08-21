package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
)

var (
	Init = gcmd.Command{
		Name:  "init",
		Usage: "init",
		Brief: "init",
		Arguments: []gcmd.Argument{
			{
				Name: "dst",
				// Short: "m",
				Brief: "the destination path",
				IsArg: true,
			},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			dst := parser.GetArg(2).String()
			g.Dump(dst)
			if dst == "" {
				dst = "."
			}
			// 如果目标路径不存在，则创建目标路径
			if gfile.IsEmpty(dst) {
				if err = gfile.Mkdir(dst); err != nil {
					return
				}
			}

			gres.Export("cool-admin-go-simple", dst, gres.ExportOption{
				RemovePrefix: "cool-admin-go-simple",
			})
			return nil
		},
	}
)

// init 初始化模块
func init() {
	Main.AddCommand(&Init)
}
