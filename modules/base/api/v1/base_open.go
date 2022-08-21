package v1

import "github.com/gogf/gf/v2/frame/g"

// login 接口请求
type BaseOpenLoginReq struct {
	g.Meta     `path:"/login" method:"POST"`
	Username   string `json:"username" p:"username" v:"required"`
	Password   string `json:"password" p:"password" v:"required"`
	CaptchaId  string `json:"captchaId" p:"captchaId" v:"required"`
	VerifyCode string `json:"verifyCode" p:"verifyCode" v:"required"`
}

// captcha 验证码接口

type BaseOpenCaptchaReq struct {
	g.Meta `path:"/captcha" method:"GET"`
	Height int `json:"height" in:"query" default:"40"`
	Width  int `json:"width"  in:"query" default:"150"`
}
