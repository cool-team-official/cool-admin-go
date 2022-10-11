package middleware

import "github.com/gogf/gf/v2/frame/g"

func init() {
	g.Server().BindMiddleware("/admin/*/open/*", BaseAuthorityMiddlewareOpen)
	g.Server().BindMiddleware("/admin/*/comm/*", BaseAuthorityMiddlewareComm)
	g.Server().BindMiddleware("/admin/*", BaseAuthorityMiddleware)
	g.Server().BindMiddleware("/admin/*", BaseLog)

}
