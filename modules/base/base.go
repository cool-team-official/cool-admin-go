package base

import (
	_ "github.com/cool-team-official/cool-admin-go/modules/base/packed"

	"github.com/cool-team-official/cool-admin-go/cool"
	_ "github.com/cool-team-official/cool-admin-go/modules/base/controller"
	_ "github.com/cool-team-official/cool-admin-go/modules/base/funcs"
	_ "github.com/cool-team-official/cool-admin-go/modules/base/middleware"
	"github.com/cool-team-official/cool-admin-go/modules/base/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {
	var (
		ctx = gctx.GetInitCtx()
	)
	g.Log().Debug(ctx, "modules/base init")

	cool.FillInitData(ctx, "base", &model.BaseSysMenu{})
	cool.FillInitData(ctx, "base", &model.BaseSysUser{})
	cool.FillInitData(ctx, "base", &model.BaseSysUserRole{})
	cool.FillInitData(ctx, "base", &model.BaseSysRole{})
	cool.FillInitData(ctx, "base", &model.BaseSysRoleMenu{})
	cool.FillInitData(ctx, "base", &model.BaseSysDepartment{})
	cool.FillInitData(ctx, "base", &model.BaseSysRoleDepartment{})
	cool.FillInitData(ctx, "base", &model.BaseSysParam{})

}
