package middleware

import (
	"github.com/cool-team-official/cool-admin-go/modules/base/config"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	if config.Config.Middleware.Authority.Enable {
		g.Server().BindMiddleware("/admin/*/open/*", BaseAuthorityMiddlewareOpen)
		g.Server().BindMiddleware("/admin/*/comm/*", BaseAuthorityMiddlewareComm)
		g.Server().BindMiddleware("/admin/*", BaseAuthorityMiddleware)
	}
	if config.Config.Middleware.Log.Enable {
		g.Server().BindMiddleware("/admin/*", BaseLog)
	}

}
