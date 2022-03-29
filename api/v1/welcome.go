package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type WelcomeReq struct {
	g.Meta `path:"/" tags:"Welcome" method:"get" summary:"首页"`
}
type WelcomeRes struct {
	g.Meta `mime:"text/html" example:"string111"`
}
