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

// ServiceAdd 方法 添加用户
func (s *BaseSysUserService) ServiceAdd(ctx context.Context, req *cool.AddReq) (data interface{}, err error) {
	var (
		m      = cool.DBM(s.Model)
		r      = g.RequestFromCtx(ctx)
		reqmap = r.GetMap()
	)
	// 如果reqmap["password"]不为空，则对密码进行md5加密
	if !r.Get("password").IsNil() {
		reqmap["password"] = gmd5.MustEncryptString(r.Get("password").String())
	}
	lastInsertId, err := m.Data(reqmap).InsertAndGetId()
	if err != nil {
		return
	}
	data = g.Map{"id": lastInsertId}
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

	// 是否来自"个人中心更新"(comm/personUpdate),该入口仅允许修改当前登录用户本人
	isPersonUpdate := r.GetCtxVar("isPersonUpdate", false).Bool()

	// 如果不传入ID代表更新当前用户
	userId := r.Get("id", admin.UserId).Uint()
	if isPersonUpdate {
		// 强制更新对象为当前登录用户,防止通过个人中心接口传入他人ID越权修改
		userId = admin.UserId
	}
	userInfo, err := m.Where("id = ?", userId).One()

	if err != nil {
		return
	}
	if userInfo.IsEmpty() {
		err = gerror.New("用户不存在")
		return
	}

	// 禁止禁用超级管理员
	if userId == 1 && (!r.Get("status").IsNil() && r.Get("status").Int() == 0) {
		err = gerror.New("禁止禁用超级管理员")
		return
	}

	// 如果请求的password不为空并且密码加密后的值有变动，说明要修改密码
	var rPassword = r.Get("password", "").String()
	if rPassword != "" && rPassword != userInfo["password"].String() {
		rMap["password"], _ = gmd5.Encrypt(rPassword)
		rMap["passwordV"] = userInfo["passwordV"].Int() + 1
		cool.CacheManager.Set(ctx, fmt.Sprintf("admin:passwordVersion:%d", userId), rMap["passwordV"], 0)
		// 密码变更后立即吊销该用户已签发的accessToken与refreshToken
		cool.CacheManager.Remove(ctx, "admin:token:"+gconv.String(userId))
		cool.CacheManager.Remove(ctx, "admin:token:refresh:"+gconv.String(userId))
	} else {
		delete(rMap, "password")
	}

	// 个人中心更新仅允许修改个人资料字段,禁止携带角色/部门/状态等管理字段,防止越权提权
	if isPersonUpdate {
		allowSelfFields := map[string]bool{
			"name":     true,
			"nickName": true,
			"headImg":  true,
			"phone":    true,
			"email":    true,
			"remark":   true,
			"password": true,
		}
		for k := range rMap {
			if !allowSelfFields[k] {
				delete(rMap, k)
			}
		}
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		roleModel := cool.DBM(model.NewBaseSysUserRole()).TX(tx).Where("userId = ?", userId)
		roleIds, err := roleModel.Fields("roleId").Array()
		if err != nil {
			return
		}

		// 个人中心更新不允许修改角色,防止普通用户给自己分配超管等高权限角色
		if isPersonUpdate {
			delete(rMap, "roleIdList")
		} else if !r.Get("roleIdList").IsNil() {
			inRoleIdSet := gset.NewFrom(r.Get("roleIdList").Ints())
			roleIdsSet := gset.NewFrom(gconv.Ints(roleIds))

			// 如果请求的角色信息未发生变化则跳过更新逻辑
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
		}

		_, err = m.TX(tx).Update(rMap)

		if err != nil {
			return err
		}
		return
	})
	return
}

// Move 移动用户部门
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
					r := g.RequestFromCtx(ctx)
					admin := cool.GetAdmin(ctx)
					// 请求中的部门范围(前端按选中部门节点下发)
					reqIds := gconv.SliceUint(r.Get("departmentIds").Val())
					// 当前管理员被授权的部门范围(登录/刷新权限时缓存,超管为全部部门)
					allowedVar, _ := cool.CacheManager.Get(ctx, "admin:department:"+gconv.String(admin.UserId))
					allowedIds := gconv.SliceUint(allowedVar.Val())
					// 授权范围缺失(缓存未命中/未配置数据权限)时,退化为仅按请求部门过滤,保证可用性
					if len(allowedIds) == 0 {
						if len(reqIds) == 0 {
							return []g.Array{{"0=1"}}
						}
						return []g.Array{{"(departmentId IN (?))", reqIds}}
					}
					// 未指定部门时,默认查询全部授权部门
					if len(reqIds) == 0 {
						return []g.Array{{"(departmentId IN (?))", allowedIds}}
					}
					// 仅允许查询“请求部门∩授权部门”,防止跨部门越权拉取用户数据
					allowedSet := gset.NewFrom(allowedIds)
					intersectIds := make([]uint, 0, len(reqIds))
					for _, id := range reqIds {
						if allowedSet.Contains(id) {
							intersectIds = append(intersectIds, id)
						}
					}
					if len(intersectIds) == 0 {
						return []g.Array{{"0=1"}}
					}
					return []g.Array{{"(departmentId IN (?))", intersectIds}}
				},
				Extend: func(ctx g.Ctx, m *gdb.Model) *gdb.Model {
					return m.Group("`base_sys_user`.`id`")
				},
				// 分页/导出结果剔除密码等敏感字段,防止密码哈希泄露
				ModifyResult: func(ctx g.Ctx, data interface{}) interface{} {
					if resultMap, ok := data.(g.Map); ok {
						if list, ok := resultMap["list"].(gdb.Result); ok {
							for _, record := range list {
								delete(record, "password")
								delete(record, "passwordV")
							}
						}
					}
					return data
				},
				KeyWordField: []string{"name", "username", "nickName"},
			},
		},
	}
}
