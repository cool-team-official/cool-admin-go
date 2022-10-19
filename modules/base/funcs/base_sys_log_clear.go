package funcs

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/base/service"
	"github.com/gogf/gf/v2/frame/g"
)

type BaseFuncClearLog struct {
}

// Func
func (f *BaseFuncClearLog) Func(ctx g.Ctx, param string) (err error) {
	g.Log().Info(ctx, "清理日志 BaseFuncClearLog.Func", "param", param)
	baseSysLogService := service.NewBaseSysLogService()
	if param == "true" {
		err = baseSysLogService.Clear(true)
	} else {
		err = baseSysLogService.Clear(false)
	}
	return
}

// IsSingleton
func (f *BaseFuncClearLog) IsSingleton() bool {
	return true
}

// IsAllWorker
func (f *BaseFuncClearLog) IsAllWorker() bool {
	return false
}

// init
func init() {
	cool.RegisterFunc("BaseFuncClearLog", &BaseFuncClearLog{})
}
