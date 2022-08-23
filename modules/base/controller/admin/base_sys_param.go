package admin

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/service"
)

type BaseSysParamController struct {
	*cool.Controller
}

func init() {
	var base_sys_param_controller = &BaseSysParamController{
		&cool.Controller{
			Perfix:  "/admin/base/sys/param",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewBaseSysParamService(),
		},
	}
	// 注册路由
	cool.RegisterController(base_sys_param_controller)
}
