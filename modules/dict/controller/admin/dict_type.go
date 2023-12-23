package admin

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/dict/service"
)

type DictTypeController struct {
	*cool.Controller
}

func init() {
	var dict_type_controller = &DictTypeController{
		&cool.Controller{
			Prefix:  "/admin/dict/type",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewDictTypeService(),
		},
	}
	// 注册路由
	cool.RegisterController(dict_type_controller)
}
