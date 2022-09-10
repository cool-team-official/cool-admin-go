package admin

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/service"
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
