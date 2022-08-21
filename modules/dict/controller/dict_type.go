package admin

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/dict/service"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type DictTypeController struct {
	*cool.Controller
}

func init() {
	var dict_type_controller = &DictTypeController{
		&cool.Controller{
			Perfix:  "/admin/dict/type",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewDictTypeService(),
		},
	}
	// 注册路由
	cool.RegisterController(dict_type_controller)
}

// 增加 Welcome 演示 方法
type DictTypeWelcomeReq struct {
	g.Meta `path:"/welcome" method:"GET"`
}
type DictTypeWelcomeRes struct {
	*cool.BaseRes
	Data interface{} `json:"data"`
}

func (c *DictTypeController) Welcome(ctx context.Context, req *DictTypeWelcomeReq) (res *DictTypeWelcomeRes, err error) {
	res = &DictTypeWelcomeRes{
		BaseRes: cool.Ok("Welcome to Cool Admin Go"),
		Data:    gjson.New(`{"name": "Cool Admin Go", "age":0}`),
	}
	return
}
