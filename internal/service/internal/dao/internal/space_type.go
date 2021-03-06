// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpaceTypeDao is the data access object for table space_type.
type SpaceTypeDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns SpaceTypeColumns // columns contains all the column names of Table for convenient usage.
}

// SpaceTypeColumns defines and stores column names for table space_type.
type SpaceTypeColumns struct {
	Id         string // ID
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	Name       string // 类别名称
	ParentId   string // 父分类ID
}

//  spaceTypeColumns holds the columns for table space_type.
var spaceTypeColumns = SpaceTypeColumns{
	Id:         "id",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	Name:       "name",
	ParentId:   "parentId",
}

// NewSpaceTypeDao creates and returns a new DAO object for table data access.
func NewSpaceTypeDao() *SpaceTypeDao {
	return &SpaceTypeDao{
		group:   "default",
		table:   "space_type",
		columns: spaceTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SpaceTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SpaceTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SpaceTypeDao) Columns() SpaceTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SpaceTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SpaceTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SpaceTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
