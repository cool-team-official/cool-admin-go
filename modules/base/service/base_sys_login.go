package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/cool-team-official/cool-admin-go/cool"
	v1 "github.com/cool-team-official/cool-admin-go/modules/base/api/v1"
	"github.com/cool-team-official/cool-admin-go/modules/base/config"
	"github.com/cool-team-official/cool-admin-go/modules/base/model"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/guid"
)

type BaseSysLoginService struct {
	*cool.Service
}

type TokenResult struct {
	Expire        uint   `json:"expire"`
	Token         string `json:"token"`
	RefreshExpire uint   `json:"refreshExpire"`
	RefreshToken  string `json:"refreshToken"`
}

// Login 登录
func (s *BaseSysLoginService) Login(ctx context.Context, req *v1.BaseOpenLoginReq) (result *TokenResult, err error) {
	var (
		captchaId   = req.CaptchaId
		verifyCode  = req.VerifyCode
		password    = req.Password
		username    = req.Username
		baseSysUser = model.NewBaseSysUser()
		// 同一账号+IP在10分钟内连续失败5次后临时锁定,降低口令爆破风险
		failKey = "admin:loginFail:" + username + ":" + g.RequestFromCtx(ctx).GetClientIp()
	)

	// 检查登录失败次数是否达到阈值
	failCount, _ := cool.CacheManager.Get(ctx, failKey)
	if failCount.Int() >= 5 {
		err = gerror.New("登录失败次数过多,请10分钟后再试~")
		return
	}

	vcode, _ := cool.CacheManager.Get(ctx, captchaId)
	// 验证码校验:不存在或错误均返回同一提示,防枚举;验证码一次性使用
	if vcode.IsNil() || vcode.String() != verifyCode {
		err = gerror.New("验证码错误")
		return
	}
	// 验证码校验成功后立即作废,防止同一验证码重放爆破
	cool.CacheManager.Remove(ctx, captchaId)

	md5password, _ := gmd5.Encrypt(password)

	var user *model.BaseSysUser
	if err = cool.DBM(baseSysUser).Where("username=?", username).Where("password=?", md5password).Where("status=?", 1).Scan(&user); err != nil {
		// 数据库异常与“密码错误”区分,避免掩盖真实故障
		g.Log().Error(ctx, "登录查询用户失败", err)
		err = gerror.New("系统繁忙,请稍后再试~")
		return
	}
	if user == nil {
		// 密码错误时累计失败次数(仅当验证码正确时累计,避免验证码问题误伤正常用户)
		cool.CacheManager.Set(ctx, failKey, failCount.Int()+1, 600*time.Second)
		// 统一错误提示,避免用户名枚举
		err = gerror.New("账户或密码不正确~")
		return
	}

	// 登录成功后清除失败计数
	cool.CacheManager.Remove(ctx, failKey)

	result, err = s.generateTokenByUser(ctx, user)
	if err != nil {
		return
	}

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
	// 验证码有效期缩短为5分钟,明文仅用于服务端比对,不再输出到日志,避免验证码泄露
	cool.CacheManager.Set(ctx, result.CaptchaId, captchaText, 300*time.Second)
	return result, err
}

// Logout 退出登录
func (*BaseSysLoginService) Logout(ctx context.Context) (err error) {
	userId := cool.GetAdmin(ctx).UserId
	cool.CacheManager.Remove(ctx, "admin:department:"+gconv.String(userId))
	cool.CacheManager.Remove(ctx, "admin:perms:"+gconv.String(userId))
	cool.CacheManager.Remove(ctx, "admin:token:"+gconv.String(userId))
	cool.CacheManager.Remove(ctx, "admin:token:refresh:"+gconv.String(userId))
	return
}

