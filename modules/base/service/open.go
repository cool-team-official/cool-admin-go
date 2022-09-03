package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
)

// login 登录接口
// func Login(username string, password string, captchaId string, verifyCode string) (data interface{}, err error) {
// 	var (
// 		ctx                g.Ctx
// 		baseSysRoleService = NewBaseSysRoleService()
// 	)
// 	vcode, _ := cool.Cache.Get(ctx, captchaId)
// 	if vcode.String() != verifyCode {
// 		err = gerror.New("验证码错误")
// 		// return
// 	}
// 	md5password, _ := gmd5.Encrypt(password)
// 	baseSysUser := model.NewBaseSysUser()

// 	var user *model.BaseSysUser
// 	g.DB(baseSysUser.GroupName()).Model(baseSysUser.TableName()).Where("username=?", username).Where("password=?", md5password).Where("status=?", 1).Scan(&user)
// 	if user == nil {
// 		err = gerror.New("账户或密码不正确~")
// 		return
// 	}
// 	// 获取用户角色
// 	roleIds := baseSysRoleService.GetByUser(user.ID)
// 	// 如果没有角色，则报错
// 	if len(roleIds) == 0 {
// 		err = gerror.New("该用户未设置任何角色，无法登录~")
// 		return
// 	}

// 	datam := gmap.New()
// 	datam.Sets(g.MapAnyAny{
// 		"refreshExpires": time.Now().Add(time.Second * 1800),
// 		"refreshToken":   guid.S(),
// 		"token":          "1231231",
// 		"expires":        time.Now().Add(time.Second * 1800),
// 	})
// 	data = datam
// 	return
// }

// 验证码接口
// func Captcha(width int, height int) (captchaID string, data string) {
// 	var (
// 		ctx         g.Ctx
// 		captchaText string
// 	)
// 	captchaText = grand.Digits(4)
// 	svg := `<svg width="150" height="50" xmlns="http://www.w3.org/2000/svg"><text x="75" y="25" text-anchor="middle" font-size="25" fill="#fff">` + captchaText + `</text></svg>`
// 	// g.Dump(svg)
// 	svgbase64 := gbase64.EncodeString(svg)
// 	// g.Dump(svgbase64)
// 	data = `data:image/svg+xml;base64,` + svgbase64
// 	captchaID = guid.S()
// 	cool.Cache.Set(ctx, captchaID, captchaText, 1800*time.Second)
// 	g.Log().Debug(ctx, "验证码", captchaID, captchaText)
// 	return
// }

// 返回AdminEPS信息,缓存部分还未写完
func AdminEps() (eps interface{}) {
	var ctx g.Ctx
	// var err error
	c := cool.CacheEPS
	epskey, err := c.Get(ctx, "eps")
	if err != nil {
		g.Log().Debug(ctx, "EPS", err)
		return
	}
	// g.DumpWithType("epskey:", epskey, epskey.String())
	if epskey.IsEmpty() {
		g.Log().Debug(ctx, "eps缓存未命中")
		eps = creatAdminEps()
		c.Set(ctx, "eps", eps, 0)
	} else {
		g.Log().Debug(ctx, "eps缓存命中")
		eps = epskey
	}

	return
}

// 生成AdminEPS信息
func creatAdminEps() (eps *gjson.Json) {
	var (
		ctx = gctx.New()
		// err    error
		method string
	)

	eps = gjson.New(`{}`)

	g.Log().Debug(ctx, "AdminEPS")

	// 清除数据库旧数据
	res, _ := g.Model("base_eps").Where("1=1").Delete()
	g.Log().Debug(ctx, "清除旧数据res", res)

	routers := g.Server().GetRoutes()
	// g.Dump(routers)
	for _, item := range routers {
		switch item.Type {
		case ghttp.HandlerTypeMiddleware, ghttp.HandlerTypeHook:
			continue
		}
		method = item.Method
		if gstr.Equal(method, "ALL") {
			method = ""
		}
		uri := item.Handler.Router.Uri
		uriSplited := gstr.Split(uri, "/")
		// g.Dump(method, uri, uriSplited)
		// method "GET"  uri "/admin/demo/sample/welcome"
		if len(uriSplited) < 5 {
			continue
		}
		if uriSplited[1] != "admin" {
			continue
		}
		// moduleName := uriSplited[2]
		module := uriSplited[2]
		path := "/" + uriSplited[len(uriSplited)-1]
		// 去除uri/Splited最后一个元素
		uriSplited = uriSplited[:len(uriSplited)-1]
		// 重新以 / 组合为prefix
		prefix := gstr.Join(uriSplited, "/")
		api := &sGroupApi{
			Module:  module,
			Path:    path,
			Prefix:  prefix,
			Method:  method,
			Summary: uri,
			Tag:     "",
			Dts:     "",
		}
		g.Model("base_eps").Insert(api)

	}
	// 查询获取module,perfix列表
	modules, _ := g.Model("base_eps").Fields("DISTINCT module,prefix").All()
	for _, item := range modules {
		module := item["module"].String()
		prefix := item["prefix"].String()
		apis, _ := g.Model("base_eps").Where("module=? AND prefix=?", module, prefix).All()
		cGroup := &sCGroup{
			Module: module,
			Prefix: prefix,
		}
		cGroupJson := gjson.New(cGroup)

		cGroupJson.Set("api", apis)
		eps.Append(module, cGroupJson)
	}
	return
}

// sCGroup struct 路由组信息
type sCGroup struct {
	Module  string   `json:"module"`
	Prefix  string   `json:"prefix"`
	Columns []string `json:"columns"`
	// api     []*sGroupApi // 接口列表 例如：[{},{},{}]

}

// sGroupApi struct 路由信息
type sGroupApi struct {
	Module  string // 所属模块名称 例如：base
	Method  string // 请求方法 例如：GET
	Path    string // 请求路径 例如：/welcome
	Prefix  string // 路由前缀 例如：/admin/base/open
	Summary string // 描述 例如：欢迎页面
	Tag     string // 标签 例如：base  好像暂时不用
	Dts     string // 未知 例如：{} 好像暂时不用
}
