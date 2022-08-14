package admin

import (
	"context"

	"cool-admin-go-simple/modules/demo/model"
	"cool-admin-go-simple/modules/demo/service"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/v2/frame/g"
)

type Sample struct {
	*cool.Controller
}

func init() {
	var sample = &Sample{
		&cool.Controller{
			Perfix:  "/admin/demo/sample",
			Api:     []string{"add", "delete", "update", "info", "list", "page"},
			Service: service.NewSample(model.NewSample()),
		},
	}
	// 注册路由
	cool.RegisterController(sample)
}

// 增加 Welcome 演示 方法
type SampleWelcomeReq struct {
	g.Meta `path:"/welcome" method:"GET"`
}
type SampleWelcomeRes struct {
	Data interface{} `json:"data"`
}

func (c *Sample) Welcome(ctx context.Context, req *SampleWelcomeReq) (res *SampleWelcomeRes, err error) {
	res = &SampleWelcomeRes{
		Data: "Welcome to Cool Admin Go!",
	}
	return
}
