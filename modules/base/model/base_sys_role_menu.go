package model

import "github.com/cool-team-official/cool-admin-go/cool"

const TableNameBaseSysRoleMenu = "base_sys_role_menu"

// BaseSysRoleMenu mapped from table <base_sys_role_menu>
type BaseSysRoleMenu struct {
	*cool.Model
	RoleID int64 `gorm:"column:roleId;type:bigint;not null" json:"roleId"` // 角色ID
	MenuID int64 `gorm:"column:menuId;type:bigint;not null" json:"menuId"` // 菜单ID
}

// TableName BaseSysRoleMenu's table name
func (*BaseSysRoleMenu) TableName() string {
	return TableNameBaseSysRoleMenu
}

// init 创建表
func init() {
	cool.CreateTable(&BaseSysRoleMenu{})
}
