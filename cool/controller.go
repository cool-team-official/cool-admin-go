package cool

import (
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type IController interface {
	// Add(ctx context.Context, req *AddReq) (res *AddRes, err error)
	Delete(ctx context.Context, req *DeleteReq) (res *DeleteRes, err error)
	Update(ctx context.Context, req *UpdateReq) (res *UpdateRes, err error)
	Info(ctx context.Context, req *InfoReq) (res *InfoRes, err error)
	List(ctx context.Context, req *ListReq) (res *ListRes, err error)
	Page(ctx context.Context, req *PageReq) (res *PageRes, err error)
}
type Controller struct {
	Perfix  string
	Api     g.ArrayStr
	Service IService
}
type AddReq struct {
	g.Meta `path:"/add" method:"POST"`
}
type AddRes struct {
	Data interface{} `json:"data"`
}
type DeleteReq struct {
	g.Meta `path:"/delete" method:"POST"`
}
type DeleteRes struct {
	Data interface{} `json:"data"`
}
type UpdateReq struct {
	g.Meta `path:"/update" method:"POST"`
}
type UpdateRes struct {
	Data interface{} `json:"data"`
}
type InfoReq struct {
	g.Meta `path:"/info" method:"GET"`
	ID     int `param:"id"`
}
type InfoRes struct {
	Data interface{} `json:"data"`
}
type ListReq struct {
	g.Meta `path:"/list" method:"POST"`
}
type ListRes struct {
	Data interface{} `json:"data"`
}
type PageReq struct {
	g.Meta `path:"/page" method:"POST"`
}
type PageRes struct {
	Data interface{} `json:"data"`
}

func (c *Controller) Add(ctx context.Context, req *AddReq) (res *AddRes, err error) {
	g.Log().Debug(ctx, "Cool Add controller~~~~~~~~~~")
	if garray.NewStrArrayFrom(c.Api).Contains("add") {
		return c.Service.ServiceAdd(ctx, req)
	}
	g.RequestFromCtx(ctx).Response.Status = 404
	return nil, nil
}
func (c *Controller) Delete(ctx context.Context, req *DeleteReq) (res *DeleteRes, err error) {
	if garray.NewStrArrayFrom(c.Api).Contains("delete") {
		return c.Service.ServiceDelete(ctx, req)
	}
	g.RequestFromCtx(ctx).Response.Status = 404
	return nil, nil
}
func (c *Controller) Update(ctx context.Context, req *UpdateReq) (res *UpdateRes, err error) {
	if garray.NewStrArrayFrom(c.Api).Contains("update") {
		return c.Service.ServiceUpdate(ctx, req)
	}
	g.RequestFromCtx(ctx).Response.Status = 404
	return nil, nil
}
func (c *Controller) Info(ctx context.Context, req *InfoReq) (res *InfoRes, err error) {
	if garray.NewStrArrayFrom(c.Api).Contains("info") {
		return c.Service.ServiceInfo(ctx, req)
	}
	g.RequestFromCtx(ctx).Response.Status = 404
	return nil, nil
}
func (c *Controller) List(ctx context.Context, req *ListReq) (res *ListRes, err error) {
	if garray.NewStrArrayFrom(c.Api).Contains("list") {
		return c.Service.ServiceList(ctx, req)
	}
	g.RequestFromCtx(ctx).Response.Status = 404
	return nil, nil
}
func (c *Controller) Page(ctx context.Context, req *PageReq) (res *PageRes, err error) {
	if garray.NewStrArrayFrom(c.Api).Contains("page") {
		return c.Service.ServicePage(ctx, req)
	}
	g.RequestFromCtx(ctx).Response.Status = 404
	return nil, nil
}

// func NewController(model interface{}) *Controller {
// 	return &Controller{
// 		Service: NewService(model),
// 	}
// }

// 注册控制器到路由
func RegisterController(c IController) {
	var sController = &Controller{}
	gconv.Struct(c, &sController)
	g.Server().BindMiddlewareDefault()
	g.Server().Group(
		sController.Perfix, func(group *ghttp.RouterGroup) {
			group.Middleware(MiddlewareHandlerResponse)
			group.Bind(
				c,
			)
		})

}
