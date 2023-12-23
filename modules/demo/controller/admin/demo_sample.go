package admin

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/demo/service"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type DemoSampleController struct {
	*cool.Controller
}

func init() {
	var demo_sample_controller = &DemoSampleController{
		&cool.Controller{
			Prefix:  "/admin/demo/demo_sample",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewDemoSampleService(),
		},
	}
	// 注册路由
	cool.RegisterController(demo_sample_controller)
}

// 增加 Welcome 演示 方法
type DemoSampleWelcomeReq struct {
	g.Meta `path:"/welcome" method:"GET"`
}
type DemoSampleWelcomeRes struct {
	*cool.BaseRes
	Data interface{} `json:"data"`
}

func (c *DemoSampleController) Welcome(ctx context.Context, req *DemoSampleWelcomeReq) (res *DemoSampleWelcomeRes, err error) {
	res = &DemoSampleWelcomeRes{
		BaseRes: cool.Ok("Welcome to Cool Admin Go"),
		Data:    gjson.New(`{"name": "Cool Admin Go", "age":0}`),
	}
	return
}
