package cool

import (
	"time"

	"github.com/cool-team-official/cool-admin-go/cool/coolfunc"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

type CoolFunc coolfunc.CoolFunc

// FuncMap 函数列表
var FuncMap = make(map[string]CoolFunc)

// RegisterFunc 注册函数
func RegisterFunc(name string, f CoolFunc) {
	FuncMap[name] = f
}

// GetFunc 获取函数
func GetFunc(name string) CoolFunc {
	return FuncMap[name]
}

// RunFunc 运行函数
func RunFunc(ctx g.Ctx, funcstring string) (err error) {
	funcName := gstr.SubStr(funcstring, 0, gstr.Pos(funcstring, "("))
	funcParam := gstr.SubStr(funcstring, gstr.Pos(funcstring, "(")+1, gstr.Pos(funcstring, ")")-gstr.Pos(funcstring, "(")-1)
	if _, ok := FuncMap[funcName]; !ok {
		err = gerror.New("函数不存在:" + funcName)
		return
	}
	if !FuncMap[funcName].IsAllWorker() {
		// 检查当前是否为主进程, 如果不是主进程, 则不执行
		if ProcessFlag != CacheManager.MustGetOrSet(ctx, "cool:masterflag", ProcessFlag, 60*time.Second).String() {
			g.Log().Debug(ctx, "当前进程不是主进程, 不执行单例函数", funcName)
			return
		}
	}
	err = FuncMap[funcName].Func(ctx, funcParam)
	return
}

// ClusterRunFunc 集群运行函数,如果是单机模式, 则直接运行函数
func ClusterRunFunc(ctx g.Ctx, funcstring string) (err error) {
	if IsRedisMode {
		conn, err := g.Redis("cool").Conn(ctx)
		if err != nil {
			return err
		}
		defer conn.Close(ctx)
		_, err = conn.Do(ctx, "publish", "cool:func", funcstring)
		return err
	} else {
		return RunFunc(ctx, funcstring)
	}
}

// ListenFunc 监听函数
func ListenFunc(ctx g.Ctx) {
	if IsRedisMode {
		conn, err := g.Redis("cool").Conn(ctx)
		if err != nil {
			panic(err)
		}
		defer conn.Close(ctx)
		_, err = conn.Do(ctx, "subscribe", "cool:func")
		if err != nil {
			panic(err)
		}
		for {
			data, err := conn.Receive(ctx)
			if err != nil {
				g.Log().Error(ctx, err)
				time.Sleep(10 * time.Second)
				continue
			}
			if data != nil {
				dataMap := data.MapStrStr()
				if dataMap["Kind"] == "subscribe" {
					continue
				}
				if dataMap["Channel"] == "cool:func" {
					g.Log().Debug(ctx, "执行函数", dataMap["Payload"])
					err := RunFunc(ctx, dataMap["Payload"])
					if err != nil {
						g.Log().Error(ctx, "执行函数失败", err)
					}
				}
			}
		}
	} else {
		panic(gerror.New("集群模式下, 请使用Redis作为缓存"))
	}
}
