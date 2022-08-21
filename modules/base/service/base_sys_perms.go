package service

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/v2/database/gdb"
)

type BaseSysPermsService struct {
}

func NewBaseSysPermsService() *BaseSysPermsService {
	return &BaseSysPermsService{}
}

// permmenu 方法
func (c *BaseSysPermsService) Permmenu(ctx context.Context, roleIds []int) (res interface{}) {
	type permmenu struct {
		Perms []string   `json:"perms"`
		Menus gdb.Result `json:"menus"`
	}
	var (
		baseSysMenuService = NewBaseSysMenuService()
		admin              = cool.GetAdmin(ctx)
	)

	res = &permmenu{
		Perms: baseSysMenuService.GetPerms(roleIds),
		Menus: baseSysMenuService.GetMenus(admin.RoleIds, admin.UserId == 1),
	}

	return

}
