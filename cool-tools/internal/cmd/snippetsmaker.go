package cmd

import (
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
)

var (
	SnippetsMaker = gcmd.Command{
		Name:  "snippetsmaker",
		Usage: "snippetsmaker",
		Brief: "snippetsmaker",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Debug(ctx, "snippetsmaker,生成工具^.^")
			files := garray.New(true)
			files.Append("modules/demo/model/demo_sample.go")
			files.Append("modules/demo/service/demo_sample.go")
			files.Append("modules/demo/controller/admin/demo_sample.go")
			// 遍历files
			for _, file := range files.Slice() {
				sArray := garray.NewStrArray()
				gfile.ReadLines(file.(string), func(line string) error {
					// g.Log().Debug(ctx, line)
					// println(line)
					// search := `Sample`
					// replace := `${TM_FILENAME_BASE/(.*)/${1:/capitalize}/}`
					// replaceArray := []string{"Sample", "${TM_FILENAME_BASE/(.*)/${1:/pascalcase}/}", "sample", "${TM_FILENAME_BASE/(.*)/${1:/downcase}/}", "demo", "${2:模块名称}", "app", "${TM_DIRECTORY/^.+[\\/\\\\]+(.*)$/$1/}"}
					replaceArray := []string{"DemoSample", "${TM_FILENAME_BASE/(.*)/${1:/pascalcase}/}", "demo_sample", "${TM_FILENAME_BASE/(.*)/${1:/downcase}/}"}

					result := gstr.ReplaceByArray(line, replaceArray)
					sArray.Append(gstr.AddSlashes(result))

					return nil
				})
				// g.Dump(sArray)
				println(file.(string))
				println("--------------------------------------code start------------------------------------------")
				println(`"body":[`)
				sArray.Iterator(
					func(index int, value string) bool {
						println("\"" + value + "\",")
						return true
					},
				)
				println("]")
				println("--------------------------------------code end------------------------------------------")
			}
			return nil
		},
	}
)

func init() {
	Main.AddCommand(&SnippetsMaker)
}
