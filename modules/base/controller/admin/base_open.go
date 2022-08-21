package admin

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	v1 "github.com/cool-team-official/cool-admin-go/modules/base/api/v1"
	"github.com/cool-team-official/cool-admin-go/modules/base/service"
	"github.com/gogf/gf/v2/frame/g"
)

type BaseOpen struct {
}

func init() {
	var open = &BaseOpen{}
	// 注册路由
	cool.RegisterControllerSimple("/admin/base/open", open)
}

// 验证码接口
func (c *BaseOpen) BaseOpenCaptcha(ctx context.Context, req *v1.BaseOpenCaptchaReq) (res *cool.BaseRes, err error) {
	var (
		baseSysLoginService = service.NewBaseSysLoginService()
	)
	data, err := baseSysLoginService.Captcha(req)
	res = cool.Ok(data)
	return
}

// eps 接口请求
type BaseOpenEpsReq struct {
	g.Meta `path:"/eps" method:"GET"`
}

// eps 接口
func (c *BaseOpen) Eps(ctx context.Context, req *BaseOpenEpsReq) (res *cool.BaseRes, err error) {
	res = cool.Ok(service.AdminEps())
	return
}

// login 接口
func (c *BaseOpen) Login(ctx context.Context, req *v1.BaseOpenLoginReq) (res *cool.BaseRes, err error) {
	var (
		baseSysLoginService = service.NewBaseSysLoginService()
	)
	data, err := baseSysLoginService.Login(ctx, req)
	if err != nil {
		return
	}
	res = cool.Ok(data)
	return
}
