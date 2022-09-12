package cool

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GDBM 获取gf的gdb.Model对象
func GDBM(m IModel) *gdb.Model {
	return g.DB(m.GroupName()).Model(m.TableName())
}
