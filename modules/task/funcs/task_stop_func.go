package funcs

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/task/model"
	"github.com/cool-team-official/cool-admin-go/modules/task/service"
	"github.com/gogf/gf/v2/frame/g"
)

type TaskStopFunc struct {
}

func (t *TaskStopFunc) Func(ctx g.Ctx, id string) error {
	taskInfo := model.NewTaskInfo()
	_, err := cool.DBM(taskInfo).Where("id = ?", id).Update(g.Map{"status": 0})
	if err != nil {
		return err
	}

	return service.DisableTask(ctx, id)
}

func (t *TaskStopFunc) IsSingleton() bool {
	return false
}

func (t *TaskStopFunc) IsAllWorker() bool {
	return true
}

func init() {
	cool.RegisterFunc("TaskStopFunc", &TaskStopFunc{})
}
