package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
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
	// 获取所有路由并更新到数据库表 base_eps_admin
	g.Model("base_eps_admin").Where("1=1").Delete()
	routers := g.Server().GetRoutes()
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
		g.DB().Model("base_eps_admin").Insert(&Api{
			Module:  module,
			Method:  method,
			Path:    path,
			Prefix:  prefix,
			Summary: summary,
			Tag:     "",
			Dts:     "",
		})
	}
	// 读取数据库表生成eps
	// var modules []*Module
	items, _ := g.Model("base_eps_admin").Fields("DISTINCT module,prefix").All()
	for _, item := range items {
		module := item["module"].String()
		prefix := item["prefix"].String()
		apis, _ := g.Model("base_eps_admin").Where("module=? AND prefix=?", module, prefix).All()
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
