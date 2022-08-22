package cmd

import (
	"context"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
)

var (
	Init = gcmd.Command{
		Name:  "init",
		Usage: "cool-tools init [dst]",
		Brief: "创建一个新的cool-admin-go项目",
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
			if dst == "" {
				dst = "."
			}
			// 如果目标路径不存在，则创建目标路径
			if gfile.IsEmpty(dst) {
				if err = gfile.Mkdir(dst); err != nil {
					return
				}
			} else {
				if !gfile.IsDir(dst) {
					g.Log().Panicf(ctx, "%s is not a directory", dst)
				} else {
					s := gcmd.Scanf(`the folder "%s" is not empty, files might be overwrote, continue? [y/n]: `, dst)
					if strings.EqualFold(s, "n") {
						return
					}

				}
			}

			err = gres.Export("cool-admin-go-simple", dst, gres.ExportOption{
				RemovePrefix: "cool-admin-go-simple",
			})
			if err != nil {
				return
			}

			err = gfile.ReplaceDir("cool-admin-go-simple", gfile.Basename(gfile.RealPath(dst)), dst, "*", true)
			if err != nil {
				return
			}
			g.Log().Infof(ctx, "init success")
			return nil
		},
	}
)

// init 初始化模块
func init() {
	Main.AddCommand(&Init)
}
