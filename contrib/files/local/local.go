// Package local 提供本地文件上传支持
package local

import (
	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/cool-team-official/cool-admin-go/cool/coolfile"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
)

type Local struct {
}

func (l *Local) Upload(ctx g.Ctx) (string, error) {
	var (
		err     error
		Request = g.RequestFromCtx(ctx)
	)

	file := Request.GetUploadFile("file")
	if file == nil {
		return "", gerror.New("上传文件为空")
	}
	// 以当前年月日为目录
	dir := gtime.Now().Format("Ymd")

	fileName, err := file.Save("./public/uploads/"+dir, true)
	if err != nil {
		return "", err
	}
	return cool.Config.File.Domain + "/public/uploads/" + dir + "/" + fileName, err
}

func (l *Local) GetMode() (data interface{}, err error) {
	data = g.MapStrStr{
		"mode": cool.Config.File.Mode,
		"type": "local",
	}
	return
}

func (l *Local) New() coolfile.Driver {
	return &Local{}
}

func New() coolfile.Driver {
	return &Local{}
}

func init() {
	var (
		err         error
		driverObj   = New()
		driverNames = g.SliceStr{"local"}
	)
	for _, driverName := range driverNames {
		if err = coolfile.Register(driverName, driverObj); err != nil {
			panic(err)
		}
	}
	s := g.Server()
	if !gfile.Exists("./public/uploads") {
		err := gfile.Mkdir("./public/uploads")
		if err != nil {
			panic(err)
		}
	}
	s.AddStaticPath("/public", "./public")
}
