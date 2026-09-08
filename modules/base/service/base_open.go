package service

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/model"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

type BaseOpenService struct {
	*cool.Service
}

func NewBaseOpenService() *BaseOpenService {
	return &BaseOpenService{
		&cool.Service{},
	}
}

// AdminEPS 获取eps
func (s *BaseOpenService) AdminEPS(ctx g.Ctx) (result *g.Var, err error) {
	c := cool.CacheEPS
	result, err = c.GetOrSetFunc(ctx, "adminEPS", func(ctx g.Ctx) (interface{}, error) {
		return s.creatAdminEPS(ctx)
	}, 0)

	return
}

// creatAdminEPS 创建eps
func (s *BaseOpenService) creatAdminEPS(ctx g.Ctx) (adminEPS interface{}, err error) {
	var (
		baseEpsAdmin = model.NewBaseEpsAdmin()
	)

	type Api struct {
		Module  string `json:"module"`  // 所属模块名称 例如：base
		Method  string `json:"method"`  // 请求方法 例如：GET
		Path    string `json:"path"`    // 请求路径 例如：/welcome
		Prefix  string `json:"prefix"`  // 路由前缀 例如：/admin/base/open
		Summary string `json:"summary"` // 描述 例如：欢迎页面
		Tag     string `json:"tag"`     // 标签 例如：base  好像暂时不用
		Dts     string `json:"dts"`     // 未知 例如：{} 好像暂时不用
	}
	// type Column struct {
	// }
	type Module struct {
		Api     []*Api             `json:"api"`
		Columns []*cool.ColumnInfo `json:"columns"`
		Module  string             `json:"module"`
		Prefix  string             `json:"prefix"`
	}
	admineps := make(map[string][]*Module)
	// 先收集全部路由,再在事务中重建路由信息表,避免半途失败/并发重建造成数据不一致
	routers := g.Server().GetRoutes()
	epsApis := make([]*Api, 0)
	for _, router := range routers {
		if router.Type == ghttp.HandlerTypeMiddleware || router.Type == ghttp.HandlerTypeHook {
			continue
		}
		if router.Method == "ALL" {
			continue
		}
		routeSplite := gstr.Split(router.Route, "/")
		if len(routeSplite) < 5 {
			continue
		}
		if routeSplite[1] != "admin" {
			continue
		}
		module := routeSplite[2]
		method := router.Method
		// 获取最后一个元素加前缀 / 为 path
		path := "/" + routeSplite[len(routeSplite)-1]
		// 获取前面的元素为prefix
		prefix := gstr.Join(routeSplite[0:len(routeSplite)-1], "/")
		// 获取最后一个元素为summary
		summary := routeSplite[len(routeSplite)-1]
		epsApis = append(epsApis, &Api{
			Module:  module,
			Method:  method,
			Path:    path,
			Prefix:  prefix,
			Summary: summary,
			Tag:     "",
			Dts:     "",
		})
	}
	// 事务内重建:先清空再批量写入,任一步失败即回滚,错误向上返回
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		m := cool.DBM(baseEpsAdmin).TX(tx)
		if _, err := m.Where("1=1").Delete(); err != nil {
			return err
		}
		if len(epsApis) > 0 {
			if _, err := m.Data(epsApis).Insert(); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	// 读取数据库表生成eps
	items, err := cool.DBM(baseEpsAdmin).Fields("DISTINCT module,prefix").All()
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		module := item["module"].String()
		prefix := item["prefix"].String()
		apis, err := cool.DBM(baseEpsAdmin).Where("module=? AND prefix=?", module, prefix).All()
		if err != nil {
			return nil, err
		}
		var apiList []*Api
		for _, api := range apis {
			apiList = append(apiList, &Api{
				Module:  api["module"].String(),
				Method:  api["method"].String(),
				Path:    api["path"].String(),
				Prefix:  api["prefix"].String(),
				Summary: api["summary"].String(),
				Tag:     api["tag"].String(),
				Dts:     api["dts"].String(),
			})
		}
		admineps[module] = append(admineps[module], &Module{
			Api:     apiList,
			Columns: cool.ModelInfo[prefix],
			Module:  module,
			Prefix:  prefix,
		})

	}

	adminEPS = gjson.New(admineps)
	return

}

