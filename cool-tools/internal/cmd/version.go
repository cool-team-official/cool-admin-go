package cmd

import (
	"context"

	"github.com/cool-team-official/cool-admin-go/cool-tools/internal/utility/mlog"
	"github.com/gogf/gf/v2/os/gbuild"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gutil"
)

type sVersion struct {
	Name     string //程序名称
	Homepage string //程序主页
	Version  string //程序版本
	GoFrame  string //goframe version
	Golang   string //golang version
	Git      string //git commit id
	Time     string //build datetime
}

var (
	Version = gcmd.Command{
		Name:  "version",
		Usage: "cool-tools version",
		Brief: "查看版本信息",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			info := gbuild.Info()
			binVersion := "v1.0.5"

			// 生成sVersion结构体
			res := sVersion{
				Name:     "cool-tools",
				Homepage: "https://cool-js.com",
				Version:  binVersion,
				GoFrame:  info.GoFrame,
				Golang:   info.Golang,
				Git:      info.Git,
				Time:     info.Time,
			}
			mlog.Printf(`CLI Installed At: %s`, gfile.SelfPath())
			gutil.Dump(res)
			return nil
		},
	}
)

// init 初始化模块
func init() {
	Main.AddCommand(&Version)
}
