// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskInfoDao is the data access object for table task_info.
type TaskInfoDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns TaskInfoColumns // columns contains all the column names of Table for convenient usage.
}

// TaskInfoColumns defines and stores column names for table task_info.
type TaskInfoColumns struct {
	Id          string // ID
	CreateTime  string // 创建时间
	UpdateTime  string // 更新时间
	JobId       string // 任务ID
	RepeatConf  string // 任务配置
	Name        string // 名称
	Cron        string // cron
	Limit       string // 最大执行次数 不传为无限次
	Every       string // 每间隔多少毫秒执行一次 如果cron设置了 这项设置就无效
	Remark      string // 备注
	Status      string // 状态 0:停止 1：运行
	StartDate   string // 开始时间
	EndDate     string // 结束时间
	Data        string // 数据
	Service     string // 执行的service实例ID
	Type        string // 状态 0:系统 1：用户
	NextRunTime string // 下一次执行时间
	TaskType    string // 状态 0:cron 1：时间间隔
}

//  taskInfoColumns holds the columns for table task_info.
var taskInfoColumns = TaskInfoColumns{
	Id:          "id",
	CreateTime:  "createTime",
	UpdateTime:  "updateTime",
	JobId:       "jobId",
	RepeatConf:  "repeatConf",
	Name:        "name",
	Cron:        "cron",
	Limit:       "limit",
	Every:       "every",
	Remark:      "remark",
	Status:      "status",
	StartDate:   "startDate",
	EndDate:     "endDate",
	Data:        "data",
	Service:     "service",
	Type:        "type",
	NextRunTime: "nextRunTime",
	TaskType:    "taskType",
}

// NewTaskInfoDao creates and returns a new DAO object for table data access.
func NewTaskInfoDao() *TaskInfoDao {
	return &TaskInfoDao{
		group:   "default",
		table:   "task_info",
		columns: taskInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskInfoDao) Columns() TaskInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
