package service

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/text/gstr"
)

func CreatModule(ctx context.Context, moduleName string) (err error) {
	// 检测当前目录是否存在go.mod文件
	if !gfile.Exists("go.mod") {
		err = gerror.New("当前目录不存在go.mod文件,请在项目根目录下执行")
		return
	}
	module := ""
	// 读取go.mod文件第一行文本
	gfile.ReadLines("go.mod", func(text string) error {
		if gstr.Contains(text, "module") && module == "" {
			// println("module:", text)
			module = gstr.StrEx(text, "module")
			// println("module:", module)
			module = gstr.TrimAll(module)
			// println("module:", module)

			return nil
		}
		return nil
	})
	if module == "" {
		err = gerror.New("go.mod文件中不存在module行")
		return
	}
	// println(module)
	// 创建模块目录
	moduleDir := gfile.Join(gfile.Pwd(), "modules", moduleName)
	if gfile.Exists(moduleDir) {
		err = gerror.New("模块已经存在,请先删除原有模块")
		return
	}
	err = gfile.Mkdir(moduleDir)
	if err != nil {
		return
	}
	// 创建模块目录结构
	err = gres.Export("cool-admin-go-simple/modules/demo", moduleDir, gres.ExportOption{
		RemovePrefix: "cool-admin-go-simple/modules/demo",
	})
	if err != nil {
		return
	}
	// 替换import路径
	err = gfile.ReplaceDir("cool-admin-go-simple/modules/demo", module+"/modules/"+moduleName, moduleDir, "*", true)
	if err != nil {
		return
	}

	// 重命名demo.go 为 moduleName.go
	err = gfile.Rename(gfile.Join(moduleDir, "demo.go"), gfile.Join(moduleDir, moduleName+".go"))
	if err != nil {
		return
	}
	println("创建模块成功:", moduleDir)
	return nil
}
