package v1

import "github.com/gogf/gf/v2/frame/g"

// BaseCommPersonReq 接口请求参数
type BaseCommPersonReq struct {
	g.Meta        `path:"/person" method:"GET"`
	Authorization string `json:"Authorization" in:"header"`
}

// BaseCommPremmenu 接口请求参数
type BaseCommPermmenuReq struct {
	g.Meta        `path:"/permmenu" method:"GET"`
	Authorization string `json:"Authorization" in:"header"`
}
