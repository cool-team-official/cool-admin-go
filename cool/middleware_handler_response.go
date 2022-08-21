// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package cool

import (
	"net/http"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

// DefaultHandlerResponse is the default implementation of HandlerResponse.
type DefaultHandlerResponse struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"message" dc:"Error message"`
	Data    interface{} `json:"data,omitempty"    dc:"Result data for certain request according API definition"`
}

// MiddlewareHandlerResponse is the default middleware handling handler response object and its error.
func MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		// ctx  g.Ctx
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err).Code()
	)

	// g.Log().Debug(ctx, code, msg, res)
	msg = "success"

	if err != nil {
		if code == -1 {
			code = 1001
		}
		msg = err.Error()
	} else if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
		msg = http.StatusText(r.Response.Status)
		switch r.Response.Status {
		case http.StatusNotFound:
			code = 404
		case http.StatusForbidden:
			code = 403
		default:
			code = 500
		}
	} else {
		// code = gcode.CodeOK
		code = 1000
	}
	// 做一些code转换适配cooladmin的错误码
	switch code {
	case 51: // 参数错误
		code = 51
	case 50: // 内部错误
		code = 1003
	default:
	}
	// g.Log().Debug(ctx, code, msg, res)
	// 如果是正常返回，直接返回res
	if code == 1000 && r.Response.Status == 200 {
		r.Response.WriteJsonExit(res)
	}
	r.Response.WriteJson(DefaultHandlerResponse{
		Code:    code,
		Message: msg,
		Data:    res,
	})
}
