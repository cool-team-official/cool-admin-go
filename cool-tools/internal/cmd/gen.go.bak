package cmd

import (
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Gen = cGen{}
)

type cGen struct {
	g.Meta `name:"gen" brief:"生成代码" description:"生成代码, 例如: cool-tools gen model"`
	cGenModel
}

func init() {
	Main.AddObject(Gen)
}
