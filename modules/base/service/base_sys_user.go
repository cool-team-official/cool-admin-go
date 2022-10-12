package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/model"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
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

// PersonUpdate 方法 更新用户信息
func (s *BaseSysUserService) PersonUpdate(ctx g.Ctx) (err error) {
	var (
		admin = cool.GetAdmin(ctx)
		req   = g.RequestFromCtx(ctx).GetMap()
		m     = cool.DBM(s.Model)
	)
	userInfo, err := m.Where("id = ?", admin.UserId).One()
	if err != nil {
		return err
	}
	if userInfo.IsEmpty() {
		return gerror.New("用户不存在")
	}
	req["id"] = admin.UserId
	// 如果 req["password"] 不为空，说明要修改密码
	if req["password"] == "" {
		delete(req, "password")
	}
	if req["password"] != nil {
		req["password"], _ = gmd5.Encrypt(req["password"].(string))
		req["passwordV"] = userInfo["passwordV"].Int() + 1
		cool.CacheManager.Set(ctx, "admin:passwordVersion:"+gconv.String(admin.UserId), gconv.String(req["passwordV"]), 0)
	}
	// g.Dump(req)
	_, err = m.Update(req)
	return
}

// NewBaseSysUserService 创建一个新的BaseSysUserService实例
func NewBaseSysUserService() *BaseSysUserService {
	return &BaseSysUserService{
		Service: &cool.Service{
			Model:              model.NewBaseSysUser(),
			InfoIgnoreProperty: "password",
			UniqueKey: map[string]string{
				"username": "用户名不能重复",
			},
		},
	}
}
