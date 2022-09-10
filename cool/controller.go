package cool

import (
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type IController interface {
	Add(ctx context.Context, req *AddReq) (res *BaseRes, err error)
	Delete(ctx context.Context, req *DeleteReq) (res *BaseRes, err error)
	Update(ctx context.Context, req *UpdateReq) (res *BaseRes, err error)
	Info(ctx context.Context, req *InfoReq) (res *BaseRes, err error)
	List(ctx context.Context, req *ListReq) (res *BaseRes, err error)
	Page(ctx context.Context, req *PageReq) (res *BaseRes, err error)
}
type Controller struct {
	Perfix  string     `json:"perfix"`
	Api     g.ArrayStr `json:"api"`
	Service IService   `json:"service"`
}

type AddReq struct {
	g.Meta `path:"/add" method:"POST"`
}

type DeleteReq struct {
	g.Meta `path:"/delete" method:"POST"`
	Ids    []int `json:"ids" v:"required#请选择要删除的数据"`
}

type UpdateReq struct {
	g.Meta `path:"/update" method:"POST"`
}

type InfoReq struct {
	g.Meta `path:"/info" method:"GET"`
	Id     int `json:"id" v:"integer|required#请选择要查询的数据"`
}
type InfoRes struct {
	*BaseRes
	Data interface{} `json:"data"`
}
type ListReq struct {
	g.Meta `path:"/list" method:"POST"`
	Order  string `json:"order"`
	Sort   string `json:"sort"`
}

type PageReq struct {
	g.Meta `path:"/page" method:"POST"`
	Page   int    `d:"1" json:"page"`
	Size   int    `d:"15" json:"size"`
	Order  string `json:"order"`
	Sort   string `json:"sort"`
}

func (c *Controller) Add(ctx context.Context, req *AddReq) (res *BaseRes, err error) {
	if garray.NewStrArrayFrom(c.Api).Contains("Add") {
		data, err := c.Service.ServiceAdd(ctx, req)
		if err != nil {
			return Fail(err.Error()), err
		}
		c.Service.ModifyAfter(ctx, g.RequestFromCtx(ctx).GetMap())
		return Ok(data), err
	}
	g.RequestFromCtx(ctx).Response.Status = 404
	return nil, nil
}
func (c *Controller) Delete(ctx context.Context, req *DeleteReq) (res *BaseRes, err error) {
	if garray.NewStrArrayFrom(c.Api).Contains("Delete") {
		data, err := c.Service.ServiceDelete(ctx, req)
		if err != nil {
			return Fail(err.Error()), err
		}
		c.Service.ModifyAfter(ctx, g.RequestFromCtx(ctx).GetMap())
		return Ok(data), err
	}
	g.RequestFromCtx(ctx).Response.Status = 404
	return nil, nil
}
func (c *Controller) Update(ctx context.Context, req *UpdateReq) (res *BaseRes, err error) {
	if garray.NewStrArrayFrom(c.Api).Contains("Update") {
		data, err := c.Service.ServiceUpdate(ctx, req)
		if err != nil {
			return Fail(err.Error()), err
		}
		c.Service.ModifyAfter(ctx, g.RequestFromCtx(ctx).GetMap())
		return Ok(data), err
	}
	g.RequestFromCtx(ctx).Response.Status = 404
	return nil, nil
}
func (c *Controller) Info(ctx context.Context, req *InfoReq) (res *BaseRes, err error) {
	if garray.NewStrArrayFrom(c.Api).Contains("Info") {
		data, err := c.Service.ServiceInfo(ctx, req)
		return Ok(data), err
	}
	g.RequestFromCtx(ctx).Response.Status = 404
	return nil, nil
}
func (c *Controller) List(ctx context.Context, req *ListReq) (res *BaseRes, err error) {
	if garray.NewStrArrayFrom(c.Api).Contains("List") {
		data, err := c.Service.ServiceList(ctx, req)
		return Ok(data), err
	}
	g.RequestFromCtx(ctx).Response.Status = 404
	return nil, nil
}
func (c *Controller) Page(ctx context.Context, req *PageReq) (res *BaseRes, err error) {
	if garray.NewStrArrayFrom(c.Api).Contains("Page") {
		data, err := c.Service.ServicePage(ctx, req)
		return Ok(data), err
	}
	g.RequestFromCtx(ctx).Response.Status = 404
	return nil, nil
}

// 注册控制器到路由
func RegisterController(c IController) {
	var ctx g.Ctx
	var sController = &Controller{}
	var sService = &Service{}
	gconv.Struct(c, &sController)
	gconv.Struct(sController.Service, &sService)
	// fields, _ := gstructs.Fields(gstructs.FieldsInput{
	// 	Pointer:         sService.Model,
	// 	RecursiveOption: 1,
	// })
	// g.Dump(fields)
	g.Log().Debug(ctx, "RegisterController", c, sController, sService.Model.TableName())
	// g.Server().BindMiddlewareDefault()
	g.Server().Group(
		sController.Perfix, func(group *ghttp.RouterGroup) {
			group.Middleware(MiddlewareHandlerResponse)
			group.Bind(
				c,
			)
		})

}
