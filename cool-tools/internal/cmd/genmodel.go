package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var (
	GenModel1 = gcmd.Command{
		Name:  "genmodel",
		Usage: "genmodel",
		Brief: "根据数据库表生成model",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			genmodel()
			return
		},
	}
)

// init 初始化模块
func init() {
	Main.AddCommand(&GenModel1)
}
func genmodel() {
	println("genmodel.............................")
	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	gene := gen.NewGenerator(gen.Config{
		OutPath: "./temp/gen/model",
		/* Mode: gen.WithoutContext|gen.WithDefaultQuery*/
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		/* FieldNullable: true,*/
		FieldNullable: true,
		//if you want to assign field which has default value in `Create` API, set FieldCoverable true, reference: https://gorm.io/docs/create.html#Default-Values
		/* FieldCoverable: true,*/
		FieldCoverable: true,
		// if you want generate field with unsigned integer type, set FieldSignable true
		/* FieldSignable: true,*/
		FieldSignable: false,
		//if you want to generate index tags from database, set FieldWithIndexTag true
		/* FieldWithIndexTag: true,*/
		FieldWithIndexTag: true,
		//if you want to generate type tags from database, set FieldWithTypeTag true
		/* FieldWithTypeTag: true,*/
		FieldWithTypeTag: true,
		//if you need unit tests for query code, set WithUnitTest true
		/* WithUnitTest: true, */
		WithUnitTest: false,
		Mode:         gen.WithoutContext,
	})

	// reuse the database connection in Project or create a connection here
	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessary or it will panic
	db, _ := gorm.Open(mysql.Open("root:123456@(127.0.0.1:3306)/cool?charset=utf8mb4&parseTime=True&loc=Local"))
	gene.UseDB(db)

	// apply basic crud api on structs or table models which is specified by table name with function
	// GenerateModel/GenerateModelAs. And generator will generate table models' code when calling Excute.
	// 想对已有的model生成crud等基础方法可以直接指定model struct ，例如model.User{}
	// 如果是想直接生成表的model和crud方法，则可以指定表的名称，例如g.GenerateModel("company")
	// 想自定义某个表生成特性，比如struct的名称/字段类型/tag等，可以指定opt，例如g.GenerateModel("company",gen.FieldIgnore("address")), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address"))
	// g.ApplyBasic(model.User{}, g.GenerateModel("company"), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address")))
	// g.GenerateModel("base_sys_user", gen.FieldIgnore("password"))

	// type Method interface {
	// 	// return "default"
	// 	GroupName() string
	// }
	// gene.ApplyInterface(func(method Method) {}, gene.GenerateModel("base_sys_user"))
	gene.GenerateAllTable(gen.FieldIgnore("id"))

	// apply diy interfaces on structs or table models
	// 如果想给某些表或者model生成自定义方法，可以用ApplyInterface，第一个参数是方法接口，可以参考DIY部分文档定义
	// g.ApplyInterface(func(method model.Method) {}, model.User{}, g.GenerateModel("company"))

	// execute the action of code generation
	gene.Execute()
}
