package admin

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/service"
)

type BaseSysRoleController struct {
	*cool.Controller
}

func init() {
	var base_sys_role_controller = &BaseSysRoleController{
		&cool.Controller{
			Prefix:  "/admin/base/sys/role",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewBaseSysRoleService(),
		},
	}
	// 注册路由
	cool.RegisterController(base_sys_role_controller)
}
