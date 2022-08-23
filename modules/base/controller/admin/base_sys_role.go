package admin

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/service"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type BaseSysRoleController struct {
	*cool.Controller
}

func init() {
	var base_sys_role_controller = &BaseSysRoleController{
		&cool.Controller{
			Perfix:  "/admin/base/sys/role",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewBaseSysRoleService(),
		},
	}
	// 注册路由
	cool.RegisterController(base_sys_role_controller)
}

// 增加 Welcome 演示 方法
type BaseSysRoleWelcomeReq struct {
	g.Meta `path:"/welcome" method:"GET"`
}
type BaseSysRoleWelcomeRes struct {
	*cool.BaseRes
	Data interface{} `json:"data"`
}

func (c *BaseSysRoleController) Welcome(ctx context.Context, req *BaseSysRoleWelcomeReq) (res *BaseSysRoleWelcomeRes, err error) {
	res = &BaseSysRoleWelcomeRes{
		BaseRes: cool.Ok("Welcome to Cool Admin Go"),
		Data:    gjson.New(`{"name": "Cool Admin Go", "age":0}`),
	}
	return
}
