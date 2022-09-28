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
	g.Log().Debug(ctx, "modules/base init")

	cool.FillInitData("base", &model.BaseSysMenu{})
	cool.FillInitData("base", &model.BaseSysUser{})
	cool.FillInitData("base", &model.BaseSysUserRole{})
	cool.FillInitData("base", &model.BaseSysRole{})
	cool.FillInitData("base", &model.BaseSysRoleMenu{})
	cool.FillInitData("base", &model.BaseSysDepartment{})
	cool.FillInitData("base", &model.BaseSysRoleDepartment{})
	cool.FillInitData("base", &model.BaseSysParam{})

}
