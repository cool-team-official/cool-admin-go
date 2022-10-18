package model

import (
	"github.com/cool-team-official/cool-admin-go/cool"
)

const TableNameTaskLog = "task_log"

// TaskLog mapped from table <task_log>
type TaskLog struct {
	*cool.Model
	TaskId uint64 `gorm:"column:taskId;comment:任务ID" json:"taskId"`
	Status uint8  `gorm:"column:status;not null;comment:状态 0:失败 1:成功" json:"status"`
	Detail string `gorm:"column:detail;comment:详情" json:"detail"`
}

// TableName TaskLog's table name
func (*TaskLog) TableName() string {
	return TableNameTaskLog
}

// GroupName TaskLog's table group
func (*TaskLog) GroupName() string {
	return "default"
}

// NewTaskLog create a new TaskLog
func NewTaskLog() *TaskLog {
	return &TaskLog{
		Model: cool.NewModel(),
	}
}

// init 创建表
func init() {
	cool.CreateTable(&TaskLog{})
}
