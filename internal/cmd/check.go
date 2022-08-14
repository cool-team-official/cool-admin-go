package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Check = gcmd.Command{
		Name:  "check",
		Usage: "check",
		Brief: "check",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Debug(ctx, "Check controller~~~~~~~~~~")
			group := "default"
			var configMap map[string]interface{}
			var configSlice []interface{}
			if v, _ := g.Cfg().Get(ctx, "database."+group); !v.IsEmpty() {
				if v.IsSlice() {
					g.Log().Debug(ctx, "v.IsSlice()")
					configSlice = v.Slice()
					configMap = configSlice[0].(map[string]interface{})
				} else if v.IsMap() {
					g.Log().Debug(ctx, "v.IsMap()")
					configMap = v.Map()
				} else {
					panic("无法解析数据库配置")
				}
			}
			g.Dump(configMap)
			return nil
		},
	}
)

// init 初始化模块
func init() {
	Main.AddCommand(&Check)
}
