package middleware

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/config"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v4"
)

// 本类接口无需权限验证
func BaseAuthorityMiddlewareOpen(r *ghttp.Request) {
	r.SetCtxVar("AuthOpen", true)
	r.Middleware.Next()
}

// 本类接口无需权限验证,只需登录验证
func BaseAuthorityMiddlewareComm(r *ghttp.Request) {
	r.SetCtxVar("AuthComm", true)
	r.Middleware.Next()
}

// 其余接口需登录验证同时需要权限验证
func BaseAuthorityMiddleware(r *ghttp.Request) {
	// g.Dump(r)
	// g.Dump(r.GetHeader("Authorization"))
	var (
		ctx = r.GetCtx()
	)
	url := r.URL.String()

	// 无需登录验证
	AuthOpen := r.GetCtxVar("AuthOpen", false)
	if AuthOpen.Bool() {
		r.Middleware.Next()
		return
	}

	tokenString := r.GetHeader("Authorization")
	token, err := jwt.ParseWithClaims(tokenString, &cool.Claims{}, func(token *jwt.Token) (interface{}, error) {

		return []byte(config.Config.Jwt.Secret), nil
	})
	if err != nil {
		g.Log().Error(ctx, "BaseAuthorityMiddleware", err)
		r.Response.WriteStatusExit(401, g.Map{
			"code":    1001,
			"message": "登陆失效～",
		})
		return
	}
	if token == nil || !token.Valid {
		g.Log().Error(ctx, "BaseAuthorityMiddleware", "token invalid")
		r.Response.WriteStatusExit(401, g.Map{
			"code":    1001,
			"message": "登陆失效～",
		})
		return
	}
	admin, ok := token.Claims.(*cool.Claims)
	if !ok || admin == nil || admin.UserId == 0 {
		g.Log().Error(ctx, "BaseAuthorityMiddleware", "claims invalid")
		r.Response.WriteStatusExit(401, g.Map{
			"code":    1001,
			"message": "登陆失效～",
		})
		return
	}
	// 将用户信息放入上下文
	r.SetCtxVar("admin", admin)

	cachetoken, _ := cool.CacheManager.Get(ctx, "admin:token:"+gconv.String(admin.UserId))
	rtoken := cachetoken.String()
	// 只验证登录不验证权限的接口(comm组):持有有效token即可访问
	AuthComm := r.GetCtxVar("AuthComm", false)
	if AuthComm.Bool() {
		r.Middleware.Next()
		return
	}
	// refreshToken不能作为访问令牌
	if admin.IsRefresh {
		g.Log().Error(ctx, "BaseAuthorityMiddleware", "token invalid")
		r.Response.WriteStatusExit(401, g.Map{
			"code":    1001,
			"message": "登陆失效～",
		})
		return
	}
	// 判断密码版本是否正确(改密后旧token立即失效)
	if admin.PasswordVersion == nil {
		g.Log().Error(ctx, "BaseAuthorityMiddleware", "passwordV invalid")
		r.Response.WriteStatusExit(401, g.Map{
			"code":    1001,
			"message": "登陆失效～",
		})
		return
	}
	passwordV, _ := cool.CacheManager.Get(ctx, "admin:passwordVersion:"+gconv.String(admin.UserId))
	if passwordV.Int32() != *admin.PasswordVersion {
		g.Log().Error(ctx, "BaseAuthorityMiddleware", "passwordV invalid")
		r.Response.WriteStatusExit(401, g.Map{
			"code":    1001,
			"message": "登陆失效～",
		})
		return
	}
	// 会话令牌缓存缺失(已登出/缓存丢失如重启)需重新登录——超管与普通用户规则一致
	if rtoken == "" {
		g.Log().Error(ctx, "BaseAuthorityMiddleware", "rtoken invalid")
		r.Response.WriteStatusExit(401, g.Map{
			"code":    1001,
			"message": "登陆失效～",
		})
		return
	}
	// 开启SSO时,仅最新签发的token有效(多端互踢)
	if tokenString != rtoken && config.Config.Jwt.Sso {
		g.Log().Error(ctx, "BaseAuthorityMiddleware", "token invalid")
		r.Response.WriteStatusExit(401, g.Map{
			"code":    1001,
			"message": "登陆失效～",
		})
		return
	}
	// 超管拥有所有权限(免perms校验,上述会话校验与普通用户保持一致)
	if admin.UserId == 1 {
		r.Middleware.Next()
		return
	}
	// 从缓存获取perms
	permsCache, _ := cool.CacheManager.Get(ctx, "admin:perms:"+gconv.String(admin.UserId))
	// 转换为数组
	permsVar := permsCache.Strings()
	// 转换为garray
	perms := garray.NewStrArrayFrom(permsVar)
	// 如果perms为空
	if perms.Len() == 0 {
		g.Log().Error(ctx, "BaseAuthorityMiddleware", "perms invalid")
		r.Response.WriteStatusExit(403, g.Map{
			"code":    1001,
			"message": "登录失效或无权限访问~",
		})
		return
	}
	// 去除url后面的参数，使用字符串分割方法，若长度等于2，则说明有参数，则我们将改写url值进行权限比对
	parts := gstr.Split(url, "?")
	if len(parts) == 2 {
		url = parts[0]
	}
	//url 转换为数组
	urls := gstr.Split(url, "/")
	// 去除第一个空字符串和admin
	urls = urls[2:]
	// 以冒号连接成新字符串url
	url = gstr.Join(urls, ":")
	// 如果perms中不包含url 则无权限
	if !perms.ContainsI(url) {
		g.Log().Error(ctx, "BaseAuthorityMiddleware", "perms invalid")
		r.Response.WriteStatusExit(403, g.Map{
			"code":    1001,
			"message": "登录失效或无权限访问~",
		})
		return
	}
	// 上面写逻辑
	r.Middleware.Next()

}
