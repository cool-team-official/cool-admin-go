package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/task/model"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type TaskInfoService struct {
	*cool.Service
}

func NewTaskInfoService() *TaskInfoService {
	return &TaskInfoService{
		&cool.Service{
			Model: model.NewTaskInfo(),
			PageQueryOp: &cool.QueryOp{
				FieldEQ: []string{"status", "type"},
			},
			UniqueKey: map[string]string{
				"name": "任务名称不能重复",
			},
		},
	}
}

func (s *TaskInfoService) ModifyAfter(ctx g.Ctx, method string, param g.MapStrAny) (err error) {
	g.Log().Info(ctx, "TaskInfoService.ModifyAfter", method, param)
	if method == "Add" {
		if gconv.Int(param["status"]) == 1 {
			id, err := cool.DBM(s.Model).Where("name = ?", param["name"]).Value("id")
			if err != nil {
				return err
			}
			return cool.ClusterRunFunc(ctx, "TaskAddTask("+id.String()+")")
		}
	}
	if method == "Update" {
		id := gconv.String(param["id"])
		if gconv.Int(param["status"]) == 1 {
			return cool.ClusterRunFunc(ctx, "TaskStartFunc("+id+")")
		} else {
			return cool.ClusterRunFunc(ctx, "TaskStopFunc("+id+")")
		}
	}
	if method == "Delete" {
		id := gconv.String(param["id"])
		return cool.ClusterRunFunc(ctx, "TaskStopFunc("+id+")")
	}
	return nil
}

// Record 保存任务记录,成功任务每个任务保留最新20条日志,失败日志不会删除
func (s *TaskInfoService) Record(ctx g.Ctx, id string, status int, detail string) error {
	taskLog := model.NewTaskLog()
	_, err := cool.DBM(taskLog).Data(g.Map{
		"taskId": id,
		"status": status,
		"detail": detail,
	}).Insert()
	if err != nil {
		return err
	}
	if status == 1 {
		// _, err = cool.DBM(taskLog).Where("taskId = ?", id).Where("status", 1).Limit(20).Delete()
		record, err := cool.DBM(taskLog).Where("taskId = ?", id).Where("status", 1).Order("id", "desc").Offset(19).One()
		if err != nil {
			return err
		}
		if record.IsEmpty() {
			return nil
		}
		minId := record["id"].Int()
		g.DumpWithType(minId)
		if err != nil {
			return err
		}
		g.Log().Info(ctx, "minId", minId)
		_, err = cool.DBM(taskLog).Where("taskId = ?", id).Where("status", 1).Where("id < ?", minId).Delete()
		if err != nil {
			return err
		}
	}
	return err
}

// Once 执行一次任务
func (s *TaskInfoService) Once(ctx g.Ctx, id int64) error {
	record, err := cool.DBM(s.Model).Where("id = ?", id).One()
	if err != nil {
		return err
	}
	if record.IsEmpty() {
		return gerror.New("任务不存在")
	}
	funcString := record["service"].String()
	return cool.ClusterRunFunc(ctx, funcString)
}