// AdminEPS 获取eps
func (s *BaseOpenService) AppEPS(ctx g.Ctx) (result *g.Var, err error) {
	c := cool.CacheEPS
	result, err = c.GetOrSetFunc(ctx, "appEPS", func(ctx g.Ctx) (interface{}, error) {
		return s.creatAppEPS(ctx)
	}, 0)

	return
}

// creatAppEPS 创建app eps
func (s *BaseOpenService) creatAppEPS(ctx g.Ctx) (appEPS interface{}, err error) {
	var (
		baseEpsApp = model.NewBaseEpsApp()
	)

	type Api struct {
		Module  string `json:"module"`  // 所属模块名称 例如：base
		Method  string `json:"method"`  // 请求方法 例如：GET
		Path    string `json:"path"`    // 请求路径 例如：/welcome
		Prefix  string `json:"prefix"`  // 路由前缀 例如：/admin/base/open
		Summary string `json:"summary"` // 描述 例如：欢迎页面
		Tag     string `json:"tag"`     // 标签 例如：base  好像暂时不用
		Dts     string `json:"dts"`     // 未知 例如：{} 好像暂时不用
	}
	// type Column struct {
	// }
	type Module struct {
		Api     []*Api             `json:"api"`
		Columns []*cool.ColumnInfo `json:"columns"`
		Module  string             `json:"module"`
		Prefix  string             `json:"prefix"`
	}
	appeps := make(map[string][]*Module)
	// 先收集全部路由,再在事务中重建路由信息表,避免半途失败/并发重建造成数据不一致
	routers := g.Server().GetRoutes()
	epsApis := make([]*Api, 0)
	for _, router := range routers {
		if router.Type == ghttp.HandlerTypeMiddleware || router.Type == ghttp.HandlerTypeHook {
			continue
		}
		if router.Method == "ALL" {
			continue
		}
		routeSplite := gstr.Split(router.Route, "/")
		if len(routeSplite) < 5 {
			continue
		}
		if routeSplite[1] != "app" {
			continue
		}
		module := routeSplite[2]
		method := router.Method
		// 获取最后一个元素加前缀 / 为 path
		path := "/" + routeSplite[len(routeSplite)-1]
		// 获取前面的元素为prefix
		prefix := gstr.Join(routeSplite[0:len(routeSplite)-1], "/")
		// 获取最后一个元素为summary
		summary := routeSplite[len(routeSplite)-1]
		epsApis = append(epsApis, &Api{
			Module:  module,
			Method:  method,
			Path:    path,
			Prefix:  prefix,
			Summary: summary,
			Tag:     "",
			Dts:     "",
		})
	}
	// 事务内重建:先清空再批量写入,任一步失败即回滚,错误向上返回
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		m := cool.DBM(baseEpsApp).TX(tx)
		if _, err := m.Where("1=1").Delete(); err != nil {
			return err
		}
		if len(epsApis) > 0 {
			if _, err := m.Data(epsApis).Insert(); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	// 读取数据库表生成eps
	items, err := cool.DBM(baseEpsApp).Fields("DISTINCT module,prefix").All()
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		module := item["module"].String()
		prefix := item["prefix"].String()
		apis, err := cool.DBM(baseEpsApp).Where("module=? AND prefix=?", module, prefix).All()
		if err != nil {
			return nil, err
		}
		var apiList []*Api
		for _, api := range apis {
			apiList = append(apiList, &Api{
				Module:  api["module"].String(),
				Method:  api["method"].String(),
				Path:    api["path"].String(),
				Prefix:  api["prefix"].String(),
				Summary: api["summary"].String(),
				Tag:     api["tag"].String(),
				Dts:     api["dts"].String(),
			})
		}
		appeps[module] = append(appeps[module], &Module{
			Api:     apiList,
			Columns: cool.ModelInfo[prefix],
			Module:  module,
			Prefix:  prefix,
		})

	}

	appEPS = gjson.New(appeps)
	return
}
