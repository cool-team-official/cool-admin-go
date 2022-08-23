package cool

import (
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type IService interface {
	ServiceAdd(ctx context.Context, req *AddReq) (data interface{}, err error)
	ServiceDelete(ctx context.Context, req *DeleteReq) (data interface{}, err error)
	ServiceUpdate(ctx context.Context, req *UpdateReq) (data interface{}, err error)
	ServiceInfo(ctx context.Context, req *InfoReq) (data interface{}, err error)
	ServiceList(ctx context.Context, req *ListReq) (data interface{}, err error)
	ServicePage(ctx context.Context, req *PageReq) (data interface{}, err error)
}
type Service struct {
	Model       IModel
	ListQueryOp *ListQueryOp
}

// List接口条件配置
type ListQueryOp struct {
	FieldEQ      []string    // 字段等于
	KeyWorkField []string    // 模糊搜索匹配的数据库字段
	AddOrderby   g.MapStrStr // 添加排序
}

func (s *Service) ServiceAdd(ctx context.Context, req *AddReq) (data interface{}, err error) {
	r := g.RequestFromCtx(ctx)
	rjson, _ := r.GetJson()
	// 如果rjson为空 则直接返回
	if rjson == nil {
		return nil, nil
	}
	m := GDBModel(s.Model)
	lastInsertId, err := m.Data(rjson).InsertAndGetId()
	data = g.Map{"id": lastInsertId}
	return
}
func (s *Service) ServiceDelete(ctx context.Context, req *DeleteReq) (data interface{}, err error) {
	m := g.DB(s.Model.GroupName()).Model(s.Model.TableName())
	data, err = m.WhereIn("id", req.Ids).Delete()
	return
}

func (s *Service) ServiceUpdate(ctx context.Context, req *UpdateReq) (data interface{}, err error) {
	r := g.RequestFromCtx(ctx)
	rmap := r.GetMap()
	// g.Dump(rmap)
	if rmap["id"] == nil {
		err = gerror.New("id不能为空")
		return
	}
	m := GDBModel(s.Model)
	_, err = m.Data(rmap).Where("id", rmap["id"]).Update()
	return
}

func (s *Service) ServiceInfo(ctx context.Context, req *InfoReq) (data interface{}, err error) {
	m := g.DB(s.Model.GroupName()).Model(s.Model.TableName())
	data, err = m.Clone().Where("id", req.ID).One()
	return
}

func (s *Service) ServiceList(ctx context.Context, req *ListReq) (data interface{}, err error) {
	r := g.RequestFromCtx(ctx)
	m := g.DB(s.Model.GroupName()).Model(s.Model.TableName())

	// 如果 req.Order 和 req.Sort 均不为空 则添加排序
	if !r.Get("order").IsEmpty() && !r.Get("sort").IsEmpty() {
		// m.Order(r.Get("order").String() + " " + r.Get("sort").String())
		m.OrderDesc("orderNum")
	}
	// 如果 ListQueryOp 不为空 则使用 ListQueryOp 进行查询
	if s.ListQueryOp != nil {
		// 如果fileldEQ不为空 则添加查询条件
		if len(s.ListQueryOp.FieldEQ) > 0 {
			for _, field := range s.ListQueryOp.FieldEQ {
				g.Log().Debug(ctx, field)
				if !r.Get(field).IsEmpty() {
					m.Where(field, r.Get(field))
				}
			}
		}
		// 如果keyWorkField不为空 则添加查询条件
		if !r.Get("keyWord").IsEmpty() {
			if len(s.ListQueryOp.KeyWorkField) > 0 {
				var sql string
				args := garray.NewArray()
				for i, field := range s.ListQueryOp.KeyWorkField {
					args.Append("%" + r.Get("keyWord").String() + "%")
					if i == 0 {
						// ((`typeId` LIKE '%cool%')
						sql = "(`" + field + "` LIKE ?) "
					} else {
						sql = sql + " OR " + "(`" + field + "` LIKE ?) "
					}
				}
				m.Where(sql, args.Slice())
			}
		}

		// 如果 addOrderby 不为空 则添加排序
		if len(s.ListQueryOp.AddOrderby) > 0 && r.Get("order").IsEmpty() && r.Get("sort").IsEmpty() {
			for field, order := range s.ListQueryOp.AddOrderby {
				m.Order(field, order)
			}
		}
	}

	result, err := m.All()
	// g.Dump(result)
	if err != nil {
		g.Log().Error(ctx, "ServiceList error:", err)
	}
	if result == nil {
		data = garray.New()
	} else {
		data = result
	}
	return
}

func (s *Service) ServicePage(ctx context.Context, req *PageReq) (data interface{}, err error) {
	type pagination struct {
		Page  int `json:"page"`
		Size  int `json:"size"`
		Total int `json:"total"`
	}
	if req.Size <= 0 {
		req.Size = 10
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	m := g.DB(s.Model.GroupName()).Model(s.Model.TableName())
	total, err := m.Clone().Count()
	if err != nil {
		return nil, err
	}
	list, err := m.Offset((req.Page - 1) * req.Size).Limit(req.Size).All()
	if err != nil {
		return nil, err
	}
	data = g.Map{
		"list": list,
		"pagination": pagination{
			Page:  req.Page,
			Size:  req.Size,
			Total: total,
		},
	}
	return
}
func NewService(model IModel) *Service {
	return &Service{
		Model: model,
	}
}
