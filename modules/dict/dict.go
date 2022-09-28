package dict

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	_ "github.com/cool-team-official/cool-admin-go/modules/dict/packed"
	"github.com/gogf/gf/v2/frame/g"

	_ "github.com/cool-team-official/cool-admin-go/modules/dict/controller"
	"github.com/cool-team-official/cool-admin-go/modules/dict/model"
)

func init() {
	var (
		ctx g.Ctx
	)
	g.Log().Debug(ctx, "modules/dict init")
	cool.FillInitData("dict", &model.DictInfo{})
	cool.FillInitData("dict", &model.DictType{})
}
