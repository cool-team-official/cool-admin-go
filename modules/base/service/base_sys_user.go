package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/model"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type BaseSysUserService struct {
	*cool.Service
}

// Person 方法 返回不带密码的用户信息
func (s *BaseSysUserService) Person(userId uint) (res gdb.Record, err error) {
	m := cool.DBM(s.Model)
	res, err = m.Where("id = ?", userId).FieldsEx("password").One()
	return
}

// ServiceInfo 方法 返回服务信息
func (s *BaseSysUserService) ServiceInfo(ctx g.Ctx, req *cool.InfoReq) (data interface{}, err error) {
	result, err := s.Service.ServiceInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if result.(gdb.Record).IsEmpty() {
		return nil, nil
	}
	// g.DumpWithType(result)
	resultMap := result.(gdb.Record).Map()

	// 获取角色
	roleIds, err := cool.DBM(model.NewBaseSysUserRole()).Where("userId = ?", resultMap["id"]).Fields("roleId").Array()
	if err != nil {
		return nil, err
	}
	resultMap["roleIdList"] = roleIds
	data = resultMap

	return
}

// NewBaseSysUserService 创建一个新的BaseSysUserService实例
func NewBaseSysUserService() *BaseSysUserService {
	return &BaseSysUserService{
		Service: &cool.Service{
			Model:              model.NewBaseSysUser(),
			InfoIgnoreProperty: "password",
		},
	}
}
