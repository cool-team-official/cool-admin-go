package model

import "github.com/cool-team-official/cool-admin-go/cool"

const TableNameBaseSysRole = "base_sys_role"

// BaseSysRole mapped from table <base_sys_role>
type BaseSysRole struct {
	*cool.Model
	UserID    string  `gorm:"column:userId;type:varchar(255);not null" json:"userId"`                                             // 用户ID
	Name      string  `gorm:"column:name;type:varchar(255);not null;index:IDX_469d49a5998170e9550cf113da,priority:1" json:"name"` // 名称
	Label     *string `gorm:"column:label;type:varchar(50);index:IDX_f3f24fbbccf00192b076e549a7,priority:1" json:"label"`         // 角色标签
	Remark    *string `gorm:"column:remark;type:varchar(255)" json:"remark"`                                                      // 备注
	Relevance *int32  `gorm:"column:relevance;type:int;not null;default:1" json:"relevance"`                                      // 数据权限是否关联上下级
}

// TableName BaseSysRole's table name
func (*BaseSysRole) TableName() string {
	return TableNameBaseSysRole
}

// NewBaseSysRole create a new BaseSysRole
func NewBaseSysRole() *BaseSysRole {
	return &BaseSysRole{
		Model: cool.NewModel(),
	}
}

// init 创建表
func init() {
	cool.CreateTable(&BaseSysRole{})
}
