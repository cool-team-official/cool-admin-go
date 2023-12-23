package admin

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/service"
)

type BaseSysMenuController struct {
	*cool.Controller
}

func init() {
	var base_sys_menu_controller = &BaseSysMenuController{
		&cool.Controller{
			Prefix:  "/admin/base/sys/menu",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewBaseSysMenuService(),
		},
	}
	// 注册路由
	cool.RegisterController(base_sys_menu_controller)
}
