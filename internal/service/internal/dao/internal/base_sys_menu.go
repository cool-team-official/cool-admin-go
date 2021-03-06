// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BaseSysMenuDao is the data access object for table base_sys_menu.
type BaseSysMenuDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns BaseSysMenuColumns // columns contains all the column names of Table for convenient usage.
}

// BaseSysMenuColumns defines and stores column names for table base_sys_menu.
type BaseSysMenuColumns struct {
	Id         string // ID
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	ParentId   string // 父菜单ID
	Name       string // 菜单名称
	Router     string // 菜单地址
	Perms      string // 权限标识
	Type       string // 类型 0：目录 1：菜单 2：按钮
	Icon       string // 图标
	OrderNum   string // 排序
	ViewPath   string // 视图地址
	KeepAlive  string // 路由缓存
	IsShow     string // 父菜单名称
}

//  baseSysMenuColumns holds the columns for table base_sys_menu.
var baseSysMenuColumns = BaseSysMenuColumns{
	Id:         "id",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	ParentId:   "parentId",
	Name:       "name",
	Router:     "router",
	Perms:      "perms",
	Type:       "type",
	Icon:       "icon",
	OrderNum:   "orderNum",
	ViewPath:   "viewPath",
	KeepAlive:  "keepAlive",
	IsShow:     "isShow",
}

// NewBaseSysMenuDao creates and returns a new DAO object for table data access.
func NewBaseSysMenuDao() *BaseSysMenuDao {
	return &BaseSysMenuDao{
		group:   "default",
		table:   "base_sys_menu",
		columns: baseSysMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BaseSysMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BaseSysMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BaseSysMenuDao) Columns() BaseSysMenuColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BaseSysMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BaseSysMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BaseSysMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
