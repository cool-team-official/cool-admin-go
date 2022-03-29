package controller

import (
	"context"

	v1 "cool-admin-go/api/v1"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Welcome = cWelcome{}
)

type cWelcome struct{}

func (c *cWelcome) Welcome(ctx context.Context, req *v1.WelcomeReq) (res *v1.WelcomeRes, err error) {

	g.RequestFromCtx(ctx).Response.WriteTpl("welcome.html", g.Map{
		"text": "HELLO COOL-ADMIN 5.x 一个项目只用COOL就够了！！！",
	})
	return
}