// RefreshToken 刷新token
func (s *BaseSysLoginService) RefreshToken(ctx context.Context, token string) (result *TokenResult, err error) {

	tokenClaims, err := jwt.ParseWithClaims(token, &cool.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.Jwt.Secret), nil
	})
	if err != nil {
		return
	}
	claims, ok := tokenClaims.Claims.(*cool.Claims)
	if !ok {
		err = gerror.New("tokenClaims.Claims.(*Claims) error")
		return
	}
	if !tokenClaims.Valid {
		err = gerror.New("tokenClaims.Valid error")
		return
	}
	if !claims.IsRefresh {
		err = gerror.New("claims.IsRefresh error")
		return
	}

	if !(claims.UserId > 0) {
		err = gerror.New("claims.UserId error")
		return
	}

	// 校验提交的refreshToken与签发时缓存的一致,防止旧refreshToken复用/盗用
	cached, _ := cool.CacheManager.Get(ctx, "admin:token:refresh:"+gconv.String(claims.UserId))
	if cached.IsNil() || cached.String() != token {
		err = gerror.New("refreshToken已失效,请重新登录~")
		return
	}

	var (
		user        *model.BaseSysUser
		baseSysUser = model.NewBaseSysUser()
	)
	if err = cool.DBM(baseSysUser).Where("id=?", claims.UserId).Where("status=?", 1).Scan(&user); err != nil {
		g.Log().Error(ctx, "刷新Token查询用户失败", err)
		err = gerror.New("系统繁忙,请稍后再试~")
		return
	}
	if user == nil {
		err = gerror.New("用户不存在")
		return
	}

	// 确认用户密码版本与当前一致,防止改密后的旧refreshToken继续换取token
	passwordV, _ := cool.CacheManager.Get(ctx, "admin:passwordVersion:"+gconv.String(claims.UserId))
	if passwordV.IsNil() || user.PasswordV == nil || passwordV.Int32() != *user.PasswordV {
		err = gerror.New("refreshToken已失效,请重新登录~")
		return
	}

	result, err = s.generateTokenByUser(ctx, user)
	return
}

// generateToken  生成token
func (*BaseSysLoginService) generateToken(ctx context.Context, user *model.BaseSysUser, roleIds []string, exprire uint, isRefresh bool) (token string) {
	err := cool.CacheManager.Set(ctx, "admin:passwordVersion:"+gconv.String(user.ID), gconv.String(user.PasswordV), 0)
	if err != nil {
		g.Log().Error(ctx, "生成token失败", err)
	}

	claims := &cool.Claims{
		IsRefresh:       isRefresh,
		RoleIds:         roleIds,
		Username:        user.Username,
		UserId:          user.ID,
		PasswordVersion: user.PasswordV,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(exprire) * time.Second)),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = tokenClaims.SignedString([]byte(config.Config.Jwt.Secret))
	if err != nil {
		g.Log().Error(ctx, "生成token失败", err)
	}
	return
}

// 根据用户生成前端需要的Token信息
func (s *BaseSysLoginService) generateTokenByUser(ctx context.Context, user *model.BaseSysUser) (result *TokenResult, err error) {
	var (
		baseSysRoleService       = NewBaseSysRoleService()
		baseSysMenuService       = NewBaseSysMenuService()
		baseSysDepartmentService = NewBaseSysDepartmentService()
	)
	// 获取用户角色
	roleIds := baseSysRoleService.GetByUser(user.ID)
	// 如果没有角色，则报错
	if len(roleIds) == 0 {
		err = gerror.New("该用户未设置任何角色，无法登录~")
		return
	}

	// 生成token
	result = &TokenResult{}
	result.Expire = config.Config.Jwt.Token.Expire
	result.RefreshExpire = config.Config.Jwt.Token.RefreshExpire
	result.Token = s.generateToken(ctx, user, roleIds, result.Expire, false)
	result.RefreshToken = s.generateToken(ctx, user, roleIds, result.RefreshExpire, true)
	// 将用户相关信息保存到缓存
	perms := baseSysMenuService.GetPerms(roleIds)
	departments := baseSysDepartmentService.GetByRoleIds(roleIds, user.Username == "admin")
	cool.CacheManager.Set(ctx, "admin:department:"+gconv.String(user.ID), departments, 0)
	cool.CacheManager.Set(ctx, "admin:perms:"+gconv.String(user.ID), perms, 0)
	cool.CacheManager.Set(ctx, "admin:token:"+gconv.String(user.ID), result.Token, 0)
	cool.CacheManager.Set(ctx, "admin:token:refresh:"+gconv.String(user.ID), result.RefreshToken, 0)

	return
}

// NewBaseSysLoginService 创建一个新的BaseSysLoginService
func NewBaseSysLoginService() *BaseSysLoginService {
	return &BaseSysLoginService{}
}
