package model

import "github.com/cool-team-official/cool-admin-go/cool"

const TableNameBaseSysUserRole = "base_sys_user_role"

// BaseSysUserRole mapped from table <base_sys_user_role>
type BaseSysUserRole struct {
	*cool.Model
	UserID uint `gorm:"column:userId;type:bigint;not null" json:"userId"` // 用户ID
	RoleID uint `gorm:"column:roleId;type:bigint;not null" json:"roleId"` // 角色ID
}

// TableName BaseSysUserRole's table name
func (*BaseSysUserRole) TableName() string {
	return TableNameBaseSysUserRole
}

// NewBaseSysUserRole create a new BaseSysUserRole
func NewBaseSysUserRole() *BaseSysUserRole {
	return &BaseSysUserRole{
		Model: cool.NewModel(),
	}
}

// init 创建表
func init() {
	cool.CreateTable(&BaseSysUserRole{})
}
