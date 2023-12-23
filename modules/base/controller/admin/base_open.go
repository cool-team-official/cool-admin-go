package admin

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	v1 "github.com/cool-team-official/cool-admin-go/modules/base/api/v1"
	"github.com/cool-team-official/cool-admin-go/modules/base/service"
	"github.com/gogf/gf/v2/frame/g"
)

type BaseOpen struct {
	*cool.ControllerSimple
	baseSysLoginService *service.BaseSysLoginService
	baseOpenService     *service.BaseOpenService
}

func init() {
	var open = &BaseOpen{
		ControllerSimple:    &cool.ControllerSimple{Prefix: "/admin/base/open"},
		baseSysLoginService: service.NewBaseSysLoginService(),
		baseOpenService:     service.NewBaseOpenService(),
	}
	// 注册路由
	cool.RegisterControllerSimple(open)
}

// 验证码接口
func (c *BaseOpen) BaseOpenCaptcha(ctx context.Context, req *v1.BaseOpenCaptchaReq) (res *cool.BaseRes, err error) {
	data, err := c.baseSysLoginService.Captcha(req)
	res = cool.Ok(data)
	return
}

// eps 接口请求
type BaseOpenEpsReq struct {
	g.Meta `path:"/eps" method:"GET"`
}

// eps 接口
func (c *BaseOpen) Eps(ctx context.Context, req *BaseOpenEpsReq) (res *cool.BaseRes, err error) {
	if !cool.Config.Eps {
		g.Log().Error(ctx, "eps is not open")
		res = cool.Ok(nil)
		return
	}
	data, err := c.baseOpenService.AdminEPS(ctx)
	if err != nil {
		g.Log().Error(ctx, "eps error", err)
		return cool.Fail(err.Error()), err
	}
	res = cool.Ok(data)
	return
}

// login 接口
func (c *BaseOpen) Login(ctx context.Context, req *v1.BaseOpenLoginReq) (res *cool.BaseRes, err error) {
	data, err := c.baseSysLoginService.Login(ctx, req)
	if err != nil {
		return
	}
	res = cool.Ok(data)
	return
}

// RefreshTokenReq 刷新token请求
type RefreshTokenReq struct {
	g.Meta       `path:"/refreshToken" method:"GET"`
	RefreshToken string `json:"refreshToken" v:"required#refreshToken不能为空"`
}

// RefreshToken 刷新token
func (c *BaseOpen) RefreshToken(ctx context.Context, req *RefreshTokenReq) (res *cool.BaseRes, err error) {
	data, err := c.baseSysLoginService.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return
	}
	res = cool.Ok(data)
	return
}
