package cool

import (
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type IService interface {
	ServiceAdd(ctx context.Context, req *AddReq) (data interface{}, err error)
	ServiceDelete(ctx context.Context, req *DeleteReq) (data interface{}, err error)
	ServiceUpdate(ctx context.Context, req *UpdateReq) (data interface{}, err error)
	ServiceInfo(ctx context.Context, req *InfoReq) (data interface{}, err error)
	ServiceList(ctx context.Context, req *ListReq) (data interface{}, err error)
	ServicePage(ctx context.Context, req *PageReq) (data interface{}, err error)
	// 新增|删除|修改后的操作
	ModifyAfter(ctx context.Context, param g.MapStrAny) (err error)
}
type Service struct {
	Model       IModel
	ListQueryOp *ListQueryOp
	PageQueryOp *PageQueryOp
	InsertParam InsertParam // Add时插入参数
}

// List接口条件配置
type ListQueryOp struct {
	FieldEQ      []string    // 字段等于
	KeyWorkField []string    // 模糊搜索匹配的数据库字段
	AddOrderby   g.MapStrStr // 添加排序
	Where        Where       // 自定义条件
}

// Add接口条件配置
type PageQueryOp struct {
	FieldEQ      []string    // 字段等于
	KeyWorkField []string    // 模糊搜索匹配的数据库字段
	AddOrderby   g.MapStrStr // 添加排序
	Where        Where       // 添加条件
}

// Add时增加的参数
type InsertParam func(ctx context.Context) g.MapStrAny

// 查询时的where条件
type Where func(ctx context.Context) []g.Array

func (s *Service) ServiceAdd(ctx context.Context, req *AddReq) (data interface{}, err error) {
	r := g.RequestFromCtx(ctx)
	rjson, _ := r.GetJson()
	// 如果rjson为空 则直接返回
	if rjson == nil {
		return nil, nil
	}
	insertParams := s.InsertParam(ctx)
	if insertParams != nil {
		for k, v := range insertParams {
			rjson.Set(k, v)
		}
	}
	m := GDBM(s.Model)
	lastInsertId, err := m.Data(rjson).InsertAndGetId()
	if err != nil {
		return
	}
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
	m := GDBM(s.Model)
	_, err = m.Data(rmap).Where("id", rmap["id"]).Update()

	return
}

func (s *Service) ServiceInfo(ctx context.Context, req *InfoReq) (data interface{}, err error) {
	m := g.DB(s.Model.GroupName()).Model(s.Model.TableName())
	data, err = m.Clone().Where("id", req.Id).One()
	return
}

func (s *Service) ServiceList(ctx context.Context, req *ListReq) (data interface{}, err error) {
	r := g.RequestFromCtx(ctx)
	m := g.DB(s.Model.GroupName()).Model(s.Model.TableName())

	// 如果 req.Order 和 req.Sort 均不为空 则添加排序
	if !r.Get("order").IsEmpty() && !r.Get("sort").IsEmpty() {
		m.Order(r.Get("order").String() + " " + r.Get("sort").String())
		// m.OrderDesc("orderNum")
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
						sql = "(`" + field + "` LIKE ?) "
					} else {
						sql = sql + " OR " + "(`" + field + "` LIKE ?) "
					}
				}
				m.Where(sql, args.Slice())
			}
		}
		if s.ListQueryOp.Where != nil {
			where := s.ListQueryOp.Where(ctx)
			if len(where) > 0 {
				for _, v := range where {
					if gconv.Bool(v[2]) {
						m.Where(v[0], v[1])
					}
				}
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
	var (
		r     = g.RequestFromCtx(ctx)
		total = 0
	)

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

	// 如果pageQueryOp不为空 则使用pageQueryOp进行查询
	if s.PageQueryOp != nil {
		// 如果fileldEQ不为空 则添加查询条件
		if len(s.PageQueryOp.FieldEQ) > 0 {
			for _, field := range s.PageQueryOp.FieldEQ {
				g.Log().Debug(ctx, field)
				if !r.Get(field).IsEmpty() {
					m.Where(field, r.Get(field))
				}
			}
		}
		// 如果keyWorkField不为空 则添加查询条件
		if !r.Get("keyWord").IsEmpty() {
			if len(s.PageQueryOp.KeyWorkField) > 0 {
				var sql string
				args := garray.NewArray()
				for i, field := range s.PageQueryOp.KeyWorkField {
					args.Append("%" + r.Get("keyWord").String() + "%")
					if i == 0 {
						sql = "(`" + field + "` LIKE ?) "
					} else {
						sql = sql + " OR " + "(`" + field + "` LIKE ?) "
					}
				}
				m.Where(sql, args.Slice())
			}
		}
		// 加入where条件
		if s.PageQueryOp.Where != nil {
			where := s.PageQueryOp.Where(ctx)
			if len(where) > 0 {
				for _, v := range where {
					if gconv.Bool(v[2]) {
						m.Where(v[0], v[1])
					}
				}
			}
		}
		// 统计总数
		total, err = m.Clone().Count()
		if err != nil {
			return nil, err
		}
		// 如果 addOrderby 不为空 则添加排序
		if len(s.PageQueryOp.AddOrderby) > 0 && r.Get("order").IsEmpty() && r.Get("sort").IsEmpty() {
			for field, order := range s.PageQueryOp.AddOrderby {
				m.Order(field, order)
			}
		}
	}

	// 如果 req.Order 和 req.Sort 均不为空 则添加排序
	if !r.Get("order").IsEmpty() && !r.Get("sort").IsEmpty() {
		m.Order(r.Get("order").String() + " " + r.Get("sort").String())
	}

	result, err := m.Offset((req.Page - 1) * req.Size).Limit(req.Size).All()
	if err != nil {
		return nil, err
	}
	if result != nil {
		data = g.Map{
			"list": result,
			"pagination": pagination{
				Page:  req.Page,
				Size:  req.Size,
				Total: total,
			},
		}
	} else {
		data = g.Map{
			"list": garray.New(),
			"pagination": pagination{
				Page:  req.Page,
				Size:  req.Size,
				Total: total,
			},
		}
	}
	return
}

// 新增|删除|修改后的操作
func (s *Service) ModifyAfter(ctx context.Context, param g.MapStrAny) (err error) {
	g.Log().Debug(ctx, param, "2222222")

	return
}

// NewService 新建一个service
func NewService(model IModel) *Service {
	return &Service{
		Model: model,
	}
}
