package cooldb

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"gorm.io/gorm"
)

// Driver 数据库驱动通用接口
type Driver interface {
	GetConn(node *gdb.ConfigNode) (db *gorm.DB, err error) // 获取数据库连接
}

var (
	driverMap = map[string]Driver{} // driverMap is the map for registered database drivers.
)

// GetConn returns the connection object of specified database driver.
func GetConn(node *gdb.ConfigNode) (db *gorm.DB, err error) {
	if driver, ok := driverMap[node.Type]; ok {
		return driver.GetConn(node)
	}
	errorMsg := "\n"
	errorMsg += `cannot find database driver for specified database type "%s"`
	errorMsg += `, did you misspell type name "%s" or forget importing the database driver? `
	errorMsg += `possible reference: https://github.com/cool-team-official/cool-admin-go/tree/master/contrib/drivers`
	// 换行
	errorMsg += "\n"
	errorMsg += `无法找到指定数据库类型的数据库驱动 "%s"`
	errorMsg += `，您是否拼写错误了类型名称 "%s" 或者忘记导入数据库驱动？`
	errorMsg += `参考:https://github.com/cool-team-official/cool-admin-go/contrib/drivers`
	err = gerror.Newf(errorMsg, node.Type, node.Type, node.Type, node.Type)
	return
}

// Register registers custom database driver to gdb.
func Register(name string, driver Driver) error {
	driverMap[name] = driver
	return nil
}
