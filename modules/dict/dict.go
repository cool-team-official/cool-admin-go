package dict

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	_ "github.com/cool-team-official/cool-admin-go/modules/dict/packed"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/cool-team-official/cool-admin-go/modules/dict/controller"
	"github.com/cool-team-official/cool-admin-go/modules/dict/model"
)

func init() {
	var (
		ctx = gctx.GetInitCtx()
	)
	g.Log().Debug(ctx, "modules/dict init")
	cool.FillInitData(ctx, "dict", &model.DictInfo{})
	cool.FillInitData(ctx, "dict", &model.DictType{})
}
