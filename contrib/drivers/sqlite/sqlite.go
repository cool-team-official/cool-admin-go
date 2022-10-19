package sqlite

import (
	"fmt"

	"github.com/cool-team-official/cool-admin-go/cool/cooldb"
	"github.com/glebarez/sqlite"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"gorm.io/gorm"
)

type DriverSqlite struct {
}

func NewSqlite() cooldb.Driver {
	return &DriverSqlite{}
}

func (d *DriverSqlite) GetConn(config *gdb.ConfigNode) (db *gorm.DB, err error) {
	var (
		source string
	)
	if config.Link != "" {
		source = config.Link
	} else {
		source = config.Name
	}
	// It searches the source file to locate its absolute path..
	if absolutePath, _ := gfile.Search(source); absolutePath != "" {
		source = absolutePath
	}
	// Multiple PRAGMAs can be specified, e.g.:
	// path/to/some.db?_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)
	if config.Extra != "" {
		var (
			options  string
			extraMap map[string]interface{}
		)
		if extraMap, err = gstr.Parse(config.Extra); err != nil {
			return nil, err
		}
		for k, v := range extraMap {
			if options != "" {
				options += "&"
			}
			options += fmt.Sprintf(`_pragma=%s(%s)`, k, gurl.Encode(gconv.String(v)))
		}
		if len(options) > 1 {
			source += "?" + options
		}
	}
	println("Will use", source, "to open DB")
	return gorm.Open(sqlite.Open(source), &gorm.Config{})
}

func init() {
	var (
		err         error
		driverObj   = NewSqlite()
		driverNames = g.SliceStr{"sqlite"}
	)
	for _, driverName := range driverNames {
		if err = cooldb.Register(driverName, driverObj); err != nil {
			panic(err)
		}
	}
}
