package mysql

import (
	"fmt"
	"net/url"

	"github.com/cool-team-official/cool-admin-go/cool/cooldb"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DriverMysql struct {
}

func NewMysql() cooldb.Driver {
	return &DriverMysql{}
}

func (d *DriverMysql) GetConn(config *gdb.ConfigNode) (db *gorm.DB, err error) {
	var (
		source string
	)
	if config.Link != "" {
		// ============================================================================
		// Deprecated from v2.2.0.
		// ============================================================================
		source = config.Link
		// Custom changing the schema in runtime.
		if config.Name != "" {
			source, _ = gregex.ReplaceString(`/([\w\.\-]+)+`, "/"+config.Name, source)
		}
	} else {
		source = fmt.Sprintf(
			"%s:%s@%s(%s:%s)/%s?charset=%s",
			config.User, config.Pass, config.Protocol, config.Host, config.Port, config.Name, config.Charset,
		)
		if config.Timezone != "" {
			source = fmt.Sprintf("%s&loc=%s", source, url.QueryEscape(config.Timezone))
		}
		if config.Extra != "" {
			source = fmt.Sprintf("%s&%s", source, config.Extra)
		}
	}

	return gorm.Open(mysql.Open(source), &gorm.Config{})
}

func init() {
	var (
		err         error
		driverObj   = NewMysql()
		driverNames = g.SliceStr{"mysql", "mariadb", "tidb"}
	)
	for _, driverName := range driverNames {
		if err = cooldb.Register(driverName, driverObj); err != nil {
			panic(err)
		}
	}
}
