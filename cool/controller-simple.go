package cool

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type IControllerSimple interface {
}
type ControllerSimple struct {
	Perfix string
}

// 注册不带crud的路由
func RegisterControllerSimple(perfix string, controller IControllerSimple) {
	g.Server().Group(
		perfix, func(group *ghttp.RouterGroup) {
			group.Middleware(MiddlewareHandlerResponse)
			group.Bind(
				controller,
			)
		})
}
