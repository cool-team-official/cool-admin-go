package admin

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/service"
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
