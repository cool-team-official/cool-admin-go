package demo

import (
	_ "github.com/cool-team-official/cool-admin-go/modules/space/controller"
	_ "github.com/cool-team-official/cool-admin-go/modules/space/middleware"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {
	var (
		ctx = gctx.GetInitCtx()
	)
	g.Log().Debug(ctx, "module space init start ...")
	g.Log().Debug(ctx, "module space init finished ...")
}
