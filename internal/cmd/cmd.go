package cmd

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool"
	i18n "github.com/cool-team-official/cool-admin-go/modules/base/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// g.Dump(g.DB("test").GetConfig())
			if cool.IsRedisMode {
				go cool.ListenFunc(ctx)
			}

			s := g.Server()

			// 如果存在 data/cool-admin-vue/dist 目录，则设置为主目录
			if gfile.IsDir("frontend/dist") {
				s.SetServerRoot("frontend/dist")
			}
			// i18n 信息
			s.BindHandler("/i18n", i18n.I18nInfo)
			s.Run()
			return nil
		},
	}
)
