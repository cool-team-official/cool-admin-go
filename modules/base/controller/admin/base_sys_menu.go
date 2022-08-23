package admin

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/service"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type BaseSysMenuController struct {
	*cool.Controller
}

func init() {
	var base_sys_menu_controller = &BaseSysMenuController{
		&cool.Controller{
			Perfix:  "/admin/base/sys/menu",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewBaseSysMenuService(),
		},
	}
	// 注册路由
	cool.RegisterController(base_sys_menu_controller)
}

// 增加 Welcome 演示 方法
type BaseSysMenuWelcomeReq struct {
	g.Meta `path:"/welcome" method:"GET"`
}
type BaseSysMenuWelcomeRes struct {
	*cool.BaseRes
	Data interface{} `json:"data"`
}

func (c *BaseSysMenuController) Welcome(ctx context.Context, req *BaseSysMenuWelcomeReq) (res *BaseSysMenuWelcomeRes, err error) {
	res = &BaseSysMenuWelcomeRes{
		BaseRes: cool.Ok("Welcome to Cool Admin Go"),
		Data:    gjson.New(`{"name": "Cool Admin Go", "age":0}`),
	}
	return
}
