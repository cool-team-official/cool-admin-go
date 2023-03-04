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
		&cool.ControllerSimple{
			Perfix: "/app/base/comm",
			//    Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			//    Service: service.NewBaseCommService(),
		},
	}
	// 注册路由
	cool.RegisterControllerSimple(base_comm_controller)
}

// eps 接口请求
type BaseCommControllerEpsReq struct {
	g.Meta `path:"/eps" method:"GET"`
}

// eps 接口
func (c *BaseCommController) Eps(ctx context.Context, req *BaseCommControllerEpsReq) (res *cool.BaseRes, err error) {
	if !cool.Config.Eps {
		g.Log().Error(ctx, "eps is not open")
		res = cool.Ok(nil)
		return
	}
	baseOpenService := service.NewBaseOpenService()
	data, err := baseOpenService.AppEPS(ctx)
	if err != nil {
		g.Log().Error(ctx, "eps error", err)
		return cool.Fail(err.Error()), err
	}
	res = cool.Ok(data)
	return
}
