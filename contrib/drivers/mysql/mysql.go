package mysql

import (
	"github.com/cool-team-official/cool-admin-go/cool/cooldb"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DriverMysql struct {
}

func New() cooldb.Driver {
	return &DriverMysql{}
}

func (d *DriverMysql) GetConn(node *gdb.ConfigNode) (db *gorm.DB, err error) {
	return gorm.Open(mysql.Open(node.Link), &gorm.Config{})
}

func init() {
	var (
		err         error
		driverObj   = New()
		driverNames = g.SliceStr{"mysql", "mariadb", "tidb"}
	)
	for _, driverName := range driverNames {
		if err = cooldb.Register(driverName, driverObj); err != nil {
			panic(err)
		}
	}
}
