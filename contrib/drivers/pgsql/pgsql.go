package pgsql

import (
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"

	"github.com/cool-team-official/cool-admin-go/cool/cooldb"
	"github.com/gogf/gf/v2/frame/g"

	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DriverPgsql struct {
}

func NewDriverPgsql() *DriverPgsql {
	return &DriverPgsql{}
}

func (d *DriverPgsql) GetConn(config *gdb.ConfigNode) (db *gorm.DB, err error) {
	var (
		source string
		// underlyingDriverName = "postgres"
	)
	if config.Link != "" {
		// ============================================================================
		// Deprecated from v2.2.0.
		// ============================================================================
		source = config.Link
		// Custom changing the schema in runtime.
		if config.Name != "" {
			source, _ = gregex.ReplaceString(`dbname=([\w\.\-]+)+`, "dbname="+config.Name, source)
		}
	} else {
		if config.Name != "" {
			source = fmt.Sprintf(
				"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
				config.User, config.Pass, config.Host, config.Port, config.Name,
			)
		} else {
			source = fmt.Sprintf(
				"user=%s password=%s host=%s port=%s sslmode=disable",
				config.User, config.Pass, config.Host, config.Port,
			)
		}

		if config.Timezone != "" {
			source = fmt.Sprintf("%s timezone=%s", source, config.Timezone)
		}

		if config.Extra != "" {
			var extraMap map[string]interface{}
			if extraMap, err = gstr.Parse(config.Extra); err != nil {
				return nil, err
			}
			for k, v := range extraMap {
				source += fmt.Sprintf(` %s=%s`, k, v)
			}
		}
	}
	db, err = gorm.Open(postgres.Open(source), &gorm.Config{})
	return
}

func init() {
	// Register the driver.
	var (
		err         error
		driverObj   = NewDriverPgsql()
		driverNames = g.SliceStr{"pgsql"}
	)
	for _, driverName := range driverNames {
		if err = cooldb.Register(driverName, driverObj); err != nil {
			panic(err)
		}
	}
}
