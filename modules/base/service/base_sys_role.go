package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/model"
)

var baseSysUserRole = model.NewBaseSysUserRole()

type BaseSysRoleService struct {
	*cool.Service
}

// GetByUser get array  roleId by userId
func (s *BaseSysRoleService) GetByUser(userId uint) []int {
	var (
		roles []int
	)
	res, _ := cool.GDBModel(baseSysUserRole).Where("userId = ?", userId).Array("roleId")
	for _, v := range res {
		roles = append(roles, v.Int())
	}
	return roles
}

// NewBaseSysRoleService create a new BaseSysRoleService
func NewBaseSysRoleService() *BaseSysRoleService {
	return &BaseSysRoleService{
		Service: &cool.Service{
			Model: model.NewBaseSysRole(),
		},
	}
}
