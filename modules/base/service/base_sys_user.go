package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/model"
	"github.com/gogf/gf/v2/database/gdb"
)

type BaseSysUserService struct {
	*cool.Service
}

// Person 方法 返回不带密码的用户信息
func (s *BaseSysUserService) Person(userId uint) (res gdb.Record, err error) {
	m := cool.GDBM(s.Model)
	res, err = m.Where("id = ?", userId).FieldsEx("password").One()
	return
}

// NewBaseSysUserService 创建一个新的BaseSysUserService实例
func NewBaseSysUserService() *BaseSysUserService {
	return &BaseSysUserService{
		Service: &cool.Service{
			Model: model.NewBaseSysUser(),
		},
	}
}
