package service

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/modules/task/model"
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
