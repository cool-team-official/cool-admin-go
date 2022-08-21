package base

import (
	_ "github.com/cool-team-official/cool-admin-go/modules/base/packed"

	"github.com/cool-team-official/cool-admin-go/cool"
	_ "github.com/cool-team-official/cool-admin-go/modules/base/controller"
	_ "github.com/cool-team-official/cool-admin-go/modules/base/middleware"
	"github.com/cool-team-official/cool-admin-go/modules/base/model"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	var (
		ctx g.Ctx
	)
	g.Log().Info(ctx, "modules/base init")
	cfg, err := g.Cfg().Get(ctx, "cool-base")
	if err != nil {
		g.Log().Panic(ctx, err)
	}
	if cfg.IsEmpty() {
		g.Log().Panic(ctx, "modules/base config is empty")
	}
	cool.FillInitData("base", &model.BaseSysMenu{})
	cool.FillInitData("base", &model.BaseSysUser{})
	cool.FillInitData("base", &model.BaseSysUserRole{})
	cool.FillInitData("base", &model.BaseSysRole{})
	cool.FillInitData("base", &model.BaseSysRoleMenu{})
	cool.FillInitData("base", &model.BaseSysDepartment{})
	cool.FillInitData("base", &model.BaseSysRoleDepartment{})

}
