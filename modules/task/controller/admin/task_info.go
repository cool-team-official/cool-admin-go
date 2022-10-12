package admin

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/task/service"
)

type TaskInfoController struct {
	*cool.Controller
}

func init() {
	var task_info_controller = &TaskInfoController{
		&cool.Controller{
			Perfix:  "/admin/task/info",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewTaskInfoService(),
		},
	}
	// 注册路由
	cool.RegisterController(task_info_controller)
}
