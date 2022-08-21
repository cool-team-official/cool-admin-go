package service

import (
	"context"
	"time"

	"github.com/cool-team-official/cool-admin-go/cool"
	v1 "github.com/cool-team-official/cool-admin-go/modules/base/api/v1"
	"github.com/cool-team-official/cool-admin-go/modules/base/model"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/golang-jwt/jwt/v4"
)

type BaseSysLoginService struct {
	*cool.Service
}

// login
func (s *BaseSysLoginService) Login(ctx context.Context, req *v1.BaseOpenLoginReq) (data interface{}, err error) {
	type Result struct {
		Expire         int64  `json:"expire"`
		Token          string `json:"token"`
		RefreshExpires int64  `json:"refreshExpires"`
		RefreshToken   string `json:"refreshToken"`
	}

	var (
		captchaId                = req.CaptchaId
		verifyCode               = req.VerifyCode
		password                 = req.Password
		username                 = req.Username
		baseSysRoleService       = NewBaseSysRoleService()
		baseSysMenuService       = NewBaseSysMenuService()
		baseSysDepartmentService = NewBaseSysDepartmentService()
		baseSysUser              = model.NewBaseSysUser()
		result                   = &Result{}
	)

	vcode, _ := cool.Cache.Get(ctx, captchaId)
	if vcode.String() != verifyCode {
		err = gerror.New("验证码错误")
		return
	}
	md5password, _ := gmd5.Encrypt(password)

	var user *model.BaseSysUser
	cool.GDBModel(baseSysUser).Where("username=?", username).Where("password=?", md5password).Where("status=?", 1).Scan(&user)
	if user == nil {
		err = gerror.New("账户或密码不正确~")
		return
	}
	// 获取用户角色
	roleIds := baseSysRoleService.GetByUser(user.ID)
	// 如果没有角色，则报错
	if len(roleIds) == 0 {
		err = gerror.New("该用户未设置任何角色，无法登录~")
		return
	}
	// 生成token
	result.Expire = g.Cfg().MustGet(ctx, "cool-base.jwt.token.expire").Int64()
	result.RefreshExpires = g.Cfg().MustGet(ctx, "cool-base.jwt.token.refresh").Int64()
	result.Token = s.GenerateToken(ctx, user, roleIds, result.Expire, false)
	result.RefreshToken = s.GenerateToken(ctx, user, roleIds, result.RefreshExpires, true)
	// 将用户相关信息保存到缓存
	perms := baseSysMenuService.GetPerms(roleIds)
	departments := baseSysDepartmentService.GetByRoleIds(roleIds, user.Username == "admin")
	cool.Cache.Set(ctx, "admin:department:"+gconv.String(user.ID), departments, 0)
	cool.Cache.Set(ctx, "admin:perms:"+gconv.String(user.ID), perms, 0)
	cool.Cache.Set(ctx, "admin:token:"+gconv.String(user.ID), result.Token, 0)
	cool.Cache.Set(ctx, "admin:token:refresh:"+gconv.String(user.ID), result.RefreshToken, 0)

	data = result
	return
}

// Captcha 图形验证码
func (*BaseSysLoginService) Captcha(req *v1.BaseOpenCaptchaReq) (interface{}, error) {
	type capchaInfo struct {
		CaptchaId string `json:"captchaId"`
		Data      string `json:"data"`
	}
	var (
		ctx g.Ctx
		err error

		result = &capchaInfo{}
	)
	captchaText := grand.Digits(4)
	svg := `<svg width="150" height="50" xmlns="http://www.w3.org/2000/svg"><text x="75" y="25" text-anchor="middle" font-size="25" fill="#fff">` + captchaText + `</text></svg>`
	svgbase64 := gbase64.EncodeString(svg)

	result.Data = `data:image/svg+xml;base64,` + svgbase64
	result.CaptchaId = guid.S()
	cool.Cache.Set(ctx, result.CaptchaId, captchaText, 1800*time.Second)
	g.Log().Debug(ctx, "验证码", result.CaptchaId, captchaText)
	return result, err
}

// generateToken 生成token
func (*BaseSysLoginService) GenerateToken(ctx context.Context, user *model.BaseSysUser, roleIds []int, exprire int64, isRefresh bool) (token string) {
	err := cool.Cache.Set(ctx, "admin:passwordVersion:"+gconv.String(user.ID), gconv.String(user.PasswordV), 0)
	if err != nil {
		g.Log().Error(ctx, "生成token失败", err)
	}
	type Claims struct {
		IsRefresh       bool   `json:"isRefresh"`
		RoleIds         []int  `json:"roleIds"`
		Username        string `json:"username"`
		UserId          uint   `json:"userId"`
		PasswordVersion *int32 `json:"passwordVersion"`
		jwt.RegisteredClaims
	}
	claims := &Claims{
		IsRefresh:       isRefresh,
		RoleIds:         roleIds,
		Username:        user.Username,
		UserId:          user.ID,
		PasswordVersion: user.PasswordV,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(exprire) * time.Second)),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = tokenClaims.SignedString(g.Cfg().MustGet(ctx, "cool-base.jwt.secret").Bytes())
	if err != nil {
		g.Log().Error(ctx, "生成token失败", err)
	}
	return
}

// NewBaseSysLoginService 创建一个新的BaseSysLoginService
func NewBaseSysLoginService() *BaseSysLoginService {
	return &BaseSysLoginService{}
}
