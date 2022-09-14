package mssql

import (
	_ "github.com/gogf/gf/contrib/drivers/mssql/v2"

	"github.com/cool-team-official/cool-admin-go/cool/cooldb"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DriverMssql struct {
}

func New() cooldb.Driver {
	return &DriverMssql{}
}

func (d *DriverMssql) GetConn(node *gdb.ConfigNode) (db *gorm.DB, err error) {
	return gorm.Open(sqlserver.Open(node.Link), &gorm.Config{})
}

func init() {
	var (
		err         error
		driverObj   = New()
		driverNames = g.SliceStr{"mssql"}
	)
	for _, driverName := range driverNames {
		if err = cooldb.Register(driverName, driverObj); err != nil {
			panic(err)
		}
	}
}
