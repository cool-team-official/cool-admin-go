package service

import (
	"context"
	"fmt"
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/model"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gset"
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

func (s *BaseSysUserService) ModifyBefore(ctx context.Context, method string, param g.MapStrAny) (err error) {
	if method == "Delete" {
		// 禁止删除超级管理员
		userIds := garray.NewIntArrayFrom(gconv.Ints(param["ids"]))
		currentId, found := userIds.Get(0)
		superAdminId := 1

		if userIds.Len() == 1 && found && currentId == superAdminId {
			err = gerror.New("超级管理员不能删除")
			return
		}

		// 删除超级管理员
		userIds.RemoveValue(1)
		g.RequestFromCtx(ctx).SetParam("ids", userIds.Slice())
	}
	return
}

func (s *BaseSysUserService) ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error) {
	if method == "Delete" {
		userIds := garray.NewIntArrayFrom(gconv.Ints(param["ids"]))
		userIds.RemoveValue(1)
		// 删除用户时删除相关数据
		cool.DBM(model.NewBaseSysUserRole()).WhereIn("userId", userIds.Slice()).Delete()
	}
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

// ServiceUpdate 方法 更新用户信息
func (s *BaseSysUserService) ServiceUpdate(ctx context.Context, req *cool.UpdateReq) (data interface{}, err error) {
	var (
		admin = cool.GetAdmin(ctx)
		m     = cool.DBM(s.Model)
	)

	r := g.RequestFromCtx(ctx)
	rMap := r.GetMap()

	// 如果不传如ID代表更新当前用户
	userId := r.Get("id", admin.UserId).Uint()

	userInfo, err := m.Where("id = ?", userId).One()

	if err != nil {
		return
	}
	if userInfo.IsEmpty() {
		err = gerror.New("用户不存在")
		return
	}

	// 如果请求的password不为空并且密码加密后的值有变动，说明要修改密码
	var rPassword = r.Get("password", "").String()
	if rPassword != "" && rPassword != userInfo["password"].String() {
		rMap["password"], _ = gmd5.Encrypt(rPassword)
		rMap["passwordV"] = userInfo["passwordV"].Int() + 1
		cool.CacheManager.Set(ctx, fmt.Sprintf("admin:passwordVersion:%d", userId), rMap["passwordV"], 0)
	} else {
		delete(rMap, "password")
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) (err error) {
		roleModel := cool.DBM(model.NewBaseSysUserRole()).TX(tx).Where("userId = ?", userId)
		roleIds, err := roleModel.Fields("roleId").Array()
		if err != nil {
			return
		}

		inRoleIdSet := gset.NewFrom(r.Get("roleIdList").Ints())
		roleIdsSet := gset.NewFrom(gconv.Ints(roleIds))

		// 判断是否相等
		if roleIdsSet.Diff(inRoleIdSet).Size() != 0 || inRoleIdSet.Diff(roleIdsSet).Size() != 0 {
			roleArray := garray.NewArray()
			inRoleIdSet.Iterator(func(v interface{}) bool {
				roleArray.PushRight(g.Map{
					"userId": gconv.Uint(userId),
					"roleId": gconv.Uint(v),
				})
				return true
			})

			_, err = roleModel.Delete()

			if err != nil {
				return err
			}

			_, err = roleModel.Fields("userId,roleId").Insert(roleArray)
			if err != nil {
				return err
			}
		}

		_, err = m.TX(tx).Update(rMap)

		if err != nil {
			return err
		}
		return
	})
	return
}

//Move 移动用户部门
func (s *BaseSysUserService) Move(ctx g.Ctx) (err error) {
	request := g.RequestFromCtx(ctx)
	departmentId := request.Get("departmentId").Int()
	userIds := request.Get("userIds").Slice()

	_, err = cool.DBM(s.Model).Where("`id` IN(?)", userIds).Data(g.Map{"departmentId": departmentId}).Update()

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
			PageQueryOp: &cool.QueryOp{
				Select: "base_sys_user.*,dept.`name` as departmentName,GROUP_CONCAT( role.`name` ) AS `roleName`",
				Join: []*cool.JoinOp{
					{
						Model:     model.NewBaseSysDepartment(),
						Alias:     "dept",
						Type:      "LeftJoin",
						Condition: "`base_sys_user`.`departmentId` = `dept`.`id`",
					},
					{
						Model:     model.NewBaseSysUserRole(),
						Alias:     "user_role",
						Type:      "LeftJoin",
						Condition: "`base_sys_user`.`id` = `user_role`.`userId`",
					},
					{
						Model:     model.NewBaseSysRole(),
						Alias:     "`role`",
						Type:      "LeftJoin",
						Condition: "`role`.`id` = `user_role`.`roleId`",
					},
				},
				Where: func(ctx context.Context) []g.Array {
					r := g.RequestFromCtx(ctx).GetMap()
					return []g.Array{
						{"(departmentId IN (?))", gconv.SliceStr(r["departmentIds"])},
					}
				},
				Extend: func(ctx g.Ctx, m *gdb.Model) *gdb.Model {
					return m.Group("`base_sys_user`.`id`")
				},
				KeyWordField: []string{"name", "username", "nickName"},
			},
		},
	}
}
