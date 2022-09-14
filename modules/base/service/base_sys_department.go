package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/model"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
)

type BaseSysDepartmentService struct {
	*cool.Service
}

// GetByRoleIds 获取部门
func (s *BaseSysDepartmentService) GetByRoleIds(roleIds []uint, isAdmin bool) (res []uint) {
	var (
		result                gdb.Result
		BaseSysRoleDepartment = model.NewBaseSysRoleDepartment()
	)
	// 如果roleIds不为空
	if len(roleIds) > 0 {
		// 如果是超级管理员，则返回所有部门
		if isAdmin {
			result, _ = cool.DBM(s.Model).Fields("id").All()
			for _, v := range result {
				vmap := v.Map()
				if vmap["id"] != nil {
					res = append(res, gconv.Uint(vmap["id"]))
				}
			}
		} else {
			// 如果不是超级管理员，则返回角色所在部门
			result, _ = cool.DBM(BaseSysRoleDepartment).Where("roleId IN (?)", roleIds).Fields("departmentId").All()
			for _, v := range result {
				vmap := v.Map()
				if vmap["departmentId"] != nil {
					res = append(res, vmap["departmentId"].(uint))
				}
			}
		}

	}
	return
}

// NewBaseSysDepartmentService 创建一个BaseSysDepartmentService实例
func NewBaseSysDepartmentService() *BaseSysDepartmentService {
	return &BaseSysDepartmentService{
		Service: &cool.Service{
			Model:       model.NewBaseSysDepartment(),
			ListQueryOp: &cool.QueryOp{},
		},
	}
}
