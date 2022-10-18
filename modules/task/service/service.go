package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
)

// EnableTask 启用任务
func EnableTask(ctx g.Ctx, cornId string, funcstring string, cron string, startDate string) (err error) {
	funcName := gstr.SubStr(funcstring, 0, gstr.Pos(funcstring, "("))
	g.Log().Debug(ctx, "启用任务", cornId, funcName, cron)
	if _, ok := cool.FuncMap[funcName]; !ok {
		err = gerror.New("函数不存在" + funcName)
		return
	}
	taskInfoService := NewTaskInfoService()

	if cool.FuncMap[funcName].IsSingleton() {
		gcron.Remove(cornId)
		_, err = gcron.AddSingleton(ctx, cron, func(ctx g.Ctx) {
			nowDate := gtime.Now().Format("Y-m-d H:i:s")
			if nowDate < startDate {
				g.Log().Debug(ctx, "当前时间小于启用时间, 不执行单例函数", funcName)
				return
			}
			err := cool.RunFunc(ctx, funcstring)
			if err != nil {
				g.Log().Error(ctx, err)
				taskInfoService.Record(ctx, cornId, 0, err.Error())
			} else {
				taskInfoService.Record(ctx, cornId, 1, "任务执行成功")
			}
		}, cornId)
	} else {
		gcron.Remove(cornId)
		_, err = gcron.Add(ctx, cron, func(ctx g.Ctx) {
			nowDate := gtime.Now().Format("Y-m-d H:i:s")
			if nowDate < startDate {
				g.Log().Debug(ctx, "当前时间小于启用时间, 不执行函数", funcName)
				return
			}
			err := cool.RunFunc(ctx, funcstring)
			if err != nil {
				g.Log().Error(ctx, err)
				taskInfoService.Record(ctx, cornId, 0, err.Error())
			} else {
				taskInfoService.Record(ctx, cornId, 1, "")
			}
		}, cornId)
	}
	return
}

// DisableTask 禁用任务
func DisableTask(ctx g.Ctx, cornId string) (err error) {
	gcron.Remove(cornId)
	return
}
