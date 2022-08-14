// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package cool

import (
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// DefaultHandlerResponse is the default implementation of HandlerResponse.
type DefaultHandlerResponse struct {
	Code    int    `json:"code"    dc:"Error code"`
	Message string `json:"message" dc:"Error message"`
	// Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}

// MiddlewareHandlerResponse is the default middleware handling handler response object and its error.
func MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
		msg = http.StatusText(r.Response.Status)
		switch r.Response.Status {
		case http.StatusNotFound:
			code = gcode.CodeNotFound
		case http.StatusForbidden:
			code = gcode.CodeNotAuthorized
		default:
			code = gcode.CodeUnknown
		}
	} else {
		code = gcode.CodeOK
	}
	// var ctx = r.Context()
	// g.Log().Debug(ctx, "[HandlerResponse]", "code", code, "message", msg, "data", res)

	coderes := &DefaultHandlerResponse{
		Code:    code.Code(),
		Message: msg,
	}

	type sData struct {
		Data interface{} `json:"data"`
	}
	var data *sData
	gconv.Struct(res, &data)
	type sRes struct {
		DefaultHandlerResponse
		sData
	}
	r.Response.WriteJson(sRes{
		DefaultHandlerResponse: *coderes,
		sData:                  *data,
	})

}
