package middleware

import (
	"github.com/cool-team-official/cool-admin-go/modules/base/service"
	"github.com/gogf/gf/v2/net/ghttp"
)

func BaseLog(r *ghttp.Request) {
	var (
		ctx               = r.GetCtx()
		BaseSysLogService = service.NewBaseSysLogService()
	)
	BaseSysLogService.Record(ctx)

	r.Middleware.Next()
}
