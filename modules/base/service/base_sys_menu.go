package service

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/model"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type BaseSysMenuService struct {
	*cool.Service
}

// GetPerms 获取菜单的权限
func (s *BaseSysMenuService) GetPerms(roleIds []string) []string {
	var (
		perms  []string
		result gdb.Result
	)
	m := cool.DBM(s.Model).As("a")
	// 如果roldIds 包含 1 则表示是超级管理员，则返回所有权限
	if garray.NewIntArrayFrom(gconv.Ints(roleIds)).Contains(1) {
		result, _ = m.Fields("a.perms").All()
	} else {
		result, _ = m.InnerJoin("base_sys_role_menu b", "a.id=b.menuId").InnerJoin("base_sys_role c", "b.roleId=c.id").Where("c.id IN (?)", roleIds).Fields("a.perms").All()
	}
	for _, v := range result {
		vmap := v.Map()
		if vmap["perms"] != nil {
			p := gstr.Split(vmap["perms"].(string), ",")
			perms = append(perms, p...)
		}
	}
	return perms
}

// GetMenus 获取菜单
func (s *BaseSysMenuService) GetMenus(roleIds []string, isAdmin bool) (result gdb.Result) {
	// 屏蔽 base_sys_role_menu.id 防止部分权限的用户登录时菜单渲染错误
	m := cool.DBM(s.Model).As("a").Fields("a.*")
	if isAdmin {
		result, _ = m.Group("a.id").Order("a.orderNum asc").All()
	} else {
		result, _ = m.InnerJoin("base_sys_role_menu b", "a.id=b.menuId").Where("b.roleId IN (?)", roleIds).Group("a.id").Order("a.orderNum asc").All()
	}
	return

}

// ModifyAfter 修改后
func (s *BaseSysMenuService) ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error) {
	if method == "Delete" {
		ids := gconv.Ints(param["ids"])
		if len(ids) > 0 {
			_, err = cool.DBM(s.Model).Where("parentId IN (?)", ids).Delete()
		}
		return
	}
	return
}

// ServiceAdd 添加
func (s *BaseSysMenuService) ServiceAdd(ctx context.Context, req *cool.AddReq) (data interface{}, err error) {
	r := g.RequestFromCtx(ctx)
	rjson, err := r.GetJson()
	if err != nil {
		return
	}
	g.DumpWithType(rjson)
	m := cool.DBM(s.Model)
	lastInsertId, err := m.Data(rjson).InsertAndGetId()
	if err != nil {
		return
	}
	data = g.Map{"id": lastInsertId}
	return
}

// NewBaseSysMenuService 创建一个BaseSysMenuService实例
func NewBaseSysMenuService() *BaseSysMenuService {
	return &BaseSysMenuService{
		&cool.Service{
			Model: model.NewBaseSysMenu(),
		},
	}
}
