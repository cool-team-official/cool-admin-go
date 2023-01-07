package cool

import (
	"github.com/cool-team-official/cool-admin-go/cool/cooldb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gres"
	"gorm.io/gorm"
)

// 初始化数据库连接供gorm使用
func InitDB(group string) (*gorm.DB, error) {
	// var ctx context.Context
	var db *gorm.DB
	// 如果group为空，则使用默认的group，否则使用group参数
	if group == "" {
		group = "default"
	}
	defer func() {
		if err := recover(); err != nil {
			panic("failed to connect database")
		}
	}()
	config := g.DB(group).GetConfig()
	db, err := cooldb.GetConn(config)
	if err != nil {
		panic(err.Error())
	}

	GormDBS[group] = db
	return db, nil
}

// 根据entity结构体获取 *gorm.DB
func getDBbyModel(model IModel) *gorm.DB {

	group := model.GroupName()
	// 判断是否存在 GormDBS[group] 字段，如果存在，则使用该字段的值作为DB，否则初始化DB
	if _, ok := GormDBS[group]; ok {
		return GormDBS[group]
	} else {

		db, err := InitDB(group)
		if err != nil {
			panic("failed to connect database")
		}
		// 把重新初始化的GormDBS存入全局变量中
		GormDBS[group] = db
		return db
	}
}

// 根据entity结构体创建表
func CreateTable(model IModel) error {
	if Config.AutoMigrate {
		db := getDBbyModel(model)
		return db.AutoMigrate(model)
	}
	return nil
}

// FillInitData 数据库填充初始数据
func FillInitData(ctx g.Ctx, moduleName string, model IModel) error {
	mInit := g.DB("default").Model("base_sys_init")
	n, err := mInit.Clone().Where("group", model.GroupName()).Where("table", model.TableName()).Count()
	if err != nil {
		g.Log().Error(ctx, "读取表 base_sys_init 失败 ", err.Error())
		return err
	}
	if n > 0 {
		g.Log().Debug(ctx, "分组", model.GroupName(), "中的表", model.TableName(), "已经初始化过,跳过本次初始化.")
		return nil
	}
	m := g.DB(model.GroupName()).Model(model.TableName())
	jsonData, _ := gjson.LoadContent(gres.GetContent("modules/" + moduleName + "/resource/initjson/" + model.TableName() + ".json"))
	if jsonData.Var().Clone().IsEmpty() {
		g.Log().Debug(ctx, "分组", model.GroupName(), "中的表", model.TableName(), "无可用的初始化数据,跳过本次初始化.")
		return nil
	}
	_, err = m.Data(jsonData).Insert()
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return err
	}
	_, err = mInit.Insert(g.Map{"group": model.GroupName(), "table": model.TableName()})
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return err
	}
	g.Log().Info(ctx, "分组", model.GroupName(), "中的表", model.TableName(), "初始化完成.")
	return nil
}
