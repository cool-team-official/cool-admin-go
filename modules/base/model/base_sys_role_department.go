package model

import "github.com/cool-team-official/cool-admin-go/cool"

const TableNameBaseSysRoleDepartment = "base_sys_role_department"

// BaseSysRoleDepartment mapped from table <base_sys_role_department>
type BaseSysRoleDepartment struct {
	*cool.Model
	RoleID       int64 `gorm:"column:roleId;type:bigint;not null" json:"roleId"`             // 角色ID
	DepartmentID int64 `gorm:"column:departmentId;type:bigint;not null" json:"departmentId"` // 部门ID
}

// TableName BaseSysRoleDepartment's table name
func (*BaseSysRoleDepartment) TableName() string {
	return TableNameBaseSysRoleDepartment
}

// NewBaseSysRoleDepartment create a new BaseSysRoleDepartment
func NewBaseSysRoleDepartment() *BaseSysRoleDepartment {
	return &BaseSysRoleDepartment{
		Model: cool.NewModel(),
	}
}

// init 创建表
func init() {
	cool.CreateTable(&BaseSysRoleDepartment{})
}
