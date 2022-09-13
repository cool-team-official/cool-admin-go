package sqlite

import (
	"github.com/cool-team-official/cool-admin-go/cool/cooldb"
	"github.com/glebarez/sqlite"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gorm"
)

type DriverSqlite struct {
}

func New() cooldb.Driver {
	return &DriverSqlite{}
}

func (d *DriverSqlite) GetConn(node *gdb.ConfigNode) (db *gorm.DB, err error) {
	return gorm.Open(sqlite.Open(node.Link), &gorm.Config{})
}

func init() {
	var (
		err         error
		driverObj   = New()
		driverNames = g.SliceStr{"sqlite"}
	)
	for _, driverName := range driverNames {
		if err = cooldb.Register(driverName, driverObj); err != nil {
			panic(err)
		}
	}
}
