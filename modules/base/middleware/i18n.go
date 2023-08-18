package middleware

import (
	"strings"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func AutoI18n(r *ghttp.Request) {
	Language := r.GetHeader("Accept-Language")
	Language = strings.Split(Language, ",")[0]
	cool.I18n.SetLanguage(Language)
	r.Response.Header().Set("Content-Language", Language)
	r.Middleware.Next()
}

func I18nInfo(r *ghttp.Request) {
	r.Response.WriteJson(g.Map{
		r.Response.Header().Get("Content-Language"): cool.I18n.Translate(r.Context(), "BaseResMessage"),
	})
}
