package cool

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gmeta"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 初始化数据库连接供gorm使用
func initDB(group string) (*gorm.DB, error) {
	var ctx context.Context
	var db *gorm.DB
	// 如果group为空，则使用默认的group，否则使用group参数
	if group == "" {
		group = "default"
	}
	var configMap map[string]interface{}
	var configSlice []interface{}
	if v, _ := g.Cfg().Get(ctx, "database."+group); !v.IsEmpty() {
		if v.IsSlice() {
			// g.Log().Debug(ctx, "v.IsSlice()")
			configSlice = v.Slice()
			configMap = configSlice[0].(map[string]interface{})
		} else if v.IsMap() {
			// g.Log().Debug(ctx, "v.IsMap()")
			configMap = v.Map()
		} else {
			panic("无法解析数据库配置")
		}
	}
	// g.Dump(configMap)

	switch configMap["type"] {
	case "sqlite":
		// g.Log().Debug(ctx, "sqlite")
		db, _ = sqliteInit(configMap["link"].(string))
	case "mysql":
		// g.Log().Debug(ctx, "mysql")
		db, _ = mysqlInit(configMap["link"].(string))
	default:
		g.Log().Error(ctx, configMap["type"], "为未知数据库类型")
		panic("cooldatabase type not found")
	}
	GromDBS[group] = db
	return db, nil
}

// 初始化sqlite类型连接
func sqliteInit(link string) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 彩色打印
		},
	)
	db, err := gorm.Open(sqlite.Open(link), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db, nil
}

// 初始化mysql类型连接
func mysqlInit(link string) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(link), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db, nil
}

// 根据entity结构体获取DB
func GetDBbyEntity(entity interface{}) *gorm.DB {
	meta := gmeta.Data(entity)
	// 判断是否存在 meta["group"] 字段，如果存在，则使用该字段的值作为group，否则使用默认的group
	var group string
	if _, ok := meta["group"]; ok {
		group = meta["group"]
	} else {
		group = "default"
	}
	// 判断是否存在 GromDBS[group] 字段，如果存在，则使用该字段的值作为DB，否则初始化DB
	if _, ok := GromDBS[group]; ok {
		return GromDBS[group]
	} else {

		db, err := initDB(group)
		if err != nil {
			panic("failed to connect database")
		}
		// 把重新初始化的GromDBS存入全局变量中
		GromDBS[group] = db
		return db
	}
}

// 根据entity结构体创建表
func CreateTable(entity interface{}) error {
	var ctx g.Ctx
	autoMigrate, _ := g.Cfg().Get(ctx, "database.autoMigrate")
	if autoMigrate.Bool() {
		g.Log().Info(ctx, "start autoMigrate! database.autoMigrate", autoMigrate.Bool())
		db := GetDBbyEntity(entity)
		return db.AutoMigrate(entity)
	}
	g.Log().Info(ctx, "autoMigrate skiped! database.autoMigrate is ", autoMigrate.Bool())
	return nil
}
