package admin

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/dict/service"

	"github.com/gogf/gf/v2/frame/g"
)

type DictInfoController struct {
	*cool.Controller
}

func init() {
	var dict_info_controller = &DictInfoController{
		&cool.Controller{
			Perfix:  "/admin/dict/info",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewDictInfoService(),
		},
	}
	// 注册路由
	cool.RegisterController(dict_info_controller)
}

// Data 方法请求
type DictInfoDataReq struct {
	g.Meta `path:"/data" method:"POST"`
	Types  []string `json:"types"`
}

// Data 方法 获得字典数据
func (c *DictInfoController) Data(ctx context.Context, req *DictInfoDataReq) (res *cool.BaseRes, err error) {
	res = cool.Ok(g.Map{})
	return
}
