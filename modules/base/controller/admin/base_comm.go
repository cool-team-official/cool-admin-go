package admin

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/service"
	"github.com/gogf/gf/v2/frame/g"
)

type BaseCommController struct {
	*cool.ControllerSimple
}

func init() {
	var base_comm_controller = &BaseCommController{
		ControllerSimple: &cool.ControllerSimple{
			Perfix: "/admin/base/comm",
		},
	}
	// 注册路由
	cool.RegisterControllerSimple(base_comm_controller)
}

// BaseCommPersonReq 接口请求参数
type BaseCommPersonReq struct {
	g.Meta        `path:"/person" method:"GET"`
	Authorization string `json:"Authorization" in:"header"`
}

// BaseCommPerson 方法
func (c *BaseCommController) Person(ctx context.Context, req *BaseCommPersonReq) (res *cool.BaseRes, err error) {
	var (
		baseSysUserService = service.NewBaseSysUserService()
		admin              = cool.GetAdmin(ctx)
	)
	data, err := baseSysUserService.Person(admin.UserId)
	res = cool.Ok(data)
	return
}

// BaseCommPremmenu 接口请求参数
type BaseCommPermmenuReq struct {
	g.Meta        `path:"/permmenu" method:"GET"`
	Authorization string `json:"Authorization" in:"header"`
}

// BaseCommPremmenu 方法
func (c *BaseCommController) Permmenu(ctx context.Context, req *BaseCommPermmenuReq) (res *cool.BaseRes, err error) {

	var (
		baseSysPermsService = service.NewBaseSysPermsService()
		admin               = cool.GetAdmin(ctx)
	)
	res = cool.Ok(baseSysPermsService.Permmenu(ctx, admin.RoleIds))
	return
}

// BaseCommLogoutReq
type BaseCommLogoutReq struct {
	g.Meta        `path:"/logout" method:"POST"`
	Authorization string `json:"Authorization" in:"header"`
}

// BaseCommLogout 方法
func (c *BaseCommController) Logout(ctx context.Context, req *BaseCommLogoutReq) (res *cool.BaseRes, err error) {
	var (
		BaseSysLoginService = service.NewBaseSysLoginService()
	)
	err = BaseSysLoginService.Logout(ctx)
	res = cool.Ok(nil)
	return
}

// BaseCommUploadModeReq
type BaseCommUploadModeReq struct {
	g.Meta        `path:"/uploadMode" method:"GET"`
	Authorization string `json:"Authorization" in:"header"`
}

// BaseCommUploadMode 方法
func (c *BaseCommController) UploadMode(ctx context.Context, req *BaseCommUploadModeReq) (res *cool.BaseRes, err error) {
	data, err := cool.File().GetMode()
	res = cool.Ok(data)
	return
}

// BaseCommUploadReq
type BaseCommUploadReq struct {
	g.Meta        `path:"/upload" method:"POST"`
	Authorization string `json:"Authorization" in:"header"`
}

// BaseCommUpload 方法
func (c *BaseCommController) Upload(ctx context.Context, req *BaseCommUploadReq) (res *cool.BaseRes, err error) {
	data, err := cool.File().Upload(ctx)
	res = cool.Ok(data)
	return
}
