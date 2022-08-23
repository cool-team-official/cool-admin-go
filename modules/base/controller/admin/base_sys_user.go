package admin

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/service"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type BaseSysUserController struct {
	*cool.Controller
}

func init() {
	var base_sys_user_controller = &BaseSysUserController{
		&cool.Controller{
			Perfix:  "/admin/base/sys/user",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewBaseSysUserService(),
		},
	}
	// 注册路由
	cool.RegisterController(base_sys_user_controller)
}

// 增加 Welcome 演示 方法
type BaseSysUserWelcomeReq struct {
	g.Meta `path:"/welcome" method:"GET"`
}
type BaseSysUserWelcomeRes struct {
	*cool.BaseRes
	Data interface{} `json:"data"`
}

func (c *BaseSysUserController) Welcome(ctx context.Context, req *BaseSysUserWelcomeReq) (res *BaseSysUserWelcomeRes, err error) {
	res = &BaseSysUserWelcomeRes{
		BaseRes: cool.Ok("Welcome to Cool Admin Go"),
		Data:    gjson.New(`{"name": "Cool Admin Go", "age":0}`),
	}
	return
}
