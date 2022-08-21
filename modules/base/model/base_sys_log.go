package model

import "github.com/cool-team-official/cool-admin-go/cool"

const TableNameBaseSysLog = "base_sys_log"

// BaseSysLog mapped from table <base_sys_log>
type BaseSysLog struct {
	*cool.Model
	UserID *int64  `gorm:"column:userId;type:bigint;index:IDX_51a2caeb5713efdfcb343a8772,priority:1" json:"userId"`                // 用户ID
	Action string  `gorm:"column:action;type:varchar(100);not null;index:IDX_938f886fb40e163db174b7f6c3,priority:1" json:"action"` // 行为
	IP     *string `gorm:"column:ip;type:varchar(50);index:IDX_24e18767659f8c7142580893f2,priority:1" json:"ip"`                   // ip
	IPAddr *string `gorm:"column:ipAddr;type:varchar(50);index:IDX_a03a27f75cf8d502b3060823e1,priority:1" json:"ipAddr"`           // ip地址
	Params *string `gorm:"column:params;type:text" json:"params"`                                                                  // 参数
}

// TableName BaseSysLog's table name
func (*BaseSysLog) TableName() string {
	return TableNameBaseSysLog
}

// init 创建表
func init() {
	cool.CreateTable(&BaseSysLog{})
}
