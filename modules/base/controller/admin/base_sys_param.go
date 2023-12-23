package admin

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type BaseSysParamController struct {
	*cool.Controller
}

func init() {
	var base_sys_param_controller = &BaseSysParamController{
		&cool.Controller{
			Prefix:  "/admin/base/sys/param",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewBaseSysParamService(),
		},
	}
	// 注册路由
	cool.RegisterController(base_sys_param_controller)
}

// BaseSysParamHtmlReq 请求参数
type BaseSysParamHtmlReq struct {
	g.Meta `path:"/html" method:"GET"`
	Key    string `v:"required#请输入key"`
}

// Html 根据配置参数key获取网页内容(富文本)
func (c *BaseSysParamController) Html(ctx g.Ctx, req *BaseSysParamHtmlReq) (res *cool.BaseRes, err error) {
	var (
		BaseSysParamService = service.NewBaseSysParamService()
		r                   = ghttp.RequestFromCtx(ctx)
	)
	r.Response.WriteExit(BaseSysParamService.HtmlByKey(req.Key))
	return
}
