package model

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/v2/os/gtime"
)

const TableNameTaskInfo = "task_info"

// TaskInfo mapped from table <task_info>
type TaskInfo struct {
	*cool.Model
	JobId       string     `json:"jobId" gorm:"column:jobId;type:varchar(255);comment:'任务ID'"`
	RepeatConf  string     `json:"repeatConf" gorm:"column:repeatConf;comment:'重复配置'"`
	Name        string     `json:"name" gorm:"column:name;type:varchar(255);comment:'任务名称'"`
	Cron        string     `json:"cron" gorm:"column:cron;type:varchar(255);comment:'cron表达式'"`
	Limit       int        `json:"limit" gorm:"column:limit;type:int(11);comment:'限制次数 不传为不限制'"`
	Every       int        `json:"every" gorm:"column:every;type:int(11);comment:'间隔时间 单位秒'"`
	Remark      string     `json:"remark" gorm:"column:remark;type:varchar(255);comment:'备注'"`
	Status      int        `json:"status" gorm:"column:status;type:int(11);comment:'状态 0:关闭 1:开启'"`
	StartDate   gtime.Time `json:"startDate" gorm:"column:startDate;type:datetime;comment:'开始时间'"`
	EndDate     gtime.Time `json:"endDate" gorm:"column:endDate;type:datetime;comment:'结束时间'"`
	Data        string     `json:"data" gorm:"column:data;type:varchar(255);comment:'数据'"`
	Service     string     `json:"service" gorm:"column:service;type:varchar(255);comment:'执行service的实例ID'"`
	Type        int        `json:"type" gorm:"column:type;type:int(11);comment:'类型 0:系统 1:用户'"`
	NextRunTime gtime.Time `json:"nextRunTime" gorm:"column:nextRunTime;type:datetime;comment:'下次执行时间'"`
	TaskType    int        `json:"taskType" gorm:"column:taskType;type:int(11);comment:'状态 0:cron 1:时间间隔'"`
}

// TableName TaskInfo's table name
func (*TaskInfo) TableName() string {
	return TableNameTaskInfo
}

// GroupName TaskInfo's table group
func (*TaskInfo) GroupName() string {
	return "default"
}

// NewTaskInfo create a new TaskInfo
func NewTaskInfo() *TaskInfo {
	return &TaskInfo{
		Model: cool.NewModel(),
	}
}

// init 创建表
func init() {
	cool.CreateTable(&TaskInfo{})
}
