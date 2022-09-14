package cool

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// Deprecated 请使用 cool.DBM 替代
func GDBM(m IModel) *gdb.Model {
	return g.DB(m.GroupName()).Model(m.TableName())
}

// DBM 根据model获取 *gdb.Model
func DBM(m IModel) *gdb.Model {
	return g.DB(m.GroupName()).Model(m.TableName())
}
