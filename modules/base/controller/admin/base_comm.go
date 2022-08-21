package admin

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	v1 "github.com/cool-team-official/cool-admin-go/modules/base/api/v1"
	"github.com/cool-team-official/cool-admin-go/modules/base/service"
)

type BaseCommController struct {
}

func init() {
	var base_comm_controller = &BaseCommController{}
	// 注册路由
	cool.RegisterControllerSimple("/admin/base/comm", base_comm_controller)
}

// BaseCommPerson 方法
func (c *BaseCommController) Person(ctx context.Context, req *v1.BaseCommPersonReq) (res *cool.BaseRes, err error) {
	var (
		baseSysUserService = service.NewBaseSysUserService()
		admin              = cool.GetAdmin(ctx)
	)
	data, err := baseSysUserService.Person(admin.UserId)
	res = cool.Ok(data)
	return
}

// BaseCommPremmenu 方法
func (c *BaseCommController) Permmenu(ctx context.Context, req *v1.BaseCommPermmenuReq) (res *cool.BaseRes, err error) {

	var (
		baseSysPermsService = service.NewBaseSysPermsService()
		admin               = cool.GetAdmin(ctx)
	)
	res = cool.Ok(baseSysPermsService.Permmenu(ctx, admin.RoleIds))
	return
}
