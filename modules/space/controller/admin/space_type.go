package admin

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/space/service"
)

type SpaceTypeController struct {
	*cool.Controller
}

func init() {
	var space_type_controller = &SpaceTypeController{
		&cool.Controller{
			Perfix:  "/admin/space/type",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewSpaceTypeService(),
		},
	}
	// 注册路由
	cool.RegisterController(space_type_controller)
}
