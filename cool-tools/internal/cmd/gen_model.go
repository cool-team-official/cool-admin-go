package cmd

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/v2/frame/g"
	"gorm.io/gen"
)

type cGenModel struct {
	g.Meta `name:"model" brief:"生成模型代码" description:"生成模型代码, 例如: cool-tools gen model"`
}

var (
	GenModel = cGenModel{}
)

type cGenModelInput struct {
	g.Meta   `name:"model" brief:"生成模型代码" description:"生成模型代码, 例如: cool-tools gen model"`
	Database string `v:"required#请输入数据库名称" arg:"true" name:"database" brief:"数据库名称" description:"数据库名称"`
}
type cGenModelOutput struct{}

func (c *cGenModel) Index(ctx g.Ctx, in cGenModelInput) (out cGenModelOutput, err error) {
	g.Log().Print(ctx, in.Database)
	// 获取数据库
	db, err := cool.InitDB(in.Database)
	if err != nil {
		panic(err.Error())
	}
	generator := gen.NewGenerator(gen.Config{
		OutPath:           "temp/gen/model",
		OutFile:           "",
		ModelPkgPath:      "",
		WithUnitTest:      false,
		FieldNullable:     true,
		FieldCoverable:    true,
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		Mode:              0,
	})
	generator.UseDB(db)
	generator.GenerateAllTable()
	generator.Execute()
	return
}
