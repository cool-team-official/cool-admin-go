package demo

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	_ "github.com/cool-team-official/cool-admin-go/modules/task/controller"
	_ "github.com/cool-team-official/cool-admin-go/modules/task/funcs"
	_ "github.com/cool-team-official/cool-admin-go/modules/task/middleware"
	"github.com/cool-team-official/cool-admin-go/modules/task/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {
	var (
		taskInfo = model.NewTaskInfo()
		ctx      = gctx.GetInitCtx()
	)
	g.Log().Debug(ctx, "modules/task init")
	result, err := cool.DBM(taskInfo).Where("status = ?", 1).All()
	if err != nil {
		panic(err)
	}
	for _, v := range result {
		id := v["id"].String()
		cool.RunFunc(ctx, "TaskAddTask("+id+")")
	}

}
