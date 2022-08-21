package cool

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gmeta"
)

type IService interface {
	ServiceAdd(ctx context.Context, req *AddReq) (res *AddRes, err error)
	ServiceDelete(ctx context.Context, req *DeleteReq) (res *DeleteRes, err error)
	ServiceUpdate(ctx context.Context, req *UpdateReq) (res *UpdateRes, err error)
	ServiceInfo(ctx context.Context, req *InfoReq) (data interface{}, err error)
	ServiceList(ctx context.Context, req *ListReq) (data interface{}, err error)
	ServicePage(ctx context.Context, req *PageReq) (res *PageRes, err error)
}
type Service struct {
	Model IModel
}

func (s *Service) ServiceAdd(ctx context.Context, req *AddReq) (res *AddRes, err error) {
	g.Log().Debug(ctx, "Cool Add service~~~~~~~~~~")
	g.Dump(s.Model)
	meta := gmeta.Data(s.Model)
	g.Dump(meta)
	request := g.RequestFromCtx(ctx)
	// g.Dump("request:", request)
	rjson, _ := request.GetJson()
	g.Dump("rjson:", rjson)

	// 如果rjson为空 则直接返回
	if rjson == nil {
		return nil, nil
	}

	m := g.DB(s.Model.GroupName()).Model(s.Model.TableName())
	result, err := m.Clone().Data(rjson).Insert()
	if err != nil {
		return nil, err
	}
	g.Dump(result)
	return &AddRes{Data: "Cool Add"}, nil
}

func (s *Service) ServiceDelete(ctx context.Context, req *DeleteReq) (res *DeleteRes, err error) {
	return &DeleteRes{Data: "Cool Delete"}, nil
}

func (s *Service) ServiceUpdate(ctx context.Context, req *UpdateReq) (res *UpdateRes, err error) {
	return &UpdateRes{Data: "Cool Update"}, nil
}

func (s *Service) ServiceInfo(ctx context.Context, req *InfoReq) (data interface{}, err error) {
	m := g.DB(s.Model.GroupName()).Model(s.Model.TableName())
	data, err = m.Clone().Where("id", req.ID).One()
	return
}

func (s *Service) ServiceList(ctx context.Context, req *ListReq) (data interface{}, err error) {
	m := g.DB(s.Model.GroupName()).Model(s.Model.TableName())
	result, err := m.All()
	data = result
	return
}

func (s *Service) ServicePage(ctx context.Context, req *PageReq) (res *PageRes, err error) {
	return &PageRes{Data: "Cool Page"}, nil
}
func NewService(model IModel) *Service {
	return &Service{
		Model: model,
	}
}
