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
		statusCode = 200
		ctx        = r.GetCtx()
	)
	url := r.URL.String()

	// 无需登录验证
	AuthOpen := r.GetCtxVar("AuthOpen", false)
	if AuthOpen.Bool() {
		g.Log().Debug(ctx, "BaseAuthorityMiddleware", "AuthOpen", url)
		r.Middleware.Next()
		return
	}
	type Claims struct {
		IsRefresh       bool   `json:"isRefresh"`
		RoleIds         []uint `json:"roleIds"`
		Username        string `json:"username"`
		UserId          uint   `json:"userId"`
		PasswordVersion *int32 `json:"passwordVersion"`
		jwt.RegisteredClaims
	}
	tokenString := r.GetHeader("Authorization")
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {

		return []byte(config.Config.Jwt.Secret), nil
	})
	if err != nil {
		g.Log().Error(ctx, "BaseAuthorityMiddleware", err)
		statusCode = 401
		r.Response.WriteStatusExit(statusCode, g.Map{
			"code":    1001,
			"message": "登陆失效～",
		})
	}
	if !token.Valid {
		g.Log().Error(ctx, "BaseAuthorityMiddleware", "token invalid")
		statusCode = 401
		r.Response.WriteStatusExit(statusCode, g.Map{
			"code":    1001,
			"message": "登陆失效～",
		})
	}
	admin := token.Claims.(*Claims)
	// 将用户信息放入上下文
	r.SetCtxVar("admin", admin)

	cachetoken, _ := cool.CacheManager.Get(ctx, "admin:token:"+gconv.String(admin.UserId))
	rtoken := cachetoken.String()
	// 超管拥有所有权限
	if admin.UserId == 1 && !admin.IsRefresh {
		if tokenString != rtoken && config.Config.Jwt.Sso {
			g.Log().Error(ctx, "BaseAuthorityMiddleware", "token invalid")
			statusCode = 401
			r.Response.WriteStatusExit(statusCode, g.Map{
				"code":    1001,
				"message": "登陆失效～",
			})
		} else {
			r.Middleware.Next()
			return
		}
	}
	// 只验证登录不验证权限的接口
	AuthComm := r.GetCtxVar("AuthComm", false)
	if AuthComm.Bool() {
		g.Log().Debug(ctx, "BaseAuthorityMiddleware", "AuthComm", url)
		r.Middleware.Next()
		return
	}
	// 如果传的token是refreshToken则校验失败
	if admin.IsRefresh {
		g.Log().Error(ctx, "BaseAuthorityMiddleware", "token invalid")
		statusCode = 401
		r.Response.WriteStatusExit(statusCode, g.Map{
			"code":    1001,
			"message": "登陆失效～",
		})
	}
	// 判断密码版本是否正确
	passwordV, _ := cool.CacheManager.Get(ctx, "admin:passwordVersion:"+gconv.String(admin.UserId))
	if passwordV.Int32() != *admin.PasswordVersion {
		g.Log().Error(ctx, "BaseAuthorityMiddleware", "passwordV invalid")
		statusCode = 401
		r.Response.WriteStatusExit(statusCode, g.Map{
			"code":    1001,
			"message": "登陆失效～",
		})
	}
	// 如果rtoken为空
	if rtoken == "" {
		g.Log().Error(ctx, "BaseAuthorityMiddleware", "rtoken invalid")
		statusCode = 401
		r.Response.WriteStatusExit(statusCode, g.Map{
			"code":    1001,
			"message": "登陆失效～",
		})
	}
	// 如果rtoken不等于token 且 sso 未开启
	if tokenString != rtoken && !config.Config.Jwt.Sso {
		g.Log().Error(ctx, "BaseAuthorityMiddleware", "token invalid")
		statusCode = 401
		r.Response.WriteStatusExit(statusCode, g.Map{
			"code":    1001,
			"message": "登陆失效～",
		})
	}
	// 从缓存获取perms
	permsCache, _ := cool.CacheManager.Get(ctx, "admin:perms:"+gconv.String(admin.UserId))
	// 转换为数组
	permsVar := permsCache.Strings()
	// 转换为garray
	perms := garray.NewStrArrayFrom(permsVar)
	// 如果perms为空 为无权限
	if perms.Len() == 0 {
		g.Log().Error(ctx, "BaseAuthorityMiddleware", "perms invalid")
		statusCode = 403
		r.Response.WriteStatusExit(statusCode, g.Map{
			"code":    1001,
			"message": "登录失效或无权限访问~",
		})
	}
	// url去除参数部分
	if gstr.Pos("?", url) > 0 {
		url = gstr.StrTillEx(url, "?")
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
		statusCode = 403
		r.Response.WriteStatusExit(statusCode, g.Map{
			"code":    1001,
			"message": "登录失效或无权限访问~",
		})
	}
	// 上面写逻辑
	r.Middleware.Next()

}
