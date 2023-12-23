package admin

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/service"
	"github.com/gogf/gf/v2/frame/g"
)

type BaseSysDepartmentController struct {
	*cool.Controller
}

func init() {
	var base_sys_department_controller = &BaseSysDepartmentController{
		&cool.Controller{
			Prefix:  "/admin/base/sys/department",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewBaseSysDepartmentService(),
		},
	}
	// 注册路由
	cool.RegisterController(base_sys_department_controller)
}

// OrderReq 接口请求参数
type OrderReq struct {
	g.Meta        `path:"/order" method:"GET"`
	Authorization string `json:"Authorization" in:"header"`
}

// Order 排序部门
func (c *BaseSysDepartmentController) Order(ctx context.Context, req *OrderReq) (res *cool.BaseRes, err error) {
	err = service.NewBaseSysDepartmentService().Order(ctx)
	res = cool.Ok(nil)
	return
}
