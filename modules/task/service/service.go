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
func EnableTask(ctx g.Ctx, cronId string, funcstring string, cron string, startDate string) (err error) {
	funcName := gstr.SubStr(funcstring, 0, gstr.Pos(funcstring, "("))
	if _, ok := cool.FuncMap[funcName]; !ok {
		err = gerror.New("函数不存在" + funcName)
		return
	}
	taskInfoService := NewTaskInfoService()

	if cool.FuncMap[funcName].IsSingleton() {
		gcron.Remove(cronId)
		_, err = gcron.AddSingleton(ctx, cron, func(ctx g.Ctx) {
			nowDate := gtime.Now().Format("Y-m-d H:i:s")
			if nowDate < startDate {
				g.Log().Debug(ctx, "当前时间小于启用时间, 不执行单例函数", funcName)
				return
			}
			err := cool.RunFunc(ctx, funcstring)
			if err != nil {
				g.Log().Error(ctx, err)
				taskInfoService.Record(ctx, cronId, 0, err.Error())
			} else {
				taskInfoService.Record(ctx, cronId, 1, "任务执行成功")
			}
		}, cronId)
	} else {
		gcron.Remove(cronId)
		_, err = gcron.Add(ctx, cron, func(ctx g.Ctx) {
			nowDate := gtime.Now().Format("Y-m-d H:i:s")
			if nowDate < startDate {
				g.Log().Debug(ctx, "当前时间小于启用时间, 不执行函数", funcName)
				return
			}
			err := cool.RunFunc(ctx, funcstring)
			if err != nil {
				g.Log().Error(ctx, err)
				taskInfoService.Record(ctx, cronId, 0, gstr.AddSlashes(err.Error()))
			} else {
				taskInfoService.Record(ctx, cronId, 1, gstr.AddSlashes("任务执行成功"))
			}
			taskInfoService.SetNextRunTime(ctx, cronId, cron)
		}, cronId)
	}
	taskInfoService.SetNextRunTime(ctx, cronId, cron)
	return
}

// DisableTask 禁用任务
func DisableTask(ctx g.Ctx, cronId string) (err error) {
	gcron.Remove(cronId)
	return
}
