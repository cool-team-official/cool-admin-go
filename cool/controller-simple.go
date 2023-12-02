package cool

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type IControllerSimple interface {
}
type ControllerSimple struct {
	Prefix string
}

// 注册不带crud的路由
func RegisterControllerSimple(c IControllerSimple) {
	var sController = &ControllerSimple{}
	// var sService = &Service{}
	gconv.Struct(c, &sController)
	g.Server().Group(
		sController.Prefix, func(group *ghttp.RouterGroup) {
			group.Middleware(MiddlewareHandlerResponse)
			group.Bind(
				c,
			)
		})
}
