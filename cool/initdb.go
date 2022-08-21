package cool

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gres"
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
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			LogLevel:                  logger.Error, // 日志级别
			IgnoreRecordNotFoundError: true,         // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,         // 彩色打印
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
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			LogLevel:                  logger.Error, // 日志级别
			IgnoreRecordNotFoundError: true,         // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,         // 彩色打印
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
func GetDBbyModel(model IModel) *gorm.DB {

	group := model.GroupName()
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
func CreateTable(model IModel) error {
	var ctx g.Ctx
	autoMigrate, _ := g.Cfg().Get(ctx, "database.autoMigrate")
	if autoMigrate.Bool() {
		g.Log().Debug(ctx, "start autoMigrate! database.autoMigrate", autoMigrate.Bool())
		g.Log().Debug(ctx, "开始在分组", model.GroupName(), "创建表", model.TableName())
		db := GetDBbyModel(model)
		return db.AutoMigrate(model)
	}
	g.Log().Info(ctx, "autoMigrate skiped! database.autoMigrate is ", autoMigrate.Bool())
	return nil
}

// 数据库填充初始数据
func FillInitData(moduleName string, model IModel) error {
	var ctx g.Ctx
	mInit := g.DB("default").Model("base_sys_init")
	n, err := mInit.Clone().Where("group", model.GroupName()).Where("table", model.TableName()).Count()
	if err != nil {
		g.Log().Error(ctx, "读取表 base_sys_init 失败 ", err)
		return err
	}
	if n > 0 {
		g.Log().Debug(ctx, "分组", model.GroupName(), "表", model.TableName(), "已经初始化过，跳过初始化")
		return nil
	}
	g.Log().Debug(ctx, "模块", moduleName, "将在分组", model.GroupName(), "表", model.TableName(), "中插入初始数据...")
	m := g.DB(model.GroupName()).Model(model.TableName())
	jsonData, _ := gjson.LoadContent(gres.GetContent("modules/" + moduleName + "/resource/initjson/" + model.TableName() + ".json"))
	if jsonData.Var().Clone().IsEmpty() {
		g.Log().Debug(ctx, "模块", moduleName, "没有初始化数据,跳过")
		return nil
	}
	m.Data(jsonData).Insert()
	mInit.Insert(g.Map{"group": model.GroupName(), "table": model.TableName()})
	return nil
}
