package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/model"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
)

type BaseSysMenuService struct {
	*cool.Service
}

// GetPerms 获取菜单的权限
func (s *BaseSysMenuService) GetPerms(roleIds []int) []string {
	var (
		perms  []string
		result gdb.Result
	)
	m := cool.GDBModel(s.Model).As("a")
	// 如果roldIds 包含 1 则表示是超级管理员，则返回所有权限
	if garray.NewIntArrayFrom(roleIds).Contains(1) {
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
func (s *BaseSysMenuService) GetMenus(roleIds []int, isAdmin bool) (result gdb.Result) {
	m := cool.GDBModel(s.Model).As("a")
	if isAdmin {
		result, _ = m.Group("a.id").Order("a.orderNum asc").All()
	} else {
		result, _ = m.InnerJoin("base_sys_role_menu b", "a.id=b.menuId").Where("b.roleId IN (?)", roleIds).Group("a.id").Order("a.orderNum asc").All()
	}
	return

}

// NewBaseSysMenuService 创建一个BaseSysMenuService实例
func NewBaseSysMenuService() *BaseSysMenuService {
	return &BaseSysMenuService{
		&cool.Service{
			Model: model.NewBaseSysMenu(),
		},
		// 	Service: cool.NewService(model.NewBaseSysMenu()),
	}
}
