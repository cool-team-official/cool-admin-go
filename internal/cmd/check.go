package cmd

import (
	"context"

	"github.com/gogf/gf/v2/debug/gdebug"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/frame/gins"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Check = gcmd.Command{
		Name:  "check",
		Usage: "check",
		Brief: "check",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Debug(ctx, "Check ～～～～～～～~~~~~~~~~~")
			_, err = g.Redis().Do(ctx, "HSET", "test", "id", 10000)

			if err != nil {
				panic(err)
			}
			// test := &gredis.Config{
			// 	Address: "127.0.0.1:6379",
			// 	Db:      0,
			// }
			redis := gins.Redis()
			redis.Do(ctx, "HSET", "test2", "id", 10000)

			println(gdebug.Caller())
			g.Redis()
			return
		},
	}
)

// init 初始化模块
func init() {
	Main.AddCommand(&Check)
}
