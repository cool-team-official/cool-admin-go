package cool

import (
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type IService interface {
	ServiceAdd(ctx context.Context, req *AddReq) (data interface{}, err error)       // 新增
	ServiceDelete(ctx context.Context, req *DeleteReq) (data interface{}, err error) // 删除
	ServiceUpdate(ctx context.Context, req *UpdateReq) (data interface{}, err error) // 修改
	ServiceInfo(ctx context.Context, req *InfoReq) (data interface{}, err error)     // 详情
	ServiceList(ctx context.Context, req *ListReq) (data interface{}, err error)     // 列表
	ServicePage(ctx context.Context, req *PageReq) (data interface{}, err error)     // 分页
	ModifyAfter(ctx context.Context, param g.MapStrAny) (err error)                  // 新增|删除|修改后的操作
	GetModel() IModel                                                                // 获取model
}
type Service struct {
	Model              IModel
	ListQueryOp        *QueryOp
	PageQueryOp        *QueryOp
	InsertParam        func(ctx context.Context) g.MapStrAny // Add时插入参数
	Before             func(ctx context.Context) (err error) // CRUD前的操作
	InfoIgnoreProperty string                                // Info时忽略的字段,多个字段用逗号隔开
}

// List/Add接口条件配置
type QueryOp struct {
	FieldEQ      []string                                 // 字段等于
	KeyWordField []string                                 // 模糊搜索匹配的数据库字段
	AddOrderby   g.MapStrStr                              // 添加排序
	Where        func(ctx context.Context) []g.Array      // 自定义条件
	Select       string                                   // 查询字段,多个字段用逗号隔开 如: id,name  或  a.id,a.name,b.name AS bname
	Join         []*JoinOp                                // 关联查询
	Extend       func(ctx g.Ctx, m *gdb.Model) *gdb.Model // 追加其他条件
}

// JoinOp 关联查询
type JoinOp struct {
	Model     IModel   // 关联的model
	Alias     string   // 别名
	Condition string   // 关联条件
	Type      JoinType // 关联类型  LeftJoin RightJoin InnerJoin
}

// JoinType 关联类型
type JoinType string

func (s *Service) ServiceAdd(ctx context.Context, req *AddReq) (data interface{}, err error) {
	if s.Before != nil {
		err = s.Before(ctx)
		if err != nil {
			return
		}
	}
	r := g.RequestFromCtx(ctx)
	rjson, _ := r.GetJson()
	// 如果rjson为空 则直接返回
	if rjson == nil {
		return nil, nil
	}
	if s.InsertParam != nil {

		insertParams := s.InsertParam(ctx)
		if len(insertParams) > 0 {
			for k, v := range insertParams {
				rjson.Set(k, v)
			}
		}
	}
	m := DBM(s.Model)
	lastInsertId, err := m.Data(rjson).InsertAndGetId()
	if err != nil {
		return
	}
	data = g.Map{"id": lastInsertId}

	return
}
func (s *Service) ServiceDelete(ctx context.Context, req *DeleteReq) (data interface{}, err error) {
	if s.Before != nil {
		err = s.Before(ctx)
		if err != nil {
			return
		}
	}
	m := g.DB(s.Model.GroupName()).Model(s.Model.TableName())
	data, err = m.WhereIn("id", req.Ids).Delete()

	return
}

func (s *Service) ServiceUpdate(ctx context.Context, req *UpdateReq) (data interface{}, err error) {
	if s.Before != nil {
		err = s.Before(ctx)
		if err != nil {
			return
		}
	}
	r := g.RequestFromCtx(ctx)
	rmap := r.GetMap()
	// g.Dump(rmap)
	if rmap["id"] == nil {
		err = gerror.New("id不能为空")
		return
	}
	m := DBM(s.Model)
	_, err = m.Data(rmap).Where("id", rmap["id"]).Update()

	return
}

func (s *Service) ServiceInfo(ctx context.Context, req *InfoReq) (data interface{}, err error) {
	if s.Before != nil {
		err = s.Before(ctx)
		if err != nil {
			return
		}
	}
	m := g.DB(s.Model.GroupName()).Model(s.Model.TableName())
	// 如果InfoIgnoreProperty不为空 则忽略相关字段
	if len(s.InfoIgnoreProperty) > 0 {
		m = m.FieldsEx(s.InfoIgnoreProperty)
	}
	data, err = m.Clone().Where("id", req.Id).One()

	return
}

func (s *Service) ServiceList(ctx context.Context, req *ListReq) (data interface{}, err error) {
	if s.Before != nil {
		err = s.Before(ctx)
		if err != nil {
			return
		}
	}
	r := g.RequestFromCtx(ctx)
	m := g.DB(s.Model.GroupName()).Model(s.Model.TableName())

	// 如果 req.Order 和 req.Sort 均不为空 则添加排序
	if !r.Get("order").IsEmpty() && !r.Get("sort").IsEmpty() {
		m.Order(r.Get("order").String() + " " + r.Get("sort").String())
		// m.OrderDesc("orderNum")
	}
	// 如果 ListQueryOp 不为空 则使用 ListQueryOp 进行查询
	if s.ListQueryOp != nil {
		if Select := s.ListQueryOp.Select; Select != "" {
			m.Fields(Select)
		}
		// 如果Join不为空 则添加Join
		if len(s.ListQueryOp.Join) > 0 {
			for _, join := range s.ListQueryOp.Join {
				switch join.Type {
				case LeftJoin:
					m.LeftJoin(join.Model.TableName(), join.Condition).As(join.Alias)
				case RightJoin:
					m.RightJoin(join.Model.TableName(), join.Condition).As(join.Alias)
				case InnerJoin:
					m.InnerJoin(join.Model.TableName(), join.Condition).As(join.Alias)
				}
			}
		}

		// 如果fileldEQ不为空 则添加查询条件
		if len(s.ListQueryOp.FieldEQ) > 0 {
			for _, field := range s.ListQueryOp.FieldEQ {
				g.Log().Debug(ctx, field)
				if !r.Get(field).IsEmpty() {
					m.Where(field, r.Get(field))
				}
			}
		}
		// 如果KeyWordField不为空 则添加查询条件
		if !r.Get("keyWord").IsEmpty() {
			if len(s.ListQueryOp.KeyWordField) > 0 {
				var sql string
				args := garray.NewArray()
				for i, field := range s.ListQueryOp.KeyWordField {
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
					if len(v) == 3 {
						if gconv.Bool(v[2]) {
							m.Where(v[0], v[1])
						}
					}
					if len(v) == 2 {
						m.Where(v[0], v[1])
					}
				}
			}
		}
		// 如果ListQueryOp的Extend不为空 则执行Extend
		if s.ListQueryOp.Extend != nil {
			m = s.ListQueryOp.Extend(ctx, m)
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
		if Select := s.PageQueryOp.Select; Select != "" {
			m.Fields(Select)
		}
		// 如果Join不为空 则添加Join
		if len(s.PageQueryOp.Join) > 0 {
			for _, join := range s.PageQueryOp.Join {
				switch join.Type {
				case LeftJoin:
					m.LeftJoin(join.Model.TableName(), join.Condition).As(join.Alias)
				case RightJoin:
					m.RightJoin(join.Model.TableName(), join.Condition).As(join.Alias)
				case InnerJoin:
					m.InnerJoin(join.Model.TableName(), join.Condition).As(join.Alias)
				}
			}
		}
		// 如果fileldEQ不为空 则添加查询条件
		if len(s.PageQueryOp.FieldEQ) > 0 {
			for _, field := range s.PageQueryOp.FieldEQ {
				g.Log().Debug(ctx, field)
				if !r.Get(field).IsEmpty() {
					m.Where(field, r.Get(field))
				}
			}
		}
		// 如果KeyWordField不为空 则添加查询条件
		if !r.Get("keyWord").IsEmpty() {
			if len(s.PageQueryOp.KeyWordField) > 0 {
				var sql string
				args := garray.NewArray()
				for i, field := range s.PageQueryOp.KeyWordField {
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
					if len(v) == 3 {
						if gconv.Bool(v[2]) {
							m.Where(v[0], v[1])
						}
					}
					if len(v) == 2 {
						m.Where(v[0], v[1])
					}
				}
			}
		}

		// 如果PageQueryOp的Extend不为空 则执行Extend
		if s.PageQueryOp.Extend != nil {
			m = s.PageQueryOp.Extend(ctx, m)
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

	// 统计总数
	total, err = m.Clone().Count()
	if err != nil {
		return nil, err
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
	// g.Log().Debug(ctx, param, "2222222")

	return
}

// GetModel 获取model
func (s *Service) GetModel() IModel {
	return s.Model
}

// NewService 新建一个service
func NewService(model IModel) *Service {
	return &Service{
		Model: model,
	}
}
