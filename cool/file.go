package cool

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
)

type sFile struct {
	Mode   string `json:"mode"`
	Domain string `json:"domain"`
}

// Upload 上传文件
func (f *sFile) Upload(ctx g.Ctx) (string, error) {
	var (
		err     error
		Request = g.RequestFromCtx(ctx)
	)

	if f.Mode == "local" {
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
		return f.Domain + "/public/uploads/" + dir + "/" + fileName, err
	}
	return "", err
}

// GetMode 获取上传模式
func (f *sFile) GetMode() (data interface{}, err error) {
	data = g.MapStrStr{
		"mode": f.Mode,
		"type": "local",
	}
	return
}

// NewFile new file
func NewFile() *sFile {
	file := &sFile{
		Mode:   Config.File.Mode,
		Domain: Config.File.Domain,
	}
	if file.Mode == "local" {
		var ctx g.Ctx
		s := g.Server()
		if !gfile.Exists("./public/uploads") {
			err := gfile.Mkdir("./public/uploads")
			if err != nil {
				s.Logger().Error(ctx, err)
			}
		}
		s.AddStaticPath("/public", "./public")

	}
	return file
}

// File file
var File = NewFile()
