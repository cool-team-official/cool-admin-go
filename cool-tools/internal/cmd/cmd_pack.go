package cmd

import (
	"context"
	"strings"

	"github.com/cool-team-official/cool-admin-go/cool-tools/internal/utility/allyes"
	"github.com/cool-team-official/cool-admin-go/cool-tools/internal/utility/mlog"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/util/gtag"
)

var (
	Pack = cPack{}
)

type cPack struct {
	g.Meta `name:"pack" usage:"{cPackUsage}" brief:"{cPackBrief}" eg:"{cPackEg}"`
}

const (
	cPackUsage = `cool-tools pack SRC DST`
	cPackBrief = `packing any file/directory to a resource file, or a go file`
	cPackEg    = `
cool-tools pack public data.bin
cool-tools pack public,template data.bin
cool-tools pack public,template packed/data.go
cool-tools pack public,template,config packed/data.go
cool-tools pack public,template,config packed/data.go -n=packed -p=/var/www/my-app
cool-tools pack /var/www/public packed/data.go -n=packed
`
	cPackSrcBrief = `source path for packing, which can be multiple source paths.`
	cPackDstBrief = `
destination file path for packed file. if extension of the filename is ".go" and "-n" option is given, 
it enables packing SRC to go file, or else it packs SRC into a binary file.
`
	cPackNameBrief     = `package name for output go file, it's set as its directory name if no name passed`
	cPackPrefixBrief   = `prefix for each file packed into the resource file`
	cPackKeepPathBrief = `keep the source path from system to resource file, usually for relative path`
)

func init() {
	gtag.Sets(g.MapStrStr{
		`cPackUsage`:         cPackUsage,
		`cPackBrief`:         cPackBrief,
		`cPackEg`:            cPackEg,
		`cPackSrcBrief`:      cPackSrcBrief,
		`cPackDstBrief`:      cPackDstBrief,
		`cPackNameBrief`:     cPackNameBrief,
		`cPackPrefixBrief`:   cPackPrefixBrief,
		`cPackKeepPathBrief`: cPackKeepPathBrief,
	})
	Main.AddObject(Pack)

}

type cPackInput struct {
	g.Meta   `name:"pack"`
	Src      string `name:"SRC" arg:"true" v:"required" brief:"{cPackSrcBrief}"`
	Dst      string `name:"DST" arg:"true" v:"required" brief:"{cPackDstBrief}"`
	Name     string `name:"name"     short:"n" brief:"{cPackNameBrief}"`
	Prefix   string `name:"prefix"   short:"p" brief:"{cPackPrefixBrief}"`
	KeepPath bool   `name:"keepPath" short:"k" brief:"{cPackKeepPathBrief}" orphan:"true"`
}

type cPackOutput struct{}

func (c cPack) Index(ctx context.Context, in cPackInput) (out *cPackOutput, err error) {
	if gfile.Exists(in.Dst) && gfile.IsDir(in.Dst) {
		mlog.Fatalf("DST path '%s' cannot be a directory", in.Dst)
	}
	if !gfile.IsEmpty(in.Dst) && !allyes.Check() {
		s := gcmd.Scanf("path '%s' is not empty, files might be overwrote, continue? [y/n]: ", in.Dst)
		if strings.EqualFold(s, "n") {
			return
		}
	}
	if in.Name == "" && gfile.ExtName(in.Dst) == "go" {
		in.Name = gfile.Basename(gfile.Dir(in.Dst))
	}
	var option = gres.Option{
		Prefix:   in.Prefix,
		KeepPath: in.KeepPath,
	}
	if in.Name != "" {
		if err = gres.PackToGoFileWithOption(in.Src, in.Dst, in.Name, option); err != nil {
			mlog.Fatalf("pack failed: %v", err)
		}
	} else {
		if err = gres.PackToFileWithOption(in.Src, in.Dst, option); err != nil {
			mlog.Fatalf("pack failed: %v", err)
		}
	}
	mlog.Print("done!")
	return
}
