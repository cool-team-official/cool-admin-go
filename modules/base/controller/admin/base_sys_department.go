package admin

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/service"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type BaseSysDepartmentController struct {
	*cool.Controller
}

func init() {
	var base_sys_department_controller = &BaseSysDepartmentController{
		&cool.Controller{
			Perfix:  "/admin/base/sys/department",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewBaseSysDepartmentService(),
		},
	}
	// 注册路由
	cool.RegisterController(base_sys_department_controller)
}

// 增加 Welcome 演示 方法
type BaseSysDepartmentWelcomeReq struct {
	g.Meta `path:"/welcome" method:"GET"`
}
type BaseSysDepartmentWelcomeRes struct {
	*cool.BaseRes
	Data interface{} `json:"data"`
}

func (c *BaseSysDepartmentController) Welcome(ctx context.Context, req *BaseSysDepartmentWelcomeReq) (res *BaseSysDepartmentWelcomeRes, err error) {
	res = &BaseSysDepartmentWelcomeRes{
		BaseRes: cool.Ok("Welcome to Cool Admin Go"),
		Data:    gjson.New(`{"name": "Cool Admin Go", "age":0}`),
	}
	return
}
