package admin

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/service"
)

type BaseSysLogController struct {
	*cool.Controller
}

func init() {
	var base_sys_log_controller = &BaseSysLogController{
		&cool.Controller{
			Perfix:  "/admin/base/sys/log",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewBaseSysLogService(),
		},
	}
	// 注册路由
	cool.RegisterController(base_sys_log_controller)
}
