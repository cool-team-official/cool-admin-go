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
	ModifyBefore(ctx context.Context, method string, param g.MapStrAny) (err error)  // 新增|删除|修改前的操作
	ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error)   // 新增|删除|修改后的操作
	GetModel() IModel                                                                // 获取model
}
type Service struct {
	Model              IModel
	ListQueryOp        *QueryOp
	PageQueryOp        *QueryOp
	InsertParam        func(ctx context.Context) g.MapStrAny // Add时插入参数
	Before             func(ctx context.Context) (err error) // CRUD前的操作
	InfoIgnoreProperty string                                // Info时忽略的字段,多个字段用逗号隔开
	UniqueKey          g.MapStrStr                           // 唯一键 key:字段名 value:错误信息
	NotNullKey         g.MapStrStr                           // 非空键 key:字段名 value:错误信息
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

// ServiceAdd 新增
func (s *Service) ServiceAdd(ctx context.Context, req *AddReq) (data interface{}, err error) {
	r := g.RequestFromCtx(ctx)

	rmap := r.GetMap()
	// 非空键
	if s.NotNullKey != nil {
		for k, v := range s.NotNullKey {
			if rmap[k] == nil {
				return nil, gerror.New(v)
			}
		}
	}
	// 唯一键
	if s.UniqueKey != nil {
		for k, v := range s.UniqueKey {
			if rmap[k] != nil {
				m := DBM(s.Model)
				count, err := m.Where(k, rmap[k]).Count()
				if err != nil {
					return nil, err
				}
				if count > 0 {
					err = gerror.New(v)
					return nil, err
				}
			}
		}
	}
	if s.InsertParam != nil {
		insertParams := s.InsertParam(ctx)
		if len(insertParams) > 0 {
			for k, v := range insertParams {
				rmap[k] = v
			}
		}
	}
	m := DBM(s.Model)
	lastInsertId, err := m.Data(rmap).InsertAndGetId()
	if err != nil {
		return
	}
	data = g.Map{"id": lastInsertId}

	return
}

// ServiceDelete 删除
func (s *Service) ServiceDelete(ctx context.Context, req *DeleteReq) (data interface{}, err error) {
	ids := g.RequestFromCtx(ctx).Get("ids").Slice()
	m := g.DB(s.Model.GroupName()).Model(s.Model.TableName())
	data, err = m.WhereIn("id", ids).Delete()

	return
}

// ServiceUpdate 修改
func (s *Service) ServiceUpdate(ctx context.Context, req *UpdateReq) (data interface{}, err error) {
	r := g.RequestFromCtx(ctx)
	rmap := r.GetMap()
	if rmap["id"] == nil {
		err = gerror.New("id不能为空")
		return
	}
	if s.UniqueKey != nil {
		for k, v := range s.UniqueKey {
			if rmap[k] != nil {
				count, err := DBM(s.Model).Where(k, rmap[k]).WhereNot("id", rmap["id"]).Count()
				if err != nil {
					return nil, err
				}
				if count > 0 {
					err = gerror.New(v)
					return nil, err
				}
			}
		}
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
				if !r.Get(field).IsEmpty() {
					m.Where(field, r.Get(field))
				}
			}
		}
		// 如果KeyWordField不为空 则添加查询条件
		if !r.Get("keyWord").IsEmpty() {
			if len(s.PageQueryOp.KeyWordField) > 0 {
				builder := m.Builder()
				for _, field := range s.PageQueryOp.KeyWordField {
					g.DumpWithType(field)
					// builder.WhereLike(field, "%"+r.Get("keyWord").String()+"%")
					builder = builder.WhereOrLike(field, "%"+r.Get("keyWord").String()+"%")
				}
				m.Where(builder)
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

	// 增加默认数据限制，防止查询所有数据
	m.Limit(10000)

	result, err := m.All()
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
		r           = g.RequestFromCtx(ctx)
		total int64 = 0
	)

	type pagination struct {
		Page  int   `json:"page"`
		Size  int   `json:"size"`
		Total int64 `json:"total"`
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
				if !r.Get(field).IsEmpty() {
					m.Where(field, r.Get(field))
				}
			}
		}
		// 如果KeyWordField不为空 则添加查询条件
		if !r.Get("keyWord").IsEmpty() {
			if len(s.PageQueryOp.KeyWordField) > 0 {
				builder := m.Builder()
				for _, field := range s.PageQueryOp.KeyWordField {
					g.DumpWithType(field)
					// builder.WhereLike(field, "%"+r.Get("keyWord").String()+"%")
					builder = builder.WhereOrLike(field, "%"+r.Get("keyWord").String()+"%")
				}
				m.Where(builder)
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

	// 统计总数
	total, err = m.Clone().Count()
	if err != nil {
		return nil, err
	}
	if s.PageQueryOp != nil {
		if Select := s.PageQueryOp.Select; Select != "" {
			m.Fields(Select)
		}
	}
	// 如果 req.Order 和 req.Sort 均不为空 则添加排序
	if !r.Get("order").IsEmpty() && !r.Get("sort").IsEmpty() {
		m.Order(r.Get("order").String() + " " + r.Get("sort").String())
	}

	// 如果req.IsExport为true 则导出数据
	if req.IsExport {
		// 如果req.MaxExportSize大于0 则限制导出数据的最大条数
		if req.MaxExportLimit > 0 {
			m.Limit(req.MaxExportLimit)
		}
		result, err := m.All()
		if err != nil {
			return nil, err
		}
		data = g.Map{
			"list":  result,
			"total": total,
		}
		return data, nil
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

// ModifyBefore 新增|删除|修改前的操作
func (s *Service) ModifyBefore(ctx context.Context, method string, param g.MapStrAny) (err error) {
	// g.Log().Debugf(ctx, "ModifyBefore: %s", method)
	return
}

// ModifyAfter 新增|删除|修改后的操作
func (s *Service) ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error) {
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
