package funcs

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/task/model"
	"github.com/cool-team-official/cool-admin-go/modules/task/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type TaskAddTask struct {
}

func (t *TaskAddTask) Func(ctx g.Ctx, id string) error {
	taskInfo := model.NewTaskInfo()
	result, err := cool.DBM(taskInfo).Where("id = ?", id).One()
	if err != nil {
		return err
	}
	if result["taskType"].Int() == 1 {
		every := result["every"].Uint() / 1000
		cron := "@every " + gconv.String(every) + "s"
		funcstring := result["service"].String()
		startDate := result["startDate"].String()
		err = service.EnableTask(ctx, id, funcstring, cron, startDate)
	} else {
		cron := result["cron"].String()
		funcstring := result["service"].String()
		startDate := result["startDate"].String()
		err = service.EnableTask(ctx, id, funcstring, cron, startDate)
	}

	return err
}

func (t *TaskAddTask) IsSingleton() bool {
	return false
}

func (t *TaskAddTask) IsAllWorker() bool {
	return true
}

func init() {
	cool.RegisterFunc("TaskAddTask", &TaskAddTask{})
}
